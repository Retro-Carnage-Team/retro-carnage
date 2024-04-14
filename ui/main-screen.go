package ui

import (
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	byw "retro-carnage/ui/buy-your-weapons"
	"retro-carnage/ui/common"
	"retro-carnage/ui/config"
	en "retro-carnage/ui/enter-name"
	"retro-carnage/ui/game"
	"retro-carnage/ui/highscore"
	lb "retro-carnage/ui/lets-begin"
	"retro-carnage/ui/loading"
	"retro-carnage/ui/mission"
	"retro-carnage/ui/shop"
	"retro-carnage/ui/start"
	"retro-carnage/ui/title"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type MainScreen struct {
	clientScreen common.Screen
	inputCtrl    input.InputController
	lastUpdate   time.Time
	Monitor      *pixelgl.Monitor
	nextScreen   common.Screen
	Window       *pixelgl.Window
}

func (ms *MainScreen) Initialize() {
	ms.inputCtrl = input.NewController(ms.Window)
	ms.inputCtrl.AssignControllersToPlayers()

	ms.clientScreen = &loading.Screen{}
	ms.setUpScreen(ms.clientScreen)

	ms.lastUpdate = time.Now()
}

func (ms *MainScreen) requireScreenChange(screenName common.ScreenName) {
	logging.Info.Printf("Changing screen to %s", screenName)

	switch screenName {
	case common.Loading:
		ms.nextScreen = &loading.Screen{}
	case common.Start:
		ms.nextScreen = &start.Screen{}
	case common.Title:
		ms.nextScreen = &title.Screen{}
	case common.ConfigurationOptions:
		ms.nextScreen = &config.OptionsScreen{}
	case common.ConfigurationResult:
		ms.nextScreen = &config.ResultScreen{}
	case common.ConfigurationSelect:
		ms.nextScreen = &config.SelectScreen{}
	case common.Mission:
		ms.nextScreen = &mission.Screen{}
	case common.BuyYourWeaponsP1:
		ms.nextScreen = &byw.Screen{PlayerIdx: 0}
	case common.BuyYourWeaponsP2:
		ms.nextScreen = &byw.Screen{PlayerIdx: 1}
	case common.ShopP1:
		ms.nextScreen = &shop.Screen{PlayerIdx: 0}
	case common.ShopP2:
		ms.nextScreen = &shop.Screen{PlayerIdx: 1}
	case common.LetTheMissionBegin:
		ms.nextScreen = &lb.Screen{}
	case common.Game:
		ms.nextScreen = &game.Screen{}
	case common.EnterNameP1:
		ms.nextScreen = &en.Screen{PlayerIdx: 0}
	case common.EnterNameP2:
		ms.nextScreen = &en.Screen{PlayerIdx: 1}
	case common.HighScore:
		ms.nextScreen = &highscore.Screen{}
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
