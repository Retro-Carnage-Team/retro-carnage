package input

import (
	"errors"
	"retro-carnage/config"
	"retro-carnage/logging"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

var joysticks = []pixel.Joystick{pixel.Joystick1, pixel.Joystick2, pixel.Joystick3, pixel.Joystick4,
	pixel.Joystick5, pixel.Joystick6, pixel.Joystick7, pixel.Joystick8, pixel.Joystick9, pixel.Joystick10,
	pixel.Joystick11, pixel.Joystick12, pixel.Joystick13, pixel.Joystick14, pixel.Joystick15,
	pixel.Joystick16}

const (
	error_invalid_player   = "invalid argument: no such player"
	log_msg_invalid_player = "Invalid player index: %d"
	rapidFireOffset        = 300
	rapidFireThreshold     = 750
)

type inputDeviceWithState struct {
	inputSource    inputDevice
	lastInputState *InputDeviceState
	rapidFireState *rapidFireState
}

type InputController struct {
	deviceConfigurations []config.InputDeviceConfiguration
	deviceStateCombined  *InputDeviceState
	inputSources         []inputDeviceWithState
	window               *opengl.Window
}

func NewController(window *opengl.Window) InputController {
	var result = InputController{window: window}
	result.inputSources = make([]inputDeviceWithState, 0)
	return result
}

func (c *InputController) AssignInputDevicesToPlayers() {
	if len(c.inputSources) > 0 {
		c.inputSources = make([]inputDeviceWithState, 0)
	}

	var deviceConfigurations = c.GetInputDeviceConfigurations()
	for _, cfg := range deviceConfigurations {
		var newRecord = inputDeviceWithState{
			inputSource: c.buildInputDevice(cfg),
		}
		c.inputSources = append(c.inputSources, newRecord)
	}
}

func (c *InputController) GetInputDeviceName(playerIdx int) (string, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf(log_msg_invalid_player, playerIdx)
		return "", errors.New(error_invalid_player)
	}
	return c.inputSources[playerIdx].inputSource.Name(), nil
}

func (c *InputController) GetInputDeviceState(playerIdx int) (*InputDeviceState, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf(log_msg_invalid_player, playerIdx)
		return nil, errors.New(error_invalid_player)
	}
	return c.inputSources[playerIdx].inputSource.State(), nil
}

func (c *InputController) getControllerDeviceStateCombined() *InputDeviceState {
	if nil == c.deviceConfigurations {
		c.deviceConfigurations = c.GetInputDeviceConfigurations()
	}

	var result *InputDeviceState = nil
	for _, j := range joysticks {
		if !c.window.JoystickPresent(j) {
			continue
		}

		var device inputDevice = nil
		for _, cfg := range c.deviceConfigurations {
			if cfg.GamepadConfiguration.JoystickIndex == int(j) && cfg.DeviceName == c.window.JoystickName(j) {
				device = c.buildInputDevice(cfg)
			}
		}

		if device == nil {
			device = &gamepad{
				configuration: config.NewGamepadConfiguration(*c.window, j),
				window:        c.window,
			}
		}

		result = device.State().Combine(result)
	}

	var keyboard = &keyboard{window: c.window}
	return keyboard.State().Combine(result)
}

