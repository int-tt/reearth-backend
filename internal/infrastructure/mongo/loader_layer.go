//go:generate go run github.com/vektah/dataloaden LayerLoader *github.com/reearth/reearth-backend/pkg/id.LayerID github.com/reearth/reearth-backend/pkg/layer.Layer

package mongo

import (
	"context"
	"fmt"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
)

type LayerDataLoader struct {
	Layer  repo.Layer
	loader *LayerLoader
}

func (l *LayerDataLoader) Filtered(filter repo.SceneFilter) repo.Layer {
	return &LayerDataLoader{Layer: l.Layer.Filtered(filter)}
}

func (l *LayerDataLoader) FindParentsByIDs(ctx context.Context, ids []id.LayerID) (layer.GroupList, error) {
	return l.Layer.FindParentsByIDs(ctx, ids)
}

func (l *LayerDataLoader) FindByPluginAndExtension(ctx context.Context, pluginID id.PluginID, extensionID *id.PluginExtensionID) (layer.List, error) {
	return l.Layer.FindByPluginAndExtension(ctx, pluginID, extensionID)
}

func (l *LayerDataLoader) FindByPluginAndExtensionOfBlocks(ctx context.Context, pluginID id.PluginID, extensionID *id.PluginExtensionID) (layer.List, error) {
	return l.Layer.FindByPluginAndExtensionOfBlocks(ctx, pluginID, extensionID)
}

func (l *LayerDataLoader) UpdatePlugin(ctx context.Context, pluginID id.PluginID, pluginID2 id.PluginID) error {
	return l.Layer.UpdatePlugin(ctx, pluginID, pluginID2)
}

type LayerDataLoaderParam struct {
	LayerID string
}

func NewLayerDataloader(ctx context.Context, layer repo.Layer) *LayerDataLoader {
	l := &LayerDataLoader{Layer: layer}
	l.loader = l.NewLayerDataLoader(ctx)
	return l
}
func (l *LayerDataLoader) Fetch(ctx context.Context, keys []*id.LayerID) ([]layer.Layer, []error) {
	fmt.Printf("LayerDataLoader: %#v\n", keys)
	ids := make([]id.LayerID, 0, len(keys))
	for _, key := range keys {
		ids = append(ids, *key)
	}

	res, err := l.FindByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}
	layerList := make([]layer.Layer, 0, len(res))
	for _, l := range res {
		layerList = append(layerList, *l)
	}
	return layerList, nil
}

func (l *LayerDataLoader) NewLayerDataLoader(ctx context.Context) *LayerLoader {
	return NewLayerLoader(LayerLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []*id.LayerID) ([]layer.Layer, []error) {
			return l.Fetch(ctx, keys)
		},
	})
}

func (l *LayerDataLoader) FindByID(ctx context.Context, id id.LayerID) (layer.Layer, error) {
	//return l.Layer.FindByID(ctx, id)
	return DataLoadersFromContext(ctx).Layer.loader.Load(&id)
}

func (l *LayerDataLoader) FindByIDs(ctx context.Context, ids []id.LayerID) (layer.List, error) {
	return l.Layer.FindByIDs(ctx, ids)
}

func (l *LayerDataLoader) FindItemByID(ctx context.Context, id id.LayerID) (*layer.Item, error) {
	return l.FindItemByID(ctx, id)
}

func (l *LayerDataLoader) FindItemByIDs(ctx context.Context, ids []id.LayerID) (layer.ItemList, error) {
	return l.Layer.FindItemByIDs(ctx, ids)
}
func (l *LayerDataLoader) FindAllByDatasetSchema(ctx context.Context, id id.DatasetSchemaID) (layer.List, error) {
	return l.Layer.FindAllByDatasetSchema(ctx, id)
}
func (l *LayerDataLoader) FindGroupByID(ctx context.Context, id id.LayerID) (*layer.Group, error) {
	return l.Layer.FindParentByID(ctx, id)
}
func (l *LayerDataLoader) FindGroupByIDs(ctx context.Context, id []id.LayerID) (layer.GroupList, error) {
	return l.Layer.FindGroupByIDs(ctx, id)
}
func (l *LayerDataLoader) FindGroupBySceneAndLinkedDatasetSchema(ctx context.Context, sceneID id.SceneID, datasetSchemaID id.DatasetSchemaID) (layer.GroupList, error) {
	return l.Layer.FindGroupBySceneAndLinkedDatasetSchema(ctx, sceneID, datasetSchemaID)
}
func (l *LayerDataLoader) FindParentByID(ctx context.Context, id id.LayerID) (*layer.Group, error) {
	return l.Layer.FindParentByID(ctx, id)
}
func (l *LayerDataLoader) FindByProperty(ctx context.Context, id id.PropertyID) (layer.Layer, error) {
	return l.Layer.FindByProperty(ctx, id)
}
func (l *LayerDataLoader) FindByScene(ctx context.Context, id id.SceneID) (layer.List, error) {
	return l.Layer.FindByScene(ctx, id)
}
func (l *LayerDataLoader) FindByTag(ctx context.Context, tagID id.TagID) (layer.List, error) {
	return l.Layer.FindByTag(ctx, tagID)
}
func (l *LayerDataLoader) Save(ctx context.Context, layer layer.Layer) error {
	return l.Layer.Save(ctx, layer)
}
func (l *LayerDataLoader) SaveAll(ctx context.Context, layer layer.List) error {
	return l.Layer.SaveAll(ctx, layer)
}
func (l *LayerDataLoader) Remove(ctx context.Context, id id.LayerID) error {
	return l.Layer.Remove(ctx, id)
}
func (l *LayerDataLoader) RemoveAll(ctx context.Context, id []id.LayerID) error {
	return l.Layer.RemoveAll(ctx, id)
}
func (l *LayerDataLoader) RemoveByScene(ctx context.Context, id id.SceneID) error {
	return l.Layer.RemoveByScene(ctx, id)
}
