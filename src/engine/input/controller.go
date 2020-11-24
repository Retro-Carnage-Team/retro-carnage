package input

import (
	"errors"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/logging"
)

type Controller interface {
	AssignControllersToPlayers()
	GetControllerDeviceState(playerIdx int) (*DeviceState, error)
	GetControllerName(playerIdx int) (string, error)
	GetControllerUiEventState(playerIdx int) (*UiEventState, error)
	GetControllerUiEventStateCombined() *UiEventState
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
	deviceStateCombined *DeviceState
	inputSources        []source
	lastInputStates     []*DeviceState
	rapidFireStates     []*rapidFireState
	window              *pixelgl.Window
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

func (c *controllerImplementation) GetControllerName(playerIdx int) (string, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf("Invalid player index: %d", playerIdx)
		return "", errors.New("invalid argument: no such player")
	}
	return c.inputSources[playerIdx].Name(), nil
}

func (c *controllerImplementation) GetControllerDeviceState(playerIdx int) (*DeviceState, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf("Invalid player index: %d", playerIdx)
		return nil, errors.New("invalid argument: no such player")
	}
	return c.inputSources[playerIdx].State(), nil
}

func (c *controllerImplementation) getControllerDeviceStateCombined() *DeviceState {
	var result *DeviceState = nil
	var padCount = 0
	for _, j := range joysticks {
		if c.window.JoystickPresent(j) && (2 > padCount) {
			padCount++
			var gamepad = &gamepad{joystick: j, window: c.window}
			var state = gamepad.State()
			if nil == result {
				result = state
			} else {
				result = result.Combine(state)
			}
		}
	}

	var keyboard = &keyboard{Window: c.window}
	var state = keyboard.State()
	if nil == result {
		result = state
	} else {
		result = result.Combine(state)
	}
	return result
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

// GetControllerUiEventStateCombined returns a UiEventState struct holding UI events. Especially the first call can
// returns nil without being in error state. Callers thus should check the result pointer before accessing it.
// The difference between GetControllerUiEventState and GetControllerUiEventStateCombined is that this method returns a
// struct that contains the values for all input devices. So you can use this method before the input devices are
// assigned to players.
func (c *controllerImplementation) GetControllerUiEventStateCombined() *UiEventState {
	var newState = c.getControllerDeviceStateCombined()
	var result *UiEventState = nil
	if nil == c.deviceStateCombined {
		c.deviceStateCombined = newState
	} else {
		var oldState = c.deviceStateCombined
		var horizontal = newState.MoveLeft || newState.MoveRight
		var vertical = newState.MoveUp || newState.MoveDown
		result = &UiEventState{
			MovedUp:       !oldState.MoveUp && newState.MoveUp && !horizontal,
			MovedDown:     !oldState.MoveDown && newState.MoveDown && !horizontal,
			MovedLeft:     !oldState.MoveLeft && newState.MoveLeft && !vertical,
			MovedRight:    !oldState.MoveRight && newState.MoveRight && !vertical,
			PressedButton: !oldState.IsButtonPressed() && newState.IsButtonPressed(),
		}
		c.deviceStateCombined = newState
	}
	return result
}
