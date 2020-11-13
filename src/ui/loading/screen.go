package loading

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"retro-carnage.net/assets"
	"retro-carnage.net/engine/geometry"
	"retro-carnage.net/util"
)

const txtFirstLine = "RETRO CARNAGE"
const txtSecondLine = "IS LOADING"

type Screen struct {
	firstLineDimensions  *geometry.Point
	secondLineDimensions *geometry.Point
	soundBoard           *util.Stereo
	Window               *pixelgl.Window
}

func (s *Screen) SetUp() {
	var txt = text.New(pixel.V(0, 0), util.DefaultAtlas)
	_, _ = fmt.Fprint(txt, txtFirstLine)
	s.firstLineDimensions = &geometry.Point{X: txt.Dot.X, Y: txt.LineHeight}

	txt.Clear()
	_, _ = fmt.Fprint(txt, txtSecondLine)
	s.secondLineDimensions = &geometry.Point{X: txt.Dot.X, Y: txt.LineHeight}

	s.soundBoard = util.NewStereo()
	s.soundBoard.PlayFx(assets.FxLoading)
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	var firstLineX = (s.Window.Bounds().Max.X - s.firstLineDimensions.X) / 2
	var firstLineY = (s.Window.Bounds().Max.Y-(3*s.firstLineDimensions.Y))/2 + s.firstLineDimensions.Y*1.5

	var secondLineX = (s.Window.Bounds().Max.X - s.secondLineDimensions.X) / 2
	var secondLineY = (s.Window.Bounds().Max.Y - (3 * s.secondLineDimensions.Y)) / 2

	var txt = text.New(pixel.V(firstLineX, firstLineY), util.DefaultAtlas)
	_, _ = fmt.Fprint(txt, txtFirstLine)
	txt.Color = colornames.Red
	txt.Draw(s.Window, pixel.IM)

	txt = text.New(pixel.V(secondLineX, secondLineY), util.DefaultAtlas)
	_, _ = fmt.Fprint(txt, txtSecondLine)
	txt.Color = colornames.Red
	txt.Draw(s.Window, pixel.IM)
}

func (s *Screen) TearDown() {

}

func (s *Screen) String() string {
	return "Loading Screen"
}
