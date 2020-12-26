package characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayerCtrl_StartNewGame(t *testing.T) {
	var playerController = &PlayerCtrl{}
	playerController.StartNewGame(1)
	assert.Equal(t, 1, len(playerController.ConfiguredPlayers()))
	assert.Equal(t, 1, len(playerController.RemainingPlayers()))

	playerController.StartNewGame(2)
	assert.Equal(t, 2, len(playerController.ConfiguredPlayers()))
	assert.Equal(t, 2, len(playerController.RemainingPlayers()))
}

func TestPlayerCtrl_KillPlayer(t *testing.T) {
	var playerController = &PlayerCtrl{}
	playerController.StartNewGame(2)

	for {
		playerController.KillPlayer(0)
		if 0 == Players[0].lives {
			break
		}
	}

	assert.Equal(t, 2, len(playerController.ConfiguredPlayers()))
	assert.Equal(t, 1, len(playerController.RemainingPlayers()))
	assert.Equal(t, 1, playerController.RemainingPlayers()[0].index)
}
