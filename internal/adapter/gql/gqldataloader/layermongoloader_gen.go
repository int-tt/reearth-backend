// Code generated by github.com/vektah/dataloaden, DO NOT EDIT.

package gqldataloader

import (
	"sync"
	"time"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
)

// LayerMongoLoaderConfig captures the config to create a new LayerMongoLoader
type LayerMongoLoaderConfig struct {
	// Fetch is a method that provides the data for the loader
	Fetch func(keys []id.LayerID) ([]*layer.Item, []error)

	// Wait is how long wait before sending a batch
	Wait time.Duration

	// MaxBatch will limit the maximum number of keys to send in one batch, 0 = not limit
	MaxBatch int
}

// NewLayerMongoLoader creates a new LayerMongoLoader given a fetch, wait, and maxBatch
func NewLayerMongoLoader(config LayerMongoLoaderConfig) *LayerMongoLoader {
	return &LayerMongoLoader{
		fetch:    config.Fetch,
		wait:     config.Wait,
		maxBatch: config.MaxBatch,
	}
}

// LayerMongoLoader batches and caches requests
type LayerMongoLoader struct {
	// this method provides the data for the loader
	fetch func(keys []id.LayerID) ([]*layer.Item, []error)

	// how long to done before sending a batch
	wait time.Duration

	// this will limit the maximum number of keys to send in one batch, 0 = no limit
	maxBatch int

	// INTERNAL

	// lazily created cache
	cache map[id.LayerID]*layer.Item

	// the current batch. keys will continue to be collected until timeout is hit,
	// then everything will be sent to the fetch method and out to the listeners
	batch *layerMongoLoaderBatch

	// mutex to prevent races
	mu sync.Mutex
}

type layerMongoLoaderBatch struct {
	keys    []id.LayerID
	data    []*layer.Item
	error   []error
	closing bool
	done    chan struct{}
}

// Load a Item by key, batching and caching will be applied automatically
func (l *LayerMongoLoader) Load(key id.LayerID) (*layer.Item, error) {
	return l.LoadThunk(key)()
}

// LoadThunk returns a function that when called will block waiting for a Item.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *LayerMongoLoader) LoadThunk(key id.LayerID) func() (*layer.Item, error) {
	l.mu.Lock()
	if it, ok := l.cache[key]; ok {
		l.mu.Unlock()
		return func() (*layer.Item, error) {
			return it, nil
		}
	}
	if l.batch == nil {
		l.batch = &layerMongoLoaderBatch{done: make(chan struct{})}
	}
	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() (*layer.Item, error) {
		<-batch.done

		var data *layer.Item
		if pos < len(batch.data) {
			data = batch.data[pos]
		}

		var err error
		// its convenient to be able to return a single error for everything
		if len(batch.error) == 1 {
			err = batch.error[0]
		} else if batch.error != nil {
			err = batch.error[pos]
		}

		if err == nil {
			l.mu.Lock()
			l.unsafeSet(key, data)
			l.mu.Unlock()
		}

		return data, err
	}
}

// LoadAll fetches many keys at once. It will be broken into appropriate sized
// sub batches depending on how the loader is configured
func (l *LayerMongoLoader) LoadAll(keys []id.LayerID) ([]*layer.Item, []error) {
	results := make([]func() (*layer.Item, error), len(keys))

	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}

	items := make([]*layer.Item, len(keys))
	errors := make([]error, len(keys))
	for i, thunk := range results {
		items[i], errors[i] = thunk()
	}
	return items, errors
}

// LoadAllThunk returns a function that when called will block waiting for a Items.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *LayerMongoLoader) LoadAllThunk(keys []id.LayerID) func() ([]*layer.Item, []error) {
	results := make([]func() (*layer.Item, error), len(keys))
	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}
	return func() ([]*layer.Item, []error) {
		items := make([]*layer.Item, len(keys))
		errors := make([]error, len(keys))
		for i, thunk := range results {
			items[i], errors[i] = thunk()
		}
		return items, errors
	}
}

// Prime the cache with the provided key and value. If the key already exists, no change is made
// and false is returned.
// (To forcefully prime the cache, clear the key first with loader.clear(key).prime(key, value).)
func (l *LayerMongoLoader) Prime(key id.LayerID, value *layer.Item) bool {
	l.mu.Lock()
	var found bool
	if _, found = l.cache[key]; !found {
		// make a copy when writing to the cache, its easy to pass a pointer in from a loop var
		// and end up with the whole cache pointing to the same value.
		cpy := *value
		l.unsafeSet(key, &cpy)
	}
	l.mu.Unlock()
	return !found
}

// Clear the value at key from the cache, if it exists
func (l *LayerMongoLoader) Clear(key id.LayerID) {
	l.mu.Lock()
	delete(l.cache, key)
	l.mu.Unlock()
}

func (l *LayerMongoLoader) unsafeSet(key id.LayerID, value *layer.Item) {
	if l.cache == nil {
		l.cache = map[id.LayerID]*layer.Item{}
	}
	l.cache[key] = value
}

// keyIndex will return the location of the key in the batch, if its not found
// it will add the key to the batch
func (b *layerMongoLoaderBatch) keyIndex(l *LayerMongoLoader, key id.LayerID) int {
	for i, existingKey := range b.keys {
		if key == existingKey {
			return i
		}
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(l)
	}

	if l.maxBatch != 0 && pos >= l.maxBatch-1 {
		if !b.closing {
			b.closing = true
			l.batch = nil
			go b.end(l)
		}
	}

	return pos
}

func (b *layerMongoLoaderBatch) startTimer(l *LayerMongoLoader) {
	time.Sleep(l.wait)
	l.mu.Lock()

	// we must have hit a batch limit and are already finalizing this batch
	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(l)
}

func (b *layerMongoLoaderBatch) end(l *LayerMongoLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}
