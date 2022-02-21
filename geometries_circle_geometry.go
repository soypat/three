//go:build TODO

package three

//go:generate go run codegen/geometry_method_generator/main.go -typeName CircleGeometry -typeSlug circle_geometry

import "syscall/js"

type CircleGeometry struct {
	*js.Object

	Radius      float64 `js:"radius"`
	Segments    int     `js:"segments"`
	ThetaStart  float64 `js:"thetaStart"`
	ThetaLength float64 `js:"thetaLength"`
}

type CircleGeometryParameters struct {
	Radius      float64
	Segments    int
	ThetaStart  float64
	ThetaLength float64
}

// NewBoxGeometry creates a new BoxGeometry.
func NewCircleGeometry(params CircleGeometryParameters) CircleGeometry {
	if params.Segments < 3 {
		params.Segments = 8
	}
	return CircleGeometry{
		Object: three.Get("CircleGeometry").New(
			params.Radius,
			params.Segments,
			params.ThetaStart,
			params.ThetaLength,
		),
	}
}
