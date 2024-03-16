package characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerCtrlStartNewGame(t *testing.T) {
	var playerController = &PlayerCtrl{}
	playerController.StartNewGame(1)
	assert.Equal(t, 1, len(playerController.ConfiguredPlayers()))
	assert.Equal(t, 1, len(playerController.RemainingPlayers()))

	playerController.StartNewGame(2)
	assert.Equal(t, 2, len(playerController.ConfiguredPlayers()))
	assert.Equal(t, 2, len(playerController.RemainingPlayers()))
}

func TestPlayerCtrlKillPlayer(t *testing.T) {
	var playerController = &PlayerCtrl{}
	playerController.StartNewGame(2)

	for {
		playerController.KillPlayer(Players[0])
		if Players[0].lives == 0 {
			break
		}
	}

	assert.Equal(t, 2, len(playerController.ConfiguredPlayers()))
	assert.Equal(t, 1, len(playerController.RemainingPlayers()))
	assert.Equal(t, 1, playerController.RemainingPlayers()[0].index)
}
