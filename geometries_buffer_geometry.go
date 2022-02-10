package three

//go:generate go run geometry_method_generator/main.go -geometryType BufferGeometry -geometrySlug buffer_geometry

import "syscall/js"

// BufferGeometry is an efficient representation of mesh, line, or point geometry. Includes vertex positions, face
// indices, normals, colors, UVs, and custom attributes within buffers, reducing the cost of passing all this
// data to the GPU.
// It's a counterpart of three.js' core/BufferGeometry object.
type BufferGeometry struct {
	// Vertices   []Vector3        `js:"vertices"`
	// Attributes *BufferAttribute `js:"attributes"`
	js.Value
}

// NewBufferGeometry creates a new Buffer Geometry.
func NewBufferGeometry() *BufferGeometry {
	return &BufferGeometry{
		Value: three.Get("BufferGeometry").New(),
	}
}

// AddVertex adds new vertex to the geometry, specified by its coordinates.
// func (bg *BufferGeometry) AddVertex(x, y, z float64) {
// 	vec := NewVector3(x, y, z)
// 	bg.Vertices = append(bg.Vertices, vec)
// }

// AddVertices adds new vertices to the geometry.
// func (bg *BufferGeometry) AddVertices(v ...Vector3) {
// 	bg.Vertices = append(bg.Vertices, v...)
// }

// AddFace adds new Face3 (triangle) to the geometry, specified by its vertice indicies.
func (bg *BufferGeometry) AddFace(a, b, c int) {
	face := NewFace3(float64(a), float64(b), float64(c))
	bg.Get("faces").Call("push", face)
}

// AddAttribute adds a new attribute like 'position' to the BufferGeometry.
func (bg *BufferGeometry) SetAttribute(name string, attr *BufferAttribute) {
	bg.Call("setAttribute", name, attr)
}

// GetAttribute retruns BufferGeometry's attribute by name (should be added first by
// AddAttribute call, see threejs docs for explanations.
func (bg *BufferGeometry) GetAttribute(name string) *BufferAttribute {
	obj := bg.Call("getAttribute", name)
	return &BufferAttribute{
		Value: obj,
	}
}
