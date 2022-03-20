package three

import "syscall/js"

// Quaternion - represents a Quaternion.
type Quaternion struct {
	// X float64 `js:"x"`
	// Y float64 `js:"y"`
	// Z float64 `js:"z"`
	// W float64 `js:"w"`
	js.Value
}

func NewQuaternion() Quaternion {
	return Quaternion{
		Value: three.Get("Quaternion").New(),
	}
}

// Returns the angle between this quaternion and quaternion a in radians.
func (q Quaternion) AngleTo(a Quaternion) float64 {
	return q.Value.Call("angleTo", a.Value).Float()
}

// Returns the rotational conjugate of this quaternion.
// The conjugate of a quaternion represents the same rotation in the opposite direction about the rotational axis.
func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{Value: q.Value.Call("conjugate")}
}

func (q Quaternion) Copy(a Quaternion) Quaternion {
	return Quaternion{Value: q.Value.Call("copy", a.Value)}
}

// Compares the x, y, z and w properties of a to the equivalent properties of this quaternion to determine if they represent the same rotation.
func (q Quaternion) Equals(a Quaternion) bool {
	return q.Value.Call("equals", a.Value).Bool()
}

// Calculates the dot product of quaternions a and receiver.
func (q Quaternion) Dot(a Quaternion) float64 {
	return q.Value.Call("dot", a.Value).Float()
}

// Sets this quaternion to the identity quaternion; that is, to the quaternion that represents "no rotation".
func (q Quaternion) Identity() Quaternion {
	return Quaternion{Value: q.Value.Call("identity")}
}

// Inverts this quaternion - calculates the conjugate. The quaternion is assumed to have unit length.
func (q Quaternion) Invert() Quaternion {
	return Quaternion{Value: q.Value.Call("invert")}
}

// Length Computes the Euclidean length (straight-line length) of this quaternion, considered as a 4 dimensional vector.
func (q Quaternion) Length() float64 {
	return q.Value.Call("length").Float()
}

// LengthSq Computes the squared Euclidean length (straight-line length) of this quaternion, considered as a 4 dimensional vector.
// This can be useful if you are comparing the lengths of two quaternions, as this is a slightly more efficient calculation than length().
func (q Quaternion) LengthSq() float64 {
	return q.Value.Call("lengthSq").Float()
}

// Normalize this quaternion - that is, calculated the quaternion that performs the same rotation as this one, but has length equal to 1.
func (q Quaternion) Normalize() Quaternion {
	return Quaternion{Value: q.Value.Call("normalize")}
}

// Multiply this quaternion by a.
func (q Quaternion) Multiply(a Quaternion) Quaternion {
	return Quaternion{Value: q.Value.Call("multiply", a.Value)}
}

// MultiplyQuaternions Sets this quaternion to a x b. Adapted from the method outlined here http://www.euclideanspace.com/maths/algebra/realNormedAlgebra/quaternions/code/index.htm
func (q Quaternion) MultiplyQuaternions(a, b Quaternion) Quaternion {
	return Quaternion{Value: q.Value.Call("multiplyQuaternions", a.Value, b.Value)}
}

// Premultiply Pre-multiplies this quaternion by a.
func (q Quaternion) Premultiply(a Quaternion) Quaternion {
	return Quaternion{Value: q.Value.Call("premultiply", a.Value)}
}

// Random Sets this quaternion to a uniformly random, normalized quaternion.
func (q Quaternion) Random() Quaternion {
	return Quaternion{Value: q.Value.Call("random")}
}

// RotateTowards Rotates this quaternion by a given angular step to the defined quaternion a.
// The method ensures that the final quaternion will not overshoot a.
func (q Quaternion) RotateTowards(step float64, a Quaternion) Quaternion {
	return Quaternion{Value: q.Value.Call("rotateTowards", a.Value, step)}
}

// Slerp Handles the spherical linear interpolation between quaternions.
// t represents the amount of rotation between this quaternion (where t is 0) and qb (where t is 1).
// This quaternion is set to the result. Also see the static version of the slerp below.
func (q Quaternion) Slerp(t float64, qb Quaternion) Quaternion {
	return Quaternion{Value: q.Value.Call("slerp", qb.Value, t)}
}

// SlerpQuaternions Performs a spherical linear interpolation between the given quaternions and stores the result in this quaternion.
func (q Quaternion) SlerpQuaternions(t float64, qa, qb Quaternion) (this Quaternion) {
	q.Value.Call("slerpQuaternions", qa.Value, qb.Value, t)
	return q
}

// Set x, y, z, w properties of this quaternion.
func (q Quaternion) Set(x, y, z, w float64) Quaternion {
	return Quaternion{
		Value: q.Value.Call("set", x, y, z, w),
	}
}

// SetFromAxisAngle sets this quaternion from rotation specified by axis and angle.
// Adapted from the method here http://www.euclideanspace.com/maths/geometry/rotations/conversions/angleToQuaternion/index.htm
// Axis is assumed to be normalized, angle is in radians.
func (q Quaternion) SetFromAxisAngle(angle float64, axis Vector3) Quaternion {
	return Quaternion{
		Value: q.Value.Call("setFromAxisAngle", axis.Value, angle),
	}
}

// SetFromEuler sets this quaternion from the rotation specified by Euler angle.
func (q Quaternion) SetFromEuler(euler Euler) Quaternion {
	return Quaternion{
		Value: q.Value.Call("setFromEuler", euler.Value),
	}
}

// SetFromRotationMatrix sets this quaternion from rotation component of m.
// Adapted from the method here: http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/index.htm
// m - a Matrix4 of which the upper 3x3 of matrix is a pure rotation matrix (i.e. unscaled).
func (q Quaternion) SetFromRotationMatrix(m Matrix4) Quaternion {
	return Quaternion{
		Value: q.Value.Call("setFromRotationMatrix", m.Value),
	}
}

// Sets this quaternion to the rotation required to rotate direction vector vFrom to direction vector vTo.
// Adapted from the method here: http://lolengine.net/blog/2013/09/18/beautiful-maths-quaternion-from-vectors
// vFrom and vTo are assumed to be normalized.
func (q Quaternion) SetFromUnitVectors(vFrom, vTo Vector3) Quaternion {
	// V1
	// 	quat quat::fromtwovectors(vec3 u, vec3 v)
	// {
	//     vec3 w = cross(u, v);
	//     quat q = quat(dot(u, v), w.x, w.y, w.z);
	//     q.w += length(q);
	//     return normalize(q);
	// }
	// V2
	// 	quat quat::fromtwovectors(vec3 u, vec3 v)
	// {
	//     float m = sqrt(2.f + 2.f * dot(u, v));
	//     vec3 w = (1.f / m) * cross(u, v);
	//     return quat(0.5f * m, w.x, w.y, w.z);
	// }
	return Quaternion{
		Value: q.Value.Call("setFromUnitVectors", vFrom.Value, vTo.Value),
	}
}
