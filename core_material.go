package three

import "syscall/js"

type Material interface {
	OnBeforeCompile()
	SetValues(values MaterialParameters)
	ToJSON(meta interface{}) js.Value
	Clone()
	Copy(source Object3D)
	Dispose()

	getInternalObject() js.Value
}

// Side defines which side of faces will be rendered - front, back or both. Default is FrontSide.
type Side float64

const (
	FrontSide  Side = 0
	BackSide   Side = 1
	DoubleSide Side = 2
)

type Shading float64

const (
	// SmoothShading is the default and linearly interpolates color between vertices.
	SmoothShading Shading = 0

	// FlatShading uses the color of the first vertex for every pixel in a face.
	FlatShading Shading = 1
)

type MaterialParameters struct {
	Color       *Color   `js:"color"`
	Shading     Shading  `js:"shading"` // THREE.MeshLambertMaterial: .shading has been removed. Use the boolean .flatShading instead.
	FlatShading bool     `js:"flatShading"`
	Side        Side     `js:"side"`
	Transparent bool     `js:"transparent"`
	Opacity     float64  `js:"opacity"`
	Map         *Texture `js:"map"`
	LineWidth   float64  `js:"linewidth"`

	// Points
	Size float64 `js:"size"`

	// Phong Materials

	BumpMap     *Texture `js:"bumpMap"`
	BumpScale   float64  `js:"bumpScale"`
	SpecularMap *Texture `js:"specularMap"`
	Specular    *Color   `js:"specular"`

	// Physical Materials

	Clearcoat             float64  `js:"clearcoat"`
	ClearcoatMap          *Texture `js:"clearcoatMap"`
	ClearcoatNormalMap    *Texture `js:"clearcoatNormalMap"`
	ClearcoatNormalScale  Vector2  `js:"clearcoatNormalScale"`
	ClearcoatRoughness    float64  `js:"clearcoatRoughness"`
	ClearcoatRoughnessMap *Texture `js:"clearcoatRoughnessMap"`
	// Index of refraction
	IOR               float64  `js:"ior"`
	Reflectivity      float64  `js:"reflectivity"`
	Sheen             float64  `js:"sheen"`
	SheenRoughness    float64  `js:"sheenRoughness"`
	SheenRoughnessMap *Texture `js:"sheenRoughnessMap"`
	SheenColor        *Color   `js:"sheenColor"`
	SheenColorMap     *Texture `js:"sheenColorMap"`
	Transmission      float64  `js:"transmission"`
	TransmissionMap   *Texture `js:"transmissionMap"`
}

// func NewMaterialParameters(param MaterialParameters) js.Value {
// 	return objectify(param)
// }
