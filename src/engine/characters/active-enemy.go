package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
)

// ActiveEnemy is an Enemy that is (becoming) visible.
type ActiveEnemy struct {
	Actions                 []assets.EnemyAction
	currentActionIdx        int
	currentActionElapsed    int64
	Dying                   bool
	DyingAnimationCountDown int64
	Movements               []*EnemyMovement
	position                geometry.Rectangle
	Skin                    EnemySkin
	SpriteSupplier          EnemySpriteSupplier
	Type                    EnemyType
	ViewingDirection        *geometry.Direction
}

// Action returns whether or not it's time to perform the given action.
// It updates the internal state of this ActiveEnemy based on the timeElapsedInMs. If it's time for this ActiveEnemy to
// perform some kind of action, the corresponding action name will be returned. If the ActiveEnemy shouldn't perform any
// action, nil will be returned.
func (e *ActiveEnemy) Action(timeElapsedInMs int64) *string {
	if 0 == len(e.Actions) {
		return nil
	}

	e.currentActionElapsed += timeElapsedInMs
	if e.currentActionElapsed > e.Actions[e.currentActionIdx].Delay {
		var result = e.Actions[e.currentActionIdx].Action
		e.currentActionElapsed = 0
		e.currentActionIdx = (e.currentActionIdx + 1) % len(e.Actions)
		return &result
	}

	return nil
}

// Die will kill this ActiveEnemy (and start it's dying animation)
func (e *ActiveEnemy) Die() {
	e.Dying = true
	e.DyingAnimationCountDown = 1
}

// Position returns the current position of this enemy.
// Implements SomethingThatExplodes for enemies that are landmines.
func (e *ActiveEnemy) Position() *geometry.Rectangle {
	return &e.position
}

// SetPosition will update the ActiveEnemy's position on screen.
func (e *ActiveEnemy) SetPosition(pos *geometry.Rectangle) {
	e.position = *pos
}
