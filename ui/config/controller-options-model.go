package config

import (
	"retro-carnage/config"
	"retro-carnage/input"
)

type controllerOptionsModel struct {
	availableDevices    []input.InputDeviceInfo
	inputConfig         config.InputDeviceConfiguration
	playerIdx           int
	selectedDeviceIndex int
	selectedOption      int
	waitForInput        bool
}
