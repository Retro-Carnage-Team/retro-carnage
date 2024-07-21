package characters

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/util"
)

// ActiveEnemy is an Enemy that is (becoming) visible.
type ActiveEnemy struct {
	Actions                 []assets.EnemyAction
	ActivationSound         string
	currentActionIdx        int
	currentActionElapsed    int64
	Dying                   bool
	DyingAnimationCountDown int64
	Movements               []EnemyMovement
	position                geometry.Rectangle
	Skin                    EnemySkin
	SpawnDelays             []int64
	SpriteSupplier          EnemySpriteSupplier
	Type                    EnemyType
	ViewingDirection        *geometry.Direction
}

// Action returns whether or not it's time to perform the given action.
// It updates the internal state of this ActiveEnemy based on the timeElapsedInMs. If it's time for this ActiveEnemy to
// perform some kind of action, the corresponding action name will be returned. If the ActiveEnemy shouldn't perform any
// action, nil will be returned.
func (e *ActiveEnemy) Action(timeElapsedInMs int64) *string {
	if !e.Type.CanFire() || len(e.Actions) == 0 {
		return nil
	}

	e.currentActionElapsed += timeElapsedInMs
	if e.currentActionElapsed > e.Actions[e.currentActionIdx].Delay {
		e.currentActionElapsed = 0
		e.currentActionIdx = (e.currentActionIdx + 1) % len(e.Actions)
		return &e.Actions[e.currentActionIdx].Action
	}

	return nil
}

// Die will kill this ActiveEnemy (and start it's dying animation)
func (e *ActiveEnemy) Die() {
	e.Dying = true
	e.DyingAnimationCountDown = 1
}

// Move will update the enemies position according to its configured movement pattern and the elapsed time.
func (e *ActiveEnemy) Move(elapsedTimeInMs int64) {
	if !e.Type.CanMove() || len(e.Movements) == 0 {
		return
	}

	var remaining = elapsedTimeInMs
	for (0 < remaining) && (0 < len(e.Movements)) {
		var currentMovement = &e.Movements[0]

		var newViewingDirection = geometry.GetDirectionByName(currentMovement.Direction)
		if newViewingDirection != nil && newViewingDirection.Name != e.ViewingDirection.Name {
			e.ViewingDirection = newViewingDirection
			e.SpriteSupplier = NewEnemyPersonSpriteSupplier(*newViewingDirection)
		}

		var duration = util.MinInt64(remaining, currentMovement.Duration-currentMovement.TimeElapsed)
		e.Position().Add(&geometry.Point{
			X: float64(duration) * currentMovement.OffsetXPerMs,
			Y: float64(duration) * currentMovement.OffsetYPerMs,
		})
		remaining -= duration
		currentMovement.TimeElapsed += duration
		if currentMovement.TimeElapsed >= currentMovement.Duration {
			e.removeFirstMovement()
		}
	}
}

// Spawn returns a new enemy instance - when it's time to do so.
// It updates the internal state of this ActiveEnemy based on the timeElapsedInMs. If it's time for this ActiveEnemy to
// spawn a new enemy, the corresponding enemy will be returned. If the ActiveEnemy shouldn't perform any/ action, nil
// will be returned.
func (e *ActiveEnemy) Spawn(timeElapsedInMs int64) *ActiveEnemy {
	if !e.Type.CanSpawn() || len(e.SpawnDelays) == 0 {
		return nil
	}

	e.currentActionElapsed += timeElapsedInMs
	if e.currentActionElapsed > e.SpawnDelays[e.currentActionIdx] {
		e.currentActionElapsed = 0
		e.currentActionIdx = (e.currentActionIdx + 1) % len(e.SpawnDelays)
		return e.spawnEnemyInstance()
	}

	return nil
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

func (e *ActiveEnemy) removeFirstMovement() {
	if len(e.Movements) == 1 {
		e.Movements = []EnemyMovement{}
		return
	}
	e.Movements = e.Movements[1:]
}

func (e *ActiveEnemy) spawnEnemyInstance() *ActiveEnemy {
	var result = &ActiveEnemy{
		Actions:                 e.Actions,
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               append(make([]EnemyMovement, 0), e.Movements...),
		Skin:                    e.Skin,
		Type:                    Person,
		ViewingDirection:        e.ViewingDirection,
	}
	result.SetPosition(e.Position().Clone())
	result.SpriteSupplier = NewEnemyPersonSpriteSupplier(*result.ViewingDirection)
	return result
}
