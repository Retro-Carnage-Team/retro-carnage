package game

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
	"retro-carnage/util"
)

// PlayerInfo models the areas left and right of the screen that display information about the players.
// This class can be used to prepare an in-memory canvas with the current state of the player info area that can be
// drawn efficiently when needed.
type playerInfo struct {
	canvas         *pixelgl.Canvas
	componentArea  *geometry.Rectangle
	playerIdx      int
	updateRequired bool
	window         *pixelgl.Window
}

// newPlayerInfo creates and returns a new instance of playerInfo.
// Use this to construct this component.
func newPlayerInfo(playerIdx int, window *pixelgl.Window) *playerInfo {
	return &playerInfo{
		canvas:         nil,
		componentArea:  nil,
		playerIdx:      playerIdx,
		updateRequired: true,
		window:         window,
	}
}

// drawToScreen draws this component to screen.
// The content is updated, if necessary.
func (pi *playerInfo) drawToScreen() {
	var sw = &util.StopWatch{Name: "playerInfo:drawToScreen"}
	sw.Start()

	if pi.updateRequired {
		pi.updateCanvas()
	}
	pi.canvas.Draw(pi.window, pixel.IM.Moved(pi.canvas.Bounds().Center()))

	sw.Stop()
	logging.Trace.Printf(sw.PrintDebugMessage())
}

// updateCanvas updates the in-memory canvas of this component.
// Should not be called from outside this class.
func (pi *playerInfo) updateCanvas() {
	if nil == pi.componentArea {
		pi.calculateScreenRect()
	}
	if nil == pi.canvas {
		pi.initializeCanvas()
	}

	if pi.updateRequired {
		pi.drawBackground()
		pi.drawPlayerPortrait()
		pi.updateRequired = false
	}
}

// calculateScreenRect gets the area of this player info component.
// Should not be called from outside this class.
func (pi *playerInfo) calculateScreenRect() {
	var playerInfoArea = geometry.Rectangle{
		X:      0,
		Y:      0,
		Width:  (pi.window.Bounds().W() - pi.window.Bounds().H()) / 2,
		Height: pi.window.Bounds().H(),
	}
	if 1 == pi.playerIdx {
		playerInfoArea.X = pi.window.Bounds().W() - playerInfoArea.Width
	}
	pi.componentArea = &playerInfoArea
}

// initializeCanvas performs the lazy initialization of the canvas.
// Should not be called from outside this class.
func (pi *playerInfo) initializeCanvas() {
	pi.canvas = pixelgl.NewCanvas(pixel.R(
		pi.componentArea.X,
		pi.componentArea.Y,
		pi.componentArea.X+pi.componentArea.Width,
		pi.componentArea.Y+pi.componentArea.Height,
	))
}

// drawBackground draws the background to the in-memory canvas.
// Should not be called from outside this class.
func (pi *playerInfo) drawBackground() {
	var backgroundSprite = assets.SpriteRepository.Get(playerInfoBgPath)
	var spriteBounds = backgroundSprite.Picture().Bounds()
	var offsetX = pi.componentArea.X + spriteBounds.W()/2
	for {
		var offsetY = spriteBounds.Max.Y / 2
		for offsetY < pi.window.Bounds().H()+spriteBounds.Max.Y/2 {
			var movedSpriteBounds = pixel.V(offsetX, offsetY)
			backgroundSprite.Draw(pi.canvas, pixel.IM.Moved(movedSpriteBounds))
			offsetY += spriteBounds.Max.Y
		}
		if offsetX >= pi.componentArea.X+pi.componentArea.Width {
			break
		} else {
			offsetX += spriteBounds.W()
		}
	}
}

// drawBackground draws player portrait to the in-memory canvas.
// Should not be called from outside this class.
// The player's portrait will have a top margin of 15 px and a height of <window height> / 4.
func (pi *playerInfo) drawPlayerPortrait() {
	var playerPortraitPath = fmt.Sprintf("images/player-%d/portrait.png", pi.playerIdx)
	var playerPortraitSprite = assets.SpriteRepository.Get(playerPortraitPath)
	var scalingFactor = (pi.window.Bounds().H() / 4) / playerPortraitSprite.Picture().Bounds().H()
	var location = pixel.Vec{
		X: pi.componentArea.X + (pi.componentArea.Width / 2),
		Y: pi.componentArea.Height - 15 - playerPortraitSprite.Picture().Bounds().H()*scalingFactor/2,
	}
	playerPortraitSprite.Draw(pi.canvas, pixel.IM.Scaled(pixel.V(0, 0), scalingFactor).Moved(location))
}
