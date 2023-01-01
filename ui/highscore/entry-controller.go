package highscore

import (
	"fmt"
	"retro-carnage/engine/characters"
	"sort"
)

type EntryController struct {
	entries     []Entry
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
		entries: []Entry{
			{Name: "Jonny", Score: 25_000},
			{Name: "Barney", Score: 22_000},
			{Name: "Christmas", Score: 18_500},
			{Name: "Trench", Score: 16_000},
			{Name: "Toll Road", Score: 14_500},
			{Name: "Caesar", Score: 10_000},
			{Name: "Gunner", Score: 8_000},
			{Name: "Doc", Score: 5_000},
			{Name: "Galgo", Score: 3_000},
			{Name: "Drummer", Score: 2_500},
		},
		playerNames: []string{"", ""},
	}
}

// AddEntry adds a new high Score Entry to the high Score table.
// Adding an Entry with a Score that is too low will have no effect.
func (ec *EntryController) AddEntry(newEntry Entry) {
	var newScoreTable = append(ec.entries, newEntry)
	sort.SliceStable(newScoreTable, func(i, j int) bool {
		return newScoreTable[i].Score > newScoreTable[j].Score
	})
	ec.entries = newScoreTable[:len(newScoreTable)-1]
}

// PlayerName returns the Name of the player as he has entered it for the last high Score Entry.
// Default is an empty string.
func (ec *EntryController) PlayerName(playerIndex int) (string, error) {
	if playerIndex < 0 || playerIndex >= len(ec.playerNames) {
		return "", fmt.Errorf("there is no player with index %d", playerIndex)
	}

	return ec.playerNames[playerIndex], nil
}

// SetPlayerName updates the Name of the player for the next high Score Entry.
func (ec *EntryController) SetPlayerName(playerIndex int, name string) error {
	if playerIndex < 0 || playerIndex >= len(ec.playerNames) {
		return fmt.Errorf("there is no player with index %d", playerIndex)
	}

	ec.playerNames[playerIndex] = name
	return nil
}

// ReachedHighScore returns a boolean value for each player - indicating whether or not the player reached a position in
// the high Score table.
func (ec *EntryController) ReachedHighScore() (bool, bool) {
	var p1Score = characters.PlayerController.ConfiguredPlayers()[0].Score()
	var p2Score = 0
	if 2 == characters.PlayerController.NumberOfPlayers() {
		p2Score = characters.PlayerController.ConfiguredPlayers()[1].Score()
	}

	return ec.reachedHighScore(p1Score, p2Score)
}

// reachedHighScore returns a boolean value for each player - indicating whether or not the player reached a position in
// the high Score table.
func (ec *EntryController) reachedHighScore(scorePlayer1, scorePlayer2 int) (bool, bool) {
	var newScoreTable = make([]Entry, len(ec.entries))
	copy(newScoreTable, ec.entries)

	newScoreTable = append(newScoreTable, Entry{Name: player1Placeholder, Score: scorePlayer1})
	newScoreTable = append(newScoreTable, Entry{Name: player2Placeholder, Score: scorePlayer2})
	sort.SliceStable(newScoreTable, func(i, j int) bool {
		return newScoreTable[i].Score > newScoreTable[j].Score
	})
	newScoreTable = newScoreTable[:len(newScoreTable)-2]

	var resultP1 = false
	var resultP2 = false
	for _, e := range newScoreTable {
		resultP1 = resultP1 || (e.Name == player1Placeholder)
		resultP2 = resultP2 || (e.Name == player2Placeholder)
	}

	return resultP1, resultP2
}
