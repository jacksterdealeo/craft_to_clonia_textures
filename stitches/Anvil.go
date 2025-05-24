package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWAnvil(input_pack_path, output_pack_path string) error {
	stitch := "Anvil"

	in_path := input_pack_path + data.CraftPaths["block"]
	out_path := output_pack_path + data.CloniaPaths["anvils"]
	abase, err := imaging.Open(in_path + "anvil.png")
	if err != nil {
		return openErrMsg(stitch, "block", "anvil.png")
	}
	a0, err := imaging.Open(in_path + "anvil_top.png")
	if err != nil {
		return openErrMsg(stitch, "block", "anvil_top.png")
	}
	a1, err := imaging.Open(in_path + "chipped_anvil_top.png")
	if err != nil {
		return openErrMsg(stitch, "block", "chipped_anvil_top.png")
	}
	a2, err := imaging.Open(in_path + "damaged_anvil_top.png")
	if err != nil {
		return openErrMsg(stitch, "block", "damaged_anvil_top.png")
	}

	new_a0, new_a1, new_a2 := Anvil(abase, a0, a1, a2)

	save_name := "mcl_anvils_anvil_top_damaged_0.png"
	if err = imaging.Save(new_a0, out_path+"mcl_anvils_anvil_top_damaged_0.png"); err != nil {
		return saveErrMsg(stitch, "anvils", save_name)
	}
	if err = imaging.Save(new_a1, out_path+"mcl_anvils_anvil_top_damaged_1.png"); err != nil {
		return saveErrMsg(stitch, "anvils", save_name)
	}
	if err = imaging.Save(new_a2, out_path+"mcl_anvils_anvil_top_damaged_2.png"); err != nil {
		return saveErrMsg(stitch, "anvils", save_name)
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
