package input

import (
	"retro-carnage/config"

	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

type keyboard struct {
	configuration config.InputDeviceConfiguration
	window        *opengl.Window
}

func (k *keyboard) State() *InputDeviceState {
	var result = InputDeviceState{
		PrimaryAction: k.window.Pressed(opengl.Button(k.configuration.InputFire)),
		MoveLeft:      k.window.Pressed(opengl.Button(k.configuration.InputLeft)),
		MoveUp:        k.window.Pressed(opengl.Button(k.configuration.InputUp)),
		MoveRight:     k.window.Pressed(opengl.Button(k.configuration.InputRight)),
		MoveDown:      k.window.Pressed(opengl.Button(k.configuration.InputDown)),
		ToggleUp:      k.window.Pressed(opengl.Button(k.configuration.InputNextWeapon)),
		ToggleDown:    k.window.Pressed(opengl.Button(k.configuration.InputPreviousWeapon)),
	}
	return &result
}

func (k *keyboard) Name() string {
	return "Keyboard"
}
