package three

//go:generate go run material_method_generator/main.go -materialName MeshPhongMaterial -materialSlug mesh_phong_material

import "syscall/js"

type MeshPhongMaterial struct {
	// BumpMap     *Texture `js:"bumpMap"`
	// BumpScale   float64  `js:"bumpScale"`
	// SpecularMap *Texture `js:"specularMap"`
	// Specular    *Color   `js:"specular"`
	js.Value
}

func NewMeshPhongMaterial(params *MaterialParameters) *MeshPhongMaterial {
	v := objectify(params)
	return &MeshPhongMaterial{
		Value: three.Get("MeshPhongMaterial").New(v),
	}
}
