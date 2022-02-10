package three

//go:generate go run object3d_method_generator/main.go -typeName AmbientLight -typeSlug ambient_light

import "syscall/js"

// AmbientLight - a light that gets emitted in a specific direction.
type AmbientLight struct {
	js.Value

	// Position         *Vector3 `js:"position"`
	// MatrixAutoUpdate bool     `js:"matrixAutoUpdate"`
}

func NewAmbientLight(color *Color, intensity float64) *AmbientLight {
	return &AmbientLight{
		Value: three.Get("AmbientLight").New(color, intensity),
	}
}

func (l *AmbientLight) SetPosition(v Vector3) {
	l.Set("position", v.Value)
}
