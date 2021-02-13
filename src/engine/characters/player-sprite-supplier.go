package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type PlayerSpriteSupplier struct {
	directionOfLastTile   geometry.Direction
	durationSinceLastTile int64
	invincibilityToggle   bool
	lastIndex             int
	lastSprite            *graphics.SpriteWithOffset
	skin                  *Skin
	wasDying              bool
}

func NewPlayerSpriteSupplier(player *Player) *PlayerSpriteSupplier {
	return &PlayerSpriteSupplier{
		directionOfLastTile:   geometry.Up,
		durationSinceLastTile: 0,
		invincibilityToggle:   true,
		lastIndex:             0,
		lastSprite:            nil,
		skin:                  playerSkins[player.index],
		wasDying:              false,
	}
}

func (pss *PlayerSpriteSupplier) Sprite(elapsedTimeInMs int64, behavior *PlayerBehavior) *graphics.SpriteWithOffset {
	if behavior.Idle() {
		return pss.spriteForIdlePlayer(behavior)
	} else {
		pss.durationSinceLastTile += elapsedTimeInMs
		if DurationOfMovementFrame <= pss.durationSinceLastTile {
			pss.durationSinceLastTile = 0
			var frames = pss.skin.MovementByDirection[behavior.Direction.Name]
			pss.lastIndex = (pss.lastIndex + 1) % len(frames)
			pss.lastSprite = frames[pss.lastIndex].ToSpriteWithOffset()
		}

		if behavior.Dying {
			return pss.spriteForDyingPlayer()
		}

		pss.wasDying = false
		if pss.directionOfLastTile != behavior.Direction {
			pss.directionOfLastTile = behavior.Direction
			pss.durationSinceLastTile = 0
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
		if DurationOfDeathAnimationFrame <= pss.durationSinceLastTile {
			pss.lastIndex = (pss.lastIndex + 1) % len(deathFrames)
			pss.lastSprite = deathFrames[pss.lastIndex].ToSpriteWithOffset()
		}
	} else {
		pss.durationSinceLastTile = 0
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
