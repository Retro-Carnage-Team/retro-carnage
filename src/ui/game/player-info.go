package game

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/util"
)

const (
	innerMargin          = 15
	livesAreaHeight      = 50
	playerInfoBgPath     = "images/other/player-info-bg.png"
	playerPortraitPath   = "images/player-%d/portrait.png"
	scoreAreaHeight      = 50
	weaponBackgroundPath = "images/other/weapon-bg.jpg"
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
		pi.drawScore()
		pi.drawWeaponBackground()
		pi.drawWeapon()
		pi.drawAmmoCounter()
		pi.drawLives()
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
	var playerPortraitPath = fmt.Sprintf(playerPortraitPath, pi.playerIdx)
	var playerPortraitSprite = assets.SpriteRepository.Get(playerPortraitPath)
	var scalingFactor = (pi.window.Bounds().H() / 4) / playerPortraitSprite.Picture().Bounds().H()
	var location = pixel.Vec{
		X: pi.componentArea.X + (pi.componentArea.Width / 2),
		Y: pi.componentArea.Height - innerMargin - playerPortraitSprite.Picture().Bounds().H()*scalingFactor/2,
	}
	playerPortraitSprite.Draw(pi.canvas, pixel.IM.Scaled(pixel.V(0, 0), scalingFactor).Moved(location))
}

func (pi *playerInfo) drawScore() {
	var draw = imdraw.New(nil)
	draw.Color = common.Black
	draw.Push(pi.areaForScore())
	draw.Rectangle(0)
	draw.Draw(pi.canvas)
}

func (pi *playerInfo) drawWeaponBackground() {
	var bottomLeft, topRight = pi.areaForWeapon()
	var weaponBackgroundSprite = assets.SpriteRepository.Get(weaponBackgroundPath)
	var scaleX = (topRight.X - bottomLeft.X) / weaponBackgroundSprite.Picture().Bounds().W()
	var scaleY = (topRight.Y - bottomLeft.Y) / weaponBackgroundSprite.Picture().Bounds().H()
	var centerX = (topRight.X + bottomLeft.X) / 2
	var centerY = (topRight.Y + bottomLeft.Y) / 2
	weaponBackgroundSprite.Draw(pi.canvas, pixel.IM.
		ScaledXY(pixel.V(0, 0), pixel.V(scaleX, scaleY)).
		Moved(pixel.V(centerX, centerY)))

	var draw = imdraw.New(nil)
	draw.Color = common.Black
	draw.Push(bottomLeft, topRight)
	draw.Rectangle(5)
	draw.Draw(pi.canvas)
}

func (pi *playerInfo) drawWeapon() {
}

func (pi *playerInfo) drawAmmoCounter() {
}

func (pi *playerInfo) drawLives() {
	var draw = imdraw.New(nil)
	draw.Color = common.Black
	draw.Push(pi.areaForLives())
	draw.Rectangle(0)
	draw.Draw(pi.canvas)
}

// areaForLives returns the two points (bottom left, top right) that define the rectangular area in which the extra
// lives are displayed.
func (pi *playerInfo) areaForLives() (pixel.Vec, pixel.Vec) {
	var bottomLeft = pixel.Vec{
		X: pi.componentArea.X + innerMargin,
		Y: innerMargin,
	}
	var topRight = pixel.Vec{
		X: pi.componentArea.X + pi.componentArea.Width - innerMargin,
		Y: innerMargin + livesAreaHeight,
	}
	return bottomLeft, topRight
}

// areaForScore returns the two points (bottom left, top right) that define the rectangular area in which the score gets
// displayed.
func (pi *playerInfo) areaForScore() (pixel.Vec, pixel.Vec) {
	var bottomLeft = pixel.Vec{
		X: pi.componentArea.X + innerMargin,
		Y: pi.window.Bounds().H() - innerMargin - (pi.window.Bounds().H() / 4) - innerMargin - scoreAreaHeight,
	}
	var topRight = pixel.Vec{
		X: pi.componentArea.X + pi.componentArea.Width - innerMargin,
		Y: bottomLeft.Y + scoreAreaHeight,
	}
	return bottomLeft, topRight
}

// areaForWeapon returns the two points (bottom left, top right) that define the rectangular area in which the weapon
// and ammo count are displayed.
func (pi *playerInfo) areaForWeapon() (pixel.Vec, pixel.Vec) {
	var _, livesTR = pi.areaForLives()
	var scoreBL, _ = pi.areaForScore()
	var bottomLeft = pixel.Vec{
		X: pi.componentArea.X + innerMargin,
		Y: livesTR.Y + innerMargin,
	}
	var topRight = pixel.Vec{
		X: pi.componentArea.X + pi.componentArea.Width - innerMargin,
		Y: scoreBL.Y - innerMargin,
	}
	return bottomLeft, topRight
}
