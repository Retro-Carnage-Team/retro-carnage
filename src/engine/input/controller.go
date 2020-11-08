package input

type Controller interface {
	GetInputState() *State
}
