package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"

	pixel "github.com/Retro-Carnage-Team/pixel2"
)

const (
	DurationOfEnemyDeathAnimationFrame = 75 // in ms
	DurationOfEnemyDeathAnimation      = DurationOfEnemyDeathAnimationFrame * 10
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
	var skinFrames = enemySkins[enemy.Skin].MovementByDirection[enemy.ViewingDirection.Name]
	if enemy.Dying {
		if !supplier.wasDying {
			supplier.durationSinceLastSprite = 0
			supplier.wasDying = true
		} else {
			supplier.durationSinceLastSprite += msSinceLastSprite
		}

		var result = skinFrames[supplier.lastIndex].ToSpriteWithOffset()
		if supplier.durationSinceLastSprite != 0 {
			var alpha = 1.0 - 1.0/float64(DurationOfEnemyDeathAnimation)*float64(supplier.durationSinceLastSprite)
			var rgba = pixel.Alpha(alpha)
			result.ColorMask = &rgba
		}
		return result
	} else {
		supplier.wasDying = false
		if supplier.lastDirection != *enemy.ViewingDirection {
			supplier.durationSinceLastSprite = 0
			supplier.lastIndex = 0
			var skinFrame = skinFrames[supplier.lastIndex]
			return skinFrame.ToSpriteWithOffset()
		} else {
			supplier.durationSinceLastSprite += msSinceLastSprite
			if supplier.durationSinceLastSprite > DurationOfEnemyMovementFrame {
				supplier.durationSinceLastSprite = 0
				supplier.lastIndex = (supplier.lastIndex + 1) % len(skinFrames)
			}
			return skinFrames[supplier.lastIndex].ToSpriteWithOffset()
		}
	}
}
