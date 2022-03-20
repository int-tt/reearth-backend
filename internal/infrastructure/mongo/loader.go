package mongo

import (
	"context"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"time"
)

const (
	dataLoaderWait     = 500 * time.Nanosecond
	dataLoaderMaxBatch = 100
)

type DataLoaders struct {
	Layer    *LayerDataLoader
	Property *PropertyDataLoader
}

type NewDataloaderKey struct{}

func DataLoadersFromContext(ctx context.Context) *DataLoaders {
	return ctx.Value(NewDataloaderKey{}).(*DataLoaders)
}
func NewDataLoaders(ctx context.Context, container *repo.Container) *DataLoaders {
	d := &DataLoaders{
		Layer:    NewLayerDataloader(ctx, container.Layer),
		Property: NewPropertyDataLoader(ctx, container.Property),
	}
	return d
}
func DataLoadersKey() interface{} {
	return NewDataloaderKey{}
}
