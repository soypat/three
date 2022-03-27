package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName Fog -typeSlug fog

import "syscall/js"

// Compile-time check that this type implements Object3D interface.
var _ Object3D = Fog{}

// readonly – Unique number for this object instance.
func (obj Fog) ID() int {
	return obj.Get("id").Int()
}

// ApplyMatrix4 applies the matrix transform to the object and updates the object's position, rotation and scale.
func (obj Fog) ApplyMatrix4(matrix Matrix4) {
	obj.Call("applyMatrix4", matrix.Value)
}

// ApplyQuaternion applies the rotation represented by q to the object.
func (obj Fog) ApplyQuaternion(q Quaternion) {
	obj.Call("applyQuaternion", q.Value)
}

// Add object as child of this object. An arbitrary number of objects may be added.
// Any current parent on an object passed in here will be removed, since an object can have at most one parent.
func (obj Fog) Add(m Object3D) {
	obj.Value.Call("add", m.getInternalObject())
}

// Attach adds m as a child of this, while maintaining m's world transform.
func (obj Fog) Attach(m Object3D) {
	obj.Value.Call("attach", m.getInternalObject())
}

// Clear removes all child objects.
func (obj Fog) Clear() {
	obj.Value.Call("clear")
}

// Remove m as child of this object. An arbitrary number of objects may be removed.
func (obj Fog) Remove(m Object3D) {
	obj.Value.Call("remove", m.getInternalObject())
}

// RemoveFromParent removes this object from its current parent.
func (obj Fog) RemoveFromParent() {
	obj.Value.Call("removeFromParent")
}

// GetObjectById searches through an object and its children, starting with the object itself, and returns the first with a matching id.
func (obj Fog) GetObjectById(id int) js.Value {
	return obj.Call("getObjectById", id)
}

// Copy the given m into this object. Note: event listeners and user-defined callbacks (.onAfterRender and .onBeforeRender) are not copied.
func (obj Fog) Copy(m Object3D, recursive bool) Fog {
	return Fog{Value: obj.Value.Call("copy", m.getInternalObject(), recursive)}
}

func (obj Fog) Clone(recursive bool) Fog {
	return Fog{Value: obj.Value.Call("clone", recursive)}
}

func (obj Fog) ToJSON() js.Value {
	return obj.Value.Call("toJSON")
}

func (obj Fog) getInternalObject() js.Value {
	return obj.Value
}

// Updates the local transform.
func (obj Fog) UpdateMatrix() {
	obj.Call("updateMatrix")
}

// Rotate an object along an axis in object space. The axis is assumed to be normalized.
// axis is a normalized vector in object space.
func (obj Fog) RotateOnAxis(angle float64, axis Vector3) (this Fog ){
	obj.Call("rotateOnAxis", axis.Value, angle)
	return obj
}

// Rotate an object along an axis in world space. The axis is assumed to be normalized. Method Assumes no rotated parent.
// axis is a normalized vector in world space.
func (obj Fog) RotateOnWorldAxis(angle float64, axis Vector3) (this Fog ){
	obj.Call("rotateOnAxis", axis.Value, angle)
	return obj
}

// Calls setFromAxisAngle( axis, angle ) on the .quaternion.
func (obj Fog) SetRotationFromAxisAngle(angle float64, axis Vector3) {
	obj.Call("setRotationFromAxisAngle", axis.Value, angle)
}

// Calls setRotationFromEuler(euler) on the .quaternion.
func (obj Fog) SetRotationFromEuler(euler Euler) {
	obj.Call("setRotationFromEuler", euler.Value)
}

// Calls setFromRotationMatrix(m) on the .quaternion.
func (obj Fog) SetRotationFromMatrix(m Matrix4) {
	obj.Call("setRotationFromMatrix", m.Value)
}

// Copy the given quaternion into .quaternion.
func (obj Fog) SetRotationFromQuaternion(q Quaternion) {
	obj.Call("setRotationFromQuaternion", q.Value)
}

func (obj Fog) SetPosition(v Vector3) {
	obj.Get("position").Call("copy", v.Value)
}

func (obj Fog) GetPosition() Vector3 {
	return Vector3{
		Value: obj.Get("position"),
	}
}

func (obj Fog) GetRotation() Euler {
	return Euler{
		Value: obj.Get("rotation"),
	}
}

// Returns a vector representing the position of the object in world space.
// The result is copied into dst.
func (obj Fog) GetWorldPosition(dst Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("getWorldPosition", dst.Value),
	}
}

// Returns a vector representing the direction of object's positive z-axis in world space.
// The result is copied into dst.
func (obj Fog) GetWorldDirection(dst Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("getWorldDirection", dst.Value),
	}
}

// Returns a vector of the scaling factors applied to the object for each axis in world space.
// The result is copied into dst.
func (obj Fog) GetWorldScale(dst Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("getWorldScale", dst.Value),
	}
}

// Returns a quaternion representing the rotation of the object in world space.
// The result is copied into dst.
func (obj Fog) GetWorldQuaternion(dst Quaternion) Quaternion {
	return Quaternion{
		Value: obj.Call("getWorldDirection", dst.Value),
	}
}

// Rotates the object to face a point in world space.
func (obj Fog) LookAt(target Vector3)  {
	obj.Call("lookAt", target.Value)
}

// Rotates the object to face a point in world space.
func (obj Fog) LookAtCoords(x, y, z float64)  {
	obj.Call("lookAt", x, y, z)
}

// Converts vec from this object's local space to world space.
func (obj Fog) LocalToWorld(vec Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("localToWorld", vec.Value),
	}
}

// Converts vec from world space to this object's local space.
// vec represents position in world space.
func (obj Fog) WorldToLocal(vec Vector3) Vector3 {
	return Vector3{
		Value: obj.Call("worldToLocal", vec.Value),
	}
}

// Translate an object by distance along an axis in object space. The axis is assumed to be normalized.
func (obj Fog) TranslateOnAxis(distance float64, axis Vector3) {
	obj.Call("translateOnAxis", axis.Value, distance)
}

// Translates object along x axis in object space by distance units.
func (obj Fog) TranslateX(distance float64) {
	obj.Call("translateX", distance)
}

// Translates object along y axis in object space by distance units.
func (obj Fog) TranslateY(distance float64) {
	obj.Call("translateY", distance)
}

// Translates object along z axis in object space by distance units.
func (obj Fog) TranslateZ(distance float64) {
	obj.Call("translateZ", distance)
}
