package game

import (
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
)

type Screen struct {
	inputController      input.Controller
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	window               *pixelgl.Window
}

func (s *Screen) SetInputController(ctrl input.Controller) {
	s.inputController = ctrl
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.stereo = assets.NewStereo()
}

func (s *Screen) Update(_ int64) {

}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(common.Game)
}
