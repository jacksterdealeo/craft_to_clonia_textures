package stitches

import (
	"image"
	"image/color"
	"path"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWLava(input_pack_path, output_pack_path string, _ *configure.Config) error {
	inPath := path.Join(input_pack_path, data.CraftPaths["block"])
	outPath := path.Join(output_pack_path, data.CloniaPaths["core"])
	voxPath := path.Join(output_pack_path, data.CloniaPaths["vl"])
	/*
		craft lava
		  still   :  16 x 512
		  flowing :  32 x 1024
		clonia lava
			still   :  16 x 256
		  flowing :  16 x 1024
	*/
	lavaFlowing, err := imaging.Open(path.Join(inPath, "lava_flow.png"))
	if err != nil {
		return openErrMsg("Lava", "block", "lava_flow.png")
	} else {
		lavaFlowX := lavaFlowing.Bounds().Dx()
		lavaFlowY := lavaFlowing.Bounds().Dy()
		dst := imaging.New(lavaFlowX/2, lavaFlowY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Overlay(dst, lavaFlowing, image.Point{0, 0}, 1.0)
		if err = imaging.Save(dst, path.Join(voxPath, "mcl_core_lava_flow_animation.png")); err != nil {
			return saveErrMsg("Lava", "vl", "mcl_core_lava_flow_animation.png")
		}
		if err = imaging.Save(dst, path.Join(outPath, "default_lava_flowing_animated.png")); err != nil {
			return saveErrMsg("Lava", "core", "default_lava_flowing_animated.png")
		}
	}
	return nil

}
