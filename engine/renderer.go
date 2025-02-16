package engine

import (
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/logging"
	"retro-carnage/ui/common"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
	"github.com/Retro-Carnage-Team/pixel2/ext/imdraw"
)

// Renderer is used to render the current state of an GameEngine to screen
type Renderer struct {
	canvas                *opengl.Canvas
	engine                *GameEngine
	outputArea            *geometry.Rectangle
	playerSpriteSuppliers []*graphics.PlayerSpriteSupplier
	scalingFactor         float64
	window                *opengl.Window
}

// NewRenderer creates and initializes a new Renderer instance.
func NewRenderer(engine *GameEngine, window *opengl.Window) *Renderer {
	var playerSpriteSuppliers = make([]*graphics.PlayerSpriteSupplier, 0)
	for _, player := range characters.PlayerController.ConfiguredPlayers() {
		var behavior = engine.playerBehaviors[player.Index()]
		var visualAdapter = characters.NewPlayerVisualsAdapter(behavior)
		playerSpriteSuppliers = append(playerSpriteSuppliers, graphics.NewPlayerSpriteSupplier(visualAdapter, DurationOfInvincibilityState))
	}

	var result = &Renderer{engine: engine, playerSpriteSuppliers: playerSpriteSuppliers, window: window}
	result.initializeGeometry()
	result.initializeCanvas()
	return result
}

// Render draws the state of the GameEngine to screen.
// The parameter elapsedTimeInMs is used to pick the right sprites for animated objects, like the player character or
// enemies. Returns a canvas that displays the current game state. Do not modify this canvas outside of this object.
func (r *Renderer) Render(elapsedTimeInMs int64) *opengl.Canvas {
	r.drawBackground()
	r.drawBurnMarks()
	r.drawEnemies(elapsedTimeInMs)
	r.drawPlayers(elapsedTimeInMs)
	r.drawBullets()
	r.drawExplosives(elapsedTimeInMs)
	r.drawExplosions(elapsedTimeInMs)

	r.canvas.Draw(r.window, pixel.IM.Moved(r.canvas.Bounds().Center()))
	return r.canvas
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
	for _, enemy := range r.engine.enemies {
		if !enemy.Type.IsVisible() {
			continue
		}

		var spriteWithOffset = enemy.SpriteSupplier.Sprite(elapsedTimeInMs)
		if nil != spriteWithOffset {
			r.drawSpriteToCanvas(spriteWithOffset, enemy.Position())
		} else {
			logging.Warning.Printf("Enemy spriteWithOffset missing for enemy %s", enemy.Skin)
		}
	}
}

// drawPlayers draws sprites for each of the players onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawPlayers(elapsedTimeInMs int64) {
	for _, player := range characters.PlayerController.RemainingPlayers() {
		var behavior = r.engine.playerBehaviors[player.Index()]
		var visualAdapter = characters.NewPlayerVisualsAdapter(behavior)
		var spriteWithOffset = r.playerSpriteSuppliers[player.Index()].Sprite(elapsedTimeInMs, visualAdapter)
		// spriteWithOffset will be null a couple of times per second when player is invincible.
		if nil != spriteWithOffset {
			r.drawSpriteToCanvas(spriteWithOffset, r.engine.playerPositions[player.Index()])
		}
	}
}

// drawBullets draws the flying bullets onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawBullets() {
	if 0 < len(r.engine.bullets) {
		var draw = imdraw.New(nil)
		draw.Color = common.White
		var outputAreaInverseRoot = pixel.V(r.outputArea.X, r.outputArea.Y+r.outputArea.Height)
		for _, bullet := range r.engine.bullets {
			draw.Push(
				pixel.Vec{
					X: bullet.Position().X,
					Y: -1 * bullet.Position().Y,
				}.Scaled(r.scalingFactor).Add(outputAreaInverseRoot),
				pixel.Vec{
					X: bullet.Position().X + bullet.Position().Width,
					Y: -1 * (bullet.Position().Y + bullet.Position().Height),
				}.Scaled(r.scalingFactor).Add(outputAreaInverseRoot),
			)
			draw.Rectangle(0)
		}
		draw.Draw(r.canvas)
	}
}

