package highscore

import "fmt"

type entry struct {
	name  string `json:"name"`
	score int    `json:"score"`
}

func (e entry) ToString(index int) string {
	return fmt.Sprintf("%2d     %7d     %s", index, e.score, e.name)
}
