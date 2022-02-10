//go:generate go run github.com/vektah/dataloaden LayerLoader string *github.com/reearth/reearth-backend/pkg/layer.Item

package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
)

type LayerDataLoader struct {
	Layer repo.Layer
	loader *LayerLoader
}

type LayerDataLoaderParam struct {
	LayerID  string
	SceneIDs []id.SceneID
}

func NewLayerDataloader(ctx context.Context, layer repo.Layer) *LayerDataLoader {
	l :=&LayerDataLoader{Layer: layer}
	l.loader = l.NewLayerDataLoader(ctx)
	return l
}
func (l *LayerDataLoader) Fetch(ctx context.Context, keys []string) ([]*layer.Item, []error) {
	ids := make([]id.LayerID, 0, len(keys))
	var sceneIDs []id.SceneID
	for _, key := range keys {
		fmt.Println("key:",key)
		var param LayerDataLoaderParam
		if  err := json.Unmarshal([]byte(key),&param); err != nil {
			return nil, []error{err}
		}
		layerID, err := id.LayerIDFrom(param.LayerID)
		if err != nil {
			return nil, []error{err}
		}
		ids = append(ids, layerID)
		sceneIDs = param.SceneIDs
	}
	//SceneIDは同一Contextであれば同じ(のはず)なので0を指定
	res, err := l.FindItemByIDs(ctx, ids, sceneIDs)
	if err != nil {
		return nil, []error{err}
	}
	layerItemList := make([]*layer.Item,0,len(res))
	for _,l := range res {
		layerItemList = append(layerItemList, l)
	}
	return layerItemList, nil
}

func (l *LayerDataLoader) NewLayerDataLoader(ctx context.Context) *LayerLoader {
	return NewLayerLoader(LayerLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []string) ([]*layer.Item,[]error) {
			return l.Fetch(ctx, keys)
		},
	})
}

func (l *LayerDataLoader) FindByID(ctx context.Context, id id.LayerID, scenes []id.SceneID) (layer.Layer, error) {
	return l.Layer.FindByID(ctx, id, scenes)
}

func (l *LayerDataLoader) FindByIDs(ctx context.Context, ids []id.LayerID, f []id.SceneID) (layer.List, error) {
	return l.Layer.FindByIDs(ctx, ids, f)
}

func (l *LayerDataLoader) FindItemByID(ctx context.Context, id id.LayerID, sids []id.SceneID) (*layer.Item, error) {
	//// Call data loader
	//
	fmt.Printf("id:%#v\n",id)
	fmt.Printf("id.String:%#v\n",id.String())
	key, err := json.Marshal(LayerDataLoaderParam{
		LayerID:  id.String(),
		SceneIDs: sids,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("FindItemById: %s", string(key))
	return DataLoadersFromContext(ctx).loader.Load(string(key))
}

func (l *LayerDataLoader) FindItemByIDs(ctx context.Context, ids []id.LayerID, f []id.SceneID) (layer.ItemList, error) {
	return l.Layer.FindItemByIDs(ctx, ids, f)
}
func (l *LayerDataLoader) FindAllByDatasetSchema(ctx context.Context, id id.DatasetSchemaID) (layer.List, error) {
	return l.Layer.FindAllByDatasetSchema(ctx, id)
}
func (l *LayerDataLoader) FindGroupByID(ctx context.Context, id id.LayerID, f []id.SceneID) (*layer.Group, error) {
	return l.Layer.FindParentByID(ctx, id, f)
}
func (l *LayerDataLoader) FindGroupByIDs(ctx context.Context, id []id.LayerID, f []id.SceneID) (layer.GroupList, error) {
	return l.Layer.FindGroupByIDs(ctx, id, f)
}
func (l *LayerDataLoader) FindGroupBySceneAndLinkedDatasetSchema(ctx context.Context, sceneID id.SceneID, datasetSchemaID id.DatasetSchemaID) (layer.GroupList, error) {
	return l.Layer.FindGroupBySceneAndLinkedDatasetSchema(ctx, sceneID, datasetSchemaID)
}
func (l *LayerDataLoader) FindParentByID(ctx context.Context, id id.LayerID, f []id.SceneID) (*layer.Group, error) {
	return l.Layer.FindParentByID(ctx, id, f)
}
func (l *LayerDataLoader) FindByProperty(ctx context.Context, id id.PropertyID, f []id.SceneID) (layer.Layer, error) {
	return l.Layer.FindByProperty(ctx, id, f)
}
func (l *LayerDataLoader) FindByScene(ctx context.Context, id id.SceneID) (layer.List, error) {
	return l.Layer.FindByScene(ctx, id)
}
func (l *LayerDataLoader) FindByTag(ctx context.Context, tagID id.TagID, f []id.SceneID) (layer.List, error) {
	return l.Layer.FindByTag(ctx, tagID, f)
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