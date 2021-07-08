package assets

// EnemyMovement is one of a series of movements of an Enemy.
type EnemyMovement struct {
	Duration     int64   `json:"duration"`
	OffsetXPerMs float64 `json:"offsetXPerMs"`
	OffsetYPerMs float64 `json:"offsetYPerMs"`
}
