//go:build TODO

package three

//go:generate go run material_method_generator/main.go -materialName MeshBasicMaterial -materialSlug mesh_basic_material

import (
	"syscall/js"
)

type MeshBasicMaterial struct {
	*js.Object
}

func NewMeshBasicMaterial(params *MaterialParameters) *MeshBasicMaterial {
	return &MeshBasicMaterial{
		Object: three.Get("MeshBasicMaterial").New(params.Object),
	}
}
