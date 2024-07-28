package characters

import "retro-carnage/engine/geometry"

type EnemyTypeGunTurret struct{}

func (et EnemyTypeGunTurret) BuildEnemySpriteSupplier(viewingDirection geometry.Direction) EnemySpriteSupplier {
	return NewGunTurretSpriteSupplier(viewingDirection)
}

func (et EnemyTypeGunTurret) CanDieWhenHitByBullet() bool {
	return false
}

func (et EnemyTypeGunTurret) CanDieWhenHitByExplosion() bool {
	return false
}

func (et EnemyTypeGunTurret) CanDieWhenHitByExplosive() bool {
	return true
}

func (et EnemyTypeGunTurret) CanFire() bool {
	return true
}

func (et EnemyTypeGunTurret) CanMove() bool {
	return false
}

func (et EnemyTypeGunTurret) CanSpawn() bool {
	return false
}

func (et EnemyTypeGunTurret) GetPointsForKill() int {
	return 15
}

func (et EnemyTypeGunTurret) IsCollisionDeadly() bool {
	return false
}

func (et EnemyTypeGunTurret) IsCollisionExplosive() bool {
	return false
}

func (et EnemyTypeGunTurret) IsVisible() bool {
	return false
}

func (et EnemyTypeGunTurret) OnActivation(e *ActiveEnemy) {
	// no logic specific to gun turrets
}

// Is called when an enemy of this type died
func (et EnemyTypeGunTurret) OnDeath(e *ActiveEnemy) {
	// no logic specific to gun turrets
}
