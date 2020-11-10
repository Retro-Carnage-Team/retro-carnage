package ui

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"retro-carnage.net/engine/input"
	"retro-carnage.net/ui/loading"
)

type MainScreen struct {
	clientScreen Screen
	nextScreen   Screen
	inputCtrl    *input.Controller
	Monitor      *pixelgl.Monitor
	Window       *pixelgl.Window
}

func (ms *MainScreen) Initialize() {
	ms.inputCtrl = &input.Controller{Window: ms.Window}
	ms.inputCtrl.HasTwoOrMoreDevices()
	ms.inputCtrl.AssignControllersToPlayers()

	ms.clientScreen = &loading.Screen{Window: ms.Window}
	ms.clientScreen.SetUp()
}

func (ms *MainScreen) RunMainLoop() {
	for !ms.Window.Closed() {
		ms.Window.Update()
		ms.Window.Clear(colornames.Black)
		ms.clientScreen.Update()

		if nil != ms.nextScreen {
			ms.clientScreen.TearDown()
			ms.clientScreen = ms.nextScreen
			ms.clientScreen.SetUp()
			ms.nextScreen = nil
		}
	}
}
