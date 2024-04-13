package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
)

// GameEngine is the heart and soul of the game screen - the class that contains the actual game logic.
// The idea is that you create an instance of GameEngine everytime you start a new mission. Once the mission is
// finished, you discard the engine and select a new one.
type GameEngine struct {
	bullets             []*Bullet
	burnMarks           []*BurnMark
	enemies             []*characters.ActiveEnemy
	explosives          []*Explosive
	explosions          []*Explosion
	inputController     input.InputController
	inventoryController []*InventoryController
	Kills               []int
	levelController     *LevelController
	Lost                bool
	mission             *assets.Mission
	playerBehaviors     []*characters.PlayerBehavior
	playerPositions     []*geometry.Rectangle
	stereo              *assets.Stereo
	Won                 bool
}

// NewGameEngine creates and initializes a new instance of GameEngine.
func NewGameEngine(mission *assets.Mission) *GameEngine {
	var result = &GameEngine{
		bullets:             make([]*Bullet, 0),
		burnMarks:           make([]*BurnMark, 0),
		enemies:             make([]*characters.ActiveEnemy, 0),
		explosives:          make([]*Explosive, 0),
		explosions:          make([]*Explosion, 0),
		inventoryController: make([]*InventoryController, 0),
		Kills:               []int{0, 0},
		levelController:     NewLevelController(mission.Segments),
		Lost:                false,
		mission:             mission,
		playerBehaviors:     make([]*characters.PlayerBehavior, 0),
		playerPositions:     make([]*geometry.Rectangle, 0),
		stereo:              assets.NewStereo(),
		Won:                 false,
	}

	for idx, p := range characters.PlayerController.ConfiguredPlayers() {
		result.playerBehaviors = append(result.playerBehaviors, characters.NewPlayerBehavior(p))
		result.playerPositions = append(result.playerPositions, &geometry.Rectangle{
			X:      float64(500 + idx*500),
			Y:      1200,
			Width:  PlayerHitRectWidth,
			Height: PlayerHitRectHeight,
		})

		var inventoryController = NewInventoryController(idx)
		result.inventoryController = append(result.inventoryController, &inventoryController)
	}
	return result
}

// SetInputController connects this GameEngine with the input.Controller to be used.
func (ge *GameEngine) SetInputController(controller input.InputController) {
	ge.inputController = controller
}

// Backgrounds is a slice of background sprites that are currently visible.
func (ge *GameEngine) Backgrounds() []graphics.SpriteWithOffset {
	return ge.levelController.VisibleBackgrounds()
}

// UpdateGameState updates the state of the game based on the milliseconds passed since the last updates.
// Once this update has been performed, you can re-render the state of the game to screen.
func (ge *GameEngine) UpdateGameState(elapsedTimeInMs int64) {
	ge.updatePlayerBehavior(elapsedTimeInMs)
	var obstacles = ge.levelController.ObstaclesOnScreen()
	ge.updatePlayerPositionWithMovement(elapsedTimeInMs, obstacles)
	ge.updateEnemies(elapsedTimeInMs)
	ge.updateBullets(elapsedTimeInMs, obstacles)
	ge.updateExplosions(elapsedTimeInMs)
	ge.updateExplosives(elapsedTimeInMs, obstacles)
	ge.handlePlayerWeaponAction(elapsedTimeInMs)

	ge.handleDeadlyCollisionsOfPlayer()
	ge.handleDeadlyCollisionsOfEnemies()
	ge.handlePlayerReachedLevelGoal()

	var scrollOffsets = ge.levelController.UpdatePosition(elapsedTimeInMs, ge.playerPositions)
	ge.scrollObjectsOnScreen(&scrollOffsets)

	var activatedEnemies = ge.levelController.ActivatedEnemies()
	for i := range activatedEnemies {
		var enemy = &activatedEnemies[i]
		ge.enemies = append(ge.enemies, enemy)
	}
}

func (ge *GameEngine) updatePlayerBehavior(elapsedTimeInMs int64) {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var behavior = ge.playerBehaviors[player.Index()]
		if behavior.Dying {
			behavior.UpdateDeath(elapsedTimeInMs)
		} else {
			if behavior.Invincible {
				behavior.UpdateInvincibility(elapsedTimeInMs)
			}

			var inputState, err = ge.inputController.ControllerDeviceState(player.Index())
			if nil != err {
				logging.Warning.Printf("Failed to get input state for player %d: %v\n", player.Index(), err)
			} else if (nil != inputState) && !behavior.Dying {
				behavior.Update(inputState)
			}
		}
	}
	ge.Lost = len(characters.PlayerController.RemainingPlayers()) == 0
}

