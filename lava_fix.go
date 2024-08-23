package main

import (
	"fmt"
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
	_ "image/png"
)

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
