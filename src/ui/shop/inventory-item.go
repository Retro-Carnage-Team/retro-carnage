package shop

import "retro-carnage/assets"

// inventoryItem defines the shared functionality of grenades, ammunition and weapons.
type inventoryItem interface {
	Description() string
	Image() string
	Name() string
	Price() int
}

func getAllInventoryItems() (result []inventoryItem) {
	result = make([]inventoryItem, 0)
	for _, weapon := range assets.WeaponCrate.GetAll() {
		result = append(result, weapon)
	}
	for _, grenade := range assets.GrenadeCrate.GetAll() {
		result = append(result, grenade)
	}
	for _, ammunition := range assets.AmmunitionCrate.GetAll() {
		result = append(result, ammunition)
	}
	return
}
