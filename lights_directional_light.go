package three

//go:generate go run object3d_method_generator/main.go -typeName DirectionalLight -typeSlug directional_light

import "syscall/js"

// DirectionalLight - a light that gets emitted in a specific direction.
type DirectionalLight struct {
	js.Value

	// Position         *Vector3 `js:"position"`
	// MatrixAutoUpdate bool     `js:"matrixAutoUpdate"`
}

func NewDirectionalLight(color *Color, intensity float64) *DirectionalLight {
	return &DirectionalLight{
		Value: three.Get("DirectionalLight").New(color, intensity),
	}
}