func (ge *GameEngine) updatePlayerPositionWithMovement(elapsedTimeInMs int64, obstacles []assets.Obstacle) {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var behavior = ge.playerBehaviors[player.Index()]
		if !behavior.Dying && behavior.Moving {
			var oldPosition = ge.playerPositions[player.Index()]
			ge.playerPositions[player.Index()] = UpdatePlayerMovement(
				elapsedTimeInMs,
				behavior.Direction,
				oldPosition,
				obstacles,
			)
		}
	}
}

func (ge *GameEngine) updateEnemies(elapsedTimeInMs int64) {
	var spawnedEnemies = make([]*characters.ActiveEnemy, 0)
	ge.updateEnemiesDeaths(elapsedTimeInMs)
	for _, enemy := range ge.enemies {
		if enemy.Dying {
			continue
		}

		enemy.Move(elapsedTimeInMs)

		var enemyAction = enemy.Action(elapsedTimeInMs)
		if nil != enemyAction {
			if assets.EnemyActionBullet == *enemyAction {
				ge.bullets = append(ge.bullets, NewBulletFiredByEnemy(enemy))
			} else if assets.EnemyActionGrenade == *enemyAction {
				var grenade = NewExplosiveGrenadeByEnemy(enemy.Position(), *enemy.ViewingDirection)
				ge.explosives = append(ge.explosives, grenade)
			} else {
				logging.Warning.Printf("Invalid enemy configuration. Unknown action %s", *enemyAction)
			}
		}

		var spawnedEnemy = enemy.Spawn(elapsedTimeInMs)
		if nil != spawnedEnemy {
			spawnedEnemies = append(spawnedEnemies, spawnedEnemy)
		}
	}

	if len(spawnedEnemies) > 0 {
		ge.enemies = append(ge.enemies, spawnedEnemies...)
	}
}

// updateEnemiesDeaths updates the dying animation countdown of all active enemies.
// Removes those enemies that have a remaining count down <= 0.
func (ge *GameEngine) updateEnemiesDeaths(elapsedTimeInMs int64) {
	for i := len(ge.enemies) - 1; i >= 0; i-- {
		if ge.enemies[i].Dying {
			ge.enemies[i].DyingAnimationCountDown -= elapsedTimeInMs
			if 0 >= ge.enemies[i].DyingAnimationCountDown {
				ge.removeEnemy(i)
			}
		}
	}
}

func (ge *GameEngine) updateExplosions(elapsedTimeInMs int64) {
	for i := len(ge.explosions) - 1; i >= 0; i-- {
		ge.explosions[i].duration += elapsedTimeInMs
		if ge.explosions[i].CreatesMark() {
			ge.burnMarks = append(ge.burnMarks, ge.explosions[i].CreateMark())
			ge.explosions[i].hasMark = true
		}
		if ge.explosions[i].duration >= durationOfExplosion {
			ge.removeExplosion(i)
		}
	}
}

func (ge *GameEngine) updateExplosives(elapsedTimeInMs int64, obstacles []assets.Obstacle) {
	for i := len(ge.explosives) - 1; i >= 0; i-- {
		var done = ge.explosives[i].Move(elapsedTimeInMs)
		if !done {
			for _, obstacle := range obstacles {
				if !done && obstacle.StopsExplosives && (nil != obstacle.Intersection(ge.explosives[i].position)) {
					done = true
				}
			}
		}
		if done {
			ge.detonateExplosive(ge.explosives[i])
			ge.removeExplosive(i)
		}
	}
}

func (ge *GameEngine) detonateExplosive(explosive *Explosive) {
	ge.explosions = append(ge.explosions, NewExplosion(explosive.FiredByPlayer, explosive.FiredByPlayerIdx, explosive))
	ge.stereo.PlayFx(assets.FxGrenade1)
}

