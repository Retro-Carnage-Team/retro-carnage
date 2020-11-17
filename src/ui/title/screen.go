// Package title contains the title screen. This screen has a cool background image. You can use the mouse or one of the
// gamepads to proceed to the next screen.
package title

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"retro-carnage.net/engine/geometry"
	"retro-carnage.net/ui/util"
)

const screenTimeout = 5000

type Screen struct {
	screenChangeRequired util.ScreenChangeCallback
	screenChangeTimeout  int64
	textDimensions       map[string]*geometry.Point
	Window               *pixelgl.Window
}

func (s *Screen) SetUp(screenChangeRequired util.ScreenChangeCallback) {
	s.screenChangeRequired = screenChangeRequired
	s.screenChangeTimeout = 0
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.screenChangeTimeout += elapsedTimeInMs
	if s.screenChangeTimeout >= screenTimeout {
		// s.screenChangeRequired(util.Title)
		// util.Info.Println("Moving forward to config screen")
	}
	s.Window.Clear(colornames.Blue)
}

func (s *Screen) TearDown() {

}

func (s *Screen) String() string {
	return string(util.Title)
}
