package assets

import "retro-carnage/engine/geometry"

// Enemy is the serialized definition of an enemy.
type Enemy struct {
	ActivationDistance float64            `json:"activationDistance"`
	Movements          []EnemyMovement    `json:"movements"`
	Direction          string             `json:"direction"`
	Position           geometry.Rectangle `json:"position"`
	Skin               string             `json:"skin"`
	Type               int                `json:"type"`
}
