package input

import (
	"retro-carnage/config"

	"github.com/faiface/pixel/pixelgl"
)

type keyboard struct {
	configuration config.InputDeviceConfiguration
	window        *pixelgl.Window
}

func (k *keyboard) State() *InputDeviceState {
	var result = InputDeviceState{
		PrimaryAction: k.window.Pressed(pixelgl.Button(k.configuration.InputFire)),
		MoveLeft:      k.window.Pressed(pixelgl.Button(k.configuration.InputLeft)),
		MoveUp:        k.window.Pressed(pixelgl.Button(k.configuration.InputUp)),
		MoveRight:     k.window.Pressed(pixelgl.Button(k.configuration.InputRight)),
		MoveDown:      k.window.Pressed(pixelgl.Button(k.configuration.InputDown)),
		ToggleUp:      k.window.Pressed(pixelgl.Button(k.configuration.InputNextWeapon)),
		ToggleDown:    k.window.Pressed(pixelgl.Button(k.configuration.InputPreviousWeapon)),
	}
	return &result
}

func (k *keyboard) Name() string {
	return "Keyboard"
}
