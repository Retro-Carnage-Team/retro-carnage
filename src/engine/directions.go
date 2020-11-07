package engine

type Direction struct {
	name  string
	up    bool
	right bool
	down  bool
	left  bool
}

var (
	Up         = Direction{name: "up", up: true, right: false, down: false, left: false}
	UpRight    = Direction{name: "up_right", up: true, right: true, down: false, left: false}
	Right      = Direction{name: "right", up: false, right: true, down: false, left: false}
	DownRight  = Direction{name: "down_right", up: false, right: true, down: true, left: false}
	Down       = Direction{name: "down", up: false, right: false, down: true, left: false}
	DownLeft   = Direction{name: "down_left", up: false, right: false, down: true, left: true}
	Left       = Direction{name: "left", up: false, right: false, down: false, left: true}
	UpLeft     = Direction{name: "up_left", up: true, right: false, down: false, left: true}
	directions = [...]Direction{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}
)

func GetDirectionForCardinals(up bool, down bool, left bool, right bool) *Direction {
	for _, dir := range directions {
		if dir.up == up && dir.right == right && dir.down == down && dir.left == left {
			return &dir
		}
	}
	return nil
}
