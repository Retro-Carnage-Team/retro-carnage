package ui

import (
	"retro-carnage/input"
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

	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"golang.org/x/image/colornames"
)

type MainScreen struct {
	clientScreen common.Screen
	inputCtrl    input.InputController
	lastUpdate   time.Time
	Monitor      *opengl.Monitor
	nextScreen   common.Screen
	Window       *opengl.Window
}

func (ms *MainScreen) Initialize() {
	ms.inputCtrl = input.NewController(ms.Window)
	ms.inputCtrl.AssignInputDevicesToPlayers()

	ms.clientScreen = loading.NewScreen()
	ms.setUpScreen()

	ms.lastUpdate = time.Now()
}

func (ms *MainScreen) requireScreenChange(screenName common.ScreenName) {
	logging.Info.Printf("Changing screen to %s", screenName)

	switch screenName {
	case common.Loading:
		ms.nextScreen = loading.NewScreen()
	case common.Start:
		ms.nextScreen = start.NewScreen()
	case common.Title:
		ms.nextScreen = title.NewScreen()
	case common.ConfigurationOptions:
		ms.nextScreen = &config.OptionsScreen{}
	case common.ConfigurationAudio:
		ms.nextScreen = config.NewAudioOptionsScreen()
	case common.ConfigurationVideo:
		ms.nextScreen = config.NewVideoOptionsScreen()
	case common.ConfigurationControls:
		ms.nextScreen = config.NewInputOptionsScreen()
	case common.ConfigurationControlsP1:
		ms.nextScreen = config.NewControllerOptionsScreen(0)
	case common.ConfigurationControlsP2:
		ms.nextScreen = config.NewControllerOptionsScreen(1)
	case common.ConfigurationResult:
		ms.nextScreen = &config.ResultScreen{}
	case common.ConfigurationSelect:
		ms.nextScreen = config.NewSelectScreen()
	case common.Mission:
		ms.nextScreen = mission.NewScreen()
	case common.BuyYourWeaponsP1:
		ms.nextScreen = byw.NewScreen(0)
	case common.BuyYourWeaponsP2:
		ms.nextScreen = byw.NewScreen(1)
	case common.ShopP1:
		ms.nextScreen = shop.NewScreen(0)
	case common.ShopP2:
		ms.nextScreen = shop.NewScreen(1)
	case common.LetTheMissionBegin:
		ms.nextScreen = lb.NewScreen()
	case common.Game:
		ms.nextScreen = &game.Screen{}
	case common.EnterNameP1:
		ms.nextScreen = en.NewScreen(0)
	case common.EnterNameP2:
		ms.nextScreen = en.NewScreen(1)
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
			ms.setUpScreen()
			ms.nextScreen = nil
		}
	}
}

func (ms *MainScreen) setUpScreen() {
	ms.clientScreen.SetInputController(ms.inputCtrl)
	ms.clientScreen.SetScreenChangeCallback(ms.requireScreenChange)
	ms.clientScreen.SetWindow(ms.Window)
	ms.clientScreen.SetUp()
}
