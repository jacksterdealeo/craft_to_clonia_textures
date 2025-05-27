package data

import "image"

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
	return CraftPaths[e.InPath] + "/" + e.InTexture
}

// Returns the texture Mineclonia outPath appended with texture name.
func (e *SimpleConversion) SavePath() string {
	return CloniaPaths[e.OutPath] + "/" + e.OutTexture
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
