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
	"retro-carnage.net/logging"
	"retro-carnage.net/ui/util"
	commonUtil "retro-carnage.net/util"
)

const backgroundImagePath = "./images/backgrounds/title.jpg"

type Screen struct {
	backgroundImageSprite *pixel.Sprite
	screenChangeRequired  util.ScreenChangeCallback
	textDimensions        map[string]*geometry.Point
	Window                *pixelgl.Window
}

func (s *Screen) SetUp(screenChangeRequired util.ScreenChangeCallback) {
	s.screenChangeRequired = screenChangeRequired

	var backgroundImage = loadBackgroundImage()
	s.backgroundImageSprite = pixel.NewSprite(backgroundImage, backgroundImage.Bounds())
}

func (s *Screen) Update(_ int64) {
	var factorX = s.Window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factorY = s.Window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factor = commonUtil.Max(factorX, factorY)

	s.backgroundImageSprite.Draw(s.Window,
		pixel.IM.Scaled(pixel.Vec{X: 0, Y: 0}, factor).Moved(s.Window.Bounds().Center()))
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(util.Title)
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
