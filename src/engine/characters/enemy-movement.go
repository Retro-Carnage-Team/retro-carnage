package characters

// EnemyMovement is (currently) not based on coordinates but on time and speed.
// This keeps the required calculation pretty simple but makes the level configuration a little harder.
type EnemyMovement struct {
	Duration     int64
	OffsetXPerMs float64
	OffsetYPerMs float64
	TimeElapsed  int64
}