func (ge *GameEngine) updateBullets(elapsedTimeInMs int64, obstacles []assets.Obstacle) {
	for i := len(ge.bullets) - 1; i >= 0; i-- {
		var reachedRange = ge.bullets[i].Move(elapsedTimeInMs)
		var hitObstacle = false
		for _, obstacle := range obstacles {
			if !hitObstacle && obstacle.StopsBullets && (nil != obstacle.Intersection(ge.bullets[i].Position())) {
				hitObstacle = true
			}
		}
		if reachedRange || hitObstacle {
			ge.removeBullet(i)
		}
	}
}

// scrollObjectsOnScreen updates the positions of all elements on screen with the given scrollOffset.
// The objects will be removed if they leave the screen with this adjustment.
func (ge *GameEngine) scrollObjectsOnScreen(scrollOffset *geometry.Point) {
	for _, playerPosition := range ge.playerPositions {
		playerPosition.Subtract(scrollOffset)
	}

	for idx := len(ge.explosives) - 1; idx >= 0; idx-- {
		ge.adjustPositionedItemWithScrollOffset(ge.explosives[idx], ge.removeExplosive, scrollOffset, idx)
	}

	for idx := len(ge.explosions) - 1; idx >= 0; idx-- {
		ge.adjustPositionedItemWithScrollOffset(ge.explosions[idx], ge.removeExplosion, scrollOffset, idx)
	}

	for idx := len(ge.enemies) - 1; idx >= 0; idx-- {
		var enemy = ge.enemies[idx]
		var hasBeenVisible = nil != enemy.Position().Intersection(screenRect)
		enemy.Position().Subtract(scrollOffset)
		var isVisible = nil != enemy.Position().Intersection(screenRect)
		if hasBeenVisible && !isVisible {
			ge.removeEnemy(idx)
		}
	}

	for idx := len(ge.bullets) - 1; idx >= 0; idx-- {
		ge.adjustPositionedItemWithScrollOffset(ge.bullets[idx], ge.removeBullet, scrollOffset, idx)
	}

	for idx := len(ge.burnMarks) - 1; idx >= 0; idx-- {
		ge.adjustPositionedItemWithScrollOffset(ge.burnMarks[idx], ge.removeBurnMark, scrollOffset, idx)
	}
}

// adjustPositionedItemWithScrollOffset adjusts the position of the given object with the given scroll offset.
// The adjusted object will be removed with deleteValueFunc if it is not on screen anymore.
func (ge *GameEngine) adjustPositionedItemWithScrollOffset(
	value geometry.Positioned,
	deleteValueFunc func(idx int),
	scrollOffset *geometry.Point,
	idx int,
) {
	value.Position().Subtract(scrollOffset)
	if nil == value.Position().Intersection(screenRect) {
		deleteValueFunc(idx)
	}
}

// handlePlayerWeaponAction updates the game state based on a weapon action of the player.
func (ge *GameEngine) handlePlayerWeaponAction(elapsedTimeInMs int64) {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var behavior = ge.playerBehaviors[player.Index()]
		if !behavior.Dying {
			var playerPosition = ge.playerPositions[player.Index()]
			if behavior.TriggerPressed {
				ge.handlePlayerWeaponTriggerPressed(player, playerPosition, behavior)
			} else if behavior.Firing && player.AutomaticWeaponSelected() {
				ge.handlePlayerWeaponTriggerHeld(behavior, elapsedTimeInMs, player)
			}

			if behavior.TriggerReleased && player.AutomaticWeaponSelected() {
				ge.stereo.StopFx(player.SelectedWeapon().Sound)
			}
		}
	}
}

// handlePlayerWeaponTriggerPressed updates the game state when a player just triggered his weapon.
// This handles both guns and explosives.
func (ge *GameEngine) handlePlayerWeaponTriggerPressed(
	player *characters.Player,
	playerPosition *geometry.Rectangle,
	behavior *characters.PlayerBehavior,
) {
	if player.GrenadeSelected() && ge.inventoryController[player.Index()].RemoveAmmunition() {
		ge.explosives = append(ge.explosives, NewExplosiveGrenadeByPlayer(
			player.Index(),
			playerPosition,
			behavior.Direction,
			player.SelectedGrenade(),
		))
	} else if player.RpgSelected() && ge.inventoryController[player.Index()].RemoveAmmunition() {
		var weapon = player.SelectedWeapon()
		ge.stereo.PlayFx(weapon.Sound)
		ge.explosives = append(
			ge.explosives,
			NewExplosiveRpg(player.Index(), playerPosition, behavior.Direction, weapon).Explosive,
		)
	} else if (player.PistolSelected() || player.AutomaticWeaponSelected()) &&
		ge.inventoryController[player.Index()].RemoveAmmunition() {
		ge.stereo.PlayFx(player.SelectedWeapon().Sound)
		ge.fireBullet(player, behavior)
	}
}

