package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type TankSpriteSupplier struct {
	durationSinceLastSprite int64
	lastDirection           geometry.Direction
	lastIndex               int
	// wasDying                bool
}

func NewTankSpriteSupplier(direction geometry.Direction) *TankSpriteSupplier {
	return &TankSpriteSupplier{
		lastDirection:           direction,
		durationSinceLastSprite: 0,
		lastIndex:               0,
	}
}

func (supplier *TankSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return 9223372036854775807
}

func (supplier *TankSpriteSupplier) Sprite(msSinceLastSprite int64, enemy ActiveEnemy) *graphics.SpriteWithOffset {
	// TODO
	return nil
}
