package game

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
	"retro-carnage/util"
)

const (
	fontSize             = 36
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
	ammunitionChanged bool
	canvas            *pixelgl.Canvas
	changeListener    *util.ChangeListener
	componentArea     *geometry.Rectangle
	livesChanged      bool
	player            *characters.Player
	playerIdx         int
	scoreChanged      bool
	weaponChanged     bool
	window            *pixelgl.Window
}

// newPlayerInfo creates and returns a new instance of playerInfo.
// Use this to construct this component.
func newPlayerInfo(playerIdx int, window *pixelgl.Window) *playerInfo {
	var players = characters.PlayerController.ConfiguredPlayers()
	var player *characters.Player = nil
	if len(players) > playerIdx {
		player = players[playerIdx]
	}
	var result = &playerInfo{
		ammunitionChanged: true,
		canvas:            nil,
		componentArea:     nil,
		livesChanged:      true,
		player:            player,
		playerIdx:         playerIdx,
		scoreChanged:      true,
		weaponChanged:     true,
		window:            window,
	}
	if nil != player {
		var changeListener = &util.ChangeListener{Callback: result.playerPropertyChanged, PropertyNames: []string{}}
		player.AddChangeListener(changeListener)
		result.changeListener = changeListener
	}
	return result
}

// drawToScreen draws this component to screen.
// The content is updated, if necessary.
func (pi *playerInfo) drawToScreen() {
	var sw = &util.StopWatch{Name: "playerInfo:drawToScreen"}
	sw.Start()

	pi.updateCanvas()
	pi.canvas.Draw(pi.window, pixel.IM.Moved(pi.canvas.Bounds().Center()))

	var _ = sw.Stop()
	logging.Trace.Printf(sw.PrintDebugMessage())
}

// updateCanvas updates the in-memory canvas of this component.
// Should not be called from outside this class.
func (pi *playerInfo) updateCanvas() {
	var initialize = nil == pi.componentArea || nil == pi.canvas
	if initialize {
		pi.calculateScreenRect()
		pi.initializeCanvas()
		pi.drawBackground()
		pi.drawPlayerPortrait()
	}

	if pi.scoreChanged {
		pi.drawScore()
		pi.scoreChanged = false
	}

	if pi.weaponChanged || pi.ammunitionChanged {
		pi.drawWeaponBackground()
		pi.drawWeapon()
		pi.drawAmmoCounter()
		pi.ammunitionChanged = false
		pi.weaponChanged = false
	}

	if pi.livesChanged {
		pi.drawLives()
		pi.livesChanged = false
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
	var bottomLeft, topRight = pi.areaForScore()
	var draw = imdraw.New(nil)
	draw.Color = common.Black
	draw.Push(bottomLeft, topRight)
	draw.Rectangle(0)
	draw.Draw(pi.canvas)

	if nil != pi.player {
		var score = fmt.Sprintf("%d", pi.player.Score())
		var textDimensions = fonts.GetTextDimension(fontSize, score)
		var lineX = pi.componentArea.X + (pi.componentArea.Width-textDimensions.X)/2
		var lineY = bottomLeft.Y + (scoreAreaHeight-textDimensions.Y)/2
		var txt = text.New(pixel.V(lineX, lineY), fonts.SizeToFontAtlas[fontSize])
		txt.Color = common.Yellow
		_, _ = fmt.Fprint(txt, score)
		txt.Draw(pi.canvas, pixel.IM)
	}
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

func (pi *playerInfo) playerPropertyChanged(_ interface{}, n string) {
	switch n {
	case characters.PlayerPropertyAmmunition:
		fallthrough
	case characters.PlayerPropertyGrenades:
		pi.ammunitionChanged = true
	case characters.PlayerPropertyLives:
		pi.livesChanged = true
	case characters.PlayerPropertyScore:
		pi.scoreChanged = true
	case characters.PlayerPropertySelectedWeapon:
		pi.weaponChanged = true
	default:
		// ignore changes to other properties
	}
}

func (pi *playerInfo) dispose() {
	if nil != pi.changeListener && nil != pi.player {
		err := pi.player.RemoveChangeListener(pi.changeListener)
		if err != nil {
			logging.Error.Fatalf("playerInfo.dispose: Failed to remove unregistered ChangeListener")
		}
		pi.changeListener = nil
	}
}
