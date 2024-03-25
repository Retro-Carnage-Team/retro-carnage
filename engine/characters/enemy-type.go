package characters

type EnemyType int

const (
	Person    EnemyType = 0
	Landmine  EnemyType = 1
	GunTurret EnemyType = 2
)

func (et EnemyType) IsCollisionDeadly() bool {
	return et == Person || et == Landmine
}
