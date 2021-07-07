package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

const (
	DurationOfPlayerDeathAnimationFrame = 75 // in ms
	DurationOfPlayerMovementFrame       = 75 // in ms
)

// PlayerSpriteSupplier returns sprites for the current state of the Player
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

// NewPlayerSpriteSupplier creates and initializes a new PlayerSpriteSupplier for a given Player.
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

// Sprite returns the graphics.SpriteWithOffset for the current state of the Player.
func (pss *PlayerSpriteSupplier) Sprite(elapsedTimeInMs int64, behavior *PlayerBehavior) *graphics.SpriteWithOffset {
	// First we remember which sprite is currently rendered - then get the next sprite.
	var lastSpriteSource = ""
	if nil != pss.lastSprite {
		lastSpriteSource = pss.lastSprite.Source
	}
	var nextSprite = pss.sprite(elapsedTimeInMs, behavior)

	// Now we drop every second sprite - when player is invincible
	if behavior.Invincible && lastSpriteSource != nextSprite.Source {
		pss.invincibilityToggle = !pss.invincibilityToggle
		if pss.invincibilityToggle {
			return nil
		}
	}

	return nextSprite
}

func (pss *PlayerSpriteSupplier) sprite(elapsedTimeInMs int64, behavior *PlayerBehavior) *graphics.SpriteWithOffset {
	if behavior.Idle() {
		pss.wasMoving = false
		var skinFrame = pss.skin.Idle[behavior.Direction.Name]
		return skinFrame.ToSpriteWithOffset()
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
