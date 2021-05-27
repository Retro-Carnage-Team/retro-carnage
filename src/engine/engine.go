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

type GameEngine struct {
	bullets         []*Bullet
	enemies         []*characters.ActiveEnemy
	explosives      []*Explosive
	explosions      []*Explosion
	inputController input.Controller
	kills           []int
	levelController *LevelController
	lost            bool
	mission         *assets.Mission
	playerBehaviors []*characters.PlayerBehavior
	playerPositions []*geometry.Rectangle
	won             bool
}

func NewGameEngine(mission *assets.Mission) *GameEngine {
	var result = &GameEngine{
		bullets:         make([]*Bullet, 0),
		enemies:         make([]*characters.ActiveEnemy, 0),
		explosives:      make([]*Explosive, 0),
		explosions:      make([]*Explosion, 0),
		kills:           []int{0, 0},
		levelController: NewLevelController(mission.Segments),
		lost:            false,
		mission:         mission,
		playerBehaviors: make([]*characters.PlayerBehavior, 0),
		playerPositions: make([]*geometry.Rectangle, 0),
		won:             false,
	}

	for idx, p := range characters.PlayerController.ConfiguredPlayers() {
		result.playerBehaviors = append(result.playerBehaviors, characters.NewPlayerBehavior(p))
		result.playerPositions = append(result.playerPositions, &geometry.Rectangle{
			X:      float64(500 + idx*500),
			Y:      1200,
			Width:  PlayerHitRectWidth,
			Height: PlayerHitRectHeight,
		})
	}
	return result
}

func (ge *GameEngine) InitializeGameState() {
}

func (ge *GameEngine) SetInputController(controller input.Controller) {
	ge.inputController = controller
}

func (ge *GameEngine) Backgrounds() []graphics.SpriteWithOffset {
	return ge.levelController.VisibleBackgrounds()
}

