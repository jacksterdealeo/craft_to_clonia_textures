package main

import ()

var (
	unusedMTGPaths = map[string]string{
		"beds":               "/beds/",
		"binoculars":         "/binoculars/",
		"boats":              "/boats/",
		"bones":              "/bones/",
		"bucket":             "/bucket/",
		"butterflies":        "/butterflies/",
		"carts":              "/carts/",
		"creative":           "/creative/",
		"default":            "/default/",
		"doors":              "/doors/",
		"dungeon_loot":       "/dungeon_loot/",
		"dye":                "/dye/",
		"env_sounds":         "/env_sounds/",
		"farming":            "/farming/",
		"fire":               "/fire/",
		"fireflies":          "/fireflies/",
		"flowers":            "/flowers/",
		"game_commands":      "/game_commands/",
		"give_initial_stuff": "/give_initial_stuff/",
		"keys":               "/keys/",
		"map":                "/map/",
		"mtg_craftguide":     "/mtg_craftguide/",
		"player_api":         "/player_api/",
		"screwdriver":        "/screwdriver/",
		"sethome":            "/sethome/",
		"sfinv":              "/sfinv/",
		"spawn":              "/spawn/",
		"stairs":             "/stairs/",
		"tnt":                "/tnt/",
		"vessels":            "/vessels/",
		"walls":              "/walls/",
		"weather":            "/weather/",
		"wool":               "/wool/",
		"xpanes":             "/xpanes/",
	}

	mtgPaths = map[string]string{
		"mtg": "/mtg/",
	}

	minetestGameItems = [...]simpleConversion{
		// -- beds
		//{"bed", "", "mtg", "beds_bed.png", 1},  // no match
		//{"bed", "", "mtg", "beds_bed_fancy.png", 1},  // no match
		{"bed", "", "mtg", "beds_bed_foot.png", 1},
		{"bed", "", "mtg", "beds_bed_head.png", 1},
		{"bed", "", "mtg", "beds_bed_side_bottom.png", 1},
		{"bed", "", "mtg", "beds_bed_side_bottom_r.png", 1},
		{"bed", "", "mtg", "beds_bed_side_top.png", 1},
		{"bed", "", "mtg", "beds_bed_side_top_r.png", 1},
		{"bed", "", "mtg", "beds_bed_side1.png", 1},
		{"bed", "", "mtg", "beds_bed_side2.png", 1},
		{"bed", "", "mtg", "beds_bed_top_bottom.png", 1},
		{"bed", "", "mtg", "beds_bed_top_top.png", 1},
		{"bed", "", "mtg", "beds_bed_top1.png", 1},
		{"bed", "", "mtg", "beds_bed_top2.png", 1},
		{"bed", "", "mtg", "beds_bed_under.png", 1},
		// -- binoculars
		{"item", "spyglass.png", "mtg", "binoculars_binoculars.png", 1},
		// -- boats
		{"item", "oak_boat.png", "mtg", "boats_inventory.png", 1},
		{"item", "oak_boat.png", "mtg", "boats_wield.png", 1},
		// -- bones -- no match
		// -- bucket
		{"item", "bucket.png", "mtg", "bucket.png", 1},
		{"item", "lava_bucket.png", "mtg", "bucket_lava.png", 1},
		//{"item", "", "mtg", "bucket_river_water.png", 1},  // no match
		{"item", "water_bucket.png", "mtg", "bucket_water.png", 1},
		// -- butterflies -- no match
		// -- carts
		// Not converting due to different rails & different cart model
		// -- creative
		{"", "", "mtg", "", 1},

		// -- default
		{"hud", "bubble.png", "mtg", "bubble.png", 1},
		//{"", "", "mtg", "crack_anylength.png", 1},  // special case
		{"block", "acacia_sapling.png", "mtg", "default_acacia_bush_sapling.png", 1},
		//{"", "", "mtg", "default_acacia_bush_stem.png", 1},  // no match
		//{"", "", "mtg", "default_acacia_leaves.png", 1},  // greenify
		{"block", "acacia_sapling.png", "mtg", "default_acacia_sapling.png", 1},
		{"block", "acacia_log.png", "mtg", "default_acacia_tree.png", 1},
		{"block", "acacia_log_top.png", "mtg", "default_acacia_tree_top.png", 1},
		{"block", "acacia_planks.png", "mtg", "default_acacia_wood.png", 1},
		{"item", "apple.png", "mtg", "default_apple.png", 1},
		//{"block", "birch_leaves.png", "mtg", "default_aspen_leaves.png", 1},  // greenify
		{"block", "birch_sapling.png", "mtg", "default_aspen_sapling.png", 1},
		{"block", "birch_log.png", "mtg", "default_aspen_tree.png", 1},
		{"block", "birch_log_top.png", "mtg", "default_aspen_tree_top.png", 1},
		{"block", "birch_planks.png", "mtg", "default_aspen_wood.png", 1},
		//{"item", "sweet_berries.png", "mtg", "default_blueberries.png", 1},  // no match?
		//{"block", "", "mtg", "default_blueberry_bush_leaves.png", 1}, // no match
		//{"block", "sweet_berry_bush_stage1.png", "mtg", "default_blueberry_bush_sapling.png", 1}, // no match?
		//{"block", "", "mtg", "default_blueberry_overlay.png", 1},  // no match
		{"item", "book.png", "books", "default_book.png", 1},
		{"item", "written_book.png", "books", "default_book_written.png", 1},
		{"block", "bookshelf.png", "books", "default_bookshelf.png", 1},
		//{"item", ".png", "books", "mcl_book_book_empty_slot.png", 1}, //no match
		{"block", "bricks.png", "mtg", "default_brick.png", 1},
		//{"block", ".png", "mtg", "default_bronze_block.png", 1},  // no match
		//{"item", ".png", "mtg", "default_bronze_ingot.png", 1},  // no match
		//{"block", ".png", "mtg", "default_bush_sapling.png", 1},  // no match
		//{"block", ".png", "mtg", "default_bush_stem.png", 1},  // no match
		//{"block", ".png", "mtg", "default_cactus_side.png", 1},  // too small
		//{"block", ".png", "mtg", "default_cactus_top.png", 1},  // too small
		//{"block", ".png", "mtg", "default_chest_front.png", 1},  // no match
		//{"block", ".png", "mtg", "default_chest_inside.png", 1},  // no match
		//{"block", ".png", "mtg", "default_chest_lock.png", 1},  // no match
		//{"block", ".png", "mtg", "default_chest_side.png", 1},  // no match
		//{"block", ".png", "mtg", "default_chest_top.png", 1},  // no match
		{"block", "clay.png", "mtg", "default_clay.png", 1},
		{"item", "brick.png", "mtg", "default_clay_brick.png", 1},
		{"item", "clay_ball.png", "mtg", "default_clay_lump.png", 1},
		//{"block", "", "mtg", "default_cloud.png", 1},  // no match
		{"block", "coal_block.png", "mtg", "default_coal_block.png", 1},
		{"item", "coal.png", "mtg", "default_coal_lump.png", 1},
		{"block", "cobblestone.png", "mtg", "default_cobble.png", 1},
		//{"block", ".png", "mtg", "default_coniferous_litter.png", 1},  // no match
		//{"block", ".png", "mtg", "default_coniferous_litter_side.png", 1},  // no match
		{"block", "copper_block.png", "mtg", "default_copper_block.png", 1},
		{"item", "copper_ingot.png", "mtg", "default_copper_ingot.png", 1},
		{"item", "raw_copper.png", "mtg", "default_copper_lump.png", 1},
		//{"block", "fire_coral_block.png", "mtg", "default_coral_brown.png", 1},  // special attention
		//{"block", "tube_coral.png", "mtg", "default_coral_cyan.png", 1},  // special attention
		//{"block", "brain_coral.png", "mtg", "default_coral_green.png", 1},  // special attention
		//{"block", "horn_coral_block.png", "mtg", "default_coral_orange.png", 1},  // special attention
		//{"block", "fire_coral.png", "mtg", "default_coral_pink.png", 1},  // special attention
		{"block", "bone_block_side.png", "mtg", "default_coral_skeleton.png", 1},
		{"block", "red_sandstone_bottom.png", "mtg", "default_desert_cobble.png", 1},
		{"block", "sand.png", "mtg", "default_desert_sand.png", 1},
		{"block", "sandstone_top.png", "mtg", "default_desert_sandstone.png", 1},
		//{"block", "cut_sandstone.png", "mtg", "default_desert_sandstone_block.png", 1},  // special case
		//{"block", ".png", "mtg", "default_desert_sandstone_brick.png", 1},  // no match || special attention
		{"block", "red_sandstone_top.png", "mtg", "default_desert_stone.png", 1},
		//{"block", "cut_red_sandstone.png", "mtg", "default_desert_stone_block.png", 1},  // special case
		//{"block", ".png", "mtg", "default_desert_stone_brick.png", 1},  // no match || special attention
		{"item", "diamond.png", "mtg", "default_diamond.png", 1},
		{"block", "diamond_block.png", "mtg", "default_diamond_block.png", 1},
		{"block", "dirt.png", "mtg", "default_dirt.png", 1},
		{"block", "coarse_dirt.png", "mtg", "default_dry_dirt.png", 1},
		// dirt_path_top.png needs ofset by 1/16 upward
		{"block", "dirt_path_top.png", "mtg", "default_dry_grass.png", 1}, // special case
		//{"block", ".png", "mtg", "default_dry_grass_1.png", 1},  // no match
		//{"block", ".png", "mtg", "default_dry_grass_2.png", 1},  // no match
		//{"block", ".png", "mtg", "default_dry_grass_3.png", 1},  // no match
		//{"block", ".png", "mtg", "default_dry_grass_4.png", 1},  // no match
		//{"block", ".png", "mtg", "default_dry_grass_5.png", 1},  // no match
		{"block", "dirt_path_side.png", "mtg", "default_dry_grass_side.png", 1},
		{"block", "dead_bush.png", "mtg", "default_dry_shrub.png", 1},
		//{"block", "jungle_sapling.png", "mtg", "default_emergent_jungle_sapling.png", 1},  // no match

		/*
			After this point, entries are random.
		*/

		//{"block", ".png", "mtg", ".png", 1},
		//

		{"item", "flint.png", "mtg", "default_flint.png", 1},
		{"block", "glass.png", "mtg", "default_glass.png", 1},
		//"default_glass_detail.png"	//no match?
		{"block", "gold_block.png", "mtg", "default_gold_block.png", 1},
		{"item", "gold_ingot.png", "mtg", "default_gold_ingot.png", 1},
		{"item", "raw_gold.png", "mtg", "default_gold_lump.png", 1},
		//{"block", "grass_block_top.png", "mtg", "default_grass.png", 1},  // greenify
		//{"block", "grass_block_side_overlay.png", "mtg", "default_grass_side.png", 1},  // greenfiy

		{"block", "gravel.png", "mtg", "default_gravel.png", 1},
		{"block", "ice.png", "mtg", "default_ice.png", 1},
		//{"block", "jungle_leaves.png", "mtg", "default_jungleleaves.png", 1},  // greenify
		{"block", "jungle_sapling.png", "mtg", "default_junglesapling.png", 1},
		{"block", "jungle_log.png", "mtg", "default_jungletree.png", 1},
		{"block", "jungle_log_top.png", "mtg", "default_jungletree_top.png", 1},
		{"block", "jungle_planks.png", "mtg", "default_junglewood.png", 1},
		{"block", "ladder.png", "mtg", "default_ladder.png", 1},
		//{"block", "lava_flow.png", "mtg", "default_lava_flowing_animated.png", 1}, // special attention
		//{"block", "lava_still.png", "mtg", "default_lava_source_animated.png", 1}, // special attention
		//{"block", "oak_leaves.png", "mtg", "default_leaves.png", 1},  // greenify
		{"block", "mossy_cobblestone.png", "mtg", "default_mossycobble.png", 1},
		{"block", "obsidian.png", "mtg", "default_obsidian.png", 1},
		{"item", "paper.png", "mtg", "default_paper.png", 1},
		{"block", "sand.png", "mtg", "default_sand.png", 1},
		{"block", "oak_sapling.png", "mtg", "default_sapling.png", 1},
		{"block", "snow.png", "mtg", "default_snow.png", 1},
		{"block", "iron_block.png", "mtg", "default_steel_block.png", 1},
		{"item", "iron_ingot.png", "mtg", "default_steel_ingot.png", 1},
		{"item", "stick.png", "mtg", "default_stick.png", 1},
		{"block", "stone.png", "mtg", "default_stone.png", 1},
		{"block", "stone_bricks.png", "mtg", "default_stone_brick.png", 1},
		{"block", "oak_log.png", "mtg", "default_tree.png", 1},
		{"block", "oak_log_top.png", "mtg", "default_tree_top.png", 1},
		//{"block", "water_flow.png", "mtg", "default_water_flowing_animated.png", 1}, // special attention
		//{"block", "water_still.png", "mtg", "default_water_source_animated.png", 1}, // special attention
		{"block", "oak_planks.png", "mtg", "default_wood.png", 1},

		{"block", ".png", "mtg", ".png", 1},
		{"block", ".png", "mtg", ".png", 1},
	}
)
