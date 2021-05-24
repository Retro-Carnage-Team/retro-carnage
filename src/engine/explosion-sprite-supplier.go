package engine

import (
	"fmt"
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/util"
)

const (
	DurationOfFrame          = 25 // in ms
	Folder                   = "images/tiles/explosion"
	NumberOfExplosionSprites = 48
)

var (
	explosionSprites = buildAnimation()
)

type ExplosionSpriteSupplier struct {
	duration int64
}

func (ess *ExplosionSpriteSupplier) Sprite(elapsedTimeInMs int64) *graphics.SpriteWithOffset {
	ess.duration += elapsedTimeInMs
	var idx = int(ess.duration / DurationOfFrame)
	return explosionSprites[util.MinInt(len(explosionSprites)-1, idx)]
}

func buildAnimation() []*graphics.SpriteWithOffset {
	var result = make([]*graphics.SpriteWithOffset, 0)
	for i := 0; i < NumberOfExplosionSprites; i++ {
		var spritePath = fmt.Sprintf("%s/%d.png", Folder, i)
		var sprite = assets.SpriteRepository.Get(spritePath)
		result = append(result, &graphics.SpriteWithOffset{
			Offset: geometry.Point{},
			Source: spritePath,
			Sprite: sprite,
		})
	}
	return result
}
