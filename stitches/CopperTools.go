package stitches

import (
	"fmt"
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWCopperTools(inputPackPath, outputPackPath string) error {
	// stitch := "CopperTools"
	errReport := ""

	texturesForCopper := data.CopperStuffWithFallback

	for _, textureOfCopper := range texturesForCopper {
		possibleTexture := textureOfCopper.In[0]

		// We assume this is going to err, because copper tool textures are new to the game.
		copperItem, err := imaging.Open(inputPackPath + possibleTexture.FullPath())
		if err == nil {
			if saveErr := imaging.Save(copperItem, outputPackPath+textureOfCopper.SavePath()); saveErr != nil {
				errReport += fmt.Sprintf("\tFailed to save copper item \"%v\", giving up.\n", textureOfCopper.OutTexture)
			}
			continue
		}
		errReport += fmt.Sprintf("\tFailed to open copper item \"%v\", trying fallback!\n", possibleTexture.Texture)

		// FALLBACK
		possibleTexture = textureOfCopper.In[1]
		ironItem, err := imaging.Open(inputPackPath + possibleTexture.FullPath())
		if err != nil {
			errReport += fmt.Sprintf("\tFailed to open iron item \"%v\", giving up.\n", possibleTexture.Texture)
		}

		dst := CopperToolsFallback(ironItem)
		if saveErr := imaging.Save(dst, outputPackPath+textureOfCopper.SavePath()); saveErr != nil {
			errReport += fmt.Sprintf("\tFailed to save iron item \"%v\", giving up.\n", textureOfCopper.OutTexture)
		}
	}

	if errReport != "" {
		return fmt.Errorf("copper tools textures report:\n%v", errReport)
	}

	return nil
}

// Converts iron item textures into copper-looking item textures.
//
// TODO: add a config option that can force the fallback.
func CopperToolsFallback(ironItem image.Image) *image.NRGBA {
	dst := imaging.New(ironItem.Bounds().Dx(), ironItem.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, ironItem, image.Point{0, 0}, 1.0)
	dst = imaging.AdjustFunc(dst,
		func(c color.NRGBA) color.NRGBA {
			r := int(c.R)
			g := int(c.G)
			b := int(c.B)

			if (r > g+20 || r < g-20) && (r > b+20 || r < b-20) {
				return c
			}

			g = (r * 55) / 100
			b = (r * 46) / 100

			return color.NRGBA{c.R, uint8(g), uint8(b), c.A}
		})
	return dst
}
