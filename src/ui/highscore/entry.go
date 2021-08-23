package highscore

import "fmt"

type Entry struct {
	Name  string `json:"Name"`
	Score int    `json:"Score"`
}

func (e Entry) ToString(index int) string {
	return fmt.Sprintf("%2d     %7d     %s", index, e.Score, e.Name)
}
