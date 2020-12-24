package shop

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"math"
	"retro-carnage/assets"
	"retro-carnage/engine"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/input"
	"retro-carnage/logging"
	"retro-carnage/ui/common"
	"retro-carnage/ui/common/fonts"
	"retro-carnage/util"
)

const backgroundImagePath = "./images/backgrounds/shop.jpg"
const bottomBarHeight = 70
const buttonPadding = 10
const checkImagePath = "./images/tiles/other/check-circle.png"
const itemMargin = 10.0
const itemPadding = 25.0
const labelAmmo = "Ammo: "
const labelLength = "Length: "
const labelPackageSize = "Package Size: "
const labelPrice = "Price: "
const labelRange = "Range: "
const labelSpeed = "Speed: "
const labelWeight = "Weight: "
const modalFontSize = 36
const modalColumnSpace = 200
const modalLabelSpace = 15
const modalTableVMargin = 30
const selectionBorderWidth = 5.0

type modalButton int

const (
	buttonBuyWeapon modalButton = iota
	buttonBuyAmmo
	buttonCloseModal
)

type Screen struct {
	backgroundImageSprite *pixel.Sprite
	checkSprite           *pixel.Sprite
	inputController       input.Controller
	inventoryController   engine.InventoryController
	itemNameToSprite      map[string]*pixel.Sprite
	items                 []*inventoryItem
	labelDimensions       map[string]*geometry.Point
	modalButtonSelection  modalButton
	modalVisible          bool
	PlayerIdx             int
	screenChangeRequired  common.ScreenChangeCallback
	selectedItemIdx       int
	stopWatch             *util.StopWatch
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
	s.checkSprite = common.LoadSprite(checkImagePath)
	s.inventoryController = engine.NewInventoryController(s.PlayerIdx)
	s.labelDimensions = fonts.GetTextDimensions(modalFontSize, labelAmmo, labelLength, labelPackageSize, labelPrice,
		labelRange, labelSpeed, labelWeight)
	s.selectedItemIdx = 0
	s.modalVisible = false

	s.items = getAllInventoryItems()
	s.itemNameToSprite = make(map[string]*pixel.Sprite)
	for _, item := range s.items {
		s.itemNameToSprite[item.Name()] = common.LoadSprite(item.Image())
	}

	s.stopWatch = &util.StopWatch{Name: "Render process"}
}

