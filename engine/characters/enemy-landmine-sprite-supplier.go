package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type EnemyLandmineSpriteSupplier struct{}

const landmineSprite = "images/environment/Tellermine-43.png"

func (supplier *EnemyLandmineSpriteSupplier) Sprite(int64, ActiveEnemy) *graphics.SpriteWithOffset {
	var sprite = assets.SpriteRepository.Get(landmineSprite)
	var offset = geometry.Point{X: 0, Y: 0}
	return &graphics.SpriteWithOffset{
		Offset: offset,
		Source: landmineSprite,
		Sprite: sprite,
	}
}
