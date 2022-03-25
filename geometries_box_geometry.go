package three

//go:generate go run codegen/geometry_method_generator/main.go -typeName BoxGeometry -typeSlug box_geometry

import "syscall/js"

// BoxGeometry is the quadrilateral primitive geometry class. It is typically
// used for creating a cube or irregular quadrilateral of the dimensions
// provided with the 'width', 'height', and 'depth' constructor arguments.
type BoxGeometry struct {
	// Width          float64 `js:"width"`
	// Height         float64 `js:"height"`
	// Depth          float64 `js:"depth"`
	// WidthSegments  int     `js:"widthSegments"`
	// HeightSegments int     `js:"heightSegments"`
	// DepthSegments  int     `js:"depthSegments"`
	js.Value
}

// BoxGeometryParameters .
type BoxGeometryParameters struct {
	Width          float64
	Height         float64
	Depth          float64
	WidthSegments  int
	HeightSegments int
	DepthSegments  int
}

// NewBoxGeometry creates a new BoxGeometry.
func NewBoxGeometry(params BoxGeometryParameters) BoxGeometry {
	if params.WidthSegments == 0 {
		params.WidthSegments = 1
	}
	if params.HeightSegments == 0 {
		params.HeightSegments = 1
	}
	if params.DepthSegments == 0 {
		params.DepthSegments = 1
	}
	return BoxGeometry{
		Value: three.Get("BoxGeometry").New(
			params.Width,
			params.Height,
			params.Depth,
			params.WidthSegments,
			params.HeightSegments,
			params.DepthSegments,
		),
	}
}
