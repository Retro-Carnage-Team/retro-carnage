package graphics

import (
	"retro-carnage/engine/geometry"
)

type TankSpriteSupplier struct {
	durationSinceLastSprite int64
	direction               geometry.Direction
	enemy                   EnemyVisuals
	lastIndex               int
}

func NewTankSpriteSupplier(enemy EnemyVisuals) *TankSpriteSupplier {
	return &TankSpriteSupplier{
		direction:               *enemy.ViewingDirection(),
		durationSinceLastSprite: 0,
		enemy:                   enemy,
		lastIndex:               0,
	}
}

func (supplier *TankSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return 9223372036854775807
}

func (supplier *TankSpriteSupplier) Sprite(msSinceLastSprite int64) *SpriteWithOffset {
	if supplier.enemy.Dying() {
		var deathAnimationFrames = enemySkins[supplier.enemy.Skin()].DeathAnimation[supplier.direction.Name]
		supplier.durationSinceLastSprite += msSinceLastSprite

		if supplier.durationSinceLastSprite > durationOfEnemyMovementFrame {
			supplier.durationSinceLastSprite = 0
			supplier.lastIndex = (supplier.lastIndex + 1) % len(deathAnimationFrames)
		}
		return deathAnimationFrames[supplier.lastIndex].ToSpriteWithOffset()
	}

	var skinFrames = enemySkins[supplier.enemy.Skin()].MovementByDirection[supplier.direction.Name]
	return skinFrames[supplier.lastIndex].ToSpriteWithOffset()
}
