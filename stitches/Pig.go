package stitches

import (
	"fmt"
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWPig(input_pack_path, output_pack_path string) error {
	in_path := input_pack_path + data.CraftPaths["entity"] + "pig/"
	out_path := output_pack_path + "/" + data.CloniaPaths["mobs_mc"] + "/" //mobs_mc_pig.png

	temperate_pig, err := imaging.Open(in_path + "/temperate_pig.png")
	if err != nil {
		return fmt.Errorf("entity::temperate_pig.png failed to open. %v, %v", in_path, out_path)
	}

	if err = imaging.Save(CommonPig(temperate_pig), out_path+"mobs_mc_pig.png"); err != nil {
		return fmt.Errorf("mobs_mc::mobs_mc_pig.png failed to save! Skipping the rest!")
	}

	return nil
}

func CommonPig(inPig image.Image) image.Image {
	pig_length := inPig.Bounds().Dx()
	pig_height := inPig.Bounds().Dy()

	dst := imaging.New(pig_length, pig_height/2, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, inPig, image.Pt(0, 0))

	return dst
}
