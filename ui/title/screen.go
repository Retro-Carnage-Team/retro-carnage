// Package title contains the title screen. This screen has a cool background image. You can use the mouse or one of the
// gamepads to proceed to the next screen.
package title

import (
	_ "image/jpeg"
	"math"
	"retro-carnage/assets"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const backgroundImagePath = "images/other/title.jpg"
const screenTimeout = 60_000

type Screen struct {
	backgroundImageSprite *pixel.Sprite
	inputController       input.Controller
	screenChangeRequired  common.ScreenChangeCallback
	screenChangeTimeout   int64
	window                *pixelgl.Window
}

func (s *Screen) SetInputController(controller input.Controller) {
	s.inputController = controller
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.backgroundImageSprite = assets.SpriteRepository.Get(backgroundImagePath)
	common.TitleScreenInit()
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.screenChangeTimeout += elapsedTimeInMs

	var factorX = s.window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factorY = s.window.Bounds().Max.Y / s.backgroundImageSprite.Picture().Bounds().Max.Y
	var factor = math.Max(factorX, factorY)

	s.backgroundImageSprite.Draw(s.window,
		pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(s.window.Bounds().Center()))

	var uiEventState = s.inputController.ControllerUiEventStateCombined()
	if (nil != uiEventState && uiEventState.PressedButton) || screenTimeout <= s.screenChangeTimeout {
		s.screenChangeRequired(common.ConfigurationSelect)
	}
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.Title)
}
