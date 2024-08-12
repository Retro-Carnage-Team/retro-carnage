package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/logging"
)

type GunTurretSpriteSupplier struct {
}

func NewGunTurretSpriteSupplier(direction geometry.Direction) *GunTurretSpriteSupplier {
	if !direction.IsDiagonal() {
		logging.Error.Fatalf("Gun turrets can have diagonal directions, only. Found %s instead", direction.Name)
	}

	return &GunTurretSpriteSupplier{}
}

func (supplier *GunTurretSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return 1
}

func (supplier *GunTurretSpriteSupplier) Sprite(msSinceLastSprite int64, enemy ActiveEnemy) *graphics.SpriteWithOffset {
	var skinFrame = enemySkins[enemy.Skin].Idle[enemy.ViewingDirection.Name]
	return skinFrame.ToSpriteWithOffset()
}
