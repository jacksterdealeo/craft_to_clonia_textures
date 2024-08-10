package main

import (
	"errors"
	"fmt"
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
	_ "image/png"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

const ()

var (
	now     = time.Now().Format("01-02-2006 15:04:05")
	version = "0.0.0"
)

func main() {
	fmt.Printf("Minecraft to Mineclonia Texture Pack Converter v%s\n", version)

	if fs.ValidPath("output") {
		if err := os.Mkdir("output", 0755); err != nil {
			if errors.Is(err, fs.ErrPermission) {
				log.Panicf("Permission was denied. %s was not made.\n", "output")
			} else if errors.Is(err, fs.ErrExist) {
				fmt.Printf("Folder %s already exists.\n", "output")
			} else {
				fmt.Printf("How.\n")
				log.Panic(err)
			}
		} else {
			fmt.Println("Made the output folder!")
		}
	}

	if fs.ValidPath("input") {
		if err := os.Mkdir("input", 0755); err != nil {
			if errors.Is(err, fs.ErrPermission) {
				log.Panicf("Permission was denied. %s was not made.\n", "input")
			} else if errors.Is(err, fs.ErrExist) {
				fmt.Printf("Folder %s already exists.\n", "input")
			} else {
				fmt.Printf("How.\n")
				log.Panic(err)
			}
		} else {
			fmt.Println("Made the input folder!")
		}
	}

	dir, err := os.Open("./input/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	for _, file := range files {
		fmt.Println(file.Name())
		outPath := fmt.Sprintf("%s_mc_converted", strings.ReplaceAll(strings.ToLower(file.Name()), " ", "_"))
		ConvertPack(file.Name(), outPath)
		fmt.Println()
	}
}

func ConvertPack(inName string, outName string) {
	inPath := "input/" + inName
	outPath := "output/" + outName
	if fs.ValidPath(outPath) {
		if err := os.Mkdir(outPath, 0755); err != nil {
			if errors.Is(err, fs.ErrInvalid) {
				log.Panicf("Folder %s is an \"invalid argument\". Maybe rename %s?\n", outPath, inPath)
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

	packConfigFile := fmt.Sprintf(`title = %s
name = %s
description = A Minecraft texture pack converted to Mineclonia on %s.`,
		inName, outName, now)

	fmt.Printf("Pack info:\n%s\n", packConfigFile)
	err := os.WriteFile(outPath+"/texture_pack.conf", []byte(packConfigFile), 0644)
	if err != nil {
		log.Panic(err)
	}

	//// pack icon
	//copyTexture(inPath+"/pack.png", outPath+"/screenshot.png")
	if src, err := imaging.Open(inPath + "/pack.png"); err != nil {
		fmt.Println("Pack icon error~")
	} else {
		background := imaging.Fill(src, 350, 233, imaging.Center, imaging.Lanczos)
		background = imaging.Blur(background, 10)
		foreground := imaging.Resize(src, 233, 0, imaging.Lanczos)

		dst := imaging.New(350, 233, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, background, image.Pt(0, 0))
		dst = imaging.OverlayCenter(dst, foreground, 1.0)
		err = imaging.Save(dst, outPath+"/screenshot.png")
	}

	craftPaths := CraftPaths()
	cloniaPaths := CloniaPaths()
	equivalents := basicITEMS()

	for _, e := range cloniaPaths {
		if err := os.MkdirAll(outPath+e, 0755); err != nil {
			log.Panic(err)
		}
	}

	copyTextureFails := []string{}
	for _, e := range equivalents {
		if !copyTexture(inPath+craftPaths[e[0]]+e[1], outPath+cloniaPaths[e[2]]+e[3]) {
			copyTextureFails = append(copyTextureFails, e[0]+e[1])
		}
	}
	if len(copyTextureFails) != 0 {
		fmt.Println("\nThe following textures couldn't be copied:", copyTextureFails)
	}

	//special casses
	anvil_fix(inPath+craftPaths["block"], outPath+cloniaPaths["anvils"])
	water_fix(inPath+craftPaths["block"], outPath+cloniaPaths["core"])
	lava_fix(inPath+craftPaths["block"], outPath+cloniaPaths["core"])
	flip_fixes(inPath, outPath)
}

func copyTexture(src string, dest string) bool {
	source, err := os.Open(src)
	if err != nil {
		//fmt.Printf(err.Error() + " ~ Copy for texture skipped.\n")
		return false
	}
	defer source.Close()
	destination, err := os.Create(dest)
	if err != nil {
		//fmt.Printf(err.Error() + " ~ Copy for texture skipped.\n")
		return false
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err)
	}
	return true
}
