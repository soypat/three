package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName Line -typeSlug line

import "syscall/js"

// Compile-time check that this type implements Object3D interface.
var _ Object3D = &Line{}

func (obj *Line) ApplyMatrix4(matrix *Matrix4) {
	obj.Call("applyMatrix4", matrix.Value)
}

func (obj *Line) Add(m Object3D) {
	obj.Value.Call("add", m.getInternalObject())
}

func (obj *Line) Remove(m js.Value) {
	obj.Value.Call("remove", m)
}

func (obj *Line) GetObjectById(id int) js.Value {
	return obj.Call("getObjectById", id)
}

// func (obj *Line) Copy() *Line {
// 	return &Line{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *Line) ToJSON() js.Value {
	return obj.Value.Call("toJSON")
}

func (obj *Line) getInternalObject() js.Value {
	return obj.Value
}

func (obj *Line) UpdateMatrix() {
	obj.Call("updateMatrix")
}

func (obj *Line) SetPosition(v Vector3) {
	obj.Get("position").Call("copy", v.Value)
}

func (obj *Line) SetRotation(euler Euler) {
	obj.Get("rotation").Call("copy", euler.Value)
}

func (obj *Line) GetPosition() Vector3 {
	return Vector3{
		Value: obj.Get("position"),
	}
}

func (obj *Line) GetRotation() Euler {
	return Euler{
		Value: obj.Get("rotation"),
	}
}