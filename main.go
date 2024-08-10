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
	for _, file := range files {
		fmt.Println(file.Name())
		outName := fmt.Sprintf("%s_mc_converted", strings.ReplaceAll(strings.ToLower(file.Name()), " ", "_"))
		ConvertPack("input/"+file.Name(), "output/"+outName)
	}

}

func ConvertPack(inName string, outName string) {
	if fs.ValidPath(outName) {
		if err := os.Mkdir(outName, 0755); err != nil {
			if errors.Is(err, fs.ErrInvalid) {
				log.Panicf("Folder %s is an \"invalid argument\". Maybe rename %s?\n", outName, inName)
			} else if errors.Is(err, fs.ErrPermission) {
				log.Panicf("Permission was denied. %s was not made.\n", outName)
			} else if errors.Is(err, fs.ErrExist) {
				fmt.Printf("Folder %s already exists. Writing into it.\n", outName)
			} else {
				fmt.Printf("How.\n")
				log.Panic(err)
			}
		}
	}

	packConfigFile := fmt.Sprintf(`title = %s
name = %s
description = A Minecraft texture pack converted to Mineclonia on %s.
author = Unknown
release = 01`, inName, outName, now)

	fmt.Printf("Pack info:\n%s\n", packConfigFile)
	err := os.WriteFile(outName+"/texture_pack.conf", []byte(packConfigFile), 0644)
	if err != nil {
		log.Panic(err)
	}

	//// pack icon
	//copyTexture(inName+"/pack.png", outName+"/screenshot.png")
	if src, err := imaging.Open(inName + "/pack.png"); err != nil {
		fmt.Println("Pack icon error~")
	} else {
		background := imaging.Fill(src, 350, 233, imaging.Center, imaging.Lanczos)
		background = imaging.Blur(background, 10)
		foreground := imaging.Resize(src, 233, 0, imaging.Lanczos)

		dst := imaging.New(350, 233, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, background, image.Pt(0, 0))
		dst = imaging.OverlayCenter(dst, foreground, 1.0)
		err = imaging.Save(dst, outName+"/screenshot.png")
	}

	craftPaths := CraftPaths()
	cloniaPaths := CloniaPaths()
	equivalents := basicITEMS()

	for _, v := range cloniaPaths {
		if err := os.MkdirAll(outName+v, 0755); err != nil {
			log.Panic(err)
		}
	}

	for _, e := range equivalents {
		copyTexture(inName+craftPaths[e[0]]+e[1], outName+cloniaPaths[e[2]]+e[3])
	}

	//special casses
	anvil_fix(inName+craftPaths["block"], outName+cloniaPaths["anvils"])
	water_fix(inName+craftPaths["block"], outName+cloniaPaths["core"])
	lava_fix(inName+craftPaths["block"], outName+cloniaPaths["core"])
	flip_fixes(inName, outName)
}

func copyTexture(src string, dest string) {
	source, err := os.Open(src)
	if err != nil {
		fmt.Printf(err.Error() + " ~ Copy for texture skipped.\n")
		return
	}
	defer source.Close()
	destination, err := os.Create(dest)
	if err != nil {
		fmt.Printf(err.Error() + " ~ Copy for texture skipped.\n")
		return
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err)
	}
}
