package main

import (
	"fmt"
	"image"
	"image/color"
	"strconv"

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

func do_fixes(inPath string, outPath string) *readWriteError {
	fails := make([]string, 0, 61_000)

	func() {
		t := []simpleConversion{
			{"block", "polished_andesite.png", "core", "mcl_stairs_andesite_smooth_slab.png", 1},
			{"block", "polished_diorite.png", "core", "mcl_stairs_diorite_smooth_slab.png", 1},
			{"block", "polished_granite.png", "core", "mcl_stairs_granite_smooth_slab.png", 1},
			{"block", "gold_block.png", "xstairs", "mcl_stairs_gold_block_slab.png", 1},
			{"block", "iron_block.png", "xstairs", "mcl_stairs_iron_block_slab.png", 1},
			{"block", "lapis_block.png", "xstairs", "mcl_stairs_lapis_block_slab.png", 1},
		}
		for _, e := range t {
			block, err := imaging.Open(inPath + craftPaths[e.inPath] + e.inTexture)
			_ = block
			if err != nil {
				fails = append(fails, e.inTexture+"failed to open!")
			} else {
				scale := block.Bounds().Dx() / 16
				dst := imaging.New(block.Bounds().Dx(), block.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
				dst = imaging.Paste(dst, block, image.Pt(0, -8*scale))
				dst = imaging.Paste(dst, block, image.Pt(0, 8*scale))

				top := imaging.Crop(block, image.Rect(0, 0, block.Bounds().Dx(), 1))
				bottom := imaging.Crop(block, image.Rect(0, 15*scale, block.Bounds().Dx(), block.Bounds().Dy()))
				dst = imaging.Paste(dst, top, image.Pt(0, 0))
				dst = imaging.Paste(dst, bottom, image.Pt(0, 15*scale))

				if err := imaging.Save(dst, outPath+cloniaPaths[e.outPath]+e.outTexture); err != nil {
					fails = append(fails, e.outTexture+" failed to save!")
				}
			}
		}
	}()

	func() { // offhand_slot
		t := "hotbar_offhand_left.png"
		if offHand, err := imaging.Open(inPath + craftPaths["hud"] + t); err != nil {
			fails = append(fails, t+" failed to open!")
		} else {
			// 29 x 24
			scale := offHand.Bounds().Dx() / 29
			// 22 x 22
			dst := imaging.New(22*scale, 22*scale, color.NRGBA{0, 0, 0, 0})
			dst = imaging.Paste(dst, offHand, image.Pt(0, -1*scale))
			if err2 := imaging.Save(dst, outPath+cloniaPaths["offhand"]+"mcl_offhand_slot.png"); err2 != nil {
				fails = append(fails, t+" failed to save!")
			}
		}
	}()

	if len(fails) > 0 {
		return &readWriteError{fails, "patched textures"}
	}
	return nil
}

func anvil_fix(inPath string, outPath string) *readWriteError {
	abase, err := imaging.Open(inPath + "anvil.png")
	if err != nil {
		return &readWriteError{[]string{"block::anvil.png failed to open! Skipping the rest!"}, "anvil textures"}
	}
	a0, err := imaging.Open(inPath + "anvil_top.png")
	if err != nil {
		return &readWriteError{[]string{"block::anvil_top.png failed to open! Skipping the rest!"}, "anvil textures"}
	}
	a1, err := imaging.Open(inPath + "chipped_anvil_top.png")
	if err != nil {
		return &readWriteError{[]string{"block::chipped_anvil_top.png failed to open! Skipping the rest!"}, "anvil textures"}
	}
	a2, err := imaging.Open(inPath + "damaged_anvil_top.png")
	if err != nil {
		return &readWriteError{[]string{"block::damaged_anvil_top.png failed to open!"}, "anvil textures"}
	}
	anvilX := abase.Bounds().Dx()
	anvilY := abase.Bounds().Dy()

	dst := imaging.New(anvilX, anvilY, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, abase, image.Pt(0, 0))
	dst = imaging.OverlayCenter(dst, a0, 1.0)

	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_0.png"); err != nil {
		return &readWriteError{[]string{"mcl_anvils_anvil_top_damaged_0.png failed to save! Skipping the rest!"}, "anvil textures"}
	}
	dst = imaging.OverlayCenter(dst, a1, 1.0)
	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_1.png"); err != nil {
		return &readWriteError{[]string{"mcl_anvils_anvil_top_damaged_1.png failed to save! Skipping the rest!"}, "anvil textures"}
	}
	dst = imaging.OverlayCenter(dst, a2, 1.0)
	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_2.png"); err != nil {
		return &readWriteError{[]string{"mcl_anvils_anvil_top_damaged_2.png failed to save!"}, "anvil textures"}
	}
	return nil
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

