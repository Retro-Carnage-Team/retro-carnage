package config

import (
	"retro-carnage/engine/characters"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/faiface/pixel/pixelgl"
)

type ResultScreen struct {
	infoTextPlayerOne    string
	infoTextPlayerTwo    string
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	timeElapsed          int64
	window               *pixelgl.Window
}

const txtOnePlayerGame = "1 PLAYER GAME"
const txtTwoPlayerGame = "2 PLAYER GAME"

func (s *ResultScreen) SetUp() {
	s.infoTextPlayerOne = "PLAYER 1: "
	s.infoTextPlayerTwo = "PLAYER 2: "

	name, err := s.inputController.GetInputDeviceName(0)
	if nil == err {
		s.infoTextPlayerOne += name
	} else {
		logging.Warning.Printf("Failed to get controller name for player 0: %v", err)
	}

	if characters.PlayerController.NumberOfPlayers() == 2 {
		name, err = s.inputController.GetInputDeviceName(1)
		if nil == err {
			s.infoTextPlayerTwo += name
		} else {
			logging.Warning.Printf("Failed to get controller name for player 1: %v", err)
		}
	}
}

func (s *ResultScreen) Update(timeElapsedInMs int64) {
	s.timeElapsed += timeElapsedInMs
	s.window.Clear(common.Black)

	renderer := fonts.TextRenderer{Window: s.window}
	if characters.PlayerController.NumberOfPlayers() == 1 {
		renderer.DrawLineToScreenCenter(txtOnePlayerGame, 2, common.Green)
		renderer.DrawLineToScreenCenter(s.infoTextPlayerOne, -1, common.Yellow)
	} else {
		renderer.DrawLineToScreenCenter(txtTwoPlayerGame, 2, common.Green)
		renderer.DrawLineToScreenCenter(s.infoTextPlayerOne, -1, common.Yellow)
		renderer.DrawLineToScreenCenter(s.infoTextPlayerTwo, -2.5, common.Yellow)
	}

	var uiEventState = s.inputController.GetUiEventStateCombined()
	if nil != uiEventState && uiEventState.PressedButton {
		s.screenChangeRequired(common.Mission)
	} else if s.timeElapsed >= 2500 {
		s.screenChangeRequired(common.Mission)
	}
}

func (s *ResultScreen) TearDown() {
	// no tear down action required
}

func (s *ResultScreen) SetInputController(controller input.InputController) {
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
