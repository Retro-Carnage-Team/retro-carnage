package graphics

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

const grenadeImagePath = "images/weapons/grenade.png"

var (
	grenadeSprite *SpriteWithOffset
)

// GrenadeSpriteSupplier is used to provide sprites for grenades.
type GrenadeSpriteSupplier struct{}

func (gss *GrenadeSpriteSupplier) Sprite(int64) *SpriteWithOffset {
	if nil == grenadeSprite {
		var sprite = assets.SpriteRepository.Get(grenadeImagePath)
		grenadeSprite = &SpriteWithOffset{
			Offset: geometry.Point{},
			Source: grenadeImagePath,
			Sprite: sprite,
		}
	}
	return grenadeSprite
}
