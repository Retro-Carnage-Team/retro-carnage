package characters

import "retro-carnage/engine/geometry"

type PlayerVisualsAdapter struct {
	playerBehavior *PlayerBehavior
}

func NewPlayerVisualsAdapter(playerBehavior *PlayerBehavior) *PlayerVisualsAdapter {
	return &PlayerVisualsAdapter{
		playerBehavior: playerBehavior,
	}
}

func (pva PlayerVisualsAdapter) Dying() bool {
	return pva.playerBehavior.Dying
}

func (pva PlayerVisualsAdapter) Idle() bool {
	return pva.playerBehavior.Idle()
}

func (pva PlayerVisualsAdapter) Invincible() bool {
	return pva.playerBehavior.Invincible
}

func (pva PlayerVisualsAdapter) PlayerIndex() int {
	return pva.playerBehavior.Player.index
}

func (pva PlayerVisualsAdapter) ViewingDirection() *geometry.Direction {
	return &pva.playerBehavior.Direction
}
