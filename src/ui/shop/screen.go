package shop

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"math"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
)

const backgroundImagePath = "./images/backgrounds/shop.jpg"
const bottomBarHeight = 70
const buttonPadding = 10
const itemMargin = 10.0
const itemPadding = 25.0
const selectionBorderWidth = 5.0

type Screen struct {
	backgroundImageSprite *pixel.Sprite
	inputController       input.Controller
	itemNameToSprite      map[string]*pixel.Sprite
	items                 []inventoryItem
	PlayerIdx             int
	screenChangeRequired  common.ScreenChangeCallback
	selectedItemIdx       int
	window                *pixelgl.Window
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
	s.selectedItemIdx = 0

	s.items = getAllInventoryItems()
	s.itemNameToSprite = make(map[string]*pixel.Sprite)
	for _, item := range s.items {
		s.itemNameToSprite[item.Name()] = common.LoadSprite(item.Image())
	}
}

func (s *Screen) Update(_ int64) {
	s.processUserInput()

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
	var itemAreas = make([]geometry.Rectangle, 0)
	for idx := range s.items {
		itemAreas = append(itemAreas, getItemRect(s.window.Bounds().Max, idx))
	}

	s.drawItemBackgrounds(itemAreas)
	s.drawItemImages(itemAreas)
	s.selectionBorder(itemAreas)
}

func (s *Screen) drawItemBackgrounds(itemAreas []geometry.Rectangle) {
	imd := imdraw.New(nil)
	imd.Color = common.Black
	for _, area := range itemAreas {
		imd.Push(pixel.V(area.X, area.Y), pixel.V(area.X+area.Width, area.Y+area.Height))
		imd.Rectangle(0)
	}
	imd.Draw(s.window)
}

func (s *Screen) selectionBorder(areas []geometry.Rectangle) {
	if 0 <= s.selectedItemIdx && 30 >= s.selectedItemIdx {
		var area = areas[s.selectedItemIdx]
		imd := imdraw.New(nil)
		imd.Color = common.Yellow
		imd.Push(pixel.V(area.X, area.Y), pixel.V(area.X+area.Width, area.Y+area.Height))
		imd.Rectangle(selectionBorderWidth)
		imd.Draw(s.window)
	}
}

func (s *Screen) drawItemImages(itemAreas []geometry.Rectangle) {
	var sampleSprite = s.itemNameToSprite[s.items[0].Name()]
	var factorX = (itemAreas[0].Width - itemPadding*2) / sampleSprite.Picture().Bounds().W()
	var factorY = (itemAreas[0].Height - itemPadding*2) / sampleSprite.Picture().Bounds().H()
	var factor = math.Min(factorX, factorY)
	for idx, item := range s.items {
		var itemArea = itemAreas[idx]
		s.itemNameToSprite[item.Name()].Draw(s.window, pixel.IM.
			Scaled(pixel.V(0, 0), factor).
			Moved(itemArea.Center().ToVec()))
	}
}

func (s *Screen) drawBottomBar() {
	s.drawExitButton()
}

func (s *Screen) processUserInput() {
	var eventState, err = s.inputController.GetControllerUiEventState(s.PlayerIdx)
	if nil != err {
		logging.Warning.Printf("Failed to get game controller state: %v", err)
	} else if nil != eventState {
		if eventState.MovedDown {
			s.processSelectionMovedDown()
		}
		if eventState.MovedUp {
			s.processSelectionMovedUp()
		}
		if eventState.MovedRight {
			s.processSelectionMovedRight()
		}
		if eventState.MovedLeft {
			s.processSelectionMovedLeft()
		}
	}
}

func (s *Screen) processSelectionMovedDown() {
	if -1 != s.selectedItemIdx {
		if 5 <= s.selectedItemIdx/5 {
			s.selectedItemIdx = -1
		} else {
			s.selectedItemIdx += 5
		}
	} else {
		s.selectedItemIdx = 4
	}
}

func (s *Screen) processSelectionMovedUp() {
	if -1 != s.selectedItemIdx {
		if 5 > s.selectedItemIdx {
			s.selectedItemIdx = -1
		} else {
			s.selectedItemIdx -= 5
		}
	} else {
		s.selectedItemIdx = len(s.items) - 1
	}
}

func (s *Screen) processSelectionMovedRight() {
	if -1 != s.selectedItemIdx {
		if 4 == s.selectedItemIdx%5 {
			s.selectedItemIdx -= 4
		} else {
			s.selectedItemIdx += 1
		}
	}
}

func (s *Screen) processSelectionMovedLeft() {
	if -1 != s.selectedItemIdx {
		if 0 == s.selectedItemIdx%5 {
			s.selectedItemIdx += 4
		} else {
			s.selectedItemIdx -= 1
		}
	}
}

func (s *Screen) drawExitButton() {
	var lineDimensions = fonts.GetTextDimension(fonts.DefaultFontSize, "EXIT SHOP")
	var lineX = s.window.Bounds().W() - lineDimensions.X - 30
	var lineY = (bottomBarHeight-lineDimensions.Y)/2 + buttonPadding
	var txt = text.New(pixel.V(lineX, lineY), fonts.SizeToFontAtlas[fonts.DefaultFontSize])
	txt.Color = common.White
	_, _ = fmt.Fprint(txt, "EXIT SHOP")
	txt.Draw(s.window, pixel.IM)

	if -1 == s.selectedItemIdx {
		imd := imdraw.New(nil)
		imd.Color = common.Yellow
		imd.Push(
			pixel.V(lineX-buttonPadding, lineY-buttonPadding),
			pixel.V(lineX+buttonPadding*2+lineDimensions.X, lineY+lineDimensions.Y))
		imd.Rectangle(selectionBorderWidth)
		imd.Draw(s.window)
	}
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
