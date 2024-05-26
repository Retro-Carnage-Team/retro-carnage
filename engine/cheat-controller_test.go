package engine

import (
	"testing"

	"github.com/Retro-Carnage-Team/pixel/pixelgl"
	"github.com/stretchr/testify/assert"
)

func TestWrongCheat(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []pixelgl.Button{
		pixelgl.KeyM, pixelgl.KeyO, pixelgl.KeyR, pixelgl.KeyK,
	}
	for _, btn := range buttons {
		var active = cheatController.HandleKeyboardInput(btn)
		assert.Equal(t, false, active)
	}
}

func TestValidCheat(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []pixelgl.Button{
		pixelgl.KeyJ,
		pixelgl.KeyO,
		pixelgl.KeyH,
		pixelgl.KeyN,
		pixelgl.KeyR,
		pixelgl.KeyA,
		pixelgl.KeyM,
		pixelgl.KeyB,
	}
	for _, btn := range buttons {
		var active = cheatController.HandleKeyboardInput(btn)
		assert.Equal(t, false, active)
	}

	var active = cheatController.HandleKeyboardInput(pixelgl.KeyO)
	assert.Equal(t, true, active)
	assert.Equal(t, true, cheatController.unlimitedAmmunition)
}

func TestValidCheatAmidstOtherInput(t *testing.T) {
	var cheatController = GetCheatController()
	cheatController.Reset()

	var buttons = []pixelgl.Button{
		pixelgl.KeyA,
		pixelgl.KeyJ,
		pixelgl.KeyO,
		pixelgl.KeyH,
		pixelgl.KeyN,
		pixelgl.KeyR,
		pixelgl.KeyA,
		pixelgl.KeyM,
		pixelgl.KeyB,
		pixelgl.KeyO,
		pixelgl.KeyX,
	}

	for _, btn := range buttons {
		cheatController.HandleKeyboardInput(btn)
	}

	assert.Equal(t, true, cheatController.unlimitedAmmunition)
}
