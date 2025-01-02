package graphics

import (
	"retro-carnage/engine/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGunTurretReturnsStaticSprite(t *testing.T) {
	InitEnemySkins("testdata/skins")
	var gunTurret = mockGunTurret{
		viewingDirection: geometry.UpRight,
	}

	var spriteSupplier = GunTurretSpriteSupplier{enemy: &gunTurret}
	assert.NotNil(t, spriteSupplier)

	var sprite = spriteSupplier.Sprite(0)
	assert.NotNil(t, sprite)
	assert.Equal(t, "images/environment/gun-turret-up_right.png", sprite.Source)

	sprite = spriteSupplier.Sprite(durationOfEnemyMovementFrame * 1.4)
	assert.NotNil(t, sprite)
	assert.Equal(t, "images/environment/gun-turret-up_right.png", sprite.Source)
}

type mockGunTurret struct {
	skin             string
	viewingDirection geometry.Direction
}

func (mgt *mockGunTurret) Dying() bool {
	return false
}

func (mgt *mockGunTurret) Skin() EnemySkin {
	return GunTurretSkin
}

func (mgt *mockGunTurret) ViewingDirection() *geometry.Direction {
	return &mgt.viewingDirection
}
