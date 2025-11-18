package stitches

import "codeberg.org/ostech/craft_to_clonia_textures/configure"

var UniversalStitches = [...]func(string, string, *configure.Config) error{
	RWAnvil,
	RWBedInv,
	RWCow,
	RWCrack,
	RWDeepslateTools,
	RWDoubleChests,
	RWEmbedded,
	RWFlowerPot,
	RWLava,
	RWLeatherArmor,
	RWMisc,
	RWPackIcon,
	RWPig,
	RWPotionIndicator,
	RWRedstoneLamp,
	RWSigns,
	RWSingleChests,
	RWUI,
	RWWater,

	// Assumed Mods
	RWCopperTools,
	RWTravelnet,
}

var CloniaStitches = [...]func(string, string, *configure.Config) error{
	RWCloniaFlipFixes,
}
