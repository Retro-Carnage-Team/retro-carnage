package config

import (
	"fmt"
	"math"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

const (
	optionOnePlayer        = 1
	optionTwoPlayers       = 2
	optionOptions          = 3
	txtSelectOnePlayerGame = "START 1 PLAYER GAME"
	txtSelectTwoPlayerGame = "START 2 PLAYER GAME"
	txtSelectOptions       = "OPTIONS"
)

type SelectScreen struct {
	defaultFontSize      int
	inputController      input.InputController
	multiplayerPossible  bool
	screenChangeRequired common.ScreenChangeCallback
	selectedOption       int
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
}

func (s *SelectScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.multiplayerPossible = len(s.inputController.GetInputDeviceInfos()) > 1
	s.selectedOption = optionOnePlayer
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtSelectOnePlayerGame, txtSelectTwoPlayerGame, txtSelectOptions)
}

func (s *SelectScreen) Update(_ int64) {
	s.processUserInput()

	var vertCenter = s.window.Bounds().Max.Y / 2

	var firstLineX = (s.window.Bounds().Max.X - s.textDimensions[txtSelectOnePlayerGame].X) / 2
	var firstLineY = vertCenter + 1.5*s.textDimensions[txtSelectOnePlayerGame].Y

	var txt = text.New(pixel.V(firstLineX, firstLineY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtSelectOnePlayerGame)
	txt.Draw(s.window, pixel.IM)

	var secondLineX = (s.window.Bounds().Max.X - s.textDimensions[txtSelectTwoPlayerGame].X) / 2
	var secondLineY = vertCenter + -1.5*s.textDimensions[txtSelectTwoPlayerGame].Y

	var startLine3 = -1.5
	if s.multiplayerPossible {
		txt = text.New(pixel.V(secondLineX, secondLineY), fonts.SizeToFontAtlas[s.defaultFontSize])
		txt.Color = common.White
		_, _ = fmt.Fprint(txt, txtSelectTwoPlayerGame)
		txt.Draw(s.window, pixel.IM)
		startLine3 = -4.5
	}

	var thirdLineX = (s.window.Bounds().Max.X - s.textDimensions[txtSelectOptions].X) / 2
	var thirdLineY = vertCenter + startLine3*s.textDimensions[txtSelectOptions].Y

	txt = text.New(pixel.V(thirdLineX, thirdLineY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtSelectOptions)
	txt.Draw(s.window, pixel.IM)

	var bottomFirst = firstLineY - buttonPadding
	var bottomSecond = secondLineY - buttonPadding
	var bottomThird = thirdLineY - buttonPadding

	var topFirst = firstLineY + s.textDimensions[txtSelectOnePlayerGame].Y
	var topSecond = secondLineY + s.textDimensions[txtSelectTwoPlayerGame].Y
	var topThird = thirdLineY + s.textDimensions[txtSelectOptions].Y

	var left = math.Min(math.Min(firstLineX, secondLineX), thirdLineX) - buttonPadding
	var right = math.Max(firstLineX, secondLineX) +
		math.Max(
			math.Max(
				s.textDimensions[txtSelectOnePlayerGame].X,
				s.textDimensions[txtSelectTwoPlayerGame].X,
			),
			s.textDimensions[txtSelectOptions].X,
		) + buttonPadding

	if s.selectedOption == optionOnePlayer {
		drawSelectionRect(s.window, left, bottomFirst, right, topFirst)
	} else if s.selectedOption == optionTwoPlayers {
		drawSelectionRect(s.window, left, bottomSecond, right, topSecond)
	} else if s.selectedOption == optionOptions {
		drawSelectionRect(s.window, left, bottomThird, right, topThird)
	}
}

func (s *SelectScreen) TearDown() {
	// no tear down action required
}

func (s *SelectScreen) SetInputController(controller input.InputController) {
	s.inputController = controller
}

func (s *SelectScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *SelectScreen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *SelectScreen) String() string {
	return string(common.ConfigurationSelect)
}

func (s *SelectScreen) processUserInput() {
	var uiEventState = s.inputController.GetUiEventStateCombined()
	if nil != uiEventState {
		if uiEventState.PressedButton {
			s.processOptionSelected()
		} else if uiEventState.MovedUp {
			if s.selectedOption > optionOnePlayer {
				s.selectedOption = s.selectedOption - 1
			}
		} else if uiEventState.MovedDown {
			if s.selectedOption == optionOnePlayer && s.multiplayerPossible {
				s.selectedOption = optionTwoPlayers
			} else {
				s.selectedOption = optionOptions
			}
		}
	}
}

func (s *SelectScreen) processOptionSelected() {
	if s.selectedOption == optionOnePlayer || s.selectedOption == optionTwoPlayers {
		s.inputController.AssignInputDevicesToPlayers()
		characters.PlayerController.StartNewGame(s.selectedOption)
		s.screenChangeRequired(common.ConfigurationResult)
	} else {
		s.screenChangeRequired(common.ConfigurationOptions)
	}
}
