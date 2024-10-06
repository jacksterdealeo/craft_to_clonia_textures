package main

import (
	"fmt"
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
)

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
