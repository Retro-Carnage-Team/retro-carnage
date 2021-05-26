package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

type GameEngine struct {
	bullets         []*Bullet
	enemies         []*characters.ActiveEnemy
	explosives      []*Explosive
	explosions      []*Explosion
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

  updatePlayerBehavior = (elapsedTimeInMs: number) => {
    PlayerController.getRemainingPlayers().forEach((p) => {
      const behavior = this.playerBehaviors[p.index];
      if (behavior.dying) {
        behavior.dyingAnimationCountDown -= Math.floor(elapsedTimeInMs);
        if (0 >= behavior.dyingAnimationCountDown) {
          behavior.dying = false;
          behavior.dyingAnimationCountDown = 0;
          PlayerController.killPlayer(p.index);
          if (p.isAlive()) {
            behavior.invincible = true;
            behavior.invincibilityCountDown = 1500;
          }
        }
      } else {
        if (behavior.invincible) {
          behavior.invincibilityCountDown -= Math.floor(elapsedTimeInMs);
          if (0 >= behavior.invincibilityCountDown) {
            behavior.invincible = false;
            behavior.invincibilityCountDown = 0;
          }
        }
        const inputState = InputController.inputProviders[p.index]();
        if (inputState && !behavior.dying) {
          behavior.update(inputState);
        }
      }
    });

    this.lost = 0 === PlayerController.getRemainingPlayers().length;
  };

*/

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

/*
  updateEnemies = (elapsedTimeInMs: number) => {
    function updateDeathAnimationCountDown(enemy: ActiveEnemy): ActiveEnemy {
      if (enemy.dying) {
        enemy.dyingAnimationCountDown -= Math.floor(elapsedTimeInMs);
      }
      return enemy;
    }

    this.enemies = this.enemies
      .map(updateDeathAnimationCountDown)
      .filter((enemy) => !enemy.dying || 0 <= enemy.dyingAnimationCountDown);

    this.enemies
      .filter((e) => !e.dying && EnemyType.Person === e.type)
      .forEach((enemy) => {
        let remaining = elapsedTimeInMs;
        let currentMovement = enemy.movements.find(
          (m) => m.timeElapsed < m.duration
        );
        while (currentMovement && 0 < remaining) {
          const duration = Math.min(
            remaining,
            currentMovement.duration - currentMovement.timeElapsed
          );
          enemy.position.add({
            x: duration * currentMovement.offsetXPerMs,
            y: duration * currentMovement.offsetYPerMs,
          });
          remaining -= duration;
          currentMovement.timeElapsed += duration;
          if (0 < remaining) {
            currentMovement = enemy.movements.find(
              (m) => m.timeElapsed < m.duration
            );
          }
        }
      });
  };
*/

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
	return explosions[:len(explosions)-1]
}

/*
  updateExplosives = (elapsedTimeInMs: number, obstacles: Rectangle[]) => {
    this.explosives = this.explosives.filter((explosive) => {
      let done = explosive.move(elapsedTimeInMs);
      if (!done && explosive.explodesOnContact) {
        const collision = obstacles.find((obstacle) =>
          obstacle.getIntersection(explosive.position)
        );
        done = !!collision;
      }
      if (done) {
        this.detonateExplosive(explosive);
      }
      return !done;
    });
  };
*/

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

  // Check if players collide with explosions / bullets / enemies
  checkEnemiesForDeadlyCollisions = () => {
    this.enemies.forEach((enemy) => {
      if (!enemy.dying) {
        let death = false;
        let killer = null;
        this.explosions.forEach((explosion) => {
          const deadlyExplosion =
            null !== enemy.position.getIntersection(explosion.position);
          if (deadlyExplosion) {
            killer = explosion.playerIdx;
          }
          death = death || deadlyExplosion;
        });

        if (EnemyType.Person === enemy.type) {
          // Bullets, flamethrowers and RPGs are useful only against persons
          this.bullets.forEach((bullet) => {
            const deadlyShot =
              null !== enemy.position.getIntersection(bullet.position);
            if (deadlyShot) {
              killer = bullet.playerIdx;
            }
            death = death || deadlyShot;
          });

          this.explosives = this.explosives.filter((explosive) => {
            let explode =
              explosive.explodesOnContact &&
              explosive.position.getIntersection(enemy.position);
            if (explode) {
              this.detonateExplosive(explosive);
              death = true;
            }
            return !explode;
          });
        }

        if (death) {
          enemy.dying = true;
          enemy.dyingAnimationCountDown = 1;
          if (null !== killer) {
            this.kills[killer] += 1;
          }
          if (EnemyType.Person === enemy.type) {
            SoundBoard.play(
              FX_DEATH_ENEMIES[Math.floor(Math.random() * Math.floor(8))]
            );
            enemy.dyingAnimationCountDown = DURATION_OF_DEATH_ANIMATION_ENEMY;
          }
        }
      }
    });
  };
*/

func (ge *GameEngine) checkIfPlayerReachedLevelGoal() {
	if ge.levelController.GoalReached(ge.playerPositions) {
		ge.won = true
	}
}
