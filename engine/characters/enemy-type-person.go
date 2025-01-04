package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/graphics"
)

type EnemyTypePerson struct{}

func (et EnemyTypePerson) BuildEnemySpriteSupplier(enemy *ActiveEnemy) graphics.EnemySpriteSupplier {
	return graphics.NewPersonSpriteSupplier(ActiveEnemyVisuals{activeEnemy: enemy})
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

func (et EnemyTypePerson) IsCollisionDeadly(e *ActiveEnemy) bool {
	return true
}

func (et EnemyTypePerson) IsCollisionExplosive() bool {
	return false
}

func (et EnemyTypePerson) IsObstacle() bool {
	return false
}

func (et EnemyTypePerson) IsStoppingBullets() bool {
	return false
}

func (et EnemyTypePerson) IsVisible() bool {
	return true
}

func (et EnemyTypePerson) OnActivation(e *ActiveEnemy) {
	// No logic specific to persons
}

func (et EnemyTypePerson) OnDeath(e *ActiveEnemy) {
	var randomDeathSound = assets.RandomEnemyDeathSoundEffect()
	assets.NewStereo().PlayFx(randomDeathSound)
}

func (et EnemyTypePerson) OnMovementStopped(e *ActiveEnemy) {
	// No logic specific to persons
}
