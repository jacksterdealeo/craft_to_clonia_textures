package main

import (
	"fmt"
	"image"
	"image/color"

	imaging "github.com/disintegration/imaging"
)

// Used for converting the textures used on *most* worn armor.
type basicArmorConversion struct {
	inTexture string

	outHelmet     string
	outChestplate string
	outLeggings   string
	outBoots      string
}

type armorSetTextures struct {
	helmet     *image.NRGBA
	chestplate *image.NRGBA
	leggings   *image.NRGBA
	boots      *image.NRGBA
}

func armor_fixes(inPack string, outPath string) *readWriteError {
	var (
		doModded         = true
		_                = doModded // This might be a config option in the future.
		armorLocation    = craftPaths["entity>equipment>humanoid"]
		leggingsLocation = craftPaths["entity>equipment>humanoid_leggings"]
		fails            = []string{}
	)

	// Returns a Helmet, Chestplate, Leggings, Boots, and a possible error.
	// Use with power.
	var getArmorSet = func(e basicArmorConversion) (armor armorSetTextures, err error) {

		glob, err := imaging.Open(inPack + armorLocation + e.inTexture)
		if err != nil {
			return armor, fmt.Errorf("Couldn't open ~ %v (the basic version). Dropping that armor set.", e.inTexture)
		}
		// IMPORTANT: Output for armor must be scaled down to 64x32 before export!
		scale := glob.Bounds().Dx() / 64

		filter_of_choice := imaging.Lanczos

		helmet := cropToScale(glob, 0, 0, 32, 16, scale)
		if scale != 1 {
			helmet = imaging.Fit(helmet, helmet.Rect.Dx()/scale, helmet.Rect.Dy()/scale, filter_of_choice)
		}
		helmet_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
		helmet_out = imaging.Paste(helmet_out, helmet, image.Pt(32, 0))

		chestplate := cropToScale(glob, 16, 16, 56, 32, scale)
		if scale != 1 {
			chestplate = imaging.Fit(chestplate, chestplate.Rect.Dx()/scale, chestplate.Rect.Dy()/scale, filter_of_choice)
		}
		chestplate_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
		chestplate_out = imaging.Paste(chestplate_out, chestplate, image.Pt(16, 16))

		boots := cropToScale(glob, 0, 16, 16, 32, scale)
		if scale != 1 {
			boots = imaging.Fit(boots, boots.Rect.Dx()/scale, boots.Rect.Dy()/scale, filter_of_choice)
		}
		boots_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
		boots_out = imaging.Paste(boots_out, boots, image.Pt(0, 16))

		leggings, err := imaging.Open(inPack + leggingsLocation + e.inTexture)
		if err != nil {
			return armor, fmt.Errorf("Couldn't open ~ %v. (the leggings version). Dropping that armor set.", e.inTexture)
		}
		if scale != 1 {
			leggings = imaging.Fit(leggings, leggings.Bounds().Dx()/scale, leggings.Bounds().Dy()/scale, filter_of_choice)
		}
		leggings_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
		leggings_out = imaging.Paste(leggings_out, leggings, image.Pt(0, 0))

		armor = armorSetTextures{
			helmet_out,
			chestplate_out,
			leggings_out,
			boots_out,
		}
		return
	}

	var saveArmorSet = func(set basicArmorConversion) {
		if textures, err := getArmorSet(set); err != nil {
			fails = append(fails, err.Error())
		} else {
			if err := imaging.Save(textures.helmet, outPath+cloniaPaths["armor"]+set.outHelmet); err != nil {
				fails = append(fails, "Couldn't save ~"+set.outHelmet+"")
			}
			if err := imaging.Save(textures.chestplate, outPath+cloniaPaths["armor"]+set.outChestplate); err != nil {
				fails = append(fails, "Couldn't save ~"+set.outChestplate+"")
			}
			if err := imaging.Save(textures.leggings, outPath+cloniaPaths["armor"]+set.outLeggings); err != nil {
				fails = append(fails, "Couldn't save ~"+set.outLeggings+"")
			}
			if err := imaging.Save(textures.boots, outPath+cloniaPaths["armor"]+set.outBoots); err != nil {
				fails = append(fails, "Couldn't save ~"+set.outBoots+"")
			}
		}
	}

	chainmail_conversion := basicArmorConversion{
		inTexture:     "chainmail.png",
		outHelmet:     "mcl_armor_helmet_chain.png",
		outChestplate: "mcl_armor_chestplate_chain.png",
		outLeggings:   "mcl_armor_leggings_chain.png",
		outBoots:      "mcl_armor_boots_chain.png",
	}
	diamond_conversion := basicArmorConversion{
		inTexture:     "diamond.png",
		outHelmet:     "mcl_armor_helmet_diamond.png",
		outChestplate: "mcl_armor_chestplate_diamond.png",
		outLeggings:   "mcl_armor_leggings_diamond.png",
		outBoots:      "mcl_armor_boots_diamond.png",
	}
	gold_conversion := basicArmorConversion{
		inTexture:     "gold.png",
		outHelmet:     "mcl_armor_helmet_gold.png",
		outChestplate: "mcl_armor_chestplate_gold.png",
		outLeggings:   "mcl_armor_leggings_gold.png",
		outBoots:      "mcl_armor_boots_gold.png",
	}
	iron_conversion := basicArmorConversion{
		inTexture:     "iron.png",
		outHelmet:     "mcl_armor_helmet_iron.png",
		outChestplate: "mcl_armor_chestplate_iron.png",
		outLeggings:   "mcl_armor_leggings_iron.png",
		outBoots:      "mcl_armor_boots_iron.png",
	}
	netherite_conversion := basicArmorConversion{
		inTexture:     "netherite.png",
		outHelmet:     "mcl_armor_helmet_netherite.png",
		outChestplate: "mcl_armor_chestplate_netherite.png",
		outLeggings:   "mcl_armor_leggings_netherite.png",
		outBoots:      "mcl_armor_boots_netherite.png",
	}

	saveArmorSet(chainmail_conversion)
	saveArmorSet(diamond_conversion)
	saveArmorSet(gold_conversion)
	saveArmorSet(iron_conversion)
	saveArmorSet(netherite_conversion)

	if !doModded {
		if len(fails) > 0 {
			return &readWriteError{fails, "armor model textures"}
		} else {
			return nil
		}
	}

	if rose_gold_armor, err := getArmorSet(netherite_conversion); err != nil {
		fails = append(fails, "Rose Gold Armor Failed ~ "+err.Error())
	} else {
		armorToChange := []*image.NRGBA{
			rose_gold_armor.helmet,
			rose_gold_armor.chestplate,
			rose_gold_armor.leggings,
			rose_gold_armor.boots}
		for _, e := range armorToChange {
			dst := imaging.New(e.Bounds().Dx(), e.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
			dst = imaging.Overlay(dst, e, image.Pt(0, 0), 1.0)
			dst = imaging.AdjustFunc(dst,
				func(c color.NRGBA) color.NRGBA {

					average := (int(c.R) + int(c.G) + int(c.B)) / 3

					r := ((average * 100) / 30) - 30
					g := ((average * 100) / 38) - 30
					b := ((average * 100) / 37) - 30

					if r > 255 {
						r = 255
					}
					if g > 255 {
						g = 255
					}
					if b > 255 {
						b = 255
					}
					if r < 0 {
						r = 0
					}
					if g < 0 {
						g = 0
					}
					if b < 0 {
						b = 0
					}

					return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
				})
			*e = *dst
		}
		if err := imaging.Save(rose_gold_armor.helmet, outPath+cloniaPaths["rose_gold_stuff"]+"mcl_rose_gold_helmet_rose_gold.png"); err != nil {
			fails = append(fails, "Rose Gold Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(rose_gold_armor.chestplate, outPath+cloniaPaths["rose_gold_stuff"]+"mcl_rose_gold_chestplate_rose_gold.png"); err != nil {
			fails = append(fails, "Rose Gold Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(rose_gold_armor.leggings, outPath+cloniaPaths["rose_gold_stuff"]+"mcl_rose_gold_leggings_rose_gold.png"); err != nil {
			fails = append(fails, "Rose Gold Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(rose_gold_armor.boots, outPath+cloniaPaths["rose_gold_stuff"]+"mcl_rose_gold_boots_rose_gold.png"); err != nil {
			fails = append(fails, "Rose Gold Armor Failed ~ "+err.Error())
		}
	}

	if emerald_armor, err := getArmorSet(diamond_conversion); err != nil {
		fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
	} else {
		armorToChange := []*image.NRGBA{
			emerald_armor.helmet,
			emerald_armor.chestplate,
			emerald_armor.leggings,
			emerald_armor.boots}
		for _, e := range armorToChange {
			dst := imaging.New(e.Bounds().Dx(), e.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
			dst = imaging.Overlay(dst, e, image.Pt(0, 0), 1.0)
			dst = imaging.AdjustFunc(dst,
				func(c color.NRGBA) color.NRGBA {
					r := int(c.R)
					g := int(c.G)
					b := int(c.B)

					if r >= g || r >= b {
						return c
					}

					b /= 2

					return color.NRGBA{c.R, c.G, uint8(b), c.A}
				})
			*e = *dst
		}
		if err := imaging.Save(emerald_armor.helmet, outPath+cloniaPaths["emerald_stuff"]+"mcl_emerald_stuff_helmet_emerald.png"); err != nil {
			fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(emerald_armor.chestplate, outPath+cloniaPaths["emerald_stuff"]+"mcl_emerald_stuff_chestplate_emerald.png"); err != nil {
			fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(emerald_armor.leggings, outPath+cloniaPaths["emerald_stuff"]+"mcl_emerald_stuff_leggings_emerald.png"); err != nil {
			fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(emerald_armor.boots, outPath+cloniaPaths["emerald_stuff"]+"mcl_emerald_stuff_boots_emerald.png"); err != nil {
			fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
		}
	}

	if copper_armor, err := getArmorSet(iron_conversion); err != nil {
		fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
	} else {
		armorToChange := []*image.NRGBA{
			copper_armor.helmet,
			copper_armor.chestplate,
			copper_armor.leggings,
			copper_armor.boots}
		for _, e := range armorToChange {
			dst := imaging.New(e.Bounds().Dx(), e.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
			dst = imaging.Overlay(dst, e, image.Pt(0, 0), 1.0)
			dst = imaging.AdjustFunc(dst,
				func(c color.NRGBA) color.NRGBA {
					r := int(c.R)
					g := int(c.G)
					b := int(c.B)

					if (r > g+20 || r < g-20) && (r > b+20 || r < b-20) {
						return c
					}

					g = (r * 55) / 100
					b = (r * 46) / 100

					return color.NRGBA{c.R, uint8(g), uint8(b), c.A}
				})
			*e = *dst
		}
		if err := imaging.Save(copper_armor.helmet, outPath+cloniaPaths["copper_stuff"]+"mcl_copper_stuff_helmet_copper.png"); err != nil {
			fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(copper_armor.chestplate, outPath+cloniaPaths["copper_stuff"]+"mcl_copper_stuff_chestplate_copper.png"); err != nil {
			fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(copper_armor.leggings, outPath+cloniaPaths["copper_stuff"]+"mcl_copper_stuff_leggings_copper.png"); err != nil {
			fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
		}
		if err := imaging.Save(copper_armor.boots, outPath+cloniaPaths["copper_stuff"]+"mcl_copper_stuff_boots_copper.png"); err != nil {
			fails = append(fails, "Emerald Armor Failed ~ "+err.Error())
		}
	}

	if len(fails) > 0 {
		return &readWriteError{fails, "armor model textures"}
	} else {
		return nil
	}
}
