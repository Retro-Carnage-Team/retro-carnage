package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/util"
)

const (
	DurationOfPlayerDeathAnimationFrame = 75 // in ms
	DurationOfPlayerMovementFrame       = 75 // in ms
)

// PlayerSpriteSupplier returns sprites for the current state of the Player
type PlayerSpriteSupplier struct {
	directionOfLastSprite        geometry.Direction
	durationOfInvincibilityState int64
	durationSinceLastSprite      int64
	durationSinceLastToggle      int64
	invincibilityToggle          bool
	lastIndex                    int
	lastSprite                   *graphics.SpriteWithOffset
	skin                         *graphics.Skin
	wasDying                     bool
	wasMoving                    bool
}

// NewPlayerSpriteSupplier creates and initializes a new PlayerSpriteSupplier for a given Player.
func NewPlayerSpriteSupplier(player *Player, durationOfInvincibilityState int64) *PlayerSpriteSupplier {
	return &PlayerSpriteSupplier{
		directionOfLastSprite:        geometry.Up,
		durationOfInvincibilityState: durationOfInvincibilityState,
		durationSinceLastSprite:      0,
		durationSinceLastToggle:      0,
		invincibilityToggle:          true,
		lastIndex:                    0,
		lastSprite:                   nil,
		skin:                         graphics.PlayerSkins[player.index],
		wasDying:                     false,
		wasMoving:                    false,
	}
}

// Sprite returns the graphics.SpriteWithOffset for the current state of the Player.
func (pss *PlayerSpriteSupplier) Sprite(elapsedTimeInMs int64, behavior *PlayerBehavior) *graphics.SpriteWithOffset {
	var nextSprite = pss.sprite(elapsedTimeInMs, behavior)
	if behavior.Invincible {
		pss.durationSinceLastToggle += elapsedTimeInMs
		if pss.durationSinceLastToggle >= pss.durationOfInvincibilityState {
			pss.durationSinceLastToggle = 0
			pss.invincibilityToggle = !pss.invincibilityToggle
		}

		if pss.invincibilityToggle {
			return nil
		}
	}

	return nextSprite
}

func (pss *PlayerSpriteSupplier) sprite(elapsedTimeInMs int64, behavior *PlayerBehavior) *graphics.SpriteWithOffset {
	pss.durationSinceLastSprite += elapsedTimeInMs

	if behavior.Dying {
		return pss.spriteForDyingPlayer()
	}
	pss.wasDying = false

	if behavior.Idle() {
		pss.wasMoving = false
		var skinFrame = pss.skin.Idle[behavior.Direction.Name]
		return skinFrame.ToSpriteWithOffset()
	} else {
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
	var deathFrames = pss.skin.DeathAnimation[geometry.Up.Name]
	if pss.wasDying {
		if DurationOfPlayerDeathAnimationFrame <= pss.durationSinceLastSprite {
			pss.lastIndex = util.MinInt(pss.lastIndex+1, len(deathFrames)-1)
			pss.lastSprite = deathFrames[pss.lastIndex].ToSpriteWithOffset()
			pss.durationSinceLastSprite = 0
		}
	} else {
		pss.durationSinceLastSprite = 0
		pss.lastIndex = 0
		pss.lastSprite = deathFrames[0].ToSpriteWithOffset()
		pss.wasDying = true
	}
	return pss.lastSprite
}
