package main

import (
	"image"
	"image/color"

	imaging "github.com/disintegration/imaging"
)

func mtg_greenify(greenery simpleConversion, inPath, outPath string) *readWriteError { // green plants
	grayImage, err := imaging.Open(inPath + greenery.readPath())
	if err != nil {
		return &readWriteError{[]string{greenery.inTexture}, " failed to open!"}
	}
	dst := imaging.New(grayImage.Bounds().Dx(), grayImage.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, grayImage, image.Point{0, 0}, 1.0)
	dst = imaging.AdjustFunc(dst,
		func(c color.NRGBA) color.NRGBA {
			r := int(c.R) / 3
			g := int(c.G) - ((int(c.G) * 2) / 7)
			b := int(c.B) / 3
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
	if err = imaging.Save(dst, outPath+mtgPaths[greenery.outPath]+"/"+greenery.outTexture); err != nil {
		return &readWriteError{[]string{greenery.inTexture}, " failed to save!"}
	}
	return nil
}
