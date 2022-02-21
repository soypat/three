package three

//go:generate go run codegen/material_method_generator/main.go -typeName MeshLambertMaterial -typeSlug mesh_lambert_material

import "syscall/js"

type MeshLambertMaterial struct {
	js.Value
}

func NewMeshLambertMaterial(params *MaterialParameters) MeshLambertMaterial {
	v := objectify(params)
	return MeshLambertMaterial{
		Value: three.Get("MeshLambertMaterial").New(v),
	}
}