// handlePlayerWeaponTriggerHeld updates the game state when a player keeps the trigger of this weapon held down.
// This handles both guns and explosives.
func (ge *GameEngine) handlePlayerWeaponTriggerHeld(
	behavior *characters.PlayerBehavior,
	elapsedTimeInMs int64,
	player *characters.Player,
) {
	behavior.TimeSinceLastBullet += elapsedTimeInMs
	var weapon = player.SelectedWeapon()
	if (int64(weapon.BulletInterval) <= behavior.TimeSinceLastBullet) &&
		ge.inventoryController[player.Index()].RemoveAmmunition() {
		ge.fireBullet(player, behavior)
	}
}

// fireBullet creates a new bullet for a player firing his weapon.
func (ge *GameEngine) fireBullet(player *characters.Player, behavior *characters.PlayerBehavior) {
	var weapon = player.SelectedWeapon()
	var position = ge.playerPositions[player.Index()]
	var bullet = NewBulletFiredByPlayer(player.Index(), position, behavior.Direction, weapon)
	ge.bullets = append(ge.bullets, bullet)
	behavior.TimeSinceLastBullet = 0
}

// handleDeadlyCollisionsOfPlayer checks the player's position for collisions with various deadly things.
// Kills the player if any of these collisions has been detected.
func (ge *GameEngine) handleDeadlyCollisionsOfPlayer() {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var behavior = ge.playerBehaviors[player.Index()]
		if !behavior.Dying && !behavior.Invincible {
			var rect = ge.playerPositions[player.Index()]
			if ge.checkPlayerForDeadlyCollisionWithEnemy(rect) ||
				ge.checkPlayerForCollisionWithExplosion(rect) ||
				ge.checkPlayerForCollisionWithBullet(rect) {
				if player.AutomaticWeaponSelected() && behavior.Firing {
					ge.stereo.StopFx(player.SelectedWeapon().Sound)
				}
				ge.stereo.PlayFx(assets.DeathFxForPlayer(player.Index()))
				behavior.Die()
			}
		}
	}
}

// checkPlayerForDeadlyCollisionWithEnemy returns true when the player collided with a deadly enemy.
func (ge *GameEngine) checkPlayerForDeadlyCollisionWithEnemy(rect *geometry.Rectangle) bool {
	for _, enemy := range ge.enemies {
		if !enemy.Dying && enemy.Type.IsCollisionDeadly() {
			var collisionWithEnemy = rect.Intersection(enemy.Position())
			if nil != collisionWithEnemy {
				if characters.Landmine == enemy.Type {
					ge.explosions = append(ge.explosions, NewExplosion(false, -1, enemy))
					ge.stereo.PlayFx(assets.FxGrenade2)
				}
				return true
			}
		}
	}
	return false
}

// checkPlayerForCollisionWithExplosion returns true when the player collided with an explosion.
func (ge *GameEngine) checkPlayerForCollisionWithExplosion(rect *geometry.Rectangle) bool {
	for _, explosion := range ge.explosions {
		if nil != rect.Intersection(explosion.Position()) {
			return true
		}
	}
	return false
}

// checkPlayerForCollisionWithBullet returns true when the player collided with a bullet.
func (ge *GameEngine) checkPlayerForCollisionWithBullet(rect *geometry.Rectangle) bool {
	for _, bullet := range ge.bullets {
		if nil != rect.Intersection(bullet.Position()) {
			return true
		}
	}
	return false
}

// handleDeadlyCollisionsOfEnemies checks of enemies collide with deadly objects - like bullets and explosions.
// Enemy will be killed if a deadly collision is detected.
func (ge *GameEngine) handleDeadlyCollisionsOfEnemies() {
	for _, enemy := range ge.enemies {
		if enemy.CanDie() {
			// Check for hits by explosion
			death, killer := ge.checkEnemyForCollisionWithExplosion(enemy)

			// Check for hits by bullets and explosives
			if !death {
				death, killer = ge.checkEnemyForCollisionWithBullet(enemy)
			}

			if !death {
				death, killer = ge.checkEnemyForCollisionWithExplosive(enemy)
			}

			if death {
				ge.killEnemy(enemy, killer)
			}
		}
	}
}

