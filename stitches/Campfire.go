package stitches

import (
	"fmt"
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWCampfire(input_pack_path, output_pack_path string, _ *configure.Config) error {
	//stitch := "Campfire"

	in_path := input_pack_path + data.CraftPaths["block"]
	out_path := output_pack_path + data.CloniaPaths["campfires"]

	fails := []string{}

	campfire_log_lit, err := imaging.Open(in_path + "campfire_log_lit.png")
	if err != nil {
		fails = append(fails, "campfires::campfire_log_lit.png failed to open!")
	} else {
		dst := imaging.New(campfire_log_lit.Bounds().Dx()*2, campfire_log_lit.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, campfire_log_lit, image.Pt(0, 0))
		if err := imaging.Save(dst, out_path+"mcl_campfires_campfire_log_lit.png"); err != nil {
			fails = append(fails, "mcl_campfires_campfire_log_lit.png failed to save!")
		}
	}

	soul_campfire_log_lit, err := imaging.Open(in_path + "soul_campfire_log_lit.png")
	if err != nil {
		fails = append(fails, "campfires::soul_campfire_log_lit.png failed to open!")
	} else {
		dst := imaging.New(soul_campfire_log_lit.Bounds().Dx()*2, soul_campfire_log_lit.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, soul_campfire_log_lit, image.Pt(0, 0))
		if err := imaging.Save(dst, out_path+"mcl_campfires_soul_campfire_log_lit.png"); err != nil {
			fails = append(fails, "mcl_campfires_soul_campfire_log_lit.png failed to save!")
		}
	}

	campfire_log, err := imaging.Open(in_path + "campfire_log.png")
	if err != nil {
		fails = append(fails, "campfires::campfire_fire.png failed to open!")
	} else {
		dst := imaging.New(campfire_log.Bounds().Dx()*2, campfire_log.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, campfire_log, image.Pt(0, 0))
		if err := imaging.Save(dst, out_path+"mcl_campfires_log.png"); err != nil {
			fails = append(fails, "mcl_campfires_log.png failed to save!")
		}
	}

	fire, err := imaging.Open(in_path + "campfire_fire.png")
	if err != nil {
		fails = append(fails, "campfires::campfire_fire.png failed to open!")
	} else {
		dst := imaging.New(fire.Bounds().Dx()*2, fire.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, fire, image.Pt(fire.Bounds().Dx(), 0))
		if err := imaging.Save(dst, out_path+"mcl_campfires_campfire_fire.png"); err != nil {
			fails = append(fails, "mcl_campfires_campfire_fire.png failed to save!")
		}
	}

	soulfire, err := imaging.Open(in_path + "soul_campfire_fire.png")
	if err != nil {
		fails = append(fails, "campfires::soul_campfire_fire.png failed to open!")
	} else {
		dst := imaging.New(soulfire.Bounds().Dx()*2, soulfire.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, soulfire, image.Pt(soulfire.Bounds().Dx(), 0))
		if err := imaging.Save(dst, out_path+"mcl_campfires_soul_campfire_fire.png"); err != nil {
			fails = append(fails, "mcl_campfires_soul_campfire_fire.png failed to save!")
		}
	}

	if len(fails) > 0 {
		return fmt.Errorf("%v\n", fails)
	}
	return nil
}
