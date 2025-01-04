package characters

import (
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type ActiveEnemyVisuals struct {
	activeEnemy *ActiveEnemy
}

func (eav ActiveEnemyVisuals) Dying() bool {
	return eav.activeEnemy.Dying
}

func (eav ActiveEnemyVisuals) Skin() graphics.EnemySkin {
	return eav.activeEnemy.Skin
}

func (eav ActiveEnemyVisuals) ViewingDirection() *geometry.Direction {
	return eav.activeEnemy.ViewingDirection
}
