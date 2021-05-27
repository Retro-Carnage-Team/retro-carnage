package characters

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"retro-carnage/engine/geometry"
	"testing"
)

func TestSpriteForIdlePlayer(t *testing.T) {
	InitPlayerSkins("../../../skins/")

	var player = Players[1]
	player.Reset()
	var behavior = NewPlayerBehavior(player)
	behavior.Moving = false
	behavior.Direction = geometry.Right
	var spriteSupplier = NewPlayerSpriteSupplier(player)
	var result = spriteSupplier.Sprite(0, behavior)

	assert.Equal(t, "images/tiles/player-1/idle/right.png", result.Source)
}

func TestSprite0IsReturnedWhenPlayerStartsToMove(t *testing.T) {
	InitPlayerSkins("../../../skins/")

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
	assert.Equal(t, "images/tiles/player-1/up/1.png", result.Source)
	result = spriteSupplier.Sprite(DurationOfPlayerMovementFrame, behavior)
	assert.Equal(t, "images/tiles/player-1/up/2.png", result.Source)
}

func TestDoesNotPlayTheAnimationWithoutDelay(t *testing.T) {
	InitPlayerSkins("../../../skins/")

	var player = Players[1]
	player.Reset()
	var behavior = NewPlayerBehavior(player)
	behavior.Moving = true
	behavior.Direction = geometry.Up
	var spriteSupplier = NewPlayerSpriteSupplier(player)

	var first = spriteSupplier.Sprite(0, behavior)
	assert.Equal(t, "images/tiles/player-1/up/1.png", first.Source)
	var second = spriteSupplier.Sprite(50, behavior)
	assert.Equal(t, "images/tiles/player-1/up/1.png", second.Source)
	var third = spriteSupplier.Sprite(50, behavior)
	assert.Equal(t, "images/tiles/player-1/up/2.png", third.Source)
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
			var expectedPath = fmt.Sprintf("images/tiles/player-1/up/%d.png", i+1)
			assert.Equal(t, expectedPath, sprite.Source)
		}
	}
}