func double_chests_fix(inPath string, outPath string) *readWriteError {
	fails := []string{}
	equals := [...][3]string{
		{"christmas_left.png", "christmas_right.png", "mcl_chests_normal_double_present.png"},
		{"normal_left.png", "normal_right.png", "mcl_chests_normal_double.png"},
		{"trapped_left.png", "trapped_right.png", "mcl_chests_trapped_double.png"},
	}

	for _, e := range equals {
		chestLeft, err := imaging.Open(inPath + e[0])
		if err != nil {
			fails = append(fails, "chest"+"::"+e[0]+" failed to open!")
			continue
		}
		chestRight, err := imaging.Open(inPath + e[1])
		if err != nil {
			fails = append(fails, "chest"+"::"+e[1]+" failed to open!")
			continue
		}
		chestX := chestLeft.Bounds().Dx()
		scale := chestX / 64
		dst := imaging.New(chestX*2, chestX, color.NRGBA{0, 0, 0, 0})
		chestF1 := flipHV(
			cropToScale(chestLeft, 14, 0, 29, 14, scale))
		chestF2 := flipV(
			cropToScale(chestLeft, 29, 0, 44, 14, scale))
		chestF3 := flipHV(
			cropToScale(chestLeft, 14, 14, 29, 19, scale))
		chestF4 := flipHV(
			cropToScale(chestLeft, 29, 14, 43, 19, scale))
		chestF5 := flipHV(
			cropToScale(chestLeft, 43, 14, 58, 19, scale))
		chestF6 := flipHV(
			cropToScale(chestLeft, 14, 19, 29, 33, scale))
		chestF7 := flipV(
			cropToScale(chestLeft, 29, 19, 44, 33, scale))
		chestF8 := flipHV(
			cropToScale(chestLeft, 14, 33, 29, 43, scale))
		chestF9 := flipHV(
			cropToScale(chestLeft, 29, 33, 43, 43, scale))
		chestF10 := flipHV(
			cropToScale(chestLeft, 43, 33, 58, 43, scale))
		chestF11 := flipHV(
			cropToScale(chestRight, 14, 0, 29, 14, scale))
		chestF12 := flipV(
			cropToScale(chestRight, 29, 0, 44, 14, scale))
		chestF13 := flipHV(
			cropToScale(chestRight, 0, 14, 14, 19, scale))
		chestF14 := flipHV(
			cropToScale(chestRight, 14, 14, 29, 19, scale))
		chestF15 := flipHV(
			cropToScale(chestRight, 43, 14, 58, 19, scale))
		chestF16 := flipHV(
			cropToScale(chestRight, 14, 19, 29, 33, scale))
		chestF17 := flipV(
			cropToScale(chestRight, 29, 19, 44, 33, scale))
		chestF18 := flipHV(
			cropToScale(chestRight, 0, 33, 14, 43, scale))
		chestF19 := flipHV(
			cropToScale(chestRight, 14, 33, 29, 43, scale))
		chestF20 := flipHV(
			cropToScale(chestRight, 43, 33, 58, 43, scale))

		dst = imaging.Overlay(dst, chestF1, image.Point{44 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF2, image.Point{29 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF4, image.Point{44 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF3, image.Point{58 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF5, image.Point{29 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF6, image.Point{44 * scale, 19 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF7, image.Point{29 * scale, 19 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF8, image.Point{58 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF9, image.Point{44 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF10, image.Point{29 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF11, image.Point{59 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF12, image.Point{14 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF13, image.Point{0 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF14, image.Point{73 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF15, image.Point{14 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF16, image.Point{59 * scale, 19 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF17, image.Point{14 * scale, 19 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF18, image.Point{0 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF19, image.Point{73 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestF20, image.Point{14 * scale, 33 * scale}, 1.0)

		chestL1 := flipV(
			imaging.Crop(chestLeft, image.Rectangle{
				image.Point{2 * scale, 0 * scale}, image.Point{3 * scale, 1 * scale}}))
		chestL2 := flipV(
			imaging.Crop(chestRight, image.Rectangle{
				image.Point{2 * scale, 0 * scale}, image.Point{3 * scale, 1 * scale}}))
		chestL3 := flipHV(
			imaging.Crop(chestLeft, image.Rectangle{
				image.Point{1 * scale, 0 * scale}, image.Point{2 * scale, 1 * scale}}))
		chestL4 := flipHV(
			imaging.Crop(chestRight, image.Rectangle{
				image.Point{1 * scale, 0 * scale}, image.Point{2 * scale, 1 * scale}}))

		dst = imaging.Overlay(dst, chestL1, image.Point{2 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestL2, image.Point{1 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestL3, image.Point{3 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestL4, image.Point{4 * scale, 0 * scale}, 1.0)

		chestLockStrip1 := cropToScale(chestLeft, 1, 1, 2, 5, scale)
		chestLockStrip1 = flipHV(chestLockStrip1)
		chestLockStrip2 := cropToScale(chestLeft, 2, 1, 3, 5, scale)
		chestLockStrip2 = flipHV(chestLockStrip2)
		chestLockStrip3 := cropToScale(chestLeft, 3, 1, 4, 5, scale)
		chestLockStrip3 = flipHV(chestLockStrip3)
		chestLockStrip4 := cropToScale(chestRight, 0, 1, 1, 5, scale)
		chestLockStrip4 = flipHV(chestLockStrip4)
		chestLockStrip5 := cropToScale(chestRight, 1, 1, 2, 5, scale)
		chestLockStrip5 = flipHV(chestLockStrip5)
		chestLockStrip6 := cropToScale(chestRight, 3, 1, 4, 5, scale)
		chestLockStrip6 = flipHV(chestLockStrip6)

		dst = imaging.Overlay(dst, chestLockStrip1, image.Point{4 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockStrip2, image.Point{3 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockStrip3, image.Point{2 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockStrip6, image.Point{1 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockStrip5, image.Point{5 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockStrip4, image.Point{0 * scale, 1 * scale}, 1.0)
		if err = imaging.Save(dst, outPath+e[2]); err != nil {
			fails = append(fails, "chest"+"::"+e[0]+" failed to save!")
			continue
		}
	}
	if len(fails) > 0 {
		return &readWriteError{fails, "double chest textures"}
	} else {
		return nil
	}
}

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

func hud_fix(inPath string, outPath string) *readWriteError {
	fails := []string{}
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
		dst := imaging.Overlay(heartContainer, heart, image.Pt(0, 0), 1.0)
		saveErr := imaging.Save(dst, outPath+cloniaPaths["hudbars"]+"hudbars_icon_health.png")
		if saveErr != nil {
			fails = append(fails, "hud::hudbars_icon_health.png failed to save!")
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
		dst := imaging.Overlay(hungerContainer, hunger, image.Pt(0, 0), 1.0)
		saveErr := imaging.Save(dst, outPath+cloniaPaths["hunger"]+"hbhunger_icon.png")
		if saveErr != nil {
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

// Restitches the extremely cursed MC version.
// TODO mcl_flowerpots_cactus.png
func flowerpot_fix(inPath string, outPath string) *readWriteError {
	pot, err := imaging.Open(inPath + "flower_pot.png")
	if err != nil {
		return &readWriteError{[]string{"block::flower_pot.png failed to open!"}, "flower pot textures"}
	}
	dirt, err := imaging.Open(inPath + "dirt.png")
	if err != nil {
		return &readWriteError{[]string{"block::dirt.png failed to open!"}, "flower pot textures"}
	}

	potX := pot.Bounds().Dx()
	potY := pot.Bounds().Dy()
	scale := potX / 16

	dirt = imaging.CropCenter(dirt, 4*scale, 4*scale)

	dst := imaging.New((potX)*2, (potY)*2, color.NRGBA{0, 0, 0, 0})
	fPotTop := imaging.Crop(pot, image.Rectangle{
		image.Point{5 * scale, 5 * scale}, image.Point{11 * scale, 11 * scale}})
	fPotSide := imaging.Crop(pot, image.Rectangle{
		image.Point{5 * scale, 10 * scale}, image.Point{11 * scale, 16 * scale}})
	fPotInner := imaging.Crop(pot, image.Rectangle{
		image.Point{5 * scale, 10 * scale}, image.Point{9 * scale, 12 * scale}})

	// pot sides
	dst = imaging.Overlay(dst, fPotSide, image.Point{0 * scale, 20 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotSide, image.Point{6 * scale, 20 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotSide, image.Point{12 * scale, 20 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotSide, image.Point{18 * scale, 20 * scale}, 1.0)
	// pot bottom
	dst = imaging.Overlay(dst, fPotSide, image.Point{22 * scale, 26 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotTop, image.Point{22 * scale, 26 * scale}, 1.0)
	// pot top
	dst = imaging.Overlay(dst, fPotTop, image.Point{16 * scale, 26 * scale}, 1.0)
	// pot inside
	dst = imaging.Overlay(dst, fPotInner, image.Point{0 * scale, 30 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotInner, image.Point{4 * scale, 30 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotInner, image.Point{8 * scale, 30 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotInner, image.Point{12 * scale, 30 * scale}, 1.0)
	// dirt
	dst = imaging.Overlay(dst, dirt, image.Point{4 * scale, 26 * scale}, 1.0)

	if err = imaging.Save(dst, outPath+"mcl_flowerpots_flowerpot.png"); err != nil {
		return &readWriteError{[]string{"mcl_flowerpots_flowerpot.png failed to save!"}, "flower pot textures"}
	}
	return nil
}

func lava_fix(inPath string, outPath string) *readWriteError {
	/*
		craft lava
		  still   :  16 x 512
		  flowing :  32 x 1024
		clonia lava
			still   :  16 x 256
		  flowing :  16 x 1024
	*/
	lavaFlowing, err := imaging.Open(inPath + "lava_flow.png")
	if err != nil {
		return &readWriteError{[]string{"lava_flow.png failed to open!"}, "lava textures"}
	} else {
		lavaStillX := lavaFlowing.Bounds().Dx()
		lavaStillY := lavaFlowing.Bounds().Dy()
		dst := imaging.New(lavaStillX/2, lavaStillY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Overlay(dst, lavaFlowing, image.Point{0, 0}, 1.0)
		if err = imaging.Save(dst, outPath+"default_lava_flowing_animated.png"); err != nil {
			return &readWriteError{[]string{"default_lava_flowing_animated.png failed to save!"}, "lava textures"}
		}
	}
	if err := copyTextureAnimated(inPath+"lava_still.png", outPath+"default_lava_source_animated.png", -1); err != nil {
		return &readWriteError{[]string{"default_lava_source_animated.png failed to copy!"}, "lava textures"}
	}
	return nil
}

func lily_fix(inPath string, outPath string) *readWriteError {
	grayLily, err := imaging.Open(inPath + "lily_pad.png")
	if err != nil {
		return &readWriteError{[]string{"block::flowers_waterlily.png failed to open!"}, "vine texture"}
	}
	_ = grayLily
	dst := imaging.New(grayLily.Bounds().Dx(), grayLily.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, grayLily, image.Point{0, 0}, 1.0)
	dst = imaging.AdjustFunc(dst,
		func(c color.NRGBA) color.NRGBA {
			r := int(c.R) - 255
			g := int(c.G) - 20
			b := int(c.B) - 255
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
	if err = imaging.Save(dst, outPath+"flowers_waterlily.png"); err != nil {
		return &readWriteError{[]string{"flowers_waterlily.png failed to save!"}, "waterlily texture"}
	}
	return nil
}

func single_chests_fix(inPath string, outPath string) *readWriteError {
	fails := []string{}
	equals := [...][2]string{
		{"christmas.png", "mcl_chests_normal_present.png"},
		{"ender.png", "mcl_chests_ender.png"},
		{"ender.png", "mcl_chests_ender_present.png"},
		{"normal.png", "mcl_chests_normal.png"},
		{"trapped.png", "mcl_chests_trapped.png"},
	}
	for _, e := range equals {
		chestSingle, err := imaging.Open(inPath + e[0])
		if err != nil {
			fails = append(fails, "chests"+"::"+e[0]+" failed to open!")
			continue
		}

		chestX := chestSingle.Bounds().Dx()
		scale := chestX / 64
		dst := imaging.New(chestX, chestX, color.NRGBA{0, 0, 0, 0})

		chestTopTop := flipV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{28 * scale, 0 * scale}, image.Point{42 * scale, 14 * scale}}))
		chestTopBot := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{14 * scale, 0 * scale}, image.Point{28 * scale, 14 * scale}}))
		chestTopFace := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{42 * scale, 14 * scale}, image.Point{56 * scale, 19 * scale}}))
		chestTopBack := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{14 * scale, 14 * scale}, image.Point{28 * scale, 19 * scale}}))
		chestTopLeft := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{0 * scale, 14 * scale}, image.Point{14 * scale, 19 * scale}}))
		chestTopRight := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{28 * scale, 14 * scale}, image.Point{42 * scale, 19 * scale}}))

		dst = imaging.Overlay(dst, chestTopTop, image.Point{14 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopBot, image.Point{28 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopFace, image.Point{14 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopBack, image.Point{42 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopLeft, image.Point{0 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopRight, image.Point{28 * scale, 14 * scale}, 1.0)

		chestBotTop := flipV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{28 * scale, 19 * scale}, image.Point{42 * scale, 33 * scale}}))
		chestBotBot := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{14 * scale, 19 * scale}, image.Point{28 * scale, 33 * scale}}))
		chestBotFace := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{42 * scale, 33 * scale}, image.Point{56 * scale, 43 * scale}}))
		chestBotBack := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{14 * scale, 33 * scale}, image.Point{28 * scale, 43 * scale}}))
		chestBotLeft := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{0 * scale, 33 * scale}, image.Point{14 * scale, 43 * scale}}))
		chestBotRight := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{28 * scale, 33 * scale}, image.Point{42 * scale, 43 * scale}}))

		dst = imaging.Overlay(dst, chestBotTop, image.Point{14 * scale, 19 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotBot, image.Point{28 * scale, 19 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotFace, image.Point{14 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotBack, image.Point{42 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotLeft, image.Point{0 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotRight, image.Point{28 * scale, 33 * scale}, 1.0)

		chestLockTop := flipV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{3 * scale, 0 * scale}, image.Point{5 * scale, 1 * scale}}))
		chestLockBot := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{1 * scale, 0 * scale}, image.Point{3 * scale, 1 * scale}})
		chestLockFace := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{4 * scale, 1 * scale}, image.Point{6 * scale, 5 * scale}}))
		chestLockBack := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{1 * scale, 1 * scale}, image.Point{3 * scale, 5 * scale}}))

		chestLockLeft := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{0 * scale, 1 * scale}, image.Point{1 * scale, 5 * scale}}))

		chestLockRight := flipHV(
			imaging.Crop(chestSingle, image.Rectangle{
				image.Point{3 * scale, 1 * scale}, image.Point{4 * scale, 5 * scale}}))

		dst = imaging.Overlay(dst, chestLockTop, image.Point{1 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockBot, image.Point{3 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockFace, image.Point{1 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockRight, image.Point{3 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockBack, image.Point{4 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockLeft, image.Point{0 * scale, 1 * scale}, 1.0)

		if err = imaging.Save(dst, outPath+e[1]); err != nil {
			fails = append(fails, "chest::"+e[1]+" failed to save!")
			continue
		}
	}
	if len(fails) > 0 {
		return &readWriteError{fails, "single chest textures"}
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
	for i := 0; i < numOfFrames; i++ {
		dst = imaging.Overlay(dst, side, image.Point{0, i * side.Bounds().Dx()}, 1.0)
	}
	if err := imaging.Save(dst, outPath+"mcl_stonecutter_saw.png"); err != nil {
		return &readWriteError{[]string{"mcl_stonecutter_saw.png failed to save!"}, "stonecutter textures"}
	}
	return nil
}

func vine_fix(inPath string, outPath string) *readWriteError {
	grayVine, err := imaging.Open(inPath + "vine.png")
	if err != nil {
		return &readWriteError{[]string{"block::vine.png failed to open!"}, "vine texture"}
	}
	_ = grayVine
	dst := imaging.New(grayVine.Bounds().Dx(), grayVine.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, grayVine, image.Point{0, 0}, 1.0)
	dst = imaging.AdjustFunc(dst,
		func(c color.NRGBA) color.NRGBA {
			r := int(c.G) - 60
			g := int(c.G)
			b := int(c.G) - 120
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
	if err = imaging.Save(dst, outPath+"mcl_core_vine.png"); err != nil {
		return &readWriteError{[]string{"mcl_core_vine.png failed to save!"}, "vine texture"}
	}

	return nil
}

func water_fix(inPath string, outPath string) *readWriteError {
	/*
		craft water
		  still   :  16 x 512
		  flowing :  32 x 1024
		clonia water (and river water)
		  still   :  16 x 256
		  flowing :  16 x 1024
	*/
	wStill, err := imaging.Open(inPath + "water_still.png")
	if err != nil {
		return &readWriteError{[]string{"block::water_still.png failed to open!"}, "water textures"}
	}
	wStillX := wStill.Bounds().Dx()
	wStillY := wStill.Bounds().Dy()
	dst := imaging.New(wStillX, wStillY, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, wStill, image.Point{0, 0}, 1.0)
	plainWater := imaging.AdjustFunc(dst,
		func(c color.NRGBA) color.NRGBA {
			r := int(c.R) - 105
			g := int(c.G) - 40
			b := int(c.B) + 20
			if r < 0 {
				r = 0
			}
			if g < 0 {
				g = 0
			}
			if b > 255 {
				b = 255
			}
			return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
		})
	if err = imaging.Save(plainWater, outPath+"default_water_source_animated.png"); err != nil {
		return &readWriteError{[]string{"default_water_source_animated.png failed to save!"}, "water textures"}
	}

	riverWater := imaging.AdjustFunc(dst,
		func(c color.NRGBA) color.NRGBA {
			r := int(c.R) - 105
			g := int(c.G) - 0
			b := int(c.B) + 45
			if r < 0 {
				r = 0
			}
			if g < 0 {
				g = 0
			}
			if b > 255 {
				b = 255
			}
			return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
		})
	if err = imaging.Save(riverWater, outPath+"default_river_water_source_animated.png"); err != nil {
		return &readWriteError{[]string{"default_river_water_source_animated.png failed to save!"}, "water textures"}
	}

	wFlowing, err := imaging.Open(inPath + "water_flow.png")
	if err != nil {
		return &readWriteError{[]string{"block::water_flow.png failed to open!"}, "water textures"}
	} else {
		wFlowingX := wFlowing.Bounds().Dx()
		wFlowingY := wFlowing.Bounds().Dy()
		dst := imaging.New(wFlowingX/2, wFlowingY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Overlay(dst, wFlowing, image.Point{0, 0}, 1.0)
		plainWater := imaging.AdjustFunc(dst,
			func(c color.NRGBA) color.NRGBA {
				r := int(c.R) - 105
				g := int(c.G) - 40
				b := int(c.B) + 20
				if r < 0 {
					r = 0
				}
				if g < 0 {
					g = 0
				}
				if b > 255 {
					b = 255
				}
				return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
			})
		if err = imaging.Save(plainWater, outPath+"default_water_flowing_animated.png"); err != nil {
			return &readWriteError{[]string{"default_water_flowing_animated.png failed to save!"}, "water textures"}
		}

		riverWater := imaging.AdjustFunc(dst,
			func(c color.NRGBA) color.NRGBA {
				r := int(c.R) - 105
				g := int(c.G) - 0
				b := int(c.B) + 45
				if r < 0 {
					r = 0
				}
				if g < 0 {
					g = 0
				}
				if b > 255 {
					b = 255
				}
				return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
			})
		if err = imaging.Save(riverWater, outPath+"default_river_water_flowing_animated.png"); err != nil {
			fmt.Println("default_river_water_flowing_animated.png save failed!")
			return &readWriteError{[]string{"default_river_water_flowing_animated.png failed to save!"}, "water textures"}
		}
	}
	return nil
}
