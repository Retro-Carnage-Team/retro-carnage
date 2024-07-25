package characters

import (
	"retro-carnage/engine/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGunTurretReturnsStaticSprite(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var gunTurret = ActiveEnemy{
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
		Type:             EnemyTypeGunTurret{},
		ViewingDirection: &geometry.UpRight,
	}

	var spriteSupplier = EnemyGunTurretSpriteSupplier{}
	assert.NotNil(t, spriteSupplier)

	var sprite = spriteSupplier.Sprite(0, gunTurret)
	assert.NotNil(t, sprite)
	assert.Equal(t, "images/environment/gun-turret-up_right.png", sprite.Source)

	sprite = spriteSupplier.Sprite(DurationOfEnemyMovementFrame*1.4, gunTurret)
	assert.NotNil(t, sprite)
	assert.Equal(t, "images/environment/gun-turret-up_right.png", sprite.Source)
}
