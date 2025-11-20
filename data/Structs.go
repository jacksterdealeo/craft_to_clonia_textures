package data

import (
	"image"
	"path/filepath"
)

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
func (e *SimpleConversion) ReadPath() string {
	return filepath.Join(CraftPaths[e.InPath], e.InTexture)
}

// Returns the texture Mineclonia outPath appended with texture name.
func (e *SimpleConversion) SavePath() string {
	return filepath.Join(CloniaPaths[e.OutPath], e.OutTexture)
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
