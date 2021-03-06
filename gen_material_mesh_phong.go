package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// material_method_generator -typeName MeshPhongMaterial -typeSlug mesh_phong

import "syscall/js"

// Compile-time check that this type implements Material interface.
var _ Material = MeshPhongMaterial{}

func (m MeshPhongMaterial) OnBeforeCompile() {
	m.Call("onBeforeCompile")
}

func (m MeshPhongMaterial) SetValues(values *MaterialParameters) {
	m.Call("setValues", objectify(values))
}

func (m MeshPhongMaterial) ToJSON(meta interface{}) js.Value {
	return m.Call("toJSON", meta)
}

func (m MeshPhongMaterial) Clone() {
	m.Call("clone")
}

func (m MeshPhongMaterial) Copy(source Object3D) {
	m.Call("copy", source.getInternalObject())
}

func (m MeshPhongMaterial) Dispose() {
	m.Call("dispose")
}

func (m MeshPhongMaterial) getInternalObject() js.Value {
	return m.Value
}
