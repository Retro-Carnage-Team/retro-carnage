package start

import (
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

var (
	lines = []string{
		"The year is 2030.",
		"",
		"Technological progress enables humanity to live in peace and harmony.",
		"A failed AI experiment gives rise to Crulgon, an all-powerful robot ruler.",
		"He quickly gains control of a vast army of highly developed war robots.",
		"With an iron grip and unstoppable precision, his robot army overruns the",
		"world, destroying entire cities and erecting robot fortresses. Humans are",
		"enslaved and the freedom of the Earth is at stake.",
		"",
		"It seems that resistance to this mechanical nightmare is in vain. But deep",
		"in the underground bunkers of a last human fortress, a spark of hope is growing.",
		"You and your team of fearless resistance fighters are the only chance to",
		"liberate the Earth.",
		"",
		"You must face Crulgon and his army, destroy his machines and free humanity",
		"from the clutches of the mechanical tyrant.",
		"",
		"Your battle begins now.",
	}
)

type Screen struct {
	controller *controller
	window     *opengl.Window
}

func NewScreen() *Screen {
	var result = Screen{
		controller: newController(),
	}
	return &result
}

func (s *Screen) SetInputController(i input.InputController) {
	s.controller.setInputController(i)
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	// No set up action required
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.controller.update(elapsedTimeInMs)
	s.drawScreen()
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.Start)
}

func (s *Screen) drawScreen() {
	var smallFontSize = fonts.DefaultFontSize() - 8
	var atlas = fonts.SizeToFontAtlas[smallFontSize]
	txt := text.New(pixel.V(50, 500), atlas)
	txt.Color = common.White

	for _, s := range lines {
		txt.WriteString(s)
		txt.WriteRune('\n')
	}

	var posX = (s.window.Bounds().W() - txt.Bounds().W()) / 2
	var posY = (s.window.Bounds().H() - txt.Bounds().H()) / 2
	txt.Orig = pixel.V(posX, posY)

	txt.Draw(s.window, pixel.IM)
}
