package characters

import "retro-carnage/engine/cheat"

// PlayerCtrl manages the set of players
type PlayerCtrl struct {
	numberOfPlayers int
}

var (
	PlayerController = &PlayerCtrl{}
)

// NumberOfPlayers returns the number of players for the current game mode. This can be 1 or 2.
func (pc *PlayerCtrl) NumberOfPlayers() int {
	return pc.numberOfPlayers
}

// ConfiguredPlayers returns a slice containing the configured players. The slide might contain dead players.
func (pc *PlayerCtrl) ConfiguredPlayers() []*Player {
	if pc.numberOfPlayers == 1 {
		return []*Player{playerOne}
	}
	return []*Player{playerOne, playerTwo}
}

// RemainingPlayers returns a slice containing the players that are currently alive.
func (pc *PlayerCtrl) RemainingPlayers() []*Player {
	var result = make([]*Player, 0)
	for _, player := range pc.ConfiguredPlayers() {
		if player.Alive() {
			result = append(result, player)
		}
	}
	return result
}

// KillPlayer decreases the number of lives of the specified player by one
func (pc *PlayerCtrl) KillPlayer(player *Player) {
	if !cheat.GetCheatController().IsNumberOfLivesUnlimited() {
		player.SetLives(player.Lives() - 1)
	}
}

// StartNewGame initializes the controller for a new game of numberOfPlayers players
func (pc *PlayerCtrl) StartNewGame(numberOfPlayers int) {
	pc.numberOfPlayers = numberOfPlayers
	playerOne.Reset()
	playerTwo.Reset()
}
