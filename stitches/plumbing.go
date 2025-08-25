package stitches

import (
	"fmt"
	"image"
	"strings"

	imaging "github.com/disintegration/imaging"
)

var (
	flipH = imaging.FlipH
	flipV = imaging.FlipV
)

// Returns a formatted error for opening or reading a single file.
func openErrMsg(stitch, dir, file string) error {
	return fmt.Errorf("%v Stitch > \tCould not open! ~ %v::%v", stitch, dir, file)
}

// Returns a formatted error for saving a single file.
func saveErrMsg(stitch, dir, file string) error {
	return fmt.Errorf("%v Stitch > \tCould not save! ~ %v::%v", stitch, dir, file)
}

// Returns a formatted error for multiple files.
// Untested!!! Please check empty pack!!!
func multiErrMsg(stitch string, readDirsAndFiles [][2]string, writeDirsAndFiles [][2]string) error {
	var msg []string
	for _, e := range readDirsAndFiles {
		msg = append(msg, fmt.Sprintf("%v Stitches > \tCould not open! ~ %v::%v\n", stitch, e[0], e[1]))
	}
	for _, e := range writeDirsAndFiles {
		msg = append(msg, fmt.Sprintf("%v Stitches > \tCould not save! ~ %v::%v\n", stitch, e[0], e[1]))
	}
	if len(msg) == 0 {
		return nil
	}
	return fmt.Errorf(strings.Join(msg, "\n"))
}

func flipHV(img image.Image) *image.NRGBA {
	return imaging.FlipH(imaging.FlipV(img))
}

func cropToScale(img image.Image, x1, y1, x2, y2, scale int) *image.NRGBA {
	return imaging.Crop(img, image.Rectangle{
		image.Point{x1 * scale, y1 * scale}, image.Point{x2 * scale, y2 * scale}})
}
