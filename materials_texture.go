package three

import "syscall/js"

type Image struct {
	js.Value
}

type Texture struct {
	// Id    int    `js:"id"`
	// UUID  string `js:"uuid"`
	// Name  string `js:"name"`
	// Image *Image `js:"image"`
	// // Array of user-specified mipmaps (optional).
	// Mipmaps        *js.Object    `js:"mipmaps"`
	// WrapS          WrappingMode  `js:"wrapS"`
	// WrapT          WrappingMode  `js:"wrapT"`
	// MagFilter      MinMagFilter  `js:"magFilter"`
	// MinFilter      MinMagFilter  `js:"minFilter"`
	// Anisotropy     int           `js:"anisotropy"`
	// Format         TextureFormat `js:"format"`
	// InternalFormat string        `js:"internalFormat"`
	// Type           TextureType   `js:"internalFormat"`
	// // How much a single repetition of the texture is offset from the beginning, in each direction U and V. Typical range is 0.0 to 1.0.
	// Offset Vector2 `js:"offset"`
	// // How many times the texture is repeated across the surface, in each direction U and V. If repeat is set greater than 1 in either direction, the
	// Repeat Vector2 `js:"repeat"`
	// // How much the texture is rotated around the center point, in radians. Positive values are counter-clockwise. Default is 0.
	// Rotation float64 `js:"rotation"`
	// // The point around which rotation occurs. A value of (0.5, 0.5) corresponds to the center of the texture. Default is (0, 0), the lower left.
	// Center           Vector2 `js:"center"`
	// MatrixAutoUpdate bool    `js:"matrixAutoUpdate"`
	// // Matrix Matrix3 `js:"matrix"`
	// GenerateMipmaps  bool            `js:"generateMipmaps"`
	// PremultiplyAlpha bool            `js:"premultiplyAlpha"`
	// FlipY            bool            `js:"flipY"`
	// UnpackAlignment  int             `js:"unpackAlignment"`
	// Encoding         TextureEncoding `js:"encoding"`
	// Version          int             `js:"version"`
	// NeedsUpdate      bool            `js:"needsUpdate"`
	// UserData *js.Object `js:"userData"`
	js.Value
}

type TextureParameters struct {
	Image *Image
	// How the image is applied to the object. An object type of THREE.UVMapping is the default, where the U,V coordinates are used to apply the map.
	Mapping MappingMode
	// This defines how the texture is wrapped horizontally and corresponds to U in UV mapping.
	WrapS WrappingMode
	// This defines how the texture is wrapped vertically and corresponds to V in UV mapping.
	WrapT WrappingMode
	// How the texture is sampled when a texel covers more than one pixel. The default is THREE.LinearFilter, which takes the four closest texels and bilinearly interpolates among them.
	MagFilter MinMagFilter
	// How the texture is sampled when a texel covers less than one pixel. The default is THREE.LinearMipmapLinearFilter, which uses mipmapping and a trilinear filter.
	MinFilter MinMagFilter

	Format TextureFormat
	Type   TextureType
	// The number of samples taken along the axis through the pixel that has the highest density of texels. By default, this value is 1. A higher value gives a less blurry result than a basic mipmap, at the cost of more texture samples being used. Use renderer.getMaxAnisotropy() to find the maximum valid anisotropy value for the GPU; this value is usually a power of 2
	Anisotropy int
	Encoding   TextureEncoding
}

func NewTexture(params TextureParameters) *Texture {
	if params.Anisotropy == 0 {
		params.Anisotropy = 1
	}
	// Set Default values if parameters are invalid
	params.WrapS.clampDefault()
	params.WrapT.clampDefault()
	params.MagFilter.clampDefault(true)
	params.MinFilter.clampDefault(false)
	params.Mapping.clampDefault()
	params.Type.clampDefault()
	params.Format.clampDefault()
	params.Encoding.clampDefault()

	return &Texture{
		Value: three.Get("Texture").New(
			params.Image,
			params.Mapping,
			params.WrapS,
			params.WrapT,
			params.MagFilter,
			params.MinFilter,
			params.Format,
			params.Type,
			params.Anisotropy,
			params.Encoding,
		),
	}
}

// Update the texture's uv-transform .matrix from the texture properties .offset, .repeat, .rotation, and .center.
func (t *Texture) UpdateMatrix() {
	t.Call("updateMatrix")
}

func (t *Texture) Clone() *Texture {
	return &Texture{
		Value: t.Call("clone"),
	}
}

func (t *Texture) ToJSON() interface{} {
	return t.Value.Call("toJSON")
}

func (t *Texture) Dispose() {
	t.Value.Call("dispose")
}

func (t *Texture) TransformUV(uv Vector2) Vector2 {
	return Vector2{
		Value: t.Value.Call("transformUV", uv.Value),
	}
}
