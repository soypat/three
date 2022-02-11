package three

//go:generate go run material_method_generator/main.go -materialName MeshLambertMaterial -materialSlug mesh_lambert_material

import "syscall/js"

type MeshLambertMaterial struct {
	js.Value
}

func NewMeshLambertMaterial(params *MaterialParameters) *MeshLambertMaterial {
	v := objectify(params)
	return &MeshLambertMaterial{
		Value: three.Get("MeshLambertMaterial").New(v),
	}
}
