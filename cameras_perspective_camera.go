package three

//go:generate go run codegen/camera_method_generator/main.go -typeName PerspectiveCamera -typeSlug perspective_camera

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
