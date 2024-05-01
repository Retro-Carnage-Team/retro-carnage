package characters

import (
	"os"
	"path/filepath"
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerInDefaultStateShouldBeAbleToMoveUp(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	var inputState = input.InputDeviceState{MoveUp: true}
	playerBehavior.Update(&inputState)

	assert.True(t, playerBehavior.Moving)
	assert.Equal(t, geometry.Up, playerBehavior.Direction)
}

func TestPlayerInDefaultStateShouldBeAbleToMoveDiagonally(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	var inputState = input.InputDeviceState{MoveDown: true, MoveRight: true}
	playerBehavior.Update(&inputState)

	assert.True(t, playerBehavior.Moving)
	assert.Equal(t, geometry.DownRight, playerBehavior.Direction)
}

func TestMovingPlayerShouldBeAbleToChangeDirection(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	playerBehavior.Update(&input.InputDeviceState{MoveUp: true})
	playerBehavior.Update(&input.InputDeviceState{MoveLeft: true})

	assert.True(t, playerBehavior.Moving)
	assert.Equal(t, geometry.Left, playerBehavior.Direction)
}

func TestPlayerShouldKeepDirectionWhenStoppingMovementAndKeepingFiring(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	playerBehavior.Update(&input.InputDeviceState{MoveUp: true, PrimaryAction: true})
	playerBehavior.Update(&input.InputDeviceState{PrimaryAction: true})

	assert.False(t, playerBehavior.Moving)
	assert.Equal(t, geometry.Up, playerBehavior.Direction)
	assert.True(t, playerBehavior.Firing)
}

func TestPlayerShouldAbleToChangeDirectionButNotStartMovingWhenFiring(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	playerBehavior.Update(&input.InputDeviceState{PrimaryAction: true})
	playerBehavior.Update(&input.InputDeviceState{MoveLeft: true, PrimaryAction: true})

	assert.False(t, playerBehavior.Moving)
	assert.Equal(t, geometry.Left, playerBehavior.Direction)
	assert.True(t, playerBehavior.Firing)
}

func TestSwitchingToNextWeaponGetsFiredOnlyOncePerButtonPress(t *testing.T) {
	assets.WeaponCrate.InitializeInTest(filepath.Join(os.Getenv("RC-ASSETS"), "items/weapons/"))

	var callCounter = 0
	var callback = func(v interface{}, n string) {
		callCounter += 1
	}
	var listener = util.ChangeListener{Callback: callback, PropertyNames: []string{}}
	var player = Players[0]
	player.Reset()
	player.SetWeaponInInventory(assets.WeaponCrate.GetAll()[0].Name, true)
	player.SetWeaponInInventory(assets.WeaponCrate.GetAll()[1].Name, true)
	player.SelectFirstWeapon()
	player.AddChangeListener(&listener)
	var playerBehavior = NewPlayerBehavior(player)
	playerBehavior.Update(&input.InputDeviceState{ToggleUp: true})
	playerBehavior.Update(&input.InputDeviceState{ToggleUp: true})

	assert.Equal(t, 1, callCounter)
	err := player.RemoveChangeListener(&listener)
	assert.Nil(t, err)
}

func TestSwitchingToPreviousWeaponGetsFiredOnlyOncePerButtonPress(t *testing.T) {
	assets.WeaponCrate.InitializeInTest(filepath.Join(os.Getenv("RC-ASSETS"), "items/weapons/"))

	var callCounter = 0
	var callback = func(v interface{}, n string) {
		callCounter += 1
	}
	var listener = util.ChangeListener{Callback: callback, PropertyNames: []string{}}
	var player = Players[0]
	player.Reset()
	player.SetWeaponInInventory(assets.WeaponCrate.GetAll()[0].Name, true)
	player.SetWeaponInInventory(assets.WeaponCrate.GetAll()[1].Name, true)
	player.SelectFirstWeapon()
	player.AddChangeListener(&listener)
	var playerBehavior = NewPlayerBehavior(player)
	playerBehavior.Update(&input.InputDeviceState{ToggleDown: true})
	playerBehavior.Update(&input.InputDeviceState{ToggleDown: true})

	assert.Equal(t, 1, callCounter)
	err := player.RemoveChangeListener(&listener)
	assert.Nil(t, err)
}

func TestTriggeredFireGetsSetWhenFiringAndUnsetWhenHeldDown(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)

	playerBehavior.Update(&input.InputDeviceState{PrimaryAction: true})
	assert.True(t, playerBehavior.TriggerPressed)

	playerBehavior.Update(&input.InputDeviceState{PrimaryAction: true})
	assert.False(t, playerBehavior.TriggerPressed)
}
