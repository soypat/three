package three

//go:generate go run material_method_generator/main.go -materialName MeshBasicMaterial -materialSlug mesh_basic_material

import "syscall/js"

type MeshBasicMaterial struct {
	js.Value
}

func NewMeshBasicMaterial(params *MaterialParameters) MeshBasicMaterial {
	v := objectify(params)
	return MeshBasicMaterial{
		Value: three.Get("MeshBasicMaterial").New(v),
	}
}
