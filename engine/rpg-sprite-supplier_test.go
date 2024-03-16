package engine

import (
	"retro-carnage/engine/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rpg_sprite_1 = "images/weapons/rpg-up-1.png"
const rpg_sprite_2 = "images/weapons/rpg-up-2.png"

func TestRpgSpritesShouldAppearWithCorrectTiming(t *testing.T) {
	var rpgSpriteSupplier = NewRpgSpriteSupplier(geometry.Up)

	var result = rpgSpriteSupplier.Sprite(0)
	assert.Equal(t, rpg_sprite_1, result.Source)

	result = rpgSpriteSupplier.Sprite(durationOfRpgFrame - 5)
	assert.Equal(t, rpg_sprite_1, result.Source)

	result = rpgSpriteSupplier.Sprite(durationOfRpgFrame + 5)
	assert.Equal(t, rpg_sprite_2, result.Source)

	result = rpgSpriteSupplier.Sprite(durationOfRpgFrame + 5)
	assert.Equal(t, rpg_sprite_1, result.Source)
}
