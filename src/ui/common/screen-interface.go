package common

import (
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/input"
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

	// SetInputController passes the input controller to the screen
	SetInputController(controller input.Controller)
	// SetScreenChangeCallback passes a callback function that cann be called to switch to another screen
	SetScreenChangeCallback(callback ScreenChangeCallback)
	// SetWindow that displays the game
	SetWindow(window *pixelgl.Window)

	/*--- Other ------------------------------------------------------------------------------------------------------*/

	// Should return the ScreenName of the Screen
	String() string
}

type ScreenName string

const (
	Loading       ScreenName = "Loading Screen"
	Start         ScreenName = "Start Screen"
	Title         ScreenName = "Title Screen"
	Configuration ScreenName = "Configuration Screen"
)
