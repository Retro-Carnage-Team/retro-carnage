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
	kills               []int
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
		kills:               []int{0, 0},
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
	for _, activatedEnemy := range activatedEnemies {
		ge.enemies = append(ge.enemies, &activatedEnemy)
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
	ge.Lost = 0 == len(characters.PlayerController.RemainingPlayers())
}

func (ge *GameEngine) updatePlayerPositionWithMovement(elapsedTimeInMs int64, obstacles []*geometry.Rectangle) {
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
	var enemies = ge.updateEnemiesDeaths(elapsedTimeInMs)
	for _, enemy := range enemies {
		if !enemy.Dying && (characters.Person == enemy.Type) && (0 < len(enemy.Movements)) {
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
	}
	ge.enemies = enemies
}

func (ge *GameEngine) removeFirstEnemyMovement(movements []*characters.EnemyMovement) []*characters.EnemyMovement {
	if 1 == len(movements) {
		return []*characters.EnemyMovement{}
	}
	movements[0] = nil
	return movements[1:]
}

// updateEnemiesDeaths updates the dying animation countdown of all active enemies.
// Returns those enemies that have a remaining count down > 0.
func (ge *GameEngine) updateEnemiesDeaths(elapsedTimeInMs int64) []*characters.ActiveEnemy {
	var enemies = ge.enemies
	for i := len(enemies) - 1; i >= 0; i-- {
		var enemy = enemies[i]
		if enemy.Dying {
			enemy.DyingAnimationCountDown -= elapsedTimeInMs
			if 0 >= enemy.DyingAnimationCountDown {
				enemies = ge.removeEnemy(enemies, i)
			}
		}
	}
	return enemies
}

func (ge *GameEngine) removeEnemy(enemies []*characters.ActiveEnemy, idx int) []*characters.ActiveEnemy {
	enemies[idx] = enemies[len(enemies)-1]
	enemies[len(enemies)-1] = nil
	return enemies[:len(enemies)-1]
}

func (ge *GameEngine) updateExplosions(elapsedTimeInMs int64) {
	var explosions = ge.explosions
	for i := len(explosions) - 1; i >= 0; i-- {
		var explosion = explosions[i]
		explosion.duration += elapsedTimeInMs
		if explosion.CreatesMark() {
			ge.burnMarks = append(ge.burnMarks, explosion.CreateMark())
			explosion.hasMark = true
		}
		if explosion.duration >= durationOfExplosion {
			explosions = ge.removeExplosion(explosions, i)
		}
	}
	ge.explosions = explosions
}

func (ge *GameEngine) removeExplosion(explosions []*Explosion, idx int) []*Explosion {
	explosions[idx] = explosions[len(explosions)-1]
	explosions[len(explosions)-1] = nil
	return explosions[:len(explosions)-1]
}

func (ge *GameEngine) updateExplosives(elapsedTimeInMs int64, obstacles []*geometry.Rectangle) {
	var explosives = ge.explosives
	for i := len(explosives) - 1; i >= 0; i-- {
		var explosive = explosives[i]
		var done = explosive.Move(elapsedTimeInMs)
		if !done && explosive.ExplodesOnContact {
			for _, obstacle := range obstacles {
				if !done && (nil != obstacle.Intersection(explosive.position)) {
					done = true
				}
			}
		}
		if done {
			ge.detonateExplosive(explosive)
			explosives = ge.removeExplosive(explosives, i)
		}
	}
	ge.explosives = explosives
}

func (ge *GameEngine) removeExplosive(explosives []*Explosive, idx int) []*Explosive {
	explosives[idx] = explosives[len(explosives)-1]
	explosives[len(explosives)-1] = nil
	return explosives[:len(explosives)-1]
}

func (ge *GameEngine) detonateExplosive(explosive *Explosive) {
	ge.explosions = append(ge.explosions, NewExplosion(explosive.FiredByPlayer, explosive.FiredByPlayerIdx, explosive))
	ge.stereo.PlayFx(assets.FxGrenade1)
}

func (ge *GameEngine) updateBullets(elapsedTimeInMs int64, obstacles []*geometry.Rectangle) {
	var bullets = ge.bullets
	for i := len(bullets) - 1; i >= 0; i-- {
		var reachedRange = bullets[i].Move(elapsedTimeInMs)
		var hitObstacle = false
		for _, obstacle := range obstacles {
			if !hitObstacle && (nil != obstacle.Intersection(bullets[i].Position)) {
				hitObstacle = true
			}
		}
		if reachedRange || hitObstacle {
			bullets = ge.removeBullet(bullets, i)
		}
	}
	ge.bullets = bullets
}

func (ge *GameEngine) removeBullet(bullets []*Bullet, idx int) []*Bullet {
	bullets[idx] = bullets[len(bullets)-1]
	bullets[len(bullets)-1] = nil
	return bullets[:len(bullets)-1]
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
			ge.explosives = ge.removeExplosive(ge.explosives, idx)
		}
	}

	for idx := len(ge.explosions) - 1; idx >= 0; idx-- {
		var explosion = ge.explosions[idx]
		explosion.Position.Subtract(scrollOffset)
		if nil == explosion.Position.Intersection(visibleArea) {
			ge.explosions = ge.removeExplosion(ge.explosions, idx)
		}
	}

	for idx := len(ge.enemies) - 1; idx >= 0; idx-- {
		var enemy = ge.enemies[idx]
		var hasBeenVisible = nil != enemy.Position().Intersection(visibleArea)
		enemy.Position().Subtract(scrollOffset)
		var isVisible = nil != enemy.Position().Intersection(visibleArea)
		if hasBeenVisible && !isVisible {
			ge.enemies = ge.removeEnemy(ge.enemies, idx)
		}
	}

	for idx := len(ge.bullets) - 1; idx >= 0; idx-- {
		var bullet = ge.bullets[idx]
		bullet.Position.Subtract(scrollOffset)
		if nil == bullet.Position.Intersection(visibleArea) {
			ge.bullets = ge.removeBullet(ge.bullets, idx)
		}
	}

	for idx := len(ge.burnMarks) - 1; idx >= 0; idx-- {
		var burnMark = ge.burnMarks[idx]
		burnMark.Position.Subtract(scrollOffset)
		if nil == burnMark.Position.Intersection(visibleArea) {
			ge.burnMarks = ge.removeBurnMark(ge.burnMarks, idx)
		}
	}
}

