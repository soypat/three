package three

//go:generate go run codegen/object3d_method_generator/main.go -typeName Line -typeSlug line

import "syscall/js"

// Line a continuous line.
type Line struct {
	// ID               int             `js:"id"`
	// Position         *Vector3        `js:"position"`
	// Rotation         *Euler          `js:"rotation"`
	// Geometry         *BufferGeometry `js:"geometry"`
	// Material         Material        `js:"material"`
	// MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	// RenderOrder      int             `js:"renderOrder"`
	js.Value
}

// NewLine creates a new material. If Material is nil, three.js will assign a randomized material to the line o_O.
func NewLine(geom Geometry, material Material) *Line {
	return &Line{
		Value: three.Get("Line").New(geom.getInternalObject(), material.getInternalObject()),
	}
}
