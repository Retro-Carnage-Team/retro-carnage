package highscore

import "fmt"

type Entry struct {
	Name  string `json:"Name"`
	Score int    `json:"Score"`
}

func (e Entry) ToString(index int) string {
	return fmt.Sprintf("%02d     %07d     %s", index, e.Score, e.Name)
}
