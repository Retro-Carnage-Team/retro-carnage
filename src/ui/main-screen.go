package ui

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"retro-carnage.net/engine/input"
	"retro-carnage.net/ui/common"
	"retro-carnage.net/ui/config"
	"retro-carnage.net/ui/loading"
	"retro-carnage.net/ui/start"
	"retro-carnage.net/ui/title"
	"time"
)

type MainScreen struct {
	clientScreen common.Screen
	lastUpdate   time.Time
	nextScreen   common.Screen
	inputCtrl    *input.Controller
	Monitor      *pixelgl.Monitor
	Window       *pixelgl.Window
}

func (ms *MainScreen) Initialize() {
	ms.inputCtrl = &input.Controller{Window: ms.Window}
	ms.inputCtrl.HasTwoOrMoreDevices()
	ms.inputCtrl.AssignControllersToPlayers()

	ms.clientScreen = &loading.Screen{}
	ms.setUpScreen(ms.clientScreen)

	ms.lastUpdate = time.Now()
}

func (ms *MainScreen) requireScreenChange(screenName common.ScreenName) {
	switch screenName {
	case common.Loading:
		ms.nextScreen = &loading.Screen{}
	case common.Start:
		ms.nextScreen = &start.Screen{}
	case common.Title:
		ms.nextScreen = &title.Screen{}
	case common.Configuration:
		ms.nextScreen = &config.Screen{}
	}
}

func (ms *MainScreen) RunMainLoop() {
	for !ms.Window.Closed() {
		duration := time.Since(ms.lastUpdate).Milliseconds()
		ms.lastUpdate = time.Now()

		ms.Window.Update()
		ms.Window.Clear(colornames.Black)
		ms.clientScreen.Update(duration)

		if nil != ms.nextScreen {
			ms.clientScreen.TearDown()
			ms.clientScreen = ms.nextScreen
			ms.setUpScreen(ms.clientScreen)
			ms.nextScreen = nil
		}
	}
}

func (ms *MainScreen) setUpScreen(aScreen common.Screen) {
	aScreen.SetInputController(ms.inputCtrl)
	aScreen.SetScreenChangeCallback(ms.requireScreenChange)
	aScreen.SetWindow(ms.Window)
	aScreen.SetUp()
}
