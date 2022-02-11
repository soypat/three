package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName Scene -typeSlug scene

import "syscall/js"

// Compile-time check that this type implements Object3D interface.
var _ Object3D = &Scene{}

func (obj *Scene) ApplyMatrix4(matrix *Matrix4) {
	obj.Call("applyMatrix4", matrix)
}

func (obj *Scene) Add(m Object3D) {
	obj.Value.Call("add", m)
}

func (obj *Scene) Remove(m js.Value) {
	obj.Value.Call("remove", m)
}

func (obj *Scene) GetObjectById(id int) js.Value {
	return obj.Call("getObjectById", id)
}

// func (obj *Scene) Copy() *Scene {
// 	return &Scene{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *Scene) ToJSON() js.Value {
	return obj.Value.Call("toJSON")
}

func (obj *Scene) getInternalObject() js.Value {
	return obj.Value
}

func (obj *Scene) UpdateMatrix() {
	obj.Call("updateMatrix")
}

func (obj *Scene) SetPosition(v Vector3) {
	x, y, z := v.Coords()
	obj.Get("position").Call("set", x, y, z)
}

func (obj *Scene) SetRotation(v Euler) {
	x, y, z := v.Angles()
	order := v.GetOrder()
	obj.Get("rotation").Call("set", x, y, z,order)
}

func (obj *Scene) GetPosition() Vector3 {
	return Vector3{
		Value: obj.Get("position"),
	}
}

func (obj *Scene) GetRotation() Euler {
	return Euler{
		Value: obj.Get("rotation"),
	}
}