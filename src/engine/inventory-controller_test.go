package engine

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"testing"
)

func init() {
	assets.AmmunitionCrate.InitializeInTest("../../items/ammunition/")
	assets.GrenadeCrate.InitializeInTest("../../items/grenades/")
	assets.WeaponCrate.InitializeInTest("../../items/weapons/")
}

func TestInitialAmmoCountIsZeroAndIncreasesByPackageSize(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	characters.Players[0].Reset()
	var ammo = assets.AmmunitionCrate.GetAll()[0]
	assert.Equal(t, 0, inventoryController.AmmunitionCount(ammo.Name))
	inventoryController.BuyAmmunition(ammo.Name)
	assert.Equal(t, ammo.PackageSize, inventoryController.AmmunitionCount(ammo.Name))
}

func TestAmmoCountDoesNotGrowLargerThanMaxCount(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	var ammo = assets.AmmunitionCrate.GetAll()[0]
	var player = characters.Players[0]
	player.Reset()
	player.SetAmmunitionCount(ammo.Name, ammo.MaxCount-1)
	assert.True(t, inventoryController.AmmunitionProcurable(ammo.Name))
	inventoryController.BuyAmmunition(ammo.Name)
	assert.Equal(t, ammo.MaxCount, inventoryController.AmmunitionCount(ammo.Name))
	assert.False(t, inventoryController.AmmunitionProcurable(ammo.Name))
}

func TestBuyingAmmoShouldDecreaseTheAmountCashAvailable(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	var player = characters.Players[0]
	player.Reset()
	var ammo = assets.AmmunitionCrate.GetAll()[0]
	var oldCash = player.Cash()
	inventoryController.BuyAmmunition(ammo.Name)
	assert.Equal(t, oldCash-ammo.Price, player.Cash())
}

func TestInitialGrenadeCountIsZeroAndIncreasesByFive(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	var player = characters.Players[0]
	player.Reset()
	var grenade = assets.GrenadeCrate.GetAll()[0]
	assert.Equal(t, 0, inventoryController.GrenadeCount(grenade.Name))
	inventoryController.BuyGrenade(grenade.Name)
	assert.Equal(t, 5, inventoryController.GrenadeCount(grenade.Name))
}

func TestBuyingAGrenadeShouldDecreaseAmountOfCashAvailable(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	var player = characters.Players[0]
	player.Reset()
	var grenade = assets.GrenadeCrate.GetAll()[0]
	var oldCash = player.Cash()
	inventoryController.BuyGrenade(grenade.Name)
	assert.Equal(t, oldCash-grenade.Price, player.Cash())
}

func TestWeaponsShouldBeProcurableWhenUserHasCashAndDidNotBuyItBefore(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	characters.Players[0].Reset()
	const weapon = "P7"
	assert.False(t, inventoryController.WeaponInInventory(weapon))
	assert.True(t, inventoryController.WeaponProcurable(weapon))
	inventoryController.BuyWeapon(weapon)
	assert.True(t, inventoryController.WeaponInInventory(weapon))
	assert.False(t, inventoryController.WeaponProcurable(weapon))
}

func TestBuyingWeaponsShouldDecreaseAmountOfCashAvailable(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	var player = characters.Players[0]
	player.Reset()
	const weapon = "P7"
	var oldCash = player.Cash()
	inventoryController.BuyWeapon(weapon)
	assert.Less(t, player.Cash(), oldCash)
}

func TestRemovingAmmunitionShouldChangeAmountOfAmountGtZero(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	var player = characters.Players[0]
	player.Reset()

	var grenade = assets.GrenadeCrate.GetAll()[0]
	inventoryController.BuyGrenade(grenade.Name)
	var count = inventoryController.GrenadeCount(grenade.Name)
	assert.Greater(t, count, 0)
	player.SelectFirstWeapon()
	var result = inventoryController.RemoveAmmunition()
	assert.True(t, result)
	assert.Equal(t, count-1, inventoryController.GrenadeCount(grenade.Name))
}
