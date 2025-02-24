package assets

type WeaponType int

const (
	NonAutomatic WeaponType = 0
	Automatic    WeaponType = 1
	RPG          WeaponType = 2
)

type Weapon struct {
	Ammo           string
	BulletInterval int     // (min) offset between two bullets in ms
	BulletRange    int     // in pixels
	BulletSpeed    float64 // pixels per ms
	CoolDownSound  SoundEffect
	Description    string
	Image          string
	ImageRotated   string
	Length         string
	Name           string
	Price          int
	Sound          SoundEffect
	WeaponType     WeaponType
	Weight         string
}

func (w *Weapon) GetDescription() string {
	return w.Description
}

func (w *Weapon) GetImage() string {
	return w.Image
}

func (w *Weapon) GetName() string {
	return w.Name
}

func (w *Weapon) GetPrice() int {
	return w.Price
}
