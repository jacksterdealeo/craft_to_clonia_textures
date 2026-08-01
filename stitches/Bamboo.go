package stitches

import (
	"fmt"
	"image"
	"image/color"
	"path/filepath"
	"strings"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWBamboo(inputPackPath, outputPackPath string, _ *configure.Config) error {
	inPath := filepath.Join(inputPackPath, data.GetCraftPath("block"))
	outPath := filepath.Join(outputPackPath, data.GetCloniaPath("bamboo"))
	failures := strings.Builder{}

	stitch := "Bamboo"

	outName := "mcl_bamboo_bamboo.png"
	inName := "bamboo_stalk.png"
	bambooImg, err := imaging.Open(filepath.Join(inPath, inName))
	if err != nil {
		fmt.Fprintf(&failures, "%v Stitch > \tCould not open! ~ %v::%v\n", stitch, inPath, inName)
	} else {
		inName = "bamboo_stage0.png"
		bambooSproutImg, err := imaging.Open(filepath.Join(inPath, inName))
		if err != nil {
			fmt.Fprintf(&failures, "%v Stitch > \tCould not open! ~ %v::%v\n", stitch, inPath, inName)
		} else {
			out := imaging.New(bambooImg.Bounds().Dx(), bambooImg.Bounds().Dy(), color.Transparent)
			stalk := imaging.Crop(bambooImg, image.Rect(0, 0, (bambooImg.Bounds().Dx()*3)/16, bambooImg.Bounds().Dy()))
			stalkTop := imaging.Crop(bambooImg, image.Rect(
				bambooImg.Bounds().Dx()-(bambooImg.Bounds().Dx()*3)/16,
				0,
				bambooImg.Bounds().Dx(),
				(bambooImg.Bounds().Dy()*3)/16))
			sprout := imaging.Crop(bambooSproutImg, image.Rect(
				(bambooSproutImg.Bounds().Dx()*5)/16,
				0,
				bambooSproutImg.Bounds().Dx()-((bambooSproutImg.Bounds().Dx()*5)/16),
				bambooSproutImg.Bounds().Dy()))

			out = imaging.Overlay(out, stalk, image.Pt(0, 0), 1)
			out = imaging.Overlay(out, stalkTop, image.Pt((out.Bounds().Dx()*3)/16, 0), 1)
			out = imaging.Overlay(out, sprout, image.Pt((out.Bounds().Dx()*6)/16, 0), 1)
			if err := imaging.Save(out, filepath.Join(outPath, outName)); err != nil {
				fmt.Fprintf(&failures, "%v Stitch > \tCould not save! ~ %v::%v\n", stitch, outPath, outName)
			}
		}
	}

	// inName := "mcl_bamboo_endcap.png" // Voxelibre only (and kinda bad)
	// inName := "mcl_bamboo_bamboo_fpm.png" // difficult

	if failures.Len() == 0 {
		return nil
	}
	return fmt.Errorf("%s", failures.String())
}
