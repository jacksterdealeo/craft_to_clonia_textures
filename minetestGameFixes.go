package main

import (
	"errors"
	"image"
	"image/color"

	imaging "github.com/disintegration/imaging"
)

func mtg_greenify(greenery simpleConversion, inPath, outPath string) *readWriteError { // green plants
	grayImage, err := imaging.Open(inPath + greenery.readPath())
	if err != nil {
		return &readWriteError{[]string{greenery.inTexture}, " failed to open!"}
	}
	dst := imaging.New(grayImage.Bounds().Dx(), grayImage.Bounds().Dx(), color.NRGBA{0, 0, 0, 0}) // disallow animated textures
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

func GlassCarveCenter(img image.Image) (*image.NRGBA, error) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	if width < 16 || height < 16 {
		return nil, errors.New("Image is too small.")
	}
	dst := imaging.New(width, height, color.Transparent)

	top := imaging.CropAnchor(img, width, height/16, imaging.Top)
	left := imaging.CropAnchor(img, width/16, height, imaging.Left)
	bottom := imaging.CropAnchor(img, width, height/16, imaging.Bottom)
	right := imaging.CropAnchor(img, width/16, height, imaging.Right)

	dst = imaging.Overlay(dst, top, image.Pt(0, 0), 1.0)
	dst = imaging.Overlay(dst, left, image.Pt(0, 0), 1.0)
	dst = imaging.Overlay(dst, right, image.Pt((width-(width/16)), 0), 1.0)
	dst = imaging.Overlay(dst, bottom, image.Pt(0, (height-(height/16))), 1.0)
	return dst, nil
}

// Makes obsidian glass fully transparent, as MTG doesn't like partial transparency.
func mtg_obsidian_glass_fix(inPath, outPath string) *readWriteError {
	tintedGlass := simpleConversion{"block", "obsidian.png", "mtg", "default_obsidian_glass.png", 1}
	// using mossy_cobblestone until I know the crop works correctly.
	inImage, err := imaging.Open(inPath + tintedGlass.readPath())
	if err != nil {
		return &readWriteError{[]string{tintedGlass.inTexture}, " failed to open!"}
	}

	obsidianGlass := imaging.New(inImage.Bounds().Dx(), inImage.Bounds().Dx(), color.NRGBA{0, 0, 0, 0}) // disallow animated textures
	obsidianGlass = imaging.Overlay(obsidianGlass, inImage, image.Point{0, 0}, 1.0)
	obsidianGlass, err = GlassCarveCenter(obsidianGlass)
	if err != nil {
		return &readWriteError{[]string{tintedGlass.inTexture}, " faild to carve out the center!"}
	}

	if err = imaging.Save(obsidianGlass, outPath+mtgPaths[tintedGlass.outPath]+"/"+tintedGlass.outTexture); err != nil {
		return &readWriteError{[]string{tintedGlass.inTexture}, " failed to save!"}
	}
	return nil
}
