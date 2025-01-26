package config

import (
	"retro-carnage/engine/characters"
	"retro-carnage/input"
	"retro-carnage/ui/common"
)

const (
	optionOnePlayer int = iota + 1
	optionTwoPlayers
	optionOptions
)

type selectController struct {
	inputController      input.InputController
	model                *selectModel
	screenChangeRequired common.ScreenChangeCallback
}

func newSelectController(model *selectModel) *selectController {
	var result = selectController{
		model: model,
	}
	return &result
}

func (sc *selectController) setInputController(controller input.InputController) {
	sc.inputController = controller
}

func (sc *selectController) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	sc.screenChangeRequired = callback
}

func (sc *selectController) update() {
	sc.processUserInput()
}

func (sc *selectController) processUserInput() {
	var uiEventState = sc.inputController.GetUiEventStateCombined()
	if nil != uiEventState {
		if uiEventState.PressedButton {
			sc.processOptionSelected()
		} else if uiEventState.MovedUp {
			if sc.model.selectedOption == optionOptions && sc.model.multiplayerPossible {
				sc.model.selectedOption = optionTwoPlayers
			} else {
				sc.model.selectedOption = optionOnePlayer
			}
		} else if uiEventState.MovedDown {
			if sc.model.selectedOption == optionOnePlayer && sc.model.multiplayerPossible {
				sc.model.selectedOption = optionTwoPlayers
			} else {
				sc.model.selectedOption = optionOptions
			}
		}
	}
}

func (sc *selectController) processOptionSelected() {
	if sc.model.selectedOption == optionOnePlayer || sc.model.selectedOption == optionTwoPlayers {
		sc.inputController.AssignInputDevicesToPlayers()
		characters.PlayerController.StartNewGame(sc.model.selectedOption)
		sc.screenChangeRequired(common.ConfigurationResult)
	} else {
		sc.screenChangeRequired(common.ConfigurationOptions)
	}
}
