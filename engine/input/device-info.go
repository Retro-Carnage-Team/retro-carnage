package input

import "fmt"

type DeviceInfo struct {
	DeviceName    string
	JoystickIndex int
}

func (di DeviceInfo) String() string {
	return fmt.Sprintf("DeviceInfo{DeviceName: %s, JoystickIndex: %d}", di.DeviceName, di.JoystickIndex)
}
