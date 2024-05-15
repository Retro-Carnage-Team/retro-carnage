package config

type focusChange struct {
	movedLeft        bool
	movedRight       bool
	movedDown        bool
	movedUp          bool
	currentSelection []int
	nextSelection    int
}
