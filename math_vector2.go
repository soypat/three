package three

import "syscall/js"

// Vector2 - represents a Vector2.
type Vector2 struct {
	// X float64 `js:"x"`
	// Y float64 `js:"y"`
	js.Value
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{
		Value: three.Get("Vector2").New(x, y),
	}
}

func (v Vector2) Set(x, y float64) Vector2 {
	v.Call("set", x, y)
	return v
}

func (v Vector2) Coords() (x, y float64) {
	x = v.Value.Get("x").Float()
	y = v.Value.Get("y").Float()
	return
}

func (v Vector2) Normalize() Vector2 {
	v.Call("normalize")
	return v
}

func (v Vector2) DistanceTo(v1 Vector2) float64 {
	return v.Call("ditanceTo", v1.Value).Float()
}
