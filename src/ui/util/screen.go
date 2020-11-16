package util

type ScreenChangeCallback = func(screenName ScreenName)

type Screen interface {
	SetUp(screenChangeRequired ScreenChangeCallback)
	Update(elapsedTimeInMs int64)
	String() string
	TearDown()
}

type ScreenName string

const (
	Loading ScreenName = "Loading Screen"
	Start   ScreenName = "Start Screen"
	Title   ScreenName = "Title Screen"
)
