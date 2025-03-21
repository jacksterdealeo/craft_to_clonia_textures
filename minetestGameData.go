package main

import ()

/* Just throwing everything from the game into one folder.
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
)
*/

var (
	mtgPaths = map[string]string{
		"mtg": "/mtg/",
	}

	minetestGreenery = [...]simpleConversion{ // note: the mtg_greenify function ignores animations at the moment.
		{"block", "acacia_leaves.png", "mtg", "default_acacia_leaves.png", 1},
		{"block", "acacia_leaves.png", "mtg", "default_acacia_leaves_simple.png", 1},
		{"block", "birch_leaves.png", "mtg", "default_aspen_leaves.png", 1},
		{"block", "grass_block_top.png", "mtg", "default_grass.png", 1},
		{"block", "grass_block_side_overlay.png", "mtg", "default_grass_side.png", 1},
		{"block", "jungle_leaves.png", "mtg", "default_jungleleaves.png", 1},
		{"block", "short_grass.png", "mtg", "default_junglegrass.png", 1},
		{"block", "jungle_leaves.png", "mtg", "default_jungleleaves_simple.png", 1},
		{"block", "oak_leaves.png", "mtg", "default_leaves.png", 1},
		{"block", "oak_leaves.png", "mtg", "default_leaves_simple.png", 1},
		{"block", "spruce_leaves.png", "mtg", "default_pine_needles.png", 1},
		//{"block", "short_grass.png", "mtg", "default_grass_5.png", 1},  // handled
	}

	minetestGameItems = [...]simpleConversion{
		// -- beds
		//{"bed", "", "mtg", "beds_bed.png", 1},  // no match
		//{"bed", "", "mtg", "beds_bed_fancy.png", 1},  // no match

		/* Beds require special attention fixes
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
		*/

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
		//{"", ".png", "mtg", "creative_clear_icon.png", 1},  // no match
		//{"", ".png", "mtg", "creative_next_icon.png", 1},  // no match
		//{"", ".png", "mtg", "creative_prev_icon.png", 1},  // no match
		{"icon", "search.png", "mtg", "creative_search_icon.png", 1},
		//{"", ".png", "mtg", "creative_trash_icon.png", 1},  // special attention
		// -- default
		{"particle", "bubble.png", "mtg", "bubble.png", 1},
		//{"", "", "mtg", "crack_anylength.png", 1},  // special attention
		{"block", "acacia_sapling.png", "mtg", "default_acacia_bush_sapling.png", 1},
		//{"", "", "mtg", "default_acacia_bush_stem.png", 1},  // no match
		//{"block", "acacia_leaves.png", "mtg", "default_acacia_leaves.png", 1},  // greenify
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
		{"item", "book.png", "mtg", "default_book.png", 1},
		{"item", "written_book.png", "mtg", "default_book_written.png", 1},
		{"block", "bookshelf.png", "mtg", "default_bookshelf.png", 1},
		//{"item", ".png", "mtg", "mcl_book_book_empty_slot.png", 1}, //no match
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
		//{"block", "dirt_path_top.png", "mtg", "default_dry_grass.png", 1}, // special case
		//{"block", ".png", "mtg", "default_dry_grass_1.png", 1},  // no match
		//{"block", ".png", "mtg", "default_dry_grass_2.png", 1},  // no match
		//{"block", ".png", "mtg", "default_dry_grass_3.png", 1},  // no match
		//{"block", ".png", "mtg", "default_dry_grass_4.png", 1},  // no match
		//{"block", ".png", "mtg", "default_dry_grass_5.png", 1},  // no match
		//{"block", "dirt_path_side.png", "mtg", "default_dry_grass_side.png", 1},  // special case
		{"block", "dead_bush.png", "mtg", "default_dry_shrub.png", 1},
		//{"block", "jungle_sapling.png", "mtg", "default_emergent_jungle_sapling.png", 1},  // no match
		{"block", "acacia_planks.png", "mtg", "default_fence_acacia_wood.png", 1},
		{"block", "birch_planks.png", "mtg", "default_fence_aspen_wood.png", 1},
		{"block", "jungle_planks.png", "mtg", "default_fence_junglewood.png", 1},
		//{"block", ".png", "mtg", "default_fence_overlay.png", 1},  // no match
		{"block", "spruce_planks.png", "mtg", "default_fence_pine_wood.png", 1},
		{"block", "acacia_planks.png", "mtg", "default_fence_rail_acacia_wood.png", 1},
		{"block", "birch_planks.png", "mtg", "default_fence_rail_aspen_wood.png", 1},
		{"block", "jungle_planks.png", "mtg", "default_fence_rail_junglewood.png", 1},
		//{"block", ".png", "mtg", "default_fence_rail_overlay.png", 1},  // no match
		{"block", "spruce_planks.png", "mtg", "default_fence_rail_pine_wood.png", 1},
		{"block", "oak_planks.png", "mtg", "default_fence_rail_wood.png", 1},
		{"block", "oak_planks.png", "mtg", "default_fence_wood.png", 1},
		//{"block", ".png", "mtg", "default_fern_1.png", 1},  // no match
		//{"block", ".png", "mtg", "default_fern_2.png", 1},  // no match
		//{"block", ".png", "mtg", "default_fern_3.png", 1},  // no match
		{"item", "flint.png", "mtg", "default_flint.png", 1},
		//{"block", ".png", "mtg", "default_footprint.png", 1},  // no match
		{"block", "blast_furnace_top.png", "mtg", "default_furnace_bottom.png", 1},
		//{"block", ".png", "mtg", "default_furnace_fire_bg.png", 1},  // special case // trim from furnace gui
		//{"gui", "sprites/container/furnace/lit_progress.png", "mtg", "default_furnace_fire_fg.png", 1},  // special case
		{"block", "furnace_front.png", "mtg", "default_furnace_front.png", 1},
		{"block", "furnace_front_on.png", "mtg", "default_furnace_front_active.png", -1},
		{"block", "furnace_side.png", "mtg", "default_furnace_side.png", 1},
		{"block", "furnace_top.png", "mtg", "default_furnace_top.png", 1},
		{"block", "glass.png", "mtg", "default_glass.png", 1},
		//{"block", "glass.png", "mtg", "default_glass_detail.png", 1},  // special case
		{"block", "gold_block.png", "mtg", "default_gold_block.png", 1},
		{"item", "gold_ingot.png", "mtg", "default_gold_ingot.png", 1},
		{"item", "raw_gold.png", "mtg", "default_gold_lump.png", 1},
		//{"block", "grass_block_top.png", "mtg", "default_grass.png", 1},  // greenify
		//{"block", "short_grass.png", "mtg", "default_grass_1.png", 1},  // greenify
		//{"block", "short_grass.png", "mtg", "default_grass_2.png", 1},  // greenify
		//{"block", "short_grass.png", "mtg", "default_grass_3.png", 1},  // greenify
		//{"block", "short_grass.png", "mtg", "default_grass_4.png", 1},  // greenify
		//{"block", "short_grass.png", "mtg", "default_grass_5.png", 1},  // greenify
		//{"block", "grass_block_side_overlay.png", "mtg", "default_grass_side.png", 1},  // greenfiy
		{"block", "gravel.png", "mtg", "default_gravel.png", 1},
		{"block", "ice.png", "mtg", "default_ice.png", 1},
		//{"block", ".png", "mtg", "default_invisible_node_overlay.png", 1},  // no match
		{"item", "raw_iron.png", "mtg", "default_iron_lump.png", 1},
		//{"block", ".png", "mtg", "default_item_smoke.png", 1},  // no match
		//{"block", "short_grass.png", "mtg", "default_junglegrass.png", 1},  // greenify
		//{"block", "jungle_leaves.png", "mtg", "default_jungleleaves.png", 1},  // greenify
		//{"block", "jungle_leaves.png", "mtg", "default_jungleleaves_simple.png", 1},  // greenify
		{"block", "jungle_sapling.png", "mtg", "default_junglesapling.png", 1},
		{"block", "jungle_log.png", "mtg", "default_jungletree.png", 1},
		{"block", "jungle_log_top.png", "mtg", "default_jungletree_top.png", 1},
		{"block", "jungle_planks.png", "mtg", "default_junglewood.png", 1},
		{"block", "kelp_plant.png", "mtg", "default_kelp.png", 1},
		//{"block", "ladder.png", "mtg", "default_ladder_steel.png", 1},  // special attention
		{"block", "ladder.png", "mtg", "default_ladder.png", 1},
		//{"block", ".png", "mtg", "default_large_cactus_seedling.png", 1},  //no match
		{"block", "lava_flow.png", "mtg", "default_lava.png", 1},
		//{"block", "lava_flow.png", "mtg", "default_lava_flowing_animated.png", -1}, // special attention
		//{"block", "lava_still.png", "mtg", "default_lava_source_animated.png", -1}, // special attention
		//{"block", "oak_leaves.png", "mtg", "default_leaves.png", 1},  // greenify
		//{"block", "oak_leaves.png", "mtg", "default_leaves_simple.png", 1},  // greenify
		//{"block", ".png", "mtg", "default_marram_grass_1.png", 1},  // no match
		//{"block", ".png", "mtg", "default_marram_grass_2.png", 1},  // no match
		//{"block", ".png", "mtg", "default_marram_grass_3.png", 1},  // no match
		//{"block", ".png", "mtg", "default_mese_block.png", 1},  // no match
		//{"block", ".png", "mtg", "default_mese_crystal.png", 1},  // no match
		//{"block", ".png", "mtg", "default_mese_crystal_fragment.png", 1},  // no match
		{"block", "redstone_lamp_on.png", "mtg", "default_meselamp.png", 1},
		{"block", "coal_ore.png", "mtg", "default_mineral_coal.png", 1},
		{"block", "copper_ore.png", "mtg", "default_mineral_copper.png", 1},
		{"block", "diamond_ore.png", "mtg", "default_mineral_diamond.png", 1},
		{"block", "gold_ore.png", "mtg", "default_mineral_gold.png", 1},
		{"block", "iron_ore.png", "mtg", "default_mineral_iron.png", 1},
		//{"block", ".png", "mtg", "default_mineral_mese.png", 1},
		//{"block", ".png", "mtg", "default_mineral_tin.png", 1},
		//{"block", ".png", "mtg", "default_moss.png", 1},  // no match
		//{"block", ".png", "mtg", "default_moss_side.png", 1},  // no match
		// {"block", "mossy_cobblestone.png", "mtg", "default_mossycobble.png", 1},  // in fixes
		{"block", "obsidian.png", "mtg", "default_obsidian.png", 1},
		//{"block", "obsidian.png", "mtg", "default_obsidian_block.png", 1},  // special case +smooth
		//{"block", "obsidian.png", "mtg", "default_obsidian_brick.png", 1},  // special case +brick
		//{"block", "tinted_glass.png", "mtg", "default_obsidian_glass.png", 1}, // in fixes
		//{"block", ".png", "mtg", "default_obsidian_glass_detail.png", 1},  // no match
		//{"block", ".png", "mtg", "default_obsidian_shard.png", 1},  // no match
		{"item", "paper.png", "mtg", "default_paper.png", 1},
		{"item", "sugar_cane.png", "mtg", "default_papyrus.png", 1},
		//{"block", ".png", "mtg", "default_permafrost.png", 1},  // no match
		//{"block", ".png", "mtg", "default_pine_bush_sapling.png", 1},  // no match
		//{"block", ".png", "mtg", "default_pine_bush_stem.png", 1},  // no match
		//{"block", "spruce_leaves.png", "mtg", "default_pine_needles.png", 1},  // greenify
		{"block", "spruce_sapling.png", "mtg", "default_pine_sapling.png", 1},
		{"block", "spruce_log.png", "mtg", "default_pine_tree.png", 1},
		{"block", "spruce_log_top.png", "mtg", "default_pine_tree_top.png", 1},
		{"block", "spruce_planks.png", "mtg", "default_pine_wood.png", 1},
		{"block", "podzol_top.png", "mtg", "default_rainforest_litter.png", 1},
		{"block", "podzol_side.png", "mtg", "default_rainforest_litter_side.png", 1},
		//{"block", ".png", "mtg", "default_river_water.png", 1},  // no match
		//{"block", ".png", "mtg", "default_river_water_flowing_animated.png", 1},  // no match
		//{"block", ".png", "mtg", "default_river_water_source_animated.png", 1},  // no match
		{"block", "sand.png", "mtg", "default_sand.png", 1},
		{"block", "sandstone_top.png", "mtg", "default_sandstone.png", 1},
		//{"block", "sandstone_top.png", "mtg", "default_sandstone_block.png", 1},  // special case +smooth
		//{"block", "sandstone_top.png", "mtg", "default_sandstone_brick.png", 1},  // special case +brick
		//{"block", "oak_sapling.png", "mtg", "default_sapling.png", 1},  // no match
		//{"block", ".png", "mtg", "default_sign_steel.png", 1},  // no match
		//{"block", ".png", "mtg", "default_sign_wall_steel.png", 1},  // no match
		//{"block", ".png", "mtg", "default_sign_wall_wood.png", 1},  // no match
		//{"block", ".png", "mtg", "default_sign_wood.png", 1},  // no match
		//{"block", ".png", "mtg", "default_silver_sand.png", 1},  // no match
		//{"block", ".png", "mtg", "default_silver_sandstone.png", 1},  // no match
		//{"block", ".png", "mtg", "default_silver_sandstone_block.png", 1},  // no match
		//{"block", ".png", "mtg", "default_silver_sandstone_brick.png", 1},  // no match
		{"block", "snow.png", "mtg", "default_snow.png", 1},
		{"block", "grass_block_snow.png", "mtg", "default_snow_side.png", 1},
		{"item", "snowball.png", "mtg", "default_snowball.png", 1},
		{"block", "iron_block.png", "mtg", "default_steel_block.png", 1},
		{"item", "iron_ingot.png", "mtg", "default_steel_ingot.png", 1},
		{"item", "stick.png", "mtg", "default_stick.png", 1},
		{"block", "stone.png", "mtg", "default_stone.png", 1},
		{"block", "stone_bricks.png", "mtg", "default_stone_brick.png", 1},
		//{"block", ".png", "mtg", "default_stones.png", 1},  // no match
		//{"block", ".png", "mtg", "default_stones_side.png", 1},  // no match
		//{"item", ".png", "mtg", "default_tin_block.png", 1},  // no match
		//{"item", ".png", "mtg", "default_tin_ingot.png", 1},  // no match
		//{"item", ".png", "mtg", "default_tin_lump.png", 1},  // no match
		//{"item", ".png", "mtg", "default_tool_bronzeaxe.png", 1},  // no match
		//{"item", ".png", "mtg", "default_tool_bronzepick.png", 1},  // no match
		//{"item", ".png", "mtg", "default_tool_bronzeshovel.png", 1},  // no match
		//{"item", ".png", "mtg", "default_tool_bronzesword.png", 1},  // no match
		{"item", "diamond_axe.png", "mtg", "default_tool_diamondaxe.png", 1},
		{"item", "diamond_pickaxe.png", "mtg", "default_tool_diamondpick.png", 1},
		{"item", "diamond_shovel.png", "mtg", "default_tool_diamondshovel.png", 1},
		{"item", "diamond_sword.png", "mtg", "default_tool_diamondsword.png", 1},
		{"item", "golden_axe.png", "mtg", "default_tool_meseaxe.png", 1},
		{"item", "golden_pickaxe.png", "mtg", "default_tool_mesepick.png", 1},
		{"item", "golden_shovel.png", "mtg", "default_tool_meseshovel.png", 1},
		{"item", "golden_sword.png", "mtg", "default_tool_mesesword.png", 1},
		{"item", "iron_axe.png", "mtg", "default_tool_steelaxe.png", 1},
		{"item", "iron_pickaxe.png", "mtg", "default_tool_steelpick.png", 1},
		{"item", "iron_shovel.png", "mtg", "default_tool_steelshovel.png", 1},
		{"item", "iron_sword.png", "mtg", "default_tool_steelsword.png", 1},
		{"item", "stone_axe.png", "mtg", "default_tool_stoneaxe.png", 1},
		{"item", "stone_pickaxe.png", "mtg", "default_tool_stonepick.png", 1},
		{"item", "stone_shovel.png", "mtg", "default_tool_stoneshovel.png", 1},
		{"item", "stone_sword.png", "mtg", "default_tool_stonesword.png", 1},
		{"item", "wooden_axe.png", "mtg", "default_tool_woodaxe.png", 1},
		{"item", "wooden_pickaxe.png", "mtg", "default_tool_woodpick.png", 1},
		{"item", "wooden_shovel.png", "mtg", "default_tool_woodshovel.png", 1},
		{"item", "wooden_sword.png", "mtg", "default_tool_woodsword.png", 1},
		//{"block", ".png", "mtg", "default_torch_animated.png", -1},  // special attention
		//{"block", ".png", "mtg", "default_torch_on_ceiling_animated.png", -1},  // special attention
		{"block", "torch.png", "mtg", "default_torch_on_floor.png", 1},
		{"block", "torch.png", "mtg", "default_torch_on_floor_animated.png", -1},
		{"block", "oak_log.png", "mtg", "default_tree.png", 1},
		{"block", "oak_log_top.png", "mtg", "default_tree_top.png", 1},
		//{"block", "water_still.png.png", "mtg", "default_water.png", 1},  // special attention
		//{"block", "water_flow.png", "mtg", "default_water_flowing_animated.png", 1}, // special attention
		//{"block", "water_still.png", "mtg", "default_water_source_animated.png", 1}, // special attention
		{"block", "oak_planks.png", "mtg", "default_wood.png", 1},
		/*  GUI
		{"gui", ".png", "mtg", "gui_formbg.png", 1},
		{"gui", ".png", "mtg", "gui_furnace_arrow_bg.png", 1},
		{"gui", ".png", "mtg", "gui_furnace_arrow_fg.png", 1},
		{"gui", ".png", "mtg", "gui_hb_bg.png", 1},
		{"gui", ".png", "mtg", "gui_hotbar.png", 1},
		{"gui", ".png", "mtg", "gui_hotbar_selected.png", 1},
		{"gui", ".png", "mtg", "heart.png", 1},
		{"gui", ".png", "mtg", "wieldhand.png", 1},
		*/
		//{"block", ".png", "mtg", ".png", 1},  // no match
		// -- doors
		//{"block", ".png", "mtg", "doors_door_glass.png", 1},  // no match
		//{"block", ".png", "mtg", "doors_door_obsidian_glass.png", 1},  // no match
		//{"block", ".png", "mtg", "doors_door_steel.png", 1},  // special attention
		//{"block", ".png", "mtg", "doors_door_wood.png", 1},  // special attention
		//{"block", ".png", "mtg", "doors_hidden_segment.png", 1},  // special attention
		//{"block", ".png", "mtg", "doors_item_glass.png", 1},  // no match
		//{"block", ".png", "mtg", "doors_item_obsidian_glass.png", 1},  // no match
		{"item", "iron_door.png", "mtg", "doors_item_steel.png", 1},
		{"item", "oak_door.png", "mtg", "doors_item_wood.png", 1},
		//{"block", ".png", "mtg", "doors_trapdoor.png", 1},  // special attention
		//{"block", ".png", "mtg", "doors_trapdoor_side.png", 1},  // special attention
		//{"block", ".png", "mtg", "doors_trapdoor_steel.png", 1},  // special attention
		//{"block", ".png", "mtg", "doors_trapdoor_steel_side.png", 1},  // special attention
		// -- dye
		{"item", "black_dye.png", "mtg", "dye_black.png", 1},
		{"item", "blue_dye.png", "mtg", "dye_blue.png", 1},
		{"item", "brown_dye.png", "mtg", "dye_brown.png", 1},
		{"item", "cyan_dye.png", "mtg", "dye_cyan.png", 1},
		{"item", "green_dye.png", "mtg", "dye_dark_green.png", 1},
		{"item", "gray_dye.png", "mtg", "dye_dark_grey.png", 1},
		{"item", "lime_dye.png", "mtg", "dye_green.png", 1},
		{"item", "light_gray_dye.png", "mtg", "dye_grey.png", 1},
		{"item", "magenta_dye.png", "mtg", "dye_magenta.png", 1},
		{"item", "orange_dye.png", "mtg", "dye_orange.png", 1},
		{"item", "pink_dye.png", "mtg", "dye_pink.png", 1},
		{"item", "red_dye.png", "mtg", "dye_red.png", 1},
		{"item", "purple_dye.png", "mtg", "dye_violet.png", 1},
		{"item", "white_dye.png", "mtg", "dye_white.png", 1},
		{"item", "yellow_dye.png", "mtg", "dye_yellow.png", 1},
		// -- farming
		{"item", "bread.png", "mtg", "farming_bread.png", 1},
		//{"item", ".png", "mtg", "farming_cotton.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_1.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_2.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_3.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_4.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_5.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_6.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_7.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_8.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_seed.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_cotton_wild.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_desert_sand_soil.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_desert_sand_soil_wet.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_desert_sand_soil_wet_side.png", 1},  // no match
		//{"item", ".png", "mtg", "farming_flour.png", 1},  // no match
		{"block", "farmland.png", "mtg", "farming_soil.png", 1},
		{"block", "farmland_moist.png", "mtg", "farming_soil_wet.png", 1}, // special attention
		//{"block", ".png", "mtg", "farming_soil_wet_side.png", 1},  // no match
		{"block", "hay_block_top.png", "mtg", "farming_straw.png", 1},
		{"item", "string.png", "mtg", "farming_string.png", 1},
		//{"item", ".png", "mtg", "farming_tool_bronzehoe.png", 1},  // no match
		{"item", "diamond_hoe.png", "mtg", "farming_tool_diamondhoe.png", 1},
		{"item", "golden_hoe.png", "mtg", "farming_tool_mesehoe.png", 1},
		{"item", "iron_hoe.png", "mtg", "farming_tool_steelhoe.png", 1},
		{"item", "stone_hoe.png", "mtg", "farming_tool_stonehoe.png", 1},
		{"item", "wooden_hoe.png", "mtg", "farming_tool_woodhoe.png", 1},
		{"item", "wheat.png", "mtg", "farming_wheat.png", 1},
		{"block", "wheat_stage0.png", "mtg", "farming_wheat_1.png", 1},
		{"block", "wheat_stage1.png", "mtg", "farming_wheat_2.png", 1},
		{"block", "wheat_stage2.png", "mtg", "farming_wheat_3.png", 1},
		{"block", "wheat_stage3.png", "mtg", "farming_wheat_4.png", 1},
		{"block", "wheat_stage4.png", "mtg", "farming_wheat_5.png", 1},
		{"block", "wheat_stage5.png", "mtg", "farming_wheat_6.png", 1},
		{"block", "wheat_stage6.png", "mtg", "farming_wheat_7.png", 1},
		{"block", "wheat_stage7.png", "mtg", "farming_wheat_8.png", 1},
		{"item", "wheat_seeds.png", "mtg", "farming_wheat_seed.png", 1},
		// -- fire
		{"block", "fire_0.png", "mtg", "fire_basic_flame.png", 1},
		{"block", "fire_0.png", "mtg", "fire_basic_flame_animated.png", 8},
		{"item", "flint_and_steel.png", "mtg", "fire_flint_steel.png", 1},
		// -- fireflies
		//{"item", ".png", "mtg", ".png", 1},
		//{"item", ".png", "mtg", ".png", 1},
		//{"item", ".png", "mtg", ".png", 1},
		//{"item", ".png", "mtg", ".png", 1},
		//{"item", ".png", "mtg", ".png", 1},
		// -- flowers
		//{"block", ".png", "mtg", "flowers_chrysanthemum_green.png", 1},  // special case
		// Maybe "greenify" another flower?
		{"block", "oxeye_daisy.png", "mtg", "flowers_dandelion_white.png", 1},
		{"block", "dandelion.png", "mtg", "flowers_dandelion_yellow.png", 1},
		{"block", "cornflower.png", "mtg", "flowers_geranium.png", 1},
		{"block", "brown_mushroom.png", "mtg", "flowers_mushroom_brown.png", 1},
		{"block", "red_mushroom.png", "mtg", "flowers_mushroom_red.png", 1},
		{"block", "red_tulip.png", "mtg", "flowers_rose.png", 1},
		{"block", "orange_tulip.png", "mtg", "flowers_tulip.png", 1},
		{"block", "wither_rose.png", "mtg", "flowers_tulip_black.png", 1},
		{"block", "allium.png", "mtg", "flowers_viola.png", 1},
		//{"block", ".png", "mtg", "flowers_waterlily.png", 1},  // special case
		//{"block", ".png", "mtg", "flowers_waterlily_bottom.png", 1},  // special case
		// -- keys
		//{"item", ".png", "mtg", ".png", 1},  // no match
		//{"item", ".png", "mtg", ".png", 1},  // no match
		// -- map
		{"item", "filled_map.png", "mtg", "map_mapping_kit.png", 1},
		// -- mtg_craftguide
		{"item", ".png", "mtg", "craftguide_clear_icon.png", 1},
		{"item", ".png", "mtg", "craftguide_furnace.png", 1},
		{"item", ".png", "mtg", "craftguide_next_icon.png", 1},
		{"item", ".png", "mtg", "craftguide_prev_icon.png", 1},
		{"item", ".png", "mtg", "craftguide_search_icon.png", 1},
		{"item", ".png", "mtg", "craftguide_shapeless.png", 1},
		// -- player_api
		//{"", ".png", "mtg", "player.png", 1},  // no match
		//{"", ".png", "mtg", "player_back.png", 1},  // no match
		// -- screwdriver
		//{"item", ".png", "mtg", "screwdriver.png", 1},  // no match
		// -- sfinv
		//{"", ".png", "mtg", "sfinv_crafting_arrow.png", 1},  // no match
		// -- stairs
		/* no matches & special cases
		{"item", ".png", "mtg", "stairs_glass_outer_stairside.png", 1},
		{"item", ".png", "mtg", "stairs_glass_split.png", 1},
		{"item", ".png", "mtg", "stairs_glass_stairside.png", 1},
		{"item", ".png", "mtg", "stairs_obsidian_glass_outer_stairside.png", 1},  // no match
		{"item", ".png", "mtg", "stairs_obsidian_glass_split.png", 1},  // no match
		{"item", ".png", "mtg", "stairs_obsidian_glass_stairside.png", 1},  // no match
		*/
		// -- tnt
		//{"item", ".png", "mtg", "tnt_blast.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_boom.png", 1},  // no match
		{"block", "tnt_bottom.png", "mtg", "tnt_bottom.png", 1},
		//{"item", ".png", "mtg", "tnt_gunpowder_burning_crossing_animated.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_gunpowder_burning_curved_animated.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_gunpowder_burning_straight_animated.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_gunpowder_burning_t_junction_animated.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_gunpowder_crossing.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_gunpowder_curved.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_gunpowder_inventory.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_gunpowder_straight.png", 1},  // no match
		//{"item", ".png", "mtg", "tnt_gunpowder_t_junction.png", 1},  // no match
		{"block", "tnt_side.png", "mtg", "tnt_side.png", 1},
		{"particle", "big_smoke_5.png", "mtg", "tnt_smoke.png", 1},
		//{"item", ".png", "mtg", "tnt_tnt_stick.png", 1},  // no match
		{"block", "tnt_top.png", "mtg", "tnt_top.png", 1},
		{"block", "tnt_top.png", "mtg", "tnt_top_burning.png", 1},          // special case
		{"block", "tnt_top.png", "mtg", "tnt_top_burning_animated.png", 1}, // special case
		// -- vessels
		//{"item", ".png", "mtg", "vessels_drinking_glass.png", 1},  // no match
		//{"item", ".png", "mtg", "vessels_drinking_glass_inv.png", 1},  // no match
		{"item", "glass_bottle.png", "mtg", "vessels_glass_bottle.png", 1},
		//{"item", ".png", "mtg", "vessels_glass_fragments.png", 1},  // no match
		//{"block", ".png", "mtg", "vessels_shelf.png", 1},  // no match
		//{"item", ".png", "mtg", "vessels_shelf_slot.png", 1},  // no match
		//{"item", ".png", "mtg", "vessels_steel_bottle.png", 1},  // no match
		// -- wool
		{"block", "black_wool.png", "mtg", "wool_black.png", 1},
		{"block", "blue_wool.png", "mtg", "wool_blue.png", 1},
		{"block", "brown_wool.png", "mtg", "wool_brown.png", 1},
		{"block", "cyan_wool.png", "mtg", "wool_cyan.png", 1},
		{"block", "green_wool.png", "mtg", "wool_dark_green.png", 1},
		{"block", "gray_wool.png", "mtg", "wool_dark_grey.png", 1},
		{"block", "lime_wool.png", "mtg", "wool_green.png", 1},
		{"block", "light_gray_wool.png", "mtg", "wool_grey.png", 1},
		{"block", "magenta_wool.png", "mtg", "wool_magenta.png", 1},
		{"block", "orange_wool.png", "mtg", "wool_orange.png", 1},
		{"block", "pink_wool.png", "mtg", "wool_pink.png", 1},
		{"block", "red_wool.png", "mtg", "wool_red.png", 1},
		{"block", "purple_wool.png", "mtg", "wool_violet.png", 1},
		{"block", "white_wool.png", "mtg", "wool_white.png", 1},
		{"block", "yellow_wool.png", "mtg", "wool_yellow.png", 1},
		// -- xpanes
		{"block", "iron_bars.png", "mtg", "xpanes_bar.png", 1},
		//{"block", ".png", "mtg", "xpanes_bar_top.png", 1},  // no match
		//{"block", ".png", "mtg", "xpanes_door_steel_bar.png", 1},  // no match
		//{"block", ".png", "mtg", "xpanes_edge.png", 1},  no match
		//{"block", ".png", "mtg", "xpanes_edge_obsidian.png", 1},  // no match
		//{"block", ".png", "mtg", "xpanes_item_steel_bar.png", 1},  // no match
		//{"block", ".png", "mtg", "xpanes_trapdoor_steel_bar.png", 1},  // no match
		//{"block", ".png", "mtg", "xpanes_trapdoor_steel_bar_side.png", 1},  // no match
	}
)
