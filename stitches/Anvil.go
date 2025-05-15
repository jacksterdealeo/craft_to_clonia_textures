package stitches

import (
	"image"
	"image/color"

	imaging "github.com/disintegration/imaging"
)

// base = anvil.png
// top = anvil_top.png
// chipped_top = chipped_anvil_top.png
// damaged_top = damaged_anvil_top.png
func Anvil(base, top, chipped_top, damaged_top image.Image) (new_top, new_chipped_top, new_damaged_top image.Image) {
	anvilX := base.Bounds().Dx()
	anvilY := base.Bounds().Dy()

	dst := imaging.New(anvilX, anvilY, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, base, image.Pt(0, 0))

	dst = imaging.OverlayCenter(dst, top, 1.0)
	new_top = dst

	dst = imaging.OverlayCenter(dst, chipped_top, 1.0)
	new_chipped_top = dst

	dst = imaging.OverlayCenter(dst, damaged_top, 1.0)
	new_damaged_top = dst

	return
}
