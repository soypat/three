package three

//go:generate go run codegen/object3d_method_generator/main.go -typeName TextSprite -typeSlug text_sprite

import "syscall/js"

type TextSprite struct {
	// Text string
	// ID               int             `js:"id"`
	// Position         *Vector3        `js:"position"`
	// Rotation         *Euler          `js:"rotation"`
	// Geometry         *BufferGeometry `js:"geometry"`
	// MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	// RenderOrder      int             `js:"renderOrder"`
	// Visible          bool            `js:"visible"`
	js.Value
}

type TextSpriteParameters struct {
	// TextSize       int                 `js:"textSize"`
	// RedrawInterval int                 `js:"redrawInterval"`
	// Material       *TextSpriteMaterial `js:"material"`
	// Texture        *TextTextureParams  `js:"texture"`
	// SEE MaterialParameters for example on how to implement this.
	js.Value
}

type TextSpriteMaterial struct {
	// Color *Color `js:"color"`
	js.Value
}

type TextTextureParams struct {
	// Text       string `js:"text"`
	// FontFamily string `js:"fontFamily"`
	// Align      string `js:"align"`
	js.Value
}

func NewTextSprite(text, align string, textSize int, color *Color) TextSprite {
	material := TextSpriteMaterial{
		Value: js.Global().Get("THREE").Get("SpriteMaterial").New(),
	}
	material.Set("color", color.Value)

	textureParams := TextTextureParams{
		Value: js.Global().Get("THREE").Get("TextTexture").New(),
	}
	textureParams.Set("text", text)
	textureParams.Set("align", align)

	params := TextSpriteParameters{
		Value: js.Global().Get("Object").New(),
	}
	params.Value.Set("texture", textureParams.Value)
	params.Value.Set("material", material.Value)
	params.Value.Set("textSize", textSize)
	obj := objectify(params)
	return TextSprite{
		Value: js.Global().Get("THREE").Get("TextSprite").New(obj),
	}
}
