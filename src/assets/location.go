package assets

// Location is the position of something based on the 1500x1500 pixel coordinate space of the screen area.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
