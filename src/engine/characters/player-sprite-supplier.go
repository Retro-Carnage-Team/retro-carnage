package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

const (
	DurationOfPlayerDeathAnimationFrame = 75 // in ms
	DurationOfPlayerMovementFrame       = 75 // in ms
)

type PlayerSpriteSupplier struct {
	directionOfLastSprite   geometry.Direction
	durationSinceLastSprite int64
	invincibilityToggle     bool
	lastIndex               int
	lastSprite              *graphics.SpriteWithOffset
	skin                    *Skin
	wasDying                bool
	wasMoving               bool
}

func NewPlayerSpriteSupplier(player *Player) *PlayerSpriteSupplier {
	return &PlayerSpriteSupplier{
		directionOfLastSprite:   geometry.Up,
		durationSinceLastSprite: 0,
		invincibilityToggle:     true,
		lastIndex:               0,
		lastSprite:              nil,
		skin:                    playerSkins[player.index],
		wasDying:                false,
		wasMoving:               false,
	}
}

func (pss *PlayerSpriteSupplier) Sprite(elapsedTimeInMs int64, behavior *PlayerBehavior) *graphics.SpriteWithOffset {
	if behavior.Idle() {
		pss.wasMoving = false
		return pss.spriteForIdlePlayer(behavior)
	} else {
		pss.durationSinceLastSprite += elapsedTimeInMs
		if (DurationOfPlayerMovementFrame <= pss.durationSinceLastSprite) || (nil == pss.lastSprite) {
			pss.durationSinceLastSprite = 0
			var frames = pss.skin.MovementByDirection[behavior.Direction.Name]
			if pss.wasMoving {
				pss.lastIndex = (pss.lastIndex + 1) % len(frames)
			} else {
				pss.lastIndex = 0
				pss.wasMoving = true
			}
			pss.lastSprite = frames[pss.lastIndex].ToSpriteWithOffset()
		}

		if behavior.Dying {
			return pss.spriteForDyingPlayer()
		}

		pss.wasDying = false
		if pss.directionOfLastSprite != behavior.Direction {
			pss.directionOfLastSprite = behavior.Direction
			pss.durationSinceLastSprite = 0
			pss.lastIndex = 0
			var frames = pss.skin.MovementByDirection[behavior.Direction.Name]
			pss.lastSprite = frames[pss.lastIndex].ToSpriteWithOffset()
		}

		return pss.lastSprite
	}
}

func (pss *PlayerSpriteSupplier) spriteForDyingPlayer() *graphics.SpriteWithOffset {
	var deathFrames = pss.skin.DeathAnimation
	if pss.wasDying {
		if DurationOfPlayerDeathAnimationFrame <= pss.durationSinceLastSprite {
			pss.lastIndex = (pss.lastIndex + 1) % len(deathFrames)
			pss.lastSprite = deathFrames[pss.lastIndex].ToSpriteWithOffset()
		}
	} else {
		pss.durationSinceLastSprite = 0
		pss.lastIndex = 0
		pss.lastSprite = deathFrames[pss.lastIndex].ToSpriteWithOffset()
		pss.wasDying = true
	}
	return pss.lastSprite
}

func (pss *PlayerSpriteSupplier) spriteForIdlePlayer(behavior *PlayerBehavior) *graphics.SpriteWithOffset {
	// TODO: Improve this: The invincibility toggle makes the player flicker with the frequency of the screen.
	if behavior.Invincible {
		pss.invincibilityToggle = !pss.invincibilityToggle
		if pss.invincibilityToggle {
			return nil
		}
	}

	var skinFrame = pss.skin.Idle[behavior.Direction.Name]
	return skinFrame.ToSpriteWithOffset()
}
