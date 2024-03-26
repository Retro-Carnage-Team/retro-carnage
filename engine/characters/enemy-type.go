package characters

type EnemyType int

const (
	Person    EnemyType = 0
	Landmine  EnemyType = 1
	GunTurret EnemyType = 2
)

var pointsByEnemyType map[EnemyType]int

func init() {
	pointsByEnemyType = make(map[EnemyType]int)
	pointsByEnemyType[Person] = 10
	pointsByEnemyType[Landmine] = 5
	pointsByEnemyType[GunTurret] = 15
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

// GetPointsForKill returns the number of points to be added to the player score when a player kills an enemy of this
// type.
func (et EnemyType) GetPointsForKill() int {
	return pointsByEnemyType[et]
}

// IsCollisionDeadly returns true when a collision of a player with this type of enemy leads to the players death.
func (et EnemyType) IsCollisionDeadly() bool {
	return et == Person || et == Landmine
}
