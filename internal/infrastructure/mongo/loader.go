package mongo

import (
	"context"
	"time"
)

const (
	dataLoaderWait     = 1 * time.Millisecond
	dataLoaderMaxBatch = 100
)

type DataLoaders struct {
	Layer LayerDataLoader
}

type NewDataloaderKey struct {}
func DataLoadersFromContext(ctx context.Context) *LayerDataLoader {
	return ctx.Value(NewDataloaderKey{}).(*LayerDataLoader)
}

func DataLoadersKey() interface{} {
	return NewDataloaderKey{}
}
