package config

import (
	"fmt"
	"retro-carnage/config"
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
	"strconv"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/imdraw"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

type VideoOptionsScreen struct {
	controller         *videoOptionsController
	defaultFontSize    int
	maxWidthScreenName float64
	model              *videoOptionsModel
	textDimensions     map[string]*geometry.Point
	window             *opengl.Window
}

func NewVideoOptionsScreen() *VideoOptionsScreen {
	var model = videoOptionsModel{
		selectedOption: optionVideoUsePrimaryMonitor,
		videoConfig:    config.GetConfigService().LoadVideoConfiguration(),
	}
	var controller = newVideoOptionsController(&model)
	var result = VideoOptionsScreen{
		controller: controller,
		model:      &model,
	}
	return &result
}

func (s *VideoOptionsScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtVideoSettings, txtSave, txtBack, txtDecrease,
		txtFullScreen, txtIncrease, txtMonitor, txtPrimaryMonitor, txtScreenmode, txtVideoSettings, txtWindowed,
		txtWindowSize, txtSelection, txtOtherMonitor, txtFiveDigits,
	)
	s.maxWidthScreenName = s.getMaxWidthOfScreenNames()
}

func (s *VideoOptionsScreen) Update(_ int64) {
	s.controller.update()

	// draw headline
	var headlineLocationY = s.window.Bounds().H() - headlineDistanceTop
	var txt = s.drawText(txtVideoSettings, headlineDistanceLeft, headlineLocationY)
	imd := imdraw.New(nil)
	imd.Color = common.White
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(txt.Bounds().Min.X, txt.Bounds().Min.Y), pixel.V(txt.Bounds().Max.X, txt.Bounds().Min.Y))
	imd.Line(4)
	imd.Draw(s.window)

	// Monitor label
	var monitorLabelLocationY = headlineLocationY - 5*s.textDimensions[txtMonitor].Y
	s.drawText(txtMonitor, headlineDistanceLeft, monitorLabelLocationY)

	// Screen mode label
	var screenModeLabelLocationY = monitorLabelLocationY - 5.5*s.textDimensions[txtScreenmode].Y
	s.drawText(txtScreenmode, headlineDistanceLeft, screenModeLabelLocationY)

	// Windows size label
	var windowSizeLabelLocationY = screenModeLabelLocationY - 3*s.textDimensions[txtWindowSize].Y
	txt = s.drawText(txtWindowSize, headlineDistanceLeft, windowSizeLabelLocationY)

	var valueDistanceLeft = headlineDistanceLeft + txt.Bounds().W()/2*3

	// Values
	s.renderMonitorValues(valueDistanceLeft, monitorLabelLocationY)
	s.renderScreenModeValues(valueDistanceLeft, screenModeLabelLocationY)
	s.renderWindowSizeValues(valueDistanceLeft, windowSizeLabelLocationY)

	var lineLocationX = s.window.Bounds().W() - headlineDistanceTop - s.textDimensions[txtBack].X
	var lineLocationY = headlineDistanceTop
	txt = s.drawText(txtBack, lineLocationX, float64(lineLocationY))
	if s.model.selectedOption == optionVideoBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	txt = s.drawText(txtSave, lineLocationX-s.textDimensions[txtSave].X-3*buttonPadding, float64(lineLocationY))
	if s.model.selectedOption == optionVideoSave {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *VideoOptionsScreen) renderMonitorValues(valueDistanceLeft float64, monitorLabelLocationY float64) {
	var primaryMonitorValueLocationY = monitorLabelLocationY
	var otherMonitorValueLocationY = monitorLabelLocationY - 2.5*s.textDimensions[txtMonitor].Y

	var txt *text.Text
	if s.model.videoConfig.UsePrimaryMonitor {
		txt = s.drawText(txtSelection, valueDistanceLeft, primaryMonitorValueLocationY)
		if s.model.selectedOption == optionVideoUsePrimaryMonitor {
			drawTextSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: -2.5 * s.textDimensions[txtMonitor].Y}))
		} else if s.model.selectedOption == optionVideoUseOtherMonitor {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawTextSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: -2.5 * s.textDimensions[txtMonitor].Y}))
		} else {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: -2.5 * s.textDimensions[txtMonitor].Y}))
		}
	} else {
		txt = s.drawText(txtSelection, valueDistanceLeft, otherMonitorValueLocationY)
		if s.model.selectedOption == optionVideoUseOtherMonitor {
			drawTextSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: 2.5 * s.textDimensions[txtMonitor].Y}))
		} else if s.model.selectedOption == optionVideoUsePrimaryMonitor {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawTextSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: 2.5 * s.textDimensions[txtMonitor].Y}))
		} else {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: 2.5 * s.textDimensions[txtMonitor].Y}))
		}
	}

	var monitorValueAfterSelectorX = txt.Bounds().Max.X + buttonPadding*3
	s.drawText(txtPrimaryMonitor, monitorValueAfterSelectorX, primaryMonitorValueLocationY)

	txt = s.drawText(txtOtherMonitor, monitorValueAfterSelectorX, otherMonitorValueLocationY)
	txt = s.drawText(txtDecrease, monitorValueAfterSelectorX+buttonPadding*3+txt.Bounds().W(), otherMonitorValueLocationY)
	if s.model.selectedOption == optionVideoPreviousMonitor {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var distanceLeftScreenName = txt.Bounds().Max.X + buttonPadding*3
	s.drawText(s.model.videoConfig.SelectedMonitor, distanceLeftScreenName, otherMonitorValueLocationY)
	txt = s.drawText(txtIncrease, distanceLeftScreenName+s.maxWidthScreenName+buttonPadding*3, otherMonitorValueLocationY)
	if s.model.selectedOption == optionVideoNextMonitor {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *VideoOptionsScreen) renderScreenModeValues(valueDistanceLeft float64, screenModeLabelLocationY float64) {
	var distanceBetweenScreenModeBoxes = s.textDimensions[txtSelection].X + s.textDimensions[txtFullScreen].X + buttonPadding*6
	if s.model.videoConfig.FullScreen {
		var txt = s.drawText(txtSelection, valueDistanceLeft, screenModeLabelLocationY)
		if s.model.selectedOption == optionVideoFullscreen {
			drawTextSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: distanceBetweenScreenModeBoxes, Y: 0}))
		} else if s.model.selectedOption == optionVideoWindowed {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawTextSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: distanceBetweenScreenModeBoxes, Y: 0}))
		} else {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: distanceBetweenScreenModeBoxes, Y: 0}))
		}
	} else {
		var txt = s.drawText(txtSelection, valueDistanceLeft+distanceBetweenScreenModeBoxes, screenModeLabelLocationY)
		if s.model.selectedOption == optionVideoWindowed {
			drawTextSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: -distanceBetweenScreenModeBoxes, Y: 0}))
		} else if s.model.selectedOption == optionVideoFullscreen {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawTextSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: -distanceBetweenScreenModeBoxes, Y: 0}))
		} else {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: -distanceBetweenScreenModeBoxes, Y: 0}))
		}
	}

	s.drawText(txtFullScreen, valueDistanceLeft+s.textDimensions[txtSelection].X+buttonPadding*3, screenModeLabelLocationY)
	s.drawText(txtWindowed, valueDistanceLeft+distanceBetweenScreenModeBoxes+buttonPadding*3, screenModeLabelLocationY)
}

