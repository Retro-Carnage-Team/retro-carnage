package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

const (
	DurationOfDeathAnimationFrame = 75 // in ms
	DurationOfMovementFrame       = 75 // in ms
)

type EnemyPersonSpriteSupplier struct {
	durationSinceLastTile int64
	lastDirection         geometry.Direction
	lastIndex             int
	wasDying              bool
}

func (supplier *EnemyPersonSpriteSupplier) Sprite(msSinceLastSprite int64, enemy Enemy) *graphics.SpriteWithOffset {
	if enemy.Dying {
		if !supplier.wasDying {
			supplier.durationSinceLastTile = 0
			supplier.lastIndex = 0
			supplier.wasDying = true
		}

		var deathFrames = enemySkins[enemy.Skin].DeathAnimation
		supplier.durationSinceLastTile += msSinceLastSprite
		if supplier.durationSinceLastTile > DurationOfDeathAnimationFrame {
			supplier.lastIndex = (supplier.lastIndex + 1) % len(deathFrames)
		}
		return deathFrames[supplier.lastIndex].ToSpriteWithOffset()
	} else {
		supplier.wasDying = false
		if supplier.lastDirection != enemy.ViewingDirection {
			supplier.durationSinceLastTile = 0
			supplier.lastIndex = 0
			var skinFrame = enemySkins[enemy.Skin].MovementByDirection[enemy.ViewingDirection.Name][supplier.lastIndex]
			return skinFrame.ToSpriteWithOffset()
		} else {
			var skinFrames = enemySkins[enemy.Skin].MovementByDirection[enemy.ViewingDirection.Name]
			supplier.durationSinceLastTile += msSinceLastSprite
			if supplier.durationSinceLastTile > DurationOfMovementFrame {
				supplier.lastIndex = (supplier.lastIndex + 1) % len(skinFrames)
			}
			return skinFrames[supplier.lastIndex].ToSpriteWithOffset()
		}
	}
}
