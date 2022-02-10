//go:build TODO

package three

//go:generate go run geometry_method_generator/main.go -geometryType ConeGeometry -geometrySlug cone_geometry

import (
	"math"

	"syscall/js"
)

// ConeGeometry a class for generating Cone geometries.
type ConeGeometry struct {
	*js.Object

	Radius         float64 `js:"radius"`
	Height         float64 `js:"height"`
	RadialSegments int     `js:"radialSegments"`
	HeightSegments int     `js:"heightSegments"`
	OpenEnded      bool    `js:"openEnded"`
	ThetaStart     float64 `js:"thetaStart"`
	ThetaLength    float64 `js:"thetaLength"`
}

// ConeGeometryParameters .
type ConeGeometryParameters struct {
	// Radius of the cone base.
	Radius float64
	// Height of the Cone. Default is 1.
	Height float64
	// Number of segmented faces around the circumference of the Cone. Default is 8
	RadialSegments int
	// Number of rows of faces along the height of the Cone. Default is 1.
	HeightSegments int
	// A Boolean indicating whether the ends of the Cone are open or capped. Default is false, meaning capped.
	OpenEnded bool
	// Start angle for first segment, default = 0 (three o'clock position).
	ThetaStart float64
	//  The central angle, often called theta, of the circular sector. The default is 2*Pi, which makes for a complete Cone.
	ThetaLength float64
}

// NewConeGeometry creates a new BoxGeometry. Set ThetaLength to NaN to create empty geometry.
func NewConeGeometry(params *ConeGeometryParameters) ConeGeometry {
	if params == nil {
		params = &ConeGeometryParameters{}
	}
	// Make sure both are defined to prevent unclear code
	if params.Height == 0 || params.Radius == 0 {
		params.Height = 1
		params.Radius = 1
	}
	// Probably don't want no Cone.
	if params.ThetaLength == 0 {
		params.ThetaLength = 2 * math.Pi
	}
	if math.IsNaN(params.ThetaLength) {
		params.ThetaLength = 0
	}
	if params.RadialSegments == 0 {
		params.RadialSegments = 8
	}
	if params.HeightSegments == 0 {
		params.RadialSegments = 1
	}
	return ConeGeometry{
		Object: three.Get("ConeGeometry").New(
			params.Radius,
			params.Height,
			params.RadialSegments,
			params.HeightSegments,
			params.OpenEnded,
			params.ThetaStart,
			params.ThetaLength,
		),
	}
}
