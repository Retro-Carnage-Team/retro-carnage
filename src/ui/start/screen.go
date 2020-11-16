// Package start contains the second screen shown to the user. The screen displays a copyright notice and the dedication
// lines. Once the screen has been loaded, the theme song gets buffered. The next screen gets displayed when the theme
// song has been fully buffered.
package start

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"retro-carnage.net/assets"
	"retro-carnage.net/engine/geometry"
	"retro-carnage.net/ui/util"
)

const screenTimeout = 2000
const txtFirstLine = "RETRO CARNAGE"
const txtSecondLine = "(C) 2020 THOMAS WERNER"
const txtThirdLine = "Dedicated to Jonathan Werner"
const txtFourthLine = "Inspired by \"DOGS OF WAR\""
const txtFifthLine = "(C) 1989 by Elite Systems Ltd."

type Screen struct {
	screenChangeRequired util.ScreenChangeCallback
	screenChangeTimeout  int64
	stereo               *util.Stereo
	textDimensions       map[string]*geometry.Point
	themeLoaded          bool
	Window               *pixelgl.Window
}

func (s *Screen) SetUp(screenChangeRequired util.ScreenChangeCallback) {
	s.screenChangeRequired = screenChangeRequired
	s.screenChangeTimeout = 0
	s.textDimensions = util.GetTextDimensions(text.New(pixel.V(0, 0), util.DefaultAtlas),
		txtFirstLine, txtSecondLine, txtThirdLine, txtFourthLine, txtFifthLine)

	s.stereo = util.NewStereo()
	s.themeLoaded = false
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.screenChangeTimeout += elapsedTimeInMs
	if s.themeLoaded {
		s.screenChangeRequired(util.Title)
	}
	s.renderScreen()
	if !s.themeLoaded && (s.screenChangeTimeout > 100) {
		s.stereo.PlaySong(assets.ThemeSong)
		s.themeLoaded = true
	}
}

func (s *Screen) TearDown() {

}

func (s *Screen) String() string {
	return string(util.Start)
}

func (s *Screen) renderScreen() {
	// TODO: Fix positions and draw remaining lines
	var firstLineDimensions = s.textDimensions[txtFirstLine]
	var firstLineX = (s.Window.Bounds().Max.X - firstLineDimensions.X) / 2
	var firstLineY = (s.Window.Bounds().Max.Y-(3*firstLineDimensions.Y))/2 + firstLineDimensions.Y*1.5

	var secondLineDimensions = s.textDimensions[txtSecondLine]
	var secondLineX = (s.Window.Bounds().Max.X - secondLineDimensions.X) / 2
	var secondLineY = (s.Window.Bounds().Max.Y - (3 * secondLineDimensions.Y)) / 2

	var txt = text.New(pixel.V(firstLineX, firstLineY), util.DefaultAtlas)
	txt.Color = util.Red
	_, _ = fmt.Fprint(txt, txtFirstLine)
	txt.Draw(s.Window, pixel.IM)

	txt = text.New(pixel.V(secondLineX, secondLineY), util.DefaultAtlas)
	txt.Color = util.Yellow
	_, _ = fmt.Fprint(txt, txtSecondLine)
	txt.Draw(s.Window, pixel.IM)
}
