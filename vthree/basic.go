//go:build js

package vthree

import (
	"syscall/js"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/three"
)

const defaultRenderPeriod = 200 * time.Millisecond

// Basic is a Vecty Component implementation for use
// with the three package. Basic works by calling
// Animate callback between rendered frames or if
// a static scene is to be displayed, by a nil Animate
// function and a initializer Init function. Init cannot be nil.
//
// All `vecty:"prop"` arguments
// should remain constant during rendering until Basic is
// unmounted.
type Basic struct {
	vecty.Core

	// Wait period is the time renderer sleeps
	// between unrendered frames when Animate returns false.
	// Default WaitPeriod on zero value is 200ms.
	WaitPeriod time.Duration `vecty:"prop"`
	// Markup should be constant during rendering.
	Markup []vecty.MarkupOrChild `vecty:"prop"`
	// Init is called upon mounting. It should not be nil.
	Init func(three.WebGLRenderer) `vecty:"prop"`
	// Animate should return true if user wants to render
	// the current scene. If Animate is nil then scene is rendered
	// as Init left it.
	Animate func(three.WebGLRenderer) bool `vecty:"prop"`
	// Shutdown is called on dismounting the vecty component.
	Shutdown func(three.WebGLRenderer) `vecty:"prop"`

	canvas   *vecty.HTML
	renderer three.WebGLRenderer
}

// Render implements vecty.MarkupOrChild interface.
func (b *Basic) Render() vecty.ComponentOrHTML {
	if b.Init == nil {
		panic("nil init function")
	}
	if b.WaitPeriod == 0 {
		b.WaitPeriod = defaultRenderPeriod
	}
	b.canvas = elem.Canvas(b.Markup...)
	return b.canvas
}

func (b *Basic) SkipRender(vecty.Component) bool {
	return true // !b.renderer.Truthy()
}

// Mount is called automatically by Vecty.
func (v *Basic) Mount() {
	err := three.Init()
	if err != nil {
		panic(err)
	}
	canvas := v.canvas.Node()
	if !canvas.Truthy() {
		panic("canvas not defined in DOM")
	}
	canvas.Set("width", 1000)
	canvas.Set("height", 800)
	v.renderer = three.NewWebGLRenderer(three.WebGLRendererParam{
		Canvas: canvas,
	})
	v.Init(v.renderer)
	v.animate(js.Null(), nil)
}

// Unmount will tear down the renderer. It is
// called automatically by vecty when Basic's
// node parent is removed from the DOM.
func (v *Basic) Unmount() {
	if v.Shutdown != nil {
		v.Shutdown(v.renderer)
	}
	v.renderer.Value = js.Null() // tear down.
}

func (v *Basic) animate(js.Value, []js.Value) interface{} {
	if v.Animate == nil || !v.renderer.Truthy() {
		// Renderer has been teared down or no more frames to render.
		// End animate cycle.
		return nil
	}
	if v.Animate(v.renderer) {
		js.Global().Call("requestAnimationFrame", js.FuncOf(v.animate))
	} else {
		time.Sleep(v.WaitPeriod) // try again in a while.
		js.Global().Call("requestAnimationFrame", js.FuncOf(v.animate))
	}
	return nil
}
