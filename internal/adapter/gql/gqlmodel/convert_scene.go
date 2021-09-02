package gqlmodel

import (
	"github.com/reearth/reearth-backend/pkg/scene"
)

func ToSceneWidget(w *scene.Widget) *SceneWidget {
	if w == nil {
		return nil
	}

	return &SceneWidget{
		ID:          w.ID().ID(),
		PluginID:    w.Plugin(),
		ExtensionID: w.Extension(),
		PropertyID:  w.Property().ID(),
		Enabled:     w.Enabled(),
		Extended:    w.WidgetLayout().Extended,
	}
}

func ToScenePlugin(sp *scene.Plugin) *ScenePlugin {
	if sp == nil {
		return nil
	}

	return &ScenePlugin{
		PluginID:   sp.Plugin(),
		PropertyID: sp.Property().IDRef(),
	}
}

func ToScene(scene *scene.Scene) *Scene {
	if scene == nil {
		return nil
	}

	sceneWidgets := scene.WidgetSystem().Widgets()
	widgets := make([]*SceneWidget, 0, len(sceneWidgets))
	for _, w := range sceneWidgets {
		widgets = append(widgets, ToSceneWidget(w))
	}

	scenePlugins := scene.PluginSystem().Plugins()
	plugins := make([]*ScenePlugin, 0, len(scenePlugins))
	for _, sp := range scenePlugins {
		plugins = append(plugins, ToScenePlugin(sp))
	}

	sceneWidgetAlignSystem := scene.WidgetAlignSystem()
	widgetAlignSystem := ToAlignSystem(sceneWidgetAlignSystem)

	return &Scene{
		ID:                scene.ID().ID(),
		ProjectID:         scene.Project().ID(),
		PropertyID:        scene.Property().ID(),
		TeamID:            scene.Team().ID(),
		RootLayerID:       scene.RootLayer().ID(),
		CreatedAt:         scene.CreatedAt(),
		UpdatedAt:         scene.UpdatedAt(),
		Widgets:           widgets,
		WidgetAlignSystem: widgetAlignSystem,
		Plugins:           plugins,
	}
}

func ToSceneLockMode(lm scene.LockMode) SceneLockMode {
	switch lm {
	case scene.LockModeFree:
		return SceneLockModeFree
	case scene.LockModePending:
		return SceneLockModePending
	case scene.LockModeDatasetSyncing:
		return SceneLockModeDatasetSyncing
	case scene.LockModePluginUpgrading:
		return SceneLockModePluginUpgrading
	case scene.LockModePublishing:
		return SceneLockModePublishing
	}
	return SceneLockMode("invalid")
}
