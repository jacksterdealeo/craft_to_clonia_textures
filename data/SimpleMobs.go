package data

var SimpleMobs = [...]SimpleConversion{
	// {"entity", "skeleton/bogged.png", "mobs_mc", ".png", 1},  // no match
	// {"entity", "skeleton/bogged_overlay.png", "mobs_mc", ".png", 1},  // no match
	{"entity", "skeleton/skeleton.png", "mobs_mc", "mobs_mc_skeleton.png", 1},
	{"entity", "skeleton/stray.png", "mobs_mc", "mobs_mc_stray.png", 1},
	{"entity", "skeleton/stray_overlay.png", "mobs_mc", "mobs_mc_stray_overlay.png", 1},
	{"entity", "skeleton/wither_skeleton.png", "mobs_mc", "mobs_mc_wither_skeleton.png", 1},

	// {"entity", "zombie/drowned.png", "mobs_mc", "", 1}, // no match
	// {"entity", "zombie/drowned_outer_layer.png", "mobs_mc", "", 1}, // no match
	{"entity", "zombie/husk.png", "mobs_mc", "mobs_mc_husk.png", 1},
	{"entity", "zombie/zombie.png", "mobs_mc", "mobs_mc_zombie.png", 1},

	// {"entity", ".png", "mobs_mc", ".png", 1},
}
