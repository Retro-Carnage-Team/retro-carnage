package config

import (
	"fmt"
	"retro-carnage/config"
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/imdraw"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

const (
	optionAudio int = iota
	optionVideo
	optionControls
	optionBack
)

type OptionsScreen struct {
	configService        *config.ConfigService
	defaultFontSize      int
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	selectedOption       int
	textDimensions       map[string]*geometry.Point
	window               *opengl.Window
}

func (s *OptionsScreen) SetUp() {
	s.configService = config.GetConfigService()
	s.defaultFontSize = fonts.DefaultFontSize()
	s.selectedOption = optionAudio
	s.textDimensions = fonts.GetTextDimensions(
		s.defaultFontSize, txtAudioSettings, txtInputSettings, txtHeadlineOptions, txtVideoSettings, txtBack,
		txtRestartRequired,
	)
}

func (s *OptionsScreen) Update(_ int64) {
	s.processUserInput()

	// draw headline
	var lineLocationY = s.window.Bounds().H() - headlineDistanceTop
	var txt = s.drawText(txtHeadlineOptions, headlineDistanceLeft, lineLocationY)
	imd := imdraw.New(nil)
	imd.Color = common.White
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(txt.Bounds().Min.X, txt.Bounds().Min.Y), pixel.V(txt.Bounds().Max.X, txt.Bounds().Min.Y))
	imd.Line(4)
	imd.Draw(s.window)

	// option 1: audio
	lineLocationY = lineLocationY - 5*s.textDimensions[txtAudioSettings].Y
	txt = s.drawText(txtAudioSettings, headlineDistanceLeft, lineLocationY)
	if s.selectedOption == optionAudio {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	// option 2: video
	lineLocationY = lineLocationY - 3*s.textDimensions[txtVideoSettings].Y
	txt = s.drawText(txtVideoSettings, headlineDistanceLeft, lineLocationY)
	if s.selectedOption == optionVideo {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	// option 3: controls
	lineLocationY = lineLocationY - 3*s.textDimensions[txtInputSettings].Y
	txt = s.drawText(txtInputSettings, headlineDistanceLeft, lineLocationY)
	if s.selectedOption == optionControls {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	// option 4: back
	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	lineLocationY = headlineDistanceTop
	txt = s.drawText(txtBack, lineLocationX, lineLocationY)
	if s.selectedOption == optionBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	// Restart hint
	if s.configService.IsRestartRequired() {
		var hintLocationX = (s.window.Bounds().W() - s.textDimensions[txtRestartRequired].X) / 2
		var hintLocationY = lineLocationY + 3*s.textDimensions[txtRestartRequired].Y
		var txt = text.New(pixel.V(hintLocationX, hintLocationY), fonts.SizeToFontAtlas[s.defaultFontSize])
		txt.Color = common.Red
		_, _ = fmt.Fprint(txt, txtRestartRequired)
		txt.Draw(s.window, pixel.IM)
	}
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

func (s *OptionsScreen) SetWindow(window *opengl.Window) {
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

func (s *OptionsScreen) drawText(output string, x float64, y float64) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}
