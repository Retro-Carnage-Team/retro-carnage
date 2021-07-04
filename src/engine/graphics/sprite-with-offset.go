package graphics

import (
	"github.com/faiface/pixel"
	"retro-carnage/engine/geometry"
)

type SpriteWithOffset struct {
	ColorMask *pixel.RGBA
	Offset    geometry.Point
	Source    string
	Sprite    *pixel.Sprite
}

func (swo *SpriteWithOffset) String() string {
	return swo.Source
}
