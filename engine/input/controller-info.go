package input

import "fmt"

type ControllerInfo struct {
	DeviceName    string
	JoystickIndex int
}

func (di ControllerInfo) String() string {
	return fmt.Sprintf("ControllerInfo{DeviceName: %s, JoystickIndex: %d}", di.DeviceName, di.JoystickIndex)
}
