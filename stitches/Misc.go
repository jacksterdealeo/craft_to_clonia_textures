package stitches

import (
	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWMisc(inputPackPath, outputPackPath string, _ *configure.Config) error {
	stitch := "Misc"

	craftPath := "block"
	cloniaPath := "vl"
	inPath := inputPackPath + data.CraftPaths[craftPath]
	outPath := outputPackPath + "/" + data.CloniaPaths[cloniaPath]

	inBlock := "sand.png"
	outBlock := "mcl_core_grey_sand.png"
	sand, err := imaging.Open(inPath + inBlock)
	if err != nil {
		return openErrMsg(stitch, craftPath, inBlock)
	}

	dst := imaging.Grayscale(sand)
	if saveErr := imaging.Save(dst, outPath+outBlock); saveErr != nil {
		return saveErrMsg(stitch, craftPath, outBlock)
	}

	return nil
}
