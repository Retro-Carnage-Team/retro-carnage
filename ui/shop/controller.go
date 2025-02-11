package shop

import (
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/engine/characters"
	"retro-carnage/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
)

type controller struct {
	inputController      input.InputController
	inventoryController  engine.InventoryController
	model                *model
	screenChangeRequired common.ScreenChangeCallback
}

func newController(model *model) *controller {
	var result = controller{
		inventoryController: engine.NewInventoryController(model.playerIdx),
		model:               model,
	}
	return &result
}

func (c *controller) setInputController(controller input.InputController) {
	c.inputController = controller
}

func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) processUserInput() {
	var eventState, err = c.inputController.GetUiEventState(c.model.playerIdx)
	if nil != err {
		logging.Warning.Printf("Failed to get game controller state: %v", err)
	} else if nil != eventState {
		if eventState.MovedDown && !c.model.modalVisible {
			c.processSelectionMovedDown()
		} else if eventState.MovedUp && !c.model.modalVisible {
			c.processSelectionMovedUp()
		} else if eventState.MovedRight {
			c.processSelectionMovedRight()
		} else if eventState.MovedLeft {
			c.processSelectionMovedLeft()
		}
		if eventState.PressedButton {
			c.processButtonPressed()
		}
	}
}

func (c *controller) processSelectionMovedDown() {
	if c.model.selectedItemIdx != -1 {
		if 5 <= c.model.selectedItemIdx/5 {
			c.model.selectedItemIdx = -1
		} else {
			c.model.selectedItemIdx += 5
		}
	} else {
		c.model.selectedItemIdx = 4
	}
}

func (c *controller) processSelectionMovedUp() {
	if c.model.selectedItemIdx != -1 {
		if 5 > c.model.selectedItemIdx {
			c.model.selectedItemIdx = -1
		} else {
			c.model.selectedItemIdx -= 5
		}
	} else {
		c.model.selectedItemIdx = len(c.model.items) - 1
	}
}

func (c *controller) processSelectionMovedRight() {
	if c.model.modalVisible {
		if c.model.modalButtonSelection == buttonBuyWeapon {
			if c.isModalButtonBuyAmmunitionAvailable() {
				c.model.modalButtonSelection = buttonBuyAmmo
			} else {
				c.model.modalButtonSelection = buttonCloseModal
			}
		} else if c.model.modalButtonSelection == buttonBuyAmmo {
			c.model.modalButtonSelection = buttonCloseModal
		}
	} else if c.model.selectedItemIdx != -1 {
		if c.model.selectedItemIdx%5 == 4 {
			c.model.selectedItemIdx -= 4
		} else {
			c.model.selectedItemIdx += 1
		}
	}
}

func (c *controller) processSelectionMovedLeft() {
	if c.model.modalVisible {
		if c.model.modalButtonSelection == buttonBuyAmmo && c.isModalButtonBuyWeaponAvailable() {
			c.model.modalButtonSelection = buttonBuyWeapon
		} else if c.model.modalButtonSelection == buttonCloseModal {
			if c.isModalButtonBuyAmmunitionAvailable() {
				c.model.modalButtonSelection = buttonBuyAmmo
			} else if c.isModalButtonBuyWeaponAvailable() {
				c.model.modalButtonSelection = buttonBuyWeapon
			}
		}
	} else if c.model.selectedItemIdx != -1 {
		if c.model.selectedItemIdx%5 == 0 {
			c.model.selectedItemIdx += 4
		} else {
			c.model.selectedItemIdx -= 1
		}
	}
}

func (c *controller) processButtonPressed() {
	if c.model.modalVisible {
		c.processButtonPressedOnModal()
	} else {
		c.processButtonPressedOnShop()
	}
}

func (c *controller) processButtonPressedOnModal() {
	var item = c.model.items[c.model.selectedItemIdx]
	switch c.model.modalButtonSelection {
	case buttonBuyWeapon:
		c.inventoryController.BuyWeapon(item.Name())
		if c.isModalButtonBuyAmmunitionAvailable() {
			c.model.modalButtonSelection = buttonBuyAmmo
		} else {
			c.model.modalButtonSelection = buttonCloseModal
		}
	case buttonBuyAmmo:
		if item.IsWeapon() {
			weapon := assets.WeaponCrate.GetByName(item.Name())
			c.inventoryController.BuyAmmunition(weapon.Ammo)
		} else if item.IsGrenade() {
			c.inventoryController.BuyGrenade(item.Name())
		}
		if !c.isModalButtonBuyAmmunitionAvailable() {
			c.model.modalButtonSelection = buttonCloseModal
		}
	case buttonCloseModal:
		c.model.modalVisible = false
	}
}

func (c *controller) processButtonPressedOnShop() {
	if c.model.selectedItemIdx == -1 {
		characters.PlayerController.ConfiguredPlayers()[c.model.playerIdx].SelectFirstWeapon()
		if (c.model.playerIdx == 0) && (characters.PlayerController.NumberOfPlayers() == 2) {
			c.screenChangeRequired(common.BuyYourWeaponsP2)
		} else {
			c.screenChangeRequired(common.LetTheMissionBegin)
		}
	} else {
		c.showModal()
	}
}

func (c *controller) showModal() {
	if c.isModalButtonBuyWeaponAvailable() {
		c.model.modalButtonSelection = buttonBuyWeapon
	} else if c.isModalButtonBuyAmmunitionAvailable() {
		c.model.modalButtonSelection = buttonBuyAmmo
	} else {
		c.model.modalButtonSelection = buttonCloseModal
	}
	c.model.modalVisible = true
}

func (c *controller) isModalButtonBuyWeaponAvailable() bool {
	return c.model.selectedItemIdx != -1 &&
		c.model.items[c.model.selectedItemIdx].IsWeapon() &&
		c.inventoryController.WeaponProcurable(c.model.items[c.model.selectedItemIdx].Name())
}

func (c *controller) isModalButtonBuyAmmunitionAvailable() bool {
	if c.model.selectedItemIdx == -1 {
		return false
	}

	var item = c.model.items[c.model.selectedItemIdx]
	if item.IsGrenade() {
		return c.inventoryController.GrenadeProcurable(item.Name())
	}

	var ammoName = item.Name()
	if item.IsWeapon() {
		var weapon = assets.WeaponCrate.GetByName(item.Name())
		ammoName = weapon.Ammo
	}
	return c.inventoryController.AmmunitionProcurable(ammoName)
}
