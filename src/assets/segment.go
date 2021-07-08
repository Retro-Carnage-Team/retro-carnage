package assets

import "retro-carnage/engine/geometry"

// Segment is a part of the Mission that follows one specific direction.
type Segment struct {
	Backgrounds []string            `json:"backgrounds"`
	Direction   string              `json:"direction"`
	Enemies     []Enemy             `json:"enemies"`
	Goal        *geometry.Rectangle `json:"goal"`
	Obstacles   []Obstacle          `json:"obstacles"`
}
