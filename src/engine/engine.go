package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
)

type GameEngine struct {
	bullets         []Bullet
	enemies         []*characters.ActiveEnemy
	explosives      []*Explosive
	explosions      []*Explosion
	kills           []int
	levelController LevelController
	lost            bool
	mission         *assets.Mission
	playerBehaviors []*characters.PlayerBehavior
	playerPositions []geometry.Rectangle
	won             bool
}
