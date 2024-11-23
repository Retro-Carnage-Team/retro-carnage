package engine

import (
	"os"
	"path/filepath"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	rc_assets = "RC-ASSETS"
	PISTOL    = "P30"
)

func init() {
	assets.AmmunitionCrate.InitializeInTest(filepath.Join(os.Getenv(rc_assets), "items/ammunition/"))
	assets.GrenadeCrate.InitializeInTest(filepath.Join(os.Getenv(rc_assets), "items/grenades/"))
	assets.WeaponCrate.InitializeInTest(filepath.Join(os.Getenv(rc_assets), "items/weapons/"))
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

	assert.False(t, inventoryController.WeaponInInventory(PISTOL))
	assert.True(t, inventoryController.WeaponProcurable(PISTOL))
	inventoryController.BuyWeapon(PISTOL)
	assert.True(t, inventoryController.WeaponInInventory(PISTOL))
	assert.False(t, inventoryController.WeaponProcurable(PISTOL))
}

func TestBuyingWeaponsShouldDecreaseAmountOfCashAvailable(t *testing.T) {
	var inventoryController = NewInventoryController(0)
	var player = characters.Players[0]
	player.Reset()
	var oldCash = player.Cash()
	inventoryController.BuyWeapon(PISTOL)
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
