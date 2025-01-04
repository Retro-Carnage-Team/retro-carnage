package graphics

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"

	pixel "github.com/Retro-Carnage-Team/pixel2"
)

const (
	durationOfEnemyDeathAnimationFrame = 75 // in ms
	durationOfEnemyMovementFrame       = 75 // in ms
)

type PersonSpriteSupplier struct {
	durationSinceLastSprite int64
	enemy                   EnemyVisuals
	lastDirection           geometry.Direction
	lastIndex               int
	wasDying                bool
}

func NewPersonSpriteSupplier(enemy EnemyVisuals) *PersonSpriteSupplier {
	return &PersonSpriteSupplier{
		enemy:                   enemy,
		lastDirection:           *enemy.ViewingDirection(),
		durationSinceLastSprite: 0,
		lastIndex:               0,
	}
}

func (supplier *PersonSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return durationOfEnemyDeathAnimationFrame * 10
}

func (supplier *PersonSpriteSupplier) Sprite(msSinceLastSprite int64) *SpriteWithOffset {
	var skinFrames = enemySkins[supplier.enemy.Skin()].MovementByDirection[supplier.enemy.ViewingDirection().Name]
	if supplier.enemy.Dying() {
		if !supplier.wasDying {
			logging.Info.Println("Enemy just died")
			supplier.durationSinceLastSprite = 0
			supplier.wasDying = true
		} else {
			supplier.durationSinceLastSprite += msSinceLastSprite
			logging.Info.Printf("Enemy is dying: %d", supplier.durationSinceLastSprite)
		}

		var result = skinFrames[supplier.lastIndex].ToSpriteWithOffset()
		if supplier.durationSinceLastSprite != 0 {
			var alpha = 1.0 - 1.0/float64(supplier.GetDurationOfEnemyDeathAnimation())*float64(supplier.durationSinceLastSprite)
			logging.Info.Printf("Alpha value is: %f", alpha)
			var rgba = pixel.Alpha(alpha)
			result.ColorMask = &rgba
		}
		return result
	} else {
		supplier.wasDying = false
		if supplier.lastDirection != *supplier.enemy.ViewingDirection() {
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
