package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	imaging "github.com/disintegration/imaging"
)

const (
	minetest_game = iota
	mineclonia
)

var (
	now      = time.Now().Format("01-02-2006 15:04:05")
	nowShort = time.Now().Format("2Jan06")
	version  = "Pre-release"
)

type readWriteError struct {
	files   []string
	message string
}

func (e *readWriteError) Error() string {
	return fmt.Sprintf("%s has %d fails:\n\t%v", e.message, len(e.files), strings.Join(e.files[:], "\n\t"))
}

func main() {
	fmt.Printf("Minecraft to Mineclonia Texture Pack Converter v%s\n\n", version)

	if config, err := loadJsonConfig(); err != nil {
		fmt.Println(err)
	} else {
		Config = config
	}

	if !Config.DefinedInput {
		Config.InputDir = "./input/"
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
	}

	if !Config.DefinedOutput {
		Config.OutputDir = "./output/"
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
	}

	var dir *os.File
	var err error
	if !Config.DefinedInput {
		dir, err = os.Open("./input/")
		if err != nil {
			log.Panic("Error:", err)
			return
		}
	} else {
		dir, err = os.Open(Config.InputDir)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				log.Println(Config.InputDir + "\n\nInput folder from config doesn't exist.")
			} else {
				log.Panic("Error:", err)
				return
			}
		}
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() &&
			filepath.Ext(file.Name()) == ".zip" {
			if _, err := os.Stat(Config.InputDir + "/" + FileNameWithoutExt(file.Name())); errors.Is(err, os.ErrNotExist) {
				fmt.Println("Unzipping:", file.Name())
				if err2 := unzipSource(Config.InputDir+"/"+file.Name(), Config.InputDir+"/"+FileNameWithoutExt(file.Name())); err2 != nil {
					fmt.Println("Extraction Error:", err)
				}
			} else {
				fmt.Println(file.Name(), "was already decompressed! :D")
			}
		}
	}

	dir, err = os.Open(Config.InputDir)
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
			if Config.ExportMineclonia {
				o := fmt.Sprintf("%s_mc_to_clonia", strings.ReplaceAll(strings.ToLower(file.Name()), " ", "_"))
				convertPackClonia(file.Name(), o)
			}
			if Config.ExportMinetest_Game {
				o := fmt.Sprintf("%s_mc_to_mtg", strings.ReplaceAll(strings.ToLower(file.Name()), " ", "_"))
				convertPackMTG(file.Name(), o)
			}

			fmt.Print("Done!\n\n")
		}
	}
}

// Copies over a texture file with no changes.
func copyTexture(src string, dest string) error {
	img, err := imaging.Open(src)
	if err != nil {
		return err
	}
	imgX := img.Bounds().Dx()
	imgY := img.Bounds().Dy()

	outImg := imaging.New(imgX, imgY, color.NRGBA{0, 0, 0, 0})
	outImg = imaging.Overlay(outImg, img, image.Point{0, 0}, 1.0)

	if err = imaging.Save(outImg, dest); err != nil {
		fmt.Println(src, "save failed!", err.Error())
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
	imgY := img.Bounds().Dy()
	maxNumOfFrames := imgY / imgX
	if maxNumOfFrames == 0 { // some 32x31 textures were causing problems.
		maxNumOfFrames = 1
	}
	if framesAllowed < maxNumOfFrames && framesAllowed >= 1 {
		maxNumOfFrames = framesAllowed
	}
	frames, err := McmetaReader(src)
	if err != nil || len(frames) == 0 {
		for i := range maxNumOfFrames {
			frames = append(frames, i)
		}
	}
	var outImgNumberOfFrames int
	if framesAllowed < 1 || framesAllowed > len(frames) {
		if len(frames) != 0 {
			outImgNumberOfFrames = len(frames)
		} else {
			outImgNumberOfFrames = maxNumOfFrames
		}
	} else {
		outImgNumberOfFrames = framesAllowed
	}
	outImg := imaging.New(imgX, imgX*outImgNumberOfFrames, color.NRGBA{0, 0, 0, 0})
	for i, e := range frames {
		frame := imaging.Crop(img, image.Rectangle{image.Point{0, e * imgX}, image.Point{imgX, (e * imgX) + imgX}})
		outImg = imaging.Overlay(outImg, frame, image.Point{0, i * imgX}, 1.0)
	}
	if err = imaging.Save(outImg, dest); err != nil {
		fmt.Println(src, "save failed!", err.Error())
		return err
	}
	return nil
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
