// Package loading contains the first screen shown to the user. This contains the message that the game is loading.
// Actually there's not much loading going on. We just play the sound of an Amiga 500 diskette drive loading a game.
package loading

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"retro-carnage.net/assets"
	"retro-carnage.net/engine/geometry"
	"retro-carnage.net/ui/util"
)

const screenTimeout = 8500
const txtFirstLine = "RETRO CARNAGE"
const txtSecondLine = "IS LOADING"

type Screen struct {
	screenChangeRequired util.ScreenChangeCallback
	screenChangeTimeout  int64
	textDimensions       map[string]*geometry.Point
	Window               *pixelgl.Window
}

func (s *Screen) SetUp(screenChangeRequired util.ScreenChangeCallback) {
	s.screenChangeRequired = screenChangeRequired
	s.screenChangeTimeout = 0
	s.textDimensions = util.GetTextDimensions(text.New(pixel.V(0, 0), util.DefaultAtlas),
		txtFirstLine, txtSecondLine)

	var stereo = util.NewStereo()
	stereo.PlayFx(assets.FxLoading)
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	var firstLineDimensions = s.textDimensions[txtFirstLine]
	var firstLineX = (s.Window.Bounds().Max.X - firstLineDimensions.X) / 2
	var firstLineY = (s.Window.Bounds().Max.Y-(3*firstLineDimensions.Y))/2 + firstLineDimensions.Y*1.5

	var secondLineDimensions = s.textDimensions[txtSecondLine]
	var secondLineX = (s.Window.Bounds().Max.X - secondLineDimensions.X) / 2
	var secondLineY = (s.Window.Bounds().Max.Y - (3 * secondLineDimensions.Y)) / 2

	var txt = text.New(pixel.V(firstLineX, firstLineY), util.DefaultAtlas)
	_, _ = fmt.Fprint(txt, txtFirstLine)
	txt.Color = colornames.Red
	txt.Draw(s.Window, pixel.IM)

	txt = text.New(pixel.V(secondLineX, secondLineY), util.DefaultAtlas)
	_, _ = fmt.Fprint(txt, txtSecondLine)
	txt.Color = colornames.Red
	txt.Draw(s.Window, pixel.IM)

	s.screenChangeTimeout += elapsedTimeInMs
	if s.screenChangeTimeout >= screenTimeout {
		s.screenChangeRequired(util.Start)
	}
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	return string(util.Loading)
}
