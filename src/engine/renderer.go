package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
)

type Renderer struct {
	canvas              *pixelgl.Canvas
	engine              *GameEngine
	outputArea          *geometry.Rectangle
	playerTileSuppliers []*characters.PlayerSpriteSupplier
	scalingFactor       float64
	window              *pixelgl.Window
}

func NewRenderer(engine *GameEngine, window *pixelgl.Window) *Renderer {
	var playerTileSuppliers = make([]*characters.PlayerSpriteSupplier, 0)
	for _, player := range characters.PlayerController.ConfiguredPlayers() {
		playerTileSuppliers = append(playerTileSuppliers, characters.NewPlayerSpriteSupplier(player))
	}

	var result = &Renderer{engine: engine, playerTileSuppliers: playerTileSuppliers, window: window}
	result.initializeGeometry()
	result.initializeCanvas()
	return result
}

func (r *Renderer) Render(elapsedTimeInMs int64) {
	r.drawBackground()
	r.drawEnemies(elapsedTimeInMs)
	r.drawPlayers(elapsedTimeInMs)
	r.drawBullets()
	r.drawExplosives()
	r.drawExplosions(elapsedTimeInMs)

	r.canvas.Draw(r.window, pixel.IM.Moved(r.canvas.Bounds().Center()))
}

func (r *Renderer) drawBackground() {
	for _, bg := range r.engine.Backgrounds() {
		if nil != bg.Sprite {
			var bgCenter = pixel.Vec{
				X: r.outputArea.X + (r.outputArea.Width / 2) + bg.Offset.X*r.scalingFactor,
				Y: r.outputArea.Y + (r.outputArea.Height / 2) - bg.Offset.Y*r.scalingFactor,
			}
			bg.Sprite.Draw(r.canvas, pixel.IM.Scaled(pixel.V(0, 0), r.scalingFactor).Moved(bgCenter))
		}
	}
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

// initializeGeometry computes the location and size of game area and the scaling factor.
// Should not be called from outside this class.
func (r *Renderer) initializeGeometry() {
	var playerInfoAreaWidth = (r.window.Bounds().W() - r.window.Bounds().H()) / 2
	var result = &geometry.Rectangle{
		X:      playerInfoAreaWidth,
		Y:      0,
		Width:  r.window.Bounds().W() - playerInfoAreaWidth - playerInfoAreaWidth,
		Height: r.window.Bounds().H(),
	}
	r.outputArea = result
	r.scalingFactor = result.Height / ScreenSize
}

// initializeCanvas performs the initialization of the canvas.
// Should not be called from outside this class.
func (r *Renderer) initializeCanvas() {
	r.canvas = pixelgl.NewCanvas(pixel.R(
		r.outputArea.X,
		r.outputArea.Y,
		r.outputArea.X+r.outputArea.Width,
		r.outputArea.Y+r.outputArea.Height,
	))
}
