//go:build TODO

package three

//go:generate go run material_method_generator/main.go -materialName MeshPhysicalMaterial -materialSlug mesh_physical_material

import "syscall/js"

type MeshPhysicalMaterial struct {
	*js.Object

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

func NewMeshPhysicalMaterial(params *MaterialParameters) *MeshBasicMaterial {
	return &MeshBasicMaterial{
		Object: three.Get("MeshBasicMaterial").New(params.Object),
	}
}
