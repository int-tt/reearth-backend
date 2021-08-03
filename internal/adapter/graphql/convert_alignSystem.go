package graphql

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

func toAlignSystem(sas *scene.WidgetAlignSystem) *WidgetAlignSystem {
	widgetAlignDoc := WidgetAlignSystem{
		Inner: toWidgetZone(sas.Zone("inner")),
		Outer: toWidgetZone(sas.Zone("outer")),
	}
	return &widgetAlignDoc
}

func IDsFrom(wids []id.WidgetID) []*id.ID {
	var nids []*id.ID
	for _, w := range wids {
		nids = append(nids, w.IDRef())
	}
	return nids
}

func toWidgetZone(z *scene.WidgetZone) *WidgetZone {
	if z == nil {
		return nil
	}
	return &WidgetZone{
		Left:   toWidgetSection(z.Section(scene.SectionLeft)),
		Center: toWidgetSection(z.Section(scene.SectionCenter)),
		Right:  toWidgetSection(z.Section(scene.SectionRight)),
	}
}

func toWidgetSection(s *scene.WidgetSection) *WidgetSection {
	if s == nil {
		return nil
	}
	return &WidgetSection{
		Top:    toWidgetArea(s.Area(scene.AreaTop)),
		Middle: toWidgetArea(s.Area(scene.AreaMiddle)),
		Bottom: toWidgetArea(s.Area(scene.AreaBottom)),
	}
}

func toWidgetArea(a *scene.WidgetArea) *WidgetArea {
	if a == nil {
		return nil
	}
	return &WidgetArea{
		WidgetIds: IDsFrom(a.WidgetIDs()),
		Align:     a.Alignment(),
	}
}
