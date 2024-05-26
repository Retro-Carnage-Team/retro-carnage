package config

import (
	"fmt"
	"retro-carnage/logging"
	"strings"

	"github.com/Retro-Carnage-Team/pixel/pixelgl"
)

const DeviceNameKeyboard = "Keyboard"

var digitalControllers = []string{"Competition Pro"}

// KeyboardConfiguration holds the configuration aspects that are specific to keyboards.
// This makes sure that a user can configure e.g. WASD or use arrow keys.
type KeyboardConfiguration struct {
	InputUp    int `json:"inputUp"`
	InputDown  int `json:"inputDown"`
	InputLeft  int `json:"inputLeft"`
	InputRight int `json:"inputRight"`
}

// GamepadConfiguration holds the configuration aspects that are specific to joysticks.
// Currently that is only the way values of the x and y axis are interpreted.
type GamepadConfiguration struct {
	HasDigitalAxis bool `json:"hasDigitalAxis"`
	JoystickIndex  int  `json:"joystickIndex"`
}

// InputDeviceConfiguration holds the configuration for a specific device. This combines common aspects
// with GamepadConfiguration and KeyboardConfigurations.
type InputDeviceConfiguration struct {
	GamepadConfiguration  `json:"gamepadConfig"`
	KeyboardConfiguration `json:"keyboardConfig"`
	DeviceName            string `json:"deviceName"`
	InputFire             int    `json:"inputFire"`
	InputNextWeapon       int    `json:"inputNextWeapon"`
	InputPreviousWeapon   int    `json:"inputPrevWeapon"`
}

func NewGamepadConfiguration(w pixelgl.Window, j pixelgl.Joystick) InputDeviceConfiguration {
	if !w.JoystickPresent(j) {
		logging.Error.Fatalf("NewGamepadConfiguration was called for joystick that is not present!")
	}

	var name = w.JoystickName(j)
	var digitalController = false
	for _, controllerName := range digitalControllers {
		if strings.Contains(strings.ToLower(name), strings.ToLower(controllerName)) {
			digitalController = true
			break
		}
	}

	var result = InputDeviceConfiguration{
		GamepadConfiguration: GamepadConfiguration{
			HasDigitalAxis: digitalController,
			JoystickIndex:  int(j),
		},
		DeviceName:          name,
		InputFire:           int(pixelgl.KeyLeftControl),
		InputNextWeapon:     int(pixelgl.KeyA),
		InputPreviousWeapon: int(pixelgl.KeyZ),
	}

	if digitalController {
		// Checked this with a SpeedLink Competition Pro USB
		result.InputFire = int(pixelgl.ButtonX)
		result.InputNextWeapon = int(pixelgl.ButtonCircle)
		result.InputPreviousWeapon = int(pixelgl.ButtonLeftBumper)
	} else {
		// Checked this with XBox360 and PlayStation controllers
		result.InputFire = int(pixelgl.ButtonA)
		result.InputNextWeapon = int(pixelgl.ButtonX)
		result.InputPreviousWeapon = int(pixelgl.ButtonY)
	}
	return result
}

func NewKeyboardConfiguration() InputDeviceConfiguration {
	return InputDeviceConfiguration{
		KeyboardConfiguration: KeyboardConfiguration{
			InputUp:    int(pixelgl.KeyUp),
			InputDown:  int(pixelgl.KeyDown),
			InputLeft:  int(pixelgl.KeyLeft),
			InputRight: int(pixelgl.KeyRight),
		},
		DeviceName:          DeviceNameKeyboard,
		InputFire:           int(pixelgl.KeyLeftControl),
		InputNextWeapon:     int(pixelgl.KeyA),
		InputPreviousWeapon: int(pixelgl.KeyZ),
	}
}

func (cc InputDeviceConfiguration) String() string {
	return fmt.Sprintf("ControllerConfiguration{DeviceName: %s}", cc.DeviceName)
}
