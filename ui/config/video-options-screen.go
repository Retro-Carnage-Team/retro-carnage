package config

import (
	"fmt"
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
	optionVideoUsePrimaryMonitor    = 1
	optionVideoUseOtherMonitor      = 2
	optionVideoPreviousMonitor      = 3
	optionVideoNextMonitor          = 4
	optionVideoFullscreen           = 5
	optionVideoWindowed             = 6
	optionVideoReduceWindowWidth    = 7
	optionVideoIncreaseWindowWidth  = 8
	optionVideoReduceWindowHeight   = 9
	optionVideoIncreaseWindowHeight = 10
	optionVideoSave                 = 11
	optionVideoBack                 = 12
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
	selectedMonitorMaxX  int
	selectedMonitorMaxY  int
}

func (s *VideoOptionsScreen) SetUp() {
	s.defaultFontSize = fonts.DefaultFontSize()
	s.selectedOption = optionVideoUsePrimaryMonitor
	s.textDimensions = fonts.GetTextDimensions(s.defaultFontSize, txtVideoSettings, txtSave, txtBack, txtDecrease,
		txtFullScreen, txtIncrease, txtMonitor, txtPrimaryMonitor, txtScreenmode, txtVideoSettings, txtWindowed,
		txtWindowSize, txtSelection,
	)
	s.videoConfig = config.GetConfigService().LoadVideoConfiguration()

	var configuredMonitorFound = false
	for i, m := range pixelgl.Monitors() {
		if s.videoConfig.SelectedMonitor != "" && m.Name() == s.videoConfig.SelectedMonitor {
			s.selectedMonitorIndex = i
			var x, y = m.Size()
			s.selectedMonitorMaxX = int(x)
			s.selectedMonitorMaxY = int(y)
			configuredMonitorFound = true
		}
	}

	if !configuredMonitorFound {
		s.videoConfig.SelectedMonitor = pixelgl.Monitors()[0].Name()
		s.selectedMonitorIndex = 0
		var x, y = pixelgl.Monitors()[0].Size()
		s.selectedMonitorMaxX = int(x)
		s.selectedMonitorMaxY = int(y)
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

	// Monitor values
	var primaryMonitorValueLocationY = monitorLabelLocationY
	var otherMonitorValueLocationY = monitorLabelLocationY - 2.5*s.textDimensions[txtMonitor].Y

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

	txt = s.drawText(txtDecrease, monitorValueAfterSelectorX, otherMonitorValueLocationY)
	if s.selectedOption == optionVideoPreviousMonitor {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	txt = s.drawText(txtIncrease, txt.Bounds().Max.X+buttonPadding*3, otherMonitorValueLocationY)
	if s.selectedOption == optionVideoNextMonitor {
		drawTextSelectionRect(s.window, txt.Bounds())
	} else {
		drawPossibleSelectionRect(s.window, txt.Bounds())
	}

	s.drawText(s.videoConfig.SelectedMonitor, txt.Bounds().Max.X+buttonPadding*3, otherMonitorValueLocationY)
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
	if nil != uiEventState {
		if uiEventState.PressedButton {
			s.processOptionSelected()
		} else if uiEventState.MovedDown && s.selectedOption == optionVideoUsePrimaryMonitor {
			s.selectedOption = optionVideoUseOtherMonitor
		} else if uiEventState.MovedUp && s.selectedOption == optionVideoUseOtherMonitor {
			s.selectedOption = optionVideoUsePrimaryMonitor
		} else if uiEventState.MovedRight && s.selectedOption == optionVideoUseOtherMonitor {
			s.selectedOption = optionVideoPreviousMonitor
		} else if uiEventState.MovedLeft && s.selectedOption == optionVideoPreviousMonitor {
			s.selectedOption = optionVideoUseOtherMonitor
		} else if uiEventState.MovedRight && s.selectedOption == optionVideoPreviousMonitor {
			s.selectedOption = optionVideoNextMonitor
		} else if uiEventState.MovedLeft && s.selectedOption == optionVideoNextMonitor {
			s.selectedOption = optionVideoPreviousMonitor
		}
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
	case optionVideoWindowed:
	case optionVideoReduceWindowWidth:
	case optionVideoIncreaseWindowWidth:
	case optionVideoReduceWindowHeight:
	case optionVideoIncreaseWindowHeight:
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
