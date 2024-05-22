package input

type inputDevice interface {
	State() *InputDeviceState
	Name() string
}
