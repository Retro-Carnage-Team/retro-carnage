package assets

type Ammunition struct {
	description string
	image       string
	maxCount    int
	name        string
	packageSize int
	price       int
}

func (a *Ammunition) Description() string {
	return a.description
}

func (a *Ammunition) Image() string {
	return a.image
}

func (a *Ammunition) MaxCount() int {
	return a.maxCount
}

func (a *Ammunition) Name() string {
	return a.name
}

func (a *Ammunition) PackageSize() int {
	return a.packageSize
}

func (a *Ammunition) Price() int {
	return a.price
}