/*
updateGameState = (elapsedTimeInMs: number) => {
    this.updatePlayerBehavior(elapsedTimeInMs);
    const obstacles = this.levelController.getObstaclesOnScreen();
    this.updatePlayerPositionWithMovement(elapsedTimeInMs, obstacles);
    this.updateEnemies(elapsedTimeInMs);
    this.updateBullets(elapsedTimeInMs, obstacles);
    this.updateExplosions(elapsedTimeInMs);
    this.updateExplosives(elapsedTimeInMs, obstacles);
    this.handleWeaponAction(elapsedTimeInMs);

    this.checkPlayersForDeadlyCollisions();
    this.checkEnemiesForDeadlyCollisions();
    this.checkIfPlayerReachedLevelGoal();

    const scrollOffsets = this.levelController.updatePosition(
      elapsedTimeInMs,
      this.playerPositions
    );
    this.updateAllPositionsWithScrollOffset(scrollOffsets);

    const activatedEnemies = this.levelController.getActivatedEnemies();
    if (0 < activatedEnemies.length) {
      this.enemies.push(...activatedEnemies.map((e) => new ActiveEnemy(e)));
    }
  };
*/

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

			var inputState, err = ge.inputController.GetControllerDeviceState(player.Index())
			if nil != err {
				logging.Warning.Printf("Failed to get input state for player %d: %v\n", player.Index(), err)
			} else if (nil != inputState) && !behavior.Dying {
				behavior.Update(inputState)
			}
		}
	}
	ge.lost = 0 == len(characters.PlayerController.RemainingPlayers())
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
				enemy.Position.Add(&geometry.Point{
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
		}
		if enemy.Dying && 0 <= enemy.DyingAnimationCountDown {
			enemies = ge.removeEnemy(enemies, i)
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
	const durationOfExplosion = DurationOfExplosionFrame * NumberOfExplosionSprites
	var explosions = ge.explosions
	for i := len(explosions) - 1; i >= 0; i-- {
		explosions[i].duration += elapsedTimeInMs
		if explosions[i].duration < durationOfExplosion {
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
	assets.NewStereo().PlayFx(assets.FxGrenade1)
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
	for _, playerPosition := range ge.playerPositions {
		playerPosition.Subtract(scrollOffset)
	}
	for _, explosive := range ge.explosives {
		explosive.Position().Subtract(scrollOffset)
	}
	for _, explosion := range ge.explosions {
		explosion.Position.Subtract(scrollOffset)
	}
	for _, enemy := range ge.enemies {
		enemy.Position.Subtract(scrollOffset)
	}
	for _, bullet := range ge.bullets {
		bullet.Position.Subtract(scrollOffset)
	}
}

/*
  handleWeaponAction = (elapsedTimeInMs: number) => {
    const _this = this;
    function fireBullet(p: Player, behavior: PlayerBehavior): void {
      const weapon = p.getSelectedWeapon() as Weapon;
      const position = _this.playerPositions[p.index];
      const bullet = new Bullet(p.index, position, behavior.direction, weapon);
      bullet.applyOffset(
        0 === p.index ? BulletOffsetForPlayer0 : BulletOffsetForPlayer1
      );
      _this.bullets.push(bullet);
      behavior.timeSinceLastBullet = 0;
    }

    PlayerController.getRemainingPlayers().forEach((p) => {
      const behavior = this.playerBehaviors[p.index];
      if (!behavior.dying) {
        const playerPosition = this.playerPositions[p.index];
        if (behavior.triggerPressed) {
          if (
            p.isGrenadeSelected() &&
            InventoryController.removeAmmunition(p.index)
          ) {
            this.explosives.push(
              new ExplosiveGrenade(
                p.index,
                new Rectangle(
                  playerPosition.x,
                  playerPosition.y,
                  GRENADE_WIDTH,
                  GRENADE_HEIGHT
                ),
                behavior.direction,
                p.getSelectedWeapon() as Grenade
              )
            );
          } else if (
            p.isRpgSelected() &&
            InventoryController.removeAmmunition(p.index)
          ) {
            const weapon = p.getSelectedWeapon() as Weapon;
            if (weapon.sound) SoundBoard.play(weapon.sound);
            this.explosives.push(
              new ExplosiveRPG(
                p.index,
                new Rectangle(
                  playerPosition.x,
                  playerPosition.y,
                  RPG_WIDTH,
                  RPG_HEIGHT
                ),
                behavior.direction,
                p.getSelectedWeapon() as Weapon
              )
            );
          } else if (
            (p.isPistolSelected() || p.isAutomaticWeaponSelected()) &&
            InventoryController.removeAmmunition(p.index)
          ) {
            const weapon = p.getSelectedWeapon() as Weapon;
            if (weapon.sound) SoundBoard.play(weapon.sound);
            fireBullet(p, behavior);
          }
        } else if (
          behavior.firing &&
          p.isAutomaticWeaponSelected() &&
          InventoryController.removeAmmunition(p.index)
        ) {
          behavior.timeSinceLastBullet += elapsedTimeInMs;
          const weapon = p.getSelectedWeapon() as Weapon;
          if (weapon.bulletInterval! <= behavior.timeSinceLastBullet) {
            fireBullet(p, behavior);
          }
        }

        if (behavior.triggerReleased && p.isAutomaticWeaponSelected()) {
          const weapon = p.getSelectedWeapon() as Weapon;
          SoundBoard.stop(weapon.sound!);
        }
      }
    });
  };

  // Check if players collide with explosions / bullets / enemies
  checkPlayersForDeadlyCollisions = () => {
    PlayerController.getRemainingPlayers().forEach((p) => {
      const behavior = this.playerBehaviors[p.index];
      if (!behavior.dying && !behavior.invincible) {
        const rect = this.playerPositions[p.index];
        let death = false;
        this.enemies.forEach((enemy) => {
          const collisionWithEnemy = rect.getIntersection(enemy.position);
          if (collisionWithEnemy) {
            if (EnemyType.Landmine === enemy.type) {
              this.explosions.push(
                new Explosion({
                  playerIdx: null,
                  position: enemy.position,
                })
              );
              SoundBoard.play(FX_GRENADE_2);
            }
            death = true;
          }
        });
        this.explosions.forEach((explosion) => {
          death = death || null !== rect.getIntersection(explosion.position);
        });
        this.bullets.forEach((bullet) => {
          death = death || null !== rect.getIntersection(bullet.position);
        });
        if (death) {
          if (p.isAutomaticWeaponSelected() && behavior.firing) {
            SoundBoard.stop((p.getSelectedWeapon() as Weapon).sound!);
          }
          SoundBoard.play(
            0 === p.index ? FX_DEATH_PLAYER_1 : FX_DEATH_PLAYER_2
          );
          behavior.dying = true;
          behavior.dyingAnimationCountDown =
            0 === p.index
              ? DURATION_OF_DEATH_ANIMATION_PLAYER_0
              : DURATION_OF_DEATH_ANIMATION_PLAYER_1;
        }
      }
    });
  };

*/

func (ge *GameEngine) checkEnemiesForDeadlyCollisions() {
	for _, enemy := range ge.enemies {
		if !enemy.Dying {
			var death = false
			var killer = -1

			// Check for hits by explosion
			for _, explosion := range ge.explosions {
				var deadlyExplosion = nil != enemy.Position.Intersection(explosion.Position)
				if deadlyExplosion {
					killer = explosion.playerIdx
				}
				death = death || deadlyExplosion
			}

			// Check for hits by bullets, flamethrowers and RPGs (useful only against persons)
			if characters.Person == enemy.Type {
				for _, bullet := range ge.bullets {
					var deadlyShot = nil != enemy.Position.Intersection(bullet.Position)
					if deadlyShot {
						killer = bullet.playerIdx
					}
					death = death || deadlyShot
				}

				var explosives = ge.explosives
				for i := len(explosives) - 1; i >= 0; i-- {
					var explosive = explosives[i]
					var explode = explosive.ExplodesOnContact && nil != explosive.position.Intersection(enemy.Position)
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
					assets.NewStereo().PlayFx(assets.RandomEnemyDeathSoundEffect())
					enemy.DyingAnimationCountDown = characters.DurationOfEnemyDeathAnimation
				}
			}
		}
	}
}

func (ge *GameEngine) checkIfPlayerReachedLevelGoal() {
	if ge.levelController.GoalReached(ge.playerPositions) {
		ge.won = true
	}
}
