package input

import (
	"github.com/faiface/pixel/pixelgl"
)

type InputController interface {
	AssignControllersToPlayers()
	ControllerDeviceState(playerIdx int) (*DeviceState, error)
	ControllerName(playerIdx int) (string, error)
	ControllerUiEventState(playerIdx int) (*UiEventState, error)
	ControllerUiEventStateCombined() *UiEventState
	HasTwoOrMoreDevices() bool
}

func NewController(window *pixelgl.Window) InputController {
	var result = &inputControllerImplementation{window: window}
	result.inputSources = make([]source, 0)
	result.lastInputStates = make([]*DeviceState, 0)
	result.rapidFireStates = make([]*rapidFireState, 0)
	return result
}
