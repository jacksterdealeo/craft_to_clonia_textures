package stitches

import (
	"image"
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

// Paintings in Mineclonia A.T.M. cannot be scaled.
func RWPaintings(inputPathPath, outputPackPath string, _ *configure.Config) error {
	stitch := "Paintings"

	inPath := filepath.Join(inputPathPath, data.GetCraftPath("painting"))
	outPath := filepath.Join(outputPackPath, data.GetCloniaPath("paintings"))

	type painting struct {
		paintingSourceFile string

		destX int
		destY int

		sizeX int
		sizeY int
	}
	var paintings = [...]painting{
		// 16x16
		{"alban.png", 16 * 0, 0, 16, 16},
		{"aztec.png", 16 * 1, 0, 16, 16},
		{"aztec2.png", 16 * 2, 0, 16, 16},
		{"meditative.png", 16 * 3, 0, 16, 16},
		{"wasteland.png", 16 * 4, 0, 16, 16},
		{"plant.png", 16 * 5, 0, 16, 16},
		{"bomb.png", 16 * 6, 0, 16, 16},
		// {"kebab.png", 16 * 7, 0, 16, 16}, // unused in game
		// 32x16
		{"sunset.png", 32 * 0, 32, 32, 16},
		{"creebet.png", 32 * 1, 32, 32, 16},
		{"sea.png", 32 * 2, 32, 32, 16},
		{"pool.png", 32 * 3, 32, 32, 16},
		{"courbet.png", 32 * 4, 32, 32, 16},
		// 16x32
		{"graham.png", 16 * 0, 64, 16, 32},
		{"wanderer.png", 16 * 1, 64, 16, 32},
		// {"prairie_ride.png", 16 * 2, 64, 16, 32}, // unused in game
		// 64x32
		{"passage.png", 64 * 0, 96, 64, 32},
		// {"lowmist.png", 64 * 1, 96, 64, 32}, // unused in game
		// {"finding.png", 64 * 2, 96, 64, 32}, // unused in game
		// {"fighters.png", 64 * 3, 96, 64, 32}, // no space // unused in game
		// {"changing.png", 64 * 4, 96, 64, 32}, // no space // unused in game
		// 32x32
		{"baroque.png", 32 * 0, 128, 32, 32},
		{"bust.png", 32 * 1, 128, 32, 32},
		{"match.png", 32 * 2, 128, 32, 32},
		{"skull_and_roses.png", 32 * 3, 128, 32, 32},
		{"stage.png", 32 * 4, 128, 32, 32},
		{"void.png", 32 * 5, 128, 32, 32},
		// ... There are 6 more of these in MC.
		// 64x64
		{"burning_skull.png", 64 * 0, 192, 64, 64},
		{"orb.png", 64 * 1, 192, 64, 64},
		{"pointer.png", 64 * 2, 192, 64, 64},
		// ... There are 2 more of these in MC.
		// 64x48
		{"donkey_kong.png", 192, 64, 64, 48},
		{"skeleton.png", 192, 112, 64, 48},
	}

	dst := imaging.New(256, 256, image.Transparent)
	for _, e := range paintings {
		src, err := imaging.Open(filepath.Join(inPath, e.paintingSourceFile))
		if err != nil {
			return openErrMsg(stitch, inPath, e.paintingSourceFile)
		}
		src = imaging.Resize(src, e.sizeX, e.sizeY, imaging.Lanczos)
		dst = imaging.Overlay(dst, src, image.Pt(e.destX, e.destY), 1.0)
	}
	//Backpeice
	src, err := imaging.Open(filepath.Join(inPath, "back.png"))
	if err != nil {
		return openErrMsg(stitch, inPath, "back.png")
	}
	src = imaging.Resize(src, 16, 16, imaging.Lanczos)
	tileX := 192
	tileY := 0
	for tileY < 64 {
		dst = imaging.Overlay(dst, src, image.Pt(tileX, tileY), 1.0)
		tileX += 16
		if tileX >= 256 {
			tileY += 16
			tileX = 192
		}
	}

	outFileName := "mcl_paintings_paintings.png"
	if err := imaging.Save(dst, filepath.Join(outPath, outFileName)); err != nil {
		return saveErrMsg(stitch, outPath, outFileName)
	}
	return nil
}
