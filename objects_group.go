//go:build TODO

package three

//go:generate go run object3d_method_generator/main.go -typeName Group -typeSlug group

import "syscall/js"

// Group is almost identical to an Object3D. Its purpose is to make working with groups of objects syntactically clearer.
type Group struct {
	*js.Object

	ID               int          `js:"id"`
	Position         *Vector3     `js:"position"`
	Rotation         *Euler       `js:"rotation"`
	Children         []*js.Object `js:"children"`
	MatrixAutoUpdate bool         `js:"matrixAutoUpdate"`
	Visible          bool         `js:"visible"`
}

func NewGroup() *Group {
	return &Group{
		Object: three.Get("Group").New(),
	}
}
