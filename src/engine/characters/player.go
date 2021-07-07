package characters

import (
	"errors"
	"fmt"
	"retro-carnage/assets"
	"retro-carnage/util"
)

const (
	PlayerPropertyAmmunition     = "ammunition"
	PlayerPropertyCash           = "cash"
	PlayerPropertyGrenades       = "grenades"
	PlayerPropertyLives          = "lives"
	PlayerPropertyScore          = "score"
	PlayerPropertySelectedWeapon = "selected-weapon"
	PlayerPropertyWeapons        = "items"
)

var (
	playerOne = newPlayer(0)
	playerTwo = newPlayer(1)
	Players   = []*Player{playerOne, playerTwo}
)

type Player struct {
	ammunition         map[string]int
	cash               int
	changeListeners    []*util.ChangeListener
	grenades           map[string]bool
	index              int
	lives              int
	score              int
	selectedWeaponName *string
	weapons            map[string]bool
}

func newPlayer(index int) *Player {
	var result = &Player{}
	result.changeListeners = make([]*util.ChangeListener, 0)
	result.index = index
	result.cash = 5000
	result.lives = 3
	result.score = 0
	result.selectedWeaponName = nil
	result.ammunition = make(map[string]int, 0)
	result.grenades = make(map[string]bool, 0)
	result.weapons = make(map[string]bool, 0)
	return result
}

func (p *Player) AddChangeListener(changeListener *util.ChangeListener) {
	p.changeListeners = append(p.changeListeners, changeListener)
}

func (p *Player) RemoveChangeListener(changeListener *util.ChangeListener) error {
	for idx, cListener := range p.changeListeners {
		if cListener == changeListener {
			p.changeListeners[idx] = p.changeListeners[len(p.changeListeners)-1]
			p.changeListeners = p.changeListeners[:len(p.changeListeners)-1]
			return nil
		}
	}
	return errors.New("the given change listener has not been registered")
}

func (p *Player) Alive() bool {
	return p.lives > 0
}

func (p *Player) AmmunitionCount(ammunition string) int {
	return p.ammunition[ammunition]
}

func (p *Player) SetAmmunitionCount(ammunition string, count int) {
	p.ammunition[ammunition] = count
	p.notifyListeners(count, PlayerPropertyAmmunition)
}

func (p *Player) AmmunitionForSelectedWeapon() int {
	if nil == p.selectedWeaponName {
		return 0
	} else {
		var weapon = p.SelectedWeapon()
		if nil != weapon {
			return p.AmmunitionCount(weapon.Ammo)
		} else {
			var grenade = p.SelectedGrenade()
			if nil == grenade {
				return 0
			}
			return p.GrenadeCount(grenade.Name)
		}
	}
}

func (p *Player) AutomaticWeaponSelected() bool {
	var weapon = p.SelectedWeapon()
	return (nil != weapon) && (assets.Automatic == weapon.WeaponType)
}

func (p *Player) Cash() int {
	return p.cash
}

func (p *Player) SetCash(cash int) {
	p.cash = cash
	p.notifyListeners(cash, PlayerPropertyCash)
}

func (p *Player) GrenadeCount(grenade string) int {
	return p.ammunition[grenade]
}

func (p *Player) SetGrenadeCount(grenade string, count int) {
	if 0 < count {
		p.grenades[grenade] = true
	}
	p.ammunition[grenade] = count
	p.notifyListeners(count, PlayerPropertyAmmunition)
}

func (p *Player) GrenadeSelected() bool {
	for _, grenade := range assets.GrenadeCrate.GetAll() {
		if grenade.Name == *p.selectedWeaponName {
			return true
		}
	}
	return false
}

func (p *Player) Index() int {
	return p.index
}

func (p *Player) Lives() int {
	return p.lives
}

func (p *Player) SetLives(lives int) {
	p.lives = lives
	p.notifyListeners(lives, PlayerPropertyLives)
}

