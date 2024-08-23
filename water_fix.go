package main

import (
	"fmt"
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
	_ "image/png"
)

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
