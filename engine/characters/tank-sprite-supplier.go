package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type TankSpriteSupplier struct {
	durationSinceLastSprite int64
	direction               geometry.Direction
	lastIndex               int
	wasDying                bool
}

func NewTankSpriteSupplier(direction geometry.Direction) *TankSpriteSupplier {
	return &TankSpriteSupplier{
		direction:               direction,
		durationSinceLastSprite: 0,
		lastIndex:               0,
		wasDying:                false,
	}
}

func (supplier *TankSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return 9223372036854775807
}

func (supplier *TankSpriteSupplier) Sprite(msSinceLastSprite int64, enemy ActiveEnemy) *graphics.SpriteWithOffset {
	var skinFrames = enemySkins[enemy.Skin].MovementByDirection[supplier.direction.Name]
	if enemy.Dying {
		if !supplier.wasDying {
			supplier.durationSinceLastSprite = 0
			supplier.wasDying = true
		} else {
			supplier.durationSinceLastSprite += msSinceLastSprite
		}

		if supplier.durationSinceLastSprite > durationOfEnemyMovementFrame {
			supplier.durationSinceLastSprite = 0
			supplier.lastIndex = (supplier.lastIndex + 1) % len(skinFrames) // TODO: death animation instead of skinFrames
		}
	}
	return skinFrames[supplier.lastIndex].ToSpriteWithOffset()
}
