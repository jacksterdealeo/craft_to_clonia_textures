package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWFlowerPot(input_pack_path, output_pack_path string, _ *configure.Config) error {
	stitch := "Flower Pot"
	craft_path := "block"
	clonia_path := "flowerpots"

	in_path := input_pack_path + data.CraftPaths[craft_path]
	out_path := output_pack_path + data.CloniaPaths[clonia_path]

	pot, err := imaging.Open(in_path + "flower_pot.png")
	if err != nil {
		return openErrMsg(stitch, craft_path, "flower_pot.png")
	}

	dirt, err := imaging.Open(in_path + "dirt.png")
	if err != nil {
		return openErrMsg(stitch, craft_path, "dirt.png")
	}

	flowerpot_texture := FlowerPot(pot, dirt)
	file_name := "mcl_flowerpots_flowerpot.png"

	if err = imaging.Save(flowerpot_texture, out_path+file_name); err != nil {
		return saveErrMsg(stitch, clonia_path, file_name)
	}

	return nil
}

func FlowerPot(pot, dirt image.Image) image.Image {
	potX := pot.Bounds().Dx()
	potY := pot.Bounds().Dy()
	scale := potX / 16

	dirt = imaging.CropCenter(dirt, 4*scale, 4*scale)

	dst := imaging.New((potX)*2, (potY)*2, color.NRGBA{0, 0, 0, 0})
	fPotTop := imaging.Crop(pot, image.Rectangle{
		image.Point{5 * scale, 5 * scale}, image.Point{11 * scale, 11 * scale}})
	fPotSide := imaging.Crop(pot, image.Rectangle{
		image.Point{5 * scale, 10 * scale}, image.Point{11 * scale, 16 * scale}})
	fPotInner := imaging.Crop(pot, image.Rectangle{
		image.Point{5 * scale, 10 * scale}, image.Point{9 * scale, 12 * scale}})

	// pot sides
	dst = imaging.Overlay(dst, fPotSide, image.Point{0 * scale, 20 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotSide, image.Point{6 * scale, 20 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotSide, image.Point{12 * scale, 20 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotSide, image.Point{18 * scale, 20 * scale}, 1.0)
	// pot bottom
	dst = imaging.Overlay(dst, fPotSide, image.Point{22 * scale, 26 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotTop, image.Point{22 * scale, 26 * scale}, 1.0)
	// pot top
	dst = imaging.Overlay(dst, fPotTop, image.Point{16 * scale, 26 * scale}, 1.0)
	// pot inside
	dst = imaging.Overlay(dst, fPotInner, image.Point{0 * scale, 30 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotInner, image.Point{4 * scale, 30 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotInner, image.Point{8 * scale, 30 * scale}, 1.0)
	dst = imaging.Overlay(dst, fPotInner, image.Point{12 * scale, 30 * scale}, 1.0)
	// dirt
	dst = imaging.Overlay(dst, dirt, image.Point{4 * scale, 26 * scale}, 1.0)

	return dst
}
