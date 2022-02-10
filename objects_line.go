//go:build TODO

package three

//go:generate go run object3d_method_generator/main.go -typeName Line -typeSlug line

import (
	"syscall/js"
)

// Line a continuous line.
type Line struct {
	*js.Object

	ID               int             `js:"id"`
	Position         *Vector3        `js:"position"`
	Rotation         *Euler          `js:"rotation"`
	Geometry         *BufferGeometry `js:"geometry"`
	Material         Material        `js:"material"`
	MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	RenderOrder      int             `js:"renderOrder"`
}

// NewLine creates a new material. If Material is nil, three.js will assign a randomized material to the line o_O.
func NewLine(geom Geometry, material Material) *Line {
	return &Line{
		Object: three.Get("Line").New(geom, material),
	}
}
