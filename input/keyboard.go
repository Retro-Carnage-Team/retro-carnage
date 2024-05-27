package input

import (
	"retro-carnage/config"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

type keyboard struct {
	configuration config.InputDeviceConfiguration
	window        *opengl.Window
}

func (k *keyboard) State() *InputDeviceState {
	var result = InputDeviceState{
		PrimaryAction: k.window.Pressed(pixel.Button(k.configuration.InputFire)),
		MoveLeft:      k.window.Pressed(pixel.Button(k.configuration.InputLeft)),
		MoveUp:        k.window.Pressed(pixel.Button(k.configuration.InputUp)),
		MoveRight:     k.window.Pressed(pixel.Button(k.configuration.InputRight)),
		MoveDown:      k.window.Pressed(pixel.Button(k.configuration.InputDown)),
		ToggleUp:      k.window.Pressed(pixel.Button(k.configuration.InputNextWeapon)),
		ToggleDown:    k.window.Pressed(pixel.Button(k.configuration.InputPreviousWeapon)),
	}
	return &result
}

func (k *keyboard) Name() string {
	return "Keyboard"
}
