package scene

import (
	"errors"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewWidget(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	wid := id.NewWidgetID()
	testCases := []struct {
		Name         string
		Id           *id.WidgetID
		Plugin       id.PluginID
		Extension    id.PluginExtensionID
		Property     id.PropertyID
		Enabled      bool
		WidgetLayout *WidgetLayout
		Expected     struct {
			Id           id.WidgetID
			Plugin       id.PluginID
			Extension    id.PluginExtensionID
			Property     id.PropertyID
			Enabled      bool
			WidgetLayout *WidgetLayout
		}
		err error
	}{
		{
			Name:         "success new widget",
			Id:           wid.Ref(),
			Plugin:       pid,
			Extension:    "eee",
			Property:     pr,
			Enabled:      true,
			WidgetLayout: nil,
			Expected: struct {
				Id           id.WidgetID
				Plugin       id.PluginID
				Extension    id.PluginExtensionID
				Property     id.PropertyID
				Enabled      bool
				WidgetLayout *WidgetLayout
			}{
				Id:           wid,
				Plugin:       pid,
				Extension:    "eee",
				Property:     pr,
				Enabled:      true,
				WidgetLayout: nil,
			},
			err: nil,
		},
		{
			Name:         "success nil id",
			Id:           nil,
			Plugin:       pid,
			Extension:    "eee",
			Property:     pr,
			Enabled:      true,
			WidgetLayout: nil,
			Expected: struct {
				Id           id.WidgetID
				Plugin       id.PluginID
				Extension    id.PluginExtensionID
				Property     id.PropertyID
				Enabled      bool
				WidgetLayout *WidgetLayout
			}{
				Id:           wid,
				Plugin:       pid,
				Extension:    "eee",
				Property:     pr,
				Enabled:      true,
				WidgetLayout: nil,
			},
			err: nil,
		},
		{
			Name:         "fail empty extension",
			Id:           wid.Ref(),
			Plugin:       pid,
			Extension:    "",
			Property:     pr,
			Enabled:      true,
			WidgetLayout: nil,
			err:          id.ErrInvalidID,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res, err := NewWidget(tc.Id, tc.Plugin, tc.Extension, tc.Property, tc.Enabled, tc.WidgetLayout)
			if err == nil {
				if tc.Id == nil {
					assert.NotNil(tt, res.ID())
				} else {
					assert.Equal(tt, tc.Expected.Id, res.ID())
				}
				assert.Equal(tt, tc.Expected.Property, res.Property())
				assert.Equal(tt, tc.Expected.Extension, res.Extension())
				assert.Equal(tt, tc.Expected.Enabled, res.Enabled())
				assert.Equal(tt, tc.Expected.WidgetLayout, res.WidgetLayout())
				assert.Equal(tt, tc.Expected.Plugin, res.Plugin())
			} else {
				assert.True(tt, errors.As(tc.err, &err))
			}
		})
	}
}
func TestMustNewWidget(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	wid := id.NewWidgetID()
	testCases := []struct {
		Name         string
		Id           *id.WidgetID
		Plugin       id.PluginID
		Extension    id.PluginExtensionID
		Property     id.PropertyID
		Enabled      bool
		WidgetLayout *WidgetLayout
		Expected     struct {
			Id           id.WidgetID
			Plugin       id.PluginID
			Extension    id.PluginExtensionID
			Property     id.PropertyID
			Enabled      bool
			WidgetLayout *WidgetLayout
		}
		err error
	}{
		{
			Name:         "success new widget",
			Id:           wid.Ref(),
			Plugin:       pid,
			Extension:    "eee",
			Property:     pr,
			Enabled:      true,
			WidgetLayout: nil,
			Expected: struct {
				Id           id.WidgetID
				Plugin       id.PluginID
				Extension    id.PluginExtensionID
				Property     id.PropertyID
				Enabled      bool
				WidgetLayout *WidgetLayout
			}{
				Id:           wid,
				Plugin:       pid,
				Extension:    "eee",
				Property:     pr,
				Enabled:      true,
				WidgetLayout: nil,
			},
			err: nil,
		},
		{
			Name:         "success nil id",
			Id:           nil,
			Plugin:       pid,
			Extension:    "eee",
			Property:     pr,
			Enabled:      true,
			WidgetLayout: nil,
			Expected: struct {
				Id           id.WidgetID
				Plugin       id.PluginID
				Extension    id.PluginExtensionID
				Property     id.PropertyID
				Enabled      bool
				WidgetLayout *WidgetLayout
			}{
				Id:           wid,
				Plugin:       pid,
				Extension:    "eee",
				Property:     pr,
				Enabled:      true,
				WidgetLayout: nil,
			},
			err: nil,
		},
		{
			Name:         "fail empty extension",
			Id:           wid.Ref(),
			Plugin:       pid,
			Extension:    "",
			Property:     pr,
			Enabled:      true,
			WidgetLayout: nil,
			err:          id.ErrInvalidID,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			var res *Widget
			defer func() {
				if r := recover(); r == nil {
					if tc.Id == nil {
						assert.NotNil(tt, res.ID())
					} else {
						assert.Equal(tt, tc.Expected.Id, res.ID())
					}
					assert.Equal(tt, tc.Expected.Property, res.Property())
					assert.Equal(tt, tc.Expected.Extension, res.Extension())
					assert.Equal(tt, tc.Expected.Enabled, res.Enabled())
					assert.Equal(tt, tc.Expected.Plugin, res.Plugin())
				}
			}()
			res = MustNewWidget(tc.Id, tc.Plugin, tc.Extension, tc.Property, tc.Enabled, tc.WidgetLayout)

		})
	}
}

func TestWidget_SetEnabled(t *testing.T) {
	res := MustNewWidget(id.NewWidgetID().Ref(), id.MustPluginID("xxx#1.1.1"), "eee", id.NewPropertyID(), false, nil)
	res.SetEnabled(true)
	assert.True(t, res.Enabled())
}
