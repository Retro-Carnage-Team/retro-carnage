package graphics

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

type LandmineSpriteSupplier struct{}

const landmineSprite = "images/environment/Tellermine-43.png"

func (supplier *LandmineSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return 1
}

func (supplier *LandmineSpriteSupplier) Sprite(int64) *SpriteWithOffset {
	var sprite = assets.SpriteRepository.Get(landmineSprite)
	var offset = geometry.Point{X: 0, Y: 0}
	return &SpriteWithOffset{
		Offset: offset,
		Source: landmineSprite,
		Sprite: sprite,
	}
}
