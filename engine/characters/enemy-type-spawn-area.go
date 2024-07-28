package characters

import "retro-carnage/engine/geometry"

type EnemyTypeSpawnArea struct{}

func (et EnemyTypeSpawnArea) BuildEnemySpriteSupplier(viewingDirection geometry.Direction) EnemySpriteSupplier {
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
	return false
}

func (et EnemyTypeSpawnArea) GetPointsForKill() int {
	return 0
}

func (et EnemyTypeSpawnArea) IsCollisionDeadly() bool {
	return false
}

func (et EnemyTypeSpawnArea) IsCollisionExplosive() bool {
	return false
}

func (et EnemyTypeSpawnArea) IsVisible() bool {
	return true
}

func (et EnemyTypeSpawnArea) OnActivation(e *ActiveEnemy) {
	// no logic specific to activation areas
}

// Is called when an enemy of this type died
func (et EnemyTypeSpawnArea) OnDeath(e *ActiveEnemy) {
	// spawn areas do not die
}
