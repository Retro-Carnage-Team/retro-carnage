package ui

type Screen interface {
	SetUp()
	Update(elapsedTimeInMs int64)
	TearDown()
}