func (s *VideoOptionsScreen) renderWindowSizeValues(valueDistanceLeft float64, windowSizeLabelLocationY float64) {
	var txt = s.drawText(txtDecrease, valueDistanceLeft, windowSizeLabelLocationY)
	if s.model.selectedOption == optionVideoReduceWindowWidth {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var widthLeft = txt.Bounds().Max.X +
		buttonPadding*3 +
		(s.textDimensions[txtFiveDigits].X-fonts.GetTextDimension(s.defaultFontSize, strconv.Itoa(s.model.videoConfig.Width)).X)/2
	s.drawText(strconv.Itoa(s.model.videoConfig.Width), widthLeft, windowSizeLabelLocationY)

	txt = s.drawText(txtIncrease, txt.Bounds().Max.X+buttonPadding*6+s.textDimensions[txtFiveDigits].X, windowSizeLabelLocationY)
	if s.model.selectedOption == optionVideoIncreaseWindowWidth {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var secondLineY = windowSizeLabelLocationY - 2.5*s.textDimensions[txtMonitor].Y
	txt = s.drawText(txtDecrease, valueDistanceLeft, secondLineY)
	if s.model.selectedOption == optionVideoReduceWindowHeight {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var heightLeft = txt.Bounds().Max.X +
		buttonPadding*3 +
		(s.textDimensions[txtFiveDigits].X-fonts.GetTextDimension(s.defaultFontSize, strconv.Itoa(s.model.videoConfig.Height)).X)/2
	s.drawText(strconv.Itoa(s.model.videoConfig.Height), heightLeft, secondLineY)

	txt = s.drawText(txtIncrease, txt.Bounds().Max.X+buttonPadding*6+s.textDimensions[txtFiveDigits].X, secondLineY)
	if s.model.selectedOption == optionVideoIncreaseWindowHeight {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *VideoOptionsScreen) TearDown() {
	// no tear down action required
}

func (s *VideoOptionsScreen) SetInputController(controller input.InputController) {
	s.controller.setInputController(controller)
}

func (s *VideoOptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *VideoOptionsScreen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *VideoOptionsScreen) String() string {
	return string(common.ConfigurationVideo)
}

func (s *VideoOptionsScreen) drawText(output string, x float64, y float64) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}

func (s *VideoOptionsScreen) getMaxWidthOfScreenNames() float64 {
	var result = 0.0
	for _, m := range opengl.Monitors() {
		var width = fonts.GetTextDimension(s.defaultFontSize, m.Name()).X
		if width > result {
			result = width
		}
	}
	return result
}
