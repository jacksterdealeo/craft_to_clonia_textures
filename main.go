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
	version = "Pre-release"
)

type readWriteError struct {
	files   []string
	message string
}

func (e *readWriteError) Error() string {
	return fmt.Sprintf("%s has %d fails: %v", e.message, len(e.files), e.files)
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
	var textureErrorsLog string
	var successes = 0
	var failures = 0
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
			copyTextureFails = append(copyTextureFails, e[0]+"::"+e[1]+" failed to copy!")
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
	if err := animated_texture_fix(inPath, outPath); err != nil {
		failures += len(err.files)
		textureErrorsLog += fmt.Sprint(err.Error() + "\n\n")
	}
	if err := anvil_fix(inPath+craftPaths["block"], outPath+cloniaPaths["anvils"]); err != nil {
		failures += len(err.files)
		textureErrorsLog += fmt.Sprint(err.Error() + "\n\n")
	}
	if err := double_chests_fix(inPath+craftPaths["entity"]+"chest/", outPath+cloniaPaths["chests"]); err != nil {
		failures += len(err.files)
		textureErrorsLog += fmt.Sprint(err.Error() + "\n\n")
	}
	if err := flip_fix(inPath, outPath); err != nil {
		failures += len(err.files)
		textureErrorsLog += fmt.Sprint(err.Error() + "\n\n")
	}
	if err := flowerpot_fix(inPath+craftPaths["block"], outPath+cloniaPaths["flowerpots"]); err != nil {
		failures += len(err.files)
		textureErrorsLog += fmt.Sprint(err.Error() + "\n\n")
	}
	if err := lava_fix(inPath+craftPaths["block"], outPath+cloniaPaths["core"]); err != nil {
		failures += len(err.files)
		textureErrorsLog += fmt.Sprint(err.Error() + "\n\n")
	}
	if err := single_chests_fix(inPath+craftPaths["entity"]+"chest/", outPath+cloniaPaths["chests"]); err != nil {
		failures += len(err.files)
		textureErrorsLog += fmt.Sprint(err.Error() + "\n\n")
	}
	if err := water_fix(inPath+craftPaths["block"], outPath+cloniaPaths["core"]); err != nil {
		failures += len(err.files)
		textureErrorsLog += fmt.Sprint(err.Error() + "\n\n")
	}

	compatibilityRating := (successes * 100) / (successes + failures)
	packConfigFile := fmt.Sprintf(`title = %s
name = %s
description = A Minecraft texture pack converted to Mineclonia on %s. %d successes, %d failures, %d%% compatible.`,
		inName, outName, now, successes, failures, compatibilityRating)
	fmt.Printf("Pack info:\n%s\n", packConfigFile)
	if err := os.WriteFile(outPath+"/texture_pack.conf", []byte(packConfigFile), 0644); err != nil {
		log.Panic(err)
	}

	if err := os.WriteFile(outPath+"/craft_to_clonia_errors_log.txt", []byte(textureErrorsLog), 0644); err != nil {
		log.Panic(err)
	}
}

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
	if framesAllowed < maxNumOfFrames && framesAllowed >= 1 {
		maxNumOfFrames = framesAllowed
	}
	frames, err := McmetaReader(src)
	if err != nil {
		return err
	}
	if len(frames) == 0 {
		for i := 0; i < maxNumOfFrames; i++ {
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
