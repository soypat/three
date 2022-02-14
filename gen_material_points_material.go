package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// material_method_generator -materialName PointsMaterial -materialSlug points_material

import "syscall/js"

// Compile-time check that this type implements Material interface.
var _ Material = PointsMaterial{}

func (m PointsMaterial) OnBeforeCompile() {
	m.Call("onBeforeCompile")
}

func (m PointsMaterial) SetValues(values MaterialParameters) {
	m.Call("setValues", objectify(values))
}

func (m PointsMaterial) ToJSON(meta interface{}) js.Value {
	return m.Call("toJSON", meta)
}

func (m PointsMaterial) Clone() {
	m.Call("clone")
}

func (m PointsMaterial) Copy(source Object3D) {
	m.Call("copy", source)
}

func (m PointsMaterial) Dispose() {
	m.Call("dispose")
}

func (m PointsMaterial) getInternalObject() js.Value {
	return m.Value
}