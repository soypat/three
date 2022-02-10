//go:build TODO

package three

//go:generate go run geometry_method_generator/main.go -geometryType CylinderGeometry -geometrySlug cylinder_geometry

import (
	"math"

	"syscall/js"
)

// CylinderGeometry a class for generating cylinder geometries.
type CylinderGeometry struct {
	*js.Object

	RadiusTop      float64 `js:"radiusTop"`
	RadiusBottom   float64 `js:"radiusBottom"`
	Height         float64 `js:"height"`
	RadialSegments int     `js:"radialSegments"`
	HeightSegments int     `js:"heightSegments"`
	OpenEnded      bool    `js:"openEnded"`
	ThetaStart     float64 `js:"thetaStart"`
	ThetaLength    float64 `js:"thetaLength"`
}

// CylinderGeometryParameters .
type CylinderGeometryParameters struct {
	RadiusTop float64
	// Radius of the cylinder at the bottom. Default is 1.
	RadiusBottom float64
	// Height of the cylinder. Default is 1.
	Height float64
	// Number of segmented faces around the circumference of the cylinder. Default is 8
	RadialSegments int
	// Number of rows of faces along the height of the cylinder. Default is 1.
	HeightSegments int
	// A Boolean indicating whether the ends of the cylinder are open or capped. Default is false, meaning capped.
	OpenEnded bool
	// Start angle for first segment, default = 0 (three o'clock position).
	ThetaStart float64
	//  The central angle, often called theta, of the circular sector. The default is 2*Pi, which makes for a complete cylinder.
	ThetaLength float64
}

// NewCylinderGeometry creates a new BoxGeometry. Set ThetaLength to NaN to create empty geometry.
func NewCylinderGeometry(params *CylinderGeometryParameters) CylinderGeometry {
	if params == nil {
		params = &CylinderGeometryParameters{}
	}
	// Make sure all are defined to prevent unclear code. This could be changed though...
	if params.Height == 0 || params.RadiusTop == 0 || params.RadiusBottom == 0 {
		params.Height = 1
		params.RadiusTop = 1
		params.RadiusBottom = 1
	}
	// Probably don't want no cylinder.
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
		params.HeightSegments = 1
	}
	obj := three.Get("CylinderGeometry").New(
		params.RadiusTop,
		params.RadiusBottom,
		params.Height,
		params.RadialSegments,
		params.HeightSegments,
		params.OpenEnded,
		params.ThetaStart,
		params.ThetaLength,
	)
	// js.Global().Get("console").Call("log", obj.Interface(), params)
	return CylinderGeometry{
		Object: obj,
	}
}
