package engine

import (
	"fmt"
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/util"
)

const (
	DurationOfExplosionFrame = 25 // in ms
	Folder                   = "images/explosion"
	NumberOfExplosionSprites = 48
)

// ExplosionSpriteSupplier provides Sprites for explosions - just as the name suggests.
type ExplosionSpriteSupplier struct {
	duration int64
}

// Sprite returns the correct SpriteWithOffset for the specified elapsedTimeInMs
func (ess *ExplosionSpriteSupplier) Sprite(elapsedTimeInMs int64) *graphics.SpriteWithOffset {
	ess.duration += elapsedTimeInMs
	var idx = util.MinInt(NumberOfExplosionSprites-1, int(ess.duration/DurationOfExplosionFrame))
	var spritePath = fmt.Sprintf("%s/%d.png", Folder, idx)
	var sprite = assets.SpriteRepository.Get(spritePath)
	return &graphics.SpriteWithOffset{
		Offset: geometry.Point{},
		Source: spritePath,
		Sprite: sprite,
	}
}
