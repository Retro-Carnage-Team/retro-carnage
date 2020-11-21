package input

import (
	"errors"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage.net/logging"
)

type Controller interface {
	AssignControllersToPlayers()
	GetControllerDeviceState(playerIdx int) (*DeviceState, error)
	GetControllerUiEventState(playerIdx int) (*UiEventState, error)
	HasTwoOrMoreDevices() bool
}

func NewController(window *pixelgl.Window) Controller {
	var result = &controllerImplementation{window: window}
	result.inputSources = make([]source, 0)
	result.lastInputStates = make([]*DeviceState, 0)
	result.rapidFireStates = make([]*rapidFireState, 0)
	return result
}

var joysticks = []pixelgl.Joystick{pixelgl.Joystick1, pixelgl.Joystick2, pixelgl.Joystick3, pixelgl.Joystick4,
	pixelgl.Joystick5, pixelgl.Joystick6, pixelgl.Joystick7, pixelgl.Joystick8, pixelgl.Joystick9, pixelgl.Joystick10,
	pixelgl.Joystick11, pixelgl.Joystick12, pixelgl.Joystick13, pixelgl.Joystick14, pixelgl.Joystick15,
	pixelgl.Joystick16}

const rapidFireOffset = 300
const rapidFireThreshold = 750

type source interface {
	State() *DeviceState
	Name() string
}

type controllerImplementation struct {
	inputSources    []source
	lastInputStates []*DeviceState
	rapidFireStates []*rapidFireState
	window          *pixelgl.Window
}

func (c *controllerImplementation) HasTwoOrMoreDevices() bool {
	for _, j := range joysticks {
		if c.window.JoystickPresent(j) {
			return true
		}
	}
	return false
}

func (c *controllerImplementation) AssignControllersToPlayers() {
	for _, j := range joysticks {
		if c.window.JoystickPresent(j) && (2 > len(c.inputSources)) {
			c.inputSources = append(c.inputSources, &gamepad{joystick: j, window: c.window})
			c.lastInputStates = append(c.lastInputStates, nil)
			c.rapidFireStates = append(c.rapidFireStates, nil)
		}
	}

	if 2 > len(c.inputSources) {
		c.inputSources = append(c.inputSources, &keyboard{Window: c.window})
		c.lastInputStates = append(c.lastInputStates, nil)
		c.rapidFireStates = append(c.rapidFireStates, nil)
	}
}

func (c *controllerImplementation) GetControllerDeviceState(playerIdx int) (*DeviceState, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf("Invalid player index: %d", playerIdx)
		return nil, errors.New("invalid argument: no such player")
	}
	return c.inputSources[playerIdx].State(), nil
}

// GetControllerUiEventState returns a UiEventState struct holding UI events. Especially the first call can returns nil
// without being in error state. Callers thus should check the result pointer before accessing it.
func (c *controllerImplementation) GetControllerUiEventState(playerIdx int) (*UiEventState, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf("Invalid player index: %d", playerIdx)
		return nil, errors.New("invalid argument: no such player")
	}

	var newState, err = c.GetControllerDeviceState(playerIdx)
	if nil != err {
		return nil, err
	}

	var result *UiEventState = nil
	if nil == c.lastInputStates[playerIdx] || nil == c.rapidFireStates[playerIdx] {
		c.lastInputStates[playerIdx] = newState
		c.rapidFireStates[playerIdx] = &rapidFireState{}
	} else {
		var oldState = c.lastInputStates[playerIdx]
		var horizontal = newState.MoveLeft || newState.MoveRight
		var vertical = newState.MoveUp || newState.MoveDown
		result = &UiEventState{
			MovedUp:       !oldState.MoveUp && newState.MoveUp && !horizontal,
			MovedDown:     !oldState.MoveDown && newState.MoveDown && !horizontal,
			MovedLeft:     !oldState.MoveLeft && newState.MoveLeft && !vertical,
			MovedRight:    !oldState.MoveRight && newState.MoveRight && !vertical,
			PressedButton: c.rapidFireStates[playerIdx].update(newState),
		}
		c.lastInputStates[playerIdx] = newState
	}
	return result, nil

}
