package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/util"
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
	inputController     input.Controller
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
func (ge *GameEngine) SetInputController(controller input.Controller) {
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
	ge.handleWeaponAction(elapsedTimeInMs)

	ge.checkPlayersForDeadlyCollisions()
	ge.checkEnemiesForDeadlyCollisions()
	ge.checkIfPlayerReachedLevelGoal()

	var scrollOffsets = ge.levelController.UpdatePosition(elapsedTimeInMs, ge.playerPositions)
	ge.updateAllPositionsWithScrollOffset(&scrollOffsets)

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
			behavior.DyingAnimationCountDown -= elapsedTimeInMs
			if 0 >= behavior.DyingAnimationCountDown {
				behavior.Dying = false
				behavior.DyingAnimationCountDown = 0
				characters.PlayerController.KillPlayer(player)
				if player.Alive() {
					behavior.StartInvincibility()
				}
			}
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
	ge.updateEnemiesDeaths(elapsedTimeInMs)
	for _, enemy := range ge.enemies {
		if !enemy.Dying && (characters.Person == enemy.Type) {
			if 0 < len(enemy.Movements) {
				var remaining = elapsedTimeInMs
				for (0 < remaining) && (0 < len(enemy.Movements)) {
					var currentMovement = enemy.Movements[0]
					var duration = util.MinInt64(remaining, currentMovement.Duration-currentMovement.TimeElapsed)
					enemy.Position().Add(&geometry.Point{
						X: float64(duration) * currentMovement.OffsetXPerMs,
						Y: float64(duration) * currentMovement.OffsetYPerMs,
					})
					remaining -= duration
					currentMovement.TimeElapsed += duration
					if currentMovement.TimeElapsed >= currentMovement.Duration {
						enemy.Movements = ge.removeFirstEnemyMovement(enemy.Movements)
					}
				}
			}

			var enemyAction = enemy.Action(elapsedTimeInMs)
			if nil != enemyAction {
				if assets.EnemyActionBullet == *enemyAction {
					ge.bullets = append(ge.bullets, NewBulletFiredByEnemy(enemy))
				} else if assets.EnemyActionGrenade == *enemyAction {
					// TODO: Throw a grenade
				} else {
					logging.Warning.Printf("Invalid enemy configuration. Unknown action %s", *enemyAction)
				}
			}
		}
	}
}

func (ge *GameEngine) removeFirstEnemyMovement(movements []*characters.EnemyMovement) []*characters.EnemyMovement {
	if len(movements) == 1 {
		return []*characters.EnemyMovement{}
	}
	movements[0] = nil
	return movements[1:]
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

func (ge *GameEngine) removeEnemy(idx int) {
	ge.enemies[idx] = ge.enemies[len(ge.enemies)-1]
	ge.enemies[len(ge.enemies)-1] = nil
	ge.enemies = ge.enemies[:len(ge.enemies)-1]
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
			if !hitObstacle && obstacle.StopsBullets && (nil != obstacle.Intersection(ge.bullets[i].Position)) {
				hitObstacle = true
			}
		}
		if reachedRange || hitObstacle {
			ge.removeBullet(i)
		}
	}
}

func (ge *GameEngine) updateAllPositionsWithScrollOffset(scrollOffset *geometry.Point) {
	var visibleArea = screenRect()
	for _, playerPosition := range ge.playerPositions {
		playerPosition.Subtract(scrollOffset)
	}

	for idx := len(ge.explosives) - 1; idx >= 0; idx-- {
		var explosive = ge.explosives[idx]
		explosive.Position().Subtract(scrollOffset)
		if nil == explosive.Position().Intersection(visibleArea) {
			ge.removeExplosive(idx)
		}
	}

	for idx := len(ge.explosions) - 1; idx >= 0; idx-- {
		var explosion = ge.explosions[idx]
		explosion.Position.Subtract(scrollOffset)
		if nil == explosion.Position.Intersection(visibleArea) {
			ge.removeExplosion(idx)
		}
	}

	for idx := len(ge.enemies) - 1; idx >= 0; idx-- {
		var enemy = ge.enemies[idx]
		var hasBeenVisible = nil != enemy.Position().Intersection(visibleArea)

		if !scrollOffset.Zero() {
			logging.Trace.Printf("Moving enemy from %s by %s", enemy.Position().String(), scrollOffset.String())
		}

		enemy.Position().Subtract(scrollOffset)
		var isVisible = nil != enemy.Position().Intersection(visibleArea)
		if hasBeenVisible && !isVisible {
			ge.removeEnemy(idx)
		}
	}

	for idx := len(ge.bullets) - 1; idx >= 0; idx-- {
		var bullet = ge.bullets[idx]
		bullet.Position.Subtract(scrollOffset)
		if nil == bullet.Position.Intersection(visibleArea) {
			ge.removeBullet(idx)
		}
	}

	for idx := len(ge.burnMarks) - 1; idx >= 0; idx-- {
		var burnMark = ge.burnMarks[idx]
		burnMark.Position.Subtract(scrollOffset)
		if nil == burnMark.Position.Intersection(visibleArea) {
			ge.removeBurnMark(idx)
		}
	}
}

