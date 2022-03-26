package three

import "syscall/js"

type objecter interface {
	getInternalObject() js.Value
}

type Object3D interface {
	objecter
	ApplyMatrix4(matrix Matrix4)
	Add(Object3D)
	Remove(Object3D)
	ToJSON() js.Value
	GetObjectById(id int) js.Value
	// Copy(source Object3D, recursive bool)
	UpdateMatrix()
}

// OnBeforeRender()
// OnAfterRender()
// ApplyMatrix(matrix float64)
// ApplyQuaternion(q Quaternion)
// SetRotationFromAxisAngle(axis Vector3, angle float64)
// SetRotationFromEuler(euler Euler)
// SetRotationFromMatrix(m Matrix4)
// SetRotationFromQuaternion(q Quaternion)
// RotateOnAxis()
// RotateX()
// RotateY()
// RotateZ()
// TranslateOnAxis()
// TranslateX()
// TranslateY()
// TranslateZ()
// LocalToWorld(vector Vector3)
// WorldToLocal()
// LookAt()
// Add(object Object3D)
// Remove(object Object3D)
// GetObjectById(id string)
// GetObjectByName(name string)
// GetObjectByProperty(name, value string)
// GetWorldPosition(optionalTarget Object3D)
// GetWorldQuaternion()
// GetWorldRotation()
// GetWorldScale()
// GetWorldDirection()
// Raycast()
// UpdateMatrixWorld(force bool)
// ToJSON(meta interface{})
// Clone(recursive bool)
// Copy(source Object3D, recursive bool)
