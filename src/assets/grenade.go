package assets

type Grenade struct {
	Description      string
	Explosive        string
	Image            string
	ImageRotated     string
	MaxCount         int
	MovementDistance int
	MovementSpeed    float64
	Name             string
	PackageSize      int
	Price            int
	Weight           string
}

func (g *Grenade) GetDescription() string {
	return g.Description
}

func (g *Grenade) GetImage() string {
	return g.Image
}

func (g *Grenade) GetName() string {
	return g.Name
}
func (g *Grenade) GetPrice() int {
	return g.Price
}
