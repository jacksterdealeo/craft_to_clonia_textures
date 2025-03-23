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

func convertPackClonia(inName string, outName string) {
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

	for _, e := range cloniaPaths {
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

	for _, e := range basicItems {
		if err := copyTextureAnimated(texturePackLocation+craftPaths[e.inPath]+e.inTexture, outPath+cloniaPaths[e.outPath]+e.outTexture, e.framesAllowed); err != nil {
			copyTextureFails = append(copyTextureFails, e.inPath+"::"+e.inTexture+" failed to copy!")
		} else {
			successes += 1
		}
	}

	for _, e := range basicHUD {
		if err := copyTextureAnimated(texturePackLocation+craftPaths[e.inPath]+e.inTexture, outPath+cloniaPaths[e.outPath]+e.outTexture, e.framesAllowed); err != nil {
			copyTextureFails = append(copyTextureFails, e.inPath+"::"+e.inTexture+" failed to copy!")
		} else {
			successes += 1
		}
	}

	for _, e := range copyAsIs {
		if err := copyTexture(texturePackLocation+craftPaths[e.inPath]+e.inTexture, outPath+cloniaPaths[e.outPath]+e.outTexture); err != nil {
			copyTextureFails = append(copyTextureFails, e.inPath+"::"+e.inTexture+" failed to copy!")
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
	logRWErr := func(e *readWriteError) {
		if e != nil {
			failures += len(e.files)
			textureErrorsLog += fmt.Sprint(e.Error() + "\n\n")
		}
	}

	double_chests_fix(texturePackLocation+craftPaths["entity"]+"chest/", outPath+cloniaPaths["chests"])
	water_fix(texturePackLocation+craftPaths["block"], outPath+cloniaPaths["core"])
	do_fixes(texturePackLocation, outPath)
	logRWErr(anvil_fix(texturePackLocation+craftPaths["block"], outPath+cloniaPaths["anvils"]))
	logRWErr(campfire_fix(texturePackLocation+craftPaths["block"], outPath+cloniaPaths["campfires"]))
	logRWErr(crack_fix(texturePackLocation+craftPaths["block"], outPath+cloniaPaths["hud_base_textures"]))
	logRWErr(flip_fix(texturePackLocation, outPath))
	logRWErr(flowerpot_fix(texturePackLocation+craftPaths["block"], outPath+cloniaPaths["flowerpots"]))
	logRWErr(hud_fix(texturePackLocation, outPath))
	logRWErr(lava_fix(texturePackLocation+craftPaths["block"], outPath+cloniaPaths["core"]))
	logRWErr(single_chests_fix(texturePackLocation+craftPaths["entity"]+"chest/", outPath+cloniaPaths["chests"]))
	logRWErr(stonecutter_fix(texturePackLocation+craftPaths["block"], outPath+cloniaPaths["stonecutter"]))
	// Achivement Icon
	if src, err := imaging.Open(texturePackLocation + craftPaths["item"] + "writable_book.png"); err != nil {
		textureErrorsLog += "Achivement Icon Construction Failed. Couldn't Find \"writable_book.png\".\n\n"
		failures++
	} else {
		achivementIcon := imaging.Grayscale(src)
		if saveErr := imaging.Save(achivementIcon, outPath+cloniaPaths["achievements"]+"mcl_achievements_button.png"); saveErr != nil {
			textureErrorsLog += "Achivement Icon Construction Failed. Couldn't Save \"writable_book.png\".\n\n"
			failures++
		}
	}
	// Experience Bar
	if expProgress, err := imaging.Open(texturePackLocation + craftPaths["hud"] + "experience_bar_progress.png"); err != nil {
		textureErrorsLog += "Full Experience Bar failed. Couldn't Open \"experience_bar_progress.png\".\n\n"
		failures++
	} else {
		if err2 := imaging.Save(imaging.Rotate90(expProgress), outPath+cloniaPaths["experience"]+"mcl_experience_bar.png"); err2 != nil {
			textureErrorsLog += "Full Experience Bar failed. Couldn't Save \"mcl_experience_bar.png\".\n\n"
			failures++
		}
	}
	if expEmpty, err := imaging.Open(texturePackLocation + craftPaths["hud"] + "experience_bar_background.png"); err != nil {
		textureErrorsLog += "Empty Experience Bar failed. Couldn't Open \"experience_bar_background.png\".\n\n"
		failures++
	} else {
		if err2 := imaging.Save(imaging.Rotate90(expEmpty), outPath+cloniaPaths["experience"]+"mcl_experience_bar_background.png"); err2 != nil {
			textureErrorsLog += "Empty Experience Bar failed. Couldn't Save \"mcl_experience_bar_background.png\".\n\n"
			failures++
		}
	}
	func() {
		sc := [...]simpleConversion{
			{"hud", "hotbar.png", "inventory", "mcl_inventory_hotbar.png", -1},
		}
		for _, e := range sc {
			err := copyTexture(texturePackLocation+craftPaths[e.inPath]+e.inTexture, outPath+cloniaPaths[e.outPath]+e.outTexture)
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
	if err := os.WriteFile(outPath+"/texture_pack.conf", []byte(packConfigFile), 0644); err != nil {
		log.Panic(err)
	}

	if err := os.WriteFile(outPath+"/craft_to_clonia_errors_log.txt", []byte(textureErrorsLog), 0644); err != nil {
		log.Panic(err)
	}
}
