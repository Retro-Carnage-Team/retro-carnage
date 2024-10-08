package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"

	pixel "github.com/Retro-Carnage-Team/pixel2"
)

const (
	durationOfEnemyDeathAnimationFrame = 75 // in ms
	durationOfEnemyMovementFrame       = 75 // in ms
)

type PersonSpriteSupplier struct {
	durationSinceLastSprite int64
	lastDirection           geometry.Direction
	lastIndex               int
	wasDying                bool
}

func NewPersonSpriteSupplier(direction geometry.Direction) *PersonSpriteSupplier {
	return &PersonSpriteSupplier{
		lastDirection:           direction,
		durationSinceLastSprite: 0,
		lastIndex:               0,
	}
}

func (supplier *PersonSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return durationOfEnemyDeathAnimationFrame * 10
}

func (supplier *PersonSpriteSupplier) Sprite(msSinceLastSprite int64, enemy ActiveEnemy) *graphics.SpriteWithOffset {
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
			var alpha = 1.0 - 1.0/float64(supplier.GetDurationOfEnemyDeathAnimation())*float64(supplier.durationSinceLastSprite)
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
			if supplier.durationSinceLastSprite > durationOfEnemyMovementFrame {
				supplier.durationSinceLastSprite = 0
				supplier.lastIndex = (supplier.lastIndex + 1) % len(skinFrames)
			}
			return skinFrames[supplier.lastIndex].ToSpriteWithOffset()
		}
	}
}
