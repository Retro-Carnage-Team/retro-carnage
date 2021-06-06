package engine

import (
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
)

type Renderer struct {
	engine              *GameEngine
	playerTileSuppliers []*characters.PlayerSpriteSupplier
	window              *pixelgl.Window
}

func NewRenderer(engine *GameEngine, window *pixelgl.Window) *Renderer {
	var playerTileSuppliers = make([]*characters.PlayerSpriteSupplier, 0)
	for _, player := range characters.PlayerController.ConfiguredPlayers() {
		playerTileSuppliers = append(playerTileSuppliers, characters.NewPlayerSpriteSupplier(player))
	}
	return &Renderer{engine: engine, playerTileSuppliers: playerTileSuppliers, window: window}
}

func (r *Renderer) Render(elapsedTimeInMs int64) {
	// this.ctx.clearRect(0, 0, SCREEN_SIZE, SCREEN_SIZE);

	r.drawBackground()
	r.drawEnemies(elapsedTimeInMs)
	r.drawPlayers(elapsedTimeInMs)
	r.drawBullets()
	r.drawExplosives()
	r.drawExplosions(elapsedTimeInMs)
}

func (r *Renderer) drawBackground() {
	//const backgroundRect = new Rectangle(0, 0, SCREEN_SIZE, SCREEN_SIZE);
	//this.engine.getBackgrounds().forEach((bg) => {
	//	const translatedPosition = bg.translate(backgroundRect);
	//	const canvas = bg.getCanvas();
	//	if (canvas && this.ctx) {
	//		this.ctx.drawImage(canvas, translatedPosition.x, translatedPosition.y);
	//	}
	//});
}

func (r *Renderer) drawEnemies(elapsedTimeInMs int64) {
	//this.engine.enemies.forEach((enemy) => {
	//	const tile = enemy.tileSupplier.getTile(elapsedTimeInMs, enemy);
	//	if (tile) {
	//		const translatedPosition = tile.translate(enemy.position);
	//		const canvas = tile.getCanvas();
	//		if (canvas && this.ctx) {
	//			this.ctx.drawImage(canvas, translatedPosition.x, translatedPosition.y);
	//		}
	//	}
	//});
}

func (r *Renderer) drawPlayers(elapsedTimeInMs int64) {
	//const positions = this.engine.playerPositions;
	//const behaviors = this.engine.playerBehaviors;
	//PlayerController.getRemainingPlayers().forEach((player) => {
	//	const tile = this.playerTileSuppliers[player.index].getTile(elapsedTimeInMs, behaviors[player.index]);
	//	if (tile) {
	//		const translatedPosition = tile.translate(positions[player.index]);
	//		const canvas = tile.getCanvas();
	//		if (canvas && this.ctx) {
	//			this.ctx.drawImage(canvas, translatedPosition.x, translatedPosition.y);
	//			// this.drawDebugRect(positions[player.index]);
	//		}
	//	}
	//});
}

func (r *Renderer) drawBullets() {
	//const ctx = this.ctx
	//if ctx {
	//	ctx.fillStyle = BULLET_COLOR
	//	this.engine.bullets.forEach((bullet) => {
	//		ctx.fillRect(bullet.position.x, bullet.position.y, bullet.position.width, bullet.position.height)
	//	})
	//}
}

func (r *Renderer) drawExplosives() {
	//this.engine.explosives.forEach((explosive) => {
	//	const canvas = explosive.tileSupplier.getTile().getCanvas();
	//	if (canvas && this.ctx) {
	//		this.ctx.drawImage(canvas, explosive.position.x, explosive.position.y);
	//	}
	//});
}

func (r *Renderer) drawExplosions(elapsedTimeInMs int64) {
	//this.engine.explosions.forEach((explosion) => {
	//	const tile = explosion.tileSupplier.getTile(elapsedTimeInMs);
	//	const canvas = tile.getCanvas();
	//	if (canvas && this.ctx) {
	//		this.ctx.drawImage(canvas, explosion.position.x, explosion.position.y);
	//	}
	//});
}

func (r *Renderer) drawDebugRect(rect geometry.Rectangle) {
	//if (this.ctx) {
	//	this.ctx.strokeStyle = "orange";
	//	this.ctx.strokeRect(rect.x, rect.y, rect.width, rect.height);
	//}
}
