package mongo

import (
	"context"
	"github.com/reearth/reearth-backend/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/internal/usecase/interactor"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
)

type LayerMongoLoader struct {
	interactor interactor.Layer
	repo repo.Layer
}

func NewLayerMongoLoader(interactor interactor.Layer, repo repo.Layer) *LayerMongoLoader {
	return &LayerMongoLoader{interactor: interactor, repo: repo}
}

func (c *LayerMongoLoader) FindItemByID(ctx context.Context, ids []id.LayerID) ([]*layer.Item, []error) {
	scene, err := c.interactor.OnlyReadableScenes(ctx, getOperator(ctx))
	if err != nil {
		return nil,[]error{err}
	}
	res, err := c.repo.FindItemByIDs(ctx, keys, scene)
	if err != nil {
		return nil,[]error{err}
	}

	layers := make([]*layer.Item,0,len(res))
	for _, l := range res {
		if l == nil {
			layers = append(layers, nil)
		} else {
			layers = append(layers, l)
		}
	}
	return layers, nil
}

type LayerMongoDataLoader interface {
	Load(id.LayerID) (*layer.Item, error)
	LoadAll([]id.LayerID) ([]*layer.Item, []error)
}

func (c *LayerMongoLoader) DataLoader(ctx context.Context) LayerMongoDataLoader {
	return gqldataloader.NewLayerMongoLoader(gqldataloader.LayerMongoLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []id.LayerID) ([]*layer.Item, []error) {
			return c.FindItemByID(ctx, keys)
		},
	})
}