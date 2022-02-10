package three

//go:generate go run object3d_method_generator/main.go -typeName AxesHelper -typeSlug axes_helper

import "syscall/js"

type AxesHelper struct {
	js.Value

	// Copied from LineSegments struct:

	// ID               int             `js:"id"`
	// Position         *Vector3        `js:"position"`
	// Rotation         *Euler          `js:"rotation"`
	// Geometry         *BufferGeometry `js:"geometry"`
	// Material         Material        `js:"material"`
	// MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	// RenderOrder      int             `js:"renderOrder"`
	// Visible          bool            `js:"visible"`
}

func NewAxesHelper(size float64) *AxesHelper {
	if size == 0 {
		size = 1
	}
	return &AxesHelper{
		Value: three.Get("AxesHelper").New(size),
	}
}

func (g AxesHelper) SetColors(xaxis, yaxis, zaxis *Color) AxesHelper {
	g.Value.Call("setColor", xaxis, yaxis, zaxis)
	return g
}

// Disposes of the internally-created material and geometry used by this helper.
func (g AxesHelper) Dispose() {
	g.Value.Call("dispose")
}
