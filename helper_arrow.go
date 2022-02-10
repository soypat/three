//go:build TODO

package three

//go:generate go run object3d_method_generator/main.go -typeName ArrowHelper -typeSlug arrow_helper

import "syscall/js"

type ArrowHelper struct {
	*js.Object

	Origin Vector3 `js:"position"`
	Line   Line    `js:"line"`
	// Quaternion Quaternion `js:"quaternion"`
	Cone Mesh `js:"cone"`
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
	novec := Vector3{}
	if params.Dir == novec {
		params.Dir = NewVector3(0, 0, 1)
	}
	if params.Origin == novec {
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
		Object: three.Get("ArrowHelper").New(
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
	g.Object.Call("setColor", color)
}

func (g ArrowHelper) SetLength(length float64, head *ArrowHeadParameters) {
	if head == nil {
		g.Object.Call("setColor", length)
		return
	}
	g.Object.Call("setColor", length, head.HeadLength, head.HeadWidth)
}

// SetDirection Sets the desired direction. Must be a unit vector.
func (g ArrowHelper) SetDirection(dir Vector3) {
	g.Object.Call("setDirection", dir)
}
