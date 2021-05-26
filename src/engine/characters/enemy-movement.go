package characters

import "retro-carnage/assets"

// EnemyMovement is (currently) not based on coordinates but on time and speed.
// This keeps the required calculation pretty simple but makes the level configuration a little harder.
type EnemyMovement struct {
	*assets.EnemyMovement
	TimeElapsed int64
}

func NewEnemyMovement(e *assets.EnemyMovement) EnemyMovement {
	return EnemyMovement{
		EnemyMovement: e,
		TimeElapsed:   0,
	}
}
