package config

import (
	"fmt"
	"retro-carnage/config"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

const (
	minWindowWidth                   = 1024
	minWindowHeight                  = 768
	optionVideoUsePrimaryMonitor int = iota
	optionVideoUseOtherMonitor
	optionVideoPreviousMonitor
	optionVideoNextMonitor
	optionVideoFullscreen
	optionVideoWindowed
	optionVideoReduceWindowWidth
	optionVideoIncreaseWindowWidth
	optionVideoReduceWindowHeight
	optionVideoIncreaseWindowHeight
	optionVideoSave
	optionVideoBack
)

var (
	optionVideoFocusChanges = []focusChange{
		{movedDown: true, currentSelection: []int{optionVideoUsePrimaryMonitor}, nextSelection: optionVideoUseOtherMonitor},
		{movedUp: true, currentSelection: []int{optionVideoUseOtherMonitor}, nextSelection: optionVideoUsePrimaryMonitor},
		{movedRight: true, currentSelection: []int{optionVideoUseOtherMonitor}, nextSelection: optionVideoPreviousMonitor},
		{movedLeft: true, currentSelection: []int{optionVideoPreviousMonitor}, nextSelection: optionVideoUseOtherMonitor},
		{movedRight: true, currentSelection: []int{optionVideoPreviousMonitor}, nextSelection: optionVideoNextMonitor},
		{movedLeft: true, currentSelection: []int{optionVideoNextMonitor}, nextSelection: optionVideoPreviousMonitor},
		{movedUp: true, currentSelection: []int{optionVideoUseOtherMonitor, optionVideoPreviousMonitor, optionVideoNextMonitor}, nextSelection: optionVideoUsePrimaryMonitor},
		{movedDown: true, currentSelection: []int{optionVideoUseOtherMonitor, optionVideoPreviousMonitor, optionVideoNextMonitor}, nextSelection: optionVideoFullscreen},
		{movedRight: true, currentSelection: []int{optionVideoFullscreen}, nextSelection: optionVideoWindowed},
		{movedLeft: true, currentSelection: []int{optionVideoWindowed}, nextSelection: optionVideoFullscreen},
		{movedUp: true, currentSelection: []int{optionVideoFullscreen, optionVideoWindowed}, nextSelection: optionVideoUseOtherMonitor},
		{movedDown: true, currentSelection: []int{optionVideoFullscreen, optionVideoWindowed}, nextSelection: optionVideoReduceWindowWidth},
		{movedRight: true, currentSelection: []int{optionVideoReduceWindowWidth}, nextSelection: optionVideoIncreaseWindowWidth},
		{movedLeft: true, currentSelection: []int{optionVideoIncreaseWindowWidth}, nextSelection: optionVideoReduceWindowWidth},
		{movedUp: true, currentSelection: []int{optionVideoIncreaseWindowWidth, optionVideoReduceWindowWidth}, nextSelection: optionVideoFullscreen},
		{movedDown: true, currentSelection: []int{optionVideoIncreaseWindowWidth, optionVideoReduceWindowWidth}, nextSelection: optionVideoReduceWindowHeight},
		{movedRight: true, currentSelection: []int{optionVideoReduceWindowHeight}, nextSelection: optionVideoIncreaseWindowHeight},
		{movedLeft: true, currentSelection: []int{optionVideoIncreaseWindowHeight}, nextSelection: optionVideoReduceWindowHeight},
		{movedUp: true, currentSelection: []int{optionVideoIncreaseWindowHeight, optionVideoReduceWindowHeight}, nextSelection: optionVideoReduceWindowWidth},
		{movedDown: true, currentSelection: []int{optionVideoIncreaseWindowHeight, optionVideoReduceWindowHeight}, nextSelection: optionVideoSave},
		{movedRight: true, currentSelection: []int{optionVideoSave}, nextSelection: optionVideoBack},
		{movedLeft: true, currentSelection: []int{optionVideoBack}, nextSelection: optionVideoSave},
		{movedUp: true, currentSelection: []int{optionVideoSave, optionVideoBack}, nextSelection: optionVideoReduceWindowHeight},
	}
)

type VideoOptionsScreen struct {
	videoConfig          config.VideoConfiguration
	defaultFontSize      int
	inputController      input.InputController
	screenChangeRequired common.ScreenChangeCallback
	selectedOption       int
	textDimensions       map[string]*geometry.Point
	window               *pixelgl.Window
	selectedMonitorIndex int
	maxWidthScreenName   float64
}

func (s *VideoOptionsScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.selectedOption = optionVideoUsePrimaryMonitor
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtVideoSettings, txtSave, txtBack, txtDecrease,
		txtFullScreen, txtIncrease, txtMonitor, txtPrimaryMonitor, txtScreenmode, txtVideoSettings, txtWindowed,
		txtWindowSize, txtSelection, txtOtherMonitor, txtFiveDigits,
	)
	s.videoConfig = config.GetConfigService().LoadVideoConfiguration()
	s.maxWidthScreenName = s.getMaxWidthOfScreenNames()

	var configuredMonitorFound = false
	for i, m := range pixelgl.Monitors() {
		if s.videoConfig.SelectedMonitor != "" && m.Name() == s.videoConfig.SelectedMonitor {
			s.selectedMonitorIndex = i
			var x, y = m.Size()
			if s.videoConfig.Width > int(x) {
				s.videoConfig.Width = int(x)
			}
			if s.videoConfig.Height > int(y) {
				s.videoConfig.Height = int(y)
			}
			configuredMonitorFound = true
		}
	}

	if !configuredMonitorFound {
		s.videoConfig.SelectedMonitor = pixelgl.Monitors()[0].Name()
		s.selectedMonitorIndex = 0
		var x, y = pixelgl.Monitors()[0].Size()
		s.videoConfig.Width = int(x)
		s.videoConfig.Height = int(y)
	}
}

