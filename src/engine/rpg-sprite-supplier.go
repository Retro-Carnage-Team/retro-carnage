package engine

import (
	"fmt"
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/logging"
)

type RpgSpriteSupplier struct {
	lastIdx int
	sprites []*graphics.SpriteWithOffset
}

func NewRpgSpriteSupplier(direction geometry.Direction) *RpgSpriteSupplier {
	var twoSprites = make([]*graphics.SpriteWithOffset, 0)
	var offset = rpgOffsetByDirection(direction)
	for i := 0; i < 2; i++ {
		var spritePath = fmt.Sprintf("images/tiles/weapons/rpg-%s-%d.png", direction.Name, 0)
		var sprite = assets.SpriteRepository.Get(spritePath)
		twoSprites = append(twoSprites, &graphics.SpriteWithOffset{
			Offset: offset,
			Source: spritePath,
			Sprite: sprite,
		})
	}
	return &RpgSpriteSupplier{lastIdx: 0, sprites: twoSprites}
}

func rpgOffsetByDirection(direction geometry.Direction) geometry.Point {
	switch {
	case direction == geometry.Up:
		return geometry.Point{X: -5, Y: 0}
	case direction == geometry.UpRight:
		return geometry.Point{X: -70, Y: 0}
	case direction == geometry.Right:
		return geometry.Point{X: -96, Y: -5}
	case direction == geometry.DownRight:
		return geometry.Point{X: -70, Y: -70}
	case direction == geometry.Down:
		return geometry.Point{X: -96, Y: -5}
	case direction == geometry.DownLeft:
		return geometry.Point{X: 0, Y: -70}
	case direction == geometry.Left:
		return geometry.Point{X: 0, Y: -5}
	case direction == geometry.UpLeft:
		return geometry.Point{X: 0, Y: 0}
	default:
		logging.Error.Fatalf("no such rpgOffset for direction: %s", direction.Name)
		return geometry.Point{}
	}
}

func (rss *RpgSpriteSupplier) Sprite() *graphics.SpriteWithOffset {
	rss.lastIdx = (rss.lastIdx + 1) % 2
	return rss.sprites[rss.lastIdx]
}
