package graphics

import (
	"retro-carnage/engine/geometry"

	"github.com/Retro-Carnage-Team/pixel"
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
