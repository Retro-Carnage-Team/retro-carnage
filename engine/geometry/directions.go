package geometry

import "math"

// Direction specifies one of eight possible directions (cardinal and diagonal).
type Direction struct {
	Name  string
	up    bool
	right bool
	down  bool
	left  bool
}

var (
	Up         = Direction{Name: "up", up: true, right: false, down: false, left: false}
	UpRight    = Direction{Name: "up_right", up: true, right: true, down: false, left: false}
	Right      = Direction{Name: "right", up: false, right: true, down: false, left: false}
	DownRight  = Direction{Name: "down_right", up: false, right: true, down: true, left: false}
	Down       = Direction{Name: "down", up: false, right: false, down: true, left: false}
	DownLeft   = Direction{Name: "down_left", up: false, right: false, down: true, left: true}
	Left       = Direction{Name: "left", up: false, right: false, down: false, left: true}
	UpLeft     = Direction{Name: "up_left", up: true, right: false, down: false, left: true}
	directions = [...]Direction{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}
)

// GetDirectionForCardinals returns the direction that is specified by the combination of given cardinal directions.
// If no such direction exists (e.g. if no parameter is true or if opposite cardinal directions are true) it will return
// nil.
func GetDirectionForCardinals(up bool, down bool, left bool, right bool) *Direction {
	for _, dir := range directions {
		if dir.up == up && dir.right == right && dir.down == down && dir.left == left {
			return &dir
		}
	}
	return nil
}

// GetDirectionByName returns the direction that is specified by name.
// If no such direction exists it will return nil.
func GetDirectionByName(name string) *Direction {
	for _, dir := range directions {
		if dir.Name == name {
			return &dir
		}
	}
	return nil
}

func (d Direction) IsDiagonal() bool {
	return d == UpRight || d == UpLeft || d == DownLeft || d == DownRight
}

// ToAngle returns the angle of the direction in radians
func (d Direction) ToAngle() float64 {
	switch d {
	case Right:
		// 0 * math.Pi / 4
		return 0
	case UpRight:
		// 1 * math.Pi / 4
		return math.Pi / 4
	case Up:
		// 2 * math.Pi / 4
		return math.Pi / 2
	case UpLeft:
		// 3 * math.Pi / 4
		return 3 * math.Pi / 4
	case Left:
		// 4 * math.Pi / 4
		return math.Pi
	case DownLeft:
		// 5 * math.Pi / 4
		return 5 * math.Pi / 4
	case Down:
		// 6 * math.Pi / 4
		return 3 * math.Pi / 2
	case DownRight:
		// 7 * math.Pi / 4
		return 7 * math.Pi / 4
	default:
		return 0
	}
}
