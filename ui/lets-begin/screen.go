package lets_begin

import (
	"fmt"
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/text"
)

const (
	displayText             = "LET THE MISSION BEGIN..."
	timeAfterLastChar       = 500
	timeBetweenChars        = 120
	timeBetweenVolumeChange = 150
)

type Screen struct {
	characterTimer       int64
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	text                 string
	textLength           int
	volumeTimer          int64
	window               *opengl.Window
}

func (s *Screen) SetInputController(_ input.InputController) {
	// Screen doesn't process user input. So no implementation required.
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.stereo = assets.NewStereo()
}

func (s *Screen) Update(elapsedTimeInMs int64) {
	s.characterTimer += elapsedTimeInMs
	s.volumeTimer += elapsedTimeInMs
	if s.textLength < len(displayText) {
		// text has not been fully typed
		if s.characterTimer >= timeBetweenChars {
			s.textLength++
			s.text = displayText[:s.textLength]
			s.characterTimer = 0
		}
		if s.volumeTimer >= timeBetweenVolumeChange {
			s.stereo.DecreaseVolume(assets.ThemeSong)
			s.volumeTimer = 0
		}
	} else if s.isMissionInitialized() {
		// text has been typed and initialization is completed
		s.screenChangeRequired(common.Game)
		s.stereo.StopSong(assets.ThemeSong)
	} else {
		// text has been typed - but initialization is not completed
		s.textLength = s.textLength - 3
		s.text = displayText[:s.textLength]
		s.characterTimer = 0
	}
	s.drawText()
}

func (s *Screen) isMissionInitialized() bool {
	var music = engine.MissionController.CurrentMission().Music
	return s.stereo.IsSongBuffered(music)
}

func (s *Screen) TearDown() {
	// No tear down action required
}

func (s *Screen) String() string {
	return string(common.LetTheMissionBegin)
}

func (s *Screen) drawText() {
	var defaultFontSize = fonts.DefaultFontSize()
	var lineDimensions = fonts.GetTextDimension(defaultFontSize, displayText)

	var vertCenter = s.window.Bounds().Max.Y / 2
	var lineX = (s.window.Bounds().Max.X - lineDimensions.X) / 2

	var txt = text.New(pixel.V(lineX, vertCenter), fonts.SizeToFontAtlas[defaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, s.text)
	txt.Draw(s.window, pixel.IM)
}
