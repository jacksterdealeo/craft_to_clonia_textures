package stitches

import (
	"image/color"
	"path"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

// {"block", "cobweb.png", "core", "mcl_core_web.png", 1},
func RWCobweb(input_pack_path, output_pack_path string, config *configure.Config) error {
	stitch := "Cobweb"

	in_path := input_pack_path + data.CraftPaths["block"]
	out_path := output_pack_path + data.CloniaPaths["core"]
	var err error

	file_name := "cobweb.png"
	img, err := imaging.Open(path.Join(in_path, file_name))
	if err != nil {
		return openErrMsg(stitch, in_path, file_name)
	}

	dst := imaging.AdjustFunc(img, func(c color.NRGBA) color.NRGBA {
		if c.A > 1 {
			c.A = 255
		}
		return c
	})

	if err = imaging.Save(dst, path.Join(out_path, "mcl_core_web.png")); err != nil {
		return saveErrMsg(stitch, out_path, "mcl_core_web.png")
	}
	if config.EnlargeObjectCrosshair == true {
		dst = imaging.Resize(img, 45, 45, imaging.Lanczos)
	}
	if err = imaging.Save(dst, path.Join(out_path, "object_crosshair.png")); err != nil {
		return saveErrMsg(stitch, out_path, "object_crosshair.png")
	}
	return nil
}
