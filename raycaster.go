package three

import "syscall/js"

// Raycaster assists with raycasting (used for mouse picking)
// See https://threejs.org/docs/#api/en/core/Raycaster
type Raycaster struct {
	// Far           float64 `js:"far"`
	// Near          float64 `js:"near"`
	// LinePrecision float64 `js:"linePrecision"`
	js.Value
}

// NewRaycaster creates a new raycaster.
func NewRaycaster() *Raycaster {
	return &Raycaster{
		Value: three.Get("Raycaster").New(),
	}
}

// SetFromCamera updates the ray with a new origin and direction.
// coords - 2D coordinates of the mouse, in normalized device coordinates (NDC)---X and Y components should be between -1 and 1.
// camera from which the ray should originate
func (r Raycaster) SetFromCamera(coords Vector2, camera Camera) {
	r.Value.Call("setFromCamera", coords, camera.getInternalObject())
}

func (r Raycaster) IntersectObjects(objects js.Value, recursive bool) js.Value {
	return r.Value.Call("intersectObjects", objects, recursive)
}