func (ge *GameEngine) removeBurnMark(burnMarks []*BurnMark, idx int) []*BurnMark {
	burnMarks[idx] = burnMarks[len(burnMarks)-1]
	burnMarks[len(burnMarks)-1] = nil
	return burnMarks[:len(burnMarks)-1]
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
			} else if behavior.Firing && player.AutomaticWeaponSelected() &&
				ge.inventoryController[player.Index()].RemoveAmmunition() {
				behavior.TimeSinceLastBullet += elapsedTimeInMs
				var weapon = player.SelectedWeapon()
				if int64(weapon.BulletInterval) <= behavior.TimeSinceLastBullet {
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
	bullet.applyPlayerOffset()
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
				var collisionWithEnemy = rect.Intersection(enemy.Position())
				if nil != collisionWithEnemy {
					if characters.Landmine == enemy.Type {
						ge.explosions = append(ge.explosions, NewExplosion(false, -1, enemy))
						ge.stereo.PlayFx(assets.FxGrenade2)
					}
					death = true
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
				var deadlyExplosion = nil != enemy.Position().Intersection(explosion.Position)
				if deadlyExplosion {
					killer = explosion.playerIdx
					if characters.Landmine == enemy.Type {
						var newExplosion = NewExplosion(explosion.causedByPlayer, explosion.playerIdx, enemy)
						ge.explosions = append(ge.explosions, newExplosion)
					}
				}
				death = death || deadlyExplosion
			}

			// Check for hits by bullets, flamethrowers and RPGs (useful only against persons)
			if characters.Person == enemy.Type {
				for _, bullet := range ge.bullets {
					var deadlyShot = nil != enemy.Position().Intersection(bullet.Position)
					if deadlyShot {
						killer = bullet.playerIdx
					}
					death = death || deadlyShot
				}

				var explosives = ge.explosives
				for i := len(explosives) - 1; i >= 0; i-- {
					var explosive = explosives[i]
					var explode = explosive.ExplodesOnContact && nil != explosive.Position().Intersection(enemy.Position())
					if explode {
						ge.detonateExplosive(explosive)
						death = true
						explosives = ge.removeExplosive(explosives, i)
					}
				}
				ge.explosives = explosives
			}

			if death {
				enemy.Dying = true
				enemy.DyingAnimationCountDown = 1
				if -1 != killer {
					ge.kills[killer] += 1
				}
				if characters.Person == enemy.Type {
					ge.stereo.PlayFx(assets.RandomEnemyDeathSoundEffect())
					enemy.DyingAnimationCountDown = characters.DurationOfEnemyDeathAnimation
				}
			}
		}
	}
}

func (ge *GameEngine) checkIfPlayerReachedLevelGoal() {
	if ge.levelController.GoalReached(ge.playerPositions) {
		ge.Won = true
	}
}
