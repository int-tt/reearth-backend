package gql

import (
	"context"

	"github.com/reearth/reearth-backend/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-backend/pkg/id"
)

type tagItemResolver struct{ *Resolver }

func (r *Resolver) TagItem() TagItemResolver {
	return &tagItemResolver{r}
}

func (t tagItemResolver) LinkedDatasetSchema(ctx context.Context, obj *gqlmodel.TagItem) (*gqlmodel.DatasetSchema, error) {
	if obj.LinkedDatasetID == nil {
		return nil, nil
	}
	return dataloaders(ctx).DatasetSchema.Load(id.DatasetSchemaID(*obj.LinkedDatasetSchemaID))
}

func (t tagItemResolver) LinkedDataset(ctx context.Context, obj *gqlmodel.TagItem) (*gqlmodel.Dataset, error) {
	if obj.LinkedDatasetID == nil {
		return nil, nil
	}
	return dataloaders(ctx).Dataset.Load(id.DatasetID(*obj.LinkedDatasetID))
}

func (t tagItemResolver) LinkedDatasetField(ctx context.Context, obj *gqlmodel.TagItem) (*gqlmodel.DatasetField, error) {
	if obj.LinkedDatasetID == nil {
		return nil, nil
	}
	ds, err := dataloaders(ctx).Dataset.Load(id.DatasetID(*obj.LinkedDatasetID))
	return ds.Field(*obj.LinkedDatasetFieldID), err
}

func (t tagItemResolver) Parent(ctx context.Context, obj *gqlmodel.TagItem) (*gqlmodel.TagGroup, error) {
	if obj.ParentID == nil {
		return nil, nil
	}
	return dataloaders(ctx).TagGroup.Load(id.TagID(*obj.ParentID))
}

func (tg tagItemResolver) Layers(ctx context.Context, obj *gqlmodel.TagItem) ([]gqlmodel.Layer, error) {
	return loaders(ctx).Layer.FetchByTag(ctx, id.TagID(obj.ID))
}

type tagGroupResolver struct{ *Resolver }

func (r *Resolver) TagGroup() TagGroupResolver {
	return &tagGroupResolver{r}
}

func (r tagGroupResolver) Tags(ctx context.Context, obj *gqlmodel.TagGroup) ([]*gqlmodel.TagItem, error) {
	tagIds := make([]id.TagID, 0, len(obj.TagIds))
	for _, i := range obj.TagIds {
		if i == nil {
			continue
		}
		tagIds = append(tagIds, id.TagID(*i))
	}
	tagItems, err := dataloaders(ctx).TagItem.LoadAll(tagIds)
	if len(err) > 0 && err[0] != nil {
		return nil, err[0]
	}
	return tagItems, nil
}

func (r tagGroupResolver) Scene(ctx context.Context, obj *gqlmodel.TagGroup) (*gqlmodel.Scene, error) {
	return dataloaders(ctx).Scene.Load(id.SceneID(obj.SceneID))
}

func (r tagGroupResolver) Layers(ctx context.Context, obj *gqlmodel.TagGroup) ([]gqlmodel.Layer, error) {
	return loaders(ctx).Layer.FetchByTag(ctx, id.TagID(obj.ID))
}
