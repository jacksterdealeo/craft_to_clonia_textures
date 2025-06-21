package main

import (
	"fmt"
	"image"
	"image/color"

	imaging "github.com/disintegration/imaging"
)

// Copies over a texture file with no changes.
func copyTexture(src string, dest string) error {
	img, err := imaging.Open(src)
	if err != nil {
		return fmt.Errorf("Could not open!")
	}
	imgX := img.Bounds().Dx()
	imgY := img.Bounds().Dy()

	outImg := imaging.New(imgX, imgY, color.NRGBA{0, 0, 0, 0})
	outImg = imaging.Overlay(outImg, img, image.Point{0, 0}, 1.0)

	if err = imaging.Save(outImg, dest); err != nil {
		fmt.Println(src, "save failed!", err.Error())
		return fmt.Errorf("Could not save!")
	}
	return nil
}

// Copies over a texture file with animation frames intact.
// Set framesAllowed to less than 1 to copy the texture with all the frames.
func copyTextureAnimated(src string, dest string, framesAllowed int) error {
	img, err := imaging.Open(src)
	if err != nil {
		return fmt.Errorf("Could not open!")
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
		return fmt.Errorf("Could not save!")
	}
	return nil
}
