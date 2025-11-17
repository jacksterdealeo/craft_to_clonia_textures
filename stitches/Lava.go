package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWLava(input_pack_path, output_pack_path string, _ *configure.Config) error {
	inPath := input_pack_path + "/" + data.CraftPaths["block"] + "/"
	outPath := output_pack_path + "/" + data.CloniaPaths["core"] + "/"
	/*
		craft lava
		  still   :  16 x 512
		  flowing :  32 x 1024
		clonia lava
			still   :  16 x 256
		  flowing :  16 x 1024
	*/
	lavaFlowing, err := imaging.Open(inPath + "lava_flow.png")
	if err != nil {
		return openErrMsg("Lava", "block", "lava_flow.png")
	} else {
		lavaStillX := lavaFlowing.Bounds().Dx()
		lavaStillY := lavaFlowing.Bounds().Dy()
		dst := imaging.New(lavaStillX/2, lavaStillY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Overlay(dst, lavaFlowing, image.Point{0, 0}, 1.0)
		if err = imaging.Save(dst, outPath+"default_lava_flowing_animated.png"); err != nil {
			return saveErrMsg("Lava", "core", "default_lava_flowing_animated.png")
		}
	}
	return nil

}
