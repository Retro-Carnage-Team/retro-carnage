package engine

import "github.com/Retro-Carnage-Team/pixel/pixelgl"

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
	// TODO:
	// - Add button to buffer
	// - Check buffer
	// - Activate or deactivate cheat
	// - Return true when cheat has been activated
	return false
}
