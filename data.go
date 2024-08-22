package main

import ()

func CraftPaths() map[string]string {
	craftPaths := map[string]string{
		"armor":    "/assets/minecraft/textures/models/armor/",
		"bed":      "/assets/minecraft/textures/entity/bed/",
		"block":    "/assets/minecraft/textures/block/",
		"entity":   "/assets/minecraft/textures/entity/",
		"item":     "/assets/minecraft/textures/item/",
		"misc":     "/assets/minecraft/textures/misc/",
		"particle": "/assets/minecraft/textures/particle/",
	}
	return craftPaths
}

func CloniaPaths() map[string]string {
	cloniaPaths := map[string]string{
		"amethyst":        "/ITEMS/mcl_amethyst/textures/",
		"anvils":          "/ITEMS/mcl_anvils/textures/",
		"armor_stand":     "/ITEMS/mcl_armor_stand/textures/",
		"bamboo":          "/ITEMS/mcl_bamboo/textures/",
		"barrels":         "/ITEMS/mcl_barrels/textures/",
		"beds":            "/ITEMS/mcl_beds/textures/",
		"beehives":        "/ITEMS/mcl_beehives/textures/",
		"blackstone":      "/ITEMS/mcl_blackstone/textures/",
		"blast_furnace":   "/ITEMS/mcl_blast_furnace/textures/",
		"bone_meal":       "/ITEMS/mcl_bone_meal/textures/",
		"books":           "/ITEMS/mcl_books/textures/",
		"cherry_blossom":  "/ITEMS/mcl_cherry_blossom/textures/",
		"copper":          "/ITEMS/mcl_copper/textures/",
		"core":            "/ITEMS/mcl_core/textures/",
		"crafting_table":  "/ITEMS/mcl_crafting_table/textures/",
		"crimson":         "/ITEMS/mcl_crimson/textures/",
		"deepslate":       "/ITEMS/mcl_deepslate/textures/",
		"doors":           "/ITEMS/mcl_doors/textures/",
		"farming":         "/ITEMS/mcl_farming/textures/",
		"fishing":         "/ITEMS/mcl_fishing/textures/",
		"fletching_table": "/ITEMS/mcl_fletching_table/textures/",
		"flowerpots":      "/ITEMS/mcl_flowerpots/textures/",
		"flowers":         "/ITEMS/mcl_flowers/textures/",
		"furnaces":        "/ITEMS/mcl_furnaces/textures/",
		"honey":           "/ITEMS/mcl_honey/textures/",
		"hopper":          "/ITEMS/mcl_hoppers/textures/",
		"mangrove":        "/ITEMS/mcl_mangrove/textures/",
		"mud":             "/ITEMS/mcl_mud/textures/",
		"mushrooms":       "/ITEMS/mcl_mushrooms/textures/",
		"nether":          "/ITEMS/mcl_nether/textures/",
		"raw_ores":        "/ITEMS/mcl_raw_ores/textures/",
		"smithing_table":  "/ITEMS/mcl_smithing_table/textures/",
		"smoker":          "/ITEMS/mcl_smoker/textures/",
		"stonecutter":     "/ITEMS/mcl_stonecutter/textures/",
		"sus_stew":        "/ITEMS/mcl_sus_stew/textures/",
		"tnt":             "/ITEMS/mcl_tnt/textures/",
		"tools":           "/ITEMS/mcl_tools/textures/",
		"torches":         "/ITEMS/mcl_torches/textures/",
		"wool":            "/ITEMS/mcl_wool/textures/",
		//"": "/ITEMS//textures/",
	}
	return cloniaPaths
}
func basicITEMS() [][4]string {
	equivalents := [][4]string{
		// -- mcl_amethyst
		{"block", "amethyst_block.png", "amethyst", "mcl_amethyst_amethyst_block.png"},
		{"block", "large_amethyst_bud.png", "amethyst", "mcl_amethyst_amethyst_bud_large.png"},
		{"block", "medium_amethyst_bud.png", "amethyst", "mcl_amethyst_amethyst_bud_medium.png"},
		{"block", "small_amethyst_bud.png", "amethyst", "mcl_amethyst_amethyst_bud_small.png"},
		{"block", "amethyst_cluster.png", "amethyst", "mcl_amethyst_amethyst_cluster.png"},
		{"item", "amethyst_shard.png", "amethyst", "mcl_amethyst_amethyst_shard.png"},
		{"block", "budding_amethyst.png", "amethyst", "mcl_amethyst_budding_amethyst.png"},
		{"block", "calcite.png", "amethyst", "mcl_amethyst_calcite_block.png"},
		{"block", "tinted_glass.png", "amethyst", "mcl_amethyst_tinted_glass.png"},
		// -- mcl_anvils
		{"block", "anvil.png", "anvils", "mcl_anvils_anvil_base.png"},
		{"block", "anvil.png", "anvils", "mcl_anvils_anvil_side.png"},
		//"mcl_anvils_anvil_top_damaged_0.png" handled elsewhere
		//"mcl_anvils_anvil_top_damaged_1.png" handled elsewhere
		//"mcl_anvils_anvil_top_damaged_2.png" handled elsewhere
		// -- mcl_armor
		// -- mcl_armor_stand
		{"item", "armor_stand.png", "armor_stand", "armor_stand.png"},
		// -- mcl_bamboo
		//{"block", ".png", "bamboo", "mcl_bamboo_bamboo.png"},
		{"block", "bamboo_block.png", "bamboo", "mcl_bamboo_bamboo_block.png"},
		{"block", "stripped_bamboo_block.png", "bamboo", "mcl_bamboo_bamboo_block_stripped.png"},
		{"block", "bamboo_block_top.png", "bamboo", "mcl_bamboo_bamboo_bottom.png"},
		{"block", "stripped_bamboo_block_top.png", "bamboo", "mcl_bamboo_bamboo_bottom_stripped.png"},
		//{"block", ".png", "bamboo", "mcl_bamboo_bamboo_fpm.png"},    //special attention
		{"block", "bamboo_planks.png", "bamboo", "mcl_bamboo_bamboo_plank.png"},
		{"block", "bamboo_mosaic.png", "bamboo", "mcl_bamboo_bamboo_plank_mosaic.png"},
		{"item", "bamboo.png", "bamboo", "mcl_bamboo_bamboo_shoot.png"},
		//{"block", ".png", "bamboo", "mcl_bamboo_bamboo_sign.png"},    //special attention
		{"item", "bamboo_sign.png", "bamboo", "mcl_bamboo_bamboo_sign_wield.png"},
		//{"block", "bamboo_door_bottom.png", "bamboo", "mcl_bamboo_door_bottom.png"},
		//{"block", "bamboo_door_bottom.png", "bamboo", "mcl_bamboo_door_bottom_alt.png"},    //test
		//{"block", "bamboo_door_top.png", "bamboo", "mcl_bamboo_door_top.png"},
		//{"block", "bamboo_door_top.png", "bamboo", "mcl_bamboo_door_top_alt.png"},    //test
		{"item", "bamboo_door.png", "bamboo", "mcl_bamboo_door_wield.png"},
		//{"block", ".png", "bamboo", "mcl_bamboo_endcap.png"}, //what is this?
		{"block", "bamboo_fence_particle.png", "bamboo", "mcl_bamboo_fence_bamboo.png"},
		{"block", "bamboo_fence_gate_particle.png", "bamboo", "mcl_bamboo_fence_gate_bamboo.png"},
		//    --scaffolding is broken in vanilla Mineclonia right now.
		//{"block", "bamboo_stage0.png", "bamboo", "mcl_bamboo_flower_pot.png"},    //broken?
		{"block", "scaffolding_bottom.png", "bamboo", "mcl_bamboo_scaffolding_bottom.png"}, //broken?
		{"block", "scaffolding_side.png", "bamboo", "mcl_bamboo_scaffolding_side.png"},     //broken?
		{"block", "scaffolding_top.png", "bamboo", "mcl_bamboo_scaffolding_top.png"},
		{"block", "bamboo_trapdoor.png", "bamboo", "mcl_bamboo_trapdoor_side.png"},
		// -- mcl_banners
		// -- mcl_barrels
		{"block", "barrel_bottom.png", "barrels", "mcl_barrels_barrel_bottom.png"},
		{"block", "barrel_side.png", "barrels", "mcl_barrels_barrel_side.png"},
		{"block", "barrel_top.png", "barrels", "mcl_barrels_barrel_top.png"},
		{"block", "barrel_top_open.png", "barrels", "mcl_barrels_barrel_top_open.png"},
		// -- mcl_beacons
		// --// special attention
		// -- mcl_beds
		{"bed", "black.png", "beds", "mcl_beds_bed_black.png"},
		//"mcl_beds_bed_black_inv.png"  no match
		{"bed", "blue.png", "beds", "mcl_beds_bed_blue.png"},
		//"mcl_beds_bed_blue_inv.png"  no match
		{"bed", "brown.png", "beds", "mcl_beds_bed_brown.png"},
		//"mcl_beds_bed_brown_inv.png"  no match
		{"bed", "cyan.png", "beds", "mcl_beds_bed_cyan.png"},
		//"mcl_beds_bed_cyan_inv.png"  no match
		{"bed", "green.png", "beds", "mcl_beds_bed_green.png"},
		//"mcl_beds_bed_green_inv.png"  no match
		{"bed", "gray.png", "beds", "mcl_beds_bed_grey.png"},
		//"mcl_beds_bed_grey_inv.png"  no match
		{"bed", "light_blue.png", "beds", "mcl_beds_bed_light_blue.png"},
		//"mcl_beds_bed_light_blue_inv.png"  no match
		{"bed", "lime.png", "beds", "mcl_beds_bed_lime.png"},
		//"mcl_beds_bed_lime_inv.png"  no match
		{"bed", "magenta.png", "beds", "mcl_beds_bed_magenta.png"},
		//"mcl_beds_bed_magenta_inv.png"  no match
		{"bed", "orange.png", "beds", "mcl_beds_bed_orange.png"},
		//"mcl_beds_bed_orange_inv.png"  no match
		{"bed", "pink.png", "beds", "mcl_beds_bed_pink.png"},
		//"mcl_beds_bed_pink_inv.png"  no match
		{"bed", "purple.png", "beds", "mcl_beds_bed_purple.png"},
		//"mcl_beds_bed_purple_inv.png"  no match
		{"bed", "red.png", "beds", "mcl_beds_bed_red.png"},
		//"mcl_beds_bed_red_inv.png"  no match
		{"bed", "light_gray.png", "beds", "mcl_beds_bed_silver.png"},
		//"mcl_beds_bed_silver_inv.png"  no match
		{"bed", "white.png", "beds", "mcl_beds_bed_white.png"},
		//"mcl_beds_bed_white_inv.png"  no match
		{"bed", "yellow.png", "beds", "mcl_beds_bed_yellow.png"},
		//"mcl_beds_bed_yellow_inv.png"  no match
		{"block", "respawn_anchor_bottom.png", "beds", "respawn_anchor_bottom.png"},
		{"block", "respawn_anchor_side0.png", "beds", "respawn_anchor_side0.png"},
		{"block", "respawn_anchor_side1.png", "beds", "respawn_anchor_side1.png"},
		{"block", "respawn_anchor_side2.png", "beds", "respawn_anchor_side2.png"},
		{"block", "respawn_anchor_side3.png", "beds", "respawn_anchor_side3.png"},
		{"block", "respawn_anchor_side4.png", "beds", "respawn_anchor_side4.png"},
		{"block", "respawn_anchor_top_off.png", "beds", "respawn_anchor_top_off.png"},
		{"block", "respawn_anchor_top.png", "beds", "respawn_anchor_top_on.png"},
		// -- mcl_beehives
		{"block", "beehive_end.png", "beehives", "mcl_beehives_beehive_end.png"},
		{"block", "beehive_front.png", "beehives", "mcl_beehives_beehive_front.png"},
		{"block", "beehive_front_honey.png", "beehives", "mcl_beehives_beehive_front_honey.png"},
		{"block", "beehive_side.png", "beehives", "mcl_beehives_beehive_side.png"},
		{"block", "bee_nest_bottom.png", "beehives", "mcl_beehives_bee_nest_bottom.png"},
		{"block", "bee_nest_front.png", "beehives", "mcl_beehives_bee_nest_front.png"},
		{"block", "bee_nest_front_honey.png", "beehives", "mcl_beehives_bee_nest_front_honey.png"},
		{"block", "bee_nest_side.png", "beehives", "mcl_beehives_bee_nest_side.png"},
		{"block", "bee_nest_top.png", "beehives", "mcl_beehives_bee_nest_top.png"},
		// -- mcl_bells
		// -- mcl_blackstone
		{"block", "quartz_bricks.png", "blackstone", "mcl_backstone_quartz_bricks.png"},
		{"block", "basalt_side.png", "blackstone", "mcl_blackstone_basalt_side.png"},
		{"block", "polished_basalt_side.png", "blackstone", "mcl_blackstone_basalt_side_polished.png"},
		{"block", "smooth_basalt.png", "blackstone", "mcl_blackstone_basalt_smooth.png"},
		{"block", "basalt_top.png", "blackstone", "mcl_blackstone_basalt_top.png"},
		{"block", "polished_basalt_top.png", "blackstone", "mcl_blackstone_basalt_top_polished.png"},
		{"block", "chiseled_polished_blackstone.png", "blackstone", "mcl_blackstone_chiseled_polished.png"},
		{"block", "gilded_blackstone.png", "blackstone", "mcl_blackstone_gilded.png"},
		{"block", "gilded_blackstone.png", "blackstone", "mcl_blackstone_gilded_side.png"},
		{"block", "polished_blackstone.png", "blackstone", "mcl_blackstone_polished.png"},
		{"block", "polished_blackstone_bricks.png", "blackstone", "mcl_blackstone_polished_bricks.png"},
		{"block", "cracked_polished_blackstone_bricks.png", "blackstone", "mcl_blackstone_polished_bricks_cracked.png"},
		{"block", "blackstone.png", "blackstone", "mcl_blackstone_side.png"},
		{"block", "soul_soil.png", "blackstone", "mcl_blackstone_soul_soil.png"},
		{"block", "blackstone.png", "blackstone", "mcl_blackstone_top.png"}, //Special Case?
		{"block", "blackstone_top.png", "blackstone", "mcl_blackstone_top.png"},
		{"block", "soul_fire_0.png", "blackstone", "soul_fire_basic_flame.png"},
		{"block", "soul_fire_0.png", "blackstone", "soul_fire_basic_flame_animated.png"}, //animated
		{"block", "soul_torch.png", "blackstone", "soul_torch_on_floor.png"},
		{"block", "soul_torch.png", "blackstone", "soul_torch_on_floor_animated.png"}, //animated
		// -- mcl_blast_furnace
		{"block", "blast_furnace_front.png", "blast_furnace", "blast_furnace_front.png"},
		{"block", "blast_furnace_front_on.png", "blast_furnace", "blast_furnace_front_on.png"},
		{"block", "blast_furnace_side.png", "blast_furnace", "blast_furnace_side.png"},
		{"block", "blast_furnace_top.png", "blast_furnace", "blast_furnace_top.png"},
		// -- mcl_bone_meal
		{"item", "bone_meal.png", "bone_meal", "mcl_bone_meal_bone_meal.png"},
		{"particle", "glint.png", "bone_meal", "mcl_particles_bonemeal.png"},
		// -- mcl_books
		{"item", ".png", "books", "default_book.png"},
		{"block", "bookshelf.png", "books", "default_bookshelf.png"},
		//{"item", ".png", "books", "mcl_book_book_empty_slot.png"}, //no match
		//{"gui", "book.png", "books", "mcl_books_book_bg.png"}, // special attention & it's not a square!
		{"block", "chiseled_bookshelf_top.png", "books", "mcl_books_bookshelf_top.png"},
		{"item", "writable_book.png", "books", "mcl_books_book_writable.png"},
		{"item", "written_book.png", "books", "mcl_books_book_written.png"},
		//{"item", ".png", "books", "mcl_books_button9.png"},         // special attention
		//{"item", ".png", "books", "mcl_books_button9_pressed.png"}, // special attention
		// -- mcl_bows
		// -- mcl_brewing
		// -- mcl_buckets
		// -- mcl_cake
		// -- mcl_campfires
		// -- mcl_cartography_table
		// -- mcl_cauldrons
		// -- mcl_cherry_blossom
		//{"block", "cherry_door_bottom.png", "cherry_blossom", "mcl_cherry_blossom_door_bottom.png"}, //flipped
		{"block", "cherry_door_bottom.png", "cherry_blossom", "mcl_cherry_blossom_door_bottom_bottompart.png"},
		{"block", "cherry_door_bottom.png", "cherry_blossom", "mcl_cherry_blossom_door_bottom_side.png"},
		{"item", "cherry_door.png", "cherry_blossom", "mcl_cherry_blossom_door_inv.png"},
		//{"block", "cherry_door_top.png", "cherry_blossom", "mcl_cherry_blossom_door_top.png"}, //flipped
		{"block", "cherry_door_top.png", "cherry_blossom", "mcl_cherry_blossom_door_top_side.png"},
		{"block", "cherry_door_top.png", "cherry_blossom", "mcl_cherry_blossom_door_top_toppart.png"},
		{"block", "cherry_leaves.png", "cherry_blossom", "mcl_cherry_blossom_leaves.png"},
		{"block", "cherry_log.png", "cherry_blossom", "mcl_cherry_blossom_log.png"},
		{"block", "stripped_cherry_log.png", "cherry_blossom", "mcl_cherry_blossom_log_stripped.png"},
		{"block", "cherry_log_top.png", "cherry_blossom", "mcl_cherry_blossom_log_top.png"},
		{"block", "stripped_cherry_log_top.png", "cherry_blossom", "mcl_cherry_blossom_log_top_stripped.png"},
		//{"block", ".png", "cherry_blossom", "mcl_cherry_blossom_particle.png"},
		//{"block", ".png", "cherry_blossom", "mcl_cherry_blossom_particle_1.png"},
		//{"block", ".png", "cherry_blossom", "mcl_cherry_blossom_particle_2.png"},
		//{"block", ".png", "cherry_blossom", "mcl_cherry_blossom_particle_3.png"},
		{"block", "pink_petals.png", "cherry_blossom", "mcl_cherry_blossom_pink_petals.png"},
		{"block", "pink_petals.png", "cherry_blossom", "mcl_cherry_blossom_pink_petals_inv.png"},
		{"block", "cherry_planks.png", "cherry_blossom", "mcl_cherry_blossom_planks.png"},
		{"block", "cherry_sapling.png", "cherry_blossom", "mcl_cherry_blossom_sapling.png"},
		{"block", "cherry_trapdoor.png", "cherry_blossom", "mcl_cherry_blossom_trapdoor.png"},
		{"block", "cherry_trapdoor.png", "cherry_blossom", "mcl_cherry_blossom_trapdoor_side.png"},
		// -- mcl_chests
		// -- mcl_clock
		// -- mcl_cocoas
		// -- mcl_colorblocks
		// -- mcl_compass
		// -- mcl_composters
		// -- mcl_conduits
		// -- mcl_copper
		//"mcl_copper_anti_oxidation_particle.png" //no match?
		{"block", "copper_block.png", "copper", "mcl_copper_block.png"},
		{"block", "copper_bulb.png", "copper", "mcl_copper_block_bulb_off.png"},
		{"block", "copper_bulb_lit.png", "copper", "mcl_copper_block_bulb_on.png"},
		{"block", "copper_bulb_powered.png", "copper", "mcl_copper_block_bulb_powered_off.png"},
		{"block", "copper_bulb_lit_powered.png", "copper", "mcl_copper_block_bulb_powered_on.png"},
		{"block", "chiseled_copper.png", "copper", "mcl_copper_block_chiseled.png"},
		{"block", "cut_copper.png", "copper", "mcl_copper_block_cut.png"},
		{"block", "copper_grate.png", "copper", "mcl_copper_block_grate.png"},
		{"block", "raw_copper_block.png", "copper", "mcl_copper_block_raw.png"},
		{"item", "copper_door.png", "copper", "mcl_copper_door.png"},
		{"block", "copper_door_bottom.png", "copper", "mcl_copper_door_bottom.png"},
		{"item", "exposed_copper_door.png", "copper", "mcl_copper_door_exposed.png"},
		{"block", "exposed_copper_door_bottom.png", "copper", "mcl_copper_door_exposed_bottom.png"},
		//{"block", "", "copper", "mcl_copper_door_exposed_lower.png"},
		//{"block", "", "copper", "mcl_copper_door_exposed_lower_bottompart.png"},
		//{"block", "", "copper", "mcl_copper_door_exposed_lower_side.png"},
		{"block", "exposed_copper_door_top.png", "copper", "mcl_copper_door_exposed_top.png"},
		//{"block", "", "copper", "mcl_copper_door_exposed_upper.png"},
		//{"block", "", "copper", "mcl_copper_door_exposed_upper_side.png"},
		//{"block", "", "copper", "mcl_copper_door_exposed_upper_toppart.png"},
		//{"block", "", "copper", "mcl_copper_door_lower.png"},
		//{"block", "", "copper", "mcl_copper_door_lower_bottompart.png"},
		//{"block", "", "copper", "mcl_copper_door_lower_side.png"},
		{"item", "oxidized_copper_door.png", "copper", "mcl_copper_door_oxidized.png"},
		{"block", "oxidized_copper_door_bottom.png", "copper", "mcl_copper_door_oxidized_bottom.png"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_lower.png"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_lower_bottompart.png"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_lower_side.png"},
		{"block", "oxidized_copper_door_top.png", "copper", "mcl_copper_door_oxidized_top.png"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_upper.png"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_upper_side.png"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_upper_toppart.png"},
		{"block", "copper_door_top.png", "copper", "mcl_copper_door_top.png"},
		//{"block", "", "copper", "mcl_copper_door_upper.png"},
		//{"block", "", "copper", "mcl_copper_door_upper_side.png"},
		//{"block", "", "copper", "mcl_copper_door_upper_toppart.png"},
		{"item", "weathered_copper_door.png", "copper", "mcl_copper_door_weathered.png"},
		{"block", "weathered_copper_door_bottom.png", "copper", "mcl_copper_door_weathered_bottom.png"},
		//{"block", "", "copper", "mcl_copper_door_weathered_lower.png"},
		//{"block", "", "copper", "mcl_copper_door_weathered_lower_bottompart.png"},
		//{"block", "", "copper", "mcl_copper_door_weathered_lower_side.png"},
		{"block", "weathered_copper_door_top.png", "copper", "mcl_copper_door_weathered_top.png"},
		//{"block", "", "copper", "mcl_copper_door_weathered_upper.png"},
		//{"block", "", "copper", "mcl_copper_door_weathered_upper_side.png"},
		//{"block", "", "copper", "mcl_copper_door_weathered_upper_toppart.png"},
		{"block", "exposed_copper.png", "copper", "mcl_copper_exposed.png"},
		{"block", "exposed_copper_bulb.png", "copper", "mcl_copper_exposed_bulb_off.png"},
		{"block", "exposed_copper_bulb_lit.png", "copper", "mcl_copper_exposed_bulb_on.png"},
		{"block", "exposed_copper_bulb_powered.png", "copper", "mcl_copper_exposed_bulb_powered_off.png"},
		{"block", "exposed_copper_bulb_lit_powered.png", "copper", "mcl_copper_exposed_bulb_powered_on.png"},
		{"block", "exposed_chiseled_copper.png", "copper", "mcl_copper_exposed_chiseled.png"},
		{"block", "exposed_cut_copper.png", "copper", "mcl_copper_exposed_cut.png"},
		{"block", "exposed_copper_grate.png", "copper", "mcl_copper_exposed_grate.png"},
		{"item", "copper_ingot.png", "copper", "mcl_copper_ingot.png"},
		{"item", "copper_door.png", "copper", "mcl_copper_item_door.png"},
		{"item", "exposed_copper_door.png", "copper", "mcl_copper_item_door_exposed.png"},
		{"item", "oxidized_copper_door.png", "copper", "mcl_copper_item_door_oxidized.png"},
		{"item", "weathered_copper_door.png", "copper", "mcl_copper_item_door_weathered.png"},
		{"block", "copper_ore.png", "copper", "mcl_copper_ore.png"},
		{"block", "oxidized_copper.png", "copper", "mcl_copper_oxidized.png"},
		{"block", "oxidized_copper_bulb.png", "copper", "mcl_copper_oxidized_bulb_off.png"},
		{"block", "oxidized_copper_bulb_lit.png", "copper", "mcl_copper_oxidized_bulb_on.png"},
		{"block", "oxidized_copper_bulb_powered.png", "copper", "mcl_copper_oxidized_bulb_powered_off.png"},
		{"block", "oxidized_copper_bulb_lit_powered.png", "copper", "mcl_copper_oxidized_bulb_powered_on.png"},
		{"block", "oxidized_chiseled_copper.png", "copper", "mcl_copper_oxidized_chiseled.png"},
		{"block", "oxidized_cut_copper.png", "copper", "mcl_copper_oxidized_cut.png"},
		{"block", "oxidized_copper_grate.png", "copper", "mcl_copper_oxidized_grate.png"},
		{"item", "raw_copper.png", "copper", "mcl_copper_raw.png"},
		{"block", "copper_trapdoor.png", "copper", "mcl_copper_trapdoor.png"},
		{"block", "exposed_copper_trapdoor.png", "copper", "mcl_copper_trapdoor_exposed.png"},
		{"block", "exposed_copper_trapdoor.png", "copper", "mcl_copper_trapdoor_exposed_side.png"},
		{"block", "oxidized_copper_trapdoor.png", "copper", "mcl_copper_trapdoor_oxidized.png"},
		{"block", "oxidized_copper_trapdoor.png", "copper", "mcl_copper_trapdoor_oxidized_side.png"},
		{"block", "copper_trapdoor.png", "copper", "mcl_copper_trapdoor_side.png"},
		{"block", "weathered_copper_trapdoor.png", "copper", "mcl_copper_trapdoor_weathered.png"},
		{"block", "weathered_copper_trapdoor.png", "copper", "mcl_copper_trapdoor_weathered_side.png"},
		{"block", "weathered_copper.png", "copper", "mcl_copper_weathered.png"},
		{"block", "weathered_copper_bulb.png", "copper", "mcl_copper_weathered_bulb_off.png"},
		{"block", "weathered_copper_bulb_lit.png", "copper", "mcl_copper_weathered_bulb_on.png"},
		{"block", "weathered_copper_bulb_powered.png", "copper", "mcl_copper_weathered_bulb_powered_off.png"},
		{"block", "weathered_copper_bulb_lit_powered.png", "copper", "mcl_copper_weathered_bulb_powered_on.png"},
		{"block", "weathered_chiseled_copper.png", "copper", "mcl_copper_weathered_chiseled.png"},
		{"block", "weathered_cut_copper.png", "copper", "mcl_copper_weathered_cut.png"},
		{"block", "weathered_copper_grate.png", "copper", "mcl_copper_weathered_grate.png"},
		// -- mcl_core
		{"block", "acacia_leaves.png", "core", "default_acacia_leaves.png"},
		{"block", "acacia_sapling.png", "core", "default_acacia_sapling.png"},
		{"block", "acacia_log.png", "core", "default_acacia_tree.png"},
		{"block", "acacia_log_top.png", "core", "default_acacia_tree_top.png"},
		{"block", "acacia_planks.png", "core", "default_acacia_wood.png"},
		{"item", "apple.png", "core", "default_apple.png"},
		{"block", "bricks.png", "core", "default_brick.png"},
		{"block", "clay.png", "core", "default_clay.png"},
		{"item", "brick.png", "core", "default_clay_brick.png"},
		{"item", "clay_ball.png", "core", "default_clay_lump.png"},
		{"block", "coal_block.png", "core", "default_coal_block.png"},
		{"item", "coal.png", "core", "default_coal_lump.png"},
		{"block", "cobblestone.png", "core", "default_cobble.png"},
		{"item", "diamond.png", "core", "default_diamond.png"},
		{"block", "diamond_block.png", "core", "default_diamond_block.png"},
		{"block", "dirt.png", "core", "default_dirt.png"},
		{"block", "dead_bush.png", "core", "default_dry_shrub.png"},
		{"item", "flint.png", "core", "default_flint.png"},
		{"block", "glass.png", "core", "default_glass.png"},
		//"default_glass_detail.png"	//no match?
		{"block", "gold_block.png", "core", "default_gold_block.png"},
		{"item", "gold_ingot.png", "core", "default_gold_ingot.png"},
		{"block", "gravel.png", "core", "default_gravel.png"},
		{"block", "ice.png", "core", "default_ice.png"},
		{"block", "jungle_leaves.png", "core", "default_jungleleaves.png"},
		{"block", "jungle_sapling.png", "core", "default_junglesapling.png"},
		{"block", "jungle_log.png", "core", "default_jungletree.png"},
		{"block", "jungle_log_top.png", "core", "default_jungletree_top.png"},
		{"block", "jungle_planks.png", "core", "default_junglewood.png"},
		{"block", "ladder.png", "core", "default_ladder.png"},
		//{"block", "lava_flow.png", "core", "default_lava_flowing_animated.png"}, //special attention
		//{"block", "lava_still.png", "core", "default_lava_source_animated.png"}, //special attention
		{"block", "oak_leaves.png", "core", "default_leaves.png"},
		{"block", "mossy_cobblestone.png", "core", "default_mossycobble.png"},
		{"block", "obsidian.png", "core", "default_obsidian.png"},
		{"item", "paper.png", "core", "default_paper.png"},
		{"block", "sand.png", "core", "default_sand.png"},
		{"block", "oak_sapling.png", "core", "default_sapling.png"},
		{"block", "snow.png", "core", "default_snow.png"},
		{"block", "iron_block.png", "core", "default_steel_block.png"},
		{"item", "iron_ingot.png", "core", "default_steel_ingot.png"},
		{"item", "stick.png", "core", "default_stick.png"},
		{"block", "stone.png", "core", "default_stone.png"},
		{"block", "stone_bricks.png", "core", "default_stone_brick.png"},
		{"block", "oak_log.png", "core", "default_tree.png"},
		{"block", "oak_log_top.png", "core", "default_tree_top.png"},
		//{"block", "water_flow.png", "core", "default_water_flowing_animated.png"}, //special attention
		//{"block", "water_still.png", "core", "default_water_source_animated.png"}, //special attention
		{"block", "oak_planks.png", "core", "default_wood.png"},
		{"block", "andesite.png", "core", "mcl_core_andesite.png"},
		{"block", "polished_andesite.png", "core", "mcl_core_andesite_smooth.png"},
		{"item", "golden_apple.png", "core", "mcl_core_apple_golden.png"},
		{"item", "barrier.png", "core", "mcl_core_barrier.png"},
		{"block", "bedrock.png", "core", "mcl_core_bedrock.png"},
		{"block", "bone_block_side.png", "core", "mcl_core_bone_block_side.png"},
		{"block", "bone_block_top.png", "core", "mcl_core_bone_block_top.png"},
		{"item", "bowl.png", "core", "mcl_core_bowl.png"},
		{"block", "cactus_bottom.png", "core", "mcl_core_cactus_bottom.png"},
		{"block", "cactus_side.png", "core", "mcl_core_cactus_side.png"},
		{"block", "cactus_top.png", "core", "mcl_core_cactus_top.png"},
		{"item", "charcoal.png", "core", "mcl_core_charcoal.png"},
		{"block", "coal_ore.png", "core", "mcl_core_coal_ore.png"},
		{"block", "coarse_dirt.png", "core", "mcl_core_coarse_dirt.png"},
		{"block", "crying_obsidian.png", "core", "mcl_core_crying_obsidian.png"}, //special attention? might be ok
		//"mcl_core_crying_obsidian_tear.png"		//no match?
		//"mcl_core_crying_obsidian_tear2.png"		//no match?
		//"mcl_core_crying_obsidian_tear3.png"		//no match?
		{"block", "diamond_ore.png", "core", "mcl_core_diamond_ore.png"},
		{"block", "diorite.png", "core", "mcl_core_diorite.png"},
		{"block", "polished_diorite.png", "core", "mcl_core_diorite_smooth.png"},
		{"block", "podzol_side.png", "core", "mcl_core_dirt_podzol_side.png"},
		{"block", "podzol_top.png", "core", "mcl_core_dirt_podzol_top.png"},
		{"item", "emerald.png", "core", "mcl_core_emerald.png"},
		{"block", "emerald_block.png", "core", "mcl_core_emerald_block.png"},
		{"block", "emerald_ore.png", "core", "mcl_core_emerald_ore.png"},
		{"block", "frosted_ice_0.png", "core", "mcl_core_frosted_ice_0.png"},
		{"block", "frosted_ice_1.png", "core", "mcl_core_frosted_ice_1.png"},
		{"block", "frosted_ice_2.png", "core", "mcl_core_frosted_ice_2.png"},
		{"block", "frosted_ice_3.png", "core", "mcl_core_frosted_ice_3.png"},
		//Glass TODO: All glass has normal and "detail" textures. Must create "detail".
		{"block", "black_stained_glass.png", "core", "mcl_core_glass_black.png"},
		{"block", "black_stained_glass.png", "core", "mcl_core_glass_black_detail.png"},
		{"block", "blue_stained_glass.png", "core", "mcl_core_glass_blue.png"},
		{"block", "blue_stained_glass.png", "core", "mcl_core_glass_blue_detail.png"},
		{"block", "brown_stained_glass.png", "core", "mcl_core_glass_brown.png"},
		{"block", "brown_stained_glass.png", "core", "mcl_core_glass_brown_detail.png"},
		{"block", "cyan_stained_glass.png", "core", "mcl_core_glass_cyan.png"},
		{"block", "cyan_stained_glass.png", "core", "mcl_core_glass_cyan_detail.png"},
		{"block", "gray_stained_glass.png", "core", "mcl_core_glass_gray.png"},
		{"block", "gray_stained_glass.png", "core", "mcl_core_glass_gray_detail.png"},
		{"block", "green_stained_glass.png", "core", "mcl_core_glass_green.png"},
		{"block", "green_stained_glass.png", "core", "mcl_core_glass_green_detail.png"},
		{"block", "light_blue_stained_glass.png", "core", "mcl_core_glass_light_blue.png"},
		{"block", "light_blue_stained_glass.png", "core", "mcl_core_glass_light_blue_detail.png"},
		{"block", "lime_stained_glass.png", "core", "mcl_core_glass_lime.png"},
		{"block", "lime_stained_glass.png", "core", "mcl_core_glass_lime_detail.png"},
		{"block", "magenta_stained_glass.png", "core", "mcl_core_glass_magenta.png"},
		{"block", "magenta_stained_glass.png", "core", "mcl_core_glass_magenta_detail.png"},
		{"block", "orange_stained_glass.png", "core", "mcl_core_glass_orange.png"},
		{"block", "orange_stained_glass.png", "core", "mcl_core_glass_orange_detail.png"},
		{"block", "pink_stained_glass.png", "core", "mcl_core_glass_pink.png"},
		{"block", "pink_stained_glass.png", "core", "mcl_core_glass_pink_detail.png"},
		{"block", "purple_stained_glass.png", "core", "mcl_core_glass_purple.png"},
		{"block", "purple_stained_glass.png", "core", "mcl_core_glass_purple_detail.png"},
		{"block", "red_stained_glass.png", "core", "mcl_core_glass_red.png"},
		{"block", "red_stained_glass.png", "core", "mcl_core_glass_red_detail.png"},
		{"block", "light_gray_stained_glass.png", "core", "mcl_core_glass_silver.png"},
		{"block", "light_gray_stained_glass.png", "core", "mcl_core_glass_silver_detail.png"},
		{"block", "white_stained_glass.png", "core", "mcl_core_glass_white.png"},
		{"block", "white_stained_glass.png", "core", "mcl_core_glass_white_detail.png"},
		{"block", "yellow_stained_glass.png", "core", "mcl_core_glass_yellow.png"},
		{"block", "yellow_stained_glass.png", "core", "mcl_core_glass_yellow_detail.png"},
		//Glass TODO
		{"item", "gold_nugget.png", "core", "mcl_core_gold_nugget.png"},
		{"block", "gold_ore.png", "core", "mcl_core_gold_ore.png"},
		{"block", "granite.png", "core", "mcl_core_granite.png"},
		{"block", "polished_granite.png", "core", "mcl_core_granite_smooth.png"},
		{"block", "grass_block_side_overlay.png", "core", "mcl_core_grass_block_side_overlay.png"},
		{"block", "grass_block_top.png", "core", "mcl_core_grass_block_top.png"},
		{"block", "dirt_path_side.png", "core", "mcl_core_grass_path_side.png"},
		{"block", "dirt_path_top.png", "core", "mcl_core_grass_path_top.png"},
		{"block", "grass_block_snow.png", "core", "mcl_core_grass_side_snowed.png"},
		{"block", "packed_ice.png", "core", "mcl_core_ice_packed.png"},
		{"item", "iron_nugget.png", "core", "mcl_core_iron_nugget.png"},
		{"block", "iron_ore.png", "core", "mcl_core_iron_ore.png"},
		{"item", "lapis_lazuli.png", "core", "mcl_core_lapis.png"},
		{"block", "lapis_block.png", "core", "mcl_core_lapis_block.png"},
		{"block", "lapis_ore.png", "core", "mcl_core_lapis_ore.png"},
		{"block", "dark_oak_leaves.png", "core", "mcl_core_leaves_big_oak.png"},
		{"block", "birch_leaves.png", "core", "mcl_core_leaves_birch.png"},
		{"block", "spruce_leaves.png", "core", "mcl_core_leaves_spruce.png"},
		{"item", "light_00.png", "core", "mcl_core_light_0.png"},
		{"item", "light_01.png", "core", "mcl_core_light_1.png"},
		{"item", "light_02.png", "core", "mcl_core_light_2.png"},
		{"item", "light_03.png", "core", "mcl_core_light_3.png"},
		{"item", "light_04.png", "core", "mcl_core_light_4.png"},
		{"item", "light_05.png", "core", "mcl_core_light_5.png"},
		{"item", "light_06.png", "core", "mcl_core_light_6.png"},
		{"item", "light_07.png", "core", "mcl_core_light_7.png"},
		{"item", "light_08.png", "core", "mcl_core_light_8.png"},
		{"item", "light_09.png", "core", "mcl_core_light_9.png"},
		{"item", "light_10.png", "core", "mcl_core_light_10.png"},
		{"item", "light_11.png", "core", "mcl_core_light_11.png"},
		{"item", "light_12.png", "core", "mcl_core_light_12.png"},
		{"item", "light_13.png", "core", "mcl_core_light_13.png"},
		{"item", "light_14.png", "core", "mcl_core_light_14.png"}, //no light_15 in Mineclonia
		{"block", "dark_oak_log.png", "core", "mcl_core_log_big_oak.png"},
		{"block", "dark_oak_log_top.png", "core", "mcl_core_log_big_oak_top.png"},
		{"block", "birch_log.png", "core", "mcl_core_log_birch.png"},
		{"block", "birch_log_top.png", "core", "mcl_core_log_birch_top.png"},
		{"block", "spruce_log.png", "core", "mcl_core_log_spruce.png"},
		{"block", "spruce_log_top.png", "core", "mcl_core_log_spruce_top.png"},
		//{"block", "", "core", "mcl_core_mycelium_particle.png"}, 	//special attention
		{"block", "mycelium_side.png", "core", "mcl_core_mycelium_side.png"},
		{"block", "mycelium_top.png", "core", "mcl_core_mycelium_top.png"},
		//{"block", "", "core", "mcl_core_palette_grass.png"}, 	//special attention
		//{"block", "", "core", "mcl_core_palette_leaves.png"}, 		//special attention
		{"block", "sugar_cane.png", "core", "mcl_core_papyrus.png"},
		{"block", "dark_oak_planks.png", "core", "mcl_core_planks_big_oak.png"},
		{"block", "birch_planks.png", "core", "mcl_core_planks_birch.png"},
		{"block", "spruce_planks.png", "core", "mcl_core_planks_spruce.png"},
		{"block", "red_sand.png", "core", "mcl_core_red_sand.png"},
		{"block", "red_sandstone_bottom.png", "core", "mcl_core_red_sandstone_bottom.png"},
		{"block", "chiseled_red_sandstone.png", "core", "mcl_core_red_sandstone_carved.png"},
		{"block", "red_sandstone.png", "core", "mcl_core_red_sandstone_normal.png"},
		{"block", "cut_red_sandstone.png", "core", "mcl_core_red_sandstone_smooth.png"},
		{"block", "red_sandstone_top.png", "core", "mcl_core_red_sandstone_top.png"},
		{"block", "redstone_ore.png", "core", "mcl_core_redstone_ore.png"},
		{"item", "sugar_cane.png", "core", "mcl_core_reeds.png"},
		{"block", "sandstone_bottom.png", "core", "mcl_core_sandstone_bottom.png"},
		{"block", "chiseled_sandstone.png", "core", "mcl_core_sandstone_carved.png"},
		{"block", "sandstone.png", "core", "mcl_core_sandstone_normal.png"},
		{"block", "cut_sandstone.png", "core", "mcl_core_sandstone_smooth.png"},
		{"block", "sandstone_top.png", "core", "mcl_core_sandstone_top.png"},
		{"block", "dark_oak_sapling.png", "core", "mcl_core_sapling_big_oak.png"},
		{"block", "birch_sapling.png", "core", "mcl_core_sapling_birch.png"},
		{"block", "spruce_sapling.png", "core", "mcl_core_sapling_spruce.png"},
		{"block", "slime_block.png", "core", "mcl_core_slime.png"},
		{"block", "chiseled_stone_bricks.png", "core", "mcl_core_stonebrick_carved.png"},
		{"block", "cracked_stone_bricks.png", "core", "mcl_core_stonebrick_cracked.png"},
		{"block", "mossy_stone_bricks.png", "core", "mcl_core_stonebrick_mossy.png"},
		{"block", "stripped_acacia_log.png", "core", "mcl_core_stripped_acacia_side.png"},
		{"block", "stripped_acacia_log_top.png", "core", "mcl_core_stripped_acacia_top.png"},
		{"block", "stripped_birch_log.png", "core", "mcl_core_stripped_birch_side.png"},
		{"block", "stripped_birch_log_top.png", "core", "mcl_core_stripped_birch_top.png"},
		{"block", "stripped_dark_oak_log.png", "core", "mcl_core_stripped_dark_oak_side.png"},
		{"block", "stripped_dark_oak_log_top.png", "core", "mcl_core_stripped_dark_oak_top.png"},
		{"block", "stripped_jungle_log.png", "core", "mcl_core_stripped_jungle_side.png"},
		{"block", "stripped_jungle_log_top.png", "core", "mcl_core_stripped_jungle_top.png"},
		{"block", "stripped_oak_log.png", "core", "mcl_core_stripped_oak_side.png"},
		{"block", "stripped_oak_log_top.png", "core", "mcl_core_stripped_oak_top.png"},
		{"block", "stripped_spruce_log.png", "core", "mcl_core_stripped_spruce_side.png"},
		{"block", "stripped_spruce_log_top.png", "core", "mcl_core_stripped_spruce_top.png"},
		{"item", "sugar.png", "core", "mcl_core_sugar.png"},
		//{"block", "vine.png", "core", "mcl_core_vine.png"}, //special attention
		//"mcl_core_void.png"     //no match
		{"block", "cobweb.png", "core", "mcl_core_web.png"},
		{"block", "grass_block_side_overlay.png", "core", "mcl_dirt_grass_shadow.png"}, //special attention?
		{"particle", "lava.png", "core", "mcl_particles_lava.png"},
		//{"block", "", "core", "mcl_stairs_andesite_smooth_slab.png"}, //special attention
		//{"block", "", "core", "mcl_stairs_diorite_smooth_slab.png"},  //special attention
		//{"block", "", "core", "mcl_stairs_granite_smooth_slab.png"},  //special attention
		// -- mcl_crafting_table
		{"block", "crafting_table_front.png", "crafting_table", "crafting_workbench_front.png"},
		{"block", "crafting_table_side.png", "crafting_table", "crafting_workbench_side.png"},
		{"block", "crafting_table_top.png", "crafting_table", "crafting_workbench_top.png"},
		//"gui_crafting_arrow.png"
		//"mcl_crafting_table_inv_fill.png" //no match
		// -- mcl_crimson
		{"block", "crimson_stem_top.png", "crimson", "crimson_hyphae.png"},
		//{"block", "crimson_stem.png", "crimson", "crimson_hyphae_side.png"},  //special attention
		{"block", "crimson_planks.png", "crimson", "crimson_hyphae_wood.png"},
		{"block", "crimson_nylium.png", "crimson", "crimson_nylium.png"},
		{"block", "crimson_nylium_side.png", "crimson", "crimson_nylium_side.png"},
		{"block", "crimson_roots.png", "crimson", "crimson_roots.png"},
		{"block", "stripped_crimson_stem.png", "crimson", "crimson_stem_stripped_side.png"},
		{"block", "stripped_crimson_stem_top.png", "crimson", "crimson_stem_stripped_top.png"},
		{"block", "crimson_fungus.png", "crimson", "farming_crimson_fungus.png"},
		{"block", "warped_fungus.png", "crimson", "farming_warped_fungus.png"},
		{"item", "crimson_door.png", "crimson", "mcl_crimson_crimson_door.png"},
		//{"block", "crimson_door_bottom.png", "crimson", "mcl_crimson_crimson_door_bottom.png"}, //flip
		//{"block", "crimson_door_top.png", "crimson", "mcl_crimson_crimson_door_top.png"},       //flip
		//{"block", "crimson_planks.png", "crimson", "mcl_crimson_crimson_fence.png"},
		//{"block", ".png", "crimson", "mcl_crimson_crimson_fence_side.png"},
		//{"block", ".png", "crimson", "mcl_crimson_crimson_fence_top.png"},
		{"block", "crimson_trapdoor.png", "crimson", "mcl_crimson_crimson_trapdoor.png"},
		{"block", "crimson_trapdoor.png", "crimson", "mcl_crimson_crimson_trapdoor_side.png"},
		{"item", "warped_door.png", "crimson", "mcl_crimson_warped_door.png"},
		//{"block", "warped_door_bottom.png", "crimson", "mcl_crimson_warped_door_bottom.png"}, //flip
		//{"block", "warped_door_top.png", "crimson", "mcl_crimson_warped_door_top.png"},       //flip
		{"block", "warped_planks.png", "crimson", "mcl_crimson_warped_fence.png"},
		//{"block", ".png", "crimson", "mcl_crimson_warped_fence_side.png"},
		//{"block", ".png", "crimson", "mcl_crimson_warped_fence_top.png"},
		{"block", "warped_trapdoor.png", "crimson", "mcl_crimson_warped_trapdoor.png"},
		{"block", "warped_trapdoor.png", "crimson", "mcl_crimson_warped_trapdoor_side.png"},
		{"block", "weeping_vines_plant.png", "crimson", "mcl_crimson_weeping_vines.png"},
		{"block", "crimson_door_bottom.png", "crimson", "mcl_doors_door_crimson_side_lower.png"},
		{"block", "crimson_door_top.png", "crimson", "mcl_doors_door_crimson_side_upper.png"},
		{"block", "warped_door_bottom.png", "crimson", "mcl_doors_door_warped_side_lower.png"},
		{"block", "warped_door_top.png", "crimson", "mcl_doors_door_warped_side_upper.png"},
		{"block", "nether_sprouts.png", "crimson", "nether_sprouts.png"},
		{"block", "nether_wart_block.png", "crimson", "nether_wart_block.png"},
		{"block", "shroomlight.png", "crimson", "shroomlight.png"},
		{"block", "stripped_crimson_stem.png", "crimson", "stripped_crimson_stem.png"},
		{"block", "stripped_crimson_stem.png", "crimson", "stripped_crimson_stem_side.png"},
		{"block", "stripped_crimson_stem_top.png", "crimson", "stripped_crimson_stem_top.png"},
		{"block", "stripped_warped_stem.png", "crimson", "stripped_warped_stem.png"},
		{"block", "stripped_warped_stem.png", "crimson", "stripped_warped_stem_side.png"},
		{"block", "stripped_warped_stem_top.png", "crimson", "stripped_warped_stem_top.png"},
		{"block", "twisting_vines.png", "crimson", "twisting_vines.png"},
		{"block", "twisting_vines_plant.png", "crimson", "twisting_vines_plant.png"},
		{"block", "warped_stem_top.png", "crimson", "warped_hyphae.png"},
		//{"block", "warped_stem.png", "crimson", "warped_hyphae_side.png"},  //special attention
		{"block", "warped_planks.png", "crimson", "warped_hyphae_wood.png"},
		{"block", "warped_nylium.png", "crimson", "warped_nylium.png"},
		{"block", "warped_nylium_side.png", "crimson", "warped_nylium_side.png"},
		{"block", "warped_roots.png", "crimson", "warped_roots.png"},
		//{"block", "stripped_warped_stem.png", "crimson", "warped_stem_stripped_side.png"},  //unused?
		{"block", "stripped_warped_stem_top.png", "crimson", "warped_stem_stripped_top.png"},
		{"block", "warped_wart_block.png", "crimson", "warped_wart_block.png"},
		// -- mcl_deepslate
		{"block", "deepslate.png", "deepslate", "mcl_deepslate.png"},
		{"block", "deepslate_bricks.png", "deepslate", "mcl_deepslate_bricks.png"},
		{"block", "cracked_deepslate_bricks.png", "deepslate", "mcl_deepslate_bricks_cracked.png"},
		{"block", "chiseled_deepslate.png", "deepslate", "mcl_deepslate_chiseled.png"},
		{"block", "deepslate_coal_ore.png", "deepslate", "mcl_deepslate_coal_ore.png"},
		{"block", "cobbled_deepslate.png", "deepslate", "mcl_deepslate_cobbled.png"},
		{"block", "deepslate_copper_ore.png", "deepslate", "mcl_deepslate_copper_ore.png"},
		{"block", "deepslate_diamond_ore.png", "deepslate", "mcl_deepslate_diamond_ore.png"},
		{"block", "deepslate_emerald_ore.png", "deepslate", "mcl_deepslate_emerald_ore.png"},
		{"block", "deepslate_gold_ore.png", "deepslate", "mcl_deepslate_gold_ore.png"},
		{"block", "deepslate_iron_ore.png", "deepslate", "mcl_deepslate_iron_ore.png"},
		{"block", "deepslate_lapis_ore.png", "deepslate", "mcl_deepslate_lapis_ore.png"},
		{"block", "polished_deepslate.png", "deepslate", "mcl_deepslate_polished.png"},
		{"block", "deepslate_redstone_ore.png", "deepslate", "mcl_deepslate_redstone_ore.png"},
		{"block", "reinforced_deepslate_side.png", "deepslate", "mcl_deepslate_reinforced.png"},
		{"block", "reinforced_deepslate_bottom.png", "deepslate", "mcl_deepslate_reinforced_bottom.png"},
		{"block", "reinforced_deepslate_top.png", "deepslate", "mcl_deepslate_reinforced_top.png"},
		{"block", "deepslate_tiles.png", "deepslate", "mcl_deepslate_tiles.png"},
		{"block", "cracked_deepslate_tiles.png", "deepslate", "mcl_deepslate_tiles_cracked.png"},
		{"block", "deepslate_top.png", "deepslate", "mcl_deepslate_top.png"},
		{"block", "tuff.png", "deepslate", "mcl_deepslate_tuff.png"},
		{"block", "tuff_bricks.png", "deepslate", "mcl_deepslate_tuff_bricks.png"},
		{"block", "chiseled_tuff.png", "deepslate", "mcl_deepslate_tuff_chiseled.png"},
		{"block", "chiseled_tuff_bricks.png", "deepslate", "mcl_deepslate_tuff_chiseled_bricks.png"},
		{"block", "chiseled_tuff_bricks_top.png", "deepslate", "mcl_deepslate_tuff_chiseled_bricks_top.png"},
		{"block", "chiseled_tuff_top.png", "deepslate", "mcl_deepslate_tuff_chiseled_top.png"},
		{"block", "polished_tuff.png", "deepslate", "mcl_deepslate_tuff_polished.png"},
		// -- mcl_doors
		{"item", "iron_door.png", "doors", "doors_item_steel.png"},
		{"item", "oak_door.png", "doors", "doors_item_wood.png"},
		{"block", "oak_trapdoor.png", "doors", "doors_trapdoor.png"},
		{"block", "oak_trapdoor.png", "doors", "doors_trapdoor_side.png"}, //
		{"block", "iron_trapdoor.png", "doors", "doors_trapdoor_steel.png"},
		{"block", "iron_trapdoor.png", "doors", "doors_trapdoor_steel_side.png"}, //
		{"item", "acacia_door.png", "doors", "mcl_doors_door_acacia.png"},
		//{"block", "acacia_door_bottom.png", "doors", "mcl_doors_door_acacia_lower.png"},
		{"block", "acacia_door_bottom.png", "doors", "mcl_doors_door_acacia_side_lower.png"}, //
		{"block", "acacia_door_top.png", "doors", "mcl_doors_door_acacia_side_upper.png"},    //
		//{"block", "acacia_door_top.png", "doors", "mcl_doors_door_acacia_upper.png"},
		{"item", "birch_door.png", "doors", "mcl_doors_door_birch.png"},
		//{"block", "birch_door_bottom.png", "doors", "mcl_doors_door_birch_lower.png"},
		{"block", "birch_door_bottom.png", "doors", "mcl_doors_door_birch_side_lower.png"}, //
		{"block", "birch_door_top.png", "doors", "mcl_doors_door_birch_side_upper.png"},    //
		//{"block", "birch_door_top.png", "doors", "mcl_doors_door_birch_upper.png"},
		{"item", "dark_oak_door.png", "doors", "mcl_doors_door_dark_oak.png"},
		//{"block", "dark_oak_door_bottom.png", "doors", "mcl_doors_door_dark_oak_lower.png"},
		{"block", "dark_oak_door_bottom.png", "doors", "mcl_doors_door_dark_oak_side_lower.png"}, //
		{"block", "dark_oak_door_top.png", "doors", "mcl_doors_door_dark_oak_side_upper.png"},    //
		//{"block", "dark_oak_door_top.png", "doors", "mcl_doors_door_dark_oak_upper.png"},
		{"block", "iron_door_bottom.png", "doors", "mcl_doors_door_iron_lower.png"},
		{"block", "iron_door_bottom.png", "doors", "mcl_doors_door_iron_side_lower.png"}, //
		{"block", "iron_door_top.png", "doors", "mcl_doors_door_iron_side_upper.png"},    //
		{"block", "iron_door_top.png", "doors", "mcl_doors_door_iron_upper.png"},
		{"item", "jungle_door.png", "doors", "mcl_doors_door_jungle.png"},
		//{"block", "jungle_door_bottom.png", "doors", "mcl_doors_door_jungle_lower.png"},
		{"block", "jungle_door_bottom.png", "doors", "mcl_doors_door_jungle_side_lower.png"}, //
		{"block", "jungle_door_top.png", "doors", "mcl_doors_door_jungle_side_upper.png"},    //
		//{"block", "jungle_door_top.png", "doors", "mcl_doors_door_jungle_upper.png"},
		{"item", "spruce_door.png", "doors", "mcl_doors_door_spruce.png"},
		//{"block", "spruce_door_bottom.png", "doors", "mcl_doors_door_spruce_lower.png"},
		{"block", "spruce_door_bottom.png", "doors", "mcl_doors_door_spruce_side_lower.png"}, //
		{"block", "spruce_door_top.png", "doors", "mcl_doors_door_spruce_side_upper.png"},    //
		//{"block", "spruce_door_top.png", "doors", "mcl_doors_door_spruce_upper.png"},
		//{"block", "oak_door_bottom.png", "doors", "mcl_doors_door_wood_lower.png"},
		{"block", "oak_door_bottom.png", "doors", "mcl_doors_door_wood_side_lower.png"},
		{"block", "oak_door_top.png", "doors", "mcl_doors_door_wood_side_upper.png"},
		//{"block", "oak_door_top.png", "doors", "mcl_doors_door_wood_upper.png"},
		{"block", "acacia_trapdoor.png", "doors", "mcl_doors_trapdoor_acacia.png"},
		{"block", "acacia_trapdoor.png", "doors", "mcl_doors_trapdoor_acacia_side.png"},
		{"block", "birch_trapdoor.png", "doors", "mcl_doors_trapdoor_birch.png"},
		{"block", "birch_trapdoor.png", "doors", "mcl_doors_trapdoor_birch_side.png"},
		{"block", "dark_oak_trapdoor.png", "doors", "mcl_doors_trapdoor_dark_oak.png"},
		{"block", "dark_oak_trapdoor.png", "doors", "mcl_doors_trapdoor_dark_oak_side.png"},
		{"block", "jungle_trapdoor.png", "doors", "mcl_doors_trapdoor_jungle.png"},
		{"block", "jungle_trapdoor.png", "doors", "mcl_doors_trapdoor_jungle_side.png"},
		{"block", "spruce_trapdoor.png", "doors", "mcl_doors_trapdoor_spruce.png"},
		{"block", "spruce_trapdoor.png", "doors", "mcl_doors_trapdoor_spruce_side.png"},
		// -- mcl_dyes
		// -- mcl_enchanting
		// -- mcl_end
		// -- mcl_farming
		{"item", "bread.png", "farming", "farming_bread.png"},
		{"item", "carrot.png", "farming", "farming_carrot.png"},
		{"block", "carrots_stage0.png", "farming", "farming_carrot_1.png"},
		{"block", "carrots_stage1.png", "farming", "farming_carrot_2.png"},
		{"block", "carrots_stage2.png", "farming", "farming_carrot_3.png"},
		{"block", "carrots_stage3.png", "farming", "farming_carrot_4.png"},
		{"item", "golden_carrot.png", "farming", "farming_carrot_gold.png"},
		{"item", "cookie.png", "farming", "farming_cookie.png"},
		{"item", "melon_slice.png", "farming", "farming_melon.png"},
		{"block", "melon_side.png", "farming", "farming_melon_side.png"},
		{"block", "melon_top.png", "farming", "farming_melon_top.png"},
		{"item", "potato.png", "farming", "farming_potato.png"},
		{"item", "baked_potato.png", "farming", "farming_potato_baked.png"},
		{"item", "poisonous_potato.png", "farming", "farming_potato_poison.png"},
		{"block", "carved_pumpkin.png", "farming", "farming_pumpkin_face.png"},
		{"block", "jack_o_lantern.png", "farming", "farming_pumpkin_face_light.png"},
		{"block", "pumpkin_side.png", "farming", "farming_pumpkin_side.png"},
		{"block", "pumpkin_top.png", "farming", "farming_pumpkin_top.png"},
		{"item", "diamond_hoe.png", "farming", "farming_tool_diamondhoe.png"},
		{"item", "golden_hoe.png", "farming", "farming_tool_goldhoe.png"},
		{"item", "netherite_hoe.png", "farming", "farming_tool_netheritehoe.png"},
		{"item", "iron_hoe.png", "farming", "farming_tool_steelhoe.png"},
		{"item", "stone_hoe.png", "farming", "farming_tool_stonehoe.png"},
		{"item", "wooden_hoe.png", "farming", "farming_tool_woodhoe.png"},
		{"item", "wheat.png", "farming", "farming_wheat_harvested.png"},
		{"item", "beetroot.png", "farming", "mcl_farming_beetroot.png"},
		{"block", "beetroots_stage0.png", "farming", "mcl_farming_beetroot_0.png"},
		{"block", "beetroots_stage1.png", "farming", "mcl_farming_beetroot_1.png"},
		{"block", "beetroots_stage2.png", "farming", "mcl_farming_beetroot_2.png"},
		{"block", "beetroots_stage3.png", "farming", "mcl_farming_beetroot_3.png"},
		{"item", "beetroot_seeds.png", "farming", "mcl_farming_beetroot_seeds.png"},
		{"item", "beetroot_soup.png", "farming", "mcl_farming_beetroot_soup.png"},
		{"block", "farmland.png", "farming", "mcl_farming_farmland_dry.png"},
		{"block", "farmland_moist.png", "farming", "mcl_farming_farmland_wet.png"},
		{"block", "hay_block_side.png", "farming", "mcl_farming_hayblock_side.png"},
		{"block", "hay_block_top.png", "farming", "mcl_farming_hayblock_top.png"},
		{"item", "melon_seeds.png", "farming", "mcl_farming_melon_seeds.png"},
		{"block", "attached_melon_stem.png", "farming", "mcl_farming_melon_stem_connected.png"},
		{"block", "melon_stem.png", "farming", "mcl_farming_melon_stem_disconnected.png"},
		{"block", "potatoes_stage0.png", "farming", "mcl_farming_potatoes_stage_0.png"},
		{"block", "potatoes_stage1.png", "farming", "mcl_farming_potatoes_stage_1.png"},
		{"block", "potatoes_stage2.png", "farming", "mcl_farming_potatoes_stage_2.png"},
		{"block", "potatoes_stage3.png", "farming", "mcl_farming_potatoes_stage_3.png"},
		//{"block", ".png", "farming", "mcl_farming_pumpkin_face.png"},  //special attention
		{"misc", "pumpkinblur.png", "farming", "mcl_farming_pumpkin_hud.png"}, //pls test
		{"item", "pumpkin_pie.png", "farming", "mcl_farming_pumpkin_pie.png"},
		{"item", "pumpkin_seeds.png", "farming", "mcl_farming_pumpkin_seeds.png"},
		{"block", "attached_pumpkin_stem.png", "farming", "mcl_farming_pumpkin_stem_connected.png"},
		{"block", "pumpkin_stem.png", "farming", "mcl_farming_pumpkin_stem_disconnected.png"},
		{"item", "sweet_berries.png", "farming", "mcl_farming_sweet_berry.png"},
		{"block", "sweet_berry_bush_stage0.png", "farming", "mcl_farming_sweet_berry_bush_0.png"},
		{"block", "sweet_berry_bush_stage1.png", "farming", "mcl_farming_sweet_berry_bush_1.png"},
		{"block", "sweet_berry_bush_stage2.png", "farming", "mcl_farming_sweet_berry_bush_2.png"},
		{"block", "sweet_berry_bush_stage3.png", "farming", "mcl_farming_sweet_berry_bush_3.png"},
		{"item", "wheat_seeds.png", "farming", "mcl_farming_wheat_seeds.png"},
		{"block", "wheat_stage0.png", "farming", "mcl_farming_wheat_stage_0.png"},
		{"block", "wheat_stage1.png", "farming", "mcl_farming_wheat_stage_1.png"},
		{"block", "wheat_stage2.png", "farming", "mcl_farming_wheat_stage_2.png"},
		{"block", "wheat_stage3.png", "farming", "mcl_farming_wheat_stage_3.png"},
		{"block", "wheat_stage4.png", "farming", "mcl_farming_wheat_stage_4.png"},
		{"block", "wheat_stage5.png", "farming", "mcl_farming_wheat_stage_5.png"},
		{"block", "wheat_stage6.png", "farming", "mcl_farming_wheat_stage_6.png"},
		{"block", "wheat_stage7.png", "farming", "mcl_farming_wheat_stage_7.png"},
		// -- mcl_fences
		// -- mcl_fire
		// -- mcl_fireworks
		// -- mcl_fishing
		{"item", "cod_bucket.png", "fishing", "cod_bucket.png"},
		{"entity", "fishing_hook.png", "fishing", "mcl_fishing_bobber.png"},
		{"item", "tropical_fish.png", "fishing", "mcl_fishing_clownfish_raw.png"},
		{"item", "cooked_cod.png", "fishing", "mcl_fishing_fish_cooked.png"},
		{"item", "fishing_rod.png", "fishing", "mcl_fishing_fishing_rod.png"},
		{"item", "cod.png", "fishing", "mcl_fishing_fish_raw.png"},
		{"item", "pufferfish.png", "fishing", "mcl_fishing_pufferfish_raw.png"},
		{"item", "cooked_salmon.png", "fishing", "mcl_fishing_salmon_cooked.png"},
		{"item", "salmon.png", "fishing", "mcl_fishing_salmon_raw.png"},
		{"item", "pufferfish_bucket.png", "fishing", "pufferfish_bucket.png"},
		{"item", "salmon_bucket.png", "fishing", "salmon_bucket.png"},
		{"item", "tropical_fish_bucket.png", "fishing", "tropical_fish_bucket.png"},
		// -- mcl_fletching_table
		//{"block", ".png", "fletching_table", "fletching_table_bottom.png"}, //no match
		{"block", "fletching_table_front.png", "fletching_table", "fletching_table_front.png"},
		{"block", "fletching_table_side.png", "fletching_table", "fletching_table_side.png"},
		{"block", "fletching_table_top.png", "fletching_table", "fletching_table_top.png"},
		// -- mcl_flowerpots
		//{"block", ".png", "flowerpots", "mcl_flowerpots_cactus.png"},    //special attention
		//{"block", ".png", "flowerpots", "mcl_flowerpots_flowerpot.png"}, //special attention
		{"item", "flower_pot.png", "flowerpots", "mcl_flowerpots_flowerpot_inventory.png"},
		// -- mcl_flowers
		{"block", "dandelion.png", "flowers", "flowers_dandelion_yellow.png"},
		{"block", "orange_tulip.png", "flowers", "flowers_tulip.png"},
		//{"block", "lily_pad.png", "flowers", "flowers_waterlily.png"}, //special attention
		{"block", "allium.png", "flowers", "mcl_flowers_allium.png"},
		{"block", "azure_bluet.png", "flowers", "mcl_flowers_azure_bluet.png"},
		{"block", "blue_orchid.png", "flowers", "mcl_flowers_blue_orchid.png"},
		{"block", "cornflower.png", "flowers", "mcl_flowers_cornflower.png"},
		{"block", "large_fern_bottom.png", "flowers", "mcl_flowers_double_plant_fern_bottom.png"},
		//{"block", "large_fern_top.png", "flowers", "mcl_flowers_double_plant_fern_inv.png"}, //special attention
		{"block", "large_fern_top.png", "flowers", "mcl_flowers_double_plant_fern_top.png"},
		{"block", "tall_grass_bottom.png", "flowers", "mcl_flowers_double_plant_grass_bottom.png"},
		//{"block", "tall_grass_top.png", "flowers", "mcl_flowers_double_plant_grass_inv.png"}, //special attention
		{"block", "tall_grass_top.png", "flowers", "mcl_flowers_double_plant_grass_top.png"},
		{"block", "peony_bottom.png", "flowers", "mcl_flowers_double_plant_paeonia_bottom.png"},
		{"block", "peony_top.png", "flowers", "mcl_flowers_double_plant_paeonia_top.png"},
		{"block", "rose_bush_bottom.png", "flowers", "mcl_flowers_double_plant_rose_bottom.png"},
		{"block", "rose_bush_top.png", "flowers", "mcl_flowers_double_plant_rose_top.png"},
		{"block", "sunflower_back.png", "flowers", "mcl_flowers_double_plant_sunflower_back.png"},
		{"block", "sunflower_bottom.png", "flowers", "mcl_flowers_double_plant_sunflower_bottom.png"},
		{"block", "sunflower_front.png", "flowers", "mcl_flowers_double_plant_sunflower_front.png"},
		{"block", "sunflower_top.png", "flowers", "mcl_flowers_double_plant_sunflower_top.png"},
		{"block", "lilac_bottom.png", "flowers", "mcl_flowers_double_plant_syringa_bottom.png"},
		{"block", "lilac_top.png", "flowers", "mcl_flowers_double_plant_syringa_top.png"},
		{"block", "fern.png", "flowers", "mcl_flowers_fern.png"},
		//{"block", "fern.png", "flowers", "mcl_flowers_fern_inv.png"}, //special attention
		{"block", "lily_of_the_valley.png", "flowers", "mcl_flowers_lily_of_the_valley.png"},
		{"block", "oxeye_daisy.png", "flowers", "mcl_flowers_oxeye_daisy.png"},
		{"block", "poppy.png", "flowers", "mcl_flowers_poppy.png"},
		{"block", "short_grass.png", "flowers", "mcl_flowers_tallgrass.png"},
		//{"block", "short_grass.png", "flowers", "mcl_flowers_tallgrass_inv.png"}, //special attention
		{"block", "pink_tulip.png", "flowers", "mcl_flowers_tulip_pink.png"},
		{"block", "red_tulip.png", "flowers", "mcl_flowers_tulip_red.png"},
		{"block", "white_tulip.png", "flowers", "mcl_flowers_tulip_white.png"},
		{"block", "wither_rose.png", "flowers", "mcl_flowers_wither_rose.png"},
		// -- mcl_furnaces
		{"block", "furnace_top.png", "furnaces", "default_furnace_bottom.png"}, //Minecraft does not have this texture.
		{"block", "furnace_front.png", "furnaces", "default_furnace_front.png"},
		{"block", "furnace_front_on.png", "furnaces", "default_furnace_front_active.png"},
		{"block", "furnace_side.png", "furnaces", "default_furnace_side.png"},
		{"block", "furnace_top.png", "furnaces", "default_furnace_top.png"},
		// -- mcl_grindstone
		// -- mcl_heads
		// -- mcl_honey
		{"block", "honey_block_bottom.png", "honey", "mcl_honey_block_bottom.png"},
		{"block", "honey_block_side.png", "honey", "mcl_honey_block_side.png"},
		{"block", "honey_block_top.png", "honey", "mcl_honey_block_top.png"},
		{"item", "honey_bottle.png", "honey", "mcl_honey_honey_bottle.png"},
		{"item", "honeycomb.png", "honey", "mcl_honey_honeycomb.png"},
		{"block", "honeycomb_block.png", "honey", "mcl_honey_honeycomb_block.png"},
		// -- mcl_hoppers
		{"block", "hopper_inside.png", "hopper", "mcl_hoppers_hopper_inside.png"},
		{"block", "hopper_outside.png", "hopper", "mcl_hoppers_hopper_outside.png"},
		{"block", "hopper_top.png", "hopper", "mcl_hoppers_hopper_top.png"},
		{"item", "hopper.png", "hopper", "mcl_hoppers_item.png"},
		// -- mcl_itemframes
		// -- mcl_jukebox
		// -- mcl_lanterns
		// -- mcl_lectern
		// -- mcl_lightning_rods
		// -- mcl_loom
		// -- mcl_lush_caves
		// -- mcl_mangrove
		//{"block", "mangrove_door_bottom.png", "mangrove", "mcl_mangrove_door_bottom.png"}, //flipped
		{"item", "mangrove_door.png", "mangrove", "mcl_mangrove_doors.png"},
		//{"block", "mangrove_door_top.png", "mangrove", "mcl_mangrove_door_top.png"}, //flipped
		{"block", "mangrove_planks.png", "mangrove", "mcl_mangrove_fence.png"},
		{"block", "mangrove_planks.png", "mangrove", "mcl_mangrove_fence_gate.png"},
		{"block", "mangrove_leaves.png", "mangrove", "mcl_mangrove_leaves.png"},
		{"block", "mangrove_log.png", "mangrove", "mcl_mangrove_log.png"},
		{"block", "mangrove_log_top.png", "mangrove", "mcl_mangrove_log_top.png"},
		{"block", "mangrove_planks.png", "mangrove", "mcl_mangrove_planks.png"},
		{"block", "mangrove_propagule.png", "mangrove", "mcl_mangrove_propagule.png"},
		{"block", "mangrove_propagule_hanging.png", "mangrove", "mcl_mangrove_propagule_hanging.png"},
		{"item", "mangrove_propagule.png", "mangrove", "mcl_mangrove_propagule_item.png"},
		{"block", "mangrove_roots_side.png", "mangrove", "mcl_mangrove_roots_side.png"},
		{"block", "mangrove_roots_top.png", "mangrove", "mcl_mangrove_roots_top.png"},
		{"block", "mangrove_trapdoor.png", "mangrove", "mcl_mangrove_trapdoor.png"},
		{"block", "mangrove_trapdoor.png", "mangrove", "mcl_mangrove_trapdoor_side.png"},
		{"block", "stripped_mangrove_log.png", "mangrove", "mcl_stripped_mangrove_log_side.png"},
		{"block", "stripped_mangrove_log_top.png", "mangrove", "mcl_stripped_mangrove_log_top.png"},
		// -- mcl_maps
		// -- mcl_mobitems
		// -- mcl_mobspawners
		// -- mcl_monster_eggs
		// -- mcl_mud
		{"block", "mud.png", "mud", "mcl_mud.png"},
		{"block", "mud_bricks.png", "mud", "mcl_mud_bricks.png"},
		{"block", "packed_mud.png", "mud", "mcl_mud_packed_mud.png"},
		// -- mcl_mushrooms
		{"block", "brown_mushroom.png", "mushrooms", "farming_mushroom_brown.png"},
		{"block", "red_mushroom.png", "mushrooms", "farming_mushroom_red.png"},
		{"item", "mushroom_stew.png", "mushrooms", "farming_mushroom_stew.png"},
		{"block", "mushroom_block_inside.png", "mushrooms", "mcl_mushrooms_mushroom_block_inside.png"},
		{"block", "brown_mushroom_block.png", "mushrooms", "mcl_mushrooms_mushroom_block_skin_brown.png"},
		{"block", "red_mushroom_block.png", "mushrooms", "mcl_mushrooms_mushroom_block_skin_red.png"},
		{"block", "mushroom_stem.png", "mushrooms", "mcl_mushrooms_mushroom_block_skin_stem.png"},
		// -- mcl_nether
		//TODO
		{"block", "ancient_debris_side.png", "nether", "mcl_nether_ancient_debris_side.png"},
		{"block", "ancient_debris_top.png", "nether", "mcl_nether_ancient_debris_top.png"},
		{"block", "chiseled_nether_bricks.png", "nether", "mcl_nether_chiseled_nether_bricks.png"},
		{"block", "cracked_nether_bricks.png", "nether", "mcl_nether_cracked_nether_bricks.png"},
		{"block", "glowstone.png", "nether", "mcl_nether_glowstone.png"},
		{"item", "glowstone_dust.png", "nether", "mcl_nether_glowstone_dust.png"},
		{"block", "nether_gold_ore.png", "nether", "mcl_nether_gold_ore.png"},
		{"block", "magma.png", "nether", "mcl_nether_magma.png"}, //is animated
		{"block", "nether_bricks.png", "nether", "mcl_nether_nether_brick.png"},
		{"item", "nether_brick.png", "nether", "mcl_nether_netherbrick.png"},
		{"block", "netherite_block.png", "nether", "mcl_nether_netheriteblock.png"},
		{"item", "netherite_ingot.png", "nether", "mcl_nether_netherite_ingot.png"},
		{"item", "netherite_scrap.png", "nether", "mcl_nether_netherite_scrap.png"},
		{"item", "netherite_upgrade_smithing_template.png", "nether", "mcl_nether_netherite_ugrade_template.png"},
		{"block", "netherrack.png", "nether", "mcl_nether_netherrack.png"},
		{"item", "nether_wart.png", "nether", "mcl_nether_nether_wart.png"},
		{"block", "nether_wart_block.png", "nether", "mcl_nether_nether_wart_block.png"},
		{"block", "nether_wart_stage0.png", "nether", "mcl_nether_nether_wart_stage_0.png"},
		{"block", "nether_wart_stage1.png", "nether", "mcl_nether_nether_wart_stage_1.png"},
		{"block", "nether_wart_stage2.png", "nether", "mcl_nether_nether_wart_stage_2.png"},
		{"item", "quartz.png", "nether", "mcl_nether_quartz.png"},
		{"block", "quartz_block_bottom.png", "nether", "mcl_nether_quartz_block_bottom.png"},
		{"block", "quartz_block_side.png", "nether", "mcl_nether_quartz_block_side.png"},
		{"block", "quartz_block_top.png", "nether", "mcl_nether_quartz_block_top.png"},
		{"block", "chiseled_quartz_block.png", "nether", "mcl_nether_quartz_chiseled_side.png"},
		{"block", "chiseled_quartz_block_top.png", "nether", "mcl_nether_quartz_chiseled_top.png"},
		{"block", "nether_quartz_ore.png", "nether", "mcl_nether_quartz_ore.png"},
		{"block", "quartz_pillar.png", "nether", "mcl_nether_quartz_pillar_side.png"},
		{"block", "quartz_pillar_top.png", "nether", "mcl_nether_quartz_pillar_top.png"},
		{"block", "red_nether_bricks.png", "nether", "mcl_nether_red_nether_brick.png"},
		{"block", "soul_sand.png", "nether", "mcl_nether_soul_sand.png"},
		// -- mcl_ocean
		// -- mcl_panes
		// -- mcl_portals
		// -- mcl_potions
		// -- mcl_pottery_sherds
		// -- mcl_raw_ores
		{"item", "raw_gold.png", "raw_ores", "mcl_raw_ores_raw_gold.png"},
		{"block", "raw_gold_block.png", "raw_ores", "mcl_raw_ores_raw_gold_block.png"},
		{"item", "raw_iron.png", "raw_ores", "mcl_raw_ores_raw_iron.png"},
		{"block", "raw_iron_block.png", "raw_ores", "mcl_raw_ores_raw_iron_block.png"},
		// -- mcl_sculk
		// -- mcl_shields
		// -- mcl_signs
		// -- mcl_smithing_table
		{"block", "smithing_table_bottom.png", "smithing_table", "mcl_smithing_table_bottom.png"},
		{"block", "smithing_table_front.png", "smithing_table", "mcl_smithing_table_front.png"},
		{"item", "empty_slot_smithing_template_armor_trim.png", "smithing_table", "mcl_smithing_table_inventory_trim_bg.png"},
		{"block", "smithing_table_side.png", "smithing_table", "mcl_smithing_table_side.png"},
		{"block", "smithing_table_top.png", "smithing_table", "mcl_smithing_table_top.png"},
		// -- mcl_smoker
		{"block", "smoker_bottom.png", "smoker", "smoker_bottom.png"},
		{"block", "smoker_front.png", "smoker", "smoker_front.png"},
		{"block", "smoker_front_on.png", "smoker", "smoker_front_on.png"},
		{"block", "smoker_side.png", "smoker", "smoker_side.png"},
		{"block", "smoker_top.png", "smoker", "smoker_top.png"},
		// -- mcl_sponges
		// -- mcl_spyglass
		// -- mcl_stairs
		// -- mcl_stonecutter
		{"block", "stonecutter_bottom.png", "stonecutter", "mcl_stonecutter_bottom.png"},
		//{"block", "stonecutter_saw.png", "stonecutter", "mcl_stonecutter_saw.png"},  //special attention
		{"block", "stonecutter_side.png", "stonecutter", "mcl_stonecutter_side.png"},
		{"block", "stonecutter_top.png", "stonecutter", "mcl_stonecutter_top.png"},
		// -- mcl_sus_nodes
		// -- mcl_sus_stew
		{"item", "suspicious_stew.png", "sus_stew", "sus_stew.png"},
		// -- mcl_throwing
		// -- mcl_tnt
		{"block", "tnt_bottom.png", "tnt", "default_tnt_bottom.png"},
		{"block", "tnt_side.png", "tnt", "default_tnt_side.png"},
		{"block", "tnt_top.png", "tnt", "default_tnt_top.png"},
		// -- mcl_tools
		{"item", "diamond_axe.png", "tools", "default_tool_diamondaxe.png"},
		{"item", "diamond_pickaxe.png", "tools", "default_tool_diamondpick.png"},
		{"item", "diamond_shovel.png", "tools", "default_tool_diamondshovel.png"},
		{"item", "diamond_sword.png", "tools", "default_tool_diamondsword.png"},
		{"item", "golden_axe.png", "tools", "default_tool_goldaxe.png"},
		{"item", "golden_pickaxe.png", "tools", "default_tool_goldpick.png"},
		{"item", "golden_shovel.png", "tools", "default_tool_goldshovel.png"},
		{"item", "golden_sword.png", "tools", "default_tool_goldsword.png"},
		{"item", "netherite_axe.png", "tools", "default_tool_netheriteaxe.png"},
		{"item", "netherite_pickaxe.png", "tools", "default_tool_netheritepick.png"},
		{"item", "netherite_shovel.png", "tools", "default_tool_netheriteshovel.png"},
		{"item", "netherite_sword.png", "tools", "default_tool_netheritesword.png"},
		{"item", "shears.png", "tools", "default_tool_shears.png"},
		{"item", "iron_axe.png", "tools", "default_tool_steelaxe.png"},
		{"item", "iron_pickaxe.png", "tools", "default_tool_steelpick.png"},
		{"item", "iron_shovel.png", "tools", "default_tool_steelshovel.png"},
		{"item", "iron_sword.png", "tools", "default_tool_steelsword.png"},
		{"item", "stone_axe.png", "tools", "default_tool_stoneaxe.png"},
		{"item", "stone_pickaxe.png", "tools", "default_tool_stonepick.png"},
		{"item", "stone_shovel.png", "tools", "default_tool_stoneshovel.png"},
		{"item", "stone_sword.png", "tools", "default_tool_stonesword.png"},
		{"item", "wooden_axe.png", "tools", "default_tool_woodaxe.png"},
		{"item", "wooden_pickaxe.png", "tools", "default_tool_woodpick.png"},
		{"item", "wooden_shovel.png", "tools", "default_tool_woodshovel.png"},
		{"item", "wooden_sword.png", "tools", "default_tool_woodsword.png"},
		//{"block", ".png", "tools", "mcl_tools_heavy_core_bottom.png"},  //special attention
		//{"block", ".png", "tools", "mcl_tools_heavy_core_side.png"},  //special attention
		//{"block", ".png", "tools", "mcl_tools_heavy_core_top.png"},  //special attention
		{"item", "mace.png", "tools", "mcl_tools_mace.png"},
		// -- mcl_torches
		{"block", "torch.png", "torches", "default_torch_on_floor.png"},
		{"block", "torch.png", "torches", "default_torch_on_floor_animated.png"},
		// -- mcl_totems
		// -- mcl_trees
		// -- mcl_walls
		// -- mcl_wool
		{"block", "light_blue_wool.png", "wool", "mcl_wool_light_blue.png"},
		{"block", "lime_wool.png", "wool", "mcl_wool_lime.png"},
		{"block", "black_wool.png", "wool", "wool_black.png"},
		{"block", "blue_wool.png", "wool", "wool_blue.png"},
		{"block", "brown_wool.png", "wool", "wool_brown.png"},
		{"block", "cyan_wool.png", "wool", "wool_cyan.png"},
		{"block", "green_wool.png", "wool", "wool_dark_green.png"},
		{"block", "gray_wool.png", "wool", "wool_dark_grey.png"},
		{"block", "light_gray_wool.png", "wool", "wool_grey.png"},
		{"block", "magenta_wool.png", "wool", "wool_magenta.png"},
		{"block", "orange_wool.png", "wool", "wool_orange.png"},
		{"block", "pink_wool.png", "wool", "wool_pink.png"},
		{"block", "red_wool.png", "wool", "wool_red.png"},
		{"block", "purple_wool.png", "wool", "wool_violet.png"},
		{"block", "white_wool.png", "wool", "wool_white.png"},
		{"block", "yellow_wool.png", "wool", "wool_yellow.png"},
		// -- mclx_core
		// -- mclx_fences
		// -- mclx_stairs
		// -- REDSTONE
		// -- screwdriver
		//{"block", ".png", "", ".png"},
	}
	return equivalents
}
