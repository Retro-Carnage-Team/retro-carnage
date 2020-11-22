// Package title contains the title screen. This screen has a cool background image. You can use the mouse or one of the
// gamepads to proceed to the next screen.
package title

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	_ "image/jpeg"
	"os"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/util"
)

const backgroundImagePath = "./images/backgrounds/title.jpg"
const screenTimeout = 2000

type Screen struct {
	backgroundImageSprite *pixel.Sprite
	inputController       input.Controller
	screenChangeRequired  common.ScreenChangeCallback
	screenChangeTimeout   int64
	textDimensions        map[string]*geometry.Point
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
	var backgroundImage = loadBackgroundImage()
	s.backgroundImageSprite = pixel.NewSprite(backgroundImage, backgroundImage.Bounds())
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.screenChangeTimeout += elapsedTimeInMs

	var factorX = s.window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factorY = s.window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factor = util.Max(factorX, factorY)

	s.backgroundImageSprite.Draw(s.window,
		pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(s.window.Bounds().Center()))

	var uiEventState = s.inputController.GetControllerUiEventStateCombined()
	if (nil != uiEventState && uiEventState.PressedButton) || screenTimeout <= s.screenChangeTimeout {
		s.screenChangeRequired(common.ConfigurationSelect)
	}
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(common.Title)
}

func loadBackgroundImage() pixel.Picture {
	file, err := os.Open(backgroundImagePath)
	if err != nil {
		logging.Error.Fatalf("Failed to load background image for title screen: %v", err)
		return nil
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		logging.Error.Fatalf("Failed to decode background image for title screen: %v", err)
		return nil
	}
	return pixel.PictureDataFromImage(img)
}
