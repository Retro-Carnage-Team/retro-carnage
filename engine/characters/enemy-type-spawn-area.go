package characters

import "retro-carnage/engine/geometry"

type EnemyTypeSpawnArea struct{}

func (et EnemyTypeSpawnArea) BuildEnemySpriteSupplier(viewingDirection *geometry.Direction) EnemySpriteSupplier {
	// spawn areas are not visible
	return nil
}

func (et EnemyTypeSpawnArea) CanDieWhenHitByBullet() bool {
	return false
}

func (et EnemyTypeSpawnArea) CanDieWhenHitByExplosion() bool {
	return false
}

func (et EnemyTypeSpawnArea) CanDieWhenHitByExplosive() bool {
	return false
}

func (et EnemyTypeSpawnArea) CanFire() bool {
	return false
}

func (et EnemyTypeSpawnArea) CanMove() bool {
	return false
}

func (et EnemyTypeSpawnArea) CanSpawn() bool {
	return true
}

func (et EnemyTypeSpawnArea) GetPointsForKill() int {
	return 0
}

func (et EnemyTypeSpawnArea) IsCollisionDeadly(e *ActiveEnemy) bool {
	return false
}

func (et EnemyTypeSpawnArea) IsCollisionExplosive() bool {
	return false
}

func (et EnemyTypeSpawnArea) IsObstacle() bool {
	return false
}

func (et EnemyTypeSpawnArea) IsStoppingBullets() bool {
	return false
}

func (et EnemyTypeSpawnArea) IsVisible() bool {
	return false
}

func (et EnemyTypeSpawnArea) OnActivation(e *ActiveEnemy) {
	// no logic specific to activation areas
}

func (et EnemyTypeSpawnArea) OnDeath(e *ActiveEnemy) {
	// spawn areas do not die
}

func (et EnemyTypeSpawnArea) OnMovementStopped(e *ActiveEnemy) {
	// spawn areas can not move
}
