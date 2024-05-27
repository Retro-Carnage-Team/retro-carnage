// Package title contains the title screen. This screen has a cool background image. You can use the mouse or one of the
// gamepads to proceed to the next screen.
package title

import (
	_ "image/jpeg"
	"math"
	"retro-carnage/assets"
	"retro-carnage/engine/cheat"
	"retro-carnage/input"
	"retro-carnage/ui/common"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const backgroundImagePath = "images/other/title.jpg"
const screenTimeout = 60_000

type Screen struct {
	backgroundImageSprite *pixel.Sprite
	cheatController       *cheat.CheatController
	inputController       input.InputController
	screenChangeRequired  common.ScreenChangeCallback
	screenChangeTimeout   int64
	stereo                *assets.Stereo
	window                *opengl.Window
}

func (s *Screen) SetInputController(controller input.InputController) {
	s.inputController = controller
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.backgroundImageSprite = assets.SpriteRepository.Get(backgroundImagePath)
	s.cheatController = cheat.GetCheatController()
	s.stereo = assets.NewStereo()
	common.TitleScreenInit()
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.screenChangeTimeout += elapsedTimeInMs

	var factorX = s.window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factorY = s.window.Bounds().Max.Y / s.backgroundImageSprite.Picture().Bounds().Max.Y
	var factor = math.Max(factorX, factorY)

	s.backgroundImageSprite.Draw(s.window,
		pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(s.window.Bounds().Center()))

	for _, btn := range common.KeyboardButtons {
		if s.window.JustPressed(btn) {
			if s.cheatController.HandleKeyboardInput(btn) {
				s.stereo.PlayFx(assets.FxCheatSwitch)
			}
			break
		}
	}

	var uiEventState = s.inputController.GetUiEventStateCombined()
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
