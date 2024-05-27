package cheat

import (
	"testing"

	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/stretchr/testify/assert"
)

func TestWrongCheat(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []opengl.Button{
		opengl.KeyM, opengl.KeyO, opengl.KeyR, opengl.KeyK,
	}
	for _, btn := range buttons {
		var active = cheatController.HandleKeyboardInput(btn)
		assert.Equal(t, false, active)
	}
}

func TestValidCheat(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []opengl.Button{
		opengl.KeyJ,
		opengl.KeyO,
		opengl.KeyH,
		opengl.KeyN,
		opengl.KeyR,
		opengl.KeyA,
		opengl.KeyM,
		opengl.KeyB,
	}
	for _, btn := range buttons {
		var active = cheatController.HandleKeyboardInput(btn)
		assert.Equal(t, false, active)
	}

	var active = cheatController.HandleKeyboardInput(opengl.KeyO)
	assert.Equal(t, true, active)
	assert.Equal(t, true, cheatController.unlimitedAmmunition)
}

func TestValidCheatAmidstOtherInput(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []opengl.Button{
		opengl.KeyA,
		opengl.KeyJ,
		opengl.KeyO,
		opengl.KeyH,
		opengl.KeyN,
		opengl.KeyR,
		opengl.KeyA,
		opengl.KeyM,
		opengl.KeyB,
		opengl.KeyO,
		opengl.KeyX,
	}

	for _, btn := range buttons {
		cheatController.HandleKeyboardInput(btn)
	}

	assert.Equal(t, true, cheatController.unlimitedAmmunition)
}
