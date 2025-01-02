package characters

type EnemyType interface {
	// BuildEnemySpriteSupplier returns a sprite supplier for this type of enemy
	BuildEnemySpriteSupplier(enemy *ActiveEnemy) EnemySpriteSupplier

	// CanDieWhenHitByBullet returns true when enemies of this type can be killed with bullets
	CanDieWhenHitByBullet() bool

	// CanDieWhenHitByExplosion returns true when enemies of this type can be killed when they get in contact with explosions
	CanDieWhenHitByExplosion() bool

	// CanDieWhenHitByExplosive returns true when enemies of this type can be killed when are hit with explosives
	CanDieWhenHitByExplosive() bool

	// CanFire returns true when this type of enemy can fire bullets. For instances of this type of enemy actions depend on
	// the configured action pattern.
	CanFire() bool

	// CanMove returns true when this type of enemy can move. For instances of this type of enemy movements depend on the
	// configured movement pattern.
	CanMove() bool

	// CanSpawn returns true when this type of enemy can create new instances of other enemies.
	CanSpawn() bool

	// GetPointsForKill returns the number of points to be added to the player score when a player kills an enemy of this
	// type.
	GetPointsForKill() int

	// IsCollisionDeadly returns true when a collision of a player with this type of enemy leads to the players death.
	IsCollisionDeadly(e *ActiveEnemy) bool

	// IsCollisionDeadly returns true when a collision of a player with this type of enemy leads to an explosion
	IsCollisionExplosive() bool

	// IsObstacle returns true when enemies of this time become an obstacle when they're dying (like tanks)
	IsObstacle() bool

	// IsStoppingBullets returns true when enemies cannot be killed by bullets but stop them (like tanks)
	IsStoppingBullets() bool

	// IsVisible returns true when this type of enemy is displayed on screen
	IsVisible() bool

	// Is called when an enemy of this type is activated
	OnActivation(e *ActiveEnemy)

	// Is called when an enemy of this type died
	OnDeath(e *ActiveEnemy)

	// Is called when all enemy movements have been processed
	OnMovementStopped(e *ActiveEnemy)
}

var (
	enemyTypePerson    EnemyType = EnemyTypePerson{}
	enemyTypeLandmine  EnemyType = EnemyTypeLandmine{}
	enemyTypeGunTurret EnemyType = EnemyTypeGunTurret{}
	enemyTypeSpawnArea EnemyType = EnemyTypeSpawnArea{}
	enemyTypeTank      EnemyType = EnemyTypeTank{}
)

// GetEnemyTypeByCode returns an instane of EnemyType for the given type of enemy.
// Code is the integer used on level files.
func GetEnemyTypeByCode(code int) EnemyType {
	switch code {
	case 0:
		return enemyTypePerson
	case 1:
		return enemyTypeLandmine
	case 2:
		return enemyTypeGunTurret
	case 3:
		return enemyTypeSpawnArea
	case 4:
		return enemyTypeTank
	default:
		return nil
	}
}
