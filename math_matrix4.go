package three

import "syscall/js"

// Matrix4 - represents a Matrix4.
type Matrix4 struct {
	// A column-major list of matrix values.
	// A *js.Object `js:"A"`
	js.Value
}

func NewMatrix4() *Matrix4 {
	return &Matrix4{
		Value: three.Get("Matrix4").New(),
	}
}

// Sets this matrix to the transformation composed of position, quaternion and scale.
func (m *Matrix4) Compose(position Vector3, q Quaternion, scale Vector3) (this *Matrix4) {
	m.Value.Call("compose", position.Value, q.Value, scale.Value)
	return m
}

// Decomposes this matrix into its position, quaternion and scale components.
// Note: Not all matrices are decomposable in this way. For example, if an object has a non-uniformly scaled parent, then the object's world matrix may not be decomposable, and this method may not be appropriate.
func (m *Matrix4) Decompose(position Vector3, q Quaternion, scale Vector3) {
	m.Value.Call("decompose", position.Value, q.Value, scale.Value)
}

// http://www.euclideanspace.com/maths/algebra/matrix/functions/inverse/fourD/index.htm
func (m *Matrix4) Determinant() float64 {
	return m.Value.Call("determinant").Float()
}

// Extracts the basis of this matrix into the three axis vectors provided. If this matrix is:
//  a, b, c, d,
//  e, f, g, h,
//  i, j, k, l,
//  m, n, o, p
// then x,y,z will be set:
//  x = (a, e, i)
//  y = (b, f, j)
//  z = (c, g, k)
func (m *Matrix4) ExtractBasis(x, y, z Vector3) (this *Matrix4) {
	m.Value.Call("extractBasis", x.Value, y.Value, z.Value)
	return m
}

func (m *Matrix4) ExtractRotation(a *Matrix4) (this *Matrix4) {
	m.Value.Call("extractRotation", a.Value)
	return m
}

// Sets the rotation component of this matrix to the rotation specified by q, as outlined here.
// The rest of the matrix is set to the identity. So, given q = w + xi + yj + zk, the resulting matrix will be:
//  1-2y²-2z²    2xy-2zw    2xz+2yw    0
//  2xy+2zw      1-2x²-2z²  2yz-2xw    0
//  2xz-2yw      2yz+2xw    1-2x²-2y²  0
//  0            0          0          1
func (m *Matrix4) MakeRotationFromQuaternion(q Quaternion) (this *Matrix4) {
	m.Value.Call("makeRotationFromQuaternion", q)
	return m
}

// Sets this matrix as a translation transform:
//  1, 0, 0, x,
//  0, 1, 0, y,
//  0, 0, 1, z,
//  0, 0, 0, 1
func (m *Matrix4) MakeTranslation(x, y, z float64) (this *Matrix4) {
	m.Value.Call("makeTranslation", x, y, z)
	return m
}

// Set this to the basis matrix consisting of the three provided basis vectors:
//  xAxis.x, yAxis.x, zAxis.x, 0,
//  xAxis.y, yAxis.y, zAxis.y, 0,
//  xAxis.z, yAxis.z, zAxis.z, 0,
//  0,       0,       0,       1
func (m *Matrix4) MakeBasis(x, y, z Vector3) (this *Matrix4) {
	m.Value.Call("makeBasis", x.Value, y.Value, z.Value)
	return m
}

// Inverts this matrix, using the analytic method. You can not invert with a determinant of zero. If you attempt this, the method produces a zero matrix instead.
func (m *Matrix4) Invert() (this *Matrix4) {
	m.Value.Call("invert")
	return m
}

// Resets this matrix to the identity matrix.
func (m *Matrix4) Identity() (this *Matrix4) {
	m.Value.Call("identity")
	return m
}
