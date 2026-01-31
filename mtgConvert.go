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

	inputPackLocation := config.InputDir + inName
	outputPackLocation := filepath.Join(config.OutputDir, outName) + "/"
	if fs.ValidPath(outputPackLocation) {
		if err := os.Mkdir(outputPackLocation, 0755); err != nil {
			if errors.Is(err, fs.ErrInvalid) {
				log.Panicf("Folder %s is an \"invalid argument\". Maybe rename %s?\n", outputPackLocation, inputPackLocation)
			} else if errors.Is(err, fs.ErrPermission) {
				log.Panicf("Permission was denied. %s was not made.\n", outputPackLocation)
			} else if errors.Is(err, fs.ErrExist) {
				fmt.Printf("Folder %s already exists. Writing into it.\n", outputPackLocation)
			} else {
				fmt.Printf("How.\n")
				log.Panic(err)
			}
		}
	}

	for _, e := range data.MTPaths {
		if err := os.MkdirAll(outputPackLocation+e, 0755); err != nil {
			log.Panic(err)
		}
	}

	stitches.RWPackIcon(inputPackLocation, outputPackLocation, config)

	copyTextureFails := []string{}
	catchReadWriteErrors := func(err *readWriteError) {
		if err != nil {
			copyTextureFails = append(copyTextureFails, err.Error())
		} else {
			successes += 1
		}
	}

	for _, e := range data.MinetestGreenery {
		catchReadWriteErrors(mtg_greenify(e, inputPackLocation, outputPackLocation))
	}
	catchReadWriteErrors(mtg_obsidian_glass_fix(inputPackLocation, outputPackLocation))
	catchReadWriteErrors(mtg_grass_fix(inputPackLocation, outputPackLocation))

	for _, e := range data.MinetestGameItems {
		//if err := copyTextureAnimated(inputPackLocation+craftPaths[e.InPath]+e.InTexture, outputPackLocation+data.MTPaths[e.OutPath]+e.OutTexture, e.FramesAllowed); err != nil {
		if err := e.Convert(inputPackLocation, outputPackLocation); err != nil {
			copyTextureFails = append(copyTextureFails, e.InPath+"::"+e.InTexture+" failed to copy!")
		} else {
			successes += 1
		}
	}
	catchReadWriteErrors(mtgLavaFix(inputPackLocation, craftPaths["block"], outputPackLocation))
	catchReadWriteErrors(mtgWaterFix(inputPackLocation+craftPaths["block"], outputPackLocation))

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
	if err := os.WriteFile(outputPackLocation+"/texture_pack.conf", []byte(packConfigFile), 0644); err != nil {
		log.Panic(err)
	}

	if err := os.WriteFile(outputPackLocation+"/craft_to_clonia_errors_log.txt", []byte(textureErrorsLog), 0644); err != nil {
		log.Panic(err)
	}

}
