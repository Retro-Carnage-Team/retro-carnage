package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

type EnemyTypeTank struct{}

func (et EnemyTypeTank) BuildEnemySpriteSupplier(viewingDirection geometry.Direction) EnemySpriteSupplier {
	return NewTankSpriteSupplier(viewingDirection)
}

func (et EnemyTypeTank) CanDieWhenHitByBullet() bool {
	return false
}

func (et EnemyTypeTank) CanDieWhenHitByExplosion() bool {
	return false
}

func (et EnemyTypeTank) CanDieWhenHitByExplosive() bool {
	return true
}

func (et EnemyTypeTank) CanFire() bool {
	return true
}

func (et EnemyTypeTank) CanMove() bool {
	return true
}

func (et EnemyTypeTank) CanSpawn() bool {
	return false
}

func (et EnemyTypeTank) GetPointsForKill() int {
	return 100
}

func (et EnemyTypeTank) IsCollisionDeadly() bool {
	return true
}

func (et EnemyTypeTank) IsCollisionExplosive() bool {
	return false
}

func (et EnemyTypeTank) IsVisible() bool {
	return false
}

func (et EnemyTypeTank) OnActivation(e *ActiveEnemy) {
	var stereoPlayer = assets.NewStereo()
	stereoPlayer.PlayFx(assets.FxTankMoving)
}

// Is called when an enemy of this type died
func (et EnemyTypeTank) OnDeath(e *ActiveEnemy) {
	var stereoPlayer = assets.NewStereo()
	stereoPlayer.StopFx(assets.FxTankMoving)
}
