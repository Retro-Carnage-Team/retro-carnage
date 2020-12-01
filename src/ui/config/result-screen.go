package config

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

type ResultScreen struct {
	infoTextPlayerOne    string
	infoTextPlayerTwo    string
	inputController      input.Controller
	screenChangeRequired common.ScreenChangeCallback
	textDimensions       map[string]*geometry.Point
	timeElapsed          int64
	window               *pixelgl.Window
}

const txtOnePlayerGame = "1 PLAYER GAME"
const txtTwoPlayerGame = "2 PLAYER GAME"

func (s *ResultScreen) SetUp() {
	s.infoTextPlayerOne = "PLAYER 1: "
	s.infoTextPlayerTwo = "PLAYER 2: "

	name, err := s.inputController.GetControllerName(0)
	if nil == err {
		s.infoTextPlayerOne += name
	} else {
		logging.Warning.Printf("Failed to get controller name for player 0: %v", err)
	}

	if 2 == characters.PlayerController.NumberOfPlayers() {
		name, err = s.inputController.GetControllerName(1)
		if nil == err {
			s.infoTextPlayerTwo += name
		} else {
			logging.Warning.Printf("Failed to get controller name for player 1: %v", err)
		}
	}

	s.textDimensions = common.GetTextDimensions(text.New(pixel.V(0, 0), common.DefaultAtlas),
		txtOnePlayerGame, txtTwoPlayerGame, s.infoTextPlayerOne, s.infoTextPlayerTwo)
}

func (s *ResultScreen) Update(timeElapsedInMs int64) {
	s.timeElapsed += timeElapsedInMs
	s.window.Clear(common.Black)

	if 1 == characters.PlayerController.NumberOfPlayers() {
		common.DrawLineToScreenCenter(s.window, txtOnePlayerGame, 2, common.Green, s.textDimensions[txtOnePlayerGame])
		common.DrawLineToScreenCenter(s.window, s.infoTextPlayerOne, -1, common.Yellow, s.textDimensions[s.infoTextPlayerOne])
	} else {
		common.DrawLineToScreenCenter(s.window, txtTwoPlayerGame, 2, common.Green, s.textDimensions[txtTwoPlayerGame])
		common.DrawLineToScreenCenter(s.window, s.infoTextPlayerOne, -1, common.Yellow, s.textDimensions[s.infoTextPlayerOne])
		common.DrawLineToScreenCenter(s.window, s.infoTextPlayerTwo, -2.5, common.Yellow, s.textDimensions[s.infoTextPlayerTwo])
	}

	if s.timeElapsed >= 2500 {
		s.screenChangeRequired(common.Mission)
	}
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
