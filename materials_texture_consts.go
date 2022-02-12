package three

type (
	MappingMode     int
	WrappingMode    int
	MinMagFilter    int
	TextureFormat   int
	TextureType     int
	TextureEncoding int
)

// These define the texture's mapping mode.
// UVMapping is the default, and maps the texture using the mesh's UV coordinates.
// The rest define environment mapping types.
const (
	UVMapping MappingMode = iota + 300
	CubeReflectionMapping
	CubeRefractionMapping
	EquirectangularReflectionMapping
	EquirectangularRefractionMapping

	// Where is 305?

	CubeUVReflectionMapping MappingMode = 306
	CubeUVRefractionMapping MappingMode = 307
)

func (m *MappingMode) setDefault() { *m = UVMapping }
func (m *MappingMode) clampDefault() {
	if *m < UVMapping || *m > CubeUVRefractionMapping {
		m.setDefault()
	}
}

// These define the texture's wrapS and wrapT properties, which define horizontal and vertical texture wrapping.
const (
	RepeatWrapping WrappingMode = iota + 1000
	ClampToEdgeWrapping
	MirroredRepeatWrapping
)

func (w *WrappingMode) setDefault() { *w = ClampToEdgeWrapping }
func (w *WrappingMode) clampDefault() {
	if *w < RepeatWrapping || *w > MirroredRepeatWrapping {
		w.setDefault()
	}
}

// For use with a texture's minFilter (and magfilter)	property, these define the texture minifying function that is used whenever the pixel being textured maps to an area greater than one texture element (texel).
const (
	NearestFilter MinMagFilter = iota + 1003
	NearestMipmapNearestFilter
	NearestMipmapLinearFilter
	LinearFilter
	LinearMipmapNearestFilter
	LinearMipmapLinearFilter
)

func (filter *MinMagFilter) setDefault(isMagnification bool) {
	if isMagnification {
		*filter = LinearFilter
	} else {
		// Is Minification Filter
		*filter = LinearMipmapLinearFilter
	}
}
func (filter *MinMagFilter) clampDefault(isMagnification bool) {
	if *filter < NearestFilter || *filter > LinearMipmapLinearFilter {
		filter.setDefault(isMagnification)
	}
}

// For use with a texture's format	property, these define how elements of a 2d texture, or texels, are read by shaders.
const (
	AlphaFormat          TextureFormat = 1021
	RGBFormat            TextureFormat = 1022
	RGBAFormat           TextureFormat = 1023
	LuminanceFormat      TextureFormat = 1024
	LuminanceAlphaFormat TextureFormat = 1025
	RGBEFormat           TextureFormat = RGBAFormat
	DepthFormat          TextureFormat = 1026
	DepthStencilFormat   TextureFormat = 1027
	RedFormat            TextureFormat = 1028
	RedIntegerFormat     TextureFormat = 1029
	RGFormat             TextureFormat = 1030
	RGIntegerFormat      TextureFormat = 1031
	RGBIntegerFormat     TextureFormat = 1032
	RGBAIntegerFormat    TextureFormat = 1033
)

func (f *TextureFormat) setDefault() { *f = RGBAFormat }
func (f *TextureFormat) clampDefault() {
	if *f < AlphaFormat || *f > RGBAIntegerFormat {
		f.setDefault()
	}
}

// For use with a texture's type property, which must correspond to the correct format. See below for details.
const (
	UnsignedByteType      TextureType = 1009
	ByteType              TextureType = 1010
	ShortType             TextureType = 1011
	UnsignedShortType     TextureType = 1012
	IntType               TextureType = 1013
	UnsignedIntType       TextureType = 1014
	FloatType             TextureType = 1015
	HalfFloatType         TextureType = 1016
	UnsignedShort4444Type TextureType = 1017
	UnsignedShort5551Type TextureType = 1018
	UnsignedShort565Type  TextureType = 1019
	UnsignedInt248Type    TextureType = 1020
)

func (t *TextureType) setDefault() { *t = UnsignedByteType }
func (t *TextureType) clampDefault() {
	if *t < UnsignedByteType || *t > UnsignedInt248Type {
		t.setDefault()
	}
}

// LinearEncoding is the default. Values other than this are only valid for a material's map, envMap and emissiveMap.
const (
	LinearEncoding    TextureEncoding = 3000
	sRGBEncoding      TextureEncoding = 3001
	GammaEncoding     TextureEncoding = 3007
	RGBEEncoding      TextureEncoding = 3002
	LogLuvEncoding    TextureEncoding = 3003
	RGBM7Encoding     TextureEncoding = 3004
	RGBM16Encoding    TextureEncoding = 3005
	RGBDEncoding      TextureEncoding = 3006
	BasicDepthPacking TextureEncoding = 3200
	RGBADepthPacking  TextureEncoding = 3201
)

func (t *TextureEncoding) setDefault() { *t = LinearEncoding }
func (t *TextureEncoding) clampDefault() {
	if *t < LinearEncoding || *t > RGBADepthPacking {
		t.setDefault()
	}
}
