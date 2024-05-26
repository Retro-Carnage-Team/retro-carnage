package game

import (
	"fmt"
	"math"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
	"retro-carnage/util"

	"github.com/Retro-Carnage-Team/pixel"
	"github.com/Retro-Carnage-Team/pixel/imdraw"
	"github.com/Retro-Carnage-Team/pixel/pixelgl"
	"github.com/Retro-Carnage-Team/pixel/text"
)

const (
	err_msg_sprint_missing = "Unable to find sprite: %s"
	fontSize               = 36
	innerMargin            = 15
	livesAreaHeight        = 50
	missingInAction        = "M.I.A."
	playerInfoBgPath       = "images/other/player-info-bg.png"
	playerLifePath         = "images/player-%d/life.png"
	playerPortraitPath     = "images/player-%d/portrait.png"
	scoreAreaHeight        = 50
	weaponBackgroundPath   = "images/other/weapon-bg.jpg"
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

// draw draws this component to the given pixel.Target.
// The content is updated, if necessary.
func (pi *playerInfo) draw(target pixel.Target) {
	pi.updateCanvas()
	pi.canvas.Draw(target, pixel.IM.Moved(pi.canvas.Bounds().Center()))
}

// updateCanvas updates the in-memory canvas of this component.
// Should not be called from outside this class.
func (pi *playerInfo) updateCanvas() {
	var initialize = nil == pi.componentArea || nil == pi.canvas
	if initialize {
		pi.calculateComponentArea(pi.window)
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

// calculateComponentArea gets the area of this player info component.
// Should not be called from outside this class.
func (pi *playerInfo) calculateComponentArea(window *pixelgl.Window) {
	var playerInfoArea = geometry.Rectangle{
		X:      0,
		Y:      0,
		Width:  (window.Bounds().W() - window.Bounds().H()) / 2,
		Height: window.Bounds().H(),
	}
	if pi.playerIdx == 1 {
		playerInfoArea.X = window.Bounds().W() - playerInfoArea.Width
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
	if nil == backgroundSprite {
		logging.Warning.Printf(err_msg_sprint_missing, playerInfoBgPath)
		return
	}

	var spriteBounds = backgroundSprite.Picture().Bounds()
	var offsetX = pi.componentArea.X + spriteBounds.W()/2
	for {
		var offsetY = spriteBounds.Max.Y / 2
		for offsetY < pi.componentArea.Height+spriteBounds.Max.Y/2 {
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
	if nil == playerPortraitSprite {
		logging.Warning.Printf(err_msg_sprint_missing, playerPortraitPath)
		return
	}

	var scalingFactor = (pi.componentArea.Height / 4) / playerPortraitSprite.Picture().Bounds().H()
	var location = pixel.Vec{
		X: pi.componentArea.X + (pi.componentArea.Width / 2),
		Y: pi.componentArea.Height - innerMargin - playerPortraitSprite.Picture().Bounds().H()*scalingFactor/2,
	}
	playerPortraitSprite.Draw(pi.canvas, pixel.IM.Scaled(pixel.V(0, 0), scalingFactor).Moved(location))
}

// drawScore draws the current score of the player onto a specific section of the in-memory canvas.
// Should not be called from outside this class.
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
		var position = pixel.Vec{
			X: pi.componentArea.X + (pi.componentArea.Width-textDimensions.X)/2,
			Y: bottomLeft.Y + (scoreAreaHeight-textDimensions.Y)/2,
		}
		var txt = text.New(position, fonts.SizeToFontAtlas[fontSize])
		txt.Color = common.Yellow
		_, _ = fmt.Fprint(txt, score)
		txt.Draw(pi.canvas, pixel.IM)
	}
}

// drawWeaponBackground draws the background of the weapon area (weapon / ammo counter) onto a specific section of the
// in-memory canvas. Should not be called from outside this class.
func (pi *playerInfo) drawWeaponBackground() {
	var bottomLeft, topRight = pi.areaForWeapon()
	var weaponBackgroundSprite = assets.SpriteRepository.Get(weaponBackgroundPath)
	if nil == weaponBackgroundSprite {
		logging.Warning.Printf(err_msg_sprint_missing, weaponBackgroundPath)
		return
	}

	var scale = pixel.Vec{
		X: (topRight.X - bottomLeft.X) / weaponBackgroundSprite.Picture().Bounds().W(),
		Y: (topRight.Y - bottomLeft.Y) / weaponBackgroundSprite.Picture().Bounds().H(),
	}
	var movement = pixel.Vec{
		X: (topRight.X + bottomLeft.X) / 2,
		Y: (topRight.Y + bottomLeft.Y) / 2,
	}
	var matrix = pixel.IM.ScaledXY(pixel.V(0, 0), scale).Moved(movement)
	weaponBackgroundSprite.Draw(pi.canvas, matrix)

	var draw = imdraw.New(nil)
	draw.Color = common.Black
	draw.Push(bottomLeft, topRight)
	draw.Rectangle(5)
	draw.Draw(pi.canvas)
}

// drawWeapon draws the rotated image of the player's currently selected weapon onto a specific section of the in-memory
// canvas. Should not be called from outside this class.
func (pi *playerInfo) drawWeapon() {
	if nil != pi.player && (nil != pi.player.SelectedWeapon() || nil != pi.player.SelectedGrenade()) {
		var imagePath string
		if nil != pi.player.SelectedWeapon() {
			imagePath = pi.player.SelectedWeapon().ImageRotated
		} else {
			imagePath = pi.player.SelectedGrenade().ImageRotated
		}
		var weaponSprite = assets.SpriteRepository.Get(imagePath)
		if nil == weaponSprite {
			logging.Warning.Printf(err_msg_sprint_missing, imagePath)
			return
		}

		var bottomLeft, topRight = pi.areaForWeapon()
		var scaleX = (topRight.X - bottomLeft.X - (2 * innerMargin)) / weaponSprite.Picture().Bounds().W()
		var scaleY = (topRight.Y - bottomLeft.Y - (3 * innerMargin) - scoreAreaHeight) / weaponSprite.Picture().Bounds().H()
		var minScale = math.Min(scaleX, scaleY)
		var scale = pixel.V(minScale, minScale)
		var movement = pixel.Vec{
			X: pi.componentArea.X + innerMargin + (topRight.X-bottomLeft.X)/2,
			Y: pi.componentArea.Y + (3 * innerMargin) + scoreAreaHeight + (topRight.Y-bottomLeft.Y)/2,
		}
		var matrix = pixel.IM.ScaledXY(pixel.V(0, 0), scale).Moved(movement)
		weaponSprite.Draw(pi.canvas, matrix)
	}
}

// drawAmmoCounter draws the number of bullets for the player's currently selected weapon onto a specific section of the
// in-memory canvas. Should not be called from outside this class.
func (pi *playerInfo) drawAmmoCounter() {
	if nil != pi.player && (nil != pi.player.SelectedWeapon() || nil != pi.player.SelectedGrenade()) {
		var bottomLeft, _ = pi.areaForWeapon()
		var ammo = fmt.Sprintf("%d", pi.player.AmmunitionForSelectedWeapon())
		var textDimensions = fonts.GetTextDimension(fontSize, ammo)
		var position = pixel.Vec{
			X: pi.componentArea.X + (pi.componentArea.Width-textDimensions.X)/2,
			Y: bottomLeft.Y + innerMargin,
		}
		var txt = text.New(position, fonts.SizeToFontAtlas[fontSize])
		txt.Color = common.Black
		_, _ = fmt.Fprint(txt, ammo)
		txt.Draw(pi.canvas, pixel.IM)
	}
}

// drawLives draws life sprites for each of the player's lives onto a specific section of the in-memory canvas.
// Should not be called from outside this class.
func (pi *playerInfo) drawLives() {
	var bottomLeft, topRight = pi.areaForLives()
	var draw = imdraw.New(nil)
	draw.Color = common.Black
	draw.Push(bottomLeft, topRight)
	draw.Rectangle(0)
	draw.Draw(pi.canvas)

	if nil != pi.player {
		pi.drawLivesForPlayer(bottomLeft)
	} else {
		pi.drawLivesForMissingInAction(bottomLeft)
	}
}

func (pi *playerInfo) drawLivesForPlayer(bottomLeft pixel.Vec) {
	var lifeSpritePath = fmt.Sprintf(playerLifePath, pi.playerIdx)
	var lifeSprite = assets.SpriteRepository.Get(lifeSpritePath)
	if nil == lifeSprite {
		logging.Warning.Printf(err_msg_sprint_missing, lifeSpritePath)
		return
	}

	var scale = (livesAreaHeight - 2) / lifeSprite.Picture().Bounds().H()
	var scaledSpriteHeight = lifeSprite.Picture().Bounds().H() * scale
	var scaledSpriteWidth = lifeSprite.Picture().Bounds().W() * scale
	var width = float64(pi.player.Lives()) * scaledSpriteWidth
	if 1 < pi.player.Lives() {
		width += float64((pi.player.Lives() - 1) * 5)
	}

	var left = pi.componentArea.X + (pi.componentArea.Width-width)/2 + scaledSpriteWidth/2
	for i := 0; i < pi.player.Lives(); i++ {
		var matrix = pixel.IM.
			Scaled(pixel.V(0, 0), scale).
			Moved(pixel.V(left, bottomLeft.Y+1+scaledSpriteHeight/2))
		lifeSprite.Draw(pi.canvas, matrix)
		left += scaledSpriteWidth + 5
	}
}

func (pi *playerInfo) drawLivesForMissingInAction(bottomLeft pixel.Vec) {
	var textDimensions = fonts.GetTextDimension(fontSize, missingInAction)
	var position = pixel.Vec{
		X: pi.componentArea.X + (pi.componentArea.Width-textDimensions.X)/2,
		Y: bottomLeft.Y + (livesAreaHeight-textDimensions.Y)/2,
	}
	var txt = text.New(position, fonts.SizeToFontAtlas[fontSize])
	txt.Color = common.Yellow
	_, _ = fmt.Fprint(txt, missingInAction)
	txt.Draw(pi.canvas, pixel.IM)
}

// areaForLives returns the two points (bottom left, top right) that define the rectangular area in which the extra
// lives are displayed.
// Should not be called from outside this class.
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
// Should not be called from outside this class.
func (pi *playerInfo) areaForScore() (pixel.Vec, pixel.Vec) {
	var bottomLeft = pixel.Vec{
		X: pi.componentArea.X + innerMargin,
		Y: pi.componentArea.Height - innerMargin - (pi.componentArea.Height / 4) - innerMargin - scoreAreaHeight,
	}
	var topRight = pixel.Vec{
		X: pi.componentArea.X + pi.componentArea.Width - innerMargin,
		Y: bottomLeft.Y + scoreAreaHeight,
	}
	return bottomLeft, topRight
}

// areaForWeapon returns the two points (bottom left, top right) that define the rectangular area in which the weapon
// and ammo count are displayed.
// Should not be called from outside this class.
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

// playerPropertyChanged is a callback function called by a Player when one of it's properties changed.
// Should not be called from outside this class.
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

// dispose frees all the resources blocked by this component.
// Should be called when you need a component instance no longer.
func (pi *playerInfo) dispose() {
	if nil != pi.changeListener && nil != pi.player {
		err := pi.player.RemoveChangeListener(pi.changeListener)
		if err != nil {
			logging.Error.Fatalf("playerInfo.dispose: Failed to remove unregistered ChangeListener")
		}
		pi.changeListener = nil
	}
}
