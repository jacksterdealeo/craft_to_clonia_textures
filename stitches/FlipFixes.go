package stitches

import (
	"log"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWCloniaFlipFixes(inputPackPath string, outputPackPath string) error {
	stitch := "FlipFixes"

	readFails := [][2]string{}
	writeFails := [][2]string{}
	flips := data.CloniaFlipsHorizontal

	for _, e := range flips {
		img, err := imaging.Open(inputPackPath + data.CraftPaths[e.InPath] + e.InTexture)
		log.Println(data.CraftPaths[e.InPath]+e.InTexture)
		if err != nil {
			readFails = append(readFails, [2]string{e.InPath, e.InTexture})
		} else {
			img = flipH(img)
			if err = imaging.Save(img, outputPackPath+data.CloniaPaths[e.OutPath]+e.OutTexture); err != nil {
				writeFails = append(readFails, [2]string{e.OutPath, e.OutTexture})
			}
		}
	}
	return multiErrMsg(stitch, readFails, writeFails)
}

