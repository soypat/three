package three

//go:generate go run codegen/object3d_method_generator/main.go -typeName Scene -typeSlug scene

import "syscall/js"

// Scene - Allows you to set up what and where is to be rendered by three.js. This is where you place objects, lights and cameras.
type Scene struct {
	// AutoUpdate bool   `js:"autoUpdate"`
	// Background *Color `js:"background"`
	js.Value
}

// NewScene - Create a new Scene object.
func NewScene() *Scene {
	return &Scene{
		Value: three.Get("Scene").New(),
	}
}
