package assets

type Ammunition struct {
	Description string
	Image       string
	MaxCount    int
	Name        string
	PackageSize int
	Price       int
}

func (a *Ammunition) GetDescription() string {
	return a.Description
}

func (a *Ammunition) GetImage() string {
	return a.Image
}

func (a *Ammunition) GetName() string {
	return a.Name
}

func (a *Ammunition) GetPrice() int {
	return a.Price
}