func (s *VideoOptionsScreen) Update(_ int64) {
	s.processUserInput()

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
	if s.selectedOption == optionVideoBack {
		drawTextSelectionRect(s.window, txt.Bounds())
	}

	txt = s.drawText(txtSave, lineLocationX-s.textDimensions[txtSave].X-3*buttonPadding, float64(lineLocationY))
	if s.selectedOption == optionVideoSave {
		drawTextSelectionRect(s.window, txt.Bounds())
	}
}

func (s *VideoOptionsScreen) renderMonitorValues(valueDistanceLeft float64, monitorLabelLocationY float64) {
	var primaryMonitorValueLocationY = monitorLabelLocationY
	var otherMonitorValueLocationY = monitorLabelLocationY - 2.5*s.textDimensions[txtMonitor].Y

	var txt *text.Text
	if s.videoConfig.UsePrimaryMonitor {
		txt = s.drawText(txtSelection, valueDistanceLeft, primaryMonitorValueLocationY)
		if s.selectedOption == optionVideoUsePrimaryMonitor {
			drawTextSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: -2.5 * s.textDimensions[txtMonitor].Y}))
		} else if s.selectedOption == optionVideoUseOtherMonitor {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawTextSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: -2.5 * s.textDimensions[txtMonitor].Y}))
		} else {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: -2.5 * s.textDimensions[txtMonitor].Y}))
		}
	} else {
		txt = s.drawText(txtSelection, valueDistanceLeft, otherMonitorValueLocationY)
		if s.selectedOption == optionVideoUseOtherMonitor {
			drawTextSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: 0, Y: 2.5 * s.textDimensions[txtMonitor].Y}))
		} else if s.selectedOption == optionVideoUsePrimaryMonitor {
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
	if s.selectedOption == optionVideoPreviousMonitor {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var distanceLeftScreenName = txt.Bounds().Max.X + buttonPadding*3
	s.drawText(s.videoConfig.SelectedMonitor, distanceLeftScreenName, otherMonitorValueLocationY)
	txt = s.drawText(txtIncrease, distanceLeftScreenName+s.maxWidthScreenName+buttonPadding*3, otherMonitorValueLocationY)
	if s.selectedOption == optionVideoNextMonitor {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *VideoOptionsScreen) renderScreenModeValues(valueDistanceLeft float64, screenModeLabelLocationY float64) {
	var distanceBetweenScreenModeBoxes = s.textDimensions[txtSelection].X + s.textDimensions[txtFullScreen].X + buttonPadding*6
	if s.videoConfig.FullScreen {
		var txt = s.drawText(txtSelection, valueDistanceLeft, screenModeLabelLocationY)
		if s.selectedOption == optionVideoFullscreen {
			drawTextSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: distanceBetweenScreenModeBoxes, Y: 0}))
		} else if s.selectedOption == optionVideoWindowed {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawTextSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: distanceBetweenScreenModeBoxes, Y: 0}))
		} else {
			drawPossibleSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: distanceBetweenScreenModeBoxes, Y: 0}))
		}
	} else {
		var txt = s.drawText(txtSelection, valueDistanceLeft+distanceBetweenScreenModeBoxes, screenModeLabelLocationY)
		if s.selectedOption == optionVideoWindowed {
			drawTextSelectionRect(s.window, txt.Bounds())
			drawPossibleSelectionRect(s.window, txt.Bounds().Moved(pixel.Vec{X: -distanceBetweenScreenModeBoxes, Y: 0}))
		} else if s.selectedOption == optionVideoFullscreen {
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
	if s.selectedOption == optionVideoReduceWindowWidth {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var widthLeft = txt.Bounds().Max.X +
		buttonPadding*3 +
		(s.textDimensions[txtFiveDigits].X-fonts.GetTextDimension(s.defaultFontSize, strconv.Itoa(s.videoConfig.Width)).X)/2
	s.drawText(strconv.Itoa(s.videoConfig.Width), widthLeft, windowSizeLabelLocationY)

	txt = s.drawText(txtIncrease, txt.Bounds().Max.X+buttonPadding*6+s.textDimensions[txtFiveDigits].X, windowSizeLabelLocationY)
	if s.selectedOption == optionVideoIncreaseWindowWidth {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var secondLineY = windowSizeLabelLocationY - 2.5*s.textDimensions[txtMonitor].Y
	txt = s.drawText(txtDecrease, valueDistanceLeft, secondLineY)
	if s.selectedOption == optionVideoReduceWindowHeight {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	var heightLeft = txt.Bounds().Max.X +
		buttonPadding*3 +
		(s.textDimensions[txtFiveDigits].X-fonts.GetTextDimension(s.defaultFontSize, strconv.Itoa(s.videoConfig.Height)).X)/2
	s.drawText(strconv.Itoa(s.videoConfig.Height), heightLeft, secondLineY)

	txt = s.drawText(txtIncrease, txt.Bounds().Max.X+buttonPadding*6+s.textDimensions[txtFiveDigits].X, secondLineY)
	if s.selectedOption == optionVideoIncreaseWindowHeight {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}
}

func (s *VideoOptionsScreen) TearDown() {
	// no tear down action required
}

func (s *VideoOptionsScreen) SetInputController(controller input.InputController) {
	s.inputController = controller
}

func (s *VideoOptionsScreen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *VideoOptionsScreen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *VideoOptionsScreen) String() string {
	return string(common.ConfigurationVideo)
}

func (s *VideoOptionsScreen) processUserInput() {
	var uiEventState = s.inputController.GetUiEventStateCombined()
	if nil == uiEventState {
		return
	}

focusHandling:
	for _, fc := range optionVideoFocusChanges {
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

func (s *VideoOptionsScreen) processOptionSelected() {
	switch s.selectedOption {
	case optionVideoUsePrimaryMonitor:
		s.videoConfig.UsePrimaryMonitor = true
	case optionVideoUseOtherMonitor:
		s.videoConfig.UsePrimaryMonitor = false
	case optionVideoPreviousMonitor:
		s.selectPreviousVideoNextMonitor()
	case optionVideoNextMonitor:
		s.selectNextVideoNextMonitor()
	case optionVideoFullscreen:
		s.videoConfig.FullScreen = true
	case optionVideoWindowed:
		s.videoConfig.FullScreen = false
	case optionVideoReduceWindowWidth:
		s.decreaseScreenWidth()
	case optionVideoIncreaseWindowWidth:
		s.increaseScreenWidth()
	case optionVideoReduceWindowHeight:
		s.decreaseScreenHeight()
	case optionVideoIncreaseWindowHeight:
		s.increaseScreenHeight()
	case optionVideoSave:
		err := config.GetConfigService().SaveVideoConfiguration(s.videoConfig)
		if nil != err {
			logging.Warning.Printf("failed to save video settings: %s", err)
		}
		s.screenChangeRequired(common.ConfigurationOptions)
	case optionVideoBack:
		s.screenChangeRequired(common.ConfigurationOptions)
	default:
		logging.Error.Fatal("Unexpected selection in VideoOptionsScreen")
	}
}

func (s *VideoOptionsScreen) drawText(output string, x float64, y float64) *text.Text {
	var txt = text.New(pixel.V(x, y), fonts.SizeToFontAtlas[s.defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, output)
	txt.Draw(s.window, pixel.IM)
	return txt
}

func (s *VideoOptionsScreen) selectPreviousVideoNextMonitor() {
	s.selectedMonitorIndex = s.selectedMonitorIndex - 1
	if s.selectedMonitorIndex < 0 {
		s.selectedMonitorIndex = len(pixelgl.Monitors()) - 1
	}
	s.videoConfig.SelectedMonitor = pixelgl.Monitors()[s.selectedMonitorIndex].Name()
}

func (s *VideoOptionsScreen) selectNextVideoNextMonitor() {
	s.selectedMonitorIndex = (s.selectedMonitorIndex + 1) % len(pixelgl.Monitors())
	s.videoConfig.SelectedMonitor = pixelgl.Monitors()[s.selectedMonitorIndex].Name()
}

func (s *VideoOptionsScreen) decreaseScreenWidth() {
	if s.videoConfig.Width > minWindowWidth+10 {
		s.videoConfig.Width = s.videoConfig.Width - 10
	}
}

func (s *VideoOptionsScreen) increaseScreenWidth() {
	s.videoConfig.Width = s.videoConfig.Width + 10
	var x, _ = pixelgl.Monitors()[s.selectedMonitorIndex].Size()
	if s.videoConfig.Width > int(x) {
		s.videoConfig.Width = int(x)
	}
}

func (s *VideoOptionsScreen) decreaseScreenHeight() {
	if s.videoConfig.Height > minWindowHeight+10 {
		s.videoConfig.Height = s.videoConfig.Height - 10
	}
}

func (s *VideoOptionsScreen) increaseScreenHeight() {
	s.videoConfig.Height = s.videoConfig.Height + 10
	var _, y = pixelgl.Monitors()[s.selectedMonitorIndex].Size()
	if s.videoConfig.Height > int(y) {
		s.videoConfig.Height = int(y)
	}
}

func (s *VideoOptionsScreen) getMaxWidthOfScreenNames() float64 {
	var result = 0.0
	for _, m := range pixelgl.Monitors() {
		var width = fonts.GetTextDimension(s.defaultFontSize, m.Name()).X
		if width > result {
			result = width
		}
	}
	return result
}
