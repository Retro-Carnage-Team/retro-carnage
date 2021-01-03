package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type EnemyLandmineSpriteSupplier struct{}

func (supplier *EnemyLandmineSpriteSupplier) Sprite(int64, Enemy) *graphics.SpriteWithOffset {
	var sprite = assets.SpriteRepository.Get("images/tiles/environment/Tellermine-43.png")
	var offset = geometry.Point{X: 0, Y: 0}
	return &graphics.SpriteWithOffset{
		Offset: offset,
		Sprite: sprite,
	}
}
