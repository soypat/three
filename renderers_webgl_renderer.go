package three

import "syscall/js"

// The WebGLRenderer displays your beautifully crafted scenes using WebGL.
type WebGLRenderer struct {
	js.Value
}

type WebGLRendererParam struct {
	// DOM node corresponding to canvas element where WebGL visualization will render.
	Canvas js.Value `js:"canvas"`
}

// WebGLRenderer creates an WebGLRenderer instance.
func NewWebGLRenderer(params WebGLRendererParam) WebGLRenderer {
	v := objectify(params)
	return WebGLRenderer{
		Value: three.Get("WebGLRenderer").New(v),
	}
}

// WebGLRenderer creates an WebGLRenderer instance.
func (r WebGLRenderer) SetSize(w float64, h float64, updateStyle bool) {
	r.Call("setSize", w, h, updateStyle)
}

func (r WebGLRenderer) SetPixelRatio(ratio float64) {
	r.Call("setPixelRatio", ratio)
}

func (r WebGLRenderer) Render(scene Scene, camera Camera) {
	r.Call("render", scene.Value, camera.getInternalObject())
}

func (r WebGLRenderer) DomElement() js.Value {
	return r.Get("domElement")
}
