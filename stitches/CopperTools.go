package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

/*
Problem:
Minecraft added copper armor and tools.
Most texture packs will not be updated for a long while, so my version is still useful.
I still need to support Minecraft's new textures for the users that want them.

Solution:
Make my old convert function a fallback,
and add a config option that can force the fallback.
*/

func RWCopperTools(input_pack_path, output_pack_path string) error {
	stitch := "CopperTools"

	textures_for_copper := data.CopperStuffMod
	// in_path := input_pack_path + data.CraftPaths["items"]
	// out_path := output_pack_path + data.CloniaPaths["copper_stuff"]

	for _, e := range textures_for_copper {
		ironItem, err := imaging.Open(input_pack_path + e.ReadPath())
		if err != nil {
			// Returns if even 1 texture fails. Maybe don't?
			return openErrMsg(stitch, "items", e.InTexture)
		} else {
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
			if err = imaging.Save(dst, output_pack_path+e.SavePath()); err != nil {
				// Returns if even 1 texture fails. Maybe don't?
				return saveErrMsg(stitch, "copper_stuff", e.SavePath())
			}
		}
	}

	return nil
}

func CopperTools() {
}
