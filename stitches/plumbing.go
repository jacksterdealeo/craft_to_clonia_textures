package stitches

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/danibezoff/perspective-transform/perspective"
	"github.com/disintegration/imaging"
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
	return fmt.Errorf(strings.Join(msg, ""))
}

func flipHV(img image.Image) *image.NRGBA {
	return imaging.FlipH(imaging.FlipV(img))
}

func cropToScale(img image.Image, x1, y1, x2, y2, scale int) *image.NRGBA {
	return imaging.Crop(img, image.Rectangle{
		image.Point{x1 * scale, y1 * scale}, image.Point{x2 * scale, y2 * scale}})
}

func PerspectiveOverlay(
	dst, src image.NRGBA, topLeftX, topLeftY, topRightX, topRightY,
	botLeftX, botLeftY, botRightX, botRightY float64) (*image.NRGBA) {
	
	var result = imaging.New(dst.Rect.Dx(), dst.Rect.Dy(), color.Transparent)

	var srcPoints = [8]float64{
			0, 0, float64(src.Rect.Dx()), 0,
			float64(src.Rect.Dx()), float64(src.Rect.Dy()), 0, float64(src.Rect.Dy()),
	}

	var dstPoints = [8]float64{
			float64(topLeftX), float64(topLeftY),
			float64(topRightX), float64(topRightY),
			float64(botRightX), float64(botRightY),
			float64(botLeftX), float64(botLeftY),
	}

	p := perspective.New(srcPoints, dstPoints)

		for x := result.Bounds().Min.X; x < result.Bounds().Max.X; x++ {
			for y := result.Bounds().Min.Y; y < result.Bounds().Max.Y; y++ {
				srcX, srcY := p.TransformInv(float64(x), float64(y))
				c := src.At(int(math.Round(srcX)), int(math.Round(srcY)))
				result.Set(x, y, c)
			}
		}

	return imaging.Overlay(&dst, result, image.Pt(0, 0), 1.0)
}
