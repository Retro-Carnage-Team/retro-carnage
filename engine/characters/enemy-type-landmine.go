package characters

import "retro-carnage/engine/geometry"

type EnemyTypeLandmine struct{}

func (et EnemyTypeLandmine) BuildEnemySpriteSupplier(viewingDirection *geometry.Direction) EnemySpriteSupplier {
	return &LandmineSpriteSupplier{}
}

func (et EnemyTypeLandmine) CanDieWhenHitByBullet() bool {
	return false
}

func (et EnemyTypeLandmine) CanDieWhenHitByExplosion() bool {
	return true
}

func (et EnemyTypeLandmine) CanDieWhenHitByExplosive() bool {
	return true
}

func (et EnemyTypeLandmine) CanFire() bool {
	return false
}

func (et EnemyTypeLandmine) CanMove() bool {
	return false
}

func (et EnemyTypeLandmine) CanSpawn() bool {
	return false
}

func (et EnemyTypeLandmine) GetPointsForKill() int {
	return 5
}

func (et EnemyTypeLandmine) IsCollisionDeadly(e *ActiveEnemy) bool {
	return true
}

func (et EnemyTypeLandmine) IsCollisionExplosive() bool {
	return true
}

func (et EnemyTypeLandmine) IsObstacle() bool {
	return false
}

func (et EnemyTypeLandmine) IsStoppingBullets() bool {
	return false
}

func (et EnemyTypeLandmine) IsVisible() bool {
	return true
}

func (et EnemyTypeLandmine) OnActivation(e *ActiveEnemy) {
	// no logic specific to landmines
}

func (et EnemyTypeLandmine) OnDeath(e *ActiveEnemy) {
	// no logic specific to landmines
}

func (et EnemyTypeLandmine) OnMovementStopped(e *ActiveEnemy) {
	// landmines can not move
}
