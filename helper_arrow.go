package three

//go:generate go run object3d_method_generator/main.go -typeName ArrowHelper -typeSlug arrow_helper

import "syscall/js"

type ArrowHelper struct {
	// Origin Vector3 `js:"position"`
	// Line   Line    `js:"line"`
	// // Quaternion Quaternion `js:"quaternion"`
	// Cone Mesh `js:"cone"`
	js.Value
}

type ArrowHelperParameters struct {
	Dir    Vector3
	Origin Vector3
	Length float64
	Color  *Color
	ArrowHeadParameters
}

type ArrowHeadParameters struct {
	HeadLength float64
	HeadWidth  float64
}

func NewArrowHelper(params *ArrowHelperParameters) *ArrowHelper {
	if params == nil {
		params = &ArrowHelperParameters{}
	}

	if !params.Dir.Truthy() {
		params.Dir = NewVector3(0, 0, 1)
	}
	if !params.Origin.Truthy() {
		params.Origin = NewVector3(0, 0, 0)
	}
	if params.Length == 0 {
		params.Length = 1
	}

	if params.Color == nil {
		params.Color = NewColorHex(0xffff00)
	}
	if params.HeadLength == 0 {
		params.HeadLength = 0.2 * params.Length
	}
	if params.HeadWidth == 0 {
		params.HeadWidth = 0.2 * params.HeadLength
	}

	return &ArrowHelper{
		Value: three.Get("ArrowHelper").New(
			params.Dir,
			params.Origin,
			params.Length,
			params.Color,
			params.HeadLength,
			params.HeadWidth,
		),
	}
}

func (g ArrowHelper) SetColor(color Color) {
	g.Value.Call("setColor", color)
}

func (g ArrowHelper) SetLength(length float64, head *ArrowHeadParameters) {
	if head == nil {
		g.Value.Call("setColor", length)
		return
	}
	g.Value.Call("setColor", length, head.HeadLength, head.HeadWidth)
}

// SetDirection Sets the desired direction. Must be a unit vector.
func (g ArrowHelper) SetDirection(dir Vector3) {
	g.Value.Call("setDirection", dir)
}
