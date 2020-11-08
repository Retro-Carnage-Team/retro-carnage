package input

import "github.com/faiface/pixel/pixelgl"

type KeyboardController struct {
	Window *pixelgl.Window
}

func (kc *KeyboardController) GetInputState() *State {
	var result State
	result.Fire = kc.Window.Pressed(pixelgl.KeyLeftControl)
	result.Grenade = kc.Window.Pressed(pixelgl.KeyLeftAlt)
	result.MoveLeft = kc.Window.Pressed(pixelgl.KeyLeft)
	result.MoveUp = kc.Window.Pressed(pixelgl.KeyUp)
	result.MoveRight = kc.Window.Pressed(pixelgl.KeyRight)
	result.MoveDown = kc.Window.Pressed(pixelgl.KeyDown)
	result.ToggleUp = kc.Window.Pressed(pixelgl.KeyA)
	result.ToggleDown = kc.Window.Pressed(pixelgl.KeyZ)
	return &result
}
