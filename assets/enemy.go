package assets

import "retro-carnage/engine/geometry"

// Enemy is the serialized definition of an enemy.
type Enemy struct {
	Actions       []EnemyAction      `json:"actions"`
	Direction     string             `json:"direction"`
	Movements     []EnemyMovement    `json:"movements"`
	Position      geometry.Rectangle `json:"position"`
	Skin          string             `json:"skin"`
	SpawnCapacity int                `json:"spawnCapacity"`
	SpawnDelays   []int64            `json:"spawnDelays"`
	Type          int                `json:"type"`
}
