package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	"codeberg.org/ostech/craft_to_clonia_textures/stitches"
)

func convertPackMTG(inName string, outName string, config *configure.Config) {
	var textureErrorsLog string = fmt.Sprintf("%v %v\n", inName, nowShort)
	var successes = 0
	var failures = 0

	texturePackLocation := config.InputDir + inName
	outPath := filepath.Join(config.OutputDir, outName) + "/"
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

	for _, e := range data.MTPaths {
		if err := os.MkdirAll(outPath+e, 0755); err != nil {
			log.Panic(err)
		}
	}

	stitches.RWPackIcon(texturePackLocation, outPath, config)

	copyTextureFails := []string{}
	catchReadWriteErrors := func(err *readWriteError) {
		if err != nil {
			copyTextureFails = append(copyTextureFails, err.Error())
		} else {
			successes += 1
		}
	}

	for _, e := range data.MinetestGreenery {
		catchReadWriteErrors(mtg_greenify(e, texturePackLocation, outPath))
	}
	catchReadWriteErrors(mtg_obsidian_glass_fix(texturePackLocation, outPath))
	catchReadWriteErrors(mtg_grass_fix(texturePackLocation, outPath))

	for _, e := range data.MinetestGameItems {
		if err := copyTextureAnimated(texturePackLocation+craftPaths[e.InPath]+e.InTexture, outPath+data.MTPaths[e.OutPath]+e.OutTexture, e.FramesAllowed); err != nil {
			copyTextureFails = append(copyTextureFails, e.InPath+"::"+e.InTexture+" failed to copy!")
		} else {
			successes += 1
		}
	}
	catchReadWriteErrors(mtgLavaFix(texturePackLocation+craftPaths["block"], outPath))
	catchReadWriteErrors(mtgWaterFix(texturePackLocation+craftPaths["block"], outPath))

	if len(copyTextureFails) > 0 {
		//fmt.Printf("\n%v\n\n", &readWriteError{copyTextureFails, "normal textures"})
		textureErrorsLog += fmt.Sprintf("%v\n\n", &readWriteError{copyTextureFails, "normal textures"})
		failures += len(copyTextureFails)
	}

	compatibilityRating := (successes * 100) / (successes + failures)
	packConfigFile := fmt.Sprintf(`title = MTG %s
name = %s
description = MTG texture pack converted from Minecraft. %d successes, %d failures, %d%% compatible, converted %v.`,
		inName, outName, successes, failures, compatibilityRating, nowShort)
	fmt.Printf("%s\n", packConfigFile)
	if err := os.WriteFile(outPath+"/texture_pack.conf", []byte(packConfigFile), 0644); err != nil {
		log.Panic(err)
	}

	if err := os.WriteFile(outPath+"/craft_to_clonia_errors_log.txt", []byte(textureErrorsLog), 0644); err != nil {
		log.Panic(err)
	}

}
