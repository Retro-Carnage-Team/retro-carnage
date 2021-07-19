package assets

const (
	EnemyActionBullet  = "bullet"
	EnemyActionGrenade = "grenade"
)

// EnemyAction is something that an Enemy does. Each Enemy has a list of these actions. Elements of this list of
// EnemyActions are repeated in an endless loop - until the enemy leaves the screen or dies.
type EnemyAction struct {
	Action string `json:"action"`
	Delay  int64  `json:"delay"`
}
