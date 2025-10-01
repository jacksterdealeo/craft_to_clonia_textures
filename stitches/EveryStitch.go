package stitches

var UniversalStitches = [...]func(string, string) error{
	RWAnvil,
	RWBedInv,
	RWCow,
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
	RWWater,

	// Assumed Mods
	RWCopperTools,
	RWTravelnet,
}

var CloniaStitches = [...]func(string, string) error{
	RWCloniaFlipFixes,
}
