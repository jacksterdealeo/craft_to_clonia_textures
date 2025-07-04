package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	data "codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

var (
	craftPaths  = data.CraftPaths
	cloniaPaths = data.CloniaPaths
)

func convertPackClonia(inName string, outName string) {
	var (
		successes          = 0
		failures           = 0
		inputPackLocation  = Config.InputDir + inName
		outputPackLocation = Config.OutputDir + outName
	)

	var textureErrorsLog = strings.Builder{}
	textureErrorsLog.WriteString(fmt.Sprintf("%v %v\n\n", inName, nowShort))

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

	for _, e := range data.CloniaPaths {
		if err := os.MkdirAll(outputPackLocation+e, 0755); err != nil {
			log.Panic(err)
		}
	}

	copyTextureFails := []string{}
	logCopyTextureAnimatedErrs := func(setsOfTextures ...[]data.SimpleConversion) {
		for _, set := range setsOfTextures {
			for _, texture := range set {
				err := copyTextureAnimated(
					inputPackLocation+data.CraftPaths[texture.InPath]+texture.InTexture,
					outputPackLocation+data.CloniaPaths[texture.OutPath]+texture.OutTexture,
					texture.FramesAllowed,
				)
				if err != nil {
					copyTextureFails = append(copyTextureFails, err.Error()+" ~ "+texture.InPath+"::"+texture.InTexture)
				} else {
					successes += 1
				}
			}
		}
	}

	logCopyTextureAnimatedErrs(
		data.SimpleItems[:],
		data.SimpleHUD[:],
	)

	for _, texture := range data.SimpleNoEdits {
		if err := copyTexture(
			inputPackLocation+craftPaths[texture.InPath]+texture.InTexture,
			outputPackLocation+cloniaPaths[texture.OutPath]+texture.OutTexture,
		); err != nil {
			copyTextureFails = append(copyTextureFails, err.Error()+" ~ "+texture.InPath+"::"+texture.InTexture)
		} else {
			successes += 1
		}
	}

	if len(copyTextureFails) > 0 {
		//fmt.Printf("\n%v\n\n", &readWriteError{copyTextureFails, "normal textures"})
		textureErrorsLog.WriteString(fmt.Sprintf("%v\n\n", &readWriteError{copyTextureFails, "normal textures"}))
		failures += len(copyTextureFails)
	}

	////special casses
	logRWErrs := func(e ...*readWriteError) {
		for _, error := range e {
			if error != nil {
				failures += len(error.files)
				textureErrorsLog.WriteString(fmt.Sprint(error.Error() + "\n\n"))
			}
		}
	}

	var stitches_error_log strings.Builder
	mcStitches(inputPackLocation, outputPackLocation, &stitches_error_log)
	textureErrorsLog.WriteString(stitches_error_log.String())
	textureErrorsLog.WriteString("\n")

	// I want to replace most or all the "fix" functions with "stitch" functions for better performace, and maintenance.
	// Any that are commented out are replaced already. I am leaving them fully listed until they are all finished.
	logRWErrs(
		//anvil_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["anvils"]),
		armor_fixes(inputPackLocation, outputPackLocation),
		campfire_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["campfires"]),
		crack_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["hud_base_textures"]),
		do_fixes(inputPackLocation, outputPackLocation),
		double_chests_fix(inputPackLocation+craftPaths["entity"]+"chest/", outputPackLocation+cloniaPaths["chests"]),
		flip_fix(inputPackLocation, outputPackLocation),
		//flowerpot_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["flowerpots"]),
		hud_fix(inputPackLocation, outputPackLocation),
		//lava_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["core"]),
		mods_fixes(inputPackLocation, outputPackLocation),
		single_chests_fix(inputPackLocation+craftPaths["entity"]+"chest/", outputPackLocation+cloniaPaths["chests"]),
		stonecutter_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["stonecutter"]),
		//water_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["core"]),

		MobConvert(inputPackLocation, outputPackLocation),
	)

	// Experience Bar
	if expProgress, err := imaging.Open(inputPackLocation + craftPaths["hud"] + "experience_bar_progress.png"); err != nil {
		textureErrorsLog.WriteString("Full Experience Bar failed. Couldn't Open \"experience_bar_progress.png\".\n\n")
		failures++
	} else {
		if err2 := imaging.Save(imaging.Rotate90(expProgress), outputPackLocation+cloniaPaths["experience"]+"mcl_experience_bar.png"); err2 != nil {
			textureErrorsLog.WriteString("Full Experience Bar failed. Couldn't Save \"mcl_experience_bar.png\".\n\n")
			failures++
		}
	}
	if expEmpty, err := imaging.Open(inputPackLocation + craftPaths["hud"] + "experience_bar_background.png"); err != nil {
		textureErrorsLog.WriteString("Empty Experience Bar failed. Couldn't Open \"experience_bar_background.png\".\n\n")
		failures++
	} else {
		if err2 := imaging.Save(imaging.Rotate90(expEmpty), outputPackLocation+cloniaPaths["experience"]+"mcl_experience_bar_background.png"); err2 != nil {
			textureErrorsLog.WriteString("Empty Experience Bar failed. Couldn't Save \"mcl_experience_bar_background.png\".\n\n")
			failures++
		}
	}
	func() {
		sc := [...]data.SimpleConversion{
			{"hud", "hotbar.png", "inventory", "mcl_inventory_hotbar.png", -1},
		}
		for _, e := range sc {
			err := copyTexture(inputPackLocation+craftPaths[e.InPath]+e.InTexture, outputPackLocation+cloniaPaths[e.OutPath]+e.OutTexture)
			if err != nil {
				textureErrorsLog.WriteString(e.OutTexture + " failed to convert.\n")
			}
		}
	}()

	compatibilityRating := (successes * 100) / (successes + failures)
	packConfigFile := fmt.Sprintf(`title = MC %s
name = %s
description = Mineclonia texture pack converted from Minecraft. %d successes, %d failures, %d%% compatible, converted %v.`,
		inName, outName, successes, failures, compatibilityRating, nowShort)
	fmt.Printf("%s\n", packConfigFile)
	if err := os.WriteFile(outputPackLocation+"/texture_pack.conf", []byte(packConfigFile), 0644); err != nil {
		log.Panic(err)
	}

	if err := os.WriteFile(outputPackLocation+"/craft_to_clonia_errors_log.txt", []byte(textureErrorsLog.String()), 0644); err != nil {
		log.Panic(err)
	}
}
