package three

//go:generate go run material_method_generator/main.go -materialName PointsMaterial -materialSlug points_material

import "syscall/js"

type PointsMaterial struct {
	js.Value
}

func NewPointsMaterial(params *MaterialParameters) *PointsMaterial {
	v := objectify(params)
	return &PointsMaterial{
		Value: three.Get("PointsMaterial").New(v),
	}
}
