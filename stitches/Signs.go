package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWSigns(input_pack_path, output_pack_path string, _ *configure.Config) error {
	stitch := "Signs"

	in_path := input_pack_path + data.CraftPaths["signs"]
	out_path := output_pack_path + data.CloniaPaths["signs"]

	sign_entity, err := imaging.Open(in_path + "spruce.png")
	if err != nil {
		return openErrMsg(stitch, "signs", "spruce.png")
	}

	new_sign_entity := Signs(sign_entity)
	if err := imaging.Save(new_sign_entity, out_path+"mcl_signs_sign_greyscale.png"); err != nil {
		return saveErrMsg(stitch, "signs", "mcl_signs_sign_greyscale.png")
	}

	in_path = input_pack_path + data.CraftPaths["item"]
	sign_item, err := imaging.Open(in_path + "spruce_sign.png")
	if err != nil {
		return openErrMsg(stitch, "signs", "spruce_sign.png")
	}

	new_sign_item := SignsItem(sign_item)
	if err := imaging.Save(new_sign_item, out_path+"default_sign_greyscale.png"); err != nil {
		return saveErrMsg(stitch, "signs", "default_sign_greyscale.png")
	}

	return nil
}

func Signs(mc_sign image.Image) image.Image {
	mc_sign = imaging.Grayscale(mc_sign)
	mc_sign = imaging.AdjustBrightness(mc_sign, 28)
	mc_sign = imaging.AdjustContrast(mc_sign, 65)
	dst := imaging.New(mc_sign.Bounds().Dx(), mc_sign.Bounds().Dy(), color.Transparent)
	dst = imaging.Paste(dst, mc_sign, image.Pt(0, 0))

	return dst
}

func SignsItem(mc_sign image.Image) image.Image {
	mc_sign = imaging.Grayscale(mc_sign)
	mc_sign = imaging.AdjustBrightness(mc_sign, 28)
	mc_sign = imaging.AdjustContrast(mc_sign, 65)
	dst := imaging.New(mc_sign.Bounds().Dx(), mc_sign.Bounds().Dy(), color.Transparent)
	dst = imaging.Paste(dst, mc_sign, image.Pt(0, 0))

	return dst
}
