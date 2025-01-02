package graphics

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
)

type GunTurretSpriteSupplier struct {
	enemy EnemyVisuals
}

func NewGunTurretSpriteSupplier(direction geometry.Direction, enemy EnemyVisuals) *GunTurretSpriteSupplier {
	if !direction.IsDiagonal() {
		logging.Error.Fatalf("Gun turrets can have diagonal directions, only. Found %s instead", direction.Name)
	}

	return &GunTurretSpriteSupplier{enemy: enemy}
}

func (supplier *GunTurretSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return 1
}

func (supplier *GunTurretSpriteSupplier) Sprite(msSinceLastSprite int64) *SpriteWithOffset {
	var skinFrame = enemySkins[supplier.enemy.Skin()].Idle[supplier.enemy.ViewingDirection().Name]
	return skinFrame.ToSpriteWithOffset()
}
