package three

import (
	"image/color"
	"syscall/js"

	"github.com/soypat/gwasm"
)

// Color - represents a color.
type Color struct {
	js.Value
}

// NewColor returns a new three.Color for the given string.
// It accepts different string color formats:
//   three.NewColor("rgb(255, 0, 0)")
//   three.NewColor("rgb(100%, 0%, 0%)")
//   three.NewColor("skyblue") // X11 color names (without CamelCase), see three.js source
//   three.NewColor("hsl(0, 100%, 50%)")
// See https://threejs.org/docs/#api/en/math/Color for additional details.
func NewColor(color string) *Color {
	return &Color{
		Value: three.Get("Color").New(color),
	}
}

// NewColorRGB returns a new three.Color for the given RGB values in 0..255 range.
func NewColorRGB(r, g, b uint8) *Color {
	rgb := gwasm.JSColor(color.RGBA{R: r, G: g, B: b, A: 0xff})
	return NewColor(rgb)
}

// NewColorRGBFloat returns a new three.Color for the given RGB values in 0.0..0.1 range.
func NewColorRGBFloat(r, g, b float64) *Color {
	return &Color{
		Value: three.Get("Color").New(r, g, b),
	}
}

// NewColorHex returns a new three.Color for the given hex integer value.
// Example:
//   three.NewColorHex(0xff0000)
func NewColorHex(i int64) *Color {
	return &Color{
		Value: three.Get("Color").New(i),
	}
}
