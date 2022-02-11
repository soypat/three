package main

import (
	"syscall/js"

	"github.com/soypat/three"
)

var (
	scene    *three.Scene
	camera   three.PerspectiveCamera
	renderer three.WebGLRenderer
	mesh     *three.Mesh
)

func main() {
	three.AddScript("../_vendor/three.js", "THREE")
	three.Init()
	document := js.Global().Get("document")

	windowWidth := js.Global().Get("innerWidth").Float()
	windowHeight := js.Global().Get("innerHeight").Float()
	devicePixelRatio := js.Global().Get("devicePixelRatio").Float()

	camera = three.NewPerspectiveCamera(70, windowWidth/windowHeight, 1, 1000)
	camera.SetPosition(three.NewVector3(0, 0, 400))

	scene = three.NewScene()

	light := three.NewDirectionalLight(three.NewColor("white"), 1)
	light.SetPosition(three.NewVector3(256, 256, 256).Normalize())

	scene.Add(light)

	ambLight := three.NewAmbientLight(three.NewColorHex(0xbbbbbb), 0.4)
	scene.Add(ambLight)

	renderer = three.NewWebGLRenderer()
	renderer.SetPixelRatio(devicePixelRatio)
	renderer.SetSize(windowWidth, windowHeight, true)
	document.Get("body").Call("appendChild", renderer.Get("domElement"))

	// Create cube
	geometry := three.NewBoxGeometry(&three.BoxGeometryParameters{
		Width:  128,
		Height: 128,
		Depth:  128,
	})

	//material := three.NewMeshBasicMaterial(materialParams)
	material := three.NewMeshLambertMaterial(&three.MaterialParameters{
		Color:       three.NewColor("blue"),
		Side:        three.FrontSide,
		FlatShading: false,
	})

	mesh = three.NewMesh(geometry, material)

	scene.Add(mesh)

	animate(js.Value{}, nil)
	select {}
}

func animate(_ js.Value, _ []js.Value) interface{} {
	rot := mesh.GetRotation()
	x, y, z := rot.Angles()
	mesh.SetRotation(three.NewEuler(x+0.01, y+0.01, z, ""))

	renderer.Render(scene, camera)

	// Best practice (soypat's opinion) to request frame after work is done.
	js.Global().Call("requestAnimationFrame", js.FuncOf(animate))
	return nil
}
