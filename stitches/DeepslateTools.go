package stitches

import (
	"fmt"
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWDeepslateTools(inputPackPath, outputPackPath string) error {
	// stitch := "DeepslateTools"
	errReport := ""

	for _, item := range data.VLDeepslate {
		stoneItem, err := imaging.Open(inputPackPath + item.ReadPath())
		if err != nil {
			errReport += fmt.Sprintf("\tFailed to open iron item \"%v\", giving up.\n", item.ReadPath())
			continue
		}

		dst := DeepslateTool(stoneItem)
		if saveErr := imaging.Save(dst, outputPackPath+item.SavePath()); saveErr != nil {
			errReport += fmt.Sprintf("\tFailed to save iron item \"%v\", giving up.\n", item.SavePath())
		}
	}

	if errReport != "" {
		return fmt.Errorf("Deepslate textures report:\n%v", errReport)
	}

	return nil
}

// Converts stone item textures into deepslate-looking item textures.
func DeepslateTool(stoneItem image.Image) *image.NRGBA {
	dst := imaging.New(stoneItem.Bounds().Dx(), stoneItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, stoneItem, image.Point{0, 0}, 1.0)
	dst = imaging.AdjustFunc(dst,
		func(c color.NRGBA) color.NRGBA {
			r := int(c.R)
			g := int(c.G)
			b := int(c.B)

			if !within10(r, g, b) {
				return c
			}

			g = g / 2
			b = b / 2
			r = r / 2

			return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
		})
	return dst
}

func within10(a, b, c int) bool {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max-min <= 10
}
