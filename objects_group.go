package three

//go:generate go run codegen/object3d_method_generator/main.go -typeName Group -typeSlug group

import "syscall/js"

// Group is almost identical to an Object3D. Its purpose is to make working with groups of objects syntactically clearer.
type Group struct {
	// ID               int          `js:"id"`
	// Position         *Vector3     `js:"position"`
	// Rotation         *Euler       `js:"rotation"`
	// Children         []*js.Object `js:"children"`
	// MatrixAutoUpdate bool         `js:"matrixAutoUpdate"`
	// Visible          bool         `js:"visible"`
	js.Value
}

func NewGroup() Group {
	return Group{
		Value: three.Get("Group").New(),
	}
}
