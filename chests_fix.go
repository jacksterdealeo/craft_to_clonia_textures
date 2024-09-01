package main

/*
import (
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
)

func chests_fix(inPath string, outPath string) error {
	chestdoubleleft, err := imaging.Open(inPath + "normal_left.png")
	if err != nil {
		return err
	}
	chestdoubleright, err := imaging.Open(inPath + "normal_right.png")
	if err != nil {
		return err
	}
	chestdoublelefttrap, err := imaging.Open(inPath + "trapped_left.png")
	if err != nil {
		return err
	}
	chestdoublerighttrap, err := imaging.Open(inPath + "trapped_right.png")
	if err != nil {
		return err
	}

	chestX := chestdoubleleft.Bounds().Dx()
	scale := chestX / 64

	dst := imaging.New(chestX, chestX*2, color.NRGBA{0, 0, 0, 0})
	fPotTop := imaging.Crop(pot, image.Rectangle{
		image.Point{5 * scale, 5 * scale}, image.Point{11 * scale, 11 * scale}})
	dst = imaging.Overlay(dst, dirt, image.Point{4 * scale, 26 * scale}, 1.0)

	if err = imaging.Save(dst, outPath+"mcl_chests_normal_double.png"); err != nil {
		return err
	}
	if err = imaging.Save(dst, outPath+"mcl_chests_trapped_double.png"); err != nil {
		return err
	}
	if err = imaging.Save(dst, outPath+"mcl_chests_normal_double_present.png"); err != nil {
		return err
	}
	return nil
}
*/
