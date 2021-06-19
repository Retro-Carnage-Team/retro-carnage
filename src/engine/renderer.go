package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
)

// Renderer is used to render the current state of an GameEngine to screen
type Renderer struct {
	canvas                *pixelgl.Canvas
	engine                *GameEngine
	outputArea            *geometry.Rectangle
	playerSpriteSuppliers []*characters.PlayerSpriteSupplier
	scalingFactor         float64
	window                *pixelgl.Window
}

// NewRenderer creates and initializes a new Renderer instance.
func NewRenderer(engine *GameEngine, window *pixelgl.Window) *Renderer {
	var playerSpriteSuppliers = make([]*characters.PlayerSpriteSupplier, 0)
	for _, player := range characters.PlayerController.ConfiguredPlayers() {
		playerSpriteSuppliers = append(playerSpriteSuppliers, characters.NewPlayerSpriteSupplier(player))
	}

	var result = &Renderer{engine: engine, playerSpriteSuppliers: playerSpriteSuppliers, window: window}
	result.initializeGeometry()
	result.initializeCanvas()
	return result
}

// Render draws the state of the GameEngine to screen. The parameter elapsedTimeInMs is used to pick the right sprites
// for animated objects, like the player character or enemies.
func (r *Renderer) Render(elapsedTimeInMs int64) {
	r.drawBackground()
	r.drawEnemies(elapsedTimeInMs)
	r.drawPlayers(elapsedTimeInMs)
	r.drawBullets()
	r.drawExplosives()
	r.drawExplosions(elapsedTimeInMs)

	r.canvas.Draw(r.window, pixel.IM.Moved(r.canvas.Bounds().Center()))
}

// drawBackground draws the background image of the current mission section onto the in-memory canvas.
// Do not call from outside this class.
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

// drawEnemies draws sprites for each of the currently visible enemies onto the in-memory canvas.
// Do not call from outside this class.
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

// drawPlayers draws sprites for each of the players onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawPlayers(elapsedTimeInMs int64) {
	var outputAreaInverseRoot = pixel.V(r.outputArea.X, r.outputArea.Y+r.outputArea.Height)
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var behavior = r.engine.playerBehaviors[player.Index()]
		var spriteWithOffset = r.playerSpriteSuppliers[player.Index()].Sprite(elapsedTimeInMs, behavior)
		if nil != spriteWithOffset {
			var position = r.engine.playerPositions[player.Index()]
			var spriteCenter = pixel.Vec{
				X: spriteWithOffset.Offset.X + position.X + spriteWithOffset.Sprite.Picture().Bounds().W()/2,
				Y: -1 * (spriteWithOffset.Offset.Y + position.Y + spriteWithOffset.Sprite.Picture().Bounds().H()/2),
			}.Scaled(r.scalingFactor).Add(outputAreaInverseRoot)
			var matrix = pixel.IM.Scaled(pixel.V(0, 0), r.scalingFactor).Moved(spriteCenter)
			spriteWithOffset.Sprite.Draw(r.canvas, matrix)
		} else {
			logging.Warning.Printf("Player spriteWithOffset missing for player %d", player.Index())
		}
	}
}

// drawBullets draws the flying bullets onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawBullets() {
	//const ctx = this.ctx
	//if ctx {
	//	ctx.fillStyle = BULLET_COLOR
	//	this.engine.bullets.forEach((bullet) => {
	//		ctx.fillRect(bullet.position.x, bullet.position.y, bullet.position.width, bullet.position.height)
	//	})
	//}
}

// drawExplosives draws the flying explosives (grenades, RPGs) onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawExplosives() {
	//this.engine.explosives.forEach((explosive) => {
	//	const canvas = explosive.tileSupplier.getTile().getCanvas();
	//	if (canvas && this.ctx) {
	//		this.ctx.drawImage(canvas, explosive.position.x, explosive.position.y);
	//	}
	//});
}

// drawExplosions draws sprites for animated explosions onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawExplosions(elapsedTimeInMs int64) {
	//this.engine.explosions.forEach((explosion) => {
	//	const tile = explosion.tileSupplier.getTile(elapsedTimeInMs);
	//	const canvas = tile.getCanvas();
	//	if (canvas && this.ctx) {
	//		this.ctx.drawImage(canvas, explosion.position.x, explosion.position.y);
	//	}
	//});
}

// drawDebugRect draws a given geometry.Rectangle onto the in-memory canvas.
// Useful for debugging.
// Do not call from outside this class.
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
