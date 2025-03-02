package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io/fs"
	"log"
	"os"

	imaging "github.com/disintegration/imaging"
)

func ConvertPackMTG(inName string, outName string) {
	var textureErrorsLog string = fmt.Sprintf("%v %v\n", inName, nowShort)
	var successes = 0
	var failures = 0
	texturePackLocation := Config.InputDir + inName
	outPath := Config.OutputDir + outName
	if fs.ValidPath(outPath) {
		if err := os.Mkdir(outPath, 0755); err != nil {
			if errors.Is(err, fs.ErrInvalid) {
				log.Panicf("Folder %s is an \"invalid argument\". Maybe rename %s?\n", outPath, texturePackLocation)
			} else if errors.Is(err, fs.ErrPermission) {
				log.Panicf("Permission was denied. %s was not made.\n", outPath)
			} else if errors.Is(err, fs.ErrExist) {
				fmt.Printf("Folder %s already exists. Writing into it.\n", outPath)
			} else {
				fmt.Printf("How.\n")
				log.Panic(err)
			}
		}
	}

	for _, e := range mtgPaths {
		if err := os.MkdirAll(outPath+e, 0755); err != nil {
			log.Panic(err)
		}
	}

	if src, err := imaging.Open(texturePackLocation + "/pack.png"); err != nil {
		fmt.Println("Pack icon error~")
	} else {
		background := imaging.Fill(src, 350, 233, imaging.Center, imaging.Lanczos)
		background = imaging.Blur(background, 10)
		foreground := imaging.Resize(src, 233, 0, imaging.Lanczos)

		dst := imaging.New(350, 233, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, background, image.Pt(0, 0))
		dst = imaging.OverlayCenter(dst, foreground, 1.0)
		err = imaging.Save(dst, outPath+"/screenshot.png")
		if err != nil {
			fmt.Println("Failed to export the pack screenshot/icon!\n", err)
		}
	}

	copyTextureFails := []string{}

	for _, e := range minetestGameItems {
		if err := copyTextureAnimated(texturePackLocation+craftPaths[e.inPath]+e.inTexture, outPath+mtgPaths[e.outPath]+e.outTexture, e.framesAllowed); err != nil {
			copyTextureFails = append(copyTextureFails, e.inPath+"::"+e.inTexture+" failed to copy!")
		} else {
			successes += 1
		}
	}
	for _, e := range minetestGreenery {
		if err := greenify(e, texturePackLocation, outPath); err != nil {
			copyTextureFails = append(copyTextureFails, err.Error())
		} else {
			successes += 1
		}
	}

	if len(copyTextureFails) > 0 {
		//fmt.Printf("\n%v\n\n", &readWriteError{copyTextureFails, "normal textures"})
		textureErrorsLog += fmt.Sprintf("%v\n\n", &readWriteError{copyTextureFails, "normal textures"})
		failures += len(copyTextureFails)
	}

	compatibilityRating := (successes * 100) / (successes + failures)
	packConfigFile := fmt.Sprintf(`title = MTG + %s
name = %s
description = A Minecraft texture pack converted to Minetest Game. %d successes, %d failures, %d%% compatible, converted %v.`,
		inName+" for MTG", outName, successes, failures, compatibilityRating, nowShort)
	fmt.Printf("%s\n", packConfigFile)
	if err := os.WriteFile(outPath+"/texture_pack.conf", []byte(packConfigFile), 0644); err != nil {
		log.Panic(err)
	}

	if err := os.WriteFile(outPath+"/craft_to_clonia_errors_log.txt", []byte(textureErrorsLog), 0644); err != nil {
		log.Panic(err)
	}

}
