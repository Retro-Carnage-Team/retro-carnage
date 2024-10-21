package config

import (
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

const (
	optionInputSelectPlayer1 int = iota
	optionInputSelectPlayer2
	optionInputBack
)

var (
	optionInputFocusChanges = []focusChange{
		{movedDown: true, currentSelection: []int{optionInputSelectPlayer1}, nextSelection: optionInputSelectPlayer2},
		{movedUp: true, currentSelection: []int{optionInputSelectPlayer2}, nextSelection: optionInputSelectPlayer1},
		{movedDown: true, currentSelection: []int{optionInputSelectPlayer2}, nextSelection: optionInputBack},
		{movedUp: true, currentSelection: []int{optionInputBack}, nextSelection: optionInputSelectPlayer2},
	}
)

type inputOptionsController struct {
	inputController      input.InputController
	model                *inputOptionsModel
	screenChangeRequired common.ScreenChangeCallback
}

func newInputOptionsController(model *inputOptionsModel) *inputOptionsController {
	var result = inputOptionsController{
		model: model,
	}
	return &result
}

func (ioc *inputOptionsController) setInputController(controller input.InputController) {
	ioc.inputController = controller
}

func (ioc *inputOptionsController) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	ioc.screenChangeRequired = callback
}

func (ioc *inputOptionsController) update() {
	ioc.processUserInput()
}

func (ioc *inputOptionsController) processUserInput() {
	var uiEventState = ioc.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

focusHandling:
	for _, fc := range optionInputFocusChanges {
		if fc.movedLeft == uiEventState.MovedLeft &&
			fc.movedRight == uiEventState.MovedRight &&
			fc.movedDown == uiEventState.MovedDown &&
			fc.movedUp == uiEventState.MovedUp {
			for _, i := range fc.currentSelection {
				if i == ioc.model.selectedOption {
					ioc.model.selectedOption = fc.nextSelection
					break focusHandling
				}
			}
		}
	}

	if uiEventState.PressedButton {
		ioc.processOptionSelected()
	}
}

func (ioc *inputOptionsController) processOptionSelected() {
	switch ioc.model.selectedOption {
	case optionInputSelectPlayer1:
		ioc.screenChangeRequired(common.ConfigurationControlsP1)
	case optionInputSelectPlayer2:
		ioc.screenChangeRequired(common.ConfigurationControlsP2)
	case optionInputBack:
		ioc.screenChangeRequired(common.ConfigurationOptions)
	default:
		logging.Error.Fatal("Unexpected selection in inputOptionsController")
	}
}
