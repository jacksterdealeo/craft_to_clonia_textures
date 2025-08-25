package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWDoubleChests(input_pack_path, output_pack_path string) error {
	stitch := "Chests"
	craft_path := "chest"
	clonia_path := "chests"

	in_path := input_pack_path + data.CraftPaths[craft_path]
	out_path := output_pack_path + data.CloniaPaths[clonia_path]

	readFails := make([][2]string, 0)
	saveFails := make([][2]string, 0)

equals := [...][3]string{
		{"christmas_left.png", "christmas_right.png", "mcl_chests_normal_double_present.png"},
		{"normal_left.png", "normal_right.png", "mcl_chests_normal_double.png"},
		{"trapped_left.png", "trapped_right.png", "mcl_chests_trapped_double.png"},
	}

	for _, e := range equals {
		chestLeft, err := imaging.Open(in_path + e[0])
		if err != nil {
			readFails = append(readFails, [2]string{craft_path, e[0]})
			continue
		}
		chestRight, err := imaging.Open(in_path + e[1])
		if err != nil {
			readFails = append(readFails, [2]string{craft_path, e[1]})
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
		if err = imaging.Save(dst, out_path+e[2]); err != nil {
			saveFails = append(saveFails, [2]string{clonia_path, e[0]})
			continue
		}
	}
	return multiErrMsg(stitch, readFails, saveFails)
}


func RWSingleChests(input_pack_path, output_pack_path string) error {
	stitch := "Chests"
	craft_path := "chest"
	clonia_path := "chests"

	in_path := input_pack_path + data.CraftPaths[craft_path]
	out_path := output_pack_path + data.CloniaPaths[clonia_path]

	readFails := make([][2]string, 0)
	saveFails := make([][2]string, 0)

	equals := [...][2]string{
		{"christmas.png", "mcl_chests_normal_present.png"},
		{"ender.png", "mcl_chests_ender.png"},
		{"ender.png", "mcl_chests_ender_present.png"},
		{"normal.png", "mcl_chests_normal.png"},
		{"trapped.png", "mcl_chests_trapped.png"},
	}

	for _, e := range equals {
		chestSingle, err := imaging.Open(in_path + e[0])
		if err != nil {
			readFails = append(readFails, [2]string{craft_path, e[0]})
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

		if err = imaging.Save(dst, out_path+e[1]); err != nil {
			saveFails = append(saveFails, [2]string{clonia_path, e[1]})
			continue
		}
	}
	return multiErrMsg(stitch, readFails, saveFails)
}
