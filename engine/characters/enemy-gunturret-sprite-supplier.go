package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/logging"
)

type EnemyGunTurretSpriteSupplier struct {
}

func NewEnemyGunTurretSpriteSupplier(direction geometry.Direction) *EnemyGunTurretSpriteSupplier {
	if !direction.IsDiagonal() {
		logging.Error.Fatalf("Gun turrets can have diagonal directions, only. Found %s instead", direction.Name)
	}

	return &EnemyGunTurretSpriteSupplier{}
}

func (supplier *EnemyGunTurretSpriteSupplier) Sprite(msSinceLastSprite int64, enemy ActiveEnemy) *graphics.SpriteWithOffset {
	var skinFrame = enemySkins[enemy.Skin].Idle[enemy.ViewingDirection.Name]
	return skinFrame.ToSpriteWithOffset()
}
