package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWCow(input_pack_path, output_pack_path string) error {
	stitch := "Cow"
	craft_path := "entity"
	clonia_path := "mobs_mc"

	in_path := input_pack_path + data.CraftPaths["entity"] + "cow/"
	out_path := output_pack_path + "/" + data.CloniaPaths["mobs_mc"] + "/" //mobs_mc_cow.png

	temperate_cow, err := imaging.Open(in_path + "temperate_cow.png")
	if err != nil {
		return openErrMsg(stitch, craft_path, "cow/temperate_cow.png")
	}

	if err = imaging.Save(CommonCow(temperate_cow), out_path+"mobs_mc_cow.png"); err != nil {
		return saveErrMsg(stitch, clonia_path, "mobs_mc_cow.png")
	}

	mooshroom_brown, err := imaging.Open(in_path + "/brown_mooshroom.png")
	if err != nil {
		return openErrMsg(stitch, craft_path, "cow/brown_mooshroom.png")
	}

	if err = imaging.Save(CommonCow(mooshroom_brown), out_path+"mobs_mc_mooshroom_brown.png"); err != nil {
		return saveErrMsg(stitch, clonia_path, "mobs_mc_mooshroom_brown.png")
	}

	mooshroom_red, err := imaging.Open(in_path + "/red_mooshroom.png")
	if err != nil {
		return openErrMsg(stitch, craft_path, "cow/red_mooshroom.png")
	}

	if err = imaging.Save(CommonCow(mooshroom_red), out_path+"mobs_mc_mooshroom.png"); err != nil {
		return saveErrMsg(stitch, clonia_path, "mobs_mc_mooshroom.png")
	}

	return nil
}

func CommonCow(inCow image.Image) image.Image {
	cow_length := inCow.Bounds().Dx()
	cow_height := inCow.Bounds().Dy()

	scale := cow_length / 64

	dst := imaging.New(cow_length, cow_height/2, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, inCow, image.Pt(0, 0))

	cow_nose := imaging.Crop(inCow, image.Rect(2*scale, 34*scale, 8*scale, 37*scale))
	dst = imaging.Paste(dst, cow_nose, image.Pt(7*scale, 11*scale))

	return dst
}
