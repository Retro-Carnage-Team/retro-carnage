package enter_name

import (
	"fmt"
	"retro-carnage/assets"
	"retro-carnage/input"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const headlineTemplate = "ENTER YOUR NAME (PLAYER %d)"

// Screen is where the players can enter their names when they reached a new high score.
type Screen struct {
	controller *controller
	model      *model
	stereo     *assets.Stereo
	window     *opengl.Window
}

func NewScreen(playerIdx int) *Screen {
	var model = &model{
		playerIdx: playerIdx,
	}
	var controller = newController(model)
	var result = Screen{
		controller: controller,
		model:      model,
	}
	return &result
}

// SetInputController passes the input controller to the screen.
func (s *Screen) SetInputController(inputCtrl input.InputController) {
	s.controller.setInputController(inputCtrl)
}

// SetScreenChangeCallback passes a callback function that cann be called to switch to another screen.
func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

// SetWindow passes the application window to the Screen.
func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

// SetUp initializes the Screen.
// This method gets called once before the Screen gets shown.
func (s *Screen) SetUp() {
	s.stereo = assets.NewStereo()

}

// Update gets called once during each rendering cycle.
// It can be used to draw the content of the Screen.
func (s *Screen) Update(elapsedTimeInMs int64) {
	s.controller.update(
		elapsedTimeInMs,
		s.window.Typed(),
		s.window.JustPressed(pixel.KeyEnter),
		s.window.JustPressed(pixel.KeyBackspace))

	var renderer = fonts.TextRenderer{Window: s.window}
	renderer.DrawLineToScreenCenter(fmt.Sprintf(headlineTemplate, s.model.playerIdx+1), 2.0, common.Green)
	renderer.DrawLineToScreenCenter(s.model.playerNameDisplay, 0.0, common.White)
}

// TearDown can be used as a life-cycle hook to release resources that a Screen blocked.
// It will be called once after the last Update.
func (s *Screen) TearDown() {
	// No tear down action required
}

// String should return the ScreenName of the Screen
func (s *Screen) String() string {
	if s.model.playerIdx == 1 {
		return string(common.EnterNameP1)
	}
	return string(common.EnterNameP2)
}
