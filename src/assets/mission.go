package assets

import "retro-carnage/engine/geometry"

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Segment struct {
	Backgrounds []string             `json:"backgrounds"`
	Direction   string               `json:"direction"`
	Goal        *geometry.Rectangle  `json:"goal"`
	Obstacles   []geometry.Rectangle `json:"obstacles"`
	// enemies  []Enemy              `json:"enemies"`
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
