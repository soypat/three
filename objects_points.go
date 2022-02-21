package three

//go:generate go run codegen/object3d_method_generator/main.go -typeName Points -typeSlug points

import "syscall/js"

// A class for displaying points. The points are rendered by the WebGLRenderer using gl.POINTS.
type Points struct {
	// ID        int             `js:"id"`
	// Position  *Vector3        `js:"position"`
	// Geometry  *BufferGeometry `js:"geometry"`
	// Material  Material        `js:"material"`
	js.Value
}

// NewLine creates a new material. If Material is nil, three.js will assign a randomized material to the line o_O.
func NewPoints(geom Geometry, material Material) *Points {
	return &Points{
		Value: three.Get("Points").New(geom.getInternalObject(), material.getInternalObject()),
	}
}

// Get intersections between a casted ray and this Points. Raycaster.intersectObject will call this method.
func (p *Points) Raycast(rc Raycaster, intersects js.Value) {
	p.Value.Call("raycast", rc.Value, intersects)
}
