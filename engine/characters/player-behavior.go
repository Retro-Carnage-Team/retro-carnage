package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/input"
)

const (
	PlayerInvincibilityTimeout = 1_500
)

// PlayerBehavior contains all player state that is valid for the duration of a single mission only.
type PlayerBehavior struct {
	Player                  *Player
	Direction               geometry.Direction
	Dying                   bool
	dyingAnimationCountDown int64
	Invincible              bool
	invincibilityCountDown  int64
	TimeSinceLastBullet     int64
	Firing                  bool // will be true as long as the player keeps the trigger pressed
	TriggerPressed          bool // will be true only when switching from "not firing" to "firing"
	TriggerReleased         bool // will be true only when switching from "firing" to "not firing"
	Moving                  bool
	NextWeapon              bool
	PreviousWeapon          bool
}

// NewPlayerBehavior creates and initializes a new PlayerBehavior instance.
func NewPlayerBehavior(player *Player) *PlayerBehavior {
	return &PlayerBehavior{
		Player:                  player,
		Direction:               geometry.Up,
		Dying:                   false,
		dyingAnimationCountDown: 0,
		Invincible:              false,
		invincibilityCountDown:  0,
		Moving:                  false,
		Firing:                  false,
		TriggerPressed:          false,
		TriggerReleased:         false,
		NextWeapon:              false,
		PreviousWeapon:          false,
		TimeSinceLastBullet:     0,
	}
}

func (pb *PlayerBehavior) Update(userInput *input.InputDeviceState) {
	if nil == userInput || !pb.Player.Alive() {
		return
	}

	var playerWantsToMove = userInput.MoveUp || userInput.MoveDown || userInput.MoveLeft || userInput.MoveRight
	pb.Moving = playerWantsToMove && !(!pb.Moving && pb.Firing)
	pb.TriggerPressed = !pb.Firing && userInput.PrimaryAction
	pb.TriggerReleased = pb.Firing && !userInput.PrimaryAction
	pb.Firing = userInput.PrimaryAction
	if playerWantsToMove {
		pb.Direction = pb.direction(userInput.MoveUp, userInput.MoveDown, userInput.MoveLeft, userInput.MoveRight)
	}

	if !pb.NextWeapon && userInput.ToggleUp {
		pb.Player.SelectNextWeapon()
	}
	pb.NextWeapon = userInput.ToggleUp

	if !pb.PreviousWeapon && userInput.ToggleDown {
		pb.Player.SelectPreviousWeapon()
	}
	pb.PreviousWeapon = userInput.ToggleDown
}

// GetDirection returns the geometry.Direction specified by combination of the given cardinal directions. Returns the
// last direction of the player of no such direction exists (e.g. if none of the given parameters is true).
func (pb *PlayerBehavior) direction(up bool, down bool, left bool, right bool) geometry.Direction {
	var direction = geometry.GetDirectionForCardinals(up, down, left, right)
	if nil == direction {
		return pb.Direction
	}
	return *direction
}

func (pb *PlayerBehavior) Die() {
	pb.Dying = true
	pb.dyingAnimationCountDown = SkinForPlayer(pb.Player.index).DurationOfDeathAnimation()
}

func (pb *PlayerBehavior) UpdateDeath(elapsedTimeInMs int64) {
	pb.dyingAnimationCountDown -= elapsedTimeInMs
	if pb.dyingAnimationCountDown <= 0 {
		pb.Dying = false
		pb.dyingAnimationCountDown = 0
		PlayerController.KillPlayer(pb.Player)
		if pb.Player.Alive() {
			pb.StartInvincibility()
		}
	}
}

func (pb *PlayerBehavior) Idle() bool {
	return !pb.Dying && !pb.Moving
}

func (pb *PlayerBehavior) StartInvincibility() {
	pb.Invincible = true
	pb.invincibilityCountDown = PlayerInvincibilityTimeout
}

func (pb *PlayerBehavior) UpdateInvincibility(elapsedTimeInMs int64) {
	pb.invincibilityCountDown -= elapsedTimeInMs
	if 0 >= pb.invincibilityCountDown {
		pb.Invincible = false
		pb.invincibilityCountDown = 0
	}
}
