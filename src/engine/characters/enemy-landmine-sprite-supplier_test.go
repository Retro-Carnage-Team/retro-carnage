package characters

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/engine/geometry"
	"testing"
)

func TestLandmineReturnsStaticSprite(t *testing.T) {
	var landmine = Enemy{
		ActivationDistance:      0,
		Active:                  false,
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               []EnemyMovement{},
		Position: &geometry.Rectangle{
			X:      100,
			Y:      100,
			Width:  50,
			Height: 50,
		},
		Skin:             "",
		SpriteSupplier:   nil,
		Type:             Landmine,
		ViewingDirection: geometry.Down,
	}
	landmine.Activate() // happens when the landmine becomes visible

	var spriteSupplier = landmine.SpriteSupplier
	assert.NotNil(t, spriteSupplier)

	var sprite = spriteSupplier.Sprite(0, landmine)
	assert.NotNil(t, sprite)
	assert.Equal(t, landmineSprite, sprite.Source)

	sprite = spriteSupplier.Sprite(DurationOfMovementFrame*1.4, landmine)
	assert.NotNil(t, sprite)
	assert.Equal(t, landmineSprite, sprite.Source)
}
