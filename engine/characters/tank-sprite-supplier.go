package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type TankSpriteSupplier struct {
	durationSinceLastSprite int64
	direction               geometry.Direction
	lastIndex               int
}

func NewTankSpriteSupplier(direction geometry.Direction) *TankSpriteSupplier {
	return &TankSpriteSupplier{
		direction:               direction,
		durationSinceLastSprite: 0,
		lastIndex:               0,
	}
}

func (supplier *TankSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return 9223372036854775807
}

func (supplier *TankSpriteSupplier) Sprite(msSinceLastSprite int64, enemy ActiveEnemy) *graphics.SpriteWithOffset {
	if enemy.Dying {
		var deathAnimationFrames = enemySkins[enemy.Skin].DeathAnimation[supplier.direction.Name]
		supplier.durationSinceLastSprite += msSinceLastSprite

		if supplier.durationSinceLastSprite > durationOfEnemyMovementFrame {
			supplier.durationSinceLastSprite = 0
			supplier.lastIndex = (supplier.lastIndex + 1) % len(deathAnimationFrames)
		}
		return deathAnimationFrames[supplier.lastIndex].ToSpriteWithOffset()
	}

	var skinFrames = enemySkins[enemy.Skin].MovementByDirection[supplier.direction.Name]
	return skinFrames[supplier.lastIndex].ToSpriteWithOffset()
}