func (p *Player) PistolSelected() bool {
	var weapon = p.SelectedWeapon()
	return (nil != weapon) && (assets.Pistol == weapon.WeaponType)
}

func (p *Player) Reset() {
	p.cash = 5000
	p.lives = 3
	p.score = 0
	p.selectedWeaponName = nil
	p.ammunition = make(map[string]int, 0)
	p.grenades = make(map[string]bool, 0)
	p.weapons = make(map[string]bool, 0)
}

func (p *Player) RpgSelected() bool {
	var weapon = p.SelectedWeapon()
	return (nil != weapon) && (assets.RPG == weapon.WeaponType)
}

func (p *Player) Score() int {
	return p.score
}

func (p *Player) SetScore(score int) {
	p.score = score
	p.notifyListeners(score, PlayerPropertyScore)
}

func (p *Player) SelectedGrenade() *assets.Grenade {
	if nil == p.selectedWeaponName {
		return nil
	}
	for _, grenade := range assets.GrenadeCrate.GetAll() {
		if grenade.Name == *p.selectedWeaponName {
			return grenade
		}
	}
	return nil
}

func (p *Player) SelectedWeapon() *assets.Weapon {
	if nil == p.selectedWeaponName {
		return nil
	}
	for _, weapon := range assets.WeaponCrate.GetAll() {
		if weapon.GetName() == *p.selectedWeaponName {
			return weapon
		}
	}
	return nil
}

func (p *Player) SelectFirstWeapon() {
	var itemNames = p.getNamesOfWeaponsAndGrenadesInInventory()
	p.selectedWeaponName = itemNames[0]
	p.notifyListeners(*p.selectedWeaponName, PlayerPropertySelectedWeapon)
}

func (p *Player) SelectNextWeapon() {
	if p.AutomaticWeaponSelected() {
		assets.NewStereo().StopFx(p.SelectedWeapon().Sound)
	}

	var itemNames = p.getNamesOfWeaponsAndGrenadesInInventory()
	for idx, name := range itemNames {
		if *name == *p.selectedWeaponName {
			if idx < len(itemNames)-1 {
				p.selectedWeaponName = itemNames[idx+1]
			} else {
				p.selectedWeaponName = itemNames[0]
			}
			break
		}
	}
	p.notifyListeners(*p.selectedWeaponName, PlayerPropertySelectedWeapon)
}

func (p *Player) SelectPreviousWeapon() {
	if p.AutomaticWeaponSelected() {
		assets.NewStereo().StopFx(p.SelectedWeapon().Sound)
	}

	var itemNames = p.getNamesOfWeaponsAndGrenadesInInventory()
	for idx, name := range itemNames {
		if *name == *p.selectedWeaponName {
			if idx > 0 {
				p.selectedWeaponName = itemNames[idx-1]
			} else {
				p.selectedWeaponName = itemNames[len(itemNames)-1]
			}
			break
		}
	}
	p.notifyListeners(*p.selectedWeaponName, PlayerPropertySelectedWeapon)
}

func (p *Player) String() string {
	return fmt.Sprintf("Player %d", p.index+1)
}

func (p *Player) WeaponInInventory(weapon string) bool {
	return p.weapons[weapon]
}

func (p *Player) SetWeaponInInventory(weapon string, value bool) {
	p.weapons[weapon] = value
	p.notifyListeners(value, PlayerPropertyWeapons)
}

func (p *Player) notifyListeners(value interface{}, property string) {
	for _, cListener := range p.changeListeners {
		cListener.Call(value, property)
	}
}

func (p *Player) getNamesOfWeaponsAndGrenadesInInventory() []*string {
	var result = make([]*string, 0)
	for _, weapon := range assets.WeaponCrate.GetAll() {
		if p.weapons[weapon.GetName()] {
			var temp = weapon.GetName()
			result = append(result, &temp)
		}
	}
	for _, grenade := range assets.GrenadeCrate.GetAll() {
		if p.grenades[grenade.GetName()] {
			var temp = grenade.Name
			result = append(result, &temp)
		}
	}
	return result
}
