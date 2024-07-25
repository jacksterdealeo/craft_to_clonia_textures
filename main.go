package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Minecraft to Mineclonia Texture Pack Converter v0.0.0\n")
	inName := "Faithful 32x - 1.21"
	outName := fmt.Sprintf("%s_mc_converted", strings.ReplaceAll(strings.ToLower(inName), " ", "_"))

	if fs.ValidPath(outName) {
		if err := os.Mkdir(outName, 0755); err != nil {
			if errors.Is(err, fs.ErrInvalid) {
				fmt.Printf("Folder %s is an \"invalid argument\". Maybe rename %s?\n", outName, inName)
			} else if errors.Is(err, fs.ErrPermission) {
				fmt.Printf("Permission was denied. %s was not made.\n", outName)
			} else if errors.Is(err, fs.ErrExist) {
				fmt.Printf("Folder %s already exists. Writing into it.\n", outName)
			} else {
				fmt.Printf("How.\n")
				log.Panic(err)
			}
		}
	}

	packConfigFile := fmt.Sprintf(`title = %s
name = %s
description = A converted Minecraft texture pack
author = Unknown
release = 01`, inName, outName)

	fmt.Printf("Pack info:\n%s\n", packConfigFile)
	err := os.WriteFile(outName+"/texture_pack.conf", []byte(packConfigFile), 0644)
	if err != nil {
		log.Panic(err)
	}

	// pack icon
	copyTexture(inName+"/pack.png", outName+"/screenshot.png")

	f := map[string]string{
		"amethyst":       "/ITEMS/mcl_amethyst/textures/",
		"anvils":         "/ITEMS/mcl_anvils/textures/",
		"blast_furnace":  "/ITEMS/mcl_blast_furnace/textures/",
		"core":           "/ITEMS/mcl_core/textures/",
		"furnaces":       "/ITEMS/mcl_furnaces/textures/",
		"mud":            "/ITEMS/mcl_mud/textures/",
		"mushrooms":      "/ITEMS/mcl_mushrooms/textures/",
		"raw_ores":       "/ITEMS/mcl_raw_ores/textures/",
		"smithing_table": "/ITEMS/mcl_smithing_table/textures/",
		"smoker":         "/ITEMS/mcl_smoker/textures/",
		"sus_stew":       "/ITEMS/mcl_sus_stew/textures/",
	}
	for _, v := range f {
		if err := os.MkdirAll(outName+v, 0755); err != nil {
			log.Panic(err)
		}
	}

	srcLocation := inName + "/assets/minecraft/textures/"
	blocksAndItems := [...][2]string{
		// mcl_amethyst
		{"block/amethyst_block.png", f["amethyst"] + "mcl_amethyst_amethyst_block.png"},
		{"block/large_amethyst_bud.png", f["amethyst"] + "mcl_amethyst_amethyst_bud_large.png"},
		{"block/medium_amethyst_bud.png", f["amethyst"] + "mcl_amethyst_amethyst_bud_medium.png"},
		{"block/small_amethyst_bud.png", f["amethyst"] + "mcl_amethyst_amethyst_bud_small.png"},
		{"block/amethyst_cluster.png", f["amethyst"] + "mcl_amethyst_amethyst_cluster.png"},
		{"item/amethyst_shard.png", f["amethyst"] + "mcl_amethyst_amethyst_shard.png"},
		{"block/budding_amethyst.png", f["amethyst"] + "mcl_amethyst_budding_amethyst.png"},
		{"block/calcite.png", f["amethyst"] + "mcl_amethyst_calcite_block.png"},
		{"block/tinted_glass.png", f["amethyst"] + "mcl_amethyst_tinted_glass.png"},
		// mcl_anvils
		// FIX: anvil damaged texture is not wide enough for all the model.
		{"block/anvil.png", f["anvils"] + "mcl_anvils_anvil_base.png"},
		{"block/anvil.png", f["anvils"] + "mcl_anvils_anvil_side.png"},
		{"block/anvil_top.png", f["anvils"] + "mcl_anvils_anvil_top_damaged_0.png"},
		{"block/chipped_anvil_top.png", f["anvils"] + "mcl_anvils_anvil_top_damaged_1.png"},
		{"block/damaged_anvil_top.png", f["anvils"] + "mcl_anvils_anvil_top_damaged_2.png"},
		// mcl_armor
		// mcl_armor_stand
		// mcl_bamboo
		// mcl_banners
		// mcl_barrels
		// mcl_beacons
		// mcl_beds
		// mcl_beehives
		// mcl_bells
		// mcl_blackstone
		// mcl_blast_furnace
		{"block/blast_furnace_front.png", f["blast_furnace"] + "blast_furnace_front.png"},
		{"block/blast_furnace_front_on.png", f["blast_furnace"] + "blast_furnace_front_on.png"},
		{"block/blast_furnace_side.png", f["blast_furnace"] + "blast_furnace_side.png"},
		{"block/blast_furnace_top.png", f["blast_furnace"] + "blast_furnace_top.png"},
		// mcl_bone_meal
		// mcl_books
		// mcl_bows
		// mcl_brewing
		// mcl_buckets
		// mcl_cake
		// mcl_campfires
		// mcl_cartography_table
		// mcl_cauldrons
		// mcl_cherry_blossom
		// mcl_chests
		// mcl_clock
		// mcl_cocoas
		// mcl_colorblocks
		// mcl_compass
		// mcl_composters
		// mcl_conduits
		// mcl_copper
		// mcl_core
		{"block/acacia_leaves.png", f["core"] + "default_acacia_leaves.png"},
		{"block/acacia_sapling.png", f["core"] + "default_acacia_sapling.png"},
		{"block/acacia_log.png", f["core"] + "default_acacia_tree.png"},
		{"block/acacia_log_top.png", f["core"] + "default_acacia_tree_top.png"},
		{"block/acacia_planks.png", f["core"] + "default_acacia_wood.png"},
		{"item/apple.png", f["core"] + "default_apple.png"},
		{"block/bricks.png", f["core"] + "default_brick.png"},
		{"block/clay.png", f["core"] + "default_clay.png"},
		{"item/brick.png", f["core"] + "default_clay_brick.png"},
		{"item/clay_ball.png", f["core"] + "default_clay_lump.png"},
		{"block/coal_block.png", f["core"] + "default_coal_block.png"},
		{"item/coal.png", f["core"] + "default_coal_lump.png"},
		{"block/cobblestone.png", f["core"] + "default_cobble.png"},
		{"item/diamond.png", f["core"] + "default_diamond.png"},
		{"block/diamond_block.png", f["core"] + "default_diamond_block.png"},
		{"block/dirt.png", f["core"] + "default_dirt.png"},
		{"block/dead_bush.png", f["core"] + "default_dry_shrub.png"},
		{"item/flint.png", f["core"] + "default_flint.png"},
		{"block/glass.png", f["core"] + "default_glass.png"},
		//{"block/", f["core"] + "default_glass_detail.png"},		//no match?
		{"block/gold_block.png", f["core"] + "default_gold_block.png"},
		{"ingot/gold_ingot.png", f["core"] + "default_gold_ingot.png"},
		{"block/gravel.png", f["core"] + "default_gravel.png"},
		{"block/ice.png", f["core"] + "default_ice.png"},
		{"block/jungle_leaves.png", f["core"] + "default_jungleleaves.png"},
		{"block/jungle_sapling.png", f["core"] + "default_junglesapling.png"},
		{"block/jungle_log.png", f["core"] + "default_jungletree.png"},
		{"block/jungle_log_top.png", f["core"] + "default_jungletree_top.png"},
		{"block/jungle_planks.png", f["core"] + "default_junglewood.png"},
		{"block/ladder.png", f["core"] + "default_ladder.png"},
		{"block/lava_flow.png", f["core"] + "default_lava_flowing_animated.png"}, //special attention
		{"block/lava_still.png", f["core"] + "default_lava_source_animated.png"}, //special attention
		{"block/oak_leaves.png", f["core"] + "default_leaves.png"},
		{"block/mossy_cobblestone.png", f["core"] + "default_mossycobble.png"},
		{"block/obsidian.png", f["core"] + "default_obsidian.png"},
		{"item/paper.png", f["core"] + "default_paper.png"},
		{"block/sand.png", f["core"] + "default_sand.png"},
		{"block/oak_sapling.png", f["core"] + "default_sapling.png"},
		{"block/snow.png", f["core"] + "default_snow.png"},
		{"block/iron_block.png", f["core"] + "default_steel_block.png"},
		{"item/iron_ingot.png", f["core"] + "default_steel_ingot.png"},
		{"item/stick.png", f["core"] + "default_stick.png"},
		{"block/stone_bricks.png", f["core"] + "default_stone_brick.png"},
		{"block/oak_log.png", f["core"] + "default_tree.png"},
		{"block/oak_log_top.png", f["core"] + "default_tree_top.png"},
		{"block/water_flow.png", f["core"] + "default_water_flowing_animated.png"}, //special attention
		{"block/water_still.png", f["core"] + "default_water_source_animated.png"}, //special attention
		{"block/oak_planks.png", f["core"] + "default_wood.png"},
		{"block/andesite.png", f["core"] + "mcl_core_andesite.png"},
		{"block/polished_andesite.png", f["core"] + "mcl_core_andesite_smooth.png"},
		{"item/golden_apple.png", f["core"] + "mcl_core_apple_golden.png"},
		{"item/barrier.png", f["core"] + "mcl_core_barrier.png"},
		{"block/bedrock.png", f["core"] + "mcl_core_bedrock.png"},
		{"block/bone_block_side.png", f["core"] + "mcl_core_bone_block_side.png"},
		{"block/bone_block_top.png", f["core"] + "mcl_core_bone_block_top.png"},
		{"item/bowl.png", f["core"] + "mcl_core_bowl.png"},
		{"block/cactus_bottom.png", f["core"] + "mcl_core_cactus_bottom.png"},
		{"block/cactus_side.png", f["core"] + "mcl_core_cactus_side.png"},
		{"block/cactus_top.png", f["core"] + "mcl_core_cactus_top.png"},
		{"item/charcoal.png", f["core"] + "mcl_core_charcoal.png"},
		{"block/coal_ore.png", f["core"] + "mcl_core_coal_ore.png"},
		{"block/coarse_dirt.png", f["core"] + "mcl_core_coarse_dirt.png"},
		{"block/crying_obsidian.png", f["core"] + "mcl_core_crying_obsidian.png"}, //special attention? might be ok
		//{"block/", f["core"] + "mcl_core_crying_obsidian_tear.png"},		//no match?
		//{"block/", f["core"] + "mcl_core_crying_obsidian_tear2.png"},		//no match?
		//{"block/", f["core"] + "mcl_core_crying_obsidian_tear3.png"},		//no match?
		{"block/diamond_ore.png", f["core"] + "mcl_core_diamond_ore.png"},
		{"block/diorite.png", f["core"] + "mcl_core_diorite.png"},
		{"block/polished_diorite.png", f["core"] + "mcl_core_diorite_smooth.png"},
		{"block/podzol_side.png", f["core"] + "mcl_core_dirt_podzol_side.png"},
		{"block/podzol_top.png", f["core"] + "mcl_core_dirt_podzol_top.png"},
		{"item/emerald.png", f["core"] + "mcl_core_emerald.png"},
		{"block/emerald_block.png", f["core"] + "mcl_core_emerald_block.png"},
		{"block/emerald_ore.png", f["core"] + "mcl_core_emerald_ore.png"},
		{"block/frosted_ice_0.png", f["core"] + "mcl_core_frosted_ice_0.png"},
		{"block/frosted_ice_1.png", f["core"] + "mcl_core_frosted_ice_1.png"},
		{"block/frosted_ice_2.png", f["core"] + "mcl_core_frosted_ice_2.png"},
		{"block/frosted_ice_3.png", f["core"] + "mcl_core_frosted_ice_3.png"},
		//Glass TODO All glass has normal and "detail" textures. Must create "detail".
		{"block/black_stained_glass.png", f["core"] + "mcl_core_glass_black.png"},
		{"block/black_stained_glass.png", f["core"] + "mcl_core_glass_black_detail.png"},
		{"block/blue_stained_glass.png", f["core"] + "mcl_core_glass_blue.png"},
		{"block/blue_stained_glass.png", f["core"] + "mcl_core_glass_blue_detail.png"},
		{"block/brown_stained_glass.png", f["core"] + "mcl_core_glass_brown.png"},
		{"block/brown_stained_glass.png", f["core"] + "mcl_core_glass_brown_detail.png"},
		{"block/cyan_stained_glass.png", f["core"] + "mcl_core_glass_cyan.png"},
		{"block/cyan_stained_glass.png", f["core"] + "mcl_core_glass_cyan_detail.png"},
		{"block/gray_stained_glass.png", f["core"] + "mcl_core_glass_gray.png"},
		{"block/gray_stained_glass.png", f["core"] + "mcl_core_glass_gray_detail.png"},
		{"block/green_stained_glass.png", f["core"] + "mcl_core_glass_green.png"},
		{"block/green_stained_glass.png", f["core"] + "mcl_core_glass_green_detail.png"},
		{"block/light_blue_stained_glass.png", f["core"] + "mcl_core_glass_light_blue.png"},
		{"block/light_blue_stained_glass.png", f["core"] + "mcl_core_glass_light_blue_detail.png"},
		{"block/lime_stained_glass.png", f["core"] + "mcl_core_glass_lime.png"},
		{"block/lime_stained_glass.png", f["core"] + "mcl_core_glass_lime_detail.png"},
		{"block/magenta_stained_glass.png", f["core"] + "mcl_core_glass_magenta.png"},
		{"block/magenta_stained_glass.png", f["core"] + "mcl_core_glass_magenta_detail.png"},
		{"block/orange_stained_glass.png", f["core"] + "mcl_core_glass_orange.png"},
		{"block/orange_stained_glass.png", f["core"] + "mcl_core_glass_orange_detail.png"},
		{"block/pink_stained_glass.png", f["core"] + "mcl_core_glass_pink.png"},
		{"block/pink_stained_glass.png", f["core"] + "mcl_core_glass_pink_detail.png"},
		{"block/purple_stained_glass.png", f["core"] + "mcl_core_glass_purple.png"},
		{"block/purple_stained_glass.png", f["core"] + "mcl_core_glass_purple_detail.png"},
		{"block/red_stained_glass.png", f["core"] + "mcl_core_glass_red.png"},
		{"block/red_stained_glass.png", f["core"] + "mcl_core_glass_red_detail.png"},
		{"block/light_gray_stained_glass.png", f["core"] + "mcl_core_glass_silver.png"},
		{"block/light_gray_stained_glass.png", f["core"] + "mcl_core_glass_silver_detail.png"},
		{"block/white_stained_glass.png", f["core"] + "mcl_core_glass_white.png"},
		{"block/white_stained_glass.png", f["core"] + "mcl_core_glass_white_detail.png"},
		{"block/yellow_stained_glass.png", f["core"] + "mcl_core_glass_yellow.png"},
		{"block/yellow_stained_glass.png", f["core"] + "mcl_core_glass_yellow_detail.png"},
		//Glass TODO
		{"item/gold_nugget.png", f["core"] + "mcl_core_gold_nugget.png"},
		{"block/", f["core"] + ""},

		// mcl_crafting_table
		// mcl_crimson
		// mcl_deepslate
		// mcl_doors
		// mcl_dyes
		// mcl_enchanting
		// mcl_end
		// mcl_farming
		// mcl_fences
		// mcl_fire
		// mcl_fireworks
		// mcl_fishing
		// mcl_fletching_table
		// mcl_flowerpots
		// mcl_flowers
		// mcl_furnaces
		{"block/furnace_top.png", f["furnaces"] + "default_furnace_bottom.png"}, //Minecraft does not have this texture.
		{"block/furnace_front.png", f["furnaces"] + "default_furnace_front.png"},
		{"block/furnace_front_on.png", f["furnaces"] + "default_furnace_front_active.png"},
		{"block/furnace_side.png", f["furnaces"] + "default_furnace_side.png"},
		{"block/furnace_top.png", f["furnaces"] + "default_furnace_top.png"},
		// mcl_grindstone
		// mcl_heads
		// mcl_honey
		// mcl_hoppers
		// mcl_itemframes
		// mcl_jukebox
		// mcl_lanterns
		// mcl_lectern
		// mcl_lightning_rods
		// mcl_loom
		// mcl_lush_caves
		// mcl_mangrove
		// mcl_maps
		// mcl_mobitems
		// mcl_mobspawners
		// mcl_monster_eggs
		// mcl_mud
		{"block/mud.png", f["mud"] + "mcl_mud.png"},
		{"block/mud_bricks.png", f["mud"] + "mcl_mud_bricks.png"},
		{"block/packed_mud.png", f["mud"] + "mcl_mud_packed_mud.png"},
		// mcl_mushrooms
		{"block/brown_mushroom.png", f["mushrooms"] + "farming_mushroom_brown.png"},
		{"block/red_mushroom.png", f["mushrooms"] + "farming_mushroom_red.png"},
		{"item/mushroom_stew.png", f["mushrooms"] + "farming_mushroom_stew.png"},
		{"block/mushroom_block_inside.png", f["mushrooms"] + "mcl_mushrooms_mushroom_block_inside.png"},
		{"block/brown_mushroom_block.png", f["mushrooms"] + "mcl_mushrooms_mushroom_block_skin_brown.png"},
		{"block/red_mushroom_block.png", f["mushrooms"] + "mcl_mushrooms_mushroom_block_skin_red.png"},
		{"block/mushroom_stem.png", f["mushrooms"] + "mcl_mushrooms_mushroom_block_skin_stem.png"},
		// mcl_nether
		// mcl_ocean
		// mcl_panes
		// mcl_portals
		// mcl_potions
		// mcl_pottery_sherds
		// mcl_raw_ores
		{"item/raw_gold.png", f["raw_ores"] + "mcl_raw_ores_raw_gold.png"},
		{"block/raw_gold_block.png", f["raw_ores"] + "mcl_raw_ores_raw_gold_block.png"},
		{"item/raw_iron.png", f["raw_ores"] + "mcl_raw_ores_raw_iron.png"},
		{"block/raw_iron_block.png", f["raw_ores"] + "mcl_raw_ores_raw_iron_block.png"},
		// mcl_sculk
		// mcl_shields
		// mcl_signs
		// mcl_smithing_table
		{"block/smithing_table_bottom.png", f["smithing_table"] + "mcl_smithing_table_bottom.png"},
		{"block/smithing_table_front.png", f["smithing_table"] + "mcl_smithing_table_front.png"},
		{"item/empty_slot_smithing_template_armor_trim.png", f["smithing_table"] + "mcl_smithing_table_inventory_trim_bg.png"},
		{"block/smithing_table_side.png", f["smithing_table"] + "mcl_smithing_table_side.png"},
		{"block/smithing_table_top.png", f["smithing_table"] + "mcl_smithing_table_top.png"},
		// mcl_smoker
		{"block/smoker_bottom.png", f["smoker"] + "smoker_bottom.png"},
		{"block/smoker_front.png", f["smoker"] + "smoker_front.png"},
		{"block/smoker_front_on.png", f["smoker"] + "smoker_front_on.png"},
		{"block/smoker_side.png", f["smoker"] + "smoker_side.png"},
		{"block/smoker_top.png", f["smoker"] + "smoker_top.png"},
		// mcl_sponges
		// mcl_spyglass
		// mcl_stairs
		// mcl_stonecutter
		// mcl_sus_nodes
		// mcl_sus_stew
		{"item/suspicious_stew.png", f["sus_stew"] + "sus_stew.png"},
		// mcl_throwing
		// mcl_tnt
		// mcl_tools
		// mcl_torches
		// mcl_totems
		// mcl_trees
		// mcl_walls
		// mcl_wool
		// mclx_core
		// mclx_fences
		// mclx_stairs
		// REDSTONE
		// screwdriver
		//{"", f[""] + ""},
	}

	for _, e := range blocksAndItems {
		copyTexture(srcLocation+e[0], outName+e[1])
	}

}

func copyTexture(src string, dest string) {
	source, err := os.Open(src)
	if err != nil {
		fmt.Printf(err.Error() + " //// Copy for texture skipped.\n")
		return
	}
	defer source.Close()
	destination, err := os.Create(dest)
	if err != nil {
		fmt.Printf(err.Error() + " //// Copy for texture skipped.\n")
		return
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err)
	}
}
