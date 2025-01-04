package graphics

import (
	"retro-carnage/engine/geometry"
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
	lastSprite                   *SpriteWithOffset
	skin                         *Skin
	wasDying                     bool
	wasMoving                    bool
}

// NewPlayerSpriteSupplier creates and initializes a new PlayerSpriteSupplier for a given Player.
func NewPlayerSpriteSupplier(player PlayerVisuals, durationOfInvincibilityState int64) *PlayerSpriteSupplier {
	return &PlayerSpriteSupplier{
		directionOfLastSprite:        geometry.Up,
		durationOfInvincibilityState: durationOfInvincibilityState,
		durationSinceLastSprite:      0,
		durationSinceLastToggle:      0,
		invincibilityToggle:          true,
		lastIndex:                    0,
		lastSprite:                   nil,
		skin:                         PlayerSkins[player.PlayerIndex()],
		wasDying:                     false,
		wasMoving:                    false,
	}
}

// Sprite returns the graphics.SpriteWithOffset for the current state of the Player.
func (pss *PlayerSpriteSupplier) Sprite(elapsedTimeInMs int64, player PlayerVisuals) *SpriteWithOffset {
	var nextSprite = pss.sprite(elapsedTimeInMs, player)
	if player.Invincible() {
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

func (pss *PlayerSpriteSupplier) sprite(elapsedTimeInMs int64, player PlayerVisuals) *SpriteWithOffset {
	pss.durationSinceLastSprite += elapsedTimeInMs

	if player.Dying() {
		return pss.spriteForDyingPlayer()
	}
	pss.wasDying = false

	if player.Idle() {
		pss.wasMoving = false
		var skinFrame = pss.skin.Idle[player.ViewingDirection().Name]
		return skinFrame.ToSpriteWithOffset()
	} else {
		if (DurationOfPlayerMovementFrame <= pss.durationSinceLastSprite) || (nil == pss.lastSprite) {
			pss.durationSinceLastSprite = 0
			var frames = pss.skin.MovementByDirection[player.ViewingDirection().Name]
			if pss.wasMoving {
				pss.lastIndex = (pss.lastIndex + 1) % len(frames)
			} else {
				pss.lastIndex = 0
				pss.wasMoving = true
			}
			pss.lastSprite = frames[pss.lastIndex].ToSpriteWithOffset()
		}

		if pss.directionOfLastSprite != *player.ViewingDirection() {
			pss.directionOfLastSprite = *player.ViewingDirection()
			pss.durationSinceLastSprite = 0
			pss.lastIndex = 0
			var frames = pss.skin.MovementByDirection[player.ViewingDirection().Name]
			pss.lastSprite = frames[pss.lastIndex].ToSpriteWithOffset()
		}

		return pss.lastSprite
	}
}

func (pss *PlayerSpriteSupplier) spriteForDyingPlayer() *SpriteWithOffset {
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
