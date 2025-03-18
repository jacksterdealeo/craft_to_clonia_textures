package main

import (
	"errors"
	"image"
	"image/color"

	imaging "github.com/disintegration/imaging"
)

func mtg_green_it(img image.Image) (result *image.NRGBA) {
	result = imaging.AdjustFunc(img,
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
	return
}

func mtg_greenify(greenery simpleConversion, inPath, outPath string) *readWriteError { // green plants
	grayImage, err := imaging.Open(inPath + greenery.readPath())
	if err != nil {
		return &readWriteError{[]string{greenery.inTexture}, " failed to open!"}
	}
	dst := imaging.New(grayImage.Bounds().Dx(), grayImage.Bounds().Dx(), color.NRGBA{0, 0, 0, 0}) // disallows animated textures
	dst = imaging.Overlay(dst, grayImage, image.Point{0, 0}, 1.0)
	dst = mtg_green_it(dst)
	if err = imaging.Save(dst, outPath+mtgPaths[greenery.outPath]+"/"+greenery.outTexture); err != nil {
		return &readWriteError{[]string{greenery.inTexture}, " failed to save!"}
	}
	return nil
}

func mtg_grass_fix(inPath, outPath string) *readWriteError { // work on this first!
	shortGrass, err := imaging.Open(inPath + craftPaths["block"] + "short_grass.png")
	if err != nil {
		return &readWriteError{[]string{"block::short_grass.png"}, " failed to open!"}
	}
	width := shortGrass.Bounds().Dx()
	scale := width / 16
	shortGrass = mtg_green_it(shortGrass)
	dst5 := imaging.New(width, width, color.NRGBA{0, 0, 0, 0}) // disallows animated textures
	dst5 = imaging.Overlay(shortGrass, shortGrass, image.Pt(0, 0), 1.0)
	//dst5 = imaging.Overlay(shortGrass)

	// dst4 := imaging.New(shortGrass.Bounds().Dx(), shortGrass.Bounds().Dx(), color.NRGBA{0, 0, 0, 0}) // disallows animated textures
	if err = imaging.Save(dst5, outPath+mtgPaths["mtg"]+"/default_grass_5.png"); err != nil {
		return &readWriteError{[]string{"default_grass_5.png"}, " failed to save!"}
	}

	dst4 := imaging.New(width, width, color.NRGBA{0, 0, 0, 0})
	dst4 = imaging.Overlay(
		dst4,
		imaging.Crop(shortGrass, image.Rect(0, 5*scale, width, width)),
		image.Pt(0, 6*scale),
		1.0,
	)
	if err = imaging.Save(dst4, outPath+mtgPaths["mtg"]+"/default_grass_4.png"); err != nil {
		return &readWriteError{[]string{"default_grass_4.png"}, " failed to save!"}
	}

	dst3 := imaging.Overlay(
		imaging.New(width, width, color.NRGBA{0, 0, 0, 0}),
		dst4,
		image.Pt(0, 3*scale),
		1.0,
	)
	if err = imaging.Save(dst3, outPath+mtgPaths["mtg"]+"/default_grass_3.png"); err != nil {
		return &readWriteError{[]string{"default_grass_3.png"}, " failed to save!"}
	}

	dst2 := imaging.Overlay(
		imaging.New(width, width, color.NRGBA{0, 0, 0, 0}),
		dst3,
		image.Pt(0, 2*scale),
		1.0,
	)
	if err = imaging.Save(dst2, outPath+mtgPaths["mtg"]+"/default_grass_2.png"); err != nil {
		return &readWriteError{[]string{"default_grass_2.png"}, " failed to save!"}
	}

	dst1 := imaging.Overlay(
		imaging.New(width, width, color.NRGBA{0, 0, 0, 0}),
		dst2,
		image.Pt(0, 2*scale),
		1.0,
	)
	if err = imaging.Save(dst1, outPath+mtgPaths["mtg"]+"/default_grass_1.png"); err != nil {
		return &readWriteError{[]string{"default_grass_1.png"}, " failed to save!"}
	}
	return nil
}

func GlassCarveCenter(img image.Image) (*image.NRGBA, error) {
	width := img.Bounds().Dx() // disallows animated textures
	if width < 16 {
		return nil, errors.New("Image is too small.")
	}
	dst := imaging.New(width, width, color.Transparent)

	top := imaging.CropAnchor(img, width, width/16, imaging.Top)
	left := imaging.CropAnchor(img, width/16, width, imaging.Left)
	bottom := imaging.CropAnchor(img, width, width/16, imaging.Bottom)
	right := imaging.CropAnchor(img, width/16, width, imaging.Right)

	dst = imaging.Overlay(dst, top, image.Pt(0, 0), 1.0)
	dst = imaging.Overlay(dst, left, image.Pt(0, 0), 1.0)
	dst = imaging.Overlay(dst, right, image.Pt((width-(width/16)), 0), 1.0)
	dst = imaging.Overlay(dst, bottom, image.Pt(0, (width-(width/16))), 1.0)
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

	obsidianGlass := imaging.New(inImage.Bounds().Dx(), inImage.Bounds().Dx(), color.NRGBA{0, 0, 0, 0}) // disallows animated textures
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
