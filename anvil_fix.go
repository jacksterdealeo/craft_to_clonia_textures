package main

import (
	"fmt"
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
	_ "image/png"
)

func anvil_fix(inPath string, outPath string) {
	abase, err := imaging.Open(inPath + "anvil.png")
	if err != nil {
		fmt.Println("AnvilBase error ~", "block, anvil.png")
		return
	}
	a0, err := imaging.Open(inPath + "anvil_top.png")
	if err != nil {
		fmt.Println("Anvil0 error ~")
		return
	}
	a1, err := imaging.Open(inPath + "chipped_anvil_top.png")
	if err != nil {
		fmt.Println("Anvil1 error ~")
		return
	}
	a2, err := imaging.Open(inPath + "damaged_anvil_top.png")
	if err != nil {
		fmt.Println("Anvil2 error ~")
		return
	}
	anvilX := abase.Bounds().Dx()
	anvilY := abase.Bounds().Dy()

	dst := imaging.New(anvilX, anvilY, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, abase, image.Pt(0, 0))
	dst = imaging.OverlayCenter(dst, a0, 1.0)

	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_0.png"); err != nil {
		fmt.Println("Anvil undamaged failed!")
	}
	dst = imaging.OverlayCenter(dst, a1, 1.0)
	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_1.png"); err != nil {
		fmt.Println("Anvil damaged1 failed!")
	}
	dst = imaging.OverlayCenter(dst, a2, 1.0)
	if err = imaging.Save(dst, outPath+"mcl_anvils_anvil_top_damaged_2.png"); err != nil {
		fmt.Println("Anvil damaged2 failed!")
	}
}
