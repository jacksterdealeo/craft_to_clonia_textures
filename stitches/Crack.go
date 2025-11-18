package stitches

import (
	"image"
	"image/color"
	"strconv"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWCrack(input_pack_path, output_pack_path string, _ *configure.Config) error {
	stitch := "Crack"

	in_path := input_pack_path + data.CraftPaths["block"]
	out_path := output_pack_path + data.CloniaPaths["hud_base_textures"]

	var err error
	var destroy_imgs [10]image.Image
	for i := 0; i <= 9; i++ {
		file_name := in_path + "destroy_stage_" + strconv.Itoa(i) + ".png"
		destroy_imgs[i], err = imaging.Open(file_name)
		if err != nil {
			return openErrMsg(stitch, in_path, file_name)
		}
	}
	dst := Crack(destroy_imgs)
	if err = imaging.Save(dst, out_path+"crack_anylength.png"); err != nil {
		return saveErrMsg(stitch, out_path, "crack_anylength.png")
	}
	return nil
}

func Crack(destroy_imgs [10]image.Image) image.Image {
	dst := imaging.New(destroy_imgs[0].Bounds().Dx(), destroy_imgs[0].Bounds().Dy()*10, color.NRGBA{0, 0, 0, 0})
	for i := 0; i <= 9; i++ {
		dst = imaging.Paste(dst, destroy_imgs[i], image.Pt(0, i*destroy_imgs[i].Bounds().Dy()))
	}
	return dst
}
