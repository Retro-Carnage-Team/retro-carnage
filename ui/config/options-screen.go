package config

import (
	"fmt"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

const (
	optionAudio        = 1
	optionVideo        = 2
	optionControls     = 3
	optionBack         = 4
	txtAudio           = "AUDIO SETTINGS"
	txtControls        = "INPUT SETTINGS"
	txtHeadlineOptions = "OPTIONS"
	txtVideo           = "VIDEO SETTINGS"
	txtBack            = "BACK"
)

type OptionsScreen struct {
	defaultFontSize      int
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	selectedOption       int
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
}

func (s *OptionsScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.selectedOption = optionAudio
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtAudio, txtControls, txtHeadlineOptions, txtVideo, txtBack)
}

func (s *OptionsScreen) Update(_ int64) {
	s.processUserInput()

	// draw headline
	var lineLocationY = s.window.Bounds().H() - headlineDistanceTop
	var txt = text.New(pixel.V(headlineDistanceLeft, lineLocationY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtHeadlineOptions)
	txt.Draw(s.window, pixel.IM)

	// underline headline
	imd := imdraw.New(nil)
	imd.Color = common.White
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(txt.Bounds().Min.X, txt.Bounds().Min.Y), pixel.V(txt.Bounds().Max.X, txt.Bounds().Min.Y))
	imd.Line(4)
	imd.Draw(s.window)

	lineLocationY = lineLocationY - 5*s.textDimensions[txtAudio].Y
	txt = text.New(pixel.V(headlineDistanceLeft, lineLocationY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtAudio)
	txt.Draw(s.window, pixel.IM)

	lineLocationY = lineLocationY - 3*s.textDimensions[txtAudio].Y
	txt = text.New(pixel.V(headlineDistanceLeft, lineLocationY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtVideo)
	txt.Draw(s.window, pixel.IM)

	lineLocationY = lineLocationY - 3*s.textDimensions[txtAudio].Y
	txt = text.New(pixel.V(headlineDistanceLeft, lineLocationY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtControls)
	txt.Draw(s.window, pixel.IM)

	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	lineLocationY = headlineDistanceTop
	txt = text.New(pixel.V(lineLocationX, lineLocationY), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, txtBack)
	txt.Draw(s.window, pixel.IM)
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

func (s *OptionsScreen) processUserInput() {
	var uiEventState = s.inputController.GetUiEventStateCombined()
	if nil != uiEventState {
		if uiEventState.PressedButton {
			s.processOptionSelected()
		} else if uiEventState.MovedUp && s.selectedOption > optionAudio {
			s.selectedOption = s.selectedOption - 1
		} else if uiEventState.MovedDown && s.selectedOption < optionBack {
			s.selectedOption = s.selectedOption + 1
		}
	}
}

func (s *OptionsScreen) processOptionSelected() {
	switch s.selectedOption {
	case optionAudio:
		s.screenChangeRequired(common.ConfigurationAudio)
	case optionVideo:
		s.screenChangeRequired(common.ConfigurationVideo)
	case optionControls:
		s.screenChangeRequired(common.ConfigurationControls)
	case optionBack:
		s.screenChangeRequired(common.ConfigurationSelect)
	default:
		logging.Error.Fatal("Unexpected selection in OptionsScreen")
	}
}
