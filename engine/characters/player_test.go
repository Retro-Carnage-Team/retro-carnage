package characters

import (
	"os"
	"path/filepath"
	"retro-carnage/assets"
	"retro-carnage/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	M1           = "M4A1"
	HAND_GRENADE = "DM51"
	PANZERFAUST  = "RPG-7"
	PISTOL       = "P30"
)

func init() {
	assets.GrenadeCrate.InitializeInTest(filepath.Join(os.Getenv("RC-ASSETS"), "items/grenades/"))
	assets.WeaponCrate.InitializeInTest(filepath.Join(os.Getenv("RC-ASSETS"), "items/weapons/"))
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
	player.SetWeaponInInventory(PISTOL, true)

	assert.Equal(t, 1, callCounter)
	assert.NotNil(t, 1, value)
	assert.Equal(t, *name, PlayerPropertyWeapons)

	err := player.RemoveChangeListener(&listener)
	assert.Nil(t, err)

	player.SetWeaponInInventory(PISTOL, false)
	assert.Equal(t, 1, callCounter)
}

func TestSelectFirstWeaponShouldSelectFirstWeaponInInventory(t *testing.T) {
	var player = Players[0]
	player.Reset()
	player.SetWeaponInInventory(M1, true)
	player.SetWeaponInInventory(PANZERFAUST, true)
	player.SelectFirstWeapon()
	assert.Equal(t, M1, player.SelectedWeapon().GetName())
}

func TestSelectFirstWeaponShouldSelectFirstGrenadeIfNoWeaponInInventory(t *testing.T) {
	var player = Players[0]
	player.Reset()
	player.SetGrenadeCount(HAND_GRENADE, 42)
	player.SelectFirstWeapon()

	assert.Nil(t, player.SelectedWeapon())
	assert.Equal(t, HAND_GRENADE, player.SelectedGrenade().Name)
}

func TestSelectNextWeaponShouldIterateAllWeaponsAndGrenadesInInventory(t *testing.T) {
	var player = Players[0]
	player.Reset()
	player.SetGrenadeCount(HAND_GRENADE, 1)
	player.SetWeaponInInventory(M1, true)
	player.SetWeaponInInventory(PANZERFAUST, true)

	player.SelectFirstWeapon()
	assert.Equal(t, M1, player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())

	player.SelectNextWeapon()
	assert.Equal(t, PANZERFAUST, player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())

	player.SelectNextWeapon()
	assert.Nil(t, player.SelectedWeapon())
	assert.Equal(t, HAND_GRENADE, player.SelectedGrenade().Name)

	player.SelectNextWeapon()
	assert.Equal(t, M1, player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())
}

func TestSelectPreviousWeaponShouldIterateAllWeaponsAndGrenadesInInventory(t *testing.T) {
	var player = Players[0]
	player.Reset()
	player.SetGrenadeCount(HAND_GRENADE, 1)
	player.SetWeaponInInventory(M1, true)
	player.SetWeaponInInventory(PANZERFAUST, true)

	player.SelectFirstWeapon()
	assert.Equal(t, M1, player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())

	player.SelectPreviousWeapon()
	assert.Nil(t, player.SelectedWeapon())
	assert.Equal(t, HAND_GRENADE, player.SelectedGrenade().Name)

	player.SelectPreviousWeapon()
	assert.Equal(t, PANZERFAUST, player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())

	player.SelectPreviousWeapon()
	assert.Equal(t, M1, player.SelectedWeapon().GetName())
	assert.Nil(t, player.SelectedGrenade())
}

func TestWeaponTypeInfoShouldWork(t *testing.T) {
	var player = Players[0]
	player.Reset()

	player.SetWeaponInInventory(PISTOL, true)
	player.SetWeaponInInventory(M1, true)
	player.SetWeaponInInventory(PANZERFAUST, true)
	player.SetGrenadeCount(HAND_GRENADE, 1)

	player.SelectFirstWeapon()
	assert.Equal(t, PISTOL, player.SelectedWeapon().GetName())
	assert.False(t, player.AutomaticWeaponSelected())
	assert.False(t, player.GrenadeSelected())
	assert.True(t, player.PistolSelected())
	assert.False(t, player.RpgSelected())

	player.SelectNextWeapon()
	assert.Equal(t, M1, player.SelectedWeapon().GetName())
	assert.True(t, player.AutomaticWeaponSelected())
	assert.False(t, player.GrenadeSelected())
	assert.False(t, player.PistolSelected())
	assert.False(t, player.RpgSelected())

	player.SelectNextWeapon()
	assert.Equal(t, PANZERFAUST, player.SelectedWeapon().GetName())
	assert.False(t, player.AutomaticWeaponSelected())
	assert.False(t, player.GrenadeSelected())
	assert.False(t, player.PistolSelected())
	assert.True(t, player.RpgSelected())

	player.SelectNextWeapon()
	assert.Equal(t, HAND_GRENADE, player.SelectedGrenade().Name)
	assert.False(t, player.AutomaticWeaponSelected())
	assert.True(t, player.GrenadeSelected())
	assert.False(t, player.PistolSelected())
	assert.False(t, player.RpgSelected())
}
