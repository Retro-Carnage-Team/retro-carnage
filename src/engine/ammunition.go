package engine

type Ammunition struct {
	Description string
	Image       string
	MaxCount    int32
	Name        string
	PackageSize int32
	Price       int32
}

var Ammunitions = initializeAmmunition()

func initializeAmmunition() []*Ammunition {
	var result = make([]*Ammunition, 0)

	result = append(result, &Ammunition{
		Description: "The 9×19mm Parabellum is a firearms cartridge that was designed by Georg Luger and introduced in 1902 by the German weapons manufacturer Deutsche Waffen- und Munitionsfabriken (DWM) (German Weapons and Munitions Factory) for its Luger semi-automatic pistol. For this reason, it is designated as the 9mm Luger by the Sporting Arms and Ammunition Manufacturers' Institute, and the 9 mm Luger by the Commission Internationale Permanente pour l'Epreuve des Armes à Feu Portatives. The name Parabellum is derived from the Latin: Si vis pacem, para bellum (\"If you seek peace, prepare for war\"), which was the motto of DWM. Under STANAG 4090, it is a standard cartridge for NATO forces as well as many non-NATO countries. According to the 2014 edition of Cartridges of the World, the 9×19mm Parabellum is the world's most popular and widely used military handgun and submachine gun cartridge. In 2007, Newsweek claimed that about 60 percent of the firearms in use by police are 9mms and credited 9×19mm Parabellum pistol sales with making semiautomatic pistols more popular than revolvers. The popularity of this cartridge can be attributed to the widely held conviction that it is effective in police and self-defense use. Its low cost and wide availability contribute to the caliber's continuing popularity.",
		Image:       "images/tiles/weapons/9x19.png",
		MaxCount:    10000,
		Name:        "9 x 19 mm",
		PackageSize: 100,
		Price:       100,
	})

	result = append(result, &Ammunition{
		Description: "The HK 4.6×30mm cartridge is a type of ammunition primarily used in the Heckler & Koch MP7 Personal defense weapon (PDW). It is designed to minimize weight and recoil while increasing penetration of body armor. It features a bottlenecked case and a pointed, steel-core, brass-jacketed bullet. Compared to standard intermediate cartridges one can carry more 4.6×30mm ammunition due to the lighter weight and relative small dimensions of the cartridge. Also, due to the lighter weight of the bullet, aiming in rapid fire is much easier as recoil depends much on the weight of the bullet. CRISAT testing shows that because of the smaller diameter and high projectile velocity of the round, body armor penetration is higher than that of traditional handgun projectiles. A series of tests performed by NATO in the United Kingdom and France indicated that 5.7×28mm was a superior cartridge. The results of the NATO tests were analyzed by a group of experts from France, the United States, Canada, and the United Kingdom and the group's conclusion was that the 5.7×28mm was undoubtedly the more efficient cartridge.",
		Image:       "images/tiles/weapons/4.6x30.png",
		MaxCount:    10000,
		Name:        "4.6 x 30 mm",
		PackageSize: 100,
		Price:       110,
	})

	result = append(result, &Ammunition{
		Description: "The 5.45×39mm cartridge is a rimless bottlenecked intermediate cartridge. It was introduced into service in 1974 by the Soviet Union for use with the new AK-74. The 5.45×39mm gradually supplemented, then largely replaced the 7.62×39mm cartridge in Soviet and Warsaw Pact service as the primary military service rifle cartridge.",
		Image:       "images/tiles/weapons/5.45x39.png",
		MaxCount:    10000,
		Name:        "5.45 x 39 mm",
		PackageSize: 100,
		Price:       120,
	})

	result = append(result, &Ammunition{
		Description: "The 5.56×45mm NATO is a rimless bottlenecked intermediate cartridge family developed in the late 1970s in Belgium by FN Herstal. It consists of the SS109, SS110, and SS111 cartridges. On 28 October 1980 under STANAG 4172 it was standardized as the second standard service rifle cartridge for NATO forces as well as many non-NATO countries. The 5.56×45mm NATO cartridge family was derived from, but is not identical to, the .223 Remington cartridge designed by Remington Arms in the early 1960s.",
		Image:       "images/tiles/weapons/5.56x45.png",
		MaxCount:    10000,
		Name:        "5.56 x 45 mm",
		PackageSize: 100,
		Price:       120,
	})

	result = append(result, &Ammunition{
		Description: "The 7.62×51mm NATO is a rimless bottlenecked rifle cartridge developed in the 1950s as a standard for small arms among NATO countries. It is sometimes confused with the similarly named Russian 7.62×54mmR cartridge, a slightly longer, rimmed cartridge. It was introduced in U.S. service in the M14 rifle and M60 machine gun in the late 1950s. The M14 was superseded in U.S. service as the infantry adopted the 5.56×45mm NATO M16 rifle. However, the M14 and many other firearms that use the 7.62×51mm NATO round remain in service, especially in the case of various marksman/sniper rifles, medium machine guns/general-purpose machine guns such as the M240, and various rifles in use by special operations forces. The cartridge is used both by infantry and on mounted and crew-served weapons mounted to vehicles, aircraft and ships. Although not identical, the 7.62×51mm NATO and the commercial .308 Winchester cartridges are similar enough that they can be loaded into rifles chambered for the other round, but the Winchester .308 cartridges are typically loaded to higher pressures than 7.62×51mm NATO service cartridges. Even though the Sporting Arms and Ammunition Manufacturers' Institute (SAAMI) does not consider it unsafe to fire the commercial round in weapons chambered for the NATO round, there is significant discussion about compatible chamber and muzzle pressures between the two cartridges based on powder loads and wall thicknesses of the military compared to commercial rounds.",
		Image:       "images/tiles/weapons/7.62x51.png",
		MaxCount:    10000,
		Name:        "7.62 x 51 mm",
		PackageSize: 100,
		Price:       130,
	})

	result = append(result, &Ammunition{
		Description: "The .30-06 Springfield cartridge (pronounced 'thirty-aught-six'), 7.62×63mm in metric notation, was introduced to the United States Army in 1906 and later standardized; it remained in use until the late-1970s. The '.30' refers to the caliber of the bullet in inches. The '06' refers to the year the cartridge was adopted, 1906. It replaced the .30-03, 6mm Lee Navy, and .30-40 Krag cartridges. The .30-06 remained the U.S. Army's primary rifle and machine gun cartridge for nearly 50 years before being replaced by the 7.62×51mm NATO and 5.56×45mm NATO, both of which remain in current U.S. and NATO service. It remains a very popular sporting round, with ammunition produced by all major manufacturers.",
		Image:       "images/tiles/weapons/7.63x63.png",
		MaxCount:    10000,
		Name:        "7.62 x 63 mm",
		PackageSize: 100,
		Price:       130,
	})

	result = append(result, &Ammunition{
		Description: "The 3.5-inch ammunition was in principle an enlarged 2.36-inch ammunition; at 60 cm it was only slightly longer, but at 4 kg it weighed almost three times as much. The rocket motor accelerated to 100 m/s muzzle velocity, resulting in a maximum range of 820 m. The warhead was tapered, but rounded at the tip. The M28 shaped-charge ammunition (in further versions M28A1 and M28A2) was loaded with 0.9 kg of composition B, resulting in a penetration of 30 cm of armour and 60 cm of concrete. The M29 was the corresponding practice ammunition without explosives in the warhead. The M30 was a smoke or incendiary grenade, loaded with 1 kg of white phosphorus. The 3.5-inch ammunition was produced by the Picatinny Arsenal.",
		Image:       "images/tiles/weapons/m28a2.png",
		MaxCount:    100,
		Name:        "89 mm",
		PackageSize: 1,
		Price:       2600,
	})

	result = append(result, &Ammunition{
		Description: "This 60mm projectile for the Panzerfaust 44 is effective against holding tanks up to 200 meters and against moving tanks up to 150 meters distance. The penetrating power is up to 375 mm of rolled armour steel.",
		Image:       "images/tiles/weapons/DM12A1.png",
		MaxCount:    100,
		Name:        "60 mm",
		PackageSize: 1,
		Price:       2800,
	})

	result = append(result, &Ammunition{
		Description: "The Panzerfaust 3 consists of the reusable grip and the cartridge. The cartridge in turn is divided into the launch tube and the bullet. The projectile head is located outside the barrel, which means that the dimensions of the head can be selected independently of the calibre of the launch barrel. It consists of two parts - the spacer tube and the active charge. Against armoured targets such as tanks, the spacer tube is pulled out, against soft targets such as trucks and buildings it remains inserted. This standard cartridge works on the hollow charge principle and can penetrate up to 700 mm of armour steel.",
		Image:       "images/tiles/weapons/BS71.png",
		MaxCount:    100,
		Name:        "110 mm",
		PackageSize: 1,
		Price:       3000,
	})

	result = append(result, &Ammunition{
		Description: "The petrol and propellant for the Flammenwerfer 41 comes in separate tanks carried on the back. These tanks hold 11.8 litres of tar and petrol mixture called Flammoel 19. The Flammenwerfer 41 proved to be more reliable and easier to operate than its predecessor, it had an increased range of 32 metres and was lighter weighing in at 28.7 kilograms.",
		Image:       "images/tiles/weapons/Napalm.png",
		MaxCount:    1000,
		Name:        "Tar and petrol mixture",
		PackageSize: 50,
		Price:       1000,
	})

	return result
}
