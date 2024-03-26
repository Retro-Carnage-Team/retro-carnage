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

// GetPointsForKill returns the number of points to be added to the player score when a player kills an enemy of this
// type.
func (et EnemyType) GetPointsForKill() int {
	return pointsByEnemyType[et]
}

// IsCollisionDeadly return true when a collision of a player with this type of enemy leads to the players death.
func (et EnemyType) IsCollisionDeadly() bool {
	return et == Person || et == Landmine
}
