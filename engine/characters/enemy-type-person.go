package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

type EnemyTypePerson struct{}

func (et EnemyTypePerson) BuildEnemySpriteSupplier(viewingDirection geometry.Direction) EnemySpriteSupplier {
	return NewEnemyPersonSpriteSupplier(viewingDirection)
}

func (et EnemyTypePerson) CanDieWhenHitByBullet() bool {
	return true
}

func (et EnemyTypePerson) CanDieWhenHitByExplosion() bool {
	return true
}

func (et EnemyTypePerson) CanDieWhenHitByExplosive() bool {
	return true
}

func (et EnemyTypePerson) CanFire() bool {
	return true
}

func (et EnemyTypePerson) CanMove() bool {
	return true
}

func (et EnemyTypePerson) CanSpawn() bool {
	return false
}

func (et EnemyTypePerson) GetPointsForKill() int {
	return 10
}

func (et EnemyTypePerson) IsCollisionDeadly() bool {
	return true
}

func (et EnemyTypePerson) IsCollisionExplosive() bool {
	return false
}

func (et EnemyTypePerson) IsVisible() bool {
	return false
}

func (et EnemyTypePerson) OnActivation(e *ActiveEnemy) {
	// No logic specific to persons
}

// Is called when an enemy of this type died
func (et EnemyTypePerson) OnDeath(e *ActiveEnemy) {
	var randomDeathSound = assets.RandomEnemyDeathSoundEffect()
	assets.NewStereo().StopFx(randomDeathSound)
}
