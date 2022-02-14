package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName AxesHelper -typeSlug axes_helper

import "syscall/js"

// Compile-time check that this type implements Object3D interface.
var _ Object3D = &AxesHelper{}

func (obj *AxesHelper) ApplyMatrix4(matrix *Matrix4) {
	obj.Call("applyMatrix4", matrix.Value)
}

func (obj *AxesHelper) Add(m Object3D) {
	obj.Value.Call("add", m.getInternalObject())
}

func (obj *AxesHelper) Remove(m js.Value) {
	obj.Value.Call("remove", m)
}

func (obj *AxesHelper) GetObjectById(id int) js.Value {
	return obj.Call("getObjectById", id)
}

// func (obj *AxesHelper) Copy() *AxesHelper {
// 	return &AxesHelper{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *AxesHelper) ToJSON() js.Value {
	return obj.Value.Call("toJSON")
}

func (obj *AxesHelper) getInternalObject() js.Value {
	return obj.Value
}

func (obj *AxesHelper) UpdateMatrix() {
	obj.Call("updateMatrix")
}

func (obj *AxesHelper) SetPosition(v Vector3) {
	obj.Get("position").Call("copy", v.Value)
}

func (obj *AxesHelper) SetRotation(euler Euler) {
	obj.Get("rotation").Call("copy", euler.Value)
}

func (obj *AxesHelper) GetPosition() Vector3 {
	return Vector3{
		Value: obj.Get("position"),
	}
}

func (obj *AxesHelper) GetRotation() Euler {
	return Euler{
		Value: obj.Get("rotation"),
	}
}