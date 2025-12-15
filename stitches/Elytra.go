package stitches

import (
	"image"
	"image/color"
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

// Elytra in Mineclonia A.T.M. cannot be scaled.
func RWElytra(inputPathPath, outputPackPath string, _ *configure.Config) error {
	stitch := "Elytra"

	inPath := filepath.Join(inputPathPath, data.GetCraftPath("wings"))
	outPath := filepath.Join(outputPackPath, data.GetCloniaPath("armor"))
	var err error

	fileName := "elytra.png"
	sourceImg, err := imaging.Open(filepath.Join(inPath, fileName))
	if err != nil {
		return openErrMsg(stitch, inPath, fileName)
	}
	sourceImg = imaging.Resize(sourceImg, 64, 32, imaging.Lanczos)

	dst := imaging.Paste(imaging.New(64, 32, color.Transparent),
		imaging.Rotate90(
			imaging.Crop(sourceImg, image.Rect(31, 0, 46, 22))),
		image.Point{0, 0})

	if err = imaging.Save(dst, filepath.Join(outPath, "mcl_armor_elytra.png")); err != nil {
		return saveErrMsg(stitch, outPath, "mcl_armor_elytra.png")
	}
	return nil
}
