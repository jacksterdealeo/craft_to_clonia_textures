package main

import (
	"fmt"
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
)

func animated_texture_fix(inName string, outName string) error {

	craftPaths := CraftPaths()
	cloniaPaths := CloniaPaths()

	animated := [][4]string{
		{"block", "respawn_anchor_top.png", "beds", "respawn_anchor_top_on.png"},
		{"block", "soul_fire_0.png", "blackstone", "soul_fire_basic_flame_animated.png"},
		{"block", "fire_0.png", "fire", "fire_basic_flame_animated.png"},
		{"block", "fire_0.png", "fire", "mcl_burning_entity_flame_animated.png"},
		{"block", "fire_0.png", "fire", "mcl_burning_hud_flame_animated.png"},
		{"block", "magma.png", "nether", "mcl_nether_magma.png"},
	}

	for _, e := range animated {
		img, err := imaging.Open(inName + craftPaths[e[0]] + e[1])
		if err != nil {
			fmt.Println(e[1], "couldn't open!")
		} else {
			if err = imaging.Save(img, outName+cloniaPaths[e[2]]+e[3]); err != nil {
				fmt.Println(e[3], "failed to save!")
			}
		}
	}
	return nil
}

func anvil_fix(inPath string, outPath string) error {
	abase, err := imaging.Open(inPath + "anvil.png")
	if err != nil {
		fmt.Println("AnvilBase error ~", "block, anvil.png")
		return err
	}
	a0, err := imaging.Open(inPath + "anvil_top.png")
	if err != nil {
		fmt.Println("Anvil0 error ~")
		return err
	}
	a1, err := imaging.Open(inPath + "chipped_anvil_top.png")
	if err != nil {
		fmt.Println("Anvil1 error ~")
		return err
	}
	a2, err := imaging.Open(inPath + "damaged_anvil_top.png")
	if err != nil {
		fmt.Println("Anvil2 error ~")
		return err
	}
	anvilX := abase.Bounds().Dx()
	anvilY := abase.Bounds().Dy()

	dst := imaging.New(anvilX, anvilY, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, abase, image.Pt(0, 0))
	dst = imaging.OverlayCenter(dst, a0, 1.0)

	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_0.png"); err != nil {
		fmt.Println("Anvil undamaged failed!")
		return err
	}
	dst = imaging.OverlayCenter(dst, a1, 1.0)
	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_1.png"); err != nil {
		fmt.Println("Anvil damaged1 failed!")
		return err
	}
	dst = imaging.OverlayCenter(dst, a2, 1.0)
	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_2.png"); err != nil {
		fmt.Println("Anvil damaged2 failed!")
		return err
	}
	return nil
}

func chests_fix(inPath string, outPath string) error {
	double_chests_fix(inPath, outPath)
	single_chests_fix(inPath, outPath)
	return nil
}

