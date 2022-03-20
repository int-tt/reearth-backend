//go:generate go run github.com/vektah/dataloaden PropertyLoader *github.com/reearth/reearth-backend/pkg/id.PropertyID *github.com/reearth/reearth-backend/pkg/property.Property

package mongo

import (
	"context"
	"fmt"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
)

type PropertyDataLoader struct {
	Property repo.Property
	loader   *PropertyLoader
}

func (p *PropertyDataLoader) Fetch(ctx context.Context, keys []*id.PropertyID) ([]*property.Property, []error) {
	fmt.Printf("PropertyDataLoader: %#v\n", keys)
	ids := make([]id.PropertyID, 0, 0)
	for _, key := range keys {
		ids = append(ids, *key)
	}
	res, err := p.Property.FindByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}
	propertyList := make([]*property.Property, 0, len(res))
	for _, r := range res {
		propertyList = append(propertyList, r)
	}
	return propertyList, nil
}
func NewPropertyDataLoader(ctx context.Context, property repo.Property) *PropertyDataLoader {
	p := &PropertyDataLoader{Property: property}
	p.loader = p.NewPropertyDataLoader(ctx)
	return p
}
func (p *PropertyDataLoader) NewPropertyDataLoader(ctx context.Context) *PropertyLoader {
	return NewPropertyLoader(PropertyLoaderConfig{
		Fetch: func(keys []*id.PropertyID) ([]*property.Property, []error) {
			return p.Fetch(ctx, keys)
		},
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
	})
}

func (p PropertyDataLoader) Filtered(filter repo.SceneFilter) repo.Property {
	return PropertyDataLoader{Property: p.Property.Filtered(filter)}
}

func (p PropertyDataLoader) FindByID(ctx context.Context, id id.PropertyID) (*property.Property, error) {
	fmt.Println("Call PropertyDataloader------------")
	return DataLoadersFromContext(ctx).Property.loader.Load(&id)
	//return p.Property.FindByID(ctx, id)
}

func (p PropertyDataLoader) FindByIDs(ctx context.Context, ids []id.PropertyID) (property.List, error) {
	//return DataLoadersFromContext(ctx).Property.loader.Load(ids)
	return p.Property.FindByIDs(ctx, ids)
}

func (p PropertyDataLoader) FindLinkedAll(ctx context.Context, id id.SceneID) (property.List, error) {
	return p.Property.FindLinkedAll(ctx, id)
}

func (p PropertyDataLoader) FindByDataset(ctx context.Context, id id.DatasetSchemaID, id2 id.DatasetID) (property.List, error) {
	return p.FindByDataset(ctx, id, id2)
}

func (p PropertyDataLoader) FindBySchema(ctx context.Context, ids []id.PropertySchemaID, id id.SceneID) (property.List, error) {
	return p.FindBySchema(ctx, ids, id)
}

func (p PropertyDataLoader) FindByPlugin(ctx context.Context, id id.PluginID, id2 id.SceneID) (property.List, error) {
	return p.Property.FindByPlugin(ctx, id, id2)
}

func (p PropertyDataLoader) Save(ctx context.Context, property *property.Property) error {
	return p.Property.Save(ctx, property)
}

func (p PropertyDataLoader) SaveAll(ctx context.Context, list property.List) error {
	return p.Property.SaveAll(ctx, list)
}

func (p PropertyDataLoader) UpdateSchemaPlugin(ctx context.Context, id id.PluginID, id2 id.PluginID, id3 id.SceneID) error {
	return p.Property.UpdateSchemaPlugin(ctx, id, id2, id3)
}

func (p PropertyDataLoader) Remove(ctx context.Context, id id.PropertyID) error {
	return p.Property.Remove(ctx, id)
}

func (p PropertyDataLoader) RemoveAll(ctx context.Context, ids []id.PropertyID) error {
	return p.Property.RemoveAll(ctx, ids)
}

func (p PropertyDataLoader) RemoveByScene(ctx context.Context, id id.SceneID) error {
	return p.Property.RemoveByScene(ctx, id)
}