func (s *Screen) Update(_ int64) {
	if s.modalVisible {
		s.stopWatch.Start()
	}

	s.processUserInput()

	s.drawBackground()
	s.drawItems()
	s.drawBottomBar()

	if s.modalVisible {
		s.drawModal()
	}

	if s.modalVisible {
		_ = s.stopWatch.Stop()
		logging.Trace.Print(s.stopWatch.PrintDebugMessage())
	}
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
	s.drawItemSelectionBorder(itemAreas)
	s.drawPurchaseStatus(itemAreas)
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

func (s *Screen) drawItemSelectionBorder(areas []geometry.Rectangle) {
	if 0 <= s.selectedItemIdx && 30 >= s.selectedItemIdx && !s.modalVisible {
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

func (s *Screen) drawPurchaseStatus(areas []geometry.Rectangle) {
	imd := imdraw.New(nil)
	imd.Color = common.White

	for idx, area := range areas {
		var item = s.items[idx]
		if item.IsWeapon() {
			if s.inventoryController.WeaponInInventory(item.Name()) {
				s.checkSprite.Draw(s.window, pixel.IM.Moved(pixel.V(
					area.X+area.Width-s.checkSprite.Picture().Bounds().W(),
					area.Y+s.checkSprite.Picture().Bounds().H())))
			}
		} else {
			var ratio = item.OwnedPortion(s.PlayerIdx)
			if ratio > 0 {
				var barWidth = (area.Width - itemPadding - itemPadding) * ratio
				imd.Push(
					pixel.V(area.X+itemPadding, area.Y+itemPadding),
					pixel.V(area.X+itemPadding+barWidth, area.Y+itemPadding+5))
				imd.Rectangle(0)
			}
		}
	}

	imd.Draw(s.window)
}

func (s *Screen) drawBottomBar() {
	s.drawCostLabel()
	s.drawCreditLabel()
	s.drawExitButton()
}

func (s *Screen) processUserInput() {
	var eventState, err = s.inputController.GetControllerUiEventState(s.PlayerIdx)
	if nil != err {
		logging.Warning.Printf("Failed to get game controller state: %v", err)
	} else if nil != eventState {
		if eventState.MovedDown && !s.modalVisible {
			s.processSelectionMovedDown()
		} else if eventState.MovedUp && !s.modalVisible {
			s.processSelectionMovedUp()
		} else if eventState.MovedRight {
			s.processSelectionMovedRight()
		} else if eventState.MovedLeft {
			s.processSelectionMovedLeft()
		}
		if eventState.PressedButton {
			s.processButtonPressed()
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

func (s *Screen) processButtonPressed() {
	if s.modalVisible {
		s.processButtonPressedOnModal()
	} else {
		s.processButtonPressedOnShop()
	}
}

func (s *Screen) processButtonPressedOnModal() {
	var item = s.items[s.selectedItemIdx]
	switch s.modalButtonSelection {
	case buttonBuyWeapon:
		s.inventoryController.BuyWeapon(item.Name())
	case buttonBuyAmmo:
		if item.IsWeapon() {
			weapon := assets.WeaponCrate.GetByName(item.Name())
			s.inventoryController.BuyAmmunition(weapon.Ammo())
		} else if item.IsGrenade() {
			s.inventoryController.BuyGrenade(item.Name())
		} else if item.IsAmmunition() {
			s.inventoryController.BuyAmmunition(item.Name())
		}
	case buttonCloseModal:
		s.modalVisible = false
	}
}

func (s *Screen) processButtonPressedOnShop() {
	if -1 == s.selectedItemIdx {
		if (0 == s.PlayerIdx) && (2 == characters.PlayerController.NumberOfPlayers()) {
			s.screenChangeRequired(common.BuyYourWeaponsP2)
		} else {
			s.screenChangeRequired(common.LetTheMissionBegin)
		}
	} else {
		s.showModal()
	}
}

func (s *Screen) showModal() {
	if s.isModalButtonBuyWeaponAvailable() {
		s.modalButtonSelection = buttonBuyWeapon
	} else if s.isModalButtonBuyAmmunitionAvailable() {
		s.modalButtonSelection = buttonBuyAmmo
	} else {
		s.modalButtonSelection = buttonCloseModal
	}
	s.modalVisible = true
}

func (s *Screen) drawCostLabel() {
	var content = "COST: -"
	if -1 != s.selectedItemIdx {
		content = fmt.Sprintf("COST: %d", s.items[s.selectedItemIdx].Price())
	}

	var lineDimensions = fonts.GetTextDimension(fonts.DefaultFontSize, content)
	var lineY = (bottomBarHeight-lineDimensions.Y)/2 + buttonPadding
	fonts.BuildText(pixel.V(30.0, lineY), fonts.DefaultFontSize, common.White, content).Draw(s.window, pixel.IM)
}

func (s *Screen) drawCreditLabel() {
	var content = fmt.Sprintf("CREDIT: %d", characters.Players[s.PlayerIdx].Cash())
	var lineDimensions = fonts.GetTextDimension(fonts.DefaultFontSize, content)
	var lineX = (s.window.Bounds().W() - lineDimensions.X) / 2
	var lineY = (bottomBarHeight-lineDimensions.Y)/2 + buttonPadding
	fonts.BuildText(pixel.V(lineX, lineY), fonts.DefaultFontSize, common.White, content).Draw(s.window, pixel.IM)
}

func (s *Screen) drawExitButton() {
	var lineDimensions = fonts.GetTextDimension(fonts.DefaultFontSize, "EXIT SHOP")
	var lineX = s.window.Bounds().W() - lineDimensions.X - 30
	var lineY = (bottomBarHeight-lineDimensions.Y)/2 + buttonPadding
	fonts.BuildText(pixel.V(lineX, lineY), fonts.DefaultFontSize, common.White, "EXIT SHOP").Draw(s.window, pixel.IM)

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

func (s *Screen) drawModal() {
	s.drawModalHeader()
	var bodyLowerBorder = s.drawModalBody()
	s.drawModalFooter(bodyLowerBorder)
}

func (s *Screen) drawModalHeader() {
	imd := imdraw.New(nil)
	imd.Color = common.ModalBg
	imd.Push(
		pixel.V(s.getModalLeftBorder(), s.window.Bounds().H()-100),
		pixel.V(s.getModalRightBorder(), s.window.Bounds().H()-100-bottomBarHeight))
	imd.Rectangle(0)
	imd.Draw(s.window)

	var itemName = s.items[s.selectedItemIdx].Name()
	var lineDimensions = fonts.GetTextDimension(fonts.DefaultFontSize, itemName)
	var lineX = s.getModalLeftBorder() + 30
	var lineY = s.window.Bounds().H() - 100 - bottomBarHeight + (bottomBarHeight-lineDimensions.Y)/2
	fonts.BuildText(pixel.V(lineX, lineY), fonts.DefaultFontSize, common.White, itemName).Draw(s.window, pixel.IM)
}

func (s *Screen) drawModalBody() float64 {
	var item = s.items[s.selectedItemIdx]

	var textRenderer = fonts.TextRenderer{Window: s.window}

	var textWidth = s.getModalRightBorder() - s.getModalLeftBorder() - modalLabelSpace - modalLabelSpace
	textLayout, err := textRenderer.CalculateTextLayout(item.Description(), modalFontSize, int(textWidth), int(s.window.Bounds().H()-300))
	if nil != err {
		logging.Warning.Fatalf("text is too large for modal")
		return 0
	}

	var tableAreaHeight = s.labelDimensions[labelPrice].Y*3.4 + modalTableVMargin*2
	var descriptionAreaHeight = textLayout.Height()
	var modalBodyLowerBorder = s.window.Bounds().H() - 100 - bottomBarHeight - tableAreaHeight - descriptionAreaHeight - modalTableVMargin

	imd := imdraw.New(nil)
	imd.Color = common.White
	imd.Push(
		pixel.V(s.getModalLeftBorder(), s.window.Bounds().H()-100-bottomBarHeight),
		pixel.V(s.getModalRightBorder(), modalBodyLowerBorder))
	imd.Rectangle(0)
	imd.Draw(s.window)

	if item.IsWeapon() {
		s.drawModalBodyWeaponTable(item)
	} else {
		s.drawModalBodyAmmoGrenadeTable(item)
	}

	var lineY = s.window.Bounds().H() - 100 - bottomBarHeight - tableAreaHeight
	var atlas = fonts.SizeToFontAtlas[modalFontSize]
	var txt = text.New(pixel.V(s.getModalLeftBorder()+modalLabelSpace, lineY-textLayout.Lines()[0].Dimension().Y), atlas)
	txt.Color = common.Black
	for _, line := range textLayout.Lines() {
		_, _ = fmt.Fprintln(txt, line.Text())
	}
	txt.Draw(s.window, pixel.IM)
	return modalBodyLowerBorder
}

func (s *Screen) drawModalBodyWeaponTable(item *inventoryItem) {
	var weapon = assets.WeaponCrate.GetByName(item.Name())
	var priceValue = fmt.Sprintf("$%d", item.Price())
	var rangeValue = fmt.Sprintf("%d m", weapon.BulletRange())
	var speedValue = "Single shot"
	if 0 < weapon.BulletInterval() {
		speedValue = fmt.Sprintf("%d / minute", 60000/weapon.BulletInterval())
	}

	var maxLabelWidth = util.Max([]float64{
		s.labelDimensions[labelPrice].X, s.labelDimensions[labelAmmo].X, s.labelDimensions[labelLength].X,
		s.labelDimensions[labelSpeed].X, s.labelDimensions[labelRange].X, s.labelDimensions[labelWeight].X,
	})

	var maxValueWidth = fonts.GetMaxTextWidth(modalFontSize, []string{
		priceValue, weapon.Ammo(), weapon.Length(), speedValue, rangeValue, weapon.Weight(),
	})

	var columnWidth = maxLabelWidth + modalLabelSpace + maxValueWidth
	var firstColumnLabelX = s.window.Bounds().W()/2 - columnWidth - modalColumnSpace/2
	var firstColumnValueX = firstColumnLabelX + maxLabelWidth + modalLabelSpace
	var secondColumnLabelX = firstColumnValueX + maxValueWidth + modalColumnSpace
	var secondColumnValueX = secondColumnLabelX + maxLabelWidth + modalLabelSpace

	var firstRowY = s.window.Bounds().H() - 100 - bottomBarHeight - modalTableVMargin - s.labelDimensions[labelPrice].Y
	fonts.BuildMultiLineText(pixel.V(firstColumnLabelX, firstRowY), modalFontSize, common.Black, []string{labelPrice, labelAmmo, labelLength}).Draw(s.window, pixel.IM)
	fonts.BuildMultiLineText(pixel.V(firstColumnValueX, firstRowY), modalFontSize, common.Black, []string{priceValue, weapon.Ammo(), weapon.Length()}).Draw(s.window, pixel.IM)
	fonts.BuildMultiLineText(pixel.V(secondColumnLabelX, firstRowY), modalFontSize, common.Black, []string{labelSpeed, labelRange, labelWeight}).Draw(s.window, pixel.IM)
	fonts.BuildMultiLineText(pixel.V(secondColumnValueX, firstRowY), modalFontSize, common.Black, []string{speedValue, rangeValue, weapon.Weight()}).Draw(s.window, pixel.IM)
}

func (s *Screen) drawModalBodyAmmoGrenadeTable(item *inventoryItem) {
	var priceValue = fmt.Sprintf("$%d", item.Price())
	var packageSizeValue = ""
	var rangeValue = ""

	if item.IsAmmunition() {
		var ammo = assets.AmmunitionCrate.GetByName(item.Name())
		packageSizeValue = fmt.Sprintf("%d", ammo.PackageSize())
	} else {
		var grenade = assets.GrenadeCrate.GetByName(item.Name())
		packageSizeValue = fmt.Sprintf("%d", grenade.PackageSize())
		rangeValue = fmt.Sprintf("%d m", grenade.MovementDistance())
	}

	var maxLabelWidth = util.Max([]float64{
		s.labelDimensions[labelPrice].X, s.labelDimensions[labelPackageSize].X, s.labelDimensions[labelRange].X,
	})
	var maxValueWidth = fonts.GetMaxTextWidth(modalFontSize, []string{
		priceValue, packageSizeValue, rangeValue,
	})

	var labelX = (s.window.Bounds().W() - maxLabelWidth - modalLabelSpace - maxValueWidth) / 2
	var valueX = labelX + maxLabelWidth + modalColumnSpace
	var labelY = s.window.Bounds().H() - 100 - bottomBarHeight - modalTableVMargin - s.labelDimensions[labelPrice].Y

	var thirdLabel = ""
	var thirdValue = ""
	if item.IsGrenade() {
		thirdLabel = labelRange
		thirdValue = rangeValue
	}

	fonts.BuildMultiLineText(pixel.V(labelX, labelY), modalFontSize, common.Black, []string{labelPrice, labelPackageSize, thirdLabel}).Draw(s.window, pixel.IM)
	fonts.BuildMultiLineText(pixel.V(valueX, labelY), modalFontSize, common.Black, []string{priceValue, packageSizeValue, thirdValue}).Draw(s.window, pixel.IM)
}

func (s *Screen) drawModalFooter(upperBorder float64) {
	imd := imdraw.New(nil)
	imd.Color = common.ModalBg
	imd.Push(
		pixel.V(s.getModalLeftBorder(), upperBorder),
		pixel.V(s.getModalRightBorder(), upperBorder-bottomBarHeight))
	imd.Rectangle(0)
	imd.Draw(s.window)

	var hint = s.getModalFooterStatusHint()
	if "" != hint {
		var lineDimensions = fonts.GetTextDimension(modalFontSize, hint)
		var lineX = s.getModalLeftBorder() + 30
		var lineY = upperBorder - lineDimensions.Y - (bottomBarHeight-lineDimensions.Y)/2
		fonts.BuildText(pixel.V(lineX, lineY), modalFontSize, common.OliveGreen, hint).Draw(s.window, pixel.IM)
	}

	var leftBorder = s.drawModalCloseButton(upperBorder, upperBorder-bottomBarHeight)

	if s.isModalButtonBuyAmmunitionAvailable() {
		leftBorder = s.drawModalBuyAmmoButton(upperBorder, upperBorder-bottomBarHeight, leftBorder)
	}

	if s.isModalButtonBuyWeaponAvailable() {
		s.drawModalBuyWeaponButton(upperBorder, upperBorder-bottomBarHeight, leftBorder)
	}
}

func (s *Screen) drawModalCloseButton(top float64, bottom float64) (leftBorder float64) {
	var lineDimensions = fonts.GetTextDimension(modalFontSize, "CLOSE")
	leftBorder = s.getModalRightBorder() - 30 - lineDimensions.X - buttonPadding - buttonPadding

	//------------------------------------------------
	// TODO: Move this following block into a function
	//------------------------------------------------
	var upperRight = pixel.V(s.getModalRightBorder()-30, top-buttonPadding)
	var lowerLeft = pixel.V(leftBorder, bottom+buttonPadding)
	imd := imdraw.New(nil)
	imd.Color = common.Black
	imd.Push(upperRight, lowerLeft)
	imd.Rectangle(0)
	imd.Draw(s.window)

	var buttonHeight = top - buttonPadding - bottom - buttonPadding
	var buttonTextX = s.getModalRightBorder() - 30 - lineDimensions.X - buttonPadding
	var buttonTextY = bottom + buttonPadding + (buttonHeight-lineDimensions.Y)/2
	fonts.
		BuildText(pixel.V(buttonTextX, buttonTextY), modalFontSize, common.White, "CLOSE").
		Draw(s.window, pixel.IM)

	if s.modalButtonSelection == buttonCloseModal {
		imd := imdraw.New(nil)
		imd.Color = common.Yellow
		imd.Push(upperRight, lowerLeft)
		imd.Rectangle(selectionBorderWidth)
		imd.Draw(s.window)
	}
	//------------------------------------------------
	return
}

func (s *Screen) drawModalBuyAmmoButton(top float64, bottom float64, closeButtonLeft float64) (leftBorder float64) {
	var labelText = s.getModalBuyAmmoButtonLabel()
	var lineDimensions = fonts.GetTextDimension(modalFontSize, labelText)
	var rightBorder = closeButtonLeft - 30
	leftBorder = rightBorder - lineDimensions.X - buttonPadding - buttonPadding

	//------------------------------------------------
	// TODO: Move this following block into a function
	//------------------------------------------------
	var upperRight = pixel.V(rightBorder, top-buttonPadding)
	var lowerLeft = pixel.V(leftBorder, bottom+buttonPadding)
	imd := imdraw.New(nil)
	imd.Color = common.Black
	imd.Push(upperRight, lowerLeft)
	imd.Rectangle(0)
	imd.Draw(s.window)

	var buttonHeight = top - buttonPadding - bottom - buttonPadding
	var buttonTextX = leftBorder + buttonPadding
	var buttonTextY = bottom + buttonPadding + (buttonHeight-lineDimensions.Y)/2
	fonts.
		BuildText(pixel.V(buttonTextX, buttonTextY), modalFontSize, common.White, labelText).
		Draw(s.window, pixel.IM)

	if s.modalButtonSelection == buttonBuyAmmo {
		imd := imdraw.New(nil)
		imd.Color = common.Yellow
		imd.Push(upperRight, lowerLeft)
		imd.Rectangle(selectionBorderWidth)
		imd.Draw(s.window)
	}
	//------------------------------------------------
	return
}

func (s *Screen) drawModalBuyWeaponButton(top float64, bottom float64, buyAmmoButtonLeft float64) {
	var labelText = "BUY WEAPON"
	var lineDimensions = fonts.GetTextDimension(modalFontSize, labelText)
	var rightBorder = buyAmmoButtonLeft - 30
	var leftBorder = rightBorder - lineDimensions.X - buttonPadding - buttonPadding

	//------------------------------------------------
	// TODO: Move this following block into a function
	//------------------------------------------------
	var upperRight = pixel.V(rightBorder, top-buttonPadding)
	var lowerLeft = pixel.V(leftBorder, bottom+buttonPadding)

	imd := imdraw.New(nil)
	imd.Color = common.Black
	imd.Push(upperRight, lowerLeft)
	imd.Rectangle(0)
	imd.Draw(s.window)

	var buttonHeight = top - buttonPadding - bottom - buttonPadding
	var buttonTextX = leftBorder + buttonPadding
	var buttonTextY = bottom + buttonPadding + (buttonHeight-lineDimensions.Y)/2
	fonts.
		BuildText(pixel.V(buttonTextX, buttonTextY), modalFontSize, common.White, labelText).
		Draw(s.window, pixel.IM)

	if s.modalButtonSelection == buttonBuyWeapon {
		imd := imdraw.New(nil)
		imd.Color = common.Yellow
		imd.Push(upperRight, lowerLeft)
		imd.Rectangle(selectionBorderWidth)
		imd.Draw(s.window)
	}
	//------------------------------------------------
}

func (s *Screen) getModalBuyAmmoButtonLabel() string {
	var item = s.items[s.selectedItemIdx]
	if item.IsWeapon() {
		var weapon = assets.WeaponCrate.GetByName(item.Name())
		var ammo = assets.AmmunitionCrate.GetByName(weapon.Ammo())
		return fmt.Sprintf("BUY %d BULLET(S) FOR $ %d", ammo.PackageSize(), ammo.Price())
	} else if item.IsAmmunition() {
		var ammo = assets.AmmunitionCrate.GetByName(item.Name())
		return fmt.Sprintf("BUY %d BULLET(S) FOR $ %d", ammo.PackageSize(), ammo.Price())
	} else {
		var grenade = assets.GrenadeCrate.GetByName(item.Name())
		return fmt.Sprintf("BUY %d BULLET(S) FOR $ %d", grenade.PackageSize(), grenade.Price())
	}
}

func (s *Screen) getModalFooterStatusHint() string {
	var item = s.items[s.selectedItemIdx]
	if item.IsWeapon() {
		var weapon = assets.WeaponCrate.GetByName(item.Name())
		if s.inventoryController.WeaponInInventory(item.Name()) {
			for _, ammo := range s.items {
				if ammo.Name() == weapon.Ammo() {
					var owned, max = ammo.OwnedFromMax(s.PlayerIdx)
					return fmt.Sprintf("You own this weapon and %d of %d bullets", owned, max)
				}
			}
			logging.Error.Fatalf("Failed to find ammo item: %s", weapon.Ammo())
		} else {
			return "You don't own this weapon yet"
		}
	}

	var _type = "bullets"
	if item.IsGrenade() {
		_type = "grenades"
	}
	var owned, max = item.OwnedFromMax(s.PlayerIdx)
	return fmt.Sprintf("You own %d of %d of these %s", owned, max, _type)
}

func (s *Screen) isModalButtonBuyWeaponAvailable() bool {
	return s.selectedItemIdx != -1 &&
		s.items[s.selectedItemIdx].IsWeapon() &&
		s.inventoryController.WeaponProcurable(s.items[s.selectedItemIdx].Name())
}

func (s *Screen) isModalButtonBuyAmmunitionAvailable() bool {
	if s.selectedItemIdx == -1 {
		return false
	}

	var item = s.items[s.selectedItemIdx]
	if item.IsGrenade() {
		return s.inventoryController.GrenadeProcurable(item.Name())
	}

	var ammoName = item.Name()
	if item.IsWeapon() {
		var weapon = assets.WeaponCrate.GetByName(item.Name())
		ammoName = weapon.Ammo()
	}
	return s.inventoryController.AmmunitionProcurable(ammoName)
}

func (s *Screen) getModalLeftBorder() float64 {
	return (s.window.Bounds().W() / 5) + 50
}

func (s *Screen) getModalRightBorder() float64 {
	return (s.window.Bounds().W() * 4 / 5) - 50
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
