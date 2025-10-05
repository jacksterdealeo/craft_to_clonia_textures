package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	"github.com/disintegration/imaging"
)

func RWUI(input_pack_path, output_pack_path string) error {
	stitch := "UI"
	readFails := make([][2]string, 0)
	saveFails := make([][2]string, 0)

	in_path := input_pack_path + "/assets/minecraft/textures/gui/sprites/container/furnace/"
	out_path := output_pack_path + data.CloniaPaths["furnaces"]

	burn_progress, err := imaging.Open(in_path + "burn_progress.png")
	if err != nil {
		readFails = append(readFails, [2]string{stitch, "burn_progress.png"})
	} else {
		gui_furnace_arrow_bg, gui_furnace_arrow_fg := FurnaceArrow(burn_progress)
		if err := imaging.Save(gui_furnace_arrow_bg, out_path+"gui_furnace_arrow_bg.png"); err != nil {
			saveFails = append(saveFails, [2]string{stitch, "gui_furnace_arrow_bg.png"})
		}
		if err := imaging.Save(gui_furnace_arrow_fg, out_path+"gui_furnace_arrow_fg.png"); err != nil {
			saveFails = append(saveFails, [2]string{stitch, "gui_furnace_arrow_fg.png"})
		}
	}

	lit_progress, err := imaging.Open(in_path + "lit_progress.png")
	if err != nil {
		readFails = append(readFails, [2]string{stitch, "lit_progress.png"})
	} else {
		default_furnace_fire_bg, default_furnace_fire_fg := FurnaceFire(lit_progress)
		if err := imaging.Save(default_furnace_fire_bg, out_path+"default_furnace_fire_bg.png"); err != nil {
			saveFails = append(saveFails, [2]string{stitch, "default_furnace_fire_bg.png"})
		}
		if err := imaging.Save(default_furnace_fire_fg, out_path+"default_furnace_fire_fg.png"); err != nil {
			saveFails = append(saveFails, [2]string{stitch, "default_furnace_fire_fg.png"})
		}
	}

	return multiErrMsg(stitch, readFails, saveFails)
}

func FurnaceArrow(burn_progress image.Image) (gui_furnace_arrow_bg, gui_furnace_arrow_fg *image.NRGBA) {
	burn_progress = imaging.Rotate90(burn_progress)
	burn_progress = imaging.Fit(burn_progress, 32, 48, imaging.NearestNeighbor)

	backgroundColor := burn_progress.At(0, 0)
	white := color.NRGBA{255, 255, 255, 255}
	darkGray := color.NRGBA{104, 104, 104, 255}

	burn_progress = imaging.AdjustFunc(burn_progress,
		func(c color.NRGBA) color.NRGBA {
			if c == backgroundColor {
				return color.NRGBA{0, 0, 0, 0}
			}
			return c
		})

	gui_furnace_arrow_fg = imaging.Clone(burn_progress)

	burn_progress = imaging.AdjustFunc(burn_progress,
		func(c color.NRGBA) color.NRGBA {
			if c == white {
				return darkGray
			}
			return c
		})

	gui_furnace_arrow_bg = imaging.Clone(burn_progress)
	return
}

func FurnaceFire(lit_progress image.Image) (default_furnace_fire_bg, default_furnace_fire_fg *image.NRGBA) {
	backgroundColor := color.NRGBAModel.Convert(lit_progress.At(0, 0)).(color.NRGBA)
	darkGray := color.NRGBA{104, 104, 104, 255}

	lit_progress = imaging.AdjustFunc(lit_progress,
		func(c color.NRGBA) color.NRGBA {
			if c == backgroundColor {
				return color.NRGBA{0, 0, 0, 0}
			}
			return c
		})

	default_furnace_fire_fg = imaging.Clone(lit_progress)

	lit_progress = imaging.AdjustFunc(lit_progress,
		func(c color.NRGBA) color.NRGBA {
			if c != (color.NRGBA{0, 0, 0, 0}) {
				return darkGray
			}
			return c
		})

	default_furnace_fire_bg = imaging.Clone(lit_progress)

	return
}
