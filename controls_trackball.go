package three

import "syscall/js"

type TrackballControls struct {
	// DynamicDampingFactor float64 `js:"dynamicDampingFactor"`
	// Enabled              bool    `js:"enabled"`
	// // How far you can zoom out. Default is Infinity.
	// MaxDistance float64 `js:"maxDistance"`
	// // How far you can zoom in. Default is 0.
	// MinDistance float64 `js:"minDistance"`
	// // This object contains references to the mouse actions used by the controls.
	// // * LEFT is assigned with THREE.MOUSE.ROTATE
	// // * MIDDLE is assigned with THREE.MOUSE.ZOOM
	// // * RIGHT is assigned with THREE.MOUSE.PAN
	// MouseButtons js.Value `js:"mouseButtons"`
	// NoPan        bool     `js:"noPan"`
	// NoRotate     bool     `js:"noRotate"`
	// NoZoom       bool     `js:"noZoom"`
	// Camera       Camera   `js:"object"`
	// PanSpeed     float64  `js:"panSpeed"`
	// RotateSpeed  float64  `js:"rotateSpeed"`
	// // Whether or not damping is disabled. Default is false.
	// StaticMoving bool    `js:"staticMoving"`
	// ZoomSpeed    float64 `js:"zoomSpeed"`
	// // This array holds keycodes for controlling interactions.
	// // When the first defined key is pressed, all mouse interactions (left, middle, right) performs orbiting.
	// // When the second defined key is pressed, all mouse interactions (left, middle, right) performs zooming.
	// // When the third defined key is pressed, all mouse interactions (left, middle, right) performs panning.
	// // Default is KeyA, KeyS, KeyD which represents A, S, D.
	// Keys []string `js:"keys"`
	// // WARNING: Not part of documentation. Use to set where zoom-in focuses on.
	// Target Vector3 `js:"target"`
	// // WARNING: Not part of documentation.
	// Target0 Vector3 `js:"target0"`
	// // WARNING: Not part of documentation. Current position of camera.
	// Position Vector3 `js:"position"`
	js.Value
}

// NewTrackballControls instances a TrackballControls. Requires TrackballControls to be
// added at a global level or on THREE. See https://github.com/turban/webgl-earth for a
// way of doing this easily using Eberhard Graether's http://egraether.com/ version.
func NewTrackballControls(camera Camera, domElement js.Value) TrackballControls {
	const namespace = "TrackballControls"
	trackball := getModule(namespace)
	if !domElement.Truthy() {
		panic("domElement must be defined")
	}
	return TrackballControls{
		Value: trackball.New(camera, domElement),
	}
}

// Ensures the controls stay in the range [minDistance, maxDistance]. Called by update().
func (t TrackballControls) CheckDistances() {
	t.Value.Call("checkDistances")
}

// Should be called if the controls is no longer required.
func (t TrackballControls) Dispose() {
	t.Value.Call("dispose")
}

// Should be called if the application window is resized.
func (t TrackballControls) HandleResizes() {
	t.Value.Call("handleResizes")
}

// Performs panning if necessary. Called by update().
func (t TrackballControls) PanCamera() {
	t.Value.Call("panCamera")
}

// Resets the controls to its initial state.
func (t TrackballControls) Reset() {
	t.Value.Call("reset")
}

// Rotates the camera if necessary. Called by update().
func (t TrackballControls) RotateCamera() {
	t.Value.Call("rotateCamera")
}

// Updates the controls. Usually called in the animation loop.
func (t TrackballControls) Update() {
	t.Value.Call("update")
}

// Performs zooming if necessary. Called by update().
func (t TrackballControls) ZoomCamera() {
	t.Value.Call("zoomCamera")
}

func (t TrackballControls) SetMaxDistance(d float64) {
	t.Value.Set("maxDistance", d)
}

func (t TrackballControls) SetMinDistance(d float64) {
	t.Value.Set("minDistance", d)
}

func (t TrackballControls) SetRotateSpeed(v float64) {
	t.Value.Set("rotateSpeed", v)
}

func (t TrackballControls) SetZoomSpeed(v float64) {
	t.Value.Set("zoomSpeed", v)
}

func (t TrackballControls) SetPanSpeed(v float64) {
	t.Value.Set("panSpeed", v)
}
