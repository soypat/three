//go:build TODO

package three

//go:generate go run object3d_method_generator/main.go -typeName LineSegments -typeSlug line_segments

import "syscall/js"

// LineSegments represents series of lines drawn between pairs of vertices.
type LineSegments struct {
	*js.Object

	ID               int             `js:"id"`
	Position         *Vector3        `js:"position"`
	Rotation         *Euler          `js:"rotation"`
	Geometry         *BufferGeometry `js:"geometry"`
	Material         Material        `js:"material"`
	MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	RenderOrder      int             `js:"renderOrder"`
	Visible          bool            `js:"visible"`
}

// NewLineSegments creates new line segments
func NewLineSegments(geom Geometry, material Material) *LineSegments {
	return &LineSegments{
		Object: three.Get("LineSegments").New(geom, material),
	}
}
