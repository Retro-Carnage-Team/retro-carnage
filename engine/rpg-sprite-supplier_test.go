package engine

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/engine/geometry"
	"testing"
)

func TestRpgSpritesShouldAppearWithCorrectTiming(t *testing.T) {
	var rpgSpriteSupplier = NewRpgSpriteSupplier(geometry.Up)

	var result = rpgSpriteSupplier.Sprite(0)
	assert.Equal(t, "images/weapons/rpg-up-1.png", result.Source)

	result = rpgSpriteSupplier.Sprite(durationOfRpgFrame - 5)
	assert.Equal(t, "images/weapons/rpg-up-1.png", result.Source)

	result = rpgSpriteSupplier.Sprite(durationOfRpgFrame + 5)
	assert.Equal(t, "images/weapons/rpg-up-2.png", result.Source)

	result = rpgSpriteSupplier.Sprite(durationOfRpgFrame + 5)
	assert.Equal(t, "images/weapons/rpg-up-1.png", result.Source)
}
