package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/logging"
)

type GunTurretSpriteSupplier struct {
	enemy ActiveEnemyVisuals
}

func NewGunTurretSpriteSupplier(direction geometry.Direction, enemy ActiveEnemyVisuals) *GunTurretSpriteSupplier {
	if !direction.IsDiagonal() {
		logging.Error.Fatalf("Gun turrets can have diagonal directions, only. Found %s instead", direction.Name)
	}

	return &GunTurretSpriteSupplier{enemy: enemy}
}

func (supplier *GunTurretSpriteSupplier) GetDurationOfEnemyDeathAnimation() int64 {
	return 1
}

func (supplier *GunTurretSpriteSupplier) Sprite(msSinceLastSprite int64) *graphics.SpriteWithOffset {
	var skinFrame = enemySkins[supplier.enemy.Skin()].Idle[supplier.enemy.ViewingDirection().Name]
	return skinFrame.ToSpriteWithOffset()
}
