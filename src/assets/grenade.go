package assets

type Grenade struct {
	description      string
	explosive        string
	image            string
	imageRotated     string
	maxCount         int
	movementDistance int
	movementSpeed    float64
	name             string
	packageSize      int
	price            int
	weight           string
}

func (g *Grenade) Description() string {
	return g.description
}

func (g *Grenade) Explosive() string {
	return g.explosive
}

func (g *Grenade) Image() string {
	return g.image
}

func (g *Grenade) ImageRotated() string {
	return g.imageRotated
}

func (g *Grenade) MaxCount() int {
	return g.maxCount
}

func (g *Grenade) MovementDistance() int {
	return g.movementDistance
}

func (g *Grenade) MovementSpeed() float64 {
	return g.movementSpeed
}

func (g *Grenade) Name() string {
	return g.name
}

func (g *Grenade) PackageSize() int {
	return g.packageSize
}

func (g *Grenade) Price() int {
	return g.price
}

func (g *Grenade) Weight() string {
	return g.weight
}
