package highscore

import (
	"fmt"
	"retro-carnage/engine/characters"
	"sort"
)

type EntryController struct {
	entries     []entry
	playerNames []string
}

var (
	EntryControllerInstance = newEntryController()
)

const (
	player1Placeholder = "___p1___"
	player2Placeholder = "___p2___"
)

// newEntryController creates a new instance of entryController
func newEntryController() *EntryController {
	return &EntryController{
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
		playerNames: []string{"", ""},
	}
}

// isHighScore returns true if the given score is high enough for a new entry in the high score table.
func (ec *EntryController) isHighScore(score int) bool {
	var lastPlace = ec.entries[len(ec.entries)-1]
	return score > lastPlace.score
}

// addEntry adds a new high score entry to the high score table.
// Adding an entry with a score that is too low will have no effect.
func (ec *EntryController) addEntry(newEntry entry) {
	var newScoreTable = append(ec.entries, newEntry)
	sort.SliceStable(newScoreTable, func(i, j int) bool {
		return newScoreTable[i].score > newScoreTable[j].score
	})
	ec.entries = newScoreTable[:len(newScoreTable)-1]
}

// PlayerName returns the name of the player as he has entered it for the last high score entry.
// Default is an empty string.
func (ec *EntryController) PlayerName(playerIndex int) (string, error) {
	if playerIndex < 0 || playerIndex >= len(ec.playerNames) {
		return "", fmt.Errorf("there is no player with index %d", playerIndex)
	}

	return ec.playerNames[playerIndex], nil
}

// SetPlayerName updates the name of the player for the next high score entry.
func (ec *EntryController) SetPlayerName(playerIndex int, name string) error {
	if playerIndex < 0 || playerIndex >= len(ec.playerNames) {
		return fmt.Errorf("there is no player with index %d", playerIndex)
	}

	ec.playerNames[playerIndex] = name
	return nil
}

// ReachedHighScore returns a boolean value for each player - indicating whether or not the player reached a position in
// the high score table.
func (ec *EntryController) ReachedHighScore() (bool, bool) {
	var p1Score = characters.PlayerController.ConfiguredPlayers()[0].Score()
	var p2Score = 0
	if 2 == characters.PlayerController.NumberOfPlayers() {
		p2Score = characters.PlayerController.ConfiguredPlayers()[1].Score()
	}

	return ec.reachedHighScore(p1Score, p2Score)
}

// reachedHighScore returns a boolean value for each player - indicating whether or not the player reached a position in
// the high score table.
func (ec *EntryController) reachedHighScore(scorePlayer1, scorePlayer2 int) (bool, bool) {
	newScoreTable := append(ec.entries, entry{name: player1Placeholder, score: scorePlayer1})
	newScoreTable = append(newScoreTable, entry{name: player2Placeholder, score: scorePlayer2})
	sort.SliceStable(newScoreTable, func(i, j int) bool {
		return newScoreTable[i].score > newScoreTable[j].score
	})
	newScoreTable = newScoreTable[:len(newScoreTable)-2]

	var resultP1 = false
	var resultP2 = false
	for _, e := range newScoreTable {
		resultP1 = resultP1 || (e.name == player1Placeholder)
		resultP2 = resultP2 || (e.name == player2Placeholder)
	}

	return resultP1, resultP2
}
