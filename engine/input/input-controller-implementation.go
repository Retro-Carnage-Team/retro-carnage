package input

import (
	"errors"
	"fmt"
	"path"
	"retro-carnage/logging"

	"encoding/json"
	"os"

	"github.com/faiface/pixel/pixelgl"
)

var joysticks = []pixelgl.Joystick{pixelgl.Joystick1, pixelgl.Joystick2, pixelgl.Joystick3, pixelgl.Joystick4,
	pixelgl.Joystick5, pixelgl.Joystick6, pixelgl.Joystick7, pixelgl.Joystick8, pixelgl.Joystick9, pixelgl.Joystick10,
	pixelgl.Joystick11, pixelgl.Joystick12, pixelgl.Joystick13, pixelgl.Joystick14, pixelgl.Joystick15,
	pixelgl.Joystick16}

const (
	error_invalid_player   = "invalid argument: no such player"
	log_msg_invalid_player = "Invalid player index: %d"
	rapidFireOffset        = 300
	rapidFireThreshold     = 750
)

type source interface {
	State() *DeviceState
	Name() string
}

type inputControllerImplementation struct {
	deviceStateCombined *DeviceState
	inputSources        []source
	lastInputStates     []*DeviceState
	rapidFireStates     []*rapidFireState
	window              *pixelgl.Window
}

func (c *inputControllerImplementation) AssignControllersToPlayers() {
	for _, j := range joysticks {
		if c.window.JoystickPresent(j) && (2 > len(c.inputSources)) {
			c.inputSources = append(c.inputSources, &gamepad{joystick: j, window: c.window})
			c.lastInputStates = append(c.lastInputStates, nil)
			c.rapidFireStates = append(c.rapidFireStates, nil)
		}
	}

	if len(c.inputSources) < 2 {
		c.inputSources = append(c.inputSources, &keyboard{Window: c.window})
		c.lastInputStates = append(c.lastInputStates, nil)
		c.rapidFireStates = append(c.rapidFireStates, nil)
	}
}

func (c *inputControllerImplementation) ControllerName(playerIdx int) (string, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf(log_msg_invalid_player, playerIdx)
		return "", errors.New(error_invalid_player)
	}
	return c.inputSources[playerIdx].Name(), nil
}

func (c *inputControllerImplementation) ControllerDeviceState(playerIdx int) (*DeviceState, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf(log_msg_invalid_player, playerIdx)
		return nil, errors.New(error_invalid_player)
	}
	return c.inputSources[playerIdx].State(), nil
}

func (c *inputControllerImplementation) getControllerDeviceStateCombined() *DeviceState {
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

// ControllerUiEventState returns a UiEventState struct holding UI events. Especially the first call can return nil
// without being in error state. Callers thus should check the result pointer before accessing it.
func (c *inputControllerImplementation) ControllerUiEventState(playerIdx int) (*UiEventState, error) {
	if (0 > playerIdx) || (playerIdx >= len(c.inputSources)) {
		logging.Error.Printf(log_msg_invalid_player, playerIdx)
		return nil, errors.New(error_invalid_player)
	}

	var newState, err = c.ControllerDeviceState(playerIdx)
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

// ControllerUiEventStateCombined returns a UiEventState struct holding UI events. Especially the first call can
// return nil without being in error state. Callers thus should check the result pointer before accessing it.
// The difference between GetControllerUiEventState and GetControllerUiEventStateCombined is that this method returns a
// struct that contains the values for all input devices. So you can use this method before the input devices are
// assigned to players.
func (c *inputControllerImplementation) ControllerUiEventStateCombined() *UiEventState {
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

// GetControllers returns a list of all controllers that are available
func (c *inputControllerImplementation) GetControllers() []ControllerInfo {
	var result = append(make([]ControllerInfo, 0), ControllerInfo{DeviceName: DeviceNameKeyboard, JoystickIndex: -1})
	for _, j := range joysticks {
		if c.window.JoystickPresent(j) {
			var joystick = ControllerInfo{DeviceName: c.window.JoystickName(j), JoystickIndex: int(j)}
			result = append(result, joystick)
		}
	}
	return result
}

func (c *inputControllerImplementation) GetControllerConfigurations() []ControllerConfiguration {
	var result = c.loadControllerConfigurations()
	// TODO: Remove invalid configurations (controller not found / not matching)
	// TODO: If len(configurations) < 2 && unconfigured controllers present: Add default configuration for gamepad
	// TODO: If len(configurations) < 2: Add default configuration for keyboard
	return result
}

// loadControllerConfigurations reads the controller configurations that is stored on disk.
// Returns empty array of not configurations can be found.
func (c *inputControllerImplementation) loadControllerConfigurations() []ControllerConfiguration {
	var result = make([]ControllerConfiguration, 0)
	for i := 0; i < 2; i++ {
		filePath, err := c.buildConfigurationFilePath(i)
		if nil != err {
			continue
		}

		logging.Trace.Printf("loading controller configuration for player %d from %s", i, filePath)
		data, err := os.ReadFile(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				logging.Info.Printf("controller config file not present %s", filePath)
			} else {
				logging.Warning.Printf("failed to read controller config %s", filePath)
			}
			continue
		}

		var config = &ControllerConfiguration{}
		err = json.Unmarshal(data, config)
		if err != nil {
			logging.Warning.Printf("failed to deserialize controller config %s", filePath)
		}
		result = append(result, *config)
	}

	return result
}

// SaveControllerConfiguration stores the given controller configuration for the specified player
func (c *inputControllerImplementation) SaveControllerConfiguration(cc ControllerConfiguration, playerIdx int) error {
	folderPath, err := c.buildConfigurationFolderPath()
	if nil != err {
		return err
	}

	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0700)
		if nil != err {
			logging.Warning.Printf("failed to create folder for configurations: %s", folderPath)
			return err
		}
	}

	filePath, err := c.buildConfigurationFilePath(playerIdx)
	if nil != err {
		logging.Warning.Printf("failed to calculate config path for controller %d", playerIdx)
		return err
	}

	logging.Trace.Printf("saving controller configuration for player %d to %s", playerIdx, filePath)
	jsonData, _ := json.Marshal(cc)
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		logging.Warning.Printf("failed to write controller config %s", filePath)
		return err
	}

	return nil
}

func (c *inputControllerImplementation) buildConfigurationFilePath(playerIdx int) (string, error) {
	var folder, err = c.buildConfigurationFolderPath()
	if err != nil {
		return "", err
	}

	return path.Join(folder, fmt.Sprintf("controller-%d.json", playerIdx)), nil
}

func (c *inputControllerImplementation) buildConfigurationFolderPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if nil != err {
		logging.Warning.Printf("failed to calculate config folder path")
		return "", err
	}

	return path.Join(homeDir, ".retro-carnage", "settings"), nil
}
