package three

import "syscall/js"

// Vector3 - represents a Vector3.
type Vector3 struct {
	// X float64 `js:"x"`
	// Y float64 `js:"y"`
	// Z float64 `js:"z"`
	js.Value
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{
		Value: three.Get("Vector3").New(x, y, z),
	}
}

// Add a to this vector.
func (v Vector3) Add(a Vector3) (this Vector3) {
	v.Call("add", a.Value)
	return v
}

// AddScalar Adds the scalar value f to this vector's x, y and z values.
func (v Vector3) AddScalar(f float64) (this Vector3) {
	v.Call("addScalar", f)
	return v
}

// AddScaledVector Adds the multiple of a and f to this vector. this = f*a
func (v Vector3) AddScaledVector(f float64, a Vector3) (this Vector3) {
	v.Call("addScaledVector", a.Value, f)
	return v
}

// ApplyAxisAngle Applies a rotation specified by a normalized axis and an angle to this vector.
func (v Vector3) ApplyAxisAngle(angle float64, axis Vector3) (this Vector3) {
	v.Call("applyAxisAngle", axis.Value, angle)
	return v
}

// ApplyEuler transform to this vector by converting the Euler object to a Quaternion and applying.
func (v Vector3) ApplyEuler(euler Euler) (this Vector3) {
	v.Call("applyEuler", euler.Value)
	return v
}

// ApplyQuaternion transform to this vector.
func (v Vector3) ApplyQuaternion(q Quaternion) (this Vector3) {
	v.Call("applyQuaternion", q.Value)
	return v
}

// AngleTo Returns the angle between this vector and vector a in radians.
func (v Vector3) AngleTo(a Vector3) (this Vector3) {
	v.Call("angleTo", a.Value)
	return v
}

// Ceil The x, y and z components of this vector are rounded up to the nearest integer value.
func (v Vector3) Ceil() (this Vector3) {
	v.Call("ceil")
	return v
}

// ClampLength If this vector's length is greater than the max value, the vector will be scaled down so its length is the max value.
//
// If this vector's length is less than the min value, the vector will be scaled up so its length is the min value.
func (v Vector3) ClampLength(min, max float64) (this Vector3) {
	v.Call("clampLength", min, max)
	return v
}

// ClampScalar clamps individual components of this vector. If this vector's x, y or z values are greater than the max value, they are replaced by the max value.
//
// If this vector's x, y or z values are less than the min value, they are replaced by the min value.
func (v Vector3) ClampScalar(min, max float64) (this Vector3) {
	v.Call("clampScalar", min, max)
	return v
}

// Cross Sets this vector to cross product of itself and a.
func (v Vector3) Cross(a Vector3) (this Vector3) {
	v.Call("cross", a.Value)
	return v
}

// CrossVectors Sets this vector to cross product of a and b.
func (v Vector3) CrossVectors(a, b Vector3) (this Vector3) {
	v.Call("cross", a.Value, b.Value)
	return v
}

// DistanceTo Computes the distance from this vector to v.
func (v Vector3) DistanceTo(a Vector3) float64 {
	return v.Call("distanceTo", a.Value).Float()
}

// ManhattanDistanceTo Computes the manhattan distance from this vector to v. https://en.wikipedia.org/wiki/Taxicab_geometry
func (v Vector3) ManhattanDistanceTo(a Vector3) float64 {
	return v.Call("manhattanDistanceTo", a.Value).Float()
}

// DistanceToSquared Computes the squared distance from this vector to v. If you are just comparing the distance
// with another distance, you should compare the distance squared instead as it is slightly more efficient to calculate.
func (v Vector3) DistanceToSquared(a Vector3) float64 {
	return v.Call("distanceToSquared", a.Value).Float()
}

// Divide this vector by a.
func (v Vector3) Divide(a Vector3) (this Vector3) {
	v.Call("divide", a.Value)
	return v
}

// DivideScalar this vector by scalar f.
func (v Vector3) DivideScalar(f float64) (this Vector3) {
	v.Call("divideScalar", f)
	return v
}

// Dot Calculate the dot product of this vector and v.
func (v Vector3) Dot(a Vector3) (this Vector3) {
	v.Call("dot", a.Value)
	return v
}

// Equals Returns true if the components of this vector and v are strictly equal; false otherwise.
func (v Vector3) Equals(a Vector3) bool {
	return v.Call("equals", a.Value).Bool()
}

// Floor The components of this vector are rounded down to the nearest integer value.
func (v Vector3) Floor() (this Vector3) {
	v.Call("floor")
	return v
}

