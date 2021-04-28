package characters

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/util"
	"testing"
)

func TestPlayerInDefaultStateShouldBeAbleToMoveUp(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	var inputState = input.DeviceState{MoveUp: true}
	playerBehavior.Update(&inputState)

	assert.True(t, playerBehavior.Moving)
	assert.Equal(t, geometry.Up, playerBehavior.Direction)
}

func TestPlayerInDefaultStateShouldBeAbleToMoveDiagonally(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	var inputState = input.DeviceState{MoveDown: true, MoveRight: true}
	playerBehavior.Update(&inputState)

	assert.True(t, playerBehavior.Moving)
	assert.Equal(t, geometry.DownRight, playerBehavior.Direction)
}

func TestMovingPlayerShouldBeAbleToChangeDirection(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	playerBehavior.Update(&input.DeviceState{MoveUp: true})
	playerBehavior.Update(&input.DeviceState{MoveLeft: true})

	assert.True(t, playerBehavior.Moving)
	assert.Equal(t, geometry.Left, playerBehavior.Direction)
}

func TestPlayerShouldKeepDirectionWhenStoppingMovementAndKeepingFiring(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	playerBehavior.Update(&input.DeviceState{MoveUp: true, Fire: true})
	playerBehavior.Update(&input.DeviceState{Fire: true})

	assert.False(t, playerBehavior.Moving)
	assert.Equal(t, geometry.Up, playerBehavior.Direction)
	assert.True(t, playerBehavior.Firing)
}

func TestPlayerShouldAbleToChangeDirectionButNotStartMovingWhenFiring(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)
	playerBehavior.Update(&input.DeviceState{Fire: true})
	playerBehavior.Update(&input.DeviceState{MoveLeft: true, Fire: true})

	assert.False(t, playerBehavior.Moving)
	assert.Equal(t, geometry.Left, playerBehavior.Direction)
	assert.True(t, playerBehavior.Firing)
}

func TestSwitchingToNextWeaponGetsFiredOnlyOncePerButtonPress(t *testing.T) {
	assets.WeaponCrate.InitializeInTest("../../../items/weapons/")

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
	playerBehavior.Update(&input.DeviceState{ToggleUp: true})
	playerBehavior.Update(&input.DeviceState{ToggleUp: true})

	assert.Equal(t, 1, callCounter)
	err := player.RemoveChangeListener(&listener)
	assert.Nil(t, err)
}

func TestSwitchingToPreviousWeaponGetsFiredOnlyOncePerButtonPress(t *testing.T) {
	assets.WeaponCrate.InitializeInTest("../../../items/weapons/")

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
	playerBehavior.Update(&input.DeviceState{ToggleDown: true})
	playerBehavior.Update(&input.DeviceState{ToggleDown: true})

	assert.Equal(t, 1, callCounter)
	err := player.RemoveChangeListener(&listener)
	assert.Nil(t, err)
}

func TestTriggeredFireGetsSetWhenFiringAndUnsetWhenHeldDown(t *testing.T) {
	var player = Players[0]
	player.Reset()
	var playerBehavior = NewPlayerBehavior(player)

	playerBehavior.Update(&input.DeviceState{Fire: true})
	assert.True(t, playerBehavior.TriggerPressed)

	playerBehavior.Update(&input.DeviceState{Fire: true})
	assert.False(t, playerBehavior.TriggerPressed)
}