func (ge *GameEngine) handleWeaponAction(elapsedTimeInMs int64) {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var behavior = ge.playerBehaviors[player.Index()]
		if !behavior.Dying {
			var playerPosition = ge.playerPositions[player.Index()]
			if behavior.TriggerPressed {
				if player.GrenadeSelected() && ge.inventoryController[player.Index()].RemoveAmmunition() {
					ge.explosives = append(ge.explosives, NewExplosiveGrenadeByPlayer(
						player.Index(),
						playerPosition,
						behavior.Direction,
						player.SelectedGrenade(),
					).Explosive)
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
			} else if behavior.Firing && player.AutomaticWeaponSelected() {
				behavior.TimeSinceLastBullet += elapsedTimeInMs
				var weapon = player.SelectedWeapon()
				if (int64(weapon.BulletInterval) <= behavior.TimeSinceLastBullet) &&
					ge.inventoryController[player.Index()].RemoveAmmunition() {
					ge.fireBullet(player, behavior)
				}
			}

			if behavior.TriggerReleased && player.AutomaticWeaponSelected() {
				ge.stereo.StopFx(player.SelectedWeapon().Sound)
			}
		}
	}
}

func (ge *GameEngine) fireBullet(player *characters.Player, behavior *characters.PlayerBehavior) {
	var weapon = player.SelectedWeapon()
	var position = ge.playerPositions[player.Index()]
	var bullet = NewBulletFiredByPlayer(player.Index(), position, behavior.Direction, weapon)
	ge.bullets = append(ge.bullets, bullet)
	behavior.TimeSinceLastBullet = 0
}

func (ge *GameEngine) checkPlayersForDeadlyCollisions() {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var behavior = ge.playerBehaviors[player.Index()]
		if !behavior.Dying && !behavior.Invincible {
			var rect = ge.playerPositions[player.Index()]
			var death = false

			for _, enemy := range ge.enemies {
				if enemy.Type.IsCollisionDeadly() {
					var collisionWithEnemy = rect.Intersection(enemy.Position())
					if nil != collisionWithEnemy {
						if characters.Landmine == enemy.Type {
							ge.explosions = append(ge.explosions, NewExplosion(false, -1, enemy))
							ge.stereo.PlayFx(assets.FxGrenade2)
						}
						death = true
					}
				}
			}

			if !death {
				for _, explosion := range ge.explosions {
					death = death || (nil != rect.Intersection(explosion.Position))
				}
			}

			if !death {
				for _, bullet := range ge.bullets {
					death = death || (nil != rect.Intersection(bullet.Position))
				}
			}

			if death {
				if player.AutomaticWeaponSelected() && behavior.Firing {
					ge.stereo.StopFx(player.SelectedWeapon().Sound)
				}
				ge.stereo.PlayFx(assets.DeathFxForPlayer(player.Index()))
				behavior.Dying = true
				behavior.DyingAnimationCountDown = characters.SkinForPlayer(player.Index()).DurationOfDeathAnimation()
			}
		}
	}
}

func (ge *GameEngine) checkEnemiesForDeadlyCollisions() {
	for _, enemy := range ge.enemies {
		if !enemy.Dying {
			var death = false
			var killer = -1

			// Check for hits by explosion
			for _, explosion := range ge.explosions {
				if nil != enemy.Position().Intersection(explosion.Position) {
					killer = explosion.playerIdx
					if characters.Landmine == enemy.Type {
						var newExplosion = NewExplosion(explosion.causedByPlayer, explosion.playerIdx, enemy)
						ge.explosions = append(ge.explosions, newExplosion)
					}
					death = true
					break
				}
			}

			// Check for hits by bullets and explosives
			if characters.Person == enemy.Type {
				if !death {
					for _, bullet := range ge.bullets {
						if nil != enemy.Position().Intersection(bullet.Position) {
							killer = bullet.playerIdx
							death = true
							break
						}
					}
				}

				if !death {
					for i := len(ge.explosives) - 1; i >= 0; i-- {
						var explosive = ge.explosives[i]
						if explosive.ExplodesOnContact && nil != explosive.Position().Intersection(enemy.Position()) {
							ge.detonateExplosive(explosive)
							killer = explosive.FiredByPlayerIdx
							ge.removeExplosive(i)
							death = true
							break
						}
					}
				}
			}

			if death {
				ge.killEnemy(enemy, killer)
			}
		}
	}
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

func (ge *GameEngine) checkIfPlayerReachedLevelGoal() {
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
