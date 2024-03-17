package input

import "github.com/faiface/pixel/pixelgl"

type keyboard struct {
	Window *pixelgl.Window
}

func (k *keyboard) State() *DeviceState {
	var result DeviceState
	result.PrimaryAction = k.Window.Pressed(pixelgl.KeyLeftControl)
	result.SecondaryAction = k.Window.Pressed(pixelgl.KeyLeftAlt)
	result.MoveLeft = k.Window.Pressed(pixelgl.KeyLeft)
	result.MoveUp = k.Window.Pressed(pixelgl.KeyUp)
	result.MoveRight = k.Window.Pressed(pixelgl.KeyRight)
	result.MoveDown = k.Window.Pressed(pixelgl.KeyDown)
	result.ToggleUp = k.Window.Pressed(pixelgl.KeyA)
	result.ToggleDown = k.Window.Pressed(pixelgl.KeyZ)
	return &result
}

func (k *keyboard) Name() string {
	return "Keyboard"
}
