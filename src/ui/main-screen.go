package ui

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"retro-carnage.net/engine/input"
	"retro-carnage.net/ui/loading"
	"retro-carnage.net/ui/start"
	"retro-carnage.net/ui/title"
	"retro-carnage.net/ui/util"
	"time"
)

type MainScreen struct {
	clientScreen util.Screen
	lastUpdate   time.Time
	nextScreen   util.Screen
	inputCtrl    *input.Controller
	Monitor      *pixelgl.Monitor
	Window       *pixelgl.Window
}

func (ms *MainScreen) Initialize() {
	ms.inputCtrl = &input.Controller{Window: ms.Window}
	ms.inputCtrl.HasTwoOrMoreDevices()
	ms.inputCtrl.AssignControllersToPlayers()

	ms.clientScreen = &loading.Screen{Window: ms.Window}
	ms.clientScreen.SetUp(ms.requireScreenChange)

	ms.lastUpdate = time.Now()
}

func (ms *MainScreen) requireScreenChange(screenName util.ScreenName) {
	switch screenName {
	case util.Loading:
		ms.nextScreen = &loading.Screen{Window: ms.Window}
	case util.Start:
		ms.nextScreen = &start.Screen{Window: ms.Window}
	case util.Title:
		ms.nextScreen = &title.Screen{Window: ms.Window}
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
			ms.clientScreen.SetUp(ms.requireScreenChange)
			ms.nextScreen = nil
		}
	}
}
