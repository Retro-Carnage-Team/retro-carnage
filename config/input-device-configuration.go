package config

import (
	"fmt"
	"retro-carnage/logging"
	"strings"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
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

func NewGamepadConfiguration(w opengl.Window, j pixel.Joystick) InputDeviceConfiguration {
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
		InputFire:           int(pixel.KeyLeftControl),
		InputNextWeapon:     int(pixel.KeyA),
		InputPreviousWeapon: int(pixel.KeyZ),
	}

	if digitalController {
		// Checked this with a SpeedLink Competition Pro USB
		result.InputFire = int(pixel.GamepadX)
		result.InputNextWeapon = int(pixel.GamepadCircle)
		result.InputPreviousWeapon = int(pixel.GamepadLeftBumper)
	} else {
		// Checked this with XBox360 and PlayStation controllers
		result.InputFire = int(pixel.GamepadA)
		result.InputNextWeapon = int(pixel.GamepadX)
		result.InputPreviousWeapon = int(pixel.GamepadY)
	}
	return result
}

func NewKeyboardConfiguration() InputDeviceConfiguration {
	return InputDeviceConfiguration{
		KeyboardConfiguration: KeyboardConfiguration{
			InputUp:    int(pixel.KeyUp),
			InputDown:  int(pixel.KeyDown),
			InputLeft:  int(pixel.KeyLeft),
			InputRight: int(pixel.KeyRight),
		},
		DeviceName:          DeviceNameKeyboard,
		InputFire:           int(pixel.KeyLeftControl),
		InputNextWeapon:     int(pixel.KeyA),
		InputPreviousWeapon: int(pixel.KeyZ),
	}
}

func (cc InputDeviceConfiguration) String() string {
	return fmt.Sprintf("ControllerConfiguration{DeviceName: %s}", cc.DeviceName)
}
