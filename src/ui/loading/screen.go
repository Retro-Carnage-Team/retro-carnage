package loading

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"retro-carnage.net/util"
)

type Screen struct {
	Window *pixelgl.Window
}

func (s *Screen) SetUp() {

}

func (s *Screen) Update() {
	txt := text.New(pixel.V(100, 500), util.DefaultAtlas)
	_, _ = fmt.Fprintln(txt, "Hello, text!")
	txt.Color = colornames.Red
	txt.Draw(s.Window, pixel.IM)
}

func (s *Screen) TearDown() {

}

func (s *Screen) String() string {
	return "Loading Screen"
}
