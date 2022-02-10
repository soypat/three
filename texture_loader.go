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
	callback := func(obj js.Value) {
		fn(&Texture{Value: obj})
	}
	if fn == nil {
		callback = func(obj js.Value) {}
	}
	return &Texture{
		Value: t.Call("load", url, callback),
	}
}
