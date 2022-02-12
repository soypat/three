package three

import "syscall/js"

// TextureLoader loads a texture from file.
// See https://threejs.org/docs/#api/en/loaders/TextureLoader
type TextureLoader struct {
	// CrossOrigin string `js:"crossOrigin"`
	js.Value
}

// NewTextureLoader creates a new texture loader.
func NewTextureLoader() *TextureLoader {
	return &TextureLoader{
		Value: three.Get("TextureLoader").New(),
	}
}

// Load loads texture from url image.
func (t *TextureLoader) Load(url string, fn func(*Texture)) *Texture {
	callback := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if fn != nil {
			fn(&Texture{Value: args[0]})
		}
		return nil
	})
	return &Texture{
		Value: t.Call("load", url, callback),
	}
}
