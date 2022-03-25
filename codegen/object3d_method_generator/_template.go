package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName {{ .Type }} -typeSlug {{ .Slug }}

import "syscall/js"

// Compile-time check that this type implements Object3D interface.
var _ Object3D = {{ .Type }}{}

func (obj {{ .Type }}) ApplyMatrix4(matrix Matrix4) {
	obj.Call("applyMatrix4", matrix.Value)
}

func (obj {{ .Type }}) Add(m Object3D) {
	obj.Value.Call("add", m.getInternalObject())
}

func (obj {{ .Type }}) Remove(m js.Value) {
	obj.Value.Call("remove", m)
}

func (obj {{ .Type }}) GetObjectById(id int) js.Value {
	return obj.Call("getObjectById", id)
}

// func (obj {{ .Type }}) Copy() {{ .Type }} {
// 	return {{ .Type }}{Object: obj.getInternalObject().Call("copy")}
// }

func (obj {{ .Type }}) ToJSON() js.Value {
	return obj.Value.Call("toJSON")
}

func (obj {{ .Type }}) getInternalObject() js.Value {
	return obj.Value
}

func (obj {{ .Type }}) UpdateMatrix() {
	obj.Call("updateMatrix")
}

func (obj {{ .Type }}) SetPosition(v Vector3) {
	obj.Get("position").Call("copy", v.Value)
}

func (obj {{ .Type }}) SetRotation(euler Euler) {
	obj.Get("rotation").Call("copy", euler.Value)
}

func (obj {{ .Type }}) GetPosition() Vector3 {
	return Vector3{
		Value: obj.Get("position"),
	}
}

func (obj {{ .Type }}) GetRotation() Euler {
	return Euler{
		Value: obj.Get("rotation"),
	}
}