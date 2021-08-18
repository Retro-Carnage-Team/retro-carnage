package highscore

import "sort"

type entryController struct {
	entries []entry
}

var (
	entryControllerInstance = newEntryController()
)

// newEntryController creates a new instance of entryController
func newEntryController() *entryController {
	return &entryController{
		entries: []entry{
			{name: "Jonny", score: 25_000},
			{name: "Barney", score: 22_000},
			{name: "Christmas", score: 18_500},
			{name: "Trench", score: 16_000},
			{name: "Toll Road", score: 14_500},
			{name: "Caesar", score: 10_000},
			{name: "Gunner", score: 8_000},
			{name: "Doc", score: 5_000},
			{name: "Galgo", score: 3_000},
			{name: "Drummer", score: 2_500},
		},
	}
}

// isHighScore returns true if the given score is high enough for a new entry in the high score table.
func (ec *entryController) isHighScore(score int) bool {
	var lastPlace = ec.entries[len(ec.entries)-1]
	return score > lastPlace.score
}

// addEntry adds a new high score entry to the high score table.
// Adding an entry with a score that is too low will have no effect.
func (ec *entryController) addEntry(newEntry entry) {
	var newScoreTable = append(ec.entries, newEntry)
	sort.SliceStable(newScoreTable, func(i, j int) bool {
		return newScoreTable[i].score > newScoreTable[j].score
	})
	ec.entries = newScoreTable[:len(newScoreTable)-1]
}
