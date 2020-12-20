package shop

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/ui/common"
)

const backgroundImagePath = "./images/backgrounds/shop.jpg"
const bottomBarHeight = 70
const itemMargin = 10.0

type Screen struct {
	backgroundImageSprite  *pixel.Sprite
	inputController        input.Controller
	itemNameToClientSprite map[string]*pixel.Sprite
	items                  []inventoryItem
	PlayerIdx              int
	screenChangeRequired   common.ScreenChangeCallback
	window                 *pixelgl.Window
}

func (s *Screen) SetInputController(controller input.Controller) {
	s.inputController = controller
}

func (s *Screen) SetScreenChangeCallback(callback common.ScreenChangeCallback) {
	s.screenChangeRequired = callback
}

func (s *Screen) SetWindow(window *pixelgl.Window) {
	s.window = window
}

func (s *Screen) SetUp() {
	s.backgroundImageSprite = common.LoadSprite(backgroundImagePath)

	s.items = getAllInventoryItems()
	s.itemNameToClientSprite = make(map[string]*pixel.Sprite)
	for _, item := range s.items {
		s.itemNameToClientSprite[item.Name()] = common.LoadSprite(item.Image())
	}
}

func (s *Screen) Update(_ int64) {
	s.drawBackground()
	s.drawItems()
	s.drawBottomBar()
}

func (s *Screen) TearDown() {}

func (s *Screen) String() string {
	if 0 == s.PlayerIdx {
		return string(common.ShopP1)
	}
	return string(common.ShopP2)
}

func (s *Screen) drawBackground() {
	var factorX = s.window.Bounds().Max.X / s.backgroundImageSprite.Picture().Bounds().Max.X
	var factorY = (s.window.Bounds().Max.Y - bottomBarHeight) / s.backgroundImageSprite.Picture().Bounds().Max.Y

	s.backgroundImageSprite.Draw(s.window, pixel.IM.
		ScaledXY(pixel.Vec{X: 0, Y: 0}, pixel.V(factorX, factorY)).
		Moved(s.window.Bounds().Center().Add(pixel.V(0, bottomBarHeight/2))))
}

func (s *Screen) drawItems() {
	imd := imdraw.New(nil)
	imd.Color = common.Black
	for idx := range s.items {
		var rect = getItemRect(s.window.Bounds().Max, idx)
		imd.Push(pixel.V(rect.X, rect.Y), pixel.V(rect.X+rect.Width, rect.Y+rect.Height))
		imd.Rectangle(0)
	}
	imd.Draw(s.window)
}

func (s *Screen) drawBottomBar() {
	// TODO: draw the content of the bottom bar
}

func getItemRect(screenSize pixel.Vec, itemIdx int) geometry.Rectangle {
	var row = float64(itemIdx / 5)
	var column = float64(itemIdx % 5)
	var width = (screenSize.X - 6*itemMargin) / 5
	var height = (screenSize.Y - bottomBarHeight - 7*itemMargin) / 6
	return geometry.Rectangle{
		X:      itemMargin + (column * itemMargin) + (column * width),
		Y:      screenSize.Y - ((row + 1) * itemMargin) - ((row + 1) * height),
		Width:  width,
		Height: height,
	}
}
