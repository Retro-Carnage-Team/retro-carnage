package config

import (
	"fmt"
	"image/color"
	"retro-carnage/config"
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/imdraw"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

type AudioOptionsScreen struct {
	controller      *audioOptionsController
	defaultFontSize int
	model           *audioOptionsModel
	textDimensions  map[string]*geometry.Point
	window          *opengl.Window
}

func NewAudioOptionsScreen() *AudioOptionsScreen {
	var model = audioOptionsModel{
		audioConfig:    config.GetConfigService().LoadAudioConfiguration(),
		selectedOption: optionAudioPlayEffects,
	}
	var controller = newAudioOptionsController(&model)
	var result = AudioOptionsScreen{
		controller: controller,
		model:      &model,
	}
	return &result
}

func (s *AudioOptionsScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtAudioSettings, txtSave, txtBack, txtEffects, txtMusic)
}

func (s *AudioOptionsScreen) Update(_ int64) {
	s.controller.update()
	s.drawScreen()
}

func (s *AudioOptionsScreen) TearDown() {
	// no tear down action required
}

func (s *AudioOptionsScreen) SetInputController(controller input.InputController) {
	s.controller.setInputController(controller)
}

func (s *AudioOptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *AudioOptionsScreen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *AudioOptionsScreen) String() string {
	return string(common.ConfigurationVideo)
}

func (s *AudioOptionsScreen) drawScreen() {
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

	if s.model.audioConfig.PlayEffects {
		txt = s.drawText(txtSelection, valueDistanceLeft, effectsLabelLocationY, common.White)
	} else {
		txt = s.drawText(txtSelection, valueDistanceLeft, effectsLabelLocationY, common.Black)
	}
	if s.model.selectedOption == optionAudioPlayEffects {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	// Music
	var musicLabelLocationY = effectsLabelLocationY - 3.5*s.textDimensions[txtMusic].Y
	s.drawText(txtMusic, headlineDistanceLeft, musicLabelLocationY, common.White)

	if s.model.audioConfig.PlayMusic {
		txt = s.drawText(txtSelection, valueDistanceLeft, musicLabelLocationY, common.White)
	} else {
		txt = s.drawText(txtSelection, valueDistanceLeft, musicLabelLocationY, common.Black)
	}
	if s.model.selectedOption == optionAudioPlayMusic {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	// Action buttons
	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	var lineLocationY = headlineDistanceTop
	txt = s.drawText(txtBack, lineLocationX, float64(lineLocationY), common.White)
	if s.model.selectedOption == optionAudioBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	txt = s.drawText(txtSave, lineLocationX-s.textDimensions[txtSave].X-3*buttonPadding, float64(lineLocationY), common.White)
	if s.model.selectedOption == optionAudioSave {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *AudioOptionsScreen) drawText(output string, x float64, y float64, col color.RGBA) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = col
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}
