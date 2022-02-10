package main

import (
	"syscall/js"

	"github.com/soypat/three"
)

/*
	Trackball controls earth visualization.
	Warning: VERY LOW RESOLUTION. You can get better images online!
*/

var (
	scene    *three.Scene
	camera   three.PerspectiveCamera
	renderer three.WebGLRenderer
	controls three.TrackballControls
)

const (
	earthSize = 6300e3 // 6300km approx
)

func main() {
	document := js.Global().Get("document")
	windowWidth := js.Global().Get("innerWidth").Float()
	windowHeight := js.Global().Get("innerHeight").Float()
	devicePixelRatio := js.Global().Get("devicePixelRatio").Float()
	camDistance := earthSize * 4
	camera = three.NewPerspectiveCamera(70, windowWidth/windowHeight, 1, camDistance+earthSize)
	camera.SetPosition(three.NewVector3(camDistance, 0, 0))
	camera.LookAt(0, 0, 0)
	scene = three.NewScene()

	light := three.NewDirectionalLight(three.NewColor("white"), 1)
	light.Position.Set(earthSize*5, earthSize, 0)
	scene.Add(light)

	ambLight := three.NewAmbientLight(three.NewColorHex(0xbbbbbb), 0.4)
	scene.Add(ambLight)

	renderer = three.NewWebGLRenderer()
	renderer.SetPixelRatio(devicePixelRatio)
	renderer.SetSize(windowWidth, windowHeight, true)
	rendererElement := renderer.Get("domElement")
	document.Get("body").Call("appendChild", rendererElement)
	earth := CreateEarth(earthSize)
	scene.Add(earth)

	// Controls to rotate camera around earth
	controls = three.NewTrackballControls(camera, rendererElement)
	controls.MaxDistance = camDistance
	controls.MinDistance = earthSize * 1.1
	controls.RotateSpeed = 1
	controls.ZoomSpeed = 1.2
	controls.PanSpeed = 0.8
	controls.RotateSpeed = 0.8

	animate()
}

func animate() {
	controls.Update()
	renderer.Render(scene, camera)

	// Best practice (soypat's opinion) to request frame after work is done.
	js.Global().Call("requestAnimationFrame", animate)
}

// CreateEarth returns a colored, reflective, beautiful, low-res, blue-marble rendition of where we live.
// Get Hi-res images from https://visibleearth.nasa.gov/ under Blue Marble collection.
func CreateEarth(radius float64) *three.Mesh {
	// Create earth geometry sphere
	geometry := three.NewSphereGeometry(three.SphereGeometryParameters{
		Radius:         earthSize,
		HeightSegments: 128,
		WidthSegments:  200,
	})

	// Add textures. Visible+topo maps taken from Blue Marble collection from https://visibleearth.nasa.gov/
	materialParams := three.NewMaterialParameters()
	materialParams.Map = three.NewTextureLoader().Load("./assets/visible_small.jpg", nil)          // visible earth
	materialParams.BumpMap = three.NewTextureLoader().Load("./assets/topo_small.jpg", nil)         // Show mountain shade
	materialParams.BumpScale = 8e3 / 6300e3 * radius                                               // Mt. everest is 8km
	materialParams.SpecularMap = three.NewTextureLoader().Load("./assets/specular_small.png", nil) // Water reflection
	materialParams.Specular = three.NewColorHex(0x202020)                                          // greyish reflection color. white too bright

	//material := three.NewMeshBasicMaterial(materialParams)
	material := three.NewMeshPhongMaterial(materialParams)
	//material := three.NewMeshPhongMaterial(materialParams)
	return three.NewMesh(geometry, material)
}
