package assets

import "retro-carnage/engine/geometry"

// Obstacle is something that blocks the movement of Players. Some obstacles block bullets, too.
type Obstacle struct {
	geometry.Rectangle
	StopsBullets    bool `json:"stopsBullets"`
	StopsExplosives bool `json:"stopsExplosives"`
}