// drawExplosives draws the flying explosives (grenades, RPGs) onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawExplosives(elapsedTimeInMs int64) {
	for _, explosive := range r.engine.explosives {
		var spriteWOffset = explosive.SpriteSupplier.Sprite(elapsedTimeInMs)
		if nil != spriteWOffset {
			r.drawSpriteToCanvas(spriteWOffset, explosive.position)
		} else {
			logging.Warning.Printf("Explosive sprite missing")
		}
	}
}

// drawExplosions draws sprites for animated explosions onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawExplosions(elapsedTimeInMs int64) {
	for _, explosion := range r.engine.explosions {
		var spriteWOffset = explosion.SpriteSupplier.Sprite(elapsedTimeInMs)
		if nil != spriteWOffset {
			r.drawSpriteToCanvas(spriteWOffset, explosion.Position())
		} else {
			logging.Warning.Printf("Explosion sprite missing")
		}
	}
}

// drawBurnMarks draws sprites for burn marks onto the in-memory canvas.
// Do not call from outside this class.
func (r *Renderer) drawBurnMarks() {
	for _, burnMark := range r.engine.burnMarks {
		var spriteWOffset = burnMark.SpriteSupplier.Sprite()
		if nil != spriteWOffset {
			r.drawSpriteToCanvas(spriteWOffset, burnMark.Position())
		} else {
			logging.Warning.Printf("Burn mark sprite missing")
		}
	}
}

// drawDebugRect draws a given geometry.Rectangle onto the in-memory canvas. Useful for debugging.
// Do not call from outside this class.
/*
func (r *Renderer) drawDebugRect(rect *geometry.Rectangle) {
	if nil == rect {
		return
	}

	var draw = imdraw.New(nil)
	draw.Color = common.Orange
	var outputAreaInverseRoot = pixel.V(r.outputArea.X, r.outputArea.Y+r.outputArea.Height)
	draw.Push(
		pixel.Vec{X: rect.X, Y: -1 * rect.Y}.Scaled(r.scalingFactor).Add(outputAreaInverseRoot),
		pixel.Vec{
			X: rect.X + rect.Width,
			Y: -1 * (rect.Y + rect.Height),
		}.Scaled(r.scalingFactor).Add(outputAreaInverseRoot),
	)
	draw.Rectangle(3)
	draw.Draw(r.canvas)
}
*/

// draws a given sprite to the given position on canvas.
func (r *Renderer) drawSpriteToCanvas(spriteWithOffset *graphics.SpriteWithOffset, position *geometry.Rectangle) {
	var outputAreaInverseRoot = pixel.V(r.outputArea.X, r.outputArea.Y+r.outputArea.Height)
	var spriteCenter = pixel.Vec{
		X: spriteWithOffset.Offset.X + position.X + spriteWithOffset.Sprite.Picture().Bounds().W()/2,
		Y: -1 * (spriteWithOffset.Offset.Y + position.Y + spriteWithOffset.Sprite.Picture().Bounds().H()/2),
	}.Scaled(r.scalingFactor).Add(outputAreaInverseRoot)
	var matrix = pixel.IM.Scaled(pixel.V(0, 0), r.scalingFactor).Moved(spriteCenter)

	if nil != spriteWithOffset.ColorMask {
		spriteWithOffset.Sprite.DrawColorMask(r.canvas, matrix, *spriteWithOffset.ColorMask)
	} else {
		spriteWithOffset.Sprite.Draw(r.canvas, matrix)
	}
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
	r.canvas = opengl.NewCanvas(pixel.R(
		r.outputArea.X,
		r.outputArea.Y,
		r.outputArea.X+r.outputArea.Width,
		r.outputArea.Y+r.outputArea.Height,
	))
}
