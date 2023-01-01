package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type BurnMarkSpriteSupplier struct{}

const (
	burnMarkSprite = "images/environment/burn-mark.png"
)

func (supplier *BurnMarkSpriteSupplier) Sprite() *graphics.SpriteWithOffset {
	var sprite = assets.SpriteRepository.Get(burnMarkSprite)
	var offset = geometry.Point{X: 0, Y: 55}
	return &graphics.SpriteWithOffset{
		Offset: offset,
		Source: burnMarkSprite,
		Sprite: sprite,
	}
}
