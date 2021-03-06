package three

//go:generate go run codegen/object3d_method_generator/main.go -typeName Sprite -typeSlug sprite

import "syscall/js"

// Sprite is a plane that always faces towards the camera, generally with a partially transparent texture applied.
type Sprite struct {
	// ID               int             `js:"id"`
	// Position         *Vector3        `js:"position"`
	// Scale            *Vector3        `js:"scale"`
	// Rotation         *Euler          `js:"rotation"`
	// Geometry         *BufferGeometry `js:"geometry"`
	// Material         Material        `js:"material"`
	// MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	// RenderOrder      int             `js:"renderOrder"`
	js.Value
}

// NewSprite creates a new sprite.
func NewSprite(material *SpriteMaterial) *Sprite {
	return &Sprite{
		Value: three.Get("Sprite").New(material),
	}
}
