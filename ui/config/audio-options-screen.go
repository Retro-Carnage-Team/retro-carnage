package config

import (
	"fmt"
	"image/color"
	"retro-carnage/config"
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
	optionAudioPlayEffects = 1
	optionAudioPlayMusic   = 2
	optionAudioSave        = 3
	optionAudioBack        = 4
)

var (
	optionAudioFocusChanges = []focusChange{
		{movedDown: true, currentSelection: []int{optionAudioPlayEffects}, nextSelection: optionAudioPlayMusic},
		{movedUp: true, currentSelection: []int{optionAudioPlayMusic}, nextSelection: optionAudioPlayEffects},
		{movedDown: true, currentSelection: []int{optionAudioPlayMusic}, nextSelection: optionAudioSave},
		{movedUp: true, currentSelection: []int{optionAudioSave}, nextSelection: optionAudioPlayMusic},
		{movedRight: true, currentSelection: []int{optionAudioSave}, nextSelection: optionAudioBack},
		{movedLeft: true, currentSelection: []int{optionAudioBack}, nextSelection: optionAudioSave},
	}
)

type AudioOptionsScreen struct {
	audioConfig          config.AudioConfiguration
	defaultFontSize      int
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	selectedOption       int
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
}

func (s *AudioOptionsScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.selectedOption = optionAudioPlayEffects
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtAudioSettings, txtSave, txtBack, txtEffects, txtMusic)
	s.audioConfig = config.GetConfigService().LoadAudioConfiguration()
}

func (s *AudioOptionsScreen) Update(_ int64) {
	s.processUserInput()

	// draw headline
	var headlineLocationY = s.window.Bounds().H() - headlineDistanceTop
	var txt = s.drawText(txtAudioSettings, headlineDistanceLeft, headlineLocationY, common.White)
	imd := imdraw.New(nil)
	imd.Color = common.White
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(txt.Bounds().Min.X, txt.Bounds().Min.Y), pixel.V(txt.Bounds().Max.X, txt.Bounds().Min.Y))
	imd.Line(4)
	imd.Draw(s.window)

	// Effects
	var effectsLabelLocationY = headlineLocationY - 5*s.textDimensions[txtEffects].Y
	txt = s.drawText(txtEffects, headlineDistanceLeft, effectsLabelLocationY, common.White)
	var valueDistanceLeft = headlineDistanceLeft + txt.Bounds().W()/2*3

	if s.audioConfig.PlayEffects {
		txt = s.drawText(txtSelection, valueDistanceLeft, effectsLabelLocationY, common.White)
	} else {
		txt = s.drawText(txtSelection, valueDistanceLeft, effectsLabelLocationY, common.Black)
	}
	if s.selectedOption == optionAudioPlayEffects {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	// Music
	var musicLabelLocationY = effectsLabelLocationY - 3.5*s.textDimensions[txtMusic].Y
	s.drawText(txtMusic, headlineDistanceLeft, musicLabelLocationY, common.White)

	if s.audioConfig.PlayMusic {
		txt = s.drawText(txtSelection, valueDistanceLeft, musicLabelLocationY, common.White)
	} else {
		txt = s.drawText(txtSelection, valueDistanceLeft, musicLabelLocationY, common.Black)
	}
	if s.selectedOption == optionAudioPlayMusic {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	// Action buttons
	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	var lineLocationY = headlineDistanceTop
	txt = s.drawText(txtBack, lineLocationX, float64(lineLocationY), common.White)
	if s.selectedOption == optionAudioBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	txt = s.drawText(txtSave, lineLocationX-s.textDimensions[txtSave].X-3*buttonPadding, float64(lineLocationY), common.White)
	if s.selectedOption == optionAudioSave {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *AudioOptionsScreen) TearDown() {
	// no tear down action required
}

func (s *AudioOptionsScreen) SetInputController(controller input.InputController) {
	s.inputController = controller
}

func (s *AudioOptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *AudioOptionsScreen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *AudioOptionsScreen) String() string {
	return string(common.ConfigurationVideo)
}

func (s *AudioOptionsScreen) processUserInput() {
	var uiEventState = s.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

focusHandling:
	for _, fc := range optionAudioFocusChanges {
		if fc.movedLeft == uiEventState.MovedLeft &&
			fc.movedRight == uiEventState.MovedRight &&
			fc.movedDown == uiEventState.MovedDown &&
			fc.movedUp == uiEventState.MovedUp {
			for _, i := range fc.currentSelection {
				if i == s.selectedOption {
					s.selectedOption = fc.nextSelection
					break focusHandling
				}
			}
		}
	}

	if uiEventState.PressedButton {
		s.processOptionSelected()
	}
}

func (s *AudioOptionsScreen) processOptionSelected() {
	switch s.selectedOption {
	case optionAudioPlayEffects:
		s.audioConfig.PlayEffects = !s.audioConfig.PlayEffects
	case optionAudioPlayMusic:
		s.audioConfig.PlayMusic = !s.audioConfig.PlayMusic
	case optionAudioSave:
		err := config.GetConfigService().SaveAudioConfiguration(s.audioConfig)
		if nil != err {
			logging.Warning.Printf("failed to save audio settings: %s", err)
		}
		s.screenChangeRequired(common.ConfigurationOptions)
	case optionAudioBack:
		s.screenChangeRequired(common.ConfigurationOptions)
	default:
		logging.Error.Fatal("Unexpected selection in AudioOptionsScreen")
	}
}

func (s *AudioOptionsScreen) drawText(output string, x float64, y float64, col color.RGBA) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = col
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}
