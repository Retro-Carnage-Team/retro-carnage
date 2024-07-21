package characters

type EnemyType int

const (
	Person    EnemyType = 0
	Landmine  EnemyType = 1
	GunTurret EnemyType = 2
	SpawnArea EnemyType = 3
)

var pointsByEnemyType map[EnemyType]int

func init() {
	pointsByEnemyType = make(map[EnemyType]int)
	pointsByEnemyType[Person] = 10
	pointsByEnemyType[Landmine] = 5
	pointsByEnemyType[GunTurret] = 15
}

func (et EnemyType) CanDie() bool {
	return et.CanDieWhenHitByBullet() || et.CanDieWhenHitByExplosion() || et.CanDieWhenHitByExplosive()
}

func (et EnemyType) CanDieWhenHitByBullet() bool {
	switch et {
	case Person:
		return true
	case Landmine:
		return false
	case GunTurret:
		return false
	case SpawnArea:
		return false
	default:
		return false
	}
}

func (et EnemyType) CanDieWhenHitByExplosion() bool {
	switch et {
	case Person:
		return true
	case Landmine:
		return true
	case GunTurret:
		return true
	case SpawnArea:
		return false
	default:
		return false
	}
}

func (et EnemyType) CanDieWhenHitByExplosive() bool {
	switch et {
	case Person:
		return true
	case Landmine:
		return true
	case GunTurret:
		return true
	case SpawnArea:
		return false
	default:
		return false
	}
}

// CanFire returns true when this type of enemy can fire bullets. For instances of this type of enemy actions depend on
// the configured action pattern.
func (et EnemyType) CanFire() bool {
	return et == Person || et == GunTurret
}

// CanMove returns true when this type of enemy can move. For instances of this type of enemy movements depend on the
// configured movement pattern.
func (et EnemyType) CanMove() bool {
	return et == Person
}

// CanSpawn returns true when this type of enemy can create new instances of other enemies.
func (et EnemyType) CanSpawn() bool {
	return et == SpawnArea
}

// GetPointsForKill returns the number of points to be added to the player score when a player kills an enemy of this
// type.
func (et EnemyType) GetPointsForKill() int {
	return pointsByEnemyType[et]
}

// IsCollisionDeadly returns true when a collision of a player with this type of enemy leads to the players death.
func (et EnemyType) IsCollisionDeadly() bool {
	return et == Person || et == Landmine
}

func (et EnemyType) IsVisible() bool {
	return et != SpawnArea
}
