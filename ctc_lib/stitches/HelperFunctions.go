package stitches

import (
	"image"
	"image/color"

	imaging "github.com/disintegration/imaging"
)

func CropToScale(img image.Image, x1, y1, x2, y2, scale int) *image.NRGBA {
	return imaging.Crop(img, image.Rectangle{
		image.Point{x1 * scale, y1 * scale}, image.Point{x2 * scale, y2 * scale}})
}

func ImageToNRGBA(img image.Image) image.NRGBA {
	return *imaging.Paste(imaging.New(img.Bounds().Dx(), img.Bounds().Dy(), color.Transparent),
		img,
		image.Point{0, 0},
	)
}

func MakeBrown(img *image.NRGBA) *image.NRGBA {
	img = imaging.AdjustFunc(img,
		func(c color.NRGBA) color.NRGBA {
			r := float64(c.R) * .60
			g := float64(c.G) * .39999
			b := float64(c.B) * .19999
			return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
		})
	return img
}

func ClampTransparency(img *image.NRGBA, line uint8) *image.NRGBA {
	img = imaging.AdjustFunc(img,
		func(c color.NRGBA) color.NRGBA {
			if c.A < line {
				return color.NRGBA{c.R, c.G, c.B, 0}
			}
			return color.NRGBA{c.R, c.G, c.B, 255}
		})
	return img

}
