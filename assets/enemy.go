package assets

import "retro-carnage/engine/geometry"

// Enemy is the serialized definition of an enemy.
type Enemy struct {
	Movements   []EnemyMovement    `json:"movements"`
	Direction   string             `json:"direction"`
	Position    geometry.Rectangle `json:"position"`
	Skin        string             `json:"skin"`
	Type        int                `json:"type"`
	Actions     []EnemyAction      `json:"actions"`
	SpawnDelays []int64            `json:"spawnDelays"`
}
