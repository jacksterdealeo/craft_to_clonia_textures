package main

import (
	"errors"
	"fmt"
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const ()

var (
	now     = time.Now().Format("01-02-2006 15:04:05")
	version = "0.2.0"
)

type readWriteError struct {
	files   []string
	message string
}

func (e *readWriteError) Error() string {
	return fmt.Sprintf("%v - %s", e.files, e.message)
}

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
		log.Panic("Error:", err)
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
		if !file.IsDir() &&
			filepath.Ext(file.Name()) == ".zip" {
			if _, err := os.Stat("input/" + FileNameWithoutExt(file.Name())); errors.Is(err, os.ErrNotExist) {
				fmt.Println("Unzipping:", file.Name())
				if unzipSource("input/"+file.Name(), "input/"+FileNameWithoutExt(file.Name())); err != nil {
					fmt.Println("Extraction Error:", err)
				}
			} else {
				fmt.Println(file.Name(), "was already decompressed! :D")
			}
		}
	}

	dir, err = os.Open("./input/")
	if err != nil {
		log.Panic("Error:", err)
		return
	}
	defer dir.Close()
	files, err = dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
			outPath := fmt.Sprintf("%s_mc_converted", strings.ReplaceAll(strings.ToLower(file.Name()), " ", "_"))
			ConvertPack(file.Name(), outPath)
			fmt.Println()
		}
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

	for _, e := range cloniaPaths {
		if err := os.MkdirAll(outPath+e, 0755); err != nil {
			log.Panic(err)
		}
	}

	copyTextureFails := []string{}
	for _, e := range basicItems {
		err := copyTexture(inPath+craftPaths[e[0]]+e[1], outPath+cloniaPaths[e[2]]+e[3])
		if err != nil {
			copyTextureFails = append(copyTextureFails, e[0]+"::"+e[1])
		}
	}
	if len(copyTextureFails) != 0 {
		fmt.Println("\nThe following textures couldn't be copied:", copyTextureFails)
	}

	//special casses
	if err := anvil_fix(inPath+craftPaths["block"], outPath+cloniaPaths["anvils"]); err != nil {
		fmt.Println(err.Error())
	}
	chests_fix(inPath+craftPaths["entity"]+"chest/", outPath+cloniaPaths["chests"])
	water_fix(inPath+craftPaths["block"], outPath+cloniaPaths["core"])
	lava_fix(inPath+craftPaths["block"], outPath+cloniaPaths["core"])
	flowerpot_fix(inPath+craftPaths["block"], outPath+cloniaPaths["flowerpots"])
	flip_fix(inPath, outPath)
	if err := animated_texture_fix(inPath, outPath); err != nil {
		fmt.Println(err.Error())
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

}

func copyTexture(src string, dest string) error {
	img, err := imaging.Open(src)
	if err != nil {
		return err
	}
	imgX := img.Bounds().Dx()
	//imgY := img.Bounds().Dy()

	outImg := imaging.New(imgX, imgX, color.NRGBA{0, 0, 0, 0})
	outImg = imaging.Overlay(outImg, img, image.Point{0, 0}, 1.0)

	if err = imaging.Save(outImg, dest); err != nil {
		fmt.Println(src, "save failed!")
		return err
	}
	return nil
}

// Copies over a texture file with animation frames intact.
// Set framesAllowed to less than 1 to copy the texture with all the frames.
func copyTextureAnimated(src string, dest string, framesAllowed int) error {
	img, err := imaging.Open(src)
	if err != nil {
		return err
	}
	imgX := img.Bounds().Dx()
	frames, err := McmetaReader(src)
	if err != nil {
		return err
	}
	var outImgNumberOfFrames int
	if framesAllowed <= 0 || framesAllowed > len(frames) {
		outImgNumberOfFrames = len(frames)
	} else {
		outImgNumberOfFrames = framesAllowed
	}
	outImg := imaging.New(imgX, imgX*outImgNumberOfFrames, color.NRGBA{0, 0, 0, 0})
	for i, e := range frames {
		frame := imaging.Crop(img, image.Rectangle{image.Point{0, e * imgX}, image.Point{imgX, (e * imgX) + imgX}})
		outImg = imaging.Overlay(outImg, frame, image.Point{0, i * imgX}, 1.0)
	}
	if err = imaging.Save(outImg, dest); err != nil {
		fmt.Println(src, "save failed!")
		return err
	}
	return nil
}

// Don't use pls.
func copyTextureBasic(src string, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		//fmt.Printf(err.Error() + " ~ Copy for texture skipped.\n")
		return err
	}
	defer source.Close()
	destination, err := os.Create(dest)
	if err != nil {
		//fmt.Printf(err.Error() + " ~ Copy for texture skipped.\n")
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err)
	}
	return nil
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
