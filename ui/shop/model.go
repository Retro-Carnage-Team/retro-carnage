package shop

import "retro-carnage/assets"

type model struct {
	items                []*inventoryItem
	modalButtonSelection modalButton
	modalVisible         bool
	playerIdx            int
	selectedItemIdx      int
}

func newModel(playerIdx int) *model {
	var result = model{
		items:           getAllInventoryItems(),
		modalVisible:    false,
		playerIdx:       playerIdx,
		selectedItemIdx: 0,
	}
	return &result
}

func getAllInventoryItems() (result []*inventoryItem) {
	result = make([]*inventoryItem, 0)
	var idx = 0
	for _, weapon := range assets.WeaponCrate.GetAll() {
		result = append(result, &inventoryItem{
			delegate: weapon,
			index:    idx,
		})
		idx += 1
	}
	for _, grenade := range assets.GrenadeCrate.GetAll() {
		result = append(result, &inventoryItem{
			delegate: grenade,
			index:    idx,
		})
		idx += 1
	}
	for _, ammunition := range assets.AmmunitionCrate.GetAll() {
		result = append(result, &inventoryItem{
			delegate: ammunition,
			index:    idx,
		})
		idx += 1
	}
	return
}