// GetUiEventState returns a UiEventState struct holding UI events. Especially the first call can return nil
// without being in error state. Callers thus should check the result pointer before accessing it.
func (c *InputController) GetUiEventState(playerIdx int) (*UiEventState, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf(log_msg_invalid_player, playerIdx)
		return nil, errors.New(error_invalid_player)
	}

	var newState, err = c.GetInputDeviceState(playerIdx)
	if nil != err {
		return nil, err
	}

	var result *UiEventState = nil
	if nil == c.inputSources[playerIdx].lastInputState || nil == c.inputSources[playerIdx].rapidFireState {
		c.inputSources[playerIdx].lastInputState = newState
		c.inputSources[playerIdx].rapidFireState = &rapidFireState{}
	} else {
		var oldState = c.inputSources[playerIdx].lastInputState
		var horizontal = newState.MoveLeft || newState.MoveRight
		var vertical = newState.MoveUp || newState.MoveDown
		result = &UiEventState{
			MovedUp:       !oldState.MoveUp && newState.MoveUp && !horizontal,
			MovedDown:     !oldState.MoveDown && newState.MoveDown && !horizontal,
			MovedLeft:     !oldState.MoveLeft && newState.MoveLeft && !vertical,
			MovedRight:    !oldState.MoveRight && newState.MoveRight && !vertical,
			PressedButton: c.inputSources[playerIdx].rapidFireState.update(newState),
		}
		c.inputSources[playerIdx].lastInputState = newState
	}
	return result, nil
}

// GetUiEventStateCombined returns a UiEventState struct holding UI events. Especially the first call can
// return nil without being in error state. Callers thus should check the result pointer before accessing it.
// The difference between GetControllerUiEventState and GetControllerUiEventStateCombined is that this method returns a
// struct that contains the values for all input devices. So you can use this method before the input devices are
// assigned to players.
func (c *InputController) GetUiEventStateCombined() *UiEventState {
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

// GetInputDeviceInfos returns a list of all controllers that are available
func (c *InputController) GetInputDeviceInfos() []InputDeviceInfo {
	var result = []InputDeviceInfo{
		{
			DeviceName:    config.DeviceNameKeyboard,
			JoystickIndex: -1,
		},
	}
	for _, j := range joysticks {
		if c.window.JoystickPresent(j) {
			var joystick = InputDeviceInfo{DeviceName: c.window.JoystickName(j), JoystickIndex: int(j)}
			result = append(result, joystick)
		}
	}
	return result
}

func (c *InputController) GetInputDeviceConfigurations() []config.InputDeviceConfiguration {
	var cs = config.ConfigService{}
	var result = cs.LoadInputDeviceConfigurations()
	result = c.filterValidConfigurations(result)

	for _, j := range joysticks {
		var joystickPresent = c.window.JoystickPresent(j)
		var deviceConfigured = false
		for _, cc := range result {
			deviceConfigured = deviceConfigured || (int(j) == cc.GamepadConfiguration.JoystickIndex)
		}
		if len(result) < 2 && joystickPresent && !deviceConfigured {
			result = append(result, config.NewGamepadConfiguration(*c.window, j))
		}

		if len(result) == 2 {
			break
		}
	}

	// Add default configuration for keyboard if there are less then two configured controllers
	if len(result) < 2 {
		var containsKeyboard = false
		for _, cc := range result {
			containsKeyboard = containsKeyboard || cc.DeviceName == config.DeviceNameKeyboard
		}
		if !containsKeyboard {
			result = append(result, config.NewKeyboardConfiguration())
		}
	}

	return result
}

// filterValidConfigurations filters the given list of ControllerConfigurations so that it contains only controllers
// that are actually present.
func (c *InputController) filterValidConfigurations(configurations []config.InputDeviceConfiguration) []config.InputDeviceConfiguration {
	var result = make([]config.InputDeviceConfiguration, 0)
	for _, cc := range configurations {
		if (cc.DeviceName == config.DeviceNameKeyboard) ||
			(c.window.JoystickPresent(pixel.Joystick(cc.JoystickIndex)) &&
				c.window.JoystickName(pixel.Joystick(cc.JoystickIndex)) == cc.DeviceName) {
			result = append(result, cc)
		}
	}
	return result
}

func (c *InputController) buildInputDevice(cfg config.InputDeviceConfiguration) inputDevice {
	if cfg.DeviceName == config.DeviceNameKeyboard {
		return &keyboard{configuration: cfg, window: c.window}
	}
	return &gamepad{configuration: cfg, window: c.window}
}
