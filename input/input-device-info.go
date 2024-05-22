package input

import "fmt"

type InputDeviceInfo struct {
	DeviceName    string
	JoystickIndex int
}

func (di InputDeviceInfo) String() string {
	return fmt.Sprintf("ControllerInfo{DeviceName: %s, JoystickIndex: %d}", di.DeviceName, di.JoystickIndex)
}
