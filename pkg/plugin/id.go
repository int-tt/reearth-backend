package plugin

import "github.com/reearth/reearth-backend/pkg/id"

type ID = id.PluginID
type ExtensionID = id.PluginExtensionID
type PropertySchemaID = id.PropertySchemaID
type SceneID = id.SceneID

var NewID = id.NewPluginID
var NewSceneID = id.NewSceneID
var NewPropertySchemaID = id.NewPropertySchemaID

var MustID = id.MustPluginID
var MustSceneID = id.MustSceneID
var MustPropertySchemaID = id.MustPropertySchemaID

var IDFrom = id.PluginIDFrom
var SceneIDFrom = id.SceneIDFrom
var PropertySchemaIDFrom = id.PropertySchemaIDFrom

var IDFromRef = id.PluginIDFromRef
var ExtensionIDFromRef = id.PluginExtensionIDFromRef
var SceneIDFromRef = id.SceneIDFromRef
var PropertySchemaIDFromRef = id.PropertySchemaIDFromRef

var SceneIDFromRefID = id.SceneIDFromRefID

var OfficialPluginID = id.OfficialPluginID
var ErrInvalidID = id.ErrInvalidID
