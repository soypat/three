package three

import "syscall/js"

// Euler angles describe a rotational transformation by
// rotating an object on its various axes in specified amounts per axis, and a specified axis order.
type Euler struct {
	// X     float64 `js:"x"`
	// Y     float64 `js:"y"`
	// Z     float64 `js:"z"`
	// Order string  `js:"order"`
	js.Value
}

// Three.js uses intrinsic Tait-Bryan angles. This means that rotations are performed with respect to the local
// coordinate system. That is, for order 'XYZ', the rotation is first around the local-X axis
// (which is the same as the world-X axis), then around local-Y (which may now be different from the world Y-axis),
// then local-Z (which may be different from the world Z-axis).
func NewEuler(x, y, z float64, order string) Euler {
	if len(order) != 3 {
		order = "XYZ"
	}
	return Euler{
		Value: three.Get("Euler").New(x, y, z, order),
	}
}

// Copy copies value of euler to receiver.
func (e Euler) Copy(src Euler) {
	e.Call("copy", src.Value)
}

func (e Euler) Reorder(newOrder string) {
	e.Call("reorder", newOrder)
}

func (e Euler) SetFromQuaternion(q Quaternion, order string) {
	e.Call("setFromQuaternion", q.Value, order)
}

func (e Euler) SetAngles(x, y, z float64, order string) {
	e.Call("set", x, y, z, order)
}

func (e Euler) Angles() (x, y, z float64) {
	x = e.Get("x").Float()
	y = e.Get("y").Float()
	z = e.Get("z").Float()
	return x, y, z
}

func (e Euler) GetOrder() string {
	return e.Get("order").String()
}
