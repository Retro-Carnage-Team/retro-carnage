package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpritesShouldAppearWithCorrectTiming(t *testing.T) {
	var explosionSpriteSupplier = ExplosionSpriteSupplier{}

	var result = explosionSpriteSupplier.Sprite(0)
	assert.Equal(t, "images/tiles/explosion/0.png", result.Source)

	result = explosionSpriteSupplier.Sprite(DurationOfExplosionFrame - 5)
	assert.Equal(t, "images/tiles/explosion/0.png", result.Source)

	result = explosionSpriteSupplier.Sprite(DurationOfExplosionFrame - 5)
	assert.Equal(t, "images/tiles/explosion/1.png", result.Source)
}

func TestNoLoop(t *testing.T) {
	var explosionSpriteSupplier = ExplosionSpriteSupplier{}
	var result = explosionSpriteSupplier.Sprite(1199)
	assert.Equal(t, "images/tiles/explosion/47.png", result.Source)

	result = explosionSpriteSupplier.Sprite(50)
	assert.Equal(t, "images/tiles/explosion/47.png", result.Source)
}