func double_chests_fix(inPath string, outPath string) error {

	flipHV := func(img image.Image) *image.NRGBA {
		return imaging.FlipH(imaging.FlipV(img))
	}

	cropToScale := func(img image.Image, x1, y1, x2, y2, scale int) *image.NRGBA {
		return imaging.Crop(img, image.Rectangle{
			image.Point{x1 * scale, y1 * scale}, image.Point{x2 * scale, y2 * scale}})
	}

	equals := [][3]string{
		{"christmas_left.png", "christmas_right.png", "mcl_chests_normal_double_present.png"},
		{"normal_left.png", "normal_right.png", "mcl_chests_normal_double.png"},
		{"trapped_left.png", "trapped_right.png", "mcl_chests_trapped_double.png"},
	}

	for _, e := range equals {
		chestLeft, err := imaging.Open(inPath + e[0])
		if err != nil {
			fmt.Println("chest::" + inPath + e[0])
			continue
			//return err
		}
		chestRight, err := imaging.Open(inPath + e[1])
		if err != nil {
			fmt.Println("chest::" + inPath + e[1])
			continue
			//return err
		}
		chestX := chestLeft.Bounds().Dx()
		scale := chestX / 64
		dst := imaging.New(chestX*2, chestX, color.NRGBA{0, 0, 0, 0})
		//NOT DONE
		chestF1 := cropToScale(chestLeft, 14, 0, 29, 14, scale)
		chestF1 = flipHV(chestF1)
		chestF2 := cropToScale(chestLeft, 29, 0, 44, 14, scale)
		chestF2 = imaging.FlipV(chestF2)
		chestF3 := cropToScale(chestLeft, 14, 14, 29, 19, scale)
		chestF3 = flipHV(chestF3)
		chestF4 := cropToScale(chestLeft, 29, 14, 43, 19, scale)
		chestF4 = flipHV(chestF4)
		chestF5 := cropToScale(chestLeft, 43, 14, 58, 19, scale)
		chestF5 = flipHV(chestF5)
		chestF6 := cropToScale(chestLeft, 14, 19, 29, 33, scale)
		chestF6 = flipHV(chestF6)
		chestF7 := cropToScale(chestLeft, 29, 19, 44, 33, scale)
		chestF7 = imaging.FlipV(chestF7)
		chestF8 := cropToScale(chestLeft, 14, 33, 29, 43, scale)
		chestF8 = flipHV(chestF8)
		chestF9 := cropToScale(chestLeft, 29, 33, 43, 43, scale)
		chestF9 = flipHV(chestF9)
		chestF10 := cropToScale(chestLeft, 43, 33, 58, 43, scale)
		chestF10 = flipHV(chestF10)
		chestF11 := cropToScale(chestRight, 14, 0, 29, 14, scale)
		chestF11 = flipHV(chestF11)
		chestF12 := cropToScale(chestRight, 29, 0, 44, 14, scale)
		chestF12 = imaging.FlipV(chestF12)
		chestF13 := cropToScale(chestRight, 0, 14, 14, 19, scale)
		chestF13 = flipHV(chestF13)
		chestF14 := cropToScale(chestRight, 14, 14, 29, 19, scale)
		chestF14 = flipHV(chestF14)
		chestF15 := cropToScale(chestRight, 43, 14, 58, 19, scale)
		chestF15 = flipHV(chestF15)
		chestF16 := cropToScale(chestRight, 14, 19, 29, 33, scale)
		chestF16 = flipHV(chestF16)
		chestF17 := cropToScale(chestRight, 29, 19, 44, 33, scale)
		chestF17 = imaging.FlipV(chestF17)
		chestF18 := cropToScale(chestRight, 0, 33, 14, 43, scale)
		chestF18 = flipHV(chestF18)
		chestF19 := cropToScale(chestRight, 14, 33, 29, 43, scale)
		chestF19 = flipHV(chestF19)
		chestF20 := cropToScale(chestRight, 43, 33, 58, 43, scale)
		chestF20 = flipHV(chestF20)

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

		chestL1 := imaging.Crop(chestLeft, image.Rectangle{
			image.Point{2 * scale, 0 * scale}, image.Point{3 * scale, 1 * scale}})
		chestL1 = imaging.FlipV(chestL1)
		chestL2 := imaging.Crop(chestRight, image.Rectangle{
			image.Point{2 * scale, 0 * scale}, image.Point{3 * scale, 1 * scale}})
		chestL2 = imaging.FlipV(chestL2)
		chestL3 := imaging.Crop(chestLeft, image.Rectangle{
			image.Point{1 * scale, 0 * scale}, image.Point{2 * scale, 1 * scale}})
		chestL3 = flipHV(chestL3)
		chestL4 := imaging.Crop(chestRight, image.Rectangle{
			image.Point{1 * scale, 0 * scale}, image.Point{2 * scale, 1 * scale}})
		chestL4 = flipHV(chestL4)

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
			continue
			//return err
		}
	}
	return nil
}

func flip_fix(inName string, outName string) error {

	craftPaths := CraftPaths()
	cloniaPaths := CloniaPaths()

	flips := [][4]string{
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
			fmt.Println(e[1], "couldn't open!")
		} else {
			img = imaging.FlipH(img)
			//anvilX := abase.Bounds().Dx()
			//anvilY := abase.Bounds().Dy()

			if err = imaging.Save(img, outName+cloniaPaths[e[2]]+e[3]); err != nil {
				fmt.Println(e[3], "failed to save!")
			}
		}
	}
	return nil
}

// Restitches the extremely cursed MC version.
// TODO mcl_flowerpots_cactus.png
func flowerpot_fix(inPath string, outPath string) error {
	pot, err := imaging.Open(inPath + "flower_pot.png")
	if err != nil {
		return err
	}
	dirt, err := imaging.Open(inPath + "dirt.png")
	if err != nil {
		return err
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
		return err
	}
	return nil
}

func lava_fix(inPath string, outPath string) error {
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
		fmt.Println("Error ~", inPath+"lava_flow.png")
		return err
	} else {
		lavaStillX := lavaFlowing.Bounds().Dx()
		lavaStillY := lavaFlowing.Bounds().Dy()
		dst := imaging.New(lavaStillX/2, lavaStillY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Overlay(dst, lavaFlowing, image.Point{0, 0}, 1.0)
		if err = imaging.Save(dst, outPath+"default_lava_flowing_animated.png"); err != nil {
			fmt.Println("default_lava_flowing_animated.png", "save failed!")
			return err
		}
	}
	copyTextureAnimated(inPath+"lava_still.png", outPath+"default_lava_source_animated.png", -1)
	return nil
}

