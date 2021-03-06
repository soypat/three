package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// material_method_generator -typeName LineBasicMaterial -typeSlug line_basic

import "syscall/js"

// Compile-time check that this type implements Material interface.
var _ Material = LineBasicMaterial{}

func (m LineBasicMaterial) OnBeforeCompile() {
	m.Call("onBeforeCompile")
}

func (m LineBasicMaterial) SetValues(values *MaterialParameters) {
	m.Call("setValues", objectify(values))
}

func (m LineBasicMaterial) ToJSON(meta interface{}) js.Value {
	return m.Call("toJSON", meta)
}

func (m LineBasicMaterial) Clone() {
	m.Call("clone")
}

func (m LineBasicMaterial) Copy(source Object3D) {
	m.Call("copy", source.getInternalObject())
}

func (m LineBasicMaterial) Dispose() {
	m.Call("dispose")
}

func (m LineBasicMaterial) getInternalObject() js.Value {
	return m.Value
}
