package three

//go:generate go run object3d_method_generator/main.go -typeName Fog -typeSlug fog

import "syscall/js"

type Fog struct {
	// Color string  `js:"color"`
	// Near  float64 `js:"near"`
	// Far   float64 `js:"far"`
	js.Value
}

func NewFog(color Color, near float64, far float64) Fog {
	return Fog{
		Value: three.Get("Fog").New(color.Value, near, far),
	}
}