func single_chests_fix(inPath string, outPath string) error {
	equals := [][2]string{
		{"christmas.png", "mcl_chests_normal_present.png"},
		{"ender.png", "mcl_chests_ender.png"},
		{"ender.png", "mcl_chests_ender_present.png"},
		{"normal.png", "mcl_chests_normal.png"},
		{"trapped.png", "mcl_chests_trapped.png"},
	}
	for _, e := range equals {

		chestSingle, err := imaging.Open(inPath + e[0])
		if err != nil {
			continue
			//return err
		}

		chestX := chestSingle.Bounds().Dx()
		scale := chestX / 64
		dst := imaging.New(chestX, chestX, color.NRGBA{0, 0, 0, 0})

		chestTopTop := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{28 * scale, 0 * scale}, image.Point{42 * scale, 14 * scale}})
		chestTopTop = imaging.FlipV(chestTopTop)
		chestTopBot := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{14 * scale, 0 * scale}, image.Point{28 * scale, 14 * scale}})
		chestTopBot = imaging.FlipV(imaging.FlipH(chestTopBot))
		chestTopFace := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{42 * scale, 14 * scale}, image.Point{56 * scale, 19 * scale}})
		chestTopFace = imaging.FlipV(imaging.FlipH(chestTopFace))
		chestTopBack := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{14 * scale, 14 * scale}, image.Point{28 * scale, 19 * scale}})
		chestTopBack = imaging.FlipV(imaging.FlipH(chestTopBack))
		chestTopLeft := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{0 * scale, 14 * scale}, image.Point{14 * scale, 19 * scale}})
		chestTopLeft = imaging.FlipV(imaging.FlipH(chestTopLeft))
		chestTopRight := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{28 * scale, 14 * scale}, image.Point{42 * scale, 19 * scale}})
		chestTopRight = imaging.FlipV(imaging.FlipH(chestTopRight))

		dst = imaging.Overlay(dst, chestTopTop, image.Point{14 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopBot, image.Point{28 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopFace, image.Point{14 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopBack, image.Point{42 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopLeft, image.Point{0 * scale, 14 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestTopRight, image.Point{28 * scale, 14 * scale}, 1.0)

		chestBotTop := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{28 * scale, 19 * scale}, image.Point{42 * scale, 33 * scale}})
		chestBotTop = imaging.FlipV(chestBotTop)
		chestBotBot := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{14 * scale, 19 * scale}, image.Point{28 * scale, 33 * scale}})
		chestBotBot = imaging.FlipV(chestBotBot)
		chestBotBot = imaging.FlipH(chestBotBot)
		chestBotFace := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{42 * scale, 33 * scale}, image.Point{56 * scale, 43 * scale}})
		chestBotFace = imaging.FlipV(imaging.FlipH(chestBotFace))
		chestBotBack := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{14 * scale, 33 * scale}, image.Point{28 * scale, 43 * scale}})
		chestBotBack = imaging.FlipV(imaging.FlipH(chestBotBack))
		chestBotLeft := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{0 * scale, 33 * scale}, image.Point{14 * scale, 43 * scale}})
		chestBotLeft = imaging.FlipV(imaging.FlipH(chestBotLeft))
		chestBotRight := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{28 * scale, 33 * scale}, image.Point{42 * scale, 43 * scale}})
		chestBotRight = imaging.FlipV(imaging.FlipH(chestBotRight))

		dst = imaging.Overlay(dst, chestBotTop, image.Point{14 * scale, 19 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotBot, image.Point{28 * scale, 19 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotFace, image.Point{14 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotBack, image.Point{42 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotLeft, image.Point{0 * scale, 33 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestBotRight, image.Point{28 * scale, 33 * scale}, 1.0)

		chestLockTop := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{3 * scale, 0 * scale}, image.Point{5 * scale, 1 * scale}})
		chestLockTop = (imaging.FlipV(chestLockTop))
		chestLockBot := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{1 * scale, 0 * scale}, image.Point{3 * scale, 1 * scale}})
		chestLockFace := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{4 * scale, 1 * scale}, image.Point{6 * scale, 5 * scale}})
		chestLockFace = imaging.FlipV(imaging.FlipH(chestLockFace))
		chestLockBack := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{1 * scale, 1 * scale}, image.Point{3 * scale, 5 * scale}})
		chestLockBack = imaging.FlipV(imaging.FlipH(chestLockBack))

		chestLockLeft := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{0 * scale, 1 * scale}, image.Point{1 * scale, 5 * scale}})
		chestLockLeft = imaging.FlipV(imaging.FlipH(chestLockLeft))

		chestLockRight := imaging.Crop(chestSingle, image.Rectangle{
			image.Point{3 * scale, 1 * scale}, image.Point{4 * scale, 5 * scale}})
		chestLockRight = imaging.FlipV(imaging.FlipH(chestLockRight))

		dst = imaging.Overlay(dst, chestLockTop, image.Point{1 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockBot, image.Point{3 * scale, 0 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockFace, image.Point{1 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockRight, image.Point{3 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockBack, image.Point{4 * scale, 1 * scale}, 1.0)
		dst = imaging.Overlay(dst, chestLockLeft, image.Point{0 * scale, 1 * scale}, 1.0)

		if err = imaging.Save(dst, outPath+e[1]); err != nil {
			continue
			//return err
		}
	}
	return nil
}

func water_fix(inPath string, outPath string) error {
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
		fmt.Println("water_still.png error ~", inPath+"water.png")
	} else {
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
			fmt.Println("default_water_source_animated.png save failed!")
			return err
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
			fmt.Println("default_river_water_source_animated.png save failed!")
			return err
		}
	}

	wFlowing, err := imaging.Open(inPath + "water_flow.png")
	if err != nil {
		return err
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
			fmt.Println("FlowingWater error~", inPath+"default_water_flowing_animated.png")
			return err
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
			return err
		}
	}
	return nil
}
