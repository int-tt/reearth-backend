package scene

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewWidgetSystem(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	wid := id.NewWidgetID().Ref()
	testCases := []struct {
		Name     string
		Input    []*Widget
		Expected *WidgetSystem
	}{
		{
			Name:     "nil widget list",
			Input:    nil,
			Expected: &WidgetSystem{widgets: []*Widget{}},
		},
		{
			Name:     "widget list with nil",
			Input:    []*Widget{nil},
			Expected: &WidgetSystem{widgets: []*Widget{}},
		},
		{
			Name: "widget list with matched values",
			Input: []*Widget{
				{
					id:           *wid,
					plugin:       pid,
					extension:    "eee",
					property:     pr,
					enabled:      true,
					widgetLayout: &WidgetLayout{},
				},
			},
			Expected: &WidgetSystem{widgets: []*Widget{
				MustNewWidget(wid, pid, "eee", pr, true, &WidgetLayout{}),
			}},
		},
		{
			Name: "widget list with matched values",
			Input: []*Widget{
				{
					id:           *wid,
					plugin:       pid,
					extension:    "eee",
					property:     pr,
					enabled:      true,
					widgetLayout: &WidgetLayout{},
				},
				{
					id:           *wid,
					plugin:       pid,
					extension:    "eee",
					property:     pr,
					enabled:      true,
					widgetLayout: &WidgetLayout{},
				},
			},
			Expected: &WidgetSystem{widgets: []*Widget{
				MustNewWidget(wid, pid, "eee", pr, true, &WidgetLayout{}),
			}},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := NewWidgetSystem(tc.Input)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetSystem_Add(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	wid := id.NewWidgetID().Ref()
	testCases := []struct {
		Name         string
		Input        *Widget
		WS, Expected *WidgetSystem
	}{
		{
			Name: "add new widget",
			Input: &Widget{
				id:           *wid,
				plugin:       pid,
				extension:    "eee",
				property:     pr,
				enabled:      true,
				widgetLayout: &WidgetLayout{},
			},
			WS:       NewWidgetSystem([]*Widget{}),
			Expected: NewWidgetSystem([]*Widget{MustNewWidget(wid, pid, "eee", pr, true, &WidgetLayout{})}),
		},
		{
			Name:     "add nil widget",
			Input:    nil,
			WS:       NewWidgetSystem([]*Widget{}),
			Expected: NewWidgetSystem([]*Widget{}),
		},
		{
			Name: "add to nil widgetSystem",
			Input: &Widget{
				id:           *wid,
				plugin:       pid,
				extension:    "eee",
				property:     pr,
				enabled:      true,
				widgetLayout: &WidgetLayout{},
			},
			WS:       nil,
			Expected: nil,
		},
		{
			Name: "add existing widget",
			Input: &Widget{
				id:           *wid,
				plugin:       pid,
				extension:    "eee",
				property:     pr,
				enabled:      true,
				widgetLayout: &WidgetLayout{},
			},
			WS:       NewWidgetSystem([]*Widget{MustNewWidget(wid, pid, "eee", pr, true, &WidgetLayout{})}),
			Expected: NewWidgetSystem([]*Widget{MustNewWidget(wid, pid, "eee", pr, true, &WidgetLayout{})}),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WS.Add(tc.Input)
			assert.Equal(tt, tc.Expected, tc.WS)
		})
	}
}

func TestWidgetSystem_Remove(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	wid := id.NewWidgetID().Ref()
	testCases := []struct {
		Name         string
		PID          id.PluginID
		EID          id.PluginExtensionID
		WS, Expected *WidgetSystem
	}{
		{
			Name:     "remove a widget",
			PID:      pid,
			EID:      "eee",
			WS:       NewWidgetSystem([]*Widget{MustNewWidget(wid, pid, "eee", pr, true, nil)}),
			Expected: NewWidgetSystem([]*Widget{}),
		},
		{
			Name:     "remove from nil widgetSystem",
			WS:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WS.Remove(tc.PID, tc.EID)
			assert.Equal(tt, tc.Expected, tc.WS)
		})
	}
}

func TestWidgetSystem_Replace(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pid2 := id.MustPluginID("zzz#1.1.1")
	pr := id.NewPropertyID()
	wid := id.NewWidgetID().Ref()
	testCases := []struct {
		Name         string
		PID, NewID   id.PluginID
		WS, Expected *WidgetSystem
	}{
		{
			Name:     "replace a widget",
			PID:      pid,
			NewID:    pid2,
			WS:       NewWidgetSystem([]*Widget{MustNewWidget(wid, pid, "eee", pr, true, nil)}),
			Expected: NewWidgetSystem([]*Widget{MustNewWidget(wid, pid2, "eee", pr, true, nil)}),
		},
		{
			Name:     "replace with nil widget",
			PID:      pid,
			WS:       NewWidgetSystem(nil),
			Expected: NewWidgetSystem(nil),
		},
		{
			Name:     "replace from nil widgetSystem",
			WS:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			tc.WS.Replace(tc.PID, tc.NewID)
			assert.Equal(tt, tc.Expected, tc.WS)
		})
	}
}

