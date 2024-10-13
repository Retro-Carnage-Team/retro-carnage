package enter_name

import (
	"retro-carnage/engine/characters"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/highscore"
)

const (
	cursorBlinkDuration = 250
	maxLengthOfName     = 10
)

type controller struct {
	duration             int64
	inputController      input.InputController
	model                *model
	screenChangeRequired common.ScreenChangeCallback
}

func newController(model *model) *controller {
	var name, error = highscore.EntryControllerInstance.PlayerName(model.playerIdx)
	if nil != error {
		logging.Error.Printf("Failed to get player name: %v", error)
		model.playerName = ""
	} else {
		model.playerName = name
	}

	var result = controller{
		model: model,
	}
	return &result
}

// SetInputController passes the input controller to the screen.
func (c *controller) setInputController(inputCtrl input.InputController) {
	c.inputController = inputCtrl
}

// SetScreenChangeCallback passes a callback function that cann be called to switch to another screen.
func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) update(elapsedTimeInMs int64, typed string, enter bool, backspace bool) {
	c.duration += elapsedTimeInMs
	if c.duration > cursorBlinkDuration {
		c.model.cursorVisible = !c.model.cursorVisible
		c.duration = 0
	}

	c.model.playerName = c.model.playerName + typed
	if maxLengthOfName < len(c.model.playerName) {
		c.model.playerName = c.model.playerName[:maxLengthOfName]
	}

	c.model.playerNameDisplay = c.model.playerName
	var uiEventStateCombined = c.inputController.GetUiEventStateCombined()
	if enter || ((nil != uiEventStateCombined) && uiEventStateCombined.PressedButton) {
		highscore.EntryControllerInstance.SetPlayerName(c.model.playerIdx, c.model.playerName)
		highscore.EntryControllerInstance.AddEntry(highscore.Entry{
			Name:  c.model.playerName,
			Score: characters.PlayerController.ConfiguredPlayers()[c.model.playerIdx].Score(),
		})
		c.exit()
	} else {
		if backspace && (0 < len(c.model.playerName)) {
			c.model.playerName = c.model.playerName[:len(c.model.playerName)-1]
			c.model.playerNameDisplay = c.model.playerName
		}

		if c.model.cursorVisible {
			c.model.playerNameDisplay = c.model.playerName + "|"
		} else {
			c.model.playerNameDisplay = c.model.playerName + " "
		}
	}
}

func (c *controller) exit() {
	if c.model.playerIdx == 0 {
		var _, p2 = highscore.EntryControllerInstance.ReachedHighScore()
		if p2 {
			c.screenChangeRequired(common.EnterNameP2)
		} else {
			c.screenChangeRequired(common.HighScore)
		}
	} else {
		c.screenChangeRequired(common.HighScore)
	}
}
