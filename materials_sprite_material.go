//go:build TODO

package three

//go:generate go run material_method_generator/main.go -materialName SpriteMaterial -materialSlug sprite_material

import "syscall/js"

type SpriteMaterial struct {
	*js.Object
}

func NewSpriteMaterial(texture *Texture) *SpriteMaterial {
	params := NewMaterialParameters()
	params.Map = texture
	return &SpriteMaterial{
		Object: three.Get("SpriteMaterial").New(params),
	}
}
