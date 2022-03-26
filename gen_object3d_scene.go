package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName Scene -typeSlug scene

import "syscall/js"

// Compile-time check that this type implements Object3D interface.
var _ Object3D = Scene{}

// readonly – Unique number for this object instance.
func (obj Scene) ID() int {
	return obj.Get("id").Int()
}

// Applies the matrix transform to the object and updates the object's position, rotation and scale.
func (obj Scene) ApplyMatrix4(matrix Matrix4) {
	obj.Call("applyMatrix4", matrix.Value)
}

// Applies the rotation represented by q to the object.
func (obj Scene) ApplyQuaternion(q Quaternion) {
	obj.Call("applyQuaternion", q.Value)
}

// Adds object as child of this object. An arbitrary number of objects may be added.
// Any current parent on an object passed in here will be removed, since an object can have at most one parent.
func (obj Scene) Add(m Object3D) {
	obj.Value.Call("add", m.getInternalObject())
}

// Adds m as a child of this, while maintaining m's world transform.
func (obj Scene) Attach(m Object3D) {
	obj.Value.Call("attach", m.getInternalObject())
}

// Removes m as child of this object. An arbitrary number of objects may be removed.
func (obj Scene) Remove(m Object3D) {
	obj.Value.Call("remove", m.getInternalObject())
}

// Searches through an object and its children, starting with the object itself, and returns the first with a matching id.
func (obj Scene) GetObjectById(id int) js.Value {
	return obj.Call("getObjectById", id)
}

// Copy the given m into this object. Note: event listeners and user-defined callbacks (.onAfterRender and .onBeforeRender) are not copied.
func (obj Scene) Copy(m Object3D, recursive bool) Scene {
	return Scene{Value: obj.Value.Call("copy", m.getInternalObject(), recursive)}
}

func (obj Scene) Clone(recursive bool) Scene {
	return Scene{Value: obj.Value.Call("clone", recursive)}
}

func (obj Scene) ToJSON() js.Value {
	return obj.Value.Call("toJSON")
}

func (obj Scene) getInternalObject() js.Value {
	return obj.Value
}

// Updates the local transform.
func (obj Scene) UpdateMatrix() {
	obj.Call("updateMatrix")
}

func (obj Scene) SetPosition(v Vector3) {
	obj.Get("position").Call("copy", v.Value)
}

func (obj Scene) SetRotation(euler Euler) {
	obj.Get("rotation").Call("copy", euler.Value)
}

func (obj Scene) GetPosition() Vector3 {
	return Vector3{
		Value: obj.Get("position"),
	}
}

func (obj Scene) GetRotation() Euler {
	return Euler{
		Value: obj.Get("rotation"),
	}
}

// Returns a vector representing the position of the object in world space.
func (obj Scene) GetWorldPosition(target Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("getWorldPosition", target.Value),
	}
}

// Returns a vector representing the direction of object's positive z-axis in world space.
func (obj Scene) GetWorldDirection(target Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("getWorldDirection", target.Value),
	}
}

// Returns a vector of the scaling factors applied to the object for each axis in world space.
func (obj Scene) GetWorldScale(target Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("getWorldScale", target.Value),
	}
}

// Returns a quaternion representing the rotation of the object in world space.
func (obj Scene) GetWorldQuaternion(target Quaternion) Quaternion {
	return Quaternion{
		Value: obj.Call("getWorldDirection", target.Value),
	}
}

// Rotates the object to face a point in world space.
func (obj Scene) LookAt(target Vector3)  {
	obj.Call("lookAt", target.Value)
}

// Rotates the object to face a point in world space.
func (obj Scene) LookAtCoords(x, y, z float64)  {
	obj.Call("lookAt", x, y, z)
}

// Converts vec from this object's local space to world space.
func (obj Scene) LocalToWorld(vec Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("localToWorld", vec.Value),
	}
}

// Converts vec from world space to this object's local space.
// vec represents position in world space.
func (obj Scene) WorldToLocal(vec Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("worldToLocal", vec.Value),
	}
}

// Translate an object by distance along an axis in object space. The axis is assumed to be normalized.
func (obj Scene) TranslateOnAxis(distance float64, axis Vector3) {
	obj.Call("translateOnAxis", axis.Value, distance)
}

// Translates object along x axis in object space by distance units.
func (obj Scene) TranslateX(distance float64) {
	obj.Call("translateX", distance)
}

// Translates object along y axis in object space by distance units.
func (obj Scene) TranslateY(distance float64) {
	obj.Call("translateY", distance)
}

// Translates object along z axis in object space by distance units.
func (obj Scene) TranslateZ(distance float64) {
	obj.Call("translateZ", distance)
}