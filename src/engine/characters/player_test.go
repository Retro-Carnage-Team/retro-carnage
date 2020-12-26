package characters

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/assets"
	"retro-carnage/util"
	"testing"
)

func init() {
	assets.GrenadeCrate.InitializeInTest("../../../items/grenades/")
	assets.WeaponCrate.InitializeInTest("../../../items/weapons/")
}

func TestChangeListenersShouldGetInformedEveryTimeAChangeHappened(t *testing.T) {
	callCounter := 0
	var value interface{} = nil
	var name *string = nil

	var callback = func(v interface{}, n string) {
		value = v
		name = &n
		callCounter += 1
	}

	var listener = util.ChangeListener{Callback: callback, PropertyNames: []string{}}

	var player = Players[0]
	player.Reset()
	player.AddChangeListener(&listener)
	player.SetWeaponInInventory("P7", true)

	assert.Equal(t, 1, callCounter)
	assert.NotNil(t, 1, value)
	assert.Equal(t, *name, PlayerPropertyWeapons)

	err := player.RemoveChangeListener(&listener)
	assert.Nil(t, err)

	player.SetWeaponInInventory("P7", false)
	assert.Equal(t, 1, callCounter)
}

func TestSelectFirstWeaponShouldSelectFirstWeaponInInventory(t *testing.T) {
	var player = Players[0]
	player.Reset()
	player.SetWeaponInInventory("AR-10", true)
	player.SetWeaponInInventory("Panzerfaust 3", true)
	player.SelectFirstWeapon()
	assert.Equal(t, "AR-10", player.SelectedWeapon().GetName())
}

func TestSelectFirstWeaponShouldSelectFirstGrenadeIfNoWeaponInInventory(t *testing.T) {
	var player = Players[0]
	player.Reset()
	player.SetGrenadeCount("Stielhandgranate 24", 42)
	player.SelectFirstWeapon()

	assert.Nil(t, player.SelectedWeapon())
	assert.Equal(t, "Stielhandgranate 24", player.SelectedGrenade().Name)
}

func TestSelectNextWeaponShouldIterateAllWeaponsAndGrenadesInInventory(t *testing.T) {
	var player = Players[0]
	player.Reset()
	player.SetGrenadeCount("Stielhandgranate 24", 1)
	player.SetWeaponInInventory("AR-10", true)
	player.SetWeaponInInventory("Panzerfaust 3", true)

	player.SelectFirstWeapon()
	assert.Equal(t, "AR-10", player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())

	player.SelectNextWeapon()
	assert.Equal(t, "Panzerfaust 3", player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())

	player.SelectNextWeapon()
	assert.Nil(t, player.SelectedWeapon())
	assert.Equal(t, "Stielhandgranate 24", player.SelectedGrenade().Name)

	player.SelectNextWeapon()
	assert.Equal(t, "AR-10", player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())
}

func TestSelectPreviousWeaponShouldIterateAllWeaponsAndGrenadesInInventory(t *testing.T) {
	var player = Players[0]
	player.Reset()
	player.SetGrenadeCount("Stielhandgranate 24", 1)
	player.SetWeaponInInventory("AR-10", true)
	player.SetWeaponInInventory("Panzerfaust 3", true)

	player.SelectFirstWeapon()
	assert.Equal(t, "AR-10", player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())

	player.SelectPreviousWeapon()
	assert.Nil(t, player.SelectedWeapon())
	assert.Equal(t, "Stielhandgranate 24", player.SelectedGrenade().Name)

	player.SelectPreviousWeapon()
	assert.Equal(t, "Panzerfaust 3", player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())

	player.SelectPreviousWeapon()
	assert.Equal(t, "AR-10", player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())
}

func TestWeaponTypeInfoShouldWork(t *testing.T) {
	var player = Players[0]
	player.Reset()

	player.SetWeaponInInventory("P210", true)
	player.SetWeaponInInventory("AR-10", true)
	player.SetWeaponInInventory("Panzerfaust 44", true)
	player.SetGrenadeCount("Stielhandgranate 24", 1)

	player.SelectFirstWeapon()
	assert.Equal(t, "P210", player.SelectedWeapon().GetName())
	assert.False(t, player.AutomaticWeaponSelected())
	assert.False(t, player.GrenadeSelected())
	assert.True(t, player.PistolSelected())
	assert.False(t, player.RpgSelected())

	player.SelectNextWeapon()
	assert.Equal(t, "AR-10", player.SelectedWeapon().GetName())
	assert.True(t, player.AutomaticWeaponSelected())
	assert.False(t, player.GrenadeSelected())
	assert.False(t, player.PistolSelected())
	assert.False(t, player.RpgSelected())

	player.SelectNextWeapon()
	assert.Equal(t, "Panzerfaust 44", player.SelectedWeapon().GetName())
	assert.False(t, player.AutomaticWeaponSelected())
	assert.False(t, player.GrenadeSelected())
	assert.False(t, player.PistolSelected())
	assert.True(t, player.RpgSelected())

	player.SelectNextWeapon()
	assert.Equal(t, "Stielhandgranate 24", player.SelectedGrenade().Name)
	assert.False(t, player.AutomaticWeaponSelected())
	assert.True(t, player.GrenadeSelected())
	assert.False(t, player.PistolSelected())
	assert.False(t, player.RpgSelected())
}
