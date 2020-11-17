package util

type ScreenChangeCallback = func(screenName ScreenName)

// Screen is a common interface shared by all detail screens. ui.MainScreen uses this interface to manage the life-cycle
// of it's detail screens.
type Screen interface {
	// SetUp initializes the Screen. This method gets called once before the Screen gets shown.
	SetUp(screenChangeRequired ScreenChangeCallback)
	// Update gets called once during each rendering cycle. It can be used to draw the content of the Screen.
	Update(elapsedTimeInMs int64)
	// Should return the ScreenName of the Screen
	String() string
	// TearDown can be used as a life-cycle hook to release resources that a Screen blocked. It will be called once
	// after the last Update.
	TearDown()
}

type ScreenName string

const (
	Loading ScreenName = "Loading Screen"
	Start   ScreenName = "Start Screen"
	Title   ScreenName = "Title Screen"
)
