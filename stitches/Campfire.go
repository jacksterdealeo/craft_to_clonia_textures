package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWCampfire(input_pack_path, output_pack_path string) error {
	stitch := "Campfire"
	craft_path := "block"
	clonia_path := "campfires"

	_ = clonia_path

	var block string

	in_path := input_pack_path + data.CraftPaths["block"]
	out_path := output_pack_path + "/" + data.CloniaPaths["campfires"]

	_ = out_path

	block = "campfire_log_lit.png"
	campfire_log_lit, err := imaging.Open(in_path + block)
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}
	_ = campfire_log_lit

	block = "soul_campfire_log_lit.png"
	soul_campfire_log_lit, err := imaging.Open(in_path + block)
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}
	_ = soul_campfire_log_lit

	campfire_log, err := imaging.Open(in_path + "campfire_log.png")
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}
	_ = campfire_log

	fire, err := imaging.Open(in_path + "campfire_fire.png")
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}
	_ = fire

	soulfire, err := imaging.Open(in_path + "soul_campfire_fire.png")
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}
	_ = soulfire

	return nil
}

func CampfireLogLit(texture image.Image) image.Image {
	dst := imaging.New(texture.Bounds().Dx()*2, texture.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, texture, image.Pt(0, 0))
	return dst
}