// GetComponent gets X value if idx=0. Gets Y if idx=1. Gets Z if idx=2.
func (v Vector3) GetComponent(idx int) float64 {
	return v.Call("getComponent", idx).Float()
}

// Length Computes the Euclidean length (straight-line length) from (0, 0, 0) to (x, y, z).
func (v Vector3) Length() float64 {
	return v.Call("length").Float()
}

// ManhattanLength Computes the Manhattan length of this vector
func (v Vector3) ManhattanLength() float64 {
	return v.Call("manhattanLength").Float()
}

// Computes the square of the Euclidean length (straight-line length) from (0, 0, 0) to (x, y, z). If you are comparing the lengths of vectors,
// you should compare the length squared instead as it is slightly more efficient to calculate.
func (v Vector3) LengthSq() float64 {
	return v.Call("lengthSq").Float()
}

// Lerp Linearly interpolate between this vector and v, where alpha is the percent distance along the line - alpha = 0 will be this vector, and alpha = 1 will be a.
func (v Vector3) Lerp(f float64, a Vector3) (this Vector3) {
	v.Call("lerp", a.Value, f)
	return v
}

// LerpVectors Sets this vector to be the vector linearly interpolated between a and b where alpha is the percent distance along the line connecting the two vectors - alpha = 0 will be a, and alpha = 1 will be b.
func (v Vector3) LerpVectors(f float64, a, b Vector3) (this Vector3) {
	v.Call("lerpVectors", a.Value, b.Value, f)
	return v
}

// Max If this vector's x, y or z value is less than a's x, y or z value, replace that value with the corresponding max value.
func (v Vector3) Max(a Vector3) (this Vector3) {
	v.Call("max", a.Value)
	return v
}

// Min If this vector's x, y or z value is greater than a's x, y or z value, replace that value with the corresponding min value.
func (v Vector3) Min(a Vector3) (this Vector3) {
	v.Call("min", a.Value)
	return v
}

// Multiply Multiplies this vector by v.
func (v Vector3) Multiply(a Vector3) (this Vector3) {
	v.Call("multiply", a.Value)
	return v
}

// MultiplyScalar Multiplies this vector by scalar f.
func (v Vector3) MultiplyScalar(f float64) (this Vector3) {
	v.Call("multiplyScalar", f)
	return v
}

// Negate Inverts this vector - i.e. sets x = -x, y = -y and z = -z.
func (v Vector3) Negate() (this Vector3) {
	v.Call("negate")
	return v
}

// Normalize Convert this vector to a unit vector - that is, sets it equal to a vector with the same direction as this one, but length 1.
func (v Vector3) Normalize() Vector3 {
	v.Call("normalize")
	return v
}

// SetComponent sets idx component to f. idx of X is 0. idx of Y is 1. idx of Z is 2.
func (v Vector3) SetComponent(idx int, f float64) {
	v.Call("setComponent", idx, f)
}

// SetLength Set this vector to a vector with the same direction as this one, but length l.
func (v Vector3) SetLength(f float64) Vector3 {
	v.Call("setLength", f)
	return v
}

// Sub Subtracts v from this vector.
func (v Vector3) Sub(a Vector3) Vector3 {
	v.Call("sub", a.Value)
	return v
}

// RandomDirection Sets this vector to a uniformly random point on a unit sphere.
func (v Vector3) RandomDirection() Vector3 {
	v.Call("randomDirection")
	return v
}

// Random Sets each component of this vector to a pseudo-random value between 0 and 1, excluding 1.
func (v Vector3) Random() Vector3 {
	v.Call("random")
	return v
}

// SubVectors Sets this vector to a - b.
func (v Vector3) SubVectors(a, b Vector3) Vector3 {
	v.Call("subVectors", a.Value, b.Value)
	return v
}

func (v Vector3) SetCoords(x, y, z float64) Vector3 {
	v.Call("set", x, y, z)
	return v
}

func (v Vector3) SetX(x float64) Vector3 {
	v.Call("setX", x)
	return v
}

func (v Vector3) SetY(y float64) Vector3 {
	v.Call("setY", y)
	return v
}
func (v Vector3) SetZ(z float64) Vector3 {
	v.Call("setZ", z)
	return v
}

func (v Vector3) Coords() (x, y, z float64) {
	x = v.Value.Get("x").Float()
	y = v.Value.Get("y").Float()
	z = v.Value.Get("z").Float()
	return
}
