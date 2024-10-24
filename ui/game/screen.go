package game

import (
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
	"time"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const pausedMessage = "GAME PAUSED"

// Screen in this package is the one that show the actual gameplay.
type Screen struct {
	controller  *controller
	fpsInfo     *fpsInfo
	model       *model
	playerInfos []*playerInfo
	renderer    *engine.Renderer
	window      *opengl.Window
}

func NewScreen() *Screen {
	var model = model{
		mission: engine.MissionController.CurrentMission(),
		paused:  false,
	}

	if nil == model.mission {
		logging.Error.Fatalf("No mission selected on game screen")
	}

	var controller = controller{
		engine: engine.NewGameEngine(model.mission),
		model:  &model,
		stereo: assets.NewStereo(),
	}
	var result = Screen{
		controller: &controller,
		model:      &model,
	}
	return &result
}

// SetInputController is used to connect Screen with the global input.Controller instance.
func (s *Screen) SetInputController(ctrl input.InputController) {
	s.controller.setInputController(ctrl)
}

// SetScreenChangeCallback is used to connect Screen with the callback method of ui.MainScreen.
func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.controller.setScreenChangeCallback(callback)
}

// SetWindow is used to connect Screen with the opengl.Window instance.
func (s *Screen) SetWindow(window *opengl.Window) {
	s.window = window
}

// SetUp is called when the screen got initialized by ui.MainScreen and is about to appear shortly.
func (s *Screen) SetUp() {
	s.controller.setUp()

	s.fpsInfo = &fpsInfo{second: time.NewTicker(time.Second).C}
	s.playerInfos = []*playerInfo{
		newPlayerInfo(0, s.window),
		newPlayerInfo(1, s.window),
	}
	s.renderer = engine.NewRenderer(s.controller.engine, s.window)
}

// Update gets called for every frame that gets displayed.
// Here we update the state of the gameplay based on the time that has elapsed since the last frame.
// Then we render the new game state to the opengl.Window.
func (s *Screen) Update(elapsedTimeInMs int64) {
	if nil == s.controller.engine || nil == s.renderer {
		return
	}

	s.controller.update(elapsedTimeInMs, s.window)

	s.drawScreen(elapsedTimeInMs)

	s.fpsInfo.update()
	s.fpsInfo.drawToScreen(s.window)
}

func (s *Screen) drawScreen(elapsedTimeInMs int64) {
	for _, playerInfo := range s.playerInfos {
		playerInfo.draw(s.window)
	}

	if s.model.inProgress {
		if s.model.paused {
			var gameCanvas = s.renderer.Render(0)
			var matrix = pixel.IM.Moved(s.window.Bounds().Center())
			gameCanvas.DrawColorMask(s.window, matrix, pixel.RGBA{A: 0.5})

			var renderer = fonts.TextRenderer{Window: s.window}
			renderer.DrawLineToScreenCenter(pausedMessage, 0, common.White)
		} else {
			s.renderer.Render(elapsedTimeInMs)
		}
	} else if s.model.won {
		var gameCanvas = s.renderer.Render(0)
		if nil == s.controller.gameWonAnimation && nil == s.controller.missionWonAnimation {
			s.controller.createGameAnimations(s.playerInfos, gameCanvas, s.window)
		}

		if nil != s.controller.gameWonAnimation {
			s.controller.gameWonAnimation.drawToScreen()
		} else if nil != s.controller.missionWonAnimation {
			s.controller.missionWonAnimation.drawToScreen()
		}
	} else if s.model.lost {
		var gameCanvas = s.renderer.Render(0)
		if nil == s.controller.gameLostAnimation {
			s.controller.createGameAnimations(s.playerInfos, gameCanvas, s.window)
		}
		s.controller.gameLostAnimation.drawToScreen()
	}
}

// TearDown is called by ui.MainWindow when the Screen has been displayed for the last time.
// Here we clean up and free used resources.
func (s *Screen) TearDown() {
	for _, playerInfo := range s.playerInfos {
		playerInfo.dispose()
	}
	s.controller.tearDown(s.window)
}

// String returns the screen name as string
func (s *Screen) String() string {
	return string(common.Game)
}
