package stitches

import (
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWCrosshair(input_pack_path, output_pack_path string, config *configure.Config) error {
	stitch := "Crosshair"

	in_path := input_pack_path + data.CraftPaths["hud"]
	out_path := output_pack_path + data.CloniaPaths["hud_base_textures"]
	var err error

	file_name := "crosshair.png"
	img, err := imaging.Open(filepath.Join(in_path, file_name))
	if err != nil {
		return openErrMsg(stitch, in_path, file_name)
	}
	dst := imaging.Resize(img, 30, 30, imaging.Lanczos)
	if err = imaging.Save(dst, filepath.Join(out_path, "crosshair.png")); err != nil {
		return saveErrMsg(stitch, out_path, "crosshair.png")
	}
	if config.EnlargeObjectCrosshair == true {
		dst = imaging.Resize(img, 45, 45, imaging.Lanczos)
	}
	if err = imaging.Save(dst, filepath.Join(out_path, "object_crosshair.png")); err != nil {
		return saveErrMsg(stitch, out_path, "object_crosshair.png")
	}
	return nil
}
