package three

//go:generate go run material_method_generator/main.go -materialName SpriteMaterial -materialSlug sprite_material

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
