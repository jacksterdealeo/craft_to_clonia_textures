package stitches

import (
	"image"
	"image/color"
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWSmallDripLeaf(inputPackPath, outputPackPath string, _ *configure.Config) error {
	inPath := filepath.Join(inputPackPath, data.GetCraftPath("block"))
	outPath := filepath.Join(outputPackPath, data.GetCloniaPath("lush_caves"))

	stitch := "SmallDripLeaf"

	fileName := "small_dripleaf_side.png"
	Side, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}
	fileName = "small_dripleaf_stem_bottom.png"
	StemBottom, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}
	fileName = "small_dripleaf_stem_top.png"
	StemTop, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}
	fileName = "small_dripleaf_top.png"
	Top, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}

	// NOTE: Small Drip Leaf cannot use higher scales, unlike Large Drip Leaf.
	Side = imaging.Fit(Side, 8, 8, imaging.Lanczos)
	StemBottom = imaging.Fit(StemBottom, 8, 8, imaging.Lanczos)
	StemTop = imaging.Fit(StemTop, 8, 8, imaging.Lanczos)
	Top = imaging.Fit(Top, 16, 16, imaging.Lanczos)
	Out := imaging.New(16, 16, color.Transparent)

	Top = imaging.Crop(Top, image.Rect(0, 0, 8, 8))
	Top = imaging.Rotate180(Top)
	StemTop = imaging.Rotate90(StemTop)
	StemBottom = imaging.Rotate90(StemBottom)

	Out = imaging.Overlay(Out, Side, image.Pt(Out.Bounds().Dx()/2, 0), 1.0)
	Out = imaging.Overlay(Out, StemTop, image.Pt(0, Out.Bounds().Dx()/2), 1.0)
	Out = imaging.Overlay(Out, StemBottom, image.Pt(Out.Bounds().Dx()/2, Out.Bounds().Dy()/2), 1.0)
	Out = imaging.Overlay(Out, Top, image.Pt(0, 0), 1.0)

	Out = ClampTransparency(Out, 60)

	fileName = "mcl_lush_caves_dripleaf_small.png"
	if err = imaging.Save(Out, filepath.Join(outPath, fileName)); err != nil {
		return saveErrMsg(stitch, outPath, fileName)
	}

	return nil
}
