package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

const (
	DurationOfEnemyDeathAnimationFrame = 75 // in ms
	DurationOfEnemyDeathAnimation      = DurationOfEnemyDeathAnimationFrame * 20
	DurationOfEnemyMovementFrame       = 75 // in ms
)

type EnemyPersonSpriteSupplier struct {
	durationSinceLastSprite int64
	lastDirection           geometry.Direction
	lastIndex               int
	wasDying                bool
}

func NewEnemyPersonSpriteSupplier(direction geometry.Direction) *EnemyPersonSpriteSupplier {
	return &EnemyPersonSpriteSupplier{
		lastDirection:           direction,
		durationSinceLastSprite: 0,
		lastIndex:               0,
	}
}

func (supplier *EnemyPersonSpriteSupplier) Sprite(msSinceLastSprite int64, enemy ActiveEnemy) *graphics.SpriteWithOffset {
	if enemy.Dying {
		if !supplier.wasDying {
			supplier.durationSinceLastSprite = 0
			supplier.lastIndex = 0
			supplier.wasDying = true
		}

		var deathSprites = enemySkins[enemy.Skin].DeathAnimation
		supplier.durationSinceLastSprite += msSinceLastSprite
		if supplier.durationSinceLastSprite > DurationOfEnemyDeathAnimationFrame {
			supplier.lastIndex = (supplier.lastIndex + 1) % len(deathSprites)
		}
		return deathSprites[supplier.lastIndex].ToSpriteWithOffset()
	} else {
		supplier.wasDying = false
		if supplier.lastDirection != *enemy.ViewingDirection {
			supplier.durationSinceLastSprite = 0
			supplier.lastIndex = 0
			var skinFrame = enemySkins[enemy.Skin].MovementByDirection[enemy.ViewingDirection.Name][supplier.lastIndex]
			return skinFrame.ToSpriteWithOffset()
		} else {
			var skinFrames = enemySkins[enemy.Skin].MovementByDirection[enemy.ViewingDirection.Name]
			supplier.durationSinceLastSprite += msSinceLastSprite
			if supplier.durationSinceLastSprite > DurationOfEnemyMovementFrame {
				supplier.lastIndex = (supplier.lastIndex + 1) % len(skinFrames)
			}
			return skinFrames[supplier.lastIndex].ToSpriteWithOffset()
		}
	}
}
