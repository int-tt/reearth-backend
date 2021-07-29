package mongodoc

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

type Location struct {
	Zone    string
	Section string
	Area    string
}

type WidgetLayout struct {
	Extendable      bool
	Extended        bool
	Floating        bool
	DefaultLocation *Location
}

type WidgetArea struct {
	WidgetIds []string
	Align     string
}

type WidgetSection struct {
	Top    WidgetArea
	Middle WidgetArea
	Bottom WidgetArea
}

type WidgetZone struct {
	Left   WidgetSection
	Center WidgetSection
	Right  WidgetSection
}

type SceneAlignSystemDocument struct {
	Inner WidgetZone
	Outer WidgetZone
}

func NewWidgetAlignSystem(was scene.WidgetAlignSystem) *SceneAlignSystemDocument {
	widgetAlignDoc := SceneAlignSystemDocument{Inner: WidgetZone{
		Left: WidgetSection{
			Top: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "left", "top")),
				Align:     *was.Alignment("inner", "left", "top"),
			},
			Middle: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "left", "middle")),
				Align:     *was.Alignment("inner", "left", "middle"),
			},
			Bottom: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "left", "bottom")),
				Align:     *was.Alignment("inner", "left", "bottom"),
			},
		},
		Center: WidgetSection{
			Top: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "center", "top")),
				Align:     *was.Alignment("inner", "center", "top"),
			},
			Middle: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "center", "middle")),
				Align:     *was.Alignment("inner", "center", "middle"),
			},
			Bottom: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "center", "bottom")),
				Align:     *was.Alignment("inner", "center", "bottom"),
			},
		},
		Right: WidgetSection{
			Top: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "right", "top")),
				Align:     *was.Alignment("inner", "right", "top"),
			},
			Middle: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "right", "middle")),
				Align:     *was.Alignment("inner", "right", "middle"),
			},
			Bottom: WidgetArea{
				WidgetIds: toString(was.WidgetIds("inner", "right", "bottom")),
				Align:     *was.Alignment("inner", "right", "bottom"),
			},
		},
	}, Outer: WidgetZone{
		Left: WidgetSection{
			Top: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "left", "top")),
				Align:     *was.Alignment("outer", "left", "top"),
			},
			Middle: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "left", "middle")),
				Align:     *was.Alignment("outer", "left", "middle"),
			},
			Bottom: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "left", "bottom")),
				Align:     *was.Alignment("outer", "left", "bottom"),
			},
		},
		Center: WidgetSection{
			Top: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "center", "top")),
				Align:     *was.Alignment("outer", "center", "top"),
			},
			Middle: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "center", "middle")),
				Align:     *was.Alignment("outer", "center", "middle"),
			},
			Bottom: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "center", "bottom")),
				Align:     *was.Alignment("outer", "center", "bottom"),
			},
		},
		Right: WidgetSection{
			Top: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "right", "top")),
				Align:     *was.Alignment("outer", "right", "top"),
			},
			Middle: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "right", "middle")),
				Align:     *was.Alignment("outer", "right", "middle"),
			},
			Bottom: WidgetArea{
				WidgetIds: toString(was.WidgetIds("outer", "right", "bottom")),
				Align:     *was.Alignment("outer", "right", "bottom"),
			},
		},
	}}
	return &widgetAlignDoc
}

func (*SceneAlignSystemDocument) ToModelAlignSystem(d SceneAlignSystemDocument) (*scene.WidgetAlignSystem, error) {
	was := scene.NewWidgetAlignSystem()
	// Inner Left
	was.WidgetAreaFrom(toStruct(d.Inner.Left.Top.WidgetIds), d.Inner.Left.Top.Align, "inner", "left", "top")
	was.WidgetAreaFrom(toStruct(d.Inner.Left.Middle.WidgetIds), d.Inner.Left.Middle.Align, "inner", "left", "middle")
	was.WidgetAreaFrom(toStruct(d.Inner.Left.Bottom.WidgetIds), d.Inner.Left.Bottom.Align, "inner", "left", "bottom")
	// Inner Center
	was.WidgetAreaFrom(toStruct(d.Inner.Center.Top.WidgetIds), d.Inner.Center.Top.Align, "inner", "center", "top")
	was.WidgetAreaFrom(toStruct(d.Inner.Center.Middle.WidgetIds), d.Inner.Center.Middle.Align, "inner", "center", "middle")
	was.WidgetAreaFrom(toStruct(d.Inner.Center.Bottom.WidgetIds), d.Inner.Center.Bottom.Align, "inner", "center", "bottom")
	// Inner Right
	was.WidgetAreaFrom(toStruct(d.Inner.Right.Top.WidgetIds), d.Inner.Right.Top.Align, "inner", "right", "top")
	was.WidgetAreaFrom(toStruct(d.Inner.Right.Middle.WidgetIds), d.Inner.Right.Middle.Align, "inner", "right", "middle")
	was.WidgetAreaFrom(toStruct(d.Inner.Right.Bottom.WidgetIds), d.Inner.Right.Bottom.Align, "inner", "right", "bottom")
	// Outer Left
	was.WidgetAreaFrom(toStruct(d.Outer.Left.Top.WidgetIds), d.Outer.Left.Top.Align, "outer", "left", "top")
	was.WidgetAreaFrom(toStruct(d.Outer.Left.Middle.WidgetIds), d.Outer.Left.Middle.Align, "outer", "left", "middle")
	was.WidgetAreaFrom(toStruct(d.Outer.Left.Bottom.WidgetIds), d.Outer.Left.Bottom.Align, "outer", "left", "bottom")
	// Outer Center
	was.WidgetAreaFrom(toStruct(d.Outer.Center.Top.WidgetIds), d.Outer.Center.Top.Align, "outer", "center", "top")
	was.WidgetAreaFrom(toStruct(d.Outer.Center.Middle.WidgetIds), d.Outer.Center.Middle.Align, "outer", "center", "middle")
	was.WidgetAreaFrom(toStruct(d.Outer.Center.Bottom.WidgetIds), d.Outer.Center.Bottom.Align, "outer", "center", "bottom")
	// Outer Right
	was.WidgetAreaFrom(toStruct(d.Outer.Right.Top.WidgetIds), d.Outer.Right.Top.Align, "outer", "right", "top")
	was.WidgetAreaFrom(toStruct(d.Outer.Right.Middle.WidgetIds), d.Outer.Right.Middle.Align, "outer", "right", "middle")
	was.WidgetAreaFrom(toStruct(d.Outer.Right.Bottom.WidgetIds), d.Outer.Right.Bottom.Align, "outer", "right", "bottom")
	return was, nil
}

func toString(wids []*id.WidgetID) []string {
	if wids == nil {
		return nil
	}
	docids := make([]string, 0, len(wids))
	for _, wid := range wids {
		if wid == nil {
			continue
		}
		docids = append(docids, wid.String())
	}
	return docids
}

func toStruct(wids []string) []*id.WidgetID {
	if wids == nil {
		return nil
	}
	var docids []*id.WidgetID
	for _, wid := range wids {
		nid, err := id.WidgetIDFrom(wid)
		if err != nil {
			continue
		}
		docids = append(docids, &nid)
	}
	return docids
}
