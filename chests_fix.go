package main

import (
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
)

func chests_fix(inPath string, outPath string) error {
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
