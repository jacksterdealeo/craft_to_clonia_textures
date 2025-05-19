package stitches

import (
	"fmt"
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWAnvil(input_pack_path, output_pack_path string) error {
	in_path := input_pack_path + data.CraftPaths["block"]
	out_path := output_pack_path + data.CloniaPaths["anvils"]
	abase, err := imaging.Open(in_path + "anvil.png")
	if err != nil {
		return fmt.Errorf("block::anvil.png failed to open! Skipping the rest!")
	}
	a0, err := imaging.Open(in_path + "anvil_top.png")
	if err != nil {
		return fmt.Errorf("block::anvil_top.png failed to open! Skipping the rest!")
	}
	a1, err := imaging.Open(in_path + "chipped_anvil_top.png")
	if err != nil {
		return fmt.Errorf("block::chipped_anvil_top.png failed to open! Skipping the rest!")
	}
	a2, err := imaging.Open(in_path + "damaged_anvil_top.png")
	if err != nil {
		return fmt.Errorf("block::damaged_anvil_top.png failed to open!")
	}

	new_a0, new_a1, new_a2 := Anvil(abase, a0, a1, a2)

	if err = imaging.Save(new_a0, out_path+"mcl_anvils_anvil_top_damaged_0.png"); err != nil {
		return fmt.Errorf("mcl_anvils_anvil_top_damaged_0.png failed to save! Skipping the rest!")
	}
	if err = imaging.Save(new_a1, out_path+"mcl_anvils_anvil_top_damaged_1.png"); err != nil {
		return fmt.Errorf("mcl_anvils_anvil_top_damaged_1.png failed to save! Skipping the rest!")
	}
	if err = imaging.Save(new_a2, out_path+"mcl_anvils_anvil_top_damaged_2.png"); err != nil {
		return fmt.Errorf("mcl_anvils_anvil_top_damaged_2.png failed to save!")
	}
	return nil
}

// base = anvil.png
// top = anvil_top.png
// chipped_top = chipped_anvil_top.png
// damaged_top = damaged_anvil_top.png
func Anvil(base, top, chipped_top, damaged_top image.Image) (new_top, new_chipped_top, new_damaged_top image.Image) {
	anvilX := base.Bounds().Dx()
	anvilY := base.Bounds().Dy()

	dst := imaging.New(anvilX, anvilY, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, base, image.Pt(0, 0))

	dst = imaging.OverlayCenter(dst, top, 1.0)
	new_top = dst

	dst = imaging.OverlayCenter(dst, chipped_top, 1.0)
	new_chipped_top = dst

	dst = imaging.OverlayCenter(dst, damaged_top, 1.0)
	new_damaged_top = dst

	return
}
