package assets

type Grenade struct {
	description      string
	explosive        string
	image            string
	imageRotated     string
	maxCount         int32
	movementDistance int32
	movementSpeed    float32
	name             string
	packageSize      int32
	price            int32
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

func (g *Grenade) MaxCount() int32 {
	return g.maxCount
}

func (g *Grenade) MovementDistance() int32 {
	return g.movementDistance
}

func (g *Grenade) MovementSpeed() float32 {
	return g.movementSpeed
}

func (g *Grenade) Name() string {
	return g.name
}

func (g *Grenade) PackageSize() int32 {
	return g.packageSize
}

func (g *Grenade) Price() int32 {
	return g.price
}

func (g *Grenade) Weight() string {
	return g.weight
}
