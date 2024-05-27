package cheat

import (
	"testing"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/stretchr/testify/assert"
)

func TestWrongCheat(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []pixel.Button{
		pixel.KeyM, pixel.KeyO, pixel.KeyR, pixel.KeyK,
	}
	for _, btn := range buttons {
		var active = cheatController.HandleKeyboardInput(btn)
		assert.Equal(t, false, active)
	}
}

func TestValidCheat(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []pixel.Button{
		pixel.KeyJ,
		pixel.KeyO,
		pixel.KeyH,
		pixel.KeyN,
		pixel.KeyR,
		pixel.KeyA,
		pixel.KeyM,
		pixel.KeyB,
	}
	for _, btn := range buttons {
		var active = cheatController.HandleKeyboardInput(btn)
		assert.Equal(t, false, active)
	}

	var active = cheatController.HandleKeyboardInput(pixel.KeyO)
	assert.Equal(t, true, active)
	assert.Equal(t, true, cheatController.unlimitedAmmunition)
}

func TestValidCheatAmidstOtherInput(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []pixel.Button{
		pixel.KeyA,
		pixel.KeyJ,
		pixel.KeyO,
		pixel.KeyH,
		pixel.KeyN,
		pixel.KeyR,
		pixel.KeyA,
		pixel.KeyM,
		pixel.KeyB,
		pixel.KeyO,
		pixel.KeyX,
	}

	for _, btn := range buttons {
		cheatController.HandleKeyboardInput(btn)
	}

	assert.Equal(t, true, cheatController.unlimitedAmmunition)
}
