package assets

import (
	"errors"
)

type WeaponContainer struct {
	weapons []*Weapon
}

func (wc *WeaponContainer) GetAll() []*Weapon {
	return wc.weapons
}

func (wc *WeaponContainer) GetByName(name string) (*Weapon, error) {
	for _, weapon := range wc.weapons {
		if weapon.Name() == name {
			return weapon, nil
		}
	}
	return nil, errors.New("no such element")
}

var (
	WeaponCrate = WeaponContainer{weapons: initializeWeapons()}
)

func initializeWeapons() []*Weapon {
	var result = make([]*Weapon, 0)

	result = append(result, &Weapon{
		ammo:           "9 x 19 mm",
		bulletInterval: 0,
		bulletRange:    350,
		bulletSpeed:    1.2,
		description:    "The P7 is a German 9×19mm semi-automatic pistol designed by Helmut Weldle and produced by Heckler & Koch GmbH (H&K) of Oberndorf am Neckar. It was revealed to the public for the first time in 1976. Prompted by the 1972 Munich Olympics Massacre, the German police decided to replace the .32 ACP Walther PP with a similarly sized but more effective 9×19mm Parabellum pistol. The new firearm was to meet the following requirements: chamber the 9×19mm Parabellum cartridge, weigh no more than 1,000 g (35 oz), the pistol's dimensions would not exceed 180 × 130 × 34 mm, it should have a muzzle energy of no less than 500 J and a service life of at least 10,000-rounds. The pistol was also to be fully ambidextrous, safe to carry with a loaded chamber (both holstered and concealed in a pocket), and able to be quickly drawn and instantly ready to fire. Series production of the P7 started in 1979. Shortly after, the pistol was adopted by the German Federal Police's counter-terrorism unit (GSG 9) and the German Army's special forces formations. The P7 was produced primarily by H&K but also under license by the Greek defense firm Hellenic Arms Industry as well as in Mexico by the Departamento de Industria Militar (DIM), as a sidearm for general officers and staff. The pistol was also exported to several countries.",
		image:          "images/tiles/weapons/HK-P7.png",
		imageRotated:   "images/tiles/weapons/HK-P7-r.png",
		length:         "17,1 cm",
		name:           "P7",
		price:          200,
		sound:          FxPistol1,
		weaponType:     Pistol,
		weight:         "0.780 kg",
	})

	result = append(result, &Weapon{
		ammo:           "9 x 19 mm",
		bulletInterval: 0,
		bulletRange:    350,
		bulletSpeed:    1.2,
		description:    "The SIG P210 is a locked breech self loading, semi-automatic pistol (which is the normal design for common modern pistols) designed and manufactured in Neuhausen am Rheinfall (Canton of Schaffhausen, Switzerland) by SIG from 1948 until 2006. It is of all-steel construction chambered in 9×19mm Parabellum and 7.65×21mm Parabellum. It was used from 1949 to 1975 by the Swiss army and police units. It was also adopted and is still in service with the Military of Denmark, in 1951 by the German Bundespolizei and in shooting sports. The pistols were decommissioned and replaced by the SIG Sauer P220 developed in 1975. Swiss production of the P210 continued until 2006. A new model, the P210 Legend, was introduced by SIG Sauer GMBH of Germany in 2010, and another, the P210A, was introduced by SIG Sauer Inc. of New Hampshire in the United States in 2017.",
		image:          "images/tiles/weapons/SIG-P210.png",
		imageRotated:   "images/tiles/weapons/SIG-P210-r.png",
		length:         "21.5 cm",
		name:           "P210",
		price:          220,
		sound:          FxPistol2,
		weaponType:     Pistol,
		weight:         "0.970 kg",
	})

	result = append(result, &Weapon{
		ammo:           "9 x 19 mm",
		bulletInterval: 80,
		bulletRange:    300,
		bulletSpeed:    1.3,
		description:    "The Uzi is a family of Israeli open-bolt, blowback-operated submachine guns. Smaller variants are often considered to be machine pistols. The Uzi was one of the first weapons to use a telescoping bolt design which allows the magazine to be housed in the pistol grip for a shorter weapon. The first Uzi submachine gun was designed by Major Uziel Gal in the late 1940s. The prototype was finished in 1950. First introduced to IDF special forces in 1954, the weapon was placed into general issue two years later. The Uzi has found use as a personal defense weapon by rear-echelon troops, officers, artillery troops and tankers, as well as a frontline weapon by elite light infantry assault forces. The Uzi has been exported to over 90 countries. Over its service lifetime, it has been manufactured by Israel Military Industries, FN Herstal, and other manufacturers. From the 1960s through the 1980s, more Uzi submachine guns were sold to more military, law enforcement and security markets than any other submachine gun ever made.",
		image:          "images/tiles/weapons/Uzi.png",
		imageRotated:   "images/tiles/weapons/Uzi-r.png",
		length:         "47 cm",
		name:           "Uzi",
		price:          1400,
		sound:          FxUzi,
		weaponType:     Automatic,
		weight:         "3.6 kg",
	})

	result = append(result, &Weapon{
		ammo:           "9 x 19 mm",
		bulletInterval: 70,
		bulletRange:    400,
		bulletSpeed:    1.3,
		description:    "The MP5 (Maschinenpistole 5) is a 9x19mm Parabellum submachine gun, developed in the 1960s by a team of engineers from the German small arms manufacturer Heckler & Koch GmbH (H&K) of Oberndorf am Neckar. There are over 100 variants of the MP5, including some semi-automatic versions. The MP5 is one of the most widely used submachine guns in the world, having been adopted by 40 nations and numerous military, law enforcement, intelligence, and security organizations. It was widely used by SWAT teams in North America, but has largely been supplanted by M16 variants in the 21st century. In 1976, a shortened machine pistol version of the MP5A2 was introduced; the MP5K (K from the German word kurz) was designed for close quarters battle use by clandestine operations and special services. The MP5K does not have a shoulder stock (the receiver end was covered with a flat end cap, featuring a buffer on the inside and a sling loop on the outside), and the bolt and receiver were shortened at the rear. The resultant lighter bolt led to a higher rate of fire than the standard MP5. The barrel, cocking handle and its cover were shortened and a vertical foregrip was used to replace the standard handguard. The barrel ends at the base of the front sight, which prevents the use of any sort of muzzle device.",
		image:          "images/tiles/weapons/MP5K.png",
		imageRotated:   "images/tiles/weapons/MP5K-r.png",
		length:         "32.5 cm",
		name:           "MP5",
		price:          1600,
		sound:          FxMp5,
		weaponType:     Automatic,
		weight:         "2 kg",
	})

	result = append(result, &Weapon{
		ammo:           "4.6 x 30 mm",
		bulletInterval: 80,
		bulletRange:    400,
		bulletSpeed:    1.3,
		description:    "The Heckler & Koch MP7 (Maschinenpistole 7) is a personal defence weapon (PDW) chambered for the HK 4.6×30mm armor-piercing cartridge designed by German defence manufacturer Heckler & Koch. It was designed together with the new cartridge to meet NATO requirements published in 1989, as these requirements call for a compact PDW-class firearm, with a greater ability to defeat body armour than the current submachine guns using conventional pistol cartridges. The MP7 went into production in 2001, and is a direct rival to the FN P90, also developed in response to NATO's requirement. The weapon has been revised since its introduction and the latest production variants are the MP7A1 and MP7A2. The proliferation of high-quality body armor has begun to make guns that fire pistol ammunition (such as Heckler & Koch's earlier MP5 submachine gun and USP pistol) ineffective. In response to this trend, Heckler & Koch designed the MP7 (along with the cancelled UCP pistol, which uses the same ammunition) to penetrate body armor while being small enough to be used in place of either a pistol or a submachine gun.",
		image:          "images/tiles/weapons/HK-MP7.png",
		imageRotated:   "images/tiles/weapons/HK-MP7-r.png",
		length:         "41.5 cm",
		name:           "MP7",
		price:          2800,
		sound:          FxMp7,
		weaponType:     Automatic,
		weight:         "1.9 kg",
	})

	result = append(result, &Weapon{
		ammo:           "7.62 x 51 mm",
		bulletInterval: 80,
		bulletRange:    500,
		bulletSpeed:    1.4,
		description:    "The FAL (French: Fusil Automatique Léger, English: Light Automatic Rifle) is a battle rifle designed by Belgian small arms designer Dieudonné Saive and manufactured by FN Herstal. During the Cold War the FAL was adopted by many countries of the North Atlantic Treaty Organization (NATO), with the notable exception of the United States. It is one of the most widely used rifles in history, having been used by more than 90 countries. Because of its prevalence and widespread usage among the militaries of many NATO and first world countries during the Cold War, it was given the title 'The right arm of the Free World'. It is chambered for the 7.62×51mm NATO cartridge (although originally designed for the .280 British intermediate cartridge). The British Commonwealth variant of the FAL was redesigned from FN's metrical FAL into British imperial units and was produced under licence as the L1A1 Self-Loading Rifle.",
		image:          "images/tiles/weapons/FAL-G1.png",
		imageRotated:   "images/tiles/weapons/FAL-G1-r.png",
		length:         "109 cm",
		name:           "FN FAL",
		price:          1600,
		sound:          FxFnfal,
		weaponType:     Automatic,
		weight:         "4.25 kg",
	})

	result = append(result, &Weapon{
		ammo:           "7.62 x 51 mm",
		bulletInterval: 100,
		bulletRange:    550,
		bulletSpeed:    1.4,
		description:    "The ArmaLite AR-10 is a 7.62×51mm NATO battle rifle developed by Eugene Stoner in the late 1950s and manufactured by ArmaLite, then a division of the Fairchild Aircraft Corporation. When first introduced in 1956, the AR-10 used an innovative straight-line barrel/stock design with phenolic composite and forged alloy parts resulting in a small arm significantly easier to control in automatic fire and over 1 lb (0.45 kg) lighter than other infantry rifles of the day. Over its production life, the original AR-10 was built in relatively small numbers, with fewer than 10,000 rifles assembled. However, the ArmaLite AR-10 would become the progenitor for a wide range of firearms. In 1957, the basic AR-10 design was rescaled and substantially modified by ArmaLite to accommodate the .223 Remington cartridge, and given the designation ArmaLite AR-15. In 1959, ArmaLite sold its rights to the AR-10 and AR-15 to Colt Firearms due to financial difficulties, and limitations in terms of manpower and production capacity. After modifications (most notably, the charging handle was re-located from under the carrying handle like AR-10 to the rear of the receiver), the new redesigned rifle was subsequently adopted by the U.S. military as the M16 Rifle. Colt continued to use the AR-15 trademark for its line of semi-automatic-only rifles, which it marketed to civilian and law-enforcement customers as the Colt AR-15. With the expiration of its patent, other manufacturers began producing their own variants, known as AR-15 style rifles.",
		image:          "images/tiles/weapons/AR10.png",
		imageRotated:   "images/tiles/weapons/AR10-r.png",
		length:         "105 cm",
		name:           "AR-10",
		price:          1800,
		sound:          FxAr10,
		weaponType:     Automatic,
		weight:         "4.3 kg",
	})

	result = append(result, &Weapon{
		ammo:           "5.45 x 39 mm",
		bulletInterval: 110,
		bulletRange:    500,
		bulletSpeed:    1.4,
		description:    "The AK-74 (Kalashnikov automatic rifle model 1974) is an assault rifle developed in the early 1970s by Russian weapons designer Mikhail Kalashnikov to replace the earlier AKM (itself a refined version of the AK-47). It uses a smaller 5.45×39mm cartridge, replacing the 7.62×39mm chambering of earlier Kalashnikov-pattern weapons. The rifle first saw service with Soviet forces in the 1979 Afghanistan conflict. The head of the Afghan bureau of the Pakistani Inter-Services Intelligence claimed that the CIA paid $5,000 for the first AK-74 captured by the Mujahideen during the Soviet–Afghan War. Today, the rifle is used by most countries of the former Soviet Union. Licensed copies were produced in Bulgaria (AK-74, AKS-74 and AKS-74U), and the former East Germany (MPi-AK-74N, MPi-AKS-74N, MPi-AKS-74NK).",
		image:          "images/tiles/weapons/AK74.png",
		imageRotated:   "images/tiles/weapons/AK74-r.png",
		length:         "94 cm",
		name:           "AK-74",
		price:          3000,
		sound:          FxAk47,
		weaponType:     Automatic,
		weight:         "3.3 kg",
	})

	result = append(result, &Weapon{
		ammo:           "5.56 x 45 mm",
		bulletInterval: 115,
		bulletRange:    550,
		bulletSpeed:    1.4,
		description:    "The G36 is a 5.56×45mm assault rifle, designed in the early 1990s by Heckler & Koch in Germany as a replacement for the heavier 7.62mm G3 battle rifle. It was accepted into service with the Bundeswehr in 1997, replacing the G3. The G36 is gas-operated and feeds from a 30-round detachable box magazine or 100-round C-Mag drum magazine.",
		image:          "images/tiles/weapons/G36.png",
		imageRotated:   "images/tiles/weapons/G36-r.png",
		length:         "99.9 cm",
		name:           "G36",
		price:          3000,
		sound:          FxG36,
		weaponType:     Automatic,
		weight:         "3.63 kg",
	})

	result = append(result, &Weapon{
		ammo:           "5.56 x 45 mm",
		bulletInterval: 100,
		bulletRange:    550,
		bulletSpeed:    1.4,
		description:    "The Heckler & Koch HK416 is an assault rifle designed and manufactured by the German company Heckler & Koch. Although its design is based on the AR-15 class of weapons, specifically the Colt M4 carbine family issued to the U.S. military, it uses a proprietary short-stroke gas piston system derived from the ArmaLite AR-18 (the same system was also used in Heckler & Koch's earlier G36 family of rifles). It is the standard assault rifle of the Norwegian Armed Forces, selected by the French Armed Forces to replace the FAMAS, and was the weapon used by SEAL Team Six to kill Osama Bin Laden in 2011.",
		image:          "images/tiles/weapons/HK416N.png",
		imageRotated:   "images/tiles/weapons/HK416N-r.png",
		length:         "103.7 cm",
		name:           "G95K",
		price:          9800,
		sound:          FxG95k,
		weaponType:     Automatic,
		weight:         "3.850 kg",
	})

	result = append(result, &Weapon{
		ammo:           "7.62 x 63 mm",
		bulletInterval: 125,
		bulletRange:    600,
		bulletSpeed:    1.4,
		description:    "The Browning Automatic Rifle (BAR) is a family of American automatic rifles and machine guns used by the United States and numerous other countries during the 20th century. The primary variant of the BAR series was the M1918, chambered for the .30-06 Springfield rifle cartridge and designed by John Browning in 1917 for the American Expeditionary Forces in Europe as a replacement for the French-made Chauchat and M1909 Benét–Mercié machine guns that US forces had previously been issued. The BAR was designed to be carried by infantrymen during an assault advance while supported by the sling over the shoulder, or to be fired from the hip. This is a concept called 'walking fire'—thought to be necessary for the individual soldier during trench warfare. The BAR never entirely lived up to the original hopes of the war department as either a rifle or a machine gun. The U.S. Army, in practice, used the BAR as a light machine gun, often fired from a bipod (introduced on models after 1938). A variant of the original M1918 BAR, the Colt Monitor Machine Rifle, remains the lightest production automatic firearm chambered for the .30-06 Springfield cartridge, though the limited capacity of its standard 20-round magazine tended to hamper its utility in that role. Although the weapon did see some action in World War I, the BAR did not become standard issue in the US Army until 1938, when it was issued to squads as a portable light machine gun. The BAR saw extensive service in both World War II and the Korean War and saw limited service in the Vietnam War. The US Army began phasing out the BAR in the late 1950s, when it was intended to be replaced by a squad automatic weapon (SAW) variant of the M14, and was without a portable light machine gun until the introduction of the M60 machine gun in 1957.",
		image:          "images/tiles/weapons/BarM1918VWM.png",
		imageRotated:   "images/tiles/weapons/BarM1918VWM-r.png",
		length:         "119.4 cm",
		name:           "B.A.R.",
		price:          11400,
		sound:          FxBar,
		weaponType:     Automatic,
		weight:         "7.25 kg",
	})

	result = append(result, &Weapon{
		ammo:           "5.56 x 45 mm",
		bulletInterval: 90,
		bulletRange:    600,
		bulletSpeed:    1.4,
		description:    "The Heckler & Koch MG4 (also known as the HK123) is a belt-fed 5.56 mm light machine gun designed and developed by German firearm manufacturer Heckler & Koch. The weapon was developed in the late 1990s and was first seen publicly in September 2001. It has been selected to replace the 7.62 mm MG3 general-purpose machine gun in the Bundeswehr at the squad support level; it will complement the MG3 in other roles. The MG4 will also be the secondary armament of the new Puma infantry fighting vehicle. Overall, it is designed to be light, provide maximum safety to the user and function reliably under adverse conditions using a wide range of ammunition from different manufacturers, without the need to adjust the gas system. The machine gun was initially known as the MG43 prior to its adoption by the Bundeswehr.",
		image:          "images/tiles/weapons/HK-MG4.png",
		imageRotated:   "images/tiles/weapons/HK-MG4-r.png",
		length:         "103 cm",
		name:           "MG4",
		price:          12400,
		sound:          FxMg4,
		weaponType:     Automatic,
		weight:         "7.35 kg",
	})

	result = append(result, &Weapon{
		ammo:           "7.62 x 51 mm",
		bulletInterval: 128,
		bulletRange:    700,
		bulletSpeed:    1.4,
		description:    "The HK21 is a German 7.62 mm general-purpose machine gun, developed in 1961 by small arms manufacturer Heckler & Koch and based on the G3 battle rifle. The weapon is in use with the armed forces of several Asian, African and Latin American countries. It was also license-manufactured by Fábrica de Braço de Prata in Portugal as the m/968 and in Mexico by SEDENA as the MG21. In the German military (Bundeswehr) and the federal police (Bundespolizei) it is designated G8.",
		image:          "images/tiles/weapons/HK-21-LMG.png",
		imageRotated:   "images/tiles/weapons/HK-21-LMG-r.png",
		length:         "102.1 cm",
		name:           "HK21",
		price:          15600,
		sound:          FxHk21,
		weaponType:     Automatic,
		weight:         "7.92 kg",
	})

	result = append(result, &Weapon{
		ammo:           "7.62 x 51 mm",
		bulletInterval: 65,
		bulletRange:    850,
		bulletSpeed:    1.6,
		description:    "The MG 42 (Maschinengewehr 42) is a 7.92×57mm Mauser general-purpose machine gun designed in Nazi Germany and used extensively by the Wehrmacht and the Waffen-SS during the second half of World War II. It was intended to replace the earlier MG 34, which was more expensive and took much longer to produce, but both weapons were produced until the end of World War II. Designed to be low-cost and easy to build, the MG 42 proved to be highly reliable and easy to operate. It is most notable for its very high cyclic rate for a gun using full power service cartridges, averaging about 1,200 rounds per minute compared to around 850 for the MG 34, and perhaps 450 to 600 for other common machine guns like the M1919 Browning or Bren. This ability made it extremely effective in providing suppressive fire, and its unique sound led to it being nicknamed \"Hitler's buzzsaw\". The MG 42 was adopted by several armed organizations after the war, and was both copied and built under licence. The MG 42's lineage continued past Nazi Germany's defeat, forming the basis for the nearly identical MG1 (MG 42/59), chambered in 7.62×51mm NATO, which subsequently evolved into the MG1A3, and later the Bundeswehr's MG 3, Italian MG 42/59 and Austrian MG 74. It also spawned the Yugoslav unlicensed nearly identical Zastava M53. The MG 42 lent many design elements to the Swiss MG 51 and SIG MG 710-3, American M60 and Belgian MAG general-purpose machine guns and the Spanish 5.56×45mm NATO Ameli light machine gun though these machine guns feature other operating mechanisms or actions than the MG 42 and its lineage.",
		image:          "images/tiles/weapons/MG42.png",
		imageRotated:   "images/tiles/weapons/MG42-r.png",
		length:         "123 cm",
		name:           "MG 42 (converted to NATO caliber)",
		price:          20800,
		sound:          FxMg42,
		weaponType:     Automatic,
		weight:         "10.6 kg",
	})

	result = append(result, &Weapon{
		ammo:           "89 mm",
		bulletInterval: 0,
		bulletRange:    750,
		bulletSpeed:    1.1,
		description:    "Bazooka is the common name for a man-portable recoilless anti-tank rocket launcher weapon, widely deployed by the United States Army. Also referred to as the 'Stovepipe', the innovative bazooka was among the first generation of rocket-propelled anti-tank weapons used in infantry combat. Featuring a solid-propellant rocket for propulsion, it allowed for high-explosive anti-tank (HEAT) warheads to be delivered against armored vehicles, machine gun nests, and fortified bunkers at ranges beyond that of a standard thrown grenade or mine. The universally-applied nickname arose from the M1 variant's vague resemblance to the musical instrument called a 'bazooka' invented and popularized by 1930s U.S. comedian Bob Burns. During World War II, German armed forces captured several bazookas in early North Africa and Eastern Front encounters and soon reverse engineered their own version, increasing the warhead diameter to 8.8 cm (among other minor changes) and widely issuing it as the Raketenpanzerbüchse 'Panzerschreck'. Near the end of the war, the Japanese developed a similar weapon, the Type 4 70 mm AT Rocket Launcher, which featured a rocket propelled grenade of a different design. The term 'bazooka' still sees informal use as a generic term referring to any ground-to-ground shoulder-fired missile weapon (mainly rocket propelled grenade launchers or recoilless rifles).",
		image:          "images/tiles/weapons/M20-Bazooka.png",
		imageRotated:   "images/tiles/weapons/M20-Bazooka-r.png",
		length:         "152 cm",
		name:           "M20A1B1 'Super Bazooka'",
		price:          5200,
		sound:          FxRocketLauncher,
		weaponType:     RPG,
		weight:         "5.9 kg",
	})

	result = append(result, &Weapon{
		ammo:           "60 mm",
		bulletInterval: 0,
		bulletRange:    900,
		bulletSpeed:    1.2,
		description:    "The Panzerfaust 44 2A1 (also called Lanze, Leichte Panzerfaust, or Panzerfaust 2) is a recoilless anti-tank weapon. The anti-tank grenade launcher was originally developed around 1960 and put into service by the Bundeswehr in 1963. It was developed to provide West German infantry with a modern replacement for the Bazooka that they had previously used. As such, it was the first German antitank rocket developed after World War II, a conflict in which German hand-held antitank weapons such as the Panzerfaust played a prominent role during 1944–45. The Lanze was a product of a period in which the German army was re-equipped with locally developed arms and equipment and retired the aging U.S. gear that had formed their initial arsenal. The full designation name by the German Army is Panzerfaust 44mm DM2 Ausführung 1 Lanze.",
		image:          "images/tiles/weapons/Panzerfaust-44.png",
		imageRotated:   "images/tiles/weapons/Panzerfaust-44-r.png",
		length:         "88 cm",
		name:           "Panzerfaust 44",
		price:          5800,
		sound:          FxRocketLauncher,
		weaponType:     RPG,
		weight:         "7.8 kg",
	})

	result = append(result, &Weapon{
		ammo:           "110 mm",
		bulletInterval: 0,
		bulletRange:    900,
		bulletSpeed:    1.3,
		description:    "The Panzerfaust 3 is a modern disposable recoilless anti-tank weapon, which was developed between 1978 and 1985 and put into service by the Bundeswehr in 1992. It was first ordered in 1973 to provide West German infantry with an effective weapon against contemporary Soviet armour, thereby replacing West Germany's aging PzF 44 Light Lanze launchers and the heavy Carl Gustaf 84 mm anti-tank recoilless rifle manufactured in Sweden. The Panzerfaust 3 is operated by at least 11 countries and has first seen combat in Afghanistan.",
		image:          "images/tiles/weapons/Panzerfaust-3.png",
		imageRotated:   "images/tiles/weapons/Panzerfaust-3-r.png",
		length:         "88 cm",
		name:           "Panzerfaust 3",
		price:          4200,
		sound:          FxRocketLauncher,
		weaponType:     RPG,
		weight:         "12.9 kg",
	})

	result = append(result, &Weapon{
		ammo:           "Tar and petrol mixture",
		bulletInterval: 100,
		bulletRange:    400,
		bulletSpeed:    0.7,
		description:    "The Flammenwerfer 41, or FmW 41 was the standard German flamethrower beginning in 1941 and an upgraded version of the earlier Flammenwerfer 35. It was designed to clear enemy trenches and buildings in highly fortified areas. The Flammenwerfer 35 had been a long needed update on a World War I version simplifying the weapon to a single operating from three and increasing its range. However the 36 kg system was cumbersome to carry and difficult to use. Thus spurred the development of the FmW 41. Similar to many other designs of the time, the FmW 41 used a hydrogen torch to ignite a tar and petrol mixture which was fired from a hand-held torch attached to a tank. The petrol and propellant were carried in separate tanks carried on the back which held 11.8 litres (2.6 imp gal; 3.1 US gal) of tar and petrol mixture called Flammöl 19. The FmW 41 proved to be more reliable and easier to operate than its predecessor, it had an increased range of 32 meters and was lighter weighing in at 28.7 kilograms. Problems surfaced against Soviet troops during the winter of 1941 as its lighting mechanism was unable to cope with the cold weather conditions. Later versions of the weapon replaced the hydrogen torch with a cartridge based system which proved more effective. Flamethrower troops are often targeted by enemy troops as they take a heavy psychological toll on enemy morale in addition to being effective weapons against well entrenched enemy troops. As such, these weapons were often dressed to look like standard infantry rifles in an attempt to disguise operators.",
		image:          "images/tiles/weapons/Flammenwerfer-42.png",
		imageRotated:   "images/tiles/weapons/Flammenwerfer-42-r.png",
		length:         "40 cm",
		name:           "Flammenwerfer 41",
		price:          19800,
		sound:          FxNone,
		weaponType:     Flamethrower,
		weight:         "22 kg",
	})

	return result
}
