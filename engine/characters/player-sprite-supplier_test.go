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
const PLAYER1_UP1 = "images/player-1/up/1.png"
const PLAYER1_UP2 = "images/player-1/up/2.png"

func TestSpriteForIdlePlayer(t *testing.T) {
	InitPlayerSkins(filepath.Join(os.Getenv(RC_ASSETS), "skins"))

	var player = Players[1]
	player.Reset()
	var behavior = NewPlayerBehavior(player)
	behavior.Moving = false
	behavior.Direction = geometry.Right
	var spriteSupplier = NewPlayerSpriteSupplier(player)
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
	var spriteSupplier = NewPlayerSpriteSupplier(player)
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
	var spriteSupplier = NewPlayerSpriteSupplier(player)

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
	var spriteSupplier = NewPlayerSpriteSupplier(player)

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
