package stitches

import (
	"fmt"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWPotionIndicator(input_pack_path, output_pack_path string) error {
	stitch := "PotionIndicator"
	var errors string

	bg_in_path := input_pack_path + data.CraftPaths["hud"]
	bg_in_texture := "effect_background.png"
	effect_bg, err := imaging.Open(bg_in_path + bg_in_texture)
	if err != nil {
		effect_bg = imaging.New(24, 24, color.Transparent)
		errors += fmt.Sprintln(openErrMsg(stitch, "hud", bg_in_texture))
	}

	for _, e := range data.SimpleMobEffect {
		input_texture_dir := input_pack_path + data.CraftPaths[e.InPath] + e.InTexture
		output_texture_dir := output_pack_path + data.CloniaPaths[e.OutPath] + e.OutTexture

		dst, err := imaging.Open(input_texture_dir)
		if err != nil {
			errors += fmt.Sprintln(openErrMsg(stitch, e.InPath, e.InTexture))
			continue
		}
		dst = imaging.OverlayCenter(effect_bg, dst, 1.0)
		err = imaging.Save(dst, output_texture_dir)
		if err != nil {
			errors += fmt.Sprintln(saveErrMsg(stitch, e.OutPath, e.OutTexture))
			continue
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("%v", errors)
	}
	return nil
}
