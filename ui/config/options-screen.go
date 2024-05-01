package config

import (
	cfg "retro-carnage/config"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"

	"github.com/faiface/pixel/pixelgl"
)

type OptionsScreen struct {
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	timeElapsed          int64
	window               *pixelgl.Window
}

func (s *OptionsScreen) SetUp() {
	for _, c := range s.inputController.GetInputDeviceInfos() {
		logging.Info.Printf("Found device %s", c.String())
	}

	var cs = cfg.ConfigService{}
	for _, cc := range s.inputController.GetInputDeviceConfigurations() {
		logging.Info.Printf("Found device configuration %s", cc.String())
	}

	var cc = cfg.InputDeviceConfiguration{
		DeviceName: "Test-Device",
	}
	cs.SaveInputDeviceConfiguration(cc, 3)
}

func (s *OptionsScreen) Update(timeElapsedInMs int64) {
	s.timeElapsed += timeElapsedInMs
	s.window.Clear(common.Black)
}

func (s *OptionsScreen) TearDown() {
	// no tear down action required
}

func (s *OptionsScreen) SetInputController(controller input.InputController) {
	s.inputController = controller
}

func (s *OptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *OptionsScreen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *OptionsScreen) String() string {
	return string(common.ConfigurationOptions)
}
