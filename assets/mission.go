package assets

type Mission struct {
	Briefing string    `json:"briefing"`
	Client   string    `json:"client"`
	Location Location  `json:"location"`
	Music    Song      `json:"music"`
	Name     string    `json:"name"`
	Reward   int       `json:"reward"`
	Segments []Segment `json:"segments"`
}
