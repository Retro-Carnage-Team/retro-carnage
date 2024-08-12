package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

type EnemyTypeTank struct{}

func (et EnemyTypeTank) BuildEnemySpriteSupplier(viewingDirection *geometry.Direction) EnemySpriteSupplier {
	return NewTankSpriteSupplier(*viewingDirection)
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

func (et EnemyTypeTank) IsCollisionDeadly(e *ActiveEnemy) bool {
	return len(e.Movements) > 0
}

func (et EnemyTypeTank) IsCollisionExplosive() bool {
	return false
}

func (et EnemyTypeTank) IsObstacle() bool {
	return true
}

func (et EnemyTypeTank) IsStoppingBullets() bool {
	return true
}

func (et EnemyTypeTank) IsVisible() bool {
	return false
}

func (et EnemyTypeTank) OnActivation(e *ActiveEnemy) {
	if len(e.Movements) > 0 {
		var stereoPlayer = assets.NewStereo()
		stereoPlayer.PlayFx(assets.FxTankMoving)
	}
}

func (et EnemyTypeTank) OnDeath(e *ActiveEnemy) {
	assets.NewStereo().StopFx(assets.FxTankMoving)
}

func (et EnemyTypeTank) OnMovementStopped(e *ActiveEnemy) {
	assets.NewStereo().StopFx(assets.FxTankMoving)
}
