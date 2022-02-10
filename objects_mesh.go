package three

//go:generate go run object3d_method_generator/main.go -typeName Mesh -typeSlug mesh

import (
	"syscall/js"
)

type Mesh struct {
	// ID               int      `js:"id"`
	// Position         *Vector3 `js:"position"`
	// Rotation         *Euler   `js:"rotation"`
	// Geometry         Geometry `js:"geometry"`
	// Material         Material `js:"material"`
	// MatrixAutoUpdate bool     `js:"matrixAutoUpdate"`
	js.Value
}

func NewMesh(geometry Geometry, material Material) *Mesh {
	return &Mesh{
		Value: three.Get("Mesh").New(geometry.getInternalObject(), material.getInternalObject()),
	}
}

func (m Mesh) SetRotationFromAxisAngle(axis string, angle float64) {
	m.Call("setRotationFromAxisAngle", axis, angle)
}

func (m Mesh) RotateX() {
	m.Call("rotateX")
}
