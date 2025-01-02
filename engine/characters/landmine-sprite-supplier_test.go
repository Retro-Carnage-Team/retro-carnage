package characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandmineReturnsStaticSprite(t *testing.T) {
	InitEnemySkins("testdata/skins")

	var spriteSupplier = LandmineSpriteSupplier{}
	assert.NotNil(t, spriteSupplier)

	var sprite = spriteSupplier.Sprite(0)
	assert.NotNil(t, sprite)
	assert.Equal(t, landmineSprite, sprite.Source)

	sprite = spriteSupplier.Sprite(durationOfEnemyMovementFrame * 1.4)
	assert.NotNil(t, sprite)
	assert.Equal(t, landmineSprite, sprite.Source)
}
