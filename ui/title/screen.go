// Package title contains the title screen. This screen has a cool background image. You can use the mouse or one of the
// gamepads to proceed to the next screen.
package title

import (
	_ "image/jpeg"
	"math"
	"retro-carnage/assets"
	"retro-carnage/input"
	"retro-carnage/ui/common"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const backgroundImagePath = "images/other/title.jpg"

type Screen struct {
	backgroundImageSprite *pixel.Sprite
	controller            *controller
	window                *opengl.Window
}

func NewScreen() *Screen {
	var result = Screen{
		controller: newController(),
	}
	return &result
}

func (s *Screen) SetInputController(controller input.InputController) {
	s.controller.setInputController(controller)
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.backgroundImageSprite = assets.SpriteRepository.Get(backgroundImagePath)
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.controller.update(elapsedTimeInMs, s.window)
	s.drawScreen()
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.Title)
}

func (s *Screen) drawScreen() {
	var factorX = s.window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factorY = s.window.Bounds().Max.Y / s.backgroundImageSprite.Picture().Bounds().Max.Y
	var factor = math.Max(factorX, factorY)
	s.backgroundImageSprite.Draw(
		s.window,
		pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(s.window.Bounds().Center()))
}