// checkEnemyForCollisionWithExplosion checks this enemy for deadly collisions with explosions.
// Returns true and index of the player that caused the explosion if such a collision is detected.
func (ge *GameEngine) checkEnemyForCollisionWithExplosion(enemy *characters.ActiveEnemy) (death bool, killer int) {
	for _, explosion := range ge.explosions {
		if nil != enemy.Position().Intersection(explosion.Position()) {
			if characters.Landmine == enemy.Type {
				var newExplosion = NewExplosion(explosion.causedByPlayer, explosion.playerIdx, enemy)
				ge.explosions = append(ge.explosions, newExplosion)
			}
			return true, explosion.playerIdx
		}
	}
	return false, -1
}

// checkEnemyForCollisionWithBullet checks this enemy for deadly collisions with bullets.
// Returns true and index of the player that fired the bullet if such a collision is detected.
func (ge *GameEngine) checkEnemyForCollisionWithBullet(enemy *characters.ActiveEnemy) (death bool, killer int) {
	if characters.Person == enemy.Type && !death {
		for _, bullet := range ge.bullets {
			if nil != enemy.Position().Intersection(bullet.Position()) {
				return true, bullet.playerIdx
			}
		}
	}
	return false, -1
}

// checkEnemyForCollisionWithExplosive checks this enemy for deadly collisions with explosives.
// Returns true and index of the player that fired the explosive if such a collision is detected.
func (ge *GameEngine) checkEnemyForCollisionWithExplosive(enemy *characters.ActiveEnemy) (death bool, killer int) {
	for i := len(ge.explosives) - 1; i >= 0; i-- {
		var explosive = ge.explosives[i]
		if explosive.ExplodesOnContact && nil != explosive.Position().Intersection(enemy.Position()) {
			ge.detonateExplosive(explosive)
			ge.removeExplosive(i)
			return true, explosive.FiredByPlayerIdx
		}
	}
	return false, -1
}

func (ge *GameEngine) killEnemy(enemy *characters.ActiveEnemy, killer int) {
	enemy.Dying = true
	enemy.DyingAnimationCountDown = 1
	if killer != -1 {
		ge.Kills[killer] += 1
		var player = ge.playerBehaviors[killer].Player
		player.SetScore(player.Score() + enemy.Type.GetPointsForKill())
	}
	if characters.Person == enemy.Type {
		ge.stereo.PlayFx(assets.RandomEnemyDeathSoundEffect())
		enemy.DyingAnimationCountDown = characters.DurationOfEnemyDeathAnimation
	}
}

func (ge *GameEngine) handlePlayerReachedLevelGoal() {
	if ge.levelController.GoalReached(ge.playerPositions) {
		ge.Won = true
	}
}

func (ge *GameEngine) removeBullet(idx int) {
	ge.bullets[idx] = ge.bullets[len(ge.bullets)-1]
	ge.bullets[len(ge.bullets)-1] = nil
	ge.bullets = ge.bullets[:len(ge.bullets)-1]
}

func (ge *GameEngine) removeBurnMark(idx int) {
	ge.burnMarks[idx] = ge.burnMarks[len(ge.burnMarks)-1]
	ge.burnMarks[len(ge.burnMarks)-1] = nil
	ge.burnMarks = ge.burnMarks[:len(ge.burnMarks)-1]
}

func (ge *GameEngine) removeEnemy(idx int) {
	ge.enemies[idx] = ge.enemies[len(ge.enemies)-1]
	ge.enemies[len(ge.enemies)-1] = nil
	ge.enemies = ge.enemies[:len(ge.enemies)-1]
}

func (ge *GameEngine) removeExplosion(idx int) {
	ge.explosions[idx] = ge.explosions[len(ge.explosions)-1]
	ge.explosions[len(ge.explosions)-1] = nil
	ge.explosions = ge.explosions[:len(ge.explosions)-1]
}

func (ge *GameEngine) removeExplosive(idx int) {
	ge.explosives[idx] = ge.explosives[len(ge.explosives)-1]
	ge.explosives[len(ge.explosives)-1] = nil
	ge.explosives = ge.explosives[:len(ge.explosives)-1]
}
