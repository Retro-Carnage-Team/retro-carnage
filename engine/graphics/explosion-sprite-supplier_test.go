package graphics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpritesShouldAppearWithCorrectTiming(t *testing.T) {
	var explosionSpriteSupplier = ExplosionSpriteSupplier{}

	var result = explosionSpriteSupplier.Sprite(0)
	assert.Equal(t, "images/explosion/0.png", result.Source)

	result = explosionSpriteSupplier.Sprite(DurationOfExplosionFrame - 5)
	assert.Equal(t, "images/explosion/0.png", result.Source)

	result = explosionSpriteSupplier.Sprite(DurationOfExplosionFrame - 5)
	assert.Equal(t, "images/explosion/1.png", result.Source)
}

func TestNoLoop(t *testing.T) {
	var explosionSpriteSupplier = ExplosionSpriteSupplier{}
	var result = explosionSpriteSupplier.Sprite(1199)
	assert.Equal(t, "images/explosion/47.png", result.Source)

	result = explosionSpriteSupplier.Sprite(50)
	assert.Equal(t, "images/explosion/47.png", result.Source)
}
