package characters

import (
	"retro-carnage/engine/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandmineReturnsStaticSprite(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var landmine = ActiveEnemy{
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               []EnemyMovement{},
		position: geometry.Rectangle{
			X:      100,
			Y:      100,
			Width:  50,
			Height: 50,
		},
		Skin:             "",
		SpriteSupplier:   nil,
		Type:             EnemyTypeLandmine{},
		ViewingDirection: &geometry.Down,
	}

	var spriteSupplier = EnemyLandmineSpriteSupplier{}
	assert.NotNil(t, spriteSupplier)

	var sprite = spriteSupplier.Sprite(0, landmine)
	assert.NotNil(t, sprite)
	assert.Equal(t, landmineSprite, sprite.Source)

	sprite = spriteSupplier.Sprite(DurationOfEnemyMovementFrame*1.4, landmine)
	assert.NotNil(t, sprite)
	assert.Equal(t, landmineSprite, sprite.Source)
}
