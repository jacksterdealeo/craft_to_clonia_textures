package data

import (
	"errors"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/mcmetahandler"
	"github.com/disintegration/imaging"
)

type GeneralConversion interface {
	Convert(string, string) error
	ReadPath() string
	SavePath() string
}

// Used for textures that should have no edits done.
// This is not to be used for most blocks, items, or anything that could possibly be animated by a custom pack.
// This just copies the file verbatim.
type StaticTexture struct {
	InPath     string
	InTexture  string
	OutPath    string
	OutTexture string
}

func (e *StaticTexture) Convert(inputPackLocation string, outputPackLocation string) error {
	err := copyTexture(
		filepath.Join(inputPackLocation, GetCraftPath(e.InPath), e.InTexture),
		filepath.Join(outputPackLocation, GetCloniaPath(e.OutPath), e.OutTexture),
	)
	if err != nil {
		return errors.New(err.Error() + " ~ " + e.InPath + "::" + e.InTexture)
	}
	return nil
}

// Returns the texture Minecraft inPath appended with texture name.
// This does not include the path to the pack itself.
func (e *StaticTexture) ReadPath() string {
	return filepath.Join(GetCraftPath(e.InPath), e.InTexture)
}

// Returns the texture Mineclonia outPath appended with texture name.
// This does not include the path to the pack itself.
func (e *StaticTexture) SavePath() string {
	return filepath.Join(GetCloniaPath(e.OutPath), e.OutTexture)
}

type SimpleTexture struct {
	Path    string
	Texture string
}

// Returns the texture path appended with texture name.
func (e *SimpleTexture) FullPath() string {
	return filepath.Join(CraftPaths[e.Path], e.Texture)
}

// Used for textures that could have animation data.
type SimpleConversion struct {
	InPath     string
	InTexture  string
	OutPath    string
	OutTexture string

	FramesAllowed int
}

// Returns the texture Minecraft inPath appended with texture name.
// This does not include the path to the pack itself.
func (e *SimpleConversion) ReadPath() string {
	return filepath.Join(GetCraftPath(e.InPath), e.InTexture)
}

// Returns the texture Mineclonia outPath appended with texture name.
// This does not include the path to the pack itself.
func (e *SimpleConversion) SavePath() string {
	return filepath.Join(GetCloniaPath(e.OutPath), e.OutTexture)
}

func (e *SimpleConversion) Convert(inputPackLocation string, outputPackLocation string) error {
	err := copyTextureAnimated(
		filepath.Join(inputPackLocation, GetCraftPath(e.InPath), e.InTexture),
		filepath.Join(outputPackLocation, GetCloniaPath(e.OutPath), e.OutTexture),
		e.FramesAllowed,
	)
	if err != nil {
		return errors.New(err.Error() + " ~ " + e.InPath + "::" + e.InTexture)
	}
	return nil
}

type ConversionWithFallback struct {
	In         [2]SimpleTexture
	OutPath    string
	OutTexture string

	FramesAllowed int
}

// Returns the texture Mineclonia outPath appended with texture name.
func (e *ConversionWithFallback) SavePath() string {
	return filepath.Join(CloniaPaths[e.OutPath], e.OutTexture)
}

type BasicArmorConversion struct {
	// This is in "humanoid" and "humanoid_leggings"
	InTexture string

	OutHelmet     string
	OutChestplate string
	OutLeggings   string
	OutBoots      string
}

type ArmorSetTextures struct {
	Helmet     *image.NRGBA
	Chestplate *image.NRGBA
	Leggings   *image.NRGBA
	Boots      *image.NRGBA
}

type MCArmorSet struct {
	MostParts *image.NRGBA
	Leggings  *image.NRGBA
}

// Used for hard-coded locations. This makes invalid paths easy for me to catch.
func GetCloniaPath(shortLocation string) string {
	path, ok := CloniaPaths[shortLocation]
	if !ok {
		log.Panic("Incorrect Location Value! ", shortLocation, "!")
	}
	return path
}

// Used for hard-coded locations. This makes invalid paths easy for me to catch.
func GetCraftPath(shortLocation string) string {
	path, ok := CraftPaths[shortLocation]
	if !ok {
		log.Panic("Incorrect Location Value! ", shortLocation, "!")
	}
	return path
}

// Copies a texture without making any changes.
func copyTexture(src string, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Copies a texture and can limit it's frames if required.
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
	frames, err := mcmetahandler.McmetaReader(src)
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
		return err
	}
	return nil
}
