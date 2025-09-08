package stitches


import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWRedstoneLamp(inputPackPath, outputPackPath string) error {
	stitch := "Campfire"
	craftPath := "block"
	cloniaPath := "vl"

	inPath := inputPackPath + data.CraftPaths[craftPath]
	outPath := outputPackPath + "/" + data.CloniaPaths[cloniaPath]

	inBlock := "redstone_lamp_on.png"
	outBlock := "mcl_lightstone_mask.png"
	lampOn, err := imaging.Open(inPath + inBlock)
	if err != nil {
		return openErrMsg(stitch, craftPath, inBlock)
	}

	dst := RedstoneLampMask(lampOn)
	if saveErr := imaging.Save(dst, outPath + outBlock); saveErr != nil {
		return saveErrMsg(stitch, craftPath, outBlock)
	}

	return nil
}

func RedstoneLampMask(lamp image.Image) *image.NRGBA {
	dst := imaging.New(lamp.Bounds().Dx(), lamp.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, lamp, image.Point{0, 0}, 1.0)
	dst = imaging.AdjustFunc(dst,
		func(c color.NRGBA) color.NRGBA {
			// 0 is clear
			if c.R < 68 {
				return color.NRGBA{255, 255, 255, 0}
			}
			return color.NRGBA{255, 255, 255, 255}
		})

	return dst
}
