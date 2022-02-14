package three

import (
	"syscall/js"
)

// BufferAttribute stores data for an attribute (such as vertex positions, face indices, normals, colors, UVs, and any custom attributes) associated with a BufferGeometry, which allows for more efficient passing of data to the GPU.
type BufferAttribute struct {
	// Array       []int `js:"array"`
	// Count       int   `js:"count"`
	// IsDynamic   bool  `js:"dynamic"`
	// ItemSize    int   `js:"itemSize"`
	// NeedsUpdate bool  `js:"needsUpdate"`
	// Normalized  bool  `js:"normalized"`
	js.Value
}

// NewBufferAttribute creates a new BufferAttribute
func NewBufferAttribute(data []float32, itemSize int) *BufferAttribute {
	return &BufferAttribute{
		Value: three.Get("BufferAttribute").New(float32ToArray(data), itemSize),
	}
}

// SetXYZ sets the x, y and z components of the vector at the given index.
func (b BufferAttribute) SetXYZ(i int, x, y, z float64) {
	b.Call("setXYZ", i, float32(x), float32(y), float32(z))
}

func (b BufferAttribute) GetX(i int) float64 {
	return b.Call("getX", i).Float()
}
