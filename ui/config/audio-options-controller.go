package config

import (
	"retro-carnage/config"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

const (
	optionAudioPlayEffects int = iota
	optionAudioPlayMusic
	optionAudioSave
	optionAudioBack
)

var (
	optionAudioFocusChanges = []focusChange{
		{movedDown: true, currentSelection: []int{optionAudioPlayEffects}, nextSelection: optionAudioPlayMusic},
		{movedUp: true, currentSelection: []int{optionAudioPlayMusic}, nextSelection: optionAudioPlayEffects},
		{movedDown: true, currentSelection: []int{optionAudioPlayMusic}, nextSelection: optionAudioSave},
		{movedUp: true, currentSelection: []int{optionAudioSave}, nextSelection: optionAudioPlayMusic},
		{movedRight: true, currentSelection: []int{optionAudioSave}, nextSelection: optionAudioBack},
		{movedLeft: true, currentSelection: []int{optionAudioBack}, nextSelection: optionAudioSave},
	}
)

type audioOptionsController struct {
	inputController      input.InputController
	model                *audioOptionsModel
	screenChangeRequired common.ScreenChangeCallback
}

func newAudioOptionsController(model *audioOptionsModel) *audioOptionsController {
	var result = audioOptionsController{
		model: model,
	}
	return &result
}

func (aoc *audioOptionsController) setInputController(controller input.InputController) {
	aoc.inputController = controller
}

func (aoc *audioOptionsController) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	aoc.screenChangeRequired = callback
}

func (aoc *audioOptionsController) update() {
	aoc.processUserInput()
}

func (aoc *audioOptionsController) processUserInput() {
	var uiEventState = aoc.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

focusHandling:
	for _, fc := range optionAudioFocusChanges {
		if fc.movedLeft == uiEventState.MovedLeft &&
			fc.movedRight == uiEventState.MovedRight &&
			fc.movedDown == uiEventState.MovedDown &&
			fc.movedUp == uiEventState.MovedUp {
			for _, i := range fc.currentSelection {
				if i == aoc.model.selectedOption {
					aoc.model.selectedOption = fc.nextSelection
					break focusHandling
				}
			}
		}
	}

	if uiEventState.PressedButton {
		aoc.processOptionSelected()
	}
}

func (aoc *audioOptionsController) processOptionSelected() {
	switch aoc.model.selectedOption {
	case optionAudioPlayEffects:
		aoc.model.audioConfig.PlayEffects = !aoc.model.audioConfig.PlayEffects
	case optionAudioPlayMusic:
		aoc.model.audioConfig.PlayMusic = !aoc.model.audioConfig.PlayMusic
	case optionAudioSave:
		err := config.GetConfigService().SaveAudioConfiguration(aoc.model.audioConfig)
		if nil != err {
			logging.Warning.Printf("failed to save audio settings: %s", err)
		}
		aoc.screenChangeRequired(common.ConfigurationOptions)
	case optionAudioBack:
		aoc.screenChangeRequired(common.ConfigurationOptions)
	default:
		logging.Error.Fatal("Unexpected selection in audioOptionsController")
	}
}
