package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/util"
)

// InventoryController offers easy access to a player's inventory and can manipulate it.
type InventoryController struct {
	playerIdx int
	stereo    *assets.Stereo
}

// NewInventoryController creates and initializes a new instance of InventoryController.
func NewInventoryController(playerIdx int) InventoryController {
	return InventoryController{
		playerIdx: playerIdx,
		stereo:    assets.NewStereo(),
	}
}

// Returns the current number of bullets of the specified ammunition type.
func (ic *InventoryController) AmmunitionCount(ammunitionName string) int {
	return characters.Players[ic.playerIdx].AmmunitionCount(ammunitionName)
}

// AmmunitionProcurable returns whether or not a player can purchase more ammunition of the specified type.
func (ic *InventoryController) AmmunitionProcurable(ammunitionName string) bool {
	var ammunition = assets.AmmunitionCrate.GetByName(ammunitionName)
	var player = characters.Players[ic.playerIdx]
	var currentAmount = player.AmmunitionCount(ammunitionName)
	return (currentAmount < ammunition.MaxCount) && (player.Cash() >= ammunition.Price)
}

// BuyAmmunition purchases a box of ammunition of the specified type.
func (ic *InventoryController) BuyAmmunition(ammunitionName string) {
	if ic.AmmunitionProcurable(ammunitionName) {
		var ammunition = assets.AmmunitionCrate.GetByName(ammunitionName)
		var player = characters.Players[ic.playerIdx]
		var increasedCount = player.AmmunitionCount(ammunitionName) + ammunition.PackageSize
		player.SetAmmunitionCount(ammunitionName, util.MinInt(increasedCount, ammunition.MaxCount))
		player.SetCash(player.Cash() - ammunition.GetPrice())
		ic.stereo.PlayFx(assets.FxCash)
	} else {
		ic.stereo.PlayFx(assets.FxError)
	}
}

// RemoveAmmunition tries to remove one piece of ammunition for the currently selected weapon. The function returns
// whether or not that was possible
func (ic *InventoryController) RemoveAmmunition() bool {
	var player = characters.Players[ic.playerIdx]
	var selectedWeapon = player.SelectedWeapon()
	if nil != selectedWeapon {
		var ammoCount = player.AmmunitionCount(selectedWeapon.Ammo)
		if 0 < ammoCount {
			player.SetAmmunitionCount(selectedWeapon.Ammo, ammoCount-1)
			return true
		} else {
			ic.stereo.StopFx(selectedWeapon.Sound)
			ic.stereo.PlayFx(assets.FxOutOfAmmo)
			return false
		}
	}

	var selectedGrenade = player.SelectedGrenade()
	if nil != selectedGrenade {
		var grenadeCount = player.GrenadeCount(selectedGrenade.Name)
		if 0 < grenadeCount {
			player.SetGrenadeCount(selectedGrenade.Name, grenadeCount-1)
			return true
		}
		// There is no "out of ammo" sound fx for grenades
		return false
	}
	return false
}

// Returns the current number of grenades of the specified type.
func (ic *InventoryController) GrenadeCount(grenadeName string) int {
	return characters.Players[ic.playerIdx].GrenadeCount(grenadeName)
}

// GrenadeProcurable returns whether or not a player can purchase more grenades of the specified type.
func (ic *InventoryController) GrenadeProcurable(grenadeName string) bool {
	var grenade = assets.GrenadeCrate.GetByName(grenadeName)
	var player = characters.Players[ic.playerIdx]
	var currentAmount = player.GrenadeCount(grenadeName)
	return (currentAmount < grenade.MaxCount) && (player.Cash() >= grenade.Price)
}

// BuyGrenade purchases a grenade of the specified type.
func (ic *InventoryController) BuyGrenade(grenadeName string) {
	if ic.GrenadeProcurable(grenadeName) {
		var grenade = assets.GrenadeCrate.GetByName(grenadeName)
		var player = characters.Players[ic.playerIdx]
		var increasedCount = player.GrenadeCount(grenadeName) + grenade.PackageSize
		player.SetGrenadeCount(grenadeName, util.MinInt(increasedCount, grenade.MaxCount))
		player.SetCash(player.Cash() - grenade.Price)
		ic.stereo.PlayFx(assets.FxCash)
	} else {
		ic.stereo.PlayFx(assets.FxError)
	}
}

// WeaponInInventory returns whether or not the player has purchased the specified weapon.
func (ic *InventoryController) WeaponInInventory(weaponName string) bool {
	return characters.Players[ic.playerIdx].WeaponInInventory(weaponName)
}

// WeaponProcurable returns whether or not a player can purchase the specified weapon.
func (ic *InventoryController) WeaponProcurable(weaponName string) bool {
	var weapon = assets.WeaponCrate.GetByName(weaponName)
	var player = characters.Players[ic.playerIdx]
	return !player.WeaponInInventory(weaponName) && player.Cash() >= weapon.GetPrice()
}

// BuyWeapon purchases the specified weapon.
func (ic *InventoryController) BuyWeapon(weaponName string) {
	if ic.WeaponProcurable(weaponName) {
		var weapon = assets.WeaponCrate.GetByName(weaponName)
		var player = characters.Players[ic.playerIdx]
		player.SetWeaponInInventory(weaponName, true)
		player.SetCash(player.Cash() - weapon.GetPrice())
		ic.stereo.PlayFx(assets.FxCash)
	} else {
		ic.stereo.PlayFx(assets.FxError)
	}
}
