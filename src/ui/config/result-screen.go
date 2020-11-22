package config

import (
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
)

type ResultScreen struct {
	inputController      input.Controller
	screenChangeRequired common.ScreenChangeCallback
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
}

func (s *ResultScreen) SetUp() {
	// s.textDimensions = common.GetTextDimensions(text.New(pixel.V(0, 0), common.DefaultAtlas), txtOnePlayerGame, txtTwoPlayerGame)
}

func (s *ResultScreen) Update(_ int64) {
	s.window.Clear(common.Black)
	// TODO: Display the name of the input device assigned to each player
	// TODO: Proceed to Map screen after 3 seconds / click
}

func (s *ResultScreen) TearDown() {}

func (s *ResultScreen) SetInputController(controller input.Controller) {
	s.inputController = controller
}

func (s *ResultScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *ResultScreen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *ResultScreen) String() string {
	return string(common.ConfigurationResult)
}
