package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	"github.com/disintegration/imaging"
)

// Only creates bed icons. Other stuff is handled elsewhere.

type Bed_t struct {
	Top       *image.NRGBA
	LongSide  *image.NRGBA
	ShortSide *image.NRGBA

	Inv *image.NRGBA
}

func RWBedInv(input_pack_path, output_pack_path string) error {
	stitch := "BedInv"

	for _, bed := range data.BedIcons {
		bedImg, err := imaging.Open(input_pack_path + bed.ReadPath())
		if err != nil {
			return openErrMsg(stitch, bed.InPath, bed.InTexture)
		}

		result := BedInv(bedImg)

		err = imaging.Save(result, output_pack_path+bed.SavePath())
		if err != nil {
			return saveErrMsg(stitch, bed.OutPath, bed.OutTexture)
		}
	}
	return nil
}

// I don't really care about making the pegs the beds stand on.
// It's a lot of effort for not much payoff.
func BedInv(src image.Image) *image.NRGBA {
	src = imaging.Resize(src, 64, 64, imaging.Box)
	dstSize := float64(64)
	dst := imaging.New(int(dstSize), int(dstSize), color.NRGBA{0, 0, 0, 0})

	var bed = Bed_t{
		Top:       imaging.New(16, 32, color.NRGBA{0, 0, 0, 0}),
		LongSide:  imaging.New(9, 32, color.NRGBA{0, 0, 0, 0}),
		ShortSide: imaging.New(16, 9, color.NRGBA{0, 0, 0, 0}),
	}

	bed.Top = imaging.Overlay(
		bed.Top,
		imaging.Crop(src, image.Rect(6, 6, 22, 22)),
		image.Pt(0, 0),
		1.0,
	)
	bed.Top = imaging.Overlay(
		bed.Top,
		imaging.Crop(src, image.Rect(6, 28, 22, 44)),
		image.Pt(0, 16),
		1.0,
	)

	bed.LongSide = imaging.Overlay(
		bed.LongSide,
		imaging.Crop(src, image.Rect(22, 6, 28, 22)),
		image.Pt(0, 0),
		1.0,
	)
	bed.LongSide = imaging.Overlay(
		bed.LongSide,
		imaging.Crop(src, image.Rect(22, 28, 28, 44)),
		image.Pt(0, 16),
		1.0,
	)
	bed.LongSide = imaging.Rotate270(bed.LongSide)

	bed.ShortSide = imaging.Overlay(
		bed.ShortSide,
		imaging.Rotate180(
			imaging.Crop(src, image.Rect(22, 22, 38, 28)),
		),
		image.Pt(0, 0),
		1.0,
	)

	// Now we make the icon.

	dst = RelativePerspectiveOverlay(
		*dst,
		*bed.Top,
		0.65, 0.25, 0.95, 0.40,
		0.05, 0.55, 0.35, 0.70,
	)

	dst = RelativePerspectiveOverlay(
		*dst,
		*bed.LongSide,
		0.35, 0.70, 0.95, 0.40,
		0.35, 0.90, 0.95, 0.60,
	)

	dst = RelativePerspectiveOverlay(
		*dst,
		*bed.ShortSide,
		0.05, 0.55, 0.35, 0.70,
		0.05, 0.75, 0.35, 0.90,
	)

	return dst
}
