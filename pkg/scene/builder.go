package scene

import (
	"time"
)

type Builder struct {
	scene *Scene
}

func New() *Builder {
	return &Builder{scene: &Scene{}}
}

func (b *Builder) Build() (*Scene, error) {
	if b.scene.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.scene.team.ID().IsNil() {
		return nil, ErrInvalidID
	}
	if b.scene.rootLayer.ID().IsNil() {
		return nil, ErrInvalidID
	}
	if b.scene.widgets == nil {
		b.scene.widgets = NewWidgets(nil, nil)
	}
	if b.scene.plugins == nil {
		b.scene.plugins = NewPlugins(nil)
	}
	if b.scene.updatedAt.IsZero() {
		b.scene.updatedAt = b.scene.CreatedAt()
	}
	return b.scene, nil
}

func (b *Builder) MustBuild() *Scene {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *Builder) ID(id ID) *Builder {
	b.scene.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.scene.id = NewID()
	return b
}

func (b *Builder) Project(prj ProjectID) *Builder {
	b.scene.project = prj
	return b
}

func (b *Builder) Team(team TeamID) *Builder {
	b.scene.team = team
	return b
}

func (b *Builder) UpdatedAt(updatedAt time.Time) *Builder {
	b.scene.updatedAt = updatedAt
	return b
}

func (b *Builder) Widgets(widgets *Widgets) *Builder {
	b.scene.widgets = widgets
	return b
}

func (b *Builder) RootLayer(rootLayer LayerID) *Builder {
	b.scene.rootLayer = rootLayer
	return b
}

func (b *Builder) Plugins(plugins *Plugins) *Builder {
	b.scene.plugins = plugins
	return b
}

func (b *Builder) Property(p PropertyID) *Builder {
	b.scene.property = p
	return b
}

func (b *Builder) Clusters(cl *ClusterList) *Builder {
	b.scene.clusters = cl
	return b
}
