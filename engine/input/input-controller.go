package input

import (
	"retro-carnage/config"

	"github.com/faiface/pixel/pixelgl"
)

type InputController interface {
	AssignInputDevicesToPlayers()
	GetInputDeviceState(playerIdx int) (*InputDeviceState, error)
	GetInputDeviceName(playerIdx int) (string, error)
	GetUiEventState(playerIdx int) (*UiEventState, error)
	GetUiEventStateCombined() *UiEventState
	GetInputDeviceInfos() []InputDeviceInfo
	GetInputDeviceConfigurations() []config.InputDeviceConfiguration
}

func NewController(window *pixelgl.Window) InputController {
	var result = &inputControllerImplementation{window: window}
	result.inputSources = make([]inputDevice, 0)
	result.lastInputStates = make([]*InputDeviceState, 0)
	result.rapidFireStates = make([]*rapidFireState, 0)
	return result
}
