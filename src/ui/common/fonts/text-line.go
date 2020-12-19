package fonts

import "retro-carnage/engine/geometry"

type TextLine struct {
	dimension *geometry.Point
	position  *geometry.Point
	text      string
}

func (tl *TextLine) Dimension() *geometry.Point {
	return tl.dimension
}

func (tl *TextLine) Position() *geometry.Point {
	return tl.position
}

func (tl *TextLine) Text() string {
	return tl.text
}
