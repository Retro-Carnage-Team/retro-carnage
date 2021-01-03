package graphics

import (
	"github.com/faiface/pixel"
	"retro-carnage/engine/geometry"
)

type SpriteWithOffset struct {
	Offset geometry.Point
	Sprite *pixel.Sprite
}
