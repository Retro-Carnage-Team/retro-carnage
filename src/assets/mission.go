package assets

import (
	"retro-carnage/engine/geometry"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type EnemyMovement struct {
	Duration     float64 `json:"duration"`
	OffsetXPerMs float64 `json:"offsetXPerMs"`
	OffsetYPerMs float64 `json:"offsetYPerMs"`
}

type Enemy struct {
	ActivationDistance float64            `json:"activationDistance"`
	Movements          []EnemyMovement    `json:"movements"`
	Direction          string             `json:"direction"`
	Position           geometry.Rectangle `json:"position"`
	Skin               string             `json:"skin"`
	Type               int                `json:"type"`
}

type Segment struct {
	Backgrounds []string             `json:"backgrounds"`
	Direction   string               `json:"direction"`
	Goal        *geometry.Rectangle  `json:"goal"`
	Obstacles   []geometry.Rectangle `json:"obstacles"`
	enemies     []Enemy              `json:"enemies"`
}

type Mission struct {
	Briefing   string    `json:"briefing"`
	Client     string    `json:"client"`
	Location   Location  `json:"location"`
	Music      Song      `json:"music"`
	Name       string    `json:"name"`
	Reward     int       `json:"reward"`
	Segments   []Segment `json:"segments"`
	Unfinished bool      `json:"unfinished"`
}
