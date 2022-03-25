package three

//go:generate go run codegen/object3d_method_generator/main.go -typeName PolarGridHelper -typeSlug polargrid_helper

import "syscall/js"

type PolarGridHelper struct {
	// Copied from LineSegments struct:

	// ID               int             `js:"id"`
	// Position         *Vector3        `js:"position"`
	// Rotation         *Euler          `js:"rotation"`
	// Geometry         *BufferGeometry `js:"geometry"`
	// Material         Material        `js:"material"`
	// MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	// RenderOrder      int             `js:"renderOrder"`
	// Visible          bool            `js:"visible"`
	js.Value
}

type PolarGridHelperParameters struct {
	// The radius of the polar grid. This can be any positive number. Default is 10.
	Radius float64
	// The number of radial lines. This can be any positive integer. Default is 16.
	Radials int
	// The number of circles. This can be any positive integer. Default is 8.
	Circles int
	// The number of line segments used for each circle.
	// This can be any positive integer that is 3 or greater. Default is 64.
	Divisions int
	// The first color used for grid elements.
	Color1 *Color
	// The second color used for grid elements.
	Color2 *Color
}

func NewPolarGridHelper(params PolarGridHelperParameters) PolarGridHelper {
	if params.Radius == 0 {
		params.Radius = 10
	}
	if params.Radials == 0 {
		params.Radius = 16
	}
	if params.Circles == 0 {
		params.Circles = 8
	}
	if params.Divisions < 3 {
		params.Circles = 64
	}
	if params.Color1 == nil {
		params.Color1 = NewColorHex(0x444444)
	}
	if params.Color2 == nil {
		params.Color2 = NewColorHex(0x888888)
	}
	return PolarGridHelper{
		Value: three.Get("PolarGridHelper").New(
			params.Radius,
			params.Radials,
			params.Circles,
			params.Divisions,
			params.Color1,
			params.Color2,
		),
	}
}
