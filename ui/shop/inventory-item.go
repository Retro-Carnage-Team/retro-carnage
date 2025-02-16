package shop

import (
	"retro-carnage/assets"
	"retro-carnage/engine"
)

// inventoryItemDelegate defines the shared functionality of grenades, ammunition and items.
type inventoryItemDelegate interface {
	GetDescription() string
	GetImage() string
	GetName() string
	GetPrice() int
}

type inventoryItem struct {
	delegate inventoryItemDelegate
	index    int
}

func (ii *inventoryItem) Description() string {
	return ii.delegate.GetDescription()
}

func (ii *inventoryItem) Image() string {
	return ii.delegate.GetImage()
}

func (ii *inventoryItem) Name() string {
	return ii.delegate.GetName()
}

func (ii *inventoryItem) Price() int {
	return ii.delegate.GetPrice()
}

func (ii *inventoryItem) IsWeapon() bool {
	return ii.index < len(assets.WeaponCrate.GetAll())
}

func (ii *inventoryItem) IsGrenade() bool {
	return ii.index >= len(assets.WeaponCrate.GetAll()) &&
		ii.index < len(assets.WeaponCrate.GetAll())+len(assets.GrenadeCrate.GetAll())
}

// OwnedFromMax returns the current number of these item in inventory and their max count.
func (ii *inventoryItem) OwnedFromMax(inventoryController *engine.InventoryController) (int, int) {
	if ii.IsWeapon() {
		var weapon = assets.WeaponCrate.GetByName(ii.Name())
		var owned = inventoryController.AmmunitionCount(weapon.Ammo)
		var ammunition = assets.AmmunitionCrate.GetByName(weapon.Ammo)
		return owned, ammunition.MaxCount
	} else {
		// item is a grenade
		var owned = inventoryController.GrenadeCount(ii.Name())
		var grenade = assets.GrenadeCrate.GetByName(ii.Name())
		return owned, grenade.MaxCount
	}
}

// OwnedPortion returns the portion of this item's max item count that the user already owns.
func (ii *inventoryItem) OwnedPortion(inventoryController *engine.InventoryController) float64 {
	var owned, max = ii.OwnedFromMax(inventoryController)
	return float64(owned) / float64(max)
}
