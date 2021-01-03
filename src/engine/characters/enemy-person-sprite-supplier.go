package characters

import (
	"github.com/faiface/pixel"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

const (
	DurationOfDeathAnimationEnemy = 75 * 20 // in m
	DurationOfMovementAnimation   = 75      // in ms
)

type EnemyPersonSpriteSupplier struct {
	direction             geometry.Direction
	durationSinceLastTile int64
	lastSprite            *pixel.Sprite
	tileSet               map[geometry.Direction][]*pixel.Sprite
}

func (supplier *EnemyPersonSpriteSupplier) Sprite(int64, Enemy) *graphics.SpriteWithOffset {
	return nil
}
