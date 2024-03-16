package enter_name

import (
	"fmt"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
	"retro-carnage/ui/highscore"

	"github.com/faiface/pixel/pixelgl"
)

// Screen is where the players can enter their names when they reached a new high score.
type Screen struct {
	cursorVisible        bool
	duration             int64
	inputController      input.Controller
	PlayerIdx            int
	playerName           string
	screenChangeRequired common.ScreenChangeCallback
	stereo               *assets.Stereo
	window               *pixelgl.Window
}

const (
	headlineTemplate = "ENTER YOUR NAME (PLAYER %d)"
	maxLengthOfName  = 10
)

// SetInputController passes the input controller to the screen.
func (s *Screen) SetInputController(inputCtrl input.Controller) {
	s.inputController = inputCtrl
}

// SetScreenChangeCallback passes a callback function that cann be called to switch to another screen.
func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

// SetWindow passes the application window to the Screen.
func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

// SetUp initializes the Screen.
// This method gets called once before the Screen gets shown.
func (s *Screen) SetUp() {
	s.stereo = assets.NewStereo()

	var name, error = highscore.EntryControllerInstance.PlayerName(s.PlayerIdx)
	if nil != error {
		logging.Error.Printf("Failed to get player name: %v", error)
	} else {
		s.playerName = name
	}
}

// Update gets called once during each rendering cycle.
// It can be used to draw the content of the Screen.
func (s *Screen) Update(elapsedTimeInMs int64) {
	s.duration += elapsedTimeInMs
	if s.duration > 250 {
		s.cursorVisible = !s.cursorVisible
		s.duration = 0
	}

	s.playerName = s.playerName + s.window.Typed()
	if maxLengthOfName < len(s.playerName) {
		s.playerName = s.playerName[:maxLengthOfName]
	}

	var playerName = s.playerName
	if s.window.JustPressed(pixelgl.KeyEnter) || s.inputController.ControllerUiEventStateCombined().PressedButton {
		highscore.EntryControllerInstance.SetPlayerName(s.PlayerIdx, s.playerName)
		highscore.EntryControllerInstance.AddEntry(highscore.Entry{
			Name:  s.playerName,
			Score: characters.PlayerController.ConfiguredPlayers()[s.PlayerIdx].Score(),
		})
		s.exit()
	} else {
		if s.window.JustPressed(pixelgl.KeyBackspace) && (0 < len(s.playerName)) {
			s.playerName = s.playerName[:len(s.playerName)-1]
			playerName = s.playerName
		}

		if s.cursorVisible {
			playerName = playerName + "|"
		} else {
			playerName = playerName + " "
		}
		var renderer = fonts.TextRenderer{Window: s.window}
		renderer.DrawLineToScreenCenter(fmt.Sprintf(headlineTemplate, s.PlayerIdx+1), 2.0, common.Green)
		renderer.DrawLineToScreenCenter(playerName, 0.0, common.White)
	}
}

// TearDown can be used as a life-cycle hook to release resources that a Screen blocked.
// It will be called once after the last Update.
func (s *Screen) TearDown() {
	// No tear down action required
}

// String should return the ScreenName of the Screen
func (s *Screen) String() string {
	if s.PlayerIdx == 1 {
		return string(common.EnterNameP1)
	}
	return string(common.EnterNameP2)
}

func (s *Screen) exit() {
	if s.PlayerIdx == 0 {
		var _, p2 = highscore.EntryControllerInstance.ReachedHighScore()
		if p2 {
			s.screenChangeRequired(common.EnterNameP2)
		} else {
			s.screenChangeRequired(common.HighScore)
		}
	} else {
		s.screenChangeRequired(common.HighScore)
	}
}
