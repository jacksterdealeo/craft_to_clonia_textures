package stitches

import (
	"image"
	"image/color"
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWBigDripLeaf(inputPackPath, outputPackPath string, _ *configure.Config) error {
	inPath := filepath.Join(inputPackPath, data.GetCraftPath("block"))
	outPath := filepath.Join(outputPackPath, data.GetCloniaPath("lush_caves"))

	stitch := "BigDripLeaf"

	fileName := "big_dripleaf_side.png"
	Side, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}

	fileName = "big_dripleaf_stem.png"
	Stem, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}
	fileName = "big_dripleaf_tip.png"
	Tip, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}
	fileName = "big_dripleaf_top.png"
	Top, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}
	Top = imaging.Rotate180(Top)

	Out := imaging.New(Top.Bounds().Dx()*2, Top.Bounds().Dy()*2, color.Transparent)
	Out = imaging.Overlay(Out, Side, image.Pt(Out.Bounds().Dx()/2, 0), 1.0)
	Out = imaging.Overlay(Out, Stem, image.Pt(Out.Bounds().Dx()/2, Out.Bounds().Dy()/2), 1.0)
	Out = imaging.Overlay(Out, Tip, image.Pt(0, Out.Bounds().Dx()/2), 1.0)
	Out = imaging.Overlay(Out, Top, image.Pt(0, 0), 1.0)

	fileName = "mcl_lush_caves_dripleaf_big.png"
	if err = imaging.Save(Out, filepath.Join(outPath, fileName)); err != nil {
		return saveErrMsg(stitch, outPath, fileName)
	}

	return nil
}
