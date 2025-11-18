package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWStonecutter(input_pack_path, output_pack_path string, _ *configure.Config) error {
	stitch := "Stonecutter"
	in_path := input_pack_path + data.CraftPaths["block"]
	out_path := output_pack_path + data.CloniaPaths["stonecutter"]

	file_name := "stonecutter_saw.png"
	saw, err := imaging.Open(in_path + file_name)
	if err != nil {
		return openErrMsg(stitch, in_path, file_name)
	}
	if saw.Bounds().Dx()%16 != 0 {
		return openErrMsg(stitch, in_path, file_name)
	}
	scale := saw.Bounds().Dx() / 16
	numOfFrames := saw.Bounds().Dy() / saw.Bounds().Dx()

	file_name = "stonecutter_side.png"
	side, err := imaging.Open(in_path + file_name)
	if err != nil {
		return openErrMsg(stitch, in_path, file_name)
	}
	dst := imaging.New(saw.Bounds().Dx(), saw.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Overlay(dst, saw, image.Point{0, -9 * scale}, 1.0)
	for i := range numOfFrames {
		dst = imaging.Overlay(dst, side, image.Point{0, i * side.Bounds().Dx()}, 1.0)
	}
	file_name = "mcl_stonecutter_saw.png"
	if err := imaging.Save(dst, out_path+file_name); err != nil {
		saveErrMsg(stitch, out_path, file_name)
	}
	return nil
}
