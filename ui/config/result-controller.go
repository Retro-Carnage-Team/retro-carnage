package config

import (
	"retro-carnage/engine/characters"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

const timeout = 2_500

type resultController struct {
	inputController      input.InputController
	model                *resultModel
	screenChangeRequired common.ScreenChangeCallback
	timeElapsed          int64
}

func newResultController(model *resultModel) *resultController {
	var result = resultController{
		model: model,
	}
	return &result
}

func (rs *resultController) setInputController(controller input.InputController) {
	rs.inputController = controller
}

func (rs *resultController) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	rs.screenChangeRequired = callback
}

func (rs *resultController) setUp() {
	name, err := rs.inputController.GetInputDeviceName(0)
	if nil == err {
		rs.model.infoTextPlayerOne += name
	} else {
		logging.Warning.Printf("Failed to get controller name for player 0: %v", err)
	}

	if characters.PlayerController.NumberOfPlayers() == 2 {
		name, err = rs.inputController.GetInputDeviceName(1)
		if nil == err {
			rs.model.infoTextPlayerTwo += name
		} else {
			logging.Warning.Printf("Failed to get controller name for player 1: %v", err)
		}
	}
}

func (rs *resultController) update(timeElapsedInMs int64) {
	rs.timeElapsed += timeElapsedInMs

	var uiEventState = rs.inputController.GetUiEventStateCombined()
	if nil != uiEventState && uiEventState.PressedButton {
		rs.screenChangeRequired(common.Mission)
	} else if rs.timeElapsed >= timeout {
		rs.screenChangeRequired(common.Mission)
	}
}
