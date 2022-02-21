package three

//go:generate go run codegen/material_method_generator/main.go -typeName LineBasicMaterial -typeSlug line_basic_material

import "syscall/js"

type LineBasicMaterial struct {
	js.Value
}

func NewLineBasicMaterial(params *MaterialParameters) LineBasicMaterial {
	v := objectify(params)
	return LineBasicMaterial{
		Value: three.Get("LineBasicMaterial").New(v),
	}
}
