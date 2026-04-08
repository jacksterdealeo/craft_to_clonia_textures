package stitches

import "codeberg.org/ostech/craft_to_clonia_textures/configure"

var UniversalStitches = [...]func(string, string, *configure.Config) error{
	RWAnvil,
	RWBedInv,
	RWBigDripLeaf,
	RWCampfire,
	RWChiseledBooks,
	RWCobweb,
	RWCow,
	RWCrack,
	RWCrosshair,
	RWDeepslateTools,
	RWDoubleChests,
	RWElytra,
	RWEmbedded,
	RWFlowerPot,
	RWLava,
	RWLeatherArmor,
	RWMisc,
	RWPackIcon,
	RWPaintings,
	RWPig,
	RWPotionIndicator,
	RWRedstoneLamp,
	RWSigns,
	RWSingleChests,
	RWSmallDripLeaf,
	RWStonecutter,
	RWUI,
	RWWater,

	// Assumed Mods
	RWCopperTools,
	RWTravelnet,
}

var CloniaStitches = [...]func(string, string, *configure.Config) error{
	RWCloniaFlipFixes,
}
