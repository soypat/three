package three

//go:generate go run codegen/material_method_generator/main.go -typeName MeshBasicMaterial -typeSlug mesh_basic_material

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
