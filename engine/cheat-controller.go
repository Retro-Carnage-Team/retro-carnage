package engine

import (
	"retro-carnage/logging"

	"github.com/Retro-Carnage-Team/pixel/pixelgl"
)

const (
	maxCheatLen = 11
)

var (
	cheatController          *CheatController
	unlimitedAmmunitionCheat = []pixelgl.Button{
		pixelgl.KeyJ,
		pixelgl.KeyO,
		pixelgl.KeyH,
		pixelgl.KeyN,
		pixelgl.KeyR,
		pixelgl.KeyA,
		pixelgl.KeyM,
		pixelgl.KeyB,
		pixelgl.KeyO,
	}
	unlimitedLivesCheat = []pixelgl.Button{
		pixelgl.KeyD,
		pixelgl.KeyU,
		pixelgl.KeyN,
		pixelgl.KeyC,
		pixelgl.KeyA,
		pixelgl.KeyN,
		pixelgl.KeyI,
		pixelgl.KeyD,
		pixelgl.KeyA,
		pixelgl.KeyH,
		pixelgl.KeyO,
	}
	unlimitedMoneyCheat = []pixelgl.Button{
		pixelgl.KeyT,
		pixelgl.KeyO,
		pixelgl.KeyN,
		pixelgl.KeyY,
		pixelgl.KeyS,
		pixelgl.KeyT,
		pixelgl.KeyA,
		pixelgl.KeyR,
		pixelgl.KeyK,
	}
)

type CheatController struct {
	input               []pixelgl.Button
	unlimitedAmmunition bool
	unlimitedLives      bool
	unlimitedMoney      bool
}

func GetCheatController() *CheatController {
	if cheatController == nil {
		cheatController = &CheatController{
			unlimitedAmmunition: false,
			unlimitedLives:      false,
			unlimitedMoney:      false,
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

func (cc *CheatController) IsMonayUnlimited() bool {
	return cc.unlimitedMoney
}

func (cc *CheatController) HandleKeyboardInput(button pixelgl.Button) bool {
	var prevInput = cc.input
	if len(cc.input) == maxCheatLen {
		prevInput = cc.input[1:]
	}
	cc.input = append(prevInput, button)

	var debugOutput = ""
	for _, btn := range cc.input {
		debugOutput = debugOutput + btn.String()
	}
	logging.Info.Printf("Cheat-Input is %s", debugOutput)

	if cc.compareInputToCheat(unlimitedAmmunitionCheat) {
		cc.unlimitedAmmunition = !cc.unlimitedAmmunition
		return true
	}

	if cc.compareInputToCheat(unlimitedLivesCheat) {
		cc.unlimitedLives = !cc.unlimitedLives
		return true
	}

	if cc.compareInputToCheat(unlimitedMoneyCheat) {
		cc.unlimitedMoney = !cc.unlimitedMoney
		return true
	}

	return false
}

func (cc *CheatController) Reset() {
	cc.input = []pixelgl.Button{}
	cc.unlimitedAmmunition = false
	cc.unlimitedLives = false
	cc.unlimitedMoney = false
}

func (cc *CheatController) compareInputToCheat(cheat []pixelgl.Button) bool {
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
