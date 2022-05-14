package three

//go:generate go run codegen/object3d_method_generator/main.go -typeName PerspectiveCamera -typeSlug perspective_camera

import "syscall/js"

type PerspectiveCamera struct {
	// Use Set method to set up vector.
	// Up               Vector3 `js:"up"`
	// Position         Vector3 `js:"position"`
	// MatrixAutoUpdate bool    `js:"matrixAutoUpdate"`
	// Aspect           float64 `js:"aspect"`
	js.Value
}

// Assert PerspectiveCamera implements Camera.
// var _ Camera = PerspectiveCamera{}

func NewPerspectiveCamera(fov, aspect, near, far float64) PerspectiveCamera {
	return PerspectiveCamera{Value: three.Get("PerspectiveCamera").New(fov, aspect, near, far)}
}

func (c PerspectiveCamera) SetFocalLength(focalLength float64) {
	c.Call("setFocalLength", focalLength)
}

func (c PerspectiveCamera) GetFocalLength() float64 {
	return c.Call("getFocalLength").Float()
}

func (c PerspectiveCamera) GetEffectiveFOV() float64 {
	return c.Call("getEffectiveFOV").Float()
}

func (c PerspectiveCamera) GetFilmWidth() float64 {
	return c.Call("getFilmWidth").Float()
}

func (c PerspectiveCamera) GetFilmHeight() float64 {
	return c.Call("getFilmHeight").Float()
}

func (c PerspectiveCamera) SetViewOffset(fullWidth, fullHeight, x, y, width, height float64) {
	c.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
}

func (c PerspectiveCamera) ClearViewOffset() {
	c.Call("clearViewOffset")
}

func (c PerspectiveCamera) UpdateProjectionMatrix() {
	c.Call("updateProjectionMatrix")
}

// SetUp sets the up direction for the camera.
//
// It is the equivalent to c.Up.Set(v.X, v.Y, v.Z)
func (obj PerspectiveCamera) SetUp(v Vector3) {
	x, y, z := v.Coords()
	obj.Get("up").Call("set", x, y, z)
}

// SetFar sets frustum far plane. Default is 2000.
// Needs call to UpdateProjectionMatrix to take effect.
func (obj PerspectiveCamera) SetFar(far float64) {
	obj.Set("far", far)
}

// SetNear sets frustum near plane. Default is 0.1.
// Needs call to UpdateProjectionMatrix to take effect.
func (obj PerspectiveCamera) SetNear(near float64) {
	obj.Set("near", near)
}

// SetAspect sets the camera aspect ratio. Usually the canvas width/canvas height.
// Needs call to UpdateProjectionMatrix to take effect.
func (obj PerspectiveCamera) SetAspect(ratio float64) {
	obj.Set("aspect", ratio)
}

// SetFOV sets Camera frustum vertical field of view, from bottom to top of view, in degrees.
// Needs call to UpdateProjectionMatrix to take effect.
func (obj PerspectiveCamera) SetFOV(fov float64) {
	obj.Set("fov", fov)
}

// Sets the zoom factor of the camera.
// Needs call to UpdateProjectionMatrix to take effect.
func (obj PerspectiveCamera) SetZoom(zoom float64) {
	obj.Set("zoom", zoom)
}

// GetUp gets the up direction for the camera.
func (obj PerspectiveCamera) GetUp() Vector3 {
	return Vector3{
		Value: obj.Get("up"),
	}
}

// GetFar gets frustum far plane. Default is 2000.
func (obj PerspectiveCamera) GetFar() float64 {
	return obj.Get("far").Float()
}

// GetNear gets frustum near plane. Default is 0.1.
func (obj PerspectiveCamera) GetNear() float64 {
	return obj.Get("near").Float()
}

// GetAspect gets the camera aspect ratio. Usually the canvas width/canvas height.
func (obj PerspectiveCamera) GetAspect() float64 {
	return obj.Get("aspect").Float()
}

// GetFOV gets Camera frustum vertical field of view, from bottom to top of view, in degrees.
func (obj PerspectiveCamera) GetFOV() float64 {
	return obj.Get("fov").Float()
}

// GetZoom gets the zoom factor of the camera.
func (obj PerspectiveCamera) GetZoom() float64 {
	return obj.Get("zoom").Float()
}
