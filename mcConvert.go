package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io/fs"
	"log"
	"os"

	data "codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func convertPackClonia(inName string, outName string) {
	var (
		textureErrorsLog   = fmt.Sprintf("%v %v\n\n", inName, nowShort)
		successes          = 0
		failures           = 0
		inputPackLocation  = Config.InputDir + inName
		outputPackLocation = Config.OutputDir + outName
	)
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

	if src, err := imaging.Open(inputPackLocation + "/pack.png"); err != nil {
		fmt.Println("Pack icon error~")
	} else {
		background := imaging.Fill(src, 350, 233, imaging.Center, imaging.Lanczos)
		background = imaging.Blur(background, 10)
		foreground := imaging.Resize(src, 233, 0, imaging.Lanczos)

		dst := imaging.New(350, 233, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, background, image.Pt(0, 0))
		dst = imaging.OverlayCenter(dst, foreground, 1.0)
		err = imaging.Save(dst, outputPackLocation+"/screenshot.png")
		if err != nil {
			fmt.Println("Failed to export the pack screenshot/icon!\n", err)
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
		data.SimpleItems[:],
	)

	for _, texture := range copyAsIs {
		if err := copyTexture(
			inputPackLocation+craftPaths[texture.inPath]+texture.inTexture,
			outputPackLocation+cloniaPaths[texture.outPath]+texture.outTexture,
		); err != nil {
			copyTextureFails = append(copyTextureFails, err.Error()+" ~ "+texture.inPath+"::"+texture.inTexture)
		} else {
			successes += 1
		}
	}

	if len(copyTextureFails) > 0 {
		//fmt.Printf("\n%v\n\n", &readWriteError{copyTextureFails, "normal textures"})
		textureErrorsLog += fmt.Sprintf("%v\n\n", &readWriteError{copyTextureFails, "normal textures"})
		failures += len(copyTextureFails)
	}

	////special casses
	logRWErrs := func(e ...*readWriteError) {
		for _, error := range e {
			if error != nil {
				failures += len(error.files)
				textureErrorsLog += fmt.Sprint(error.Error() + "\n\n")
			}
		}
	}

	logRWErrs(
		anvil_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["anvils"]),
		armor_fixes(inputPackLocation, outputPackLocation),
		campfire_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["campfires"]),
		crack_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["hud_base_textures"]),
		do_fixes(inputPackLocation, outputPackLocation),
		double_chests_fix(inputPackLocation+craftPaths["entity"]+"chest/", outputPackLocation+cloniaPaths["chests"]),
		flip_fix(inputPackLocation, outputPackLocation),
		flowerpot_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["flowerpots"]),
		hud_fix(inputPackLocation, outputPackLocation),
		lava_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["core"]),
		mods_fixes(inputPackLocation, outputPackLocation),
		single_chests_fix(inputPackLocation+craftPaths["entity"]+"chest/", outputPackLocation+cloniaPaths["chests"]),
		stonecutter_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["stonecutter"]),
		water_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["core"]),
	)

	// Achivement Icon
	if src, err := imaging.Open(inputPackLocation + craftPaths["item"] + "writable_book.png"); err != nil {
		textureErrorsLog += "Achivement Icon Construction Failed. Couldn't Find \"writable_book.png\".\n\n"
		failures++
	} else {
		achivementIcon := imaging.Grayscale(src)
		if saveErr := imaging.Save(achivementIcon, outputPackLocation+cloniaPaths["achievements"]+"mcl_achievements_button.png"); saveErr != nil {
			textureErrorsLog += "Achivement Icon Construction Failed. Couldn't Save \"writable_book.png\".\n\n"
			failures++
		}
	}
	// Experience Bar
	if expProgress, err := imaging.Open(inputPackLocation + craftPaths["hud"] + "experience_bar_progress.png"); err != nil {
		textureErrorsLog += "Full Experience Bar failed. Couldn't Open \"experience_bar_progress.png\".\n\n"
		failures++
	} else {
		if err2 := imaging.Save(imaging.Rotate90(expProgress), outputPackLocation+cloniaPaths["experience"]+"mcl_experience_bar.png"); err2 != nil {
			textureErrorsLog += "Full Experience Bar failed. Couldn't Save \"mcl_experience_bar.png\".\n\n"
			failures++
		}
	}
	if expEmpty, err := imaging.Open(inputPackLocation + craftPaths["hud"] + "experience_bar_background.png"); err != nil {
		textureErrorsLog += "Empty Experience Bar failed. Couldn't Open \"experience_bar_background.png\".\n\n"
		failures++
	} else {
		if err2 := imaging.Save(imaging.Rotate90(expEmpty), outputPackLocation+cloniaPaths["experience"]+"mcl_experience_bar_background.png"); err2 != nil {
			textureErrorsLog += "Empty Experience Bar failed. Couldn't Save \"mcl_experience_bar_background.png\".\n\n"
			failures++
		}
	}
	func() {
		sc := [...]simpleConversion{
			{"hud", "hotbar.png", "inventory", "mcl_inventory_hotbar.png", -1},
		}
		for _, e := range sc {
			err := copyTexture(inputPackLocation+craftPaths[e.inPath]+e.inTexture, outputPackLocation+cloniaPaths[e.outPath]+e.outTexture)
			if err != nil {
				textureErrorsLog += (e.outTexture + " failed to convert.\n")
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

	if err := os.WriteFile(outputPackLocation+"/craft_to_clonia_errors_log.txt", []byte(textureErrorsLog), 0644); err != nil {
		log.Panic(err)
	}
}
