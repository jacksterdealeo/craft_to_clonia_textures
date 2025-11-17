package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	imaging "github.com/disintegration/imaging"
)

func RWPackIcon(input_pack_path, output_pack_path string, _ *configure.Config) error {
	stitch := "Pack Icon"

	src, err := imaging.Open(input_pack_path + "/pack.png")
	if err != nil {
		return openErrMsg(stitch, "", "pack.png")
	}

	background := imaging.Fill(src, 350, 233, imaging.Center, imaging.Lanczos)
	background = imaging.Blur(background, 10)
	foreground := imaging.Resize(src, 233, 0, imaging.Lanczos)

	dst := imaging.New(350, 233, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, background, image.Pt(0, 0))
	dst = imaging.OverlayCenter(dst, foreground, 1.0)
	err = imaging.Save(dst, output_pack_path+"/screenshot.png")
	if err != nil {
		return saveErrMsg(stitch, "", "screenshot.png")
	}
	return nil
}
