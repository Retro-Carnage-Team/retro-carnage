package cheat

import (
	"retro-carnage/logging"

	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

const (
	maxCheatLen = 11
)

var (
	cheatController          *CheatController
	unlimitedAmmunitionCheat = []opengl.Button{
		opengl.KeyJ,
		opengl.KeyO,
		opengl.KeyH,
		opengl.KeyN,
		opengl.KeyR,
		opengl.KeyA,
		opengl.KeyM,
		opengl.KeyB,
		opengl.KeyO,
	}
	unlimitedLivesCheat = []opengl.Button{
		opengl.KeyD,
		opengl.KeyU,
		opengl.KeyN,
		opengl.KeyC,
		opengl.KeyA,
		opengl.KeyN,
		opengl.KeyI,
		opengl.KeyD,
		opengl.KeyA,
		opengl.KeyH,
		opengl.KeyO,
	}
)

type CheatController struct {
	input               []opengl.Button
	unlimitedAmmunition bool
	unlimitedLives      bool
}

func GetCheatController() *CheatController {
	if cheatController == nil {
		cheatController = &CheatController{
			unlimitedAmmunition: false,
			unlimitedLives:      false,
		}
	}
	return cheatController
}

func (cc *CheatController) IsAmmunitionUnlimited() bool {
	return cc.unlimitedAmmunition
}

func (cc *CheatController) IsNumberOfLivesUnlimited() bool {
	return cc.unlimitedLives
}

func (cc *CheatController) HandleKeyboardInput(button opengl.Button) bool {
	var prevInput = cc.input
	if len(cc.input) == maxCheatLen {
		prevInput = cc.input[1:]
	}
	cc.input = append(prevInput, button)

	if cc.compareInputToCheat(unlimitedAmmunitionCheat) {
		cc.unlimitedAmmunition = !cc.unlimitedAmmunition
		cc.logCheatActivation("ammo", cc.unlimitedAmmunition)
		return true
	}

	if cc.compareInputToCheat(unlimitedLivesCheat) {
		cc.unlimitedLives = !cc.unlimitedLives
		cc.logCheatActivation("lives", cc.unlimitedLives)
		return true
	}

	return false
}

func (cc *CheatController) Reset() {
	cc.input = []opengl.Button{}
	cc.unlimitedAmmunition = false
	cc.unlimitedLives = false
}

func (cc *CheatController) compareInputToCheat(cheat []opengl.Button) bool {
	if len(cc.input) < len(cheat) {
		return false
	}

	for i, c := range cheat {
		if cc.input[len(cc.input)-len(cheat)+i] != c {
			return false
		}
	}

	return true
}

func (cc *CheatController) logCheatActivation(topic string, status bool) {
	var statusString = "OFF"
	if status {
		statusString = "ON"
	}
	logging.Info.Printf("Cheat for unlimited %s switched %s", topic, statusString)
}
