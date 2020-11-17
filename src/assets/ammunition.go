package assets

type Ammunition struct {
	description string
	image       string
	maxCount    int32
	name        string
	packageSize int32
	price       int32
}

func (a *Ammunition) Description() string {
	return a.description
}

func (a *Ammunition) Image() string {
	return a.image
}

func (a *Ammunition) MaxCount() int32 {
	return a.maxCount
}

func (a *Ammunition) Name() string {
	return a.name
}

func (a *Ammunition) PackageSize() int32 {
	return a.packageSize
}

func (a *Ammunition) Price() int32 {
	return a.price
}
