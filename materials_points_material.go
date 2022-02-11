//go:build TODO

package three

//go:generate go run material_method_generator/main.go -materialName PointsMaterial -materialSlug points_material

import "syscall/js"

type PointsMaterial struct {
	*js.Object
}

func NewPointsMaterial(params *MaterialParameters) *PointsMaterial {
	return &PointsMaterial{
		Object: three.Get("PointsMaterial").New(params.Object),
	}
}
