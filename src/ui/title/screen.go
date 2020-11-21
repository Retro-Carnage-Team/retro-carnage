// Package title contains the title screen. This screen has a cool background image. You can use the mouse or one of the
// gamepads to proceed to the next screen.
package title

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	_ "image/jpeg"
	"os"
	"retro-carnage.net/engine/geometry"
	"retro-carnage.net/engine/input"
	"retro-carnage.net/logging"
	"retro-carnage.net/ui/common"
	util "retro-carnage.net/util"
)

const backgroundImagePath = "./images/backgrounds/title.jpg"

type Screen struct {
	backgroundImageSprite *pixel.Sprite
	screenChangeRequired  common.ScreenChangeCallback
	textDimensions        map[string]*geometry.Point
	window                *pixelgl.Window
}

func (s *Screen) SetInputController(_ *input.Controller) {}

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

func (s *Screen) Update(_ int64) {
	var factorX = s.window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factorY = s.window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factor = util.Max(factorX, factorY)

	s.backgroundImageSprite.Draw(s.window,
		pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(s.window.Bounds().Center()))
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
