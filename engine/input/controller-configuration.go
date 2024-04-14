package input

const (
	DeviceNameKeyboard = "Keyboard"
)

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

// ControllerConfiguration holds the configuration for a specific device. This combines common aspects
// with GamepadConfiguration and KeyboardConfigurations.
type ControllerConfiguration struct {
	GamepadConfiguration  `json:"gamepadConfig"`
	KeyboardConfiguration `json:"keyboardConfig"`
	DeviceName            string `json:"deviceName"`
	InputFire             int    `json:"inputFire"`
	InputNextWeapon       int    `json:"inputNextWeapon"`
	InputPreviousWeapon   int    `json:"inputPrevWeapon"`
}
