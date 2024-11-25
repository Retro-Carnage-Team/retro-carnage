package characters

import (
	"fmt"
	"os"
	"path/filepath"
	"retro-carnage/engine/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

const RC_ASSETS = "RC-ASSETS"
const PLAYER1_IDLE_UP = "images/player-1/idle/up.png"
const PLAYER1_UP1 = "images/player-1/up/1.png"
const PLAYER1_UP2 = "images/player-1/up/2.png"

func TestSpriteForIdlePlayer(t *testing.T) {
	InitPlayerSkins(filepath.Join(os.Getenv(RC_ASSETS), "skins"))

	var player = Players[1]
	player.Reset()
	var behavior = NewPlayerBehavior(player)
	behavior.Moving = false
	behavior.Direction = geometry.Right
	var spriteSupplier = NewPlayerSpriteSupplier(player, 200)
	var result = spriteSupplier.Sprite(0, behavior)

	assert.Equal(t, "images/player-1/idle/right.png", result.Source)
}

func TestSprite0IsReturnedWhenPlayerStartsToMove(t *testing.T) {
	InitPlayerSkins(filepath.Join(os.Getenv(RC_ASSETS), "skins"))

	var player = Players[1]
	player.Reset()
	var behavior = NewPlayerBehavior(player)
	behavior.Moving = false
	behavior.Direction = geometry.Right
	var spriteSupplier = NewPlayerSpriteSupplier(player, 200)
	spriteSupplier.Sprite(0, behavior)

	behavior.Moving = true
	behavior.Direction = geometry.Up
	var result = spriteSupplier.Sprite(DurationOfPlayerMovementFrame, behavior)
	assert.Equal(t, PLAYER1_UP1, result.Source)
	result = spriteSupplier.Sprite(DurationOfPlayerMovementFrame, behavior)
	assert.Equal(t, PLAYER1_UP2, result.Source)
}

func TestDoesNotPlayTheAnimationWithoutDelay(t *testing.T) {
	InitPlayerSkins(filepath.Join(os.Getenv(RC_ASSETS), "skins"))

	var player = Players[1]
	player.Reset()
	var behavior = NewPlayerBehavior(player)
	behavior.Moving = true
	behavior.Direction = geometry.Up
	var spriteSupplier = NewPlayerSpriteSupplier(player, 200)

	var first = spriteSupplier.Sprite(0, behavior)
	assert.Equal(t, PLAYER1_UP1, first.Source)
	var second = spriteSupplier.Sprite(50, behavior)
	assert.Equal(t, PLAYER1_UP1, second.Source)
	var third = spriteSupplier.Sprite(50, behavior)
	assert.Equal(t, PLAYER1_UP2, third.Source)
}

func TestPlaysAnimationInLoop(t *testing.T) {
	var player = Players[1]
	player.Reset()
	var behavior = NewPlayerBehavior(player)
	behavior.Moving = true
	behavior.Direction = geometry.Up
	var spriteSupplier = NewPlayerSpriteSupplier(player, 200)

	var spriteCount = len(spriteSupplier.skin.MovementByDirection[geometry.Up.Name])
	assert.True(t, 0 < spriteCount)

	for j := 0; j < 2; j++ {
		for i := 0; i < spriteCount; i++ {
			var sprite = spriteSupplier.Sprite(DurationOfPlayerMovementFrame, behavior)
			var expectedPath = fmt.Sprintf("images/player-1/up/%d.png", i+1)
			assert.Equal(t, expectedPath, sprite.Source)
		}
	}
}

func TestInvincibilityAnimationWithoutMovement(t *testing.T) {
	InitPlayerSkins(filepath.Join(os.Getenv(RC_ASSETS), "skins"))

	var player = Players[1]
	player.Reset()

	var behavior = NewPlayerBehavior(player)
	behavior.Direction = geometry.Up
	behavior.Moving = false
	behavior.StartInvincibility()

	var spriteSupplier = NewPlayerSpriteSupplier(player, 200)

	// Invicibility lasts for 1.500 ms
	AssertNextSpriteIsNil(0, spriteSupplier, behavior, t)
	AssertNextSpriteIsNil(100, spriteSupplier, behavior, t)

	// After 200 ms the sprite has to be returned
	AssertNextSpriteIsIdleUp(100, spriteSupplier, behavior, t)
	AssertNextSpriteIsIdleUp(100, spriteSupplier, behavior, t)

	// After 400 ms nil has to be returned
	AssertNextSpriteIsNil(100, spriteSupplier, behavior, t)
	AssertNextSpriteIsNil(100, spriteSupplier, behavior, t)

	// After 600 ms the sprite has to be returned
	AssertNextSpriteIsIdleUp(100, spriteSupplier, behavior, t)
	AssertNextSpriteIsIdleUp(100, spriteSupplier, behavior, t)

	// After 800 ms nil has to be returned
	AssertNextSpriteIsNil(100, spriteSupplier, behavior, t)
	AssertNextSpriteIsNil(100, spriteSupplier, behavior, t)

	// After 1.000 ms the sprite has to be returned
	AssertNextSpriteIsIdleUp(100, spriteSupplier, behavior, t)
	AssertNextSpriteIsIdleUp(100, spriteSupplier, behavior, t)

	// After 1.200 ms nil has to be returned
	AssertNextSpriteIsNil(100, spriteSupplier, behavior, t)
	AssertNextSpriteIsNil(100, spriteSupplier, behavior, t)

	// After 1.400 ms the sprite has to be returned
	AssertNextSpriteIsIdleUp(100, spriteSupplier, behavior, t)
	AssertNextSpriteIsIdleUp(100, spriteSupplier, behavior, t)

	// Player is no longer invicible. Test that result is no longer toggled.
	AssertNextSpriteIsIdleUp(150, spriteSupplier, behavior, t)
	AssertNextSpriteIsIdleUp(150, spriteSupplier, behavior, t)
}

func AssertNextSpriteIsNil(elapsedTimeInMs int64, spriteSupplier *PlayerSpriteSupplier, behavior *PlayerBehavior, t *testing.T) {
	var result = spriteSupplier.Sprite(elapsedTimeInMs, behavior)
	behavior.UpdateInvincibility(elapsedTimeInMs)
	assert.Nil(t, result)
}

func AssertNextSpriteIsIdleUp(elapsedTimeInMs int64, spriteSupplier *PlayerSpriteSupplier, behavior *PlayerBehavior, t *testing.T) {
	var result = spriteSupplier.Sprite(elapsedTimeInMs, behavior)
	behavior.UpdateInvincibility(elapsedTimeInMs)
	assert.NotNil(t, result)
	assert.Equal(t, PLAYER1_IDLE_UP, result.Source)
}
