package main

import ()

func CraftPaths() map[string]string {
	craftPaths := map[string]string{
		"armor":    "/assets/minecraft/textures/models/armor/",
		"bed":      "/assets/minecraft/textures/entity/bed/",
		"block":    "/assets/minecraft/textures/block/",
		"item":     "/assets/minecraft/textures/item/",
		"particle": "/assets/minecraft/textures/particle/",
	}
	return craftPaths
}

func CloniaPaths() map[string]string {
	cloniaPaths := map[string]string{
		"amethyst":       "/ITEMS/mcl_amethyst/textures/",
		"anvils":         "/ITEMS/mcl_anvils/textures/",
		"armor_stand":    "/ITEMS/mcl_armor_stand/textures/",
		"bamboo":         "/ITEMS/mcl_bamboo/textures/",
		"barrels":        "/ITEMS/mcl_barrels/textures/",
		"beds":           "/ITEMS/mcl_beds/textures/",
		"blast_furnace":  "/ITEMS/mcl_blast_furnace/textures/",
		"cherry_blossom": "/ITEMS/mcl_cherry_blossom/textures/",
		"copper":         "/ITEMS/mcl_copper/textures/",
		"core":           "/ITEMS/mcl_core/textures/",
		"crafting_table": "/ITEMS/mcl_crafting_table/textures/",
		"crimson":        "/ITEMS/mcl_crimson/textures/",
		"deepslate":      "/ITEMS/mcl_deepslate/textures/",
		"doors":          "/ITEMS/mcl_doors/textures/",
		"furnaces":       "/ITEMS/mcl_furnaces/textures/",
		"mud":            "/ITEMS/mcl_mud/textures/",
		"mushrooms":      "/ITEMS/mcl_mushrooms/textures/",
		"nether":         "/ITEMS/mcl_nether/textures/",
		"raw_ores":       "/ITEMS/mcl_raw_ores/textures/",
		"smithing_table": "/ITEMS/mcl_smithing_table/textures/",
		"smoker":         "/ITEMS/mcl_smoker/textures/",
		"stonecutter":    "/ITEMS/mcl_stonecutter/textures/",
		"sus_stew":       "/ITEMS/mcl_sus_stew/textures/",
		"tnt":            "/ITEMS/mcl_tnt/textures/",
		"tools":          "/ITEMS/mcl_tools/textures/",
		"torches":        "/ITEMS/mcl_torches/textures/",
		"wool":           "/ITEMS/mcl_wool/textures/",
		//"": "/ITEMS//textures/",
	}
	return cloniaPaths
}
func basicITEMS() [][4]string {
	equivalents := [][4]string{
		//// mcl_amethyst
		{"block", "amethyst_block", "amethyst", "mcl_amethyst_amethyst_block"},
		{"block", "large_amethyst_bud", "amethyst", "mcl_amethyst_amethyst_bud_large"},
		{"block", "medium_amethyst_bud", "amethyst", "mcl_amethyst_amethyst_bud_medium"},
		{"block", "small_amethyst_bud", "amethyst", "mcl_amethyst_amethyst_bud_small"},
		{"block", "amethyst_cluster", "amethyst", "mcl_amethyst_amethyst_cluster"},
		{"item", "amethyst_shard", "amethyst", "mcl_amethyst_amethyst_shard"},
		{"block", "budding_amethyst", "amethyst", "mcl_amethyst_budding_amethyst"},
		{"block", "calcite", "amethyst", "mcl_amethyst_calcite_block"},
		{"block", "tinted_glass", "amethyst", "mcl_amethyst_tinted_glass"},
		//// mcl_anvils
		{"block", "anvil", "anvils", "mcl_anvils_anvil_base"},
		{"block", "anvil", "anvils", "mcl_anvils_anvil_side"},
		//"mcl_anvils_anvil_top_damaged_0" handled elsewhere
		//"mcl_anvils_anvil_top_damaged_1" handled elsewhere
		//"mcl_anvils_anvil_top_damaged_2" handled elsewhere
		//// mcl_armor
		//// mcl_armor_stand
		{"item", "armor_stand", "armor_stand", "armor_stand"},
		//// mcl_bamboo
		//{"block", "", "bamboo", "mcl_bamboo_bamboo"},
		{"block", "bamboo_block", "bamboo", "mcl_bamboo_bamboo_block"},
		{"block", "stripped_bamboo_block", "bamboo", "mcl_bamboo_bamboo_block_stripped"},
		{"block", "bamboo_block_top", "bamboo", "mcl_bamboo_bamboo_bottom"},
		{"block", "stripped_bamboo_block_top", "bamboo", "mcl_bamboo_bamboo_bottom_stripped"},
		//{"block", "", "bamboo", "mcl_bamboo_bamboo_fpm"},    //special attention
		{"block", "bamboo_planks", "bamboo", "mcl_bamboo_bamboo_plank"},
		{"block", "bamboo_mosaic", "bamboo", "mcl_bamboo_bamboo_plank_mosaic"},
		{"item", "bamboo", "bamboo", "mcl_bamboo_bamboo_shoot"},
		//{"block", "", "bamboo", "mcl_bamboo_bamboo_sign"},    //special attention
		{"item", "bamboo_sign", "bamboo", "mcl_bamboo_bamboo_sign_wield"},
		//{"block", "bamboo_door_bottom", "bamboo", "mcl_bamboo_door_bottom"},
		//{"block", "bamboo_door_bottom", "bamboo", "mcl_bamboo_door_bottom_alt"},    //test
		//{"block", "bamboo_door_top", "bamboo", "mcl_bamboo_door_top"},
		//{"block", "bamboo_door_top", "bamboo", "mcl_bamboo_door_top_alt"},    //test
		{"item", "bamboo_door", "bamboo", "mcl_bamboo_door_wield"},
		//{"block", "", "bamboo", "mcl_bamboo_endcap"}, //what is this?
		{"block", "bamboo_fence_particle", "bamboo", "mcl_bamboo_fence_bamboo"},
		{"block", "bamboo_fence_gate_particle", "bamboo", "mcl_bamboo_fence_gate_bamboo"},
		////scaffolding is broken in vanilla Mineclonia right now.
		//{"block", "bamboo_stage0", "bamboo", "mcl_bamboo_flower_pot"},    //broken?
		{"block", "scaffolding_bottom", "bamboo", "mcl_bamboo_scaffolding_bottom"}, //broken?
		{"block", "scaffolding_side", "bamboo", "mcl_bamboo_scaffolding_side"},     //broken?
		{"block", "scaffolding_top", "bamboo", "mcl_bamboo_scaffolding_top"},
		{"block", "bamboo_trapdoor", "bamboo", "mcl_bamboo_trapdoor_side"},
		//// mcl_banners
		//// mcl_barrels
		{"block", "barrel_bottom", "barrels", "mcl_barrels_barrel_bottom"},
		{"block", "barrel_side", "barrels", "mcl_barrels_barrel_side"},
		{"block", "barrel_top", "barrels", "mcl_barrels_barrel_top"},
		{"block", "barrel_top_open", "barrels", "mcl_barrels_barrel_top_open"},
		//// mcl_beacons
		////// special attention
		//// mcl_beds
		{"bed", "black", "beds", "mcl_beds_bed_black"},
		//"mcl_beds_bed_black_inv"  no match
		{"bed", "blue", "beds", "mcl_beds_bed_blue"},
		//"mcl_beds_bed_blue_inv"  no match
		{"bed", "brown", "beds", "mcl_beds_bed_brown"},
		//"mcl_beds_bed_brown_inv"  no match
		{"bed", "cyan", "beds", "mcl_beds_bed_cyan"},
		//"mcl_beds_bed_cyan_inv"  no match
		{"bed", "green", "beds", "mcl_beds_bed_green"},
		//"mcl_beds_bed_green_inv"  no match
		{"bed", "gray", "beds", "mcl_beds_bed_grey"},
		//"mcl_beds_bed_grey_inv"  no match
		{"bed", "light_blue", "beds", "mcl_beds_bed_light_blue"},
		//"mcl_beds_bed_light_blue_inv"  no match
		{"bed", "lime", "beds", "mcl_beds_bed_lime"},
		//"mcl_beds_bed_lime_inv"  no match
		{"bed", "magenta", "beds", "mcl_beds_bed_magenta"},
		//"mcl_beds_bed_magenta_inv"  no match
		{"bed", "orange", "beds", "mcl_beds_bed_orange"},
		//"mcl_beds_bed_orange_inv"  no match
		{"bed", "pink", "beds", "mcl_beds_bed_pink"},
		//"mcl_beds_bed_pink_inv"  no match
		{"bed", "purple", "beds", "mcl_beds_bed_purple"},
		//"mcl_beds_bed_purple_inv"  no match
		{"bed", "red", "beds", "mcl_beds_bed_red"},
		//"mcl_beds_bed_red_inv"  no match
		{"bed", "light_gray", "beds", "mcl_beds_bed_silver"},
		//"mcl_beds_bed_silver_inv"  no match
		{"bed", "white", "beds", "mcl_beds_bed_white"},
		//"mcl_beds_bed_white_inv"  no match
		{"bed", "yellow", "beds", "mcl_beds_bed_yellow"},
		//"mcl_beds_bed_yellow_inv"  no match
		{"block", "respawn_anchor_bottom", "beds", "respawn_anchor_bottom"},
		{"block", "respawn_anchor_side0", "beds", "respawn_anchor_side0"},
		{"block", "respawn_anchor_side1", "beds", "respawn_anchor_side1"},
		{"block", "respawn_anchor_side2", "beds", "respawn_anchor_side2"},
		{"block", "respawn_anchor_side3", "beds", "respawn_anchor_side3"},
		{"block", "respawn_anchor_side4", "beds", "respawn_anchor_side4"},
		{"block", "respawn_anchor_top_off", "beds", "respawn_anchor_top_off"},
		{"block", "respawn_anchor_top", "beds", "respawn_anchor_top_on"},
		//// mcl_beehives
		//// mcl_bells
		//// mcl_blackstone
		//// mcl_blast_furnace
		{"block", "blast_furnace_front", "blast_furnace", "blast_furnace_front"},
		{"block", "blast_furnace_front_on", "blast_furnace", "blast_furnace_front_on"},
		{"block", "blast_furnace_side", "blast_furnace", "blast_furnace_side"},
		{"block", "blast_furnace_top", "blast_furnace", "blast_furnace_top"},
		//// mcl_bone_meal
		//// mcl_books
		//// mcl_bows
		//// mcl_brewing
		//// mcl_buckets
		//// mcl_cake
		//// mcl_campfires
		//// mcl_cartography_table
		//// mcl_cauldrons
		//// mcl_cherry_blossom
		//{"block", "cherry_door_bottom", "cherry_blossom", "mcl_cherry_blossom_door_bottom"}, //flipped
		{"block", "cherry_door_bottom", "cherry_blossom", "mcl_cherry_blossom_door_bottom_bottompart"},
		{"block", "cherry_door_bottom", "cherry_blossom", "mcl_cherry_blossom_door_bottom_side"},
		{"item", "cherry_door", "cherry_blossom", "mcl_cherry_blossom_door_inv"},
		//{"block", "cherry_door_top", "cherry_blossom", "mcl_cherry_blossom_door_top"}, //flipped
		{"block", "cherry_door_top", "cherry_blossom", "mcl_cherry_blossom_door_top_side"},
		{"block", "cherry_door_top", "cherry_blossom", "mcl_cherry_blossom_door_top_toppart"},
		{"block", "cherry_leaves", "cherry_blossom", "mcl_cherry_blossom_leaves"},
		{"block", "cherry_log", "cherry_blossom", "mcl_cherry_blossom_log"},
		{"block", "stripped_cherry_log", "cherry_blossom", "mcl_cherry_blossom_log_stripped"},
		{"block", "cherry_log_top", "cherry_blossom", "mcl_cherry_blossom_log_top"},
		{"block", "stripped_cherry_log_top", "cherry_blossom", "mcl_cherry_blossom_log_top_stripped"},
		//{"block", "", "cherry_blossom", "mcl_cherry_blossom_particle"},
		//{"block", "", "cherry_blossom", "mcl_cherry_blossom_particle_1"},
		//{"block", "", "cherry_blossom", "mcl_cherry_blossom_particle_2"},
		//{"block", "", "cherry_blossom", "mcl_cherry_blossom_particle_3"},
		{"block", "pink_petals", "cherry_blossom", "mcl_cherry_blossom_pink_petals"},
		{"block", "pink_petals", "cherry_blossom", "mcl_cherry_blossom_pink_petals_inv"},
		{"block", "cherry_planks", "cherry_blossom", "mcl_cherry_blossom_planks"},
		{"block", "cherry_sapling", "cherry_blossom", "mcl_cherry_blossom_sapling"},
		{"block", "cherry_trapdoor", "cherry_blossom", "mcl_cherry_blossom_trapdoor"},
		{"block", "cherry_trapdoor", "cherry_blossom", "mcl_cherry_blossom_trapdoor_side"},
		//// mcl_chests
		//// mcl_clock
		//// mcl_cocoas
		//// mcl_colorblocks
		//// mcl_compass
		//// mcl_composters
		//// mcl_conduits
		//// mcl_copper
		//"mcl_copper_anti_oxidation_particle" //no match?
		{"block", "copper_block", "copper", "mcl_copper_block"},
		{"block", "copper_bulb", "copper", "mcl_copper_block_bulb_off"},
		{"block", "copper_bulb_lit", "copper", "mcl_copper_block_bulb_on"},
		{"block", "copper_bulb_powered", "copper", "mcl_copper_block_bulb_powered_off"},
		{"block", "copper_bulb_lit_powered", "copper", "mcl_copper_block_bulb_powered_on"},
		{"block", "chiseled_copper", "copper", "mcl_copper_block_chiseled"},
		{"block", "cut_copper", "copper", "mcl_copper_block_cut"},
		{"block", "copper_grate", "copper", "mcl_copper_block_grate"},
		{"block", "raw_copper_block", "copper", "mcl_copper_block_raw"},
		{"item", "copper_door", "copper", "mcl_copper_door"},
		{"block", "copper_door_bottom", "copper", "mcl_copper_door_bottom"},
		{"item", "exposed_copper_door", "copper", "mcl_copper_door_exposed"},
		{"block", "exposed_copper_door_bottom", "copper", "mcl_copper_door_exposed_bottom"},
		//{"block", "", "copper", "mcl_copper_door_exposed_lower"},
		//{"block", "", "copper", "mcl_copper_door_exposed_lower_bottompart"},
		//{"block", "", "copper", "mcl_copper_door_exposed_lower_side"},
		{"block", "exposed_copper_door_top", "copper", "mcl_copper_door_exposed_top"},
		//{"block", "", "copper", "mcl_copper_door_exposed_upper"},
		//{"block", "", "copper", "mcl_copper_door_exposed_upper_side"},
		//{"block", "", "copper", "mcl_copper_door_exposed_upper_toppart"},
		//{"block", "", "copper", "mcl_copper_door_lower"},
		//{"block", "", "copper", "mcl_copper_door_lower_bottompart"},
		//{"block", "", "copper", "mcl_copper_door_lower_side"},
		{"item", "oxidized_copper_door", "copper", "mcl_copper_door_oxidized"},
		{"block", "oxidized_copper_door_bottom", "copper", "mcl_copper_door_oxidized_bottom"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_lower"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_lower_bottompart"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_lower_side"},
		{"block", "oxidized_copper_door_top", "copper", "mcl_copper_door_oxidized_top"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_upper"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_upper_side"},
		//{"block", "", "copper", "mcl_copper_door_oxidized_upper_toppart"},
		{"block", "copper_door_top", "copper", "mcl_copper_door_top"},
		//{"block", "", "copper", "mcl_copper_door_upper"},
		//{"block", "", "copper", "mcl_copper_door_upper_side"},
		//{"block", "", "copper", "mcl_copper_door_upper_toppart"},
		{"item", "weathered_copper_door", "copper", "mcl_copper_door_weathered"},
		{"block", "weathered_copper_door_bottom", "copper", "mcl_copper_door_weathered_bottom"},
		//{"block", "", "copper", "mcl_copper_door_weathered_lower"},
		//{"block", "", "copper", "mcl_copper_door_weathered_lower_bottompart"},
		//{"block", "", "copper", "mcl_copper_door_weathered_lower_side"},
		{"block", "weathered_copper_door_top", "copper", "mcl_copper_door_weathered_top"},
		//{"block", "", "copper", "mcl_copper_door_weathered_upper"},
		//{"block", "", "copper", "mcl_copper_door_weathered_upper_side"},
		//{"block", "", "copper", "mcl_copper_door_weathered_upper_toppart"},
		{"block", "exposed_copper", "copper", "mcl_copper_exposed"},
		{"block", "exposed_copper_bulb", "copper", "mcl_copper_exposed_bulb_off"},
		{"block", "exposed_copper_bulb_lit", "copper", "mcl_copper_exposed_bulb_on"},
		{"block", "exposed_copper_bulb_powered", "copper", "mcl_copper_exposed_bulb_powered_off"},
		{"block", "exposed_copper_bulb_lit_powered", "copper", "mcl_copper_exposed_bulb_powered_on"},
		{"block", "exposed_chiseled_copper", "copper", "mcl_copper_exposed_chiseled"},
		{"block", "exposed_cut_copper", "copper", "mcl_copper_exposed_cut"},
		{"block", "exposed_copper_grate", "copper", "mcl_copper_exposed_grate"},
		{"item", "copper_ingot", "copper", "mcl_copper_ingot"},
		{"item", "copper_door", "copper", "mcl_copper_item_door"},
		{"item", "exposed_copper_door", "copper", "mcl_copper_item_door_exposed"},
		{"item", "oxidized_copper_door", "copper", "mcl_copper_item_door_oxidized"},
		{"item", "weathered_copper_door", "copper", "mcl_copper_item_door_weathered"},
		{"block", "copper_ore", "copper", "mcl_copper_ore"},
		{"block", "oxidized_copper", "copper", "mcl_copper_oxidized"},
		{"block", "oxidized_copper_bulb", "copper", "mcl_copper_oxidized_bulb_off"},
		{"block", "oxidized_copper_bulb_lit", "copper", "mcl_copper_oxidized_bulb_on"},
		{"block", "oxidized_copper_bulb_powered", "copper", "mcl_copper_oxidized_bulb_powered_off"},
		{"block", "oxidized_copper_bulb_lit_powered", "copper", "mcl_copper_oxidized_bulb_powered_on"},
		{"block", "oxidized_chiseled_copper", "copper", "mcl_copper_oxidized_chiseled"},
		{"block", "oxidized_cut_copper", "copper", "mcl_copper_oxidized_cut"},
		{"block", "oxidized_copper_grate", "copper", "mcl_copper_oxidized_grate"},
		{"item", "raw_copper", "copper", "mcl_copper_raw"},
		{"block", "copper_trapdoor", "copper", "mcl_copper_trapdoor"},
		{"block", "exposed_copper_trapdoor", "copper", "mcl_copper_trapdoor_exposed"},
		{"block", "exposed_copper_trapdoor", "copper", "mcl_copper_trapdoor_exposed_side"},
		{"block", "oxidized_copper_trapdoor", "copper", "mcl_copper_trapdoor_oxidized"},
		{"block", "oxidized_copper_trapdoor", "copper", "mcl_copper_trapdoor_oxidized_side"},
		{"block", "copper_trapdoor", "copper", "mcl_copper_trapdoor_side"},
		{"block", "weathered_copper_trapdoor", "copper", "mcl_copper_trapdoor_weathered"},
		{"block", "weathered_copper_trapdoor", "copper", "mcl_copper_trapdoor_weathered_side"},
		{"block", "weathered_copper", "copper", "mcl_copper_weathered"},
		{"block", "weathered_copper_bulb", "copper", "mcl_copper_weathered_bulb_off"},
		{"block", "weathered_copper_bulb_lit", "copper", "mcl_copper_weathered_bulb_on"},
		{"block", "weathered_copper_bulb_powered", "copper", "mcl_copper_weathered_bulb_powered_off"},
		{"block", "weathered_copper_bulb_lit_powered", "copper", "mcl_copper_weathered_bulb_powered_on"},
		{"block", "weathered_chiseled_copper", "copper", "mcl_copper_weathered_chiseled"},
		{"block", "weathered_cut_copper", "copper", "mcl_copper_weathered_cut"},
		{"block", "weathered_copper_grate", "copper", "mcl_copper_weathered_grate"},
		//// mcl_core
		{"block", "acacia_leaves", "core", "default_acacia_leaves"},
		{"block", "acacia_sapling", "core", "default_acacia_sapling"},
		{"block", "acacia_log", "core", "default_acacia_tree"},
		{"block", "acacia_log_top", "core", "default_acacia_tree_top"},
		{"block", "acacia_planks", "core", "default_acacia_wood"},
		{"item", "apple", "core", "default_apple"},
		{"block", "bricks", "core", "default_brick"},
		{"block", "clay", "core", "default_clay"},
		{"item", "brick", "core", "default_clay_brick"},
		{"item", "clay_ball", "core", "default_clay_lump"},
		{"block", "coal_block", "core", "default_coal_block"},
		{"item", "coal", "core", "default_coal_lump"},
		{"block", "cobblestone", "core", "default_cobble"},
		{"item", "diamond", "core", "default_diamond"},
		{"block", "diamond_block", "core", "default_diamond_block"},
		{"block", "dirt", "core", "default_dirt"},
		{"block", "dead_bush", "core", "default_dry_shrub"},
		{"item", "flint", "core", "default_flint"},
		{"block", "glass", "core", "default_glass"},
		//"default_glass_detail"	//no match?
		{"block", "gold_block", "core", "default_gold_block"},
		{"item", "gold_ingot", "core", "default_gold_ingot"},
		{"block", "gravel", "core", "default_gravel"},
		{"block", "ice", "core", "default_ice"},
		{"block", "jungle_leaves", "core", "default_jungleleaves"},
		{"block", "jungle_sapling", "core", "default_junglesapling"},
		{"block", "jungle_log", "core", "default_jungletree"},
		{"block", "jungle_log_top", "core", "default_jungletree_top"},
		{"block", "jungle_planks", "core", "default_junglewood"},
		{"block", "ladder", "core", "default_ladder"},
		{"block", "lava_flow", "core", "default_lava_flowing_animated"}, //special attention
		{"block", "lava_still", "core", "default_lava_source_animated"}, //special attention
		{"block", "oak_leaves", "core", "default_leaves"},
		{"block", "mossy_cobblestone", "core", "default_mossycobble"},
		{"block", "obsidian", "core", "default_obsidian"},
		{"item", "paper", "core", "default_paper"},
		{"block", "sand", "core", "default_sand"},
		{"block", "oak_sapling", "core", "default_sapling"},
		{"block", "snow", "core", "default_snow"},
		{"block", "iron_block", "core", "default_steel_block"},
		{"item", "iron_ingot", "core", "default_steel_ingot"},
		{"item", "stick", "core", "default_stick"},
		{"block", "stone", "core", "default_stone"},
		{"block", "stone_bricks", "core", "default_stone_brick"},
		{"block", "oak_log", "core", "default_tree"},
		{"block", "oak_log_top", "core", "default_tree_top"},
		//{"block", "water_flow", "core", "default_water_flowing_animated"}, //special attention
		//{"block", "water_still", "core", "default_water_source_animated"}, //special attention
		{"block", "oak_planks", "core", "default_wood"},
		{"block", "andesite", "core", "mcl_core_andesite"},
		{"block", "polished_andesite", "core", "mcl_core_andesite_smooth"},
		{"item", "golden_apple", "core", "mcl_core_apple_golden"},
		{"item", "barrier", "core", "mcl_core_barrier"},
		{"block", "bedrock", "core", "mcl_core_bedrock"},
		{"block", "bone_block_side", "core", "mcl_core_bone_block_side"},
		{"block", "bone_block_top", "core", "mcl_core_bone_block_top"},
		{"item", "bowl", "core", "mcl_core_bowl"},
		{"block", "cactus_bottom", "core", "mcl_core_cactus_bottom"},
		{"block", "cactus_side", "core", "mcl_core_cactus_side"},
		{"block", "cactus_top", "core", "mcl_core_cactus_top"},
		{"item", "charcoal", "core", "mcl_core_charcoal"},
		{"block", "coal_ore", "core", "mcl_core_coal_ore"},
		{"block", "coarse_dirt", "core", "mcl_core_coarse_dirt"},
		{"block", "crying_obsidian", "core", "mcl_core_crying_obsidian"}, //special attention? might be ok
		//"mcl_core_crying_obsidian_tear"		//no match?
		//"mcl_core_crying_obsidian_tear2"		//no match?
		//"mcl_core_crying_obsidian_tear3"		//no match?
		{"block", "diamond_ore", "core", "mcl_core_diamond_ore"},
		{"block", "diorite", "core", "mcl_core_diorite"},
		{"block", "polished_diorite", "core", "mcl_core_diorite_smooth"},
		{"block", "podzol_side", "core", "mcl_core_dirt_podzol_side"},
		{"block", "podzol_top", "core", "mcl_core_dirt_podzol_top"},
		{"item", "emerald", "core", "mcl_core_emerald"},
		{"block", "emerald_block", "core", "mcl_core_emerald_block"},
		{"block", "emerald_ore", "core", "mcl_core_emerald_ore"},
		{"block", "frosted_ice_0", "core", "mcl_core_frosted_ice_0"},
		{"block", "frosted_ice_1", "core", "mcl_core_frosted_ice_1"},
		{"block", "frosted_ice_2", "core", "mcl_core_frosted_ice_2"},
		{"block", "frosted_ice_3", "core", "mcl_core_frosted_ice_3"},
		//Glass TODO: All glass has normal and "detail" textures. Must create "detail".
		{"block", "black_stained_glass", "core", "mcl_core_glass_black"},
		{"block", "black_stained_glass", "core", "mcl_core_glass_black_detail"},
		{"block", "blue_stained_glass", "core", "mcl_core_glass_blue"},
		{"block", "blue_stained_glass", "core", "mcl_core_glass_blue_detail"},
		{"block", "brown_stained_glass", "core", "mcl_core_glass_brown"},
		{"block", "brown_stained_glass", "core", "mcl_core_glass_brown_detail"},
		{"block", "cyan_stained_glass", "core", "mcl_core_glass_cyan"},
		{"block", "cyan_stained_glass", "core", "mcl_core_glass_cyan_detail"},
		{"block", "gray_stained_glass", "core", "mcl_core_glass_gray"},
		{"block", "gray_stained_glass", "core", "mcl_core_glass_gray_detail"},
		{"block", "green_stained_glass", "core", "mcl_core_glass_green"},
		{"block", "green_stained_glass", "core", "mcl_core_glass_green_detail"},
		{"block", "light_blue_stained_glass", "core", "mcl_core_glass_light_blue"},
		{"block", "light_blue_stained_glass", "core", "mcl_core_glass_light_blue_detail"},
		{"block", "lime_stained_glass", "core", "mcl_core_glass_lime"},
		{"block", "lime_stained_glass", "core", "mcl_core_glass_lime_detail"},
		{"block", "magenta_stained_glass", "core", "mcl_core_glass_magenta"},
		{"block", "magenta_stained_glass", "core", "mcl_core_glass_magenta_detail"},
		{"block", "orange_stained_glass", "core", "mcl_core_glass_orange"},
		{"block", "orange_stained_glass", "core", "mcl_core_glass_orange_detail"},
		{"block", "pink_stained_glass", "core", "mcl_core_glass_pink"},
		{"block", "pink_stained_glass", "core", "mcl_core_glass_pink_detail"},
		{"block", "purple_stained_glass", "core", "mcl_core_glass_purple"},
		{"block", "purple_stained_glass", "core", "mcl_core_glass_purple_detail"},
		{"block", "red_stained_glass", "core", "mcl_core_glass_red"},
		{"block", "red_stained_glass", "core", "mcl_core_glass_red_detail"},
		{"block", "light_gray_stained_glass", "core", "mcl_core_glass_silver"},
		{"block", "light_gray_stained_glass", "core", "mcl_core_glass_silver_detail"},
		{"block", "white_stained_glass", "core", "mcl_core_glass_white"},
		{"block", "white_stained_glass", "core", "mcl_core_glass_white_detail"},
		{"block", "yellow_stained_glass", "core", "mcl_core_glass_yellow"},
		{"block", "yellow_stained_glass", "core", "mcl_core_glass_yellow_detail"},
		//Glass TODO
		{"item", "gold_nugget", "core", "mcl_core_gold_nugget"},
		{"block", "gold_ore", "core", "mcl_core_gold_ore"},
		{"block", "granite", "core", "mcl_core_granite"},
		{"block", "polished_granite", "core", "mcl_core_granite_smooth"},
		{"block", "grass_block_side_overlay", "core", "mcl_core_grass_block_side_overlay"},
		{"block", "grass_block_top", "core", "mcl_core_grass_block_top"},
		{"block", "dirt_path_side", "core", "mcl_core_grass_path_side"},
		{"block", "dirt_path_top", "core", "mcl_core_grass_path_top"},
		{"block", "grass_block_snow", "core", "mcl_core_grass_side_snowed"},
		{"block", "packed_ice", "core", "mcl_core_ice_packed"},
		{"item", "iron_nugget", "core", "mcl_core_iron_nugget"},
		{"block", "iron_ore", "core", "mcl_core_iron_ore"},
		{"item", "lapis_lazuli", "core", "mcl_core_lapis"},
		{"block", "lapis_block", "core", "mcl_core_lapis_block"},
		{"block", "lapis_ore", "core", "mcl_core_lapis_ore"},
		{"block", "dark_oak_leaves", "core", "mcl_core_leaves_big_oak"},
		{"block", "birch_leaves", "core", "mcl_core_leaves_birch"},
		{"block", "spruce_leaves", "core", "mcl_core_leaves_spruce"},
		{"item", "light_00", "core", "mcl_core_light_0"},
		{"item", "light_01", "core", "mcl_core_light_1"},
		{"item", "light_02", "core", "mcl_core_light_2"},
		{"item", "light_03", "core", "mcl_core_light_3"},
		{"item", "light_04", "core", "mcl_core_light_4"},
		{"item", "light_05", "core", "mcl_core_light_5"},
		{"item", "light_06", "core", "mcl_core_light_6"},
		{"item", "light_07", "core", "mcl_core_light_7"},
		{"item", "light_08", "core", "mcl_core_light_8"},
		{"item", "light_09", "core", "mcl_core_light_9"},
		{"item", "light_10", "core", "mcl_core_light_10"},
		{"item", "light_11", "core", "mcl_core_light_11"},
		{"item", "light_12", "core", "mcl_core_light_12"},
		{"item", "light_13", "core", "mcl_core_light_13"},
		{"item", "light_14", "core", "mcl_core_light_14"}, //no light_15 in Mineclonia
		{"block", "dark_oak_log", "core", "mcl_core_log_big_oak"},
		{"block", "dark_oak_log_top", "core", "mcl_core_log_big_oak_top"},
		{"block", "birch_log", "core", "mcl_core_log_birch"},
		{"block", "birch_log_top", "core", "mcl_core_log_birch_top"},
		{"block", "spruce_log", "core", "mcl_core_log_spruce"},
		{"block", "spruce_log_top", "core", "mcl_core_log_spruce_top"},
		//{"block", "", "core", "mcl_core_mycelium_particle"}, 	//special attention
		{"block", "mycelium_side", "core", "mcl_core_mycelium_side"},
		{"block", "mycelium_top", "core", "mcl_core_mycelium_top"},
		//{"block", "", "core", "mcl_core_palette_grass"}, 	//special attention
		//{"block", "", "core", "mcl_core_palette_leaves"}, 		//special attention
		{"block", "sugar_cane", "core", "mcl_core_papyrus"},
		{"block", "dark_oak_planks", "core", "mcl_core_planks_big_oak"},
		{"block", "birch_planks", "core", "mcl_core_planks_birch"},
		{"block", "spruce_planks", "core", "mcl_core_planks_spruce"},
		{"block", "red_sand", "core", "mcl_core_red_sand"},
		{"block", "red_sandstone_bottom", "core", "mcl_core_red_sandstone_bottom"},
		{"block", "chiseled_red_sandstone", "core", "mcl_core_red_sandstone_carved"},
		{"block", "red_sandstone", "core", "mcl_core_red_sandstone_normal"},
		{"block", "cut_red_sandstone", "core", "mcl_core_red_sandstone_smooth"},
		{"block", "red_sandstone_top", "core", "mcl_core_red_sandstone_top"},
		{"block", "redstone_ore", "core", "mcl_core_redstone_ore"},
		{"item", "sugar_cane", "core", "mcl_core_reeds"},
		{"block", "sandstone_bottom", "core", "mcl_core_sandstone_bottom"},
		{"block", "chiseled_sandstone", "core", "mcl_core_sandstone_carved"},
		{"block", "sandstone", "core", "mcl_core_sandstone_normal"},
		{"block", "cut_sandstone", "core", "mcl_core_sandstone_smooth"},
		{"block", "sandstone_top", "core", "mcl_core_sandstone_top"},
		{"block", "dark_oak_sapling", "core", "mcl_core_sapling_big_oak"},
		{"block", "birch_sapling", "core", "mcl_core_sapling_birch"},
		{"block", "spruce_sapling", "core", "mcl_core_sapling_spruce"},
		{"block", "slime_block", "core", "mcl_core_slime"},
		{"block", "chiseled_stone_bricks", "core", "mcl_core_stonebrick_carved"},
		{"block", "cracked_stone_bricks", "core", "mcl_core_stonebrick_cracked"},
		{"block", "mossy_stone_bricks", "core", "mcl_core_stonebrick_mossy"},
		{"block", "stripped_acacia_log", "core", "mcl_core_stripped_acacia_side"},
		{"block", "stripped_acacia_log_top", "core", "mcl_core_stripped_acacia_top"},
		{"block", "stripped_birch_log", "core", "mcl_core_stripped_birch_side"},
		{"block", "stripped_birch_log_top", "core", "mcl_core_stripped_birch_top"},
		{"block", "stripped_dark_oak_log", "core", "mcl_core_stripped_dark_oak_side"},
		{"block", "stripped_dark_oak_log_top", "core", "mcl_core_stripped_dark_oak_top"},
		{"block", "stripped_jungle_log", "core", "mcl_core_stripped_jungle_side"},
		{"block", "stripped_jungle_log_top", "core", "mcl_core_stripped_jungle_top"},
		{"block", "stripped_oak_log", "core", "mcl_core_stripped_oak_side"},
		{"block", "stripped_oak_log_top", "core", "mcl_core_stripped_oak_top"},
		{"block", "stripped_spruce_log", "core", "mcl_core_stripped_spruce_side"},
		{"block", "stripped_spruce_log_top", "core", "mcl_core_stripped_spruce_top"},
		{"item", "sugar", "core", "mcl_core_sugar"},
		//{"block", "vine", "core", "mcl_core_vine"}, //special attention
		//"mcl_core_void"     //no match
		{"block", "cobweb", "core", "mcl_core_web"},
		{"block", "grass_block_side_overlay", "core", "mcl_dirt_grass_shadow"}, //special attention?
		{"particle", "lava", "core", "mcl_particles_lava"},
		//{"block", "", "core", "mcl_stairs_andesite_smooth_slab"}, //special attention
		//{"block", "", "core", "mcl_stairs_diorite_smooth_slab"},  //special attention
		//{"block", "", "core", "mcl_stairs_granite_smooth_slab"},  //special attention
		//// mcl_crafting_table
		{"block", "crafting_table_front", "crafting_table", "crafting_workbench_front"},
		{"block", "crafting_table_side", "crafting_table", "crafting_workbench_side"},
		{"block", "crafting_table_top", "crafting_table", "crafting_workbench_top"},
		//"gui_crafting_arrow"
		//"mcl_crafting_table_inv_fill" //no match
		//// mcl_crimson
		{"block", "crimson_stem_top", "crimson", "crimson_hyphae"},
		//{"block", "crimson_stem", "crimson", "crimson_hyphae_side"},  //special attention
		{"block", "crimson_planks", "crimson", "crimson_hyphae_wood"},
		{"block", "crimson_nylium", "crimson", "crimson_nylium"},
		{"block", "crimson_nylium_side", "crimson", "crimson_nylium_side"},
		{"block", "crimson_roots", "crimson", "crimson_roots"},
		{"block", "stripped_crimson_stem", "crimson", "crimson_stem_stripped_side"},
		{"block", "stripped_crimson_stem_top", "crimson", "crimson_stem_stripped_top"},
		{"block", "crimson_fungus", "crimson", "farming_crimson_fungus"},
		{"block", "warped_fungus", "crimson", "farming_warped_fungus"},
		{"item", "crimson_door", "crimson", "mcl_crimson_crimson_door"},
		//{"block", "crimson_door_bottom", "crimson", "mcl_crimson_crimson_door_bottom"}, //flip
		//{"block", "crimson_door_top", "crimson", "mcl_crimson_crimson_door_top"},       //flip
		//{"block", "crimson_planks", "crimson", "mcl_crimson_crimson_fence"},
		//{"block", "", "crimson", "mcl_crimson_crimson_fence_side"},
		//{"block", "", "crimson", "mcl_crimson_crimson_fence_top"},
		{"block", "crimson_trapdoor", "crimson", "mcl_crimson_crimson_trapdoor"},
		{"block", "crimson_trapdoor", "crimson", "mcl_crimson_crimson_trapdoor_side"},
		{"item", "warped_door", "crimson", "mcl_crimson_warped_door"},
		//{"block", "warped_door_bottom", "crimson", "mcl_crimson_warped_door_bottom"}, //flip
		//{"block", "warped_door_top", "crimson", "mcl_crimson_warped_door_top"},       //flip
		{"block", "warped_planks", "crimson", "mcl_crimson_warped_fence"},
		//{"block", "", "crimson", "mcl_crimson_warped_fence_side"},
		//{"block", "", "crimson", "mcl_crimson_warped_fence_top"},
		{"block", "warped_trapdoor", "crimson", "mcl_crimson_warped_trapdoor"},
		{"block", "warped_trapdoor", "crimson", "mcl_crimson_warped_trapdoor_side"},
		{"block", "weeping_vines_plant", "crimson", "mcl_crimson_weeping_vines"},
		{"block", "crimson_door_bottom", "crimson", "mcl_doors_door_crimson_side_lower"},
		{"block", "crimson_door_top", "crimson", "mcl_doors_door_crimson_side_upper"},
		{"block", "warped_door_bottom", "crimson", "mcl_doors_door_warped_side_lower"},
		{"block", "warped_door_top", "crimson", "mcl_doors_door_warped_side_upper"},
		{"block", "nether_sprouts", "crimson", "nether_sprouts"},
		{"block", "nether_wart_block", "crimson", "nether_wart_block"},
		{"block", "shroomlight", "crimson", "shroomlight"},
		{"block", "stripped_crimson_stem", "crimson", "stripped_crimson_stem"},
		{"block", "stripped_crimson_stem", "crimson", "stripped_crimson_stem_side"},
		{"block", "stripped_crimson_stem_top", "crimson", "stripped_crimson_stem_top"},
		{"block", "stripped_warped_stem", "crimson", "stripped_warped_stem"},
		{"block", "stripped_warped_stem", "crimson", "stripped_warped_stem_side"},
		{"block", "stripped_warped_stem_top", "crimson", "stripped_warped_stem_top"},
		{"block", "twisting_vines", "crimson", "twisting_vines"},
		{"block", "twisting_vines_plant", "crimson", "twisting_vines_plant"},
		{"block", "warped_stem_top", "crimson", "warped_hyphae"},
		//{"block", "warped_stem", "crimson", "warped_hyphae_side"},  //special attention
		{"block", "warped_planks", "crimson", "warped_hyphae_wood"},
		{"block", "warped_nylium", "crimson", "warped_nylium"},
		{"block", "warped_nylium_side", "crimson", "warped_nylium_side"},
		{"block", "warped_roots", "crimson", "warped_roots"},
		//{"block", "stripped_warped_stem", "crimson", "warped_stem_stripped_side"},  //unused?
		{"block", "stripped_warped_stem_top", "crimson", "warped_stem_stripped_top"},
		{"block", "warped_wart_block", "crimson", "warped_wart_block"},
		//// mcl_deepslate
		{"block", "deepslate", "deepslate", "mcl_deepslate"},
		{"block", "deepslate_bricks", "deepslate", "mcl_deepslate_bricks"},
		{"block", "cracked_deepslate_bricks", "deepslate", "mcl_deepslate_bricks_cracked"},
		{"block", "chiseled_deepslate", "deepslate", "mcl_deepslate_chiseled"},
		{"block", "deepslate_coal_ore", "deepslate", "mcl_deepslate_coal_ore"},
		{"block", "cobbled_deepslate", "deepslate", "mcl_deepslate_cobbled"},
		{"block", "deepslate_copper_ore", "deepslate", "mcl_deepslate_copper_ore"},
		{"block", "deepslate_diamond_ore", "deepslate", "mcl_deepslate_diamond_ore"},
		{"block", "deepslate_emerald_ore", "deepslate", "mcl_deepslate_emerald_ore"},
		{"block", "deepslate_gold_ore", "deepslate", "mcl_deepslate_gold_ore"},
		{"block", "deepslate_iron_ore", "deepslate", "mcl_deepslate_iron_ore"},
		{"block", "deepslate_lapis_ore", "deepslate", "mcl_deepslate_lapis_ore"},
		{"block", "polished_deepslate", "deepslate", "mcl_deepslate_polished"},
		{"block", "deepslate_redstone_ore", "deepslate", "mcl_deepslate_redstone_ore"},
		{"block", "reinforced_deepslate_side", "deepslate", "mcl_deepslate_reinforced"},
		{"block", "reinforced_deepslate_bottom", "deepslate", "mcl_deepslate_reinforced_bottom"},
		{"block", "reinforced_deepslate_top", "deepslate", "mcl_deepslate_reinforced_top"},
		{"block", "deepslate_tiles", "deepslate", "mcl_deepslate_tiles"},
		{"block", "cracked_deepslate_tiles", "deepslate", "mcl_deepslate_tiles_cracked"},
		{"block", "deepslate_top", "deepslate", "mcl_deepslate_top"},
		{"block", "tuff", "deepslate", "mcl_deepslate_tuff"},
		{"block", "tuff_bricks", "deepslate", "mcl_deepslate_tuff_bricks"},
		{"block", "chiseled_tuff", "deepslate", "mcl_deepslate_tuff_chiseled"},
		{"block", "chiseled_tuff_bricks", "deepslate", "mcl_deepslate_tuff_chiseled_bricks"},
		{"block", "chiseled_tuff_bricks_top", "deepslate", "mcl_deepslate_tuff_chiseled_bricks_top"},
		{"block", "chiseled_tuff_top", "deepslate", "mcl_deepslate_tuff_chiseled_top"},
		{"block", "polished_tuff", "deepslate", "mcl_deepslate_tuff_polished"},
		//// mcl_doors
		{"item", "iron_door", "doors", "doors_item_steel"},
		{"item", "oak_door", "doors", "doors_item_wood"},
		{"block", "oak_trapdoor", "doors", "doors_trapdoor"},
		{"block", "oak_trapdoor", "doors", "doors_trapdoor_side"}, //
		{"block", "iron_trapdoor", "doors", "doors_trapdoor_steel"},
		{"block", "iron_trapdoor", "doors", "doors_trapdoor_steel_side"}, //
		{"item", "acacia_door", "doors", "mcl_doors_door_acacia"},
		//{"block", "acacia_door_bottom", "doors", "mcl_doors_door_acacia_lower"},
		{"block", "acacia_door_bottom", "doors", "mcl_doors_door_acacia_side_lower"}, //
		{"block", "acacia_door_top", "doors", "mcl_doors_door_acacia_side_upper"},    //
		//{"block", "acacia_door_top", "doors", "mcl_doors_door_acacia_upper"},
		{"item", "birch_door", "doors", "mcl_doors_door_birch"},
		//{"block", "birch_door_bottom", "doors", "mcl_doors_door_birch_lower"},
		{"block", "birch_door_bottom", "doors", "mcl_doors_door_birch_side_lower"}, //
		{"block", "birch_door_top", "doors", "mcl_doors_door_birch_side_upper"},    //
		//{"block", "birch_door_top", "doors", "mcl_doors_door_birch_upper"},
		{"item", "dark_oak_door", "doors", "mcl_doors_door_dark_oak"},
		//{"block", "dark_oak_door_bottom", "doors", "mcl_doors_door_dark_oak_lower"},
		{"block", "dark_oak_door_bottom", "doors", "mcl_doors_door_dark_oak_side_lower"}, //
		{"block", "dark_oak_door_top", "doors", "mcl_doors_door_dark_oak_side_upper"},    //
		//{"block", "dark_oak_door_top", "doors", "mcl_doors_door_dark_oak_upper"},
		{"block", "iron_door_bottom", "doors", "mcl_doors_door_iron_lower"},
		{"block", "iron_door_bottom", "doors", "mcl_doors_door_iron_side_lower"}, //
		{"block", "iron_door_top", "doors", "mcl_doors_door_iron_side_upper"},    //
		{"block", "iron_door_top", "doors", "mcl_doors_door_iron_upper"},
		{"item", "jungle_door", "doors", "mcl_doors_door_jungle"},
		//{"block", "jungle_door_bottom", "doors", "mcl_doors_door_jungle_lower"},
		{"block", "jungle_door_bottom", "doors", "mcl_doors_door_jungle_side_lower"}, //
		{"block", "jungle_door_top", "doors", "mcl_doors_door_jungle_side_upper"},    //
		//{"block", "jungle_door_top", "doors", "mcl_doors_door_jungle_upper"},
		{"item", "spruce_door", "doors", "mcl_doors_door_spruce"},
		//{"block", "spruce_door_bottom", "doors", "mcl_doors_door_spruce_lower"},
		{"block", "spruce_door_bottom", "doors", "mcl_doors_door_spruce_side_lower"}, //
		{"block", "spruce_door_top", "doors", "mcl_doors_door_spruce_side_upper"},    //
		//{"block", "spruce_door_top", "doors", "mcl_doors_door_spruce_upper"},
		//{"block", "oak_door_bottom", "doors", "mcl_doors_door_wood_lower"},
		{"block", "oak_door_bottom", "doors", "mcl_doors_door_wood_side_lower"},
		{"block", "oak_door_top", "doors", "mcl_doors_door_wood_side_upper"},
		//{"block", "oak_door_top", "doors", "mcl_doors_door_wood_upper"},
		{"block", "acacia_trapdoor", "doors", "mcl_doors_trapdoor_acacia"},
		{"block", "acacia_trapdoor", "doors", "mcl_doors_trapdoor_acacia_side"},
		{"block", "birch_trapdoor", "doors", "mcl_doors_trapdoor_birch"},
		{"block", "birch_trapdoor", "doors", "mcl_doors_trapdoor_birch_side"},
		{"block", "dark_oak_trapdoor", "doors", "mcl_doors_trapdoor_dark_oak"},
		{"block", "dark_oak_trapdoor", "doors", "mcl_doors_trapdoor_dark_oak_side"},
		{"block", "jungle_trapdoor", "doors", "mcl_doors_trapdoor_jungle"},
		{"block", "jungle_trapdoor", "doors", "mcl_doors_trapdoor_jungle_side"},
		{"block", "spruce_trapdoor", "doors", "mcl_doors_trapdoor_spruce"},
		{"block", "spruce_trapdoor", "doors", "mcl_doors_trapdoor_spruce_side"},
		//// mcl_dyes
		//// mcl_enchanting
		//// mcl_end
		//// mcl_farming
		//// mcl_fences
		//// mcl_fire
		//// mcl_fireworks
		//// mcl_fishing
		//// mcl_fletching_table
		//// mcl_flowerpots
		//// mcl_flowers
		//// mcl_furnaces
		{"block", "furnace_top", "furnaces", "default_furnace_bottom"}, //Minecraft does not have this texture.
		{"block", "furnace_front", "furnaces", "default_furnace_front"},
		{"block", "furnace_front_on", "furnaces", "default_furnace_front_active"},
		{"block", "furnace_side", "furnaces", "default_furnace_side"},
		{"block", "furnace_top", "furnaces", "default_furnace_top"},
		//// mcl_grindstone
		//// mcl_heads
		//// mcl_honey
		//// mcl_hoppers
		//// mcl_itemframes
		//// mcl_jukebox
		//// mcl_lanterns
		//// mcl_lectern
		//// mcl_lightning_rods
		//// mcl_loom
		//// mcl_lush_caves
		//// mcl_mangrove
		//// mcl_maps
		//// mcl_mobitems
		//// mcl_mobspawners
		//// mcl_monster_eggs
		//// mcl_mud
		{"block", "mud", "mud", "mcl_mud"},
		{"block", "mud_bricks", "mud", "mcl_mud_bricks"},
		{"block", "packed_mud", "mud", "mcl_mud_packed_mud"},
		//// mcl_mushrooms
		{"block", "brown_mushroom", "mushrooms", "farming_mushroom_brown"},
		{"block", "red_mushroom", "mushrooms", "farming_mushroom_red"},
		{"item", "mushroom_stew", "mushrooms", "farming_mushroom_stew"},
		{"block", "mushroom_block_inside", "mushrooms", "mcl_mushrooms_mushroom_block_inside"},
		{"block", "brown_mushroom_block", "mushrooms", "mcl_mushrooms_mushroom_block_skin_brown"},
		{"block", "red_mushroom_block", "mushrooms", "mcl_mushrooms_mushroom_block_skin_red"},
		{"block", "mushroom_stem", "mushrooms", "mcl_mushrooms_mushroom_block_skin_stem"},
		//// mcl_nether
		//TODO
		{"block", "ancient_debris_side", "nether", "mcl_nether_ancient_debris_side"},
		{"block", "ancient_debris_top", "nether", "mcl_nether_ancient_debris_top"},
		{"block", "chiseled_nether_bricks", "nether", "mcl_nether_chiseled_nether_bricks"},
		{"block", "cracked_nether_bricks", "nether", "mcl_nether_cracked_nether_bricks"},
		{"block", "glowstone", "nether", "mcl_nether_glowstone"},
		{"item", "glowstone_dust", "nether", "mcl_nether_glowstone_dust"},
		{"block", "nether_gold_ore", "nether", "mcl_nether_gold_ore"},
		{"block", "magma", "nether", "mcl_nether_magma"}, //is animated
		{"block", "nether_bricks", "nether", "mcl_nether_nether_brick"},
		{"item", "nether_brick", "nether", "mcl_nether_netherbrick"},
		{"block", "netherite_block", "nether", "mcl_nether_netheriteblock"},
		{"item", "netherite_ingot", "nether", "mcl_nether_netherite_ingot"},
		{"item", "netherite_scrap", "nether", "mcl_nether_netherite_scrap"},
		{"item", "netherite_upgrade_smithing_template", "nether", "mcl_nether_netherite_ugrade_template"},
		{"block", "netherrack", "nether", "mcl_nether_netherrack"},
		{"item", "nether_wart", "nether", "mcl_nether_nether_wart"},
		{"block", "nether_wart_block", "nether", "mcl_nether_nether_wart_block"},
		{"block", "nether_wart_stage0", "nether", "mcl_nether_nether_wart_stage_0"},
		{"block", "nether_wart_stage1", "nether", "mcl_nether_nether_wart_stage_1"},
		{"block", "nether_wart_stage2", "nether", "mcl_nether_nether_wart_stage_2"},
		{"item", "quartz", "nether", "mcl_nether_quartz"},
		{"block", "quartz_block_bottom", "nether", "mcl_nether_quartz_block_bottom"},
		{"block", "quartz_block_side", "nether", "mcl_nether_quartz_block_side"},
		{"block", "quartz_block_top", "nether", "mcl_nether_quartz_block_top"},
		{"block", "chiseled_quartz_block", "nether", "mcl_nether_quartz_chiseled_side"},
		{"block", "chiseled_quartz_block_top", "nether", "mcl_nether_quartz_chiseled_top"},
		{"block", "nether_quartz_ore", "nether", "mcl_nether_quartz_ore"},
		{"block", "quartz_pillar", "nether", "mcl_nether_quartz_pillar_side"},
		{"block", "quartz_pillar_top", "nether", "mcl_nether_quartz_pillar_top"},
		{"block", "red_nether_bricks", "nether", "mcl_nether_red_nether_brick"},
		{"block", "soul_sand", "nether", "mcl_nether_soul_sand"},
		//// mcl_ocean
		//// mcl_panes
		//// mcl_portals
		//// mcl_potions
		//// mcl_pottery_sherds
		//// mcl_raw_ores
		{"item", "raw_gold", "raw_ores", "mcl_raw_ores_raw_gold"},
		{"block", "raw_gold_block", "raw_ores", "mcl_raw_ores_raw_gold_block"},
		{"item", "raw_iron", "raw_ores", "mcl_raw_ores_raw_iron"},
		{"block", "raw_iron_block", "raw_ores", "mcl_raw_ores_raw_iron_block"},
		//// mcl_sculk
		//// mcl_shields
		//// mcl_signs
		//// mcl_smithing_table
		{"block", "smithing_table_bottom", "smithing_table", "mcl_smithing_table_bottom"},
		{"block", "smithing_table_front", "smithing_table", "mcl_smithing_table_front"},
		{"item", "empty_slot_smithing_template_armor_trim", "smithing_table", "mcl_smithing_table_inventory_trim_bg"},
		{"block", "smithing_table_side", "smithing_table", "mcl_smithing_table_side"},
		{"block", "smithing_table_top", "smithing_table", "mcl_smithing_table_top"},
		//// mcl_smoker
		{"block", "smoker_bottom", "smoker", "smoker_bottom"},
		{"block", "smoker_front", "smoker", "smoker_front"},
		{"block", "smoker_front_on", "smoker", "smoker_front_on"},
		{"block", "smoker_side", "smoker", "smoker_side"},
		{"block", "smoker_top", "smoker", "smoker_top"},
		//// mcl_sponges
		//// mcl_spyglass
		//// mcl_stairs
		//// mcl_stonecutter
		{"block", "stonecutter_bottom", "stonecutter", "mcl_stonecutter_bottom"},
		//{"block", "stonecutter_saw", "stonecutter", "mcl_stonecutter_saw"},  //special attention
		{"block", "stonecutter_side", "stonecutter", "mcl_stonecutter_side"},
		{"block", "stonecutter_top", "stonecutter", "mcl_stonecutter_top"},
		//// mcl_sus_nodes
		//// mcl_sus_stew
		{"item", "suspicious_stew", "sus_stew", "sus_stew"},
		//// mcl_throwing
		//// mcl_tnt
		{"block", "tnt_bottom", "tnt", "default_tnt_bottom"},
		{"block", "tnt_side", "tnt", "default_tnt_side"},
		{"block", "tnt_top", "tnt", "default_tnt_top"},
		//// mcl_tools
		{"item", "diamond_axe", "tools", "default_tool_diamondaxe"},
		{"item", "diamond_pickaxe", "tools", "default_tool_diamondpick"},
		{"item", "diamond_shovel", "tools", "default_tool_diamondshovel"},
		{"item", "diamond_sword", "tools", "default_tool_diamondsword"},
		{"item", "golden_axe", "tools", "default_tool_goldaxe"},
		{"item", "golden_pickaxe", "tools", "default_tool_goldpick"},
		{"item", "golden_shovel", "tools", "default_tool_goldshovel"},
		{"item", "golden_sword", "tools", "default_tool_goldsword"},
		{"item", "netherite_axe", "tools", "default_tool_netheriteaxe"},
		{"item", "netherite_pickaxe", "tools", "default_tool_netheritepick"},
		{"item", "netherite_shovel", "tools", "default_tool_netheriteshovel"},
		{"item", "netherite_sword", "tools", "default_tool_netheritesword"},
		{"item", "shears", "tools", "default_tool_shears"},
		{"item", "iron_axe", "tools", "default_tool_steelaxe"},
		{"item", "iron_pickaxe", "tools", "default_tool_steelpick"},
		{"item", "iron_shovel", "tools", "default_tool_steelshovel"},
		{"item", "iron_sword", "tools", "default_tool_steelsword"},
		{"item", "stone_axe", "tools", "default_tool_stoneaxe"},
		{"item", "stone_pickaxe", "tools", "default_tool_stonepick"},
		{"item", "stone_shovel", "tools", "default_tool_stoneshovel"},
		{"item", "stone_sword", "tools", "default_tool_stonesword"},
		{"item", "wooden_axe", "tools", "default_tool_woodaxe"},
		{"item", "wooden_pickaxe", "tools", "default_tool_woodpick"},
		{"item", "wooden_shovel", "tools", "default_tool_woodshovel"},
		{"item", "wooden_sword", "tools", "default_tool_woodsword"},
		//{"block", "", "tools", "mcl_tools_heavy_core_bottom"},  //special attention
		//{"block", "", "tools", "mcl_tools_heavy_core_side"},  //special attention
		//{"block", "", "tools", "mcl_tools_heavy_core_top"},  //special attention
		{"item", "mace", "tools", "mcl_tools_mace"},
		//// mcl_torches
		{"block", "torch", "torches", "default_torch_on_floor"},
		{"block", "torch", "torches", "default_torch_on_floor_animated"},
		//// mcl_totems
		//// mcl_trees
		//// mcl_walls
		//// mcl_wool
		{"block", "light_blue_wool", "wool", "mcl_wool_light_blue"},
		{"block", "lime_wool", "wool", "mcl_wool_lime"},
		{"block", "black_wool", "wool", "wool_black"},
		{"block", "blue_wool", "wool", "wool_blue"},
		{"block", "brown_wool", "wool", "wool_brown"},
		{"block", "cyan_wool", "wool", "wool_cyan"},
		{"block", "green_wool", "wool", "wool_dark_green"},
		{"block", "gray_wool", "wool", "wool_dark_grey"},
		{"block", "light_gray_wool", "wool", "wool_grey"},
		{"block", "magenta_wool", "wool", "wool_magenta"},
		{"block", "orange_wool", "wool", "wool_orange"},
		{"block", "pink_wool", "wool", "wool_pink"},
		{"block", "red_wool", "wool", "wool_red"},
		{"block", "purple_wool", "wool", "wool_violet"},
		{"block", "white_wool", "wool", "wool_white"},
		{"block", "yellow_wool", "wool", "wool_yellow"},
		//// mclx_core
		//// mclx_fences
		//// mclx_stairs
		//// REDSTONE
		//// screwdriver
		//{"block", "", "", ""},
	}
	return equivalents
}
