package geometry

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
