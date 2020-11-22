package config

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
)

const txtOnePlayerGame = "START 1 PLAYER GAME"
const txtTwoPlayerGame = "START 2 PLAYER GAME"

type Screen struct {
	inputController      input.Controller
	screenChangeRequired common.ScreenChangeCallback
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
}

func (s *Screen) SetUp() {
	s.textDimensions = common.GetTextDimensions(text.New(pixel.V(0, 0), common.DefaultAtlas),
		txtOnePlayerGame, txtTwoPlayerGame)
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.window.Clear(common.Yellow)
}

func (s *Screen) TearDown() {}

func (s *Screen) SetInputController(controller input.Controller) {
	s.inputController = controller
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) String() string {
	return string(common.Loading)
}
