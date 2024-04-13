package common

import (
	"retro-carnage/engine/input"

	"github.com/faiface/pixel/pixelgl"
)

type ScreenChangeCallback = func(screenName ScreenName)

// Screen is a common interface shared by all detail screens. ui.MainScreen uses this interface to manage the life-cycle
// of it's detail screens.
type Screen interface {
	/*--- Life cycle methods -----------------------------------------------------------------------------------------*/

	// SetUp initializes the Screen. This method gets called once before the Screen gets shown.
	SetUp()
	// Update gets called once during each rendering cycle. It can be used to draw the content of the Screen.
	Update(elapsedTimeInMs int64)
	// TearDown can be used as a life-cycle hook to release resources that a Screen blocked. It will be called once
	// after the last Update.
	TearDown()

	/*--- Initializers for shared properties -------------------------------------------------------------------------*/

	// SetInputController passes the input controller to the screen.
	SetInputController(controller input.InputController)
	// SetScreenChangeCallback passes a callback function that cann be called to switch to another screen.
	SetScreenChangeCallback(callback ScreenChangeCallback)
	// SetWindow passes the application window to the Screen.
	SetWindow(window *pixelgl.Window)

	/*--- Other ------------------------------------------------------------------------------------------------------*/

	// String should return the ScreenName of the Screen
	String() string
}

type ScreenName string

const (
	Loading              ScreenName = "Loading screen"
	Start                ScreenName = "Start screen"
	Title                ScreenName = "Title screen"
	ConfigurationSelect  ScreenName = "Configuration select screen"
	ConfigurationResult  ScreenName = "Configuration result screen"
	ConfigurationOptions ScreenName = "Configuration options screen"
	Mission              ScreenName = "Mission screen"
	BuyYourWeaponsP1     ScreenName = "Buy your items screen (Player 1)"
	BuyYourWeaponsP2     ScreenName = "Buy your items screen (Player 2)"
	ShopP1               ScreenName = "Shop screen (Player 1)"
	ShopP2               ScreenName = "Shop screen (Player 2)"
	LetTheMissionBegin   ScreenName = "Let the mission begin screen"
	Game                 ScreenName = "Game screen"
	EnterNameP1          ScreenName = "Enter name screen (Player 1)"
	EnterNameP2          ScreenName = "Enter name screen (Player 2)"
	HighScore            ScreenName = "High score table screen"
)
