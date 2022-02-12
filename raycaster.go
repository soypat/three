//go:build TODO

package three

import "syscall/js"

// Raycaster assists with raycasting (used for mouse picking)
// See https://threejs.org/docs/#api/en/core/Raycaster
type Raycaster struct {
	*js.Object

	Far           float64 `js:"far"`
	Near          float64 `js:"near"`
	LinePrecision float64 `js:"linePrecision"`
}

// NewRaycaster creates a new raycaster.
func NewRaycaster() *Raycaster {
	return &Raycaster{
		Object: three.Get("Raycaster").New(),
	}
}

// SetFromCamera updates the ray with a new origin and direction.
// coords - 2D coordinates of the mouse, in normalized device coordinates (NDC)---X and Y components should be between -1 and 1.
// camera from which the ray should originate
// TODO(divan): abstract camera away
func (r Raycaster) SetFromCamera(coords Vector2, camera PerspectiveCamera) {
	r.Object.Call("setFromCamera", coords, camera)
}

func (r Raycaster) IntersectObjects(objects *js.Object, recursive bool) *js.Object {
	return r.Object.Call("intersectObjects", objects, recursive)
}
