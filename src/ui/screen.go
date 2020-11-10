package ui

type Screen interface {
	SetUp()
	Update()
	TearDown()
}
