package graphics

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

type BurnMarkSpriteSupplier struct{}

const (
	burnMarkSprite = "images/environment/burn-mark.png"
)

func (supplier *BurnMarkSpriteSupplier) Sprite() *SpriteWithOffset {
	var sprite = assets.SpriteRepository.Get(burnMarkSprite)
	var offset = geometry.Point{X: 0, Y: 55}
	return &SpriteWithOffset{
		Offset: offset,
		Source: burnMarkSprite,
		Sprite: sprite,
	}
}
