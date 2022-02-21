package three

//go:generate go run codegen/material_method_generator/main.go -typeName SpriteMaterial -typeSlug sprite_material

import "syscall/js"

type SpriteMaterial struct {
	js.Value
}

func NewSpriteMaterial(texture *Texture) *SpriteMaterial {
	v := objectify(&MaterialParameters{Map: texture})
	return &SpriteMaterial{
		Value: three.Get("SpriteMaterial").New(v),
	}
}
