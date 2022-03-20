package gqlmodel

import (
	"github.com/reearth/reearth-backend/pkg/asset"
)

func ToAsset(a *asset.Asset) *Asset {
	if a == nil {
		return nil
	}
	return &Asset{
		ID:          a.ID().ID(),
		CreatedAt:   a.CreatedAt(),
		TeamID:      a.Team().ID(),
		Name:        a.Name(),
		Size:        a.Size(),
		URL:         a.URL(),
		ContentType: a.ContentType(),
	}
}

func AssetSortTypeFrom(ast *AssetSortType) *asset.SortType {
	if ast == nil {
		return nil
	}
	switch *ast {
	case AssetSortTypeDate:
		return &asset.SortTypeID
	case AssetSortTypeName:
		return &asset.SortTypeName
	case AssetSortTypeSize:
		return &asset.SortTypeSize
	}
	return &asset.SortTypeID
}
