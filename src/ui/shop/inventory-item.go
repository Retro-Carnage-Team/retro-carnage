package shop

import (
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/logging"
)

// inventoryItemDelegate defines the shared functionality of grenades, ammunition and weapons.
type inventoryItemDelegate interface {
	Description() string
	Image() string
	Name() string
	Price() int
}

type inventoryItem struct {
	delegate inventoryItemDelegate
	index    int
}

func (ii *inventoryItem) Description() string {
	return ii.delegate.Description()
}

func (ii *inventoryItem) Image() string {
	return ii.delegate.Image()
}

func (ii *inventoryItem) Name() string {
	return ii.delegate.Name()
}

func (ii *inventoryItem) Price() int {
	return ii.delegate.Price()
}

func (ii *inventoryItem) IsWeapon() bool {
	return ii.index < len(assets.WeaponCrate.GetAll())
}

func (ii *inventoryItem) IsGrenade() bool {
	return ii.index >= len(assets.WeaponCrate.GetAll()) &&
		ii.index < len(assets.WeaponCrate.GetAll())+len(assets.GrenadeCrate.GetAll())
}

func (ii *inventoryItem) IsAmmunition() bool {
	return !ii.IsWeapon() && !ii.IsGrenade()
}

// OwnedFromMax returns the current number of these item in inventory and their max count.
func (ii *inventoryItem) OwnedFromMax(playerIdx int) (int, int) {
	if ii.IsWeapon() {
		logging.Error.Fatal("this method should be used for grenades and ammunition only")
		return 0, 0
	} else if ii.IsGrenade() {
		var owned = characters.Players[playerIdx].GrenadeCount(ii.Name())
		var grenade = assets.GrenadeCrate.GetByName(ii.Name())
		return owned, grenade.MaxCount()
	} else {
		var owned = characters.Players[playerIdx].AmmunitionCount(ii.Name())
		var ammunition = assets.AmmunitionCrate.GetByName(ii.Name())
		return owned, ammunition.MaxCount()
	}
}

// OwnedPortion returns the portion of this item's max item count that the user already owns.
func (ii *inventoryItem) OwnedPortion(playerIdx int) float64 {
	var owned, max = ii.OwnedFromMax(playerIdx)
	return float64(owned) / float64(max)
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