func TestWidgetSystem_Properties(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	pr2 := id.NewPropertyID()
	wid := id.NewWidgetID().Ref()
	wid2 := id.NewWidgetID().Ref()
	testCases := []struct {
		Name     string
		WS       *WidgetSystem
		Expected []id.PropertyID
	}{
		{
			Name: "get properties",
			WS: NewWidgetSystem([]*Widget{
				MustNewWidget(wid, pid, "eee", pr, true, nil),
				MustNewWidget(wid2, pid, "eee", pr2, true, nil),
			}),
			Expected: []id.PropertyID{pr, pr2},
		},
		{
			Name:     "get properties from nil widgetSystem",
			WS:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WS.Properties()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetSystem_Widgets(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	pr2 := id.NewPropertyID()
	wid := id.NewWidgetID().Ref()
	wid2 := id.NewWidgetID().Ref()
	testCases := []struct {
		Name     string
		WS       *WidgetSystem
		Expected []*Widget
	}{
		{
			Name: "get widgets",
			WS: NewWidgetSystem([]*Widget{
				MustNewWidget(wid, pid, "eee", pr, true, nil),
				MustNewWidget(wid2, pid, "eee", pr2, true, nil),
			}),
			Expected: []*Widget{
				MustNewWidget(wid, pid, "eee", pr, true, nil),
				MustNewWidget(wid2, pid, "eee", pr2, true, nil),
			},
		},
		{
			Name:     "get widgets from nil widgetSystem",
			WS:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WS.Widgets()
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetSystem_Widget(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	wid := id.NewWidgetID().Ref()
	testCases := []struct {
		Name     string
		PID      id.PluginID
		EID      id.PluginExtensionID
		WS       *WidgetSystem
		Expected *Widget
	}{
		{
			Name:     "get a widget",
			PID:      pid,
			EID:      "eee",
			WS:       NewWidgetSystem([]*Widget{MustNewWidget(wid, pid, "eee", pr, true, nil)}),
			Expected: MustNewWidget(wid, pid, "eee", pr, true, nil),
		},
		{
			Name:     "dont has the widget",
			PID:      pid,
			EID:      "eee",
			WS:       NewWidgetSystem([]*Widget{}),
			Expected: nil,
		},
		{
			Name:     "get widget from nil widgetSystem",
			WS:       nil,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WS.Widget(tc.PID, tc.EID)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}

func TestWidgetSystem_Has(t *testing.T) {
	pid := id.MustPluginID("xxx#1.1.1")
	pr := id.NewPropertyID()
	wid := id.NewWidgetID().Ref()
	testCases := []struct {
		Name     string
		PID      id.PluginID
		EID      id.PluginExtensionID
		WS       *WidgetSystem
		Expected bool
	}{
		{
			Name:     "has a widget",
			PID:      pid,
			EID:      "eee",
			WS:       NewWidgetSystem([]*Widget{MustNewWidget(wid, pid, "eee", pr, true, nil)}),
			Expected: true,
		},
		{
			Name:     "dont has a widget",
			PID:      pid,
			EID:      "eee",
			WS:       NewWidgetSystem([]*Widget{}),
			Expected: false,
		},
		{
			Name:     "has from nil widgetSystem",
			WS:       nil,
			Expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res := tc.WS.Has(tc.PID, tc.EID)
			assert.Equal(tt, tc.Expected, res)
		})
	}
}
