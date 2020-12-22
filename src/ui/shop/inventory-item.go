package shop

import (
	"errors"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
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
	return ii.index < len(assets.WeaponCrate.GetAll())+len(assets.GrenadeCrate.GetAll())
}

func (ii *inventoryItem) IsAmmunition() bool {
	return !ii.IsWeapon() && !ii.IsGrenade()
}

func (ii *inventoryItem) OwnedPortion(playerIdx int) (float64, error) {
	if ii.IsWeapon() {
		return 0, errors.New("this method should be used for grenades and ammunition only")
	} else if ii.IsGrenade() {
		var owned = characters.Players[playerIdx].GrenadeCount(ii.Name())
		var grenade, err = assets.GrenadeCrate.GetByName(ii.Name())
		if nil != err {
			return 0, err
		}
		return float64(owned) / float64(grenade.MaxCount()), nil
	} else {
		var owned = characters.Players[playerIdx].AmmunitionCount(ii.Name())
		var ammunition, err = assets.AmmunitionCrate.GetByName(ii.Name())
		if nil != err {
			return 0, err
		}
		return float64(owned) / float64(ammunition.MaxCount()), nil
	}
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
