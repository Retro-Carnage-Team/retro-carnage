package game

import (
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
	playerIdx      int
	screenRect     *geometry.Rectangle
	updateRequired bool
	window         *pixelgl.Window
}

// newPlayerInfo creates and returns a new instance of playerInfo. Use this to construct this component.
func newPlayerInfo(playerIdx int, window *pixelgl.Window) *playerInfo {
	return &playerInfo{
		canvas:         nil,
		playerIdx:      playerIdx,
		screenRect:     nil,
		updateRequired: true,
		window:         window,
	}
}

// drawToScreen draws this component to screen. The content is updated, if necessary.
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

// updateCanvas updates the in-memory canvas of this component. Should not be called from outside this class.
func (pi *playerInfo) updateCanvas() {
	if nil == pi.screenRect {
		pi.calculateScreenRect()
	}
	if nil == pi.canvas {
		pi.initializeCanvas()
	}

	if pi.updateRequired {
		pi.drawBackground()
		pi.updateRequired = false
	}
}

// calculateScreenRect gets the area of this player info component. Should not be called from outside this class.
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
	pi.screenRect = &playerInfoArea
}

// initializeCanvas performs the lazy initialization of the canvas. Should not be called from outside this class.
func (pi *playerInfo) initializeCanvas() {
	pi.canvas = pixelgl.NewCanvas(pixel.R(
		pi.screenRect.X,
		pi.screenRect.Y,
		pi.screenRect.X+pi.screenRect.Width,
		pi.screenRect.Y+pi.screenRect.Height,
	))
}

// drawBackground draws the background to the in-memory canvas. Should not be called from outside this class.
func (pi *playerInfo) drawBackground() {
	var backgroundSprite = assets.SpriteRepository.Get(playerInfoBgPath)
	var spriteBounds = backgroundSprite.Picture().Bounds()
	var offsetX = pi.screenRect.X + spriteBounds.W()/2
	for {
		var offsetY = spriteBounds.Max.Y / 2
		for offsetY < pi.window.Bounds().H()+spriteBounds.Max.Y/2 {
			var movedSpriteBounds = pixel.V(offsetX, offsetY)
			backgroundSprite.Draw(pi.canvas, pixel.IM.Moved(movedSpriteBounds))
			offsetY += spriteBounds.Max.Y
		}
		if offsetX >= pi.screenRect.X+pi.screenRect.Width {
			break
		} else {
			offsetX += spriteBounds.W()
		}
	}
}
