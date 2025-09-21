package main

import (
	"image"
	"image/color"
	"strconv"

	data "codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

var (
	flipH = imaging.FlipH
	flipV = imaging.FlipV
)

func flipHV(img image.Image) *image.NRGBA {
	return imaging.FlipH(imaging.FlipV(img))
}

func cropToScale(img image.Image, x1, y1, x2, y2, scale int) *image.NRGBA {
	return imaging.Crop(img, image.Rectangle{
		image.Point{x1 * scale, y1 * scale}, image.Point{x2 * scale, y2 * scale}})
}

/*
func animated_texture_fix(inName string, outName string) *readWriteError {
	fails := []string{}
	animated := [...][4]string{
		{"block", "respawn_anchor_top.png", "beds", "respawn_anchor_top_on.png"},         //
		{"block", "soul_fire_0.png", "blackstone", "soul_fire_basic_flame_animated.png"}, //
		{"block", "fire_0.png", "fire", "fire_basic_flame_animated.png"},                 //
		{"block", "fire_0.png", "fire", "mcl_burning_entity_flame_animated.png"},         //
		{"block", "fire_0.png", "fire", "mcl_burning_hud_flame_animated.png"},            //
		{"block", "magma.png", "nether", "mcl_nether_magma.png"},                         //
	}

	for _, e := range animated {
		img, err := imaging.Open(inName + craftPaths[e[0]] + e[1])
		if err != nil {
			fails = append(fails, e[0]+"::"+e[1]+" failed to open!")
		} else {
			if err = imaging.Save(img, outName+cloniaPaths[e[2]]+e[3]); err != nil {
				fails = append(fails, e[3]+" failed to save!")
			}
		}
	}
	if len(fails) > 0 {
		return &readWriteError{fails, "animated textures"}
	} else {
		return nil
	}
}
*/

func do_fixes(inPack string, outPack string) *readWriteError {
	fails := []string{}

	func() { // special slabs
		t := [...]data.SimpleConversion{
			{"block", "polished_andesite.png", "core", "mcl_stairs_andesite_smooth_slab.png", 1},
			{"block", "polished_diorite.png", "core", "mcl_stairs_diorite_smooth_slab.png", 1},
			{"block", "polished_granite.png", "core", "mcl_stairs_granite_smooth_slab.png", 1},
			{"block", "gold_block.png", "xstairs", "mcl_stairs_gold_block_slab.png", 1},
			{"block", "iron_block.png", "xstairs", "mcl_stairs_iron_block_slab.png", 1},
			{"block", "lapis_block.png", "xstairs", "mcl_stairs_lapis_block_slab.png", 1},
		}
		for _, e := range t {
			block, err := imaging.Open(inPack + craftPaths[e.InPath] + e.InTexture)
			_ = block
			if err != nil {
				fails = append(fails, e.InTexture+"failed to open!")
			} else {
				scale := block.Bounds().Dx() / 16
				dst := imaging.New(block.Bounds().Dx(), block.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Paste(dst, block, image.Pt(0, -8*scale))
				dst = imaging.Paste(dst, block, image.Pt(0, 8*scale))

				top := imaging.Crop(block, image.Rect(0, 0, block.Bounds().Dx(), 1))
				bottom := imaging.Crop(block, image.Rect(0, 15*scale, block.Bounds().Dx(), block.Bounds().Dy()))
				dst = imaging.Paste(dst, top, image.Pt(0, 0))
				dst = imaging.Paste(dst, bottom, image.Pt(0, 15*scale))

				if err := imaging.Save(dst, outPack+cloniaPaths[e.OutPath]+e.OutTexture); err != nil {
					fails = append(fails, e.OutTexture+" failed to save!")
				}
			}
		}
	}()

	func() { // green plants
		t := [...]data.SimpleConversion{
			{"block", "vine.png", "core", "mcl_core_vine.png", 1},
			{"block", "lily_pad.png", "flowers", "flowers_waterlily.png", 1},
			{"block", "large_fern_top.png", "flowers", "mcl_flowers_double_plant_fern_inv.png", 1},
			{"block", "tall_grass_top.png", "flowers", "mcl_flowers_double_plant_grass_inv.png", 1},
			{"block", "fern.png", "flowers", "mcl_flowers_fern_inv.png", 1},
			{"block", "short_grass.png", "flowers", "mcl_flowers_tallgrass_inv.png", 1},
		}
		for _, e := range t {
			grayImage, err := imaging.Open(inPack + e.ReadPath())
			if err != nil {
				fails = append(fails, e.InTexture+" failed to open!")
			} else {
				dst := imaging.New(grayImage.Bounds().Dx(), grayImage.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Overlay(dst, grayImage, image.Point{0, 0}, 1.0)
				dst = mtg_green_it(dst)
				if err = imaging.Save(dst, outPack+e.SavePath()); err != nil {
					fails = append(fails, e.OutTexture+" failed to save!")
				}
			}
		}
	}()

	func() { // offhand_slot
		t := "hotbar_offhand_left.png"
		if offHand, err := imaging.Open(inPack + craftPaths["hud"] + t); err != nil {
			fails = append(fails, t+" failed to open!")
		} else {
			// 29 x 24
			scale := offHand.Bounds().Dx() / 29
			// 22 x 22
			dst := imaging.New(22*scale, 22*scale, color.NRGBA{0, 0, 0, 0})
			dst = imaging.Paste(dst, offHand, image.Pt(0, -1*scale))
			if err2 := imaging.Save(dst, outPack+cloniaPaths["offhand"]+"mcl_offhand_slot.png"); err2 != nil {
				fails = append(fails, t+" failed to save!")
			}
		}
	}()

	if len(fails) > 0 {
		return &readWriteError{fails, "patched textures"}
	} else {
		return nil
	}
}

func campfire_fix(inPath string, outPath string) *readWriteError {
	fails := []string{}

	campfire_log_lit, err := imaging.Open(inPath + "campfire_log_lit.png")
	if err != nil {
		fails = append(fails, "campfires::campfire_log_lit.png failed to open!")
	} else {
		dst := imaging.New(campfire_log_lit.Bounds().Dx()*2, campfire_log_lit.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, campfire_log_lit, image.Pt(0, 0))
		if err := imaging.Save(dst, outPath+"mcl_campfires_campfire_log_lit.png"); err != nil {
			fails = append(fails, "mcl_campfires_campfire_log_lit.png failed to save!")
		}
	}

	soul_campfire_log_lit, err := imaging.Open(inPath + "soul_campfire_log_lit.png")
	if err != nil {
		fails = append(fails, "campfires::soul_campfire_log_lit.png failed to open!")
	} else {
		dst := imaging.New(soul_campfire_log_lit.Bounds().Dx()*2, soul_campfire_log_lit.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, soul_campfire_log_lit, image.Pt(0, 0))
		if err := imaging.Save(dst, outPath+"mcl_campfires_soul_campfire_log_lit.png"); err != nil {
			fails = append(fails, "mcl_campfires_soul_campfire_log_lit.png failed to save!")
		}
	}

	campfire_log, err := imaging.Open(inPath + "campfire_log.png")
	if err != nil {
		fails = append(fails, "campfires::campfire_fire.png failed to open!")
	} else {
		dst := imaging.New(campfire_log.Bounds().Dx()*2, campfire_log.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, campfire_log, image.Pt(0, 0))
		if err := imaging.Save(dst, outPath+"mcl_campfires_log.png"); err != nil {
			fails = append(fails, "mcl_campfires_log.png failed to save!")
		}
	}

	fire, err := imaging.Open(inPath + "campfire_fire.png")
	if err != nil {
		fails = append(fails, "campfires::campfire_fire.png failed to open!")
	} else {
		dst := imaging.New(fire.Bounds().Dx()*2, fire.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, fire, image.Pt(fire.Bounds().Dx(), 0))
		if err := imaging.Save(dst, outPath+"mcl_campfires_campfire_fire.png"); err != nil {
			fails = append(fails, "mcl_campfires_campfire_fire.png failed to save!")
		}
	}

	soulfire, err := imaging.Open(inPath + "soul_campfire_fire.png")
	if err != nil {
		fails = append(fails, "campfires::soul_campfire_fire.png failed to open!")
	} else {
		dst := imaging.New(soulfire.Bounds().Dx()*2, soulfire.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, soulfire, image.Pt(soulfire.Bounds().Dx(), 0))
		if err := imaging.Save(dst, outPath+"mcl_campfires_soul_campfire_fire.png"); err != nil {
			fails = append(fails, "mcl_campfires_soul_campfire_fire.png failed to save!")
		}
	}

	if len(fails) > 0 {
		return &readWriteError{fails, "campfire textures"}
	}
	return nil
}

func crack_fix(inPath string, outPath string) *readWriteError {
	destroy0, err := imaging.Open(inPath + "destroy_stage_0.png")
	if err != nil {
		return &readWriteError{[]string{"block::destroy_stage_0 failed to open!"}, "crack textures"}
	}
	dst := imaging.New(destroy0.Bounds().Dx(), destroy0.Bounds().Dy()*10, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, destroy0, image.Pt(0, 0))
	for i := 1; i <= 9; i++ {
		destroyPartI, err := imaging.Open(inPath + "destroy_stage_" + strconv.Itoa(i) + ".png")
		if err != nil {
			return &readWriteError{[]string{"block::destroy_stage_" + strconv.Itoa(i) + " failed to open!"}, "crack textures"}
		}
		dst = imaging.Paste(dst, destroyPartI, image.Pt(0, i*destroy0.Bounds().Dy()))
	}
	if err := imaging.Save(dst, outPath+"crack_anylength.png"); err != nil {
		return &readWriteError{[]string{"crack_anylength.png failed to save!"}, "crack textures"}
	}
	return nil
}

/*
func flip_fix(inName string, outName string) *readWriteError {
	fails := []string{}
	flips := [...][4]string{
		////mcl_bamboo
		{"block", "bamboo_door_bottom.png", "bamboo", "mcl_bamboo_door_bottom.png"},
		{"block", "bamboo_door_top.png", "bamboo", "mcl_bamboo_door_top.png"},
		////mcl_cherry_blossom
		{"block", "cherry_door_bottom.png", "cherry_blossom", "mcl_cherry_blossom_door_bottom.png"},
		{"block", "cherry_door_top.png", "cherry_blossom", "mcl_cherry_blossom_door_top.png"},
		////mcl_crimson
		{"block", "crimson_door_bottom.png", "crimson", "mcl_crimson_crimson_door_bottom.png"},
		{"block", "crimson_door_top.png", "crimson", "mcl_crimson_crimson_door_top.png"},
		{"block", "warped_door_bottom.png", "crimson", "mcl_crimson_warped_door_bottom.png"},
		{"block", "warped_door_top.png", "crimson", "mcl_crimson_warped_door_top.png"},
		////mcl_doors
		{"block", "acacia_door_bottom.png", "doors", "mcl_doors_door_acacia_lower.png"},
		{"block", "acacia_door_top.png", "doors", "mcl_doors_door_acacia_upper.png"},
		{"block", "birch_door_bottom.png", "doors", "mcl_doors_door_birch_lower.png"},
		{"block", "birch_door_top.png", "doors", "mcl_doors_door_birch_upper.png"},
		{"block", "dark_oak_door_bottom.png", "doors", "mcl_doors_door_dark_oak_lower.png"},
		{"block", "dark_oak_door_top.png", "doors", "mcl_doors_door_dark_oak_upper.png"},
		{"block", "jungle_door_bottom.png", "doors", "mcl_doors_door_jungle_lower.png"},
		{"block", "jungle_door_top.png", "doors", "mcl_doors_door_jungle_upper.png"},
		{"block", "spruce_door_bottom.png", "doors", "mcl_doors_door_spruce_lower.png"},
		{"block", "spruce_door_top.png", "doors", "mcl_doors_door_spruce_upper.png"},
		{"block", "oak_door_bottom.png", "doors", "mcl_doors_door_wood_lower.png"},
		{"block", "oak_door_top.png", "doors", "mcl_doors_door_wood_upper.png"},
		////mcl_mangrove
		{"block", "mangrove_door_bottom.png", "mangrove", "mcl_mangrove_door_bottom.png"},
		{"block", "mangrove_door_top.png", "mangrove", "mcl_mangrove_door_top.png"},
	}

	for _, e := range flips {
		img, err := imaging.Open(inName + craftPaths[e[0]] + e[1])
		if err != nil {
			fails = append(fails, e[0]+"::"+e[1]+" failed to open!")
		} else {
			img = flipH(img)
			if err = imaging.Save(img, outName+cloniaPaths[e[2]]+e[3]); err != nil {
				fails = append(fails, e[2]+"::"+e[3]+" failed to save!")
			}
		}
	}
	if len(fails) > 0 {
		return &readWriteError{fails, "flip textures"}
	} else {
		return nil
	}
}
*/

func hud_fix(inPath string, outPath string) *readWriteError {
	fails := []string{}

	make_pink := func(c color.NRGBA) color.NRGBA {
		return color.NRGBA{120, 40, 90, c.A}
	}

	func() { // health HUD
		heartLocation := inPath + craftPaths["hud"]
		heart, err := imaging.Open(heartLocation + "heart/full.png")
		if err != nil {
			fails = append(fails, "hud::heart/full.png failed to open!")
			return
		}
		heartContainer, err := imaging.Open(heartLocation + "heart/container.png")
		if err != nil {
			fails = append(fails, "hud::heart/container.png failed to open!")
			return
		}
		heartContainerRegen := imaging.AdjustFunc(heartContainer, make_pink)
		dst := imaging.Overlay(heartContainerRegen, heart, image.Pt(0, 0), 1.0)
		if err := imaging.Save(dst, outPath+cloniaPaths["potions"]+"hudbars_icon_regenerate.png"); err != nil {
			fails = append(fails, "potions::hudbars_icon_regenerate.png failed to save!")
		}

		dst = imaging.Overlay(heartContainer, heart, image.Pt(0, 0), 1.0)
		saveErr := imaging.Save(dst, outPath+cloniaPaths["hudbars"]+"hudbars_icon_health.png")
		if saveErr != nil {
			fails = append(fails, "hud::hudbars_icon_health.png failed to save!")
			return
		}
		saveErr = imaging.Save(dst, outPath+cloniaPaths["hud_base_textures"]+"heart.png")
		if saveErr != nil {
			fails = append(fails, "hud_base_textures::heart.png failed to save!")
			return
		}

		heartAbsorbing, err := imaging.Open(heartLocation + "heart/absorbing_full.png")
		if err != nil {
			fails = append(fails, "hud::heart/absorbing_full.png failed to open!")
		} else {
			dst = imaging.Overlay(heartContainer, heartAbsorbing, image.Pt(0, 0), 1.0)
			saveErr = imaging.Save(dst, outPath+cloniaPaths["potions"]+"mcl_potions_icon_absorb.png")
			if saveErr != nil {
				fails = append(fails, "potions::mcl_potions_icon_absorb.png failed to save!")
			}
		}

		heartWither, err := imaging.Open(heartLocation + "heart/withered_full.png")
		if err != nil {
			fails = append(fails, "hud::heart/withered_full.png failed to open!")
		} else {
			dst = imaging.Overlay(heartContainer, heartWither, image.Pt(0, 0), 1.0)
			saveErr = imaging.Save(dst, outPath+cloniaPaths["potions"]+"mcl_potions_icon_wither.png")
			if saveErr != nil {
				fails = append(fails, "potions::mcl_potions_icon_wither.png failed to save!")
			}
		}

		heartPoison, err := imaging.Open(heartLocation + "heart/poisoned_full.png")
		if err != nil {
			fails = append(fails, "hud::heart/poisoned_full.png failed to open!")
		} else {
			dst = imaging.Overlay(heartContainer, heartPoison, image.Pt(0, 0), 1.0)
			saveErr = imaging.Save(dst, outPath+cloniaPaths["hunger"]+"hbhunger_icon_health_poison.png")
			if saveErr != nil {
				fails = append(fails, "hunger::hbhunger_icon_health_poison.png failed to save!")
			}
		}

		heartFrozen, err := imaging.Open(heartLocation + "heart/frozen_full.png")
		if err != nil {
			fails = append(fails, "hud::heart/frozen_full.png failed to open!")
		} else {
			dst = imaging.Overlay(heartContainer, heartFrozen, image.Pt(0, 0), 1.0)
			saveErr = imaging.Save(dst, outPath+cloniaPaths["powder_snow"]+"frozen_heart.png")
			if saveErr != nil {
				fails = append(fails, "powder_snow::frozen_heart.png failed to save!")
			}
		}

		dst = imaging.Overlay(heartContainerRegen, heartPoison, image.Pt(0, 0), 1.0)
		if err := imaging.Save(dst, outPath+cloniaPaths["potions"]+"hbhunger_icon_regen_poison.png"); err != nil {
			fails = append(fails, "potions::hbhunger_icon_regen_poison.png failed to save!")
		}

		dst = imaging.Overlay(heartContainerRegen, heartWither, image.Pt(0, 0), 1.0)
		if err := imaging.Save(dst, outPath+cloniaPaths["potions"]+"mcl_potions_icon_regen_wither.png"); err != nil {
			fails = append(fails, "potions::mcl_potions_icon_regen_wither.png failed to save!")
		}

	}()
	func() { // hunger HUD
		hungerLocation := inPath + craftPaths["hud"]
		hunger, err := imaging.Open(hungerLocation + "food_full.png")
		if err != nil {
			fails = append(fails, "hud::food_full.png failed to open!")
			return
		}
		hungerContainer, err := imaging.Open(hungerLocation + "food_empty.png")
		if err != nil {
			fails = append(fails, "hud::food_empty.png failed to open!")
			return
		}
		hungerContainerPoison, err := imaging.Open(hungerLocation + "food_empty_hunger.png")
		if err != nil {
			fails = append(fails, "hud::food_empty.png failed to open!")
			return
		}
		hungerFoodPoison, err := imaging.Open(hungerLocation + "food_full_hunger.png")
		if err != nil {
			fails = append(fails, "hud::food_empty.png failed to open!")
			return
		}
		dst := imaging.Overlay(hungerContainer, hunger, image.Pt(0, 0), 1.0)
		err = imaging.Save(dst, outPath+cloniaPaths["hunger"]+"hbhunger_icon.png")
		if err != nil {
			fails = append(fails, "hud::hbhunger_icon.png failed to save!")
			return
		}
		dst = imaging.Overlay(hungerContainerPoison, hungerFoodPoison, image.Pt(0, 0), 1.0)
		err = imaging.Save(dst, outPath+cloniaPaths["hunger"]+"mcl_hunger_icon_foodpoison.png")
		if err != nil {
			fails = append(fails, "hud::hbhunger_icon.png failed to save!")
			return
		}
	}()
	if len(fails) > 0 {
		return &readWriteError{fails, "hud textures"}
	} else {
		return nil
	}
}

func mods_fixes(inPath, outPack string) *readWriteError {
	fails := []string{}
	mod := "copper_stuff"
	/*
		textures_for_copper := data.CopperStuffMod

		for _, e := range textures_for_copper {
			ironItem, err := imaging.Open(inPath + e.ReadPath())
			if err != nil {
				fails = append(fails, e.InTexture+" failed to open for mod", mod)
			} else {
				dst := imaging.New(ironItem.Bounds().Dx(), ironItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Overlay(dst, ironItem, image.Point{0, 0}, 1.0)
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
				if err = imaging.Save(dst, outPack+e.SavePath()); err != nil {
					fails = append(fails, e.OutTexture+" failed to save!")
				}
			}
		}
	*/

	func() {
		mod := "rose_gold_stuff"
		// They used diamond for the tools, netherite for the armor, and obviously, iron for the shears.

		// Netherite is hard to consistantly make look good in pink. :(
		netherite_to_rose_gold := data.RoseGoldStuffMod_NetheriteToRoseGold
		for _, e := range netherite_to_rose_gold {
			netheriteItem, err := imaging.Open(inPath + e.ReadPath())
			if err != nil {
				fails = append(fails, e.InTexture+" failed to open for mod", mod)
			} else {
				dst := imaging.New(netheriteItem.Bounds().Dx(), netheriteItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Overlay(dst, netheriteItem, image.Point{0, 0}, 1.0)
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
				if err = imaging.Save(dst, outPack+e.SavePath()); err != nil {
					fails = append(fails, e.OutTexture+" failed to save!")
				}
			}
		}

		/* Diamond stuff is not looking good :(
		diamond_to_rose_gold := [...]simpleConversion{
			{"item", "diamond_hoe.png", mod, "mcl_rose_gold_rose_gold_hoe.png", 1},
			{"item", "diamond_axe.png", mod, "mcl_rose_gold_rose_gold_axe.png", 1},
			{"item", "diamond_pickaxe.png", mod, "mcl_rose_gold_rose_gold_pick.png", 1},
			{"item", "diamond_shovel.png", mod, "mcl_rose_gold_rose_gold_shovel.png", 1},
			{"item", "diamond_sword.png", mod, "mcl_rose_gold_rose_gold_sword.png", 1},
		}
		for _, e := range diamond_to_rose_gold {
			diamondItem, err := imaging.Open(inPath + e.readPath())
			if err != nil {
				fails = append(fails, e.inTexture+" failed to open for mod", mod)
			} else {
				dst := imaging.New(diamondItem.Bounds().Dx(), diamondItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Overlay(dst, diamondItem, image.Point{0, 0}, 1.0)
				dst = imaging.AdjustFunc(dst,
					func(c color.NRGBA) color.NRGBA {

						r := int(c.R)
						g := int(c.G)
						b := int(c.B)

						if r > g || r > b {
							return c
						}

						r = ((r * 100) / 21) + 20
						g = ((g * 100) / 131) + 20
						b = ((b * 100) / 115) + 20

						if r > 255 {
							r = 255
						}
						if g > 255 {
							g = 255
						}
						if b > 255 {
							b = 255
						}

						return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
					})
				if err = imaging.Save(dst, outPack+e.savePath()); err != nil {
					fails = append(fails, e.outTexture+" failed to save!")
				}
			}
		}
		*/

		copper_to_rose_gold_exposed := data.RoseGoldStuffMod_CopperToRoseGoldExposed
		for _, e := range copper_to_rose_gold_exposed {
			copperItem, err := imaging.Open(inPath + e.ReadPath())
			if err != nil {
				fails = append(fails, e.InTexture+" failed to open for mod", mod)
			} else {
				dst := imaging.New(copperItem.Bounds().Dx(), copperItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Overlay(dst, copperItem, image.Point{0, 0}, 1.0)
				dst = imaging.AdjustFunc(dst,
					func(c color.NRGBA) color.NRGBA {

						average := (int(c.R) + int(c.G) + int(c.B)) / 3

						r := (average) + 10
						g := ((average * 100) / 110) + 10
						b := ((average * 100) / 72) + 10

						if r > 255 {
							r = 255
						}
						if b > 255 {
							b = 255
						}

						return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
					})
				if err = imaging.Save(dst, outPack+e.SavePath()); err != nil {
					fails = append(fails, e.OutTexture+" failed to save!")
				}
			}
		}

		copper_to_rose_gold := data.RoseGoldStuffMod_CopperToRoseGold
		for _, e := range copper_to_rose_gold {
			copperItem, err := imaging.Open(inPath + e.ReadPath())
			if err != nil {
				fails = append(fails, e.InTexture+" failed to open for mod", mod)
			} else {
				dst := imaging.New(copperItem.Bounds().Dx(), copperItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Overlay(dst, copperItem, image.Point{0, 0}, 1.0)
				dst = imaging.AdjustFunc(dst,
					func(c color.NRGBA) color.NRGBA {

						average := (int(c.R) + int(c.G) + int(c.B)) / 3

						r := ((average * 100) / 56) - 40
						g := ((average * 100) / 70) - 40
						b := ((average * 100) / 68) - 40

						if r > 255 {
							r = 255
						}
						if g > 255 {
							g = 255
						}
						if b > 255 {
							b = 255
						}

						return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
					})
				if err = imaging.Save(dst, outPack+e.SavePath()); err != nil {
					fails = append(fails, e.OutTexture+" failed to save!")
				}
			}
		}

		iron_to_rose_gold := data.RoseGoldStuffMod_IronToRoseGold
		for _, e := range iron_to_rose_gold {
			ironItem, err := imaging.Open(inPath + e.ReadPath())
			if err != nil {
				fails = append(fails, e.InTexture+" failed to open for mod", mod)
			} else {
				dst := imaging.New(ironItem.Bounds().Dx(), ironItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Overlay(dst, ironItem, image.Point{0, 0}, 1.0)
				dst = imaging.AdjustFunc(dst,
					func(c color.NRGBA) color.NRGBA {

						r := int(c.R)
						g := int(c.G)
						b := int(c.B)

						if (r > g+20 || r < g-20) && (r > b+20 || r < b-20) {
							return c
						}

						average := (int(c.R) + int(c.G) + int(c.B)) / 3

						r = ((average * 100) / 92)
						g = ((average * 100) / 123)
						b = ((average * 100) / 124)

						if r > 255 {
							r = 255
						}

						return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
					})
				if err = imaging.Save(dst, outPack+e.SavePath()); err != nil {
					fails = append(fails, e.OutTexture+" failed to save!")
				}
			}
		}
		iron_to_rose_gold_no_filter := data.RoseGoldStuffMod_IronToRoseGoldNoFilter
		for _, e := range iron_to_rose_gold_no_filter {
			ironItem, err := imaging.Open(inPath + e.ReadPath())
			if err != nil {
				fails = append(fails, e.InTexture+" failed to open for mod", mod)
			} else {
				dst := imaging.New(ironItem.Bounds().Dx(), ironItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Overlay(dst, ironItem, image.Point{0, 0}, 1.0)
				dst = imaging.AdjustFunc(dst,
					func(c color.NRGBA) color.NRGBA {

						r := int(c.R)
						g := int(c.G)
						b := int(c.B)

						average := (int(c.R) + int(c.G) + int(c.B)) / 3

						r = ((average * 100) / 92)
						g = ((average * 100) / 123)
						b = ((average * 100) / 124)

						if r > 255 {
							r = 255
						}

						return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
					})
				if err = imaging.Save(dst, outPack+e.SavePath()); err != nil {
					fails = append(fails, e.OutTexture+" failed to save!")
				}
			}
		}

	}() // end of "Rose Gold Stuff"

	mod = "emerald_stuff"
	emerald_stuff_textures := data.EmeraldStuffMod

	for _, e := range emerald_stuff_textures {
		diamondItem, err := imaging.Open(inPath + e.ReadPath())
		if err != nil {
			fails = append(fails, e.InTexture+" failed to open for mod", mod)
		} else {
			dst := imaging.New(diamondItem.Bounds().Dx(), diamondItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
			dst = imaging.Overlay(dst, diamondItem, image.Point{0, 0}, 1.0)
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
			if err = imaging.Save(dst, outPack+e.SavePath()); err != nil {
				fails = append(fails, e.OutTexture+" failed to save!")
			}
		}
	}

	mod = "rose_gold_stuff"

	if len(fails) > 0 {
		return &readWriteError{fails, "patched textures"}
	} else {
		return nil
	}
}

func stonecutter_fix(inPath string, outPath string) *readWriteError {
	saw, err := imaging.Open(inPath + "stonecutter_saw.png")
	if err != nil {
		return &readWriteError{[]string{"block::stonecutter_saw.png failed to open!"}, "stonecutter textures"}
	}
	if saw.Bounds().Dx()%16 != 0 {
		return &readWriteError{[]string{"block::stonecutter_saw.png has an incompatible image size!"}, "stonecutter textures"}
	}
	scale := saw.Bounds().Dx() / 16
	numOfFrames := saw.Bounds().Dy() / saw.Bounds().Dx()
	side, err := imaging.Open(inPath + "stonecutter_side.png")
	if err != nil {
		return &readWriteError{[]string{"block::stonecutter_side.png failed to open!"}, "stonecutter textures"}
	}
	dst := imaging.New(saw.Bounds().Dx(), saw.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, saw, image.Point{0, -9 * scale}, 1.0)
	for i := range numOfFrames {
		dst = imaging.Overlay(dst, side, image.Point{0, i * side.Bounds().Dx()}, 1.0)
	}
	if err := imaging.Save(dst, outPath+"mcl_stonecutter_saw.png"); err != nil {
		return &readWriteError{[]string{"mcl_stonecutter_saw.png failed to save!"}, "stonecutter textures"}
	}
	return nil
}

