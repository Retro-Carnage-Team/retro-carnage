package assets

import "errors"

type GrenadeContainer struct {
	grenades []*Grenade
}

func (gc *GrenadeContainer) GetAll() []*Grenade {
	return gc.grenades
}

func (gc *GrenadeContainer) GetByName(name string) (*Grenade, error) {
	for _, grenade := range gc.grenades {
		if grenade.Name() == name {
			return grenade, nil
		}
	}
	return nil, errors.New("no such element")
}

var (
	GrenadeCrate = GrenadeContainer{grenades: initializeGrenades()}
)

func initializeGrenades() []*Grenade {
	var result = make([]*Grenade, 0)

	result = append(result, &Grenade{
		description:      "The DM41 is a fragmentation hand grenade and based on the US-American M26A2 hand grenade with fuse M215. The M26 entered service around 1952 and was used in combat during the Korean War. Its distinct lemon shape led it to being nicknamed the 'lemon grenade' (compare the Russian F1 grenade and American Mk 2 'pineapple' grenade, with similar nicknames). Fragmentation is enhanced by a special pre-notched fragmentation coil that lies along the inside of the grenade's body. This coil had a circular cross-section in the M26 grenade and an improved square cross-section in the M26A1 and later designs.",
		explosive:        "150 g",
		image:            "images/tiles/weapons/DM41.png",
		imageRotated:     "images/tiles/weapons/DM41-r.png",
		maxCount:         100,
		movementDistance: 450,
		movementSpeed:    0.8,
		name:             "DM41",
		packageSize:      5,
		price:            500,
		weight:           "0.450 kg",
	})

	result = append(result, &Grenade{
		description:      "The Stielhandgranate (German for 'stick hand grenade') was a German hand grenade of unique design. It was the standard issue of the German Empire during World War I, and became the widespread issue of Nazi Germany's Wehrmacht during World War II. The very distinctive appearance led to it being called a 'stick grenade', or 'potato masher' in British Army slang, and is today one of the most easily recognized infantry weapons of the 20th century.",
		explosive:        "170 g",
		image:            "images/tiles/weapons/M24.png",
		imageRotated:     "images/tiles/weapons/M24-r.png",
		maxCount:         100,
		movementDistance: 550,
		movementSpeed:    0.85,
		name:             "Stielhandgranate 24",
		packageSize:      5,
		price:            600,
		weight:           "0.595 kg",
	})

	return result
}
