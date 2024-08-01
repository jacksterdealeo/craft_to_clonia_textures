package main

import (
	"errors"
	"fmt"
	imaging "github.com/disintegration/imaging"
	"image"
	"image/color"
	_ "image/png"
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
				log.Panicf("Folder %s is an \"invalid argument\". Maybe rename %s?\n", outName, inName)
			} else if errors.Is(err, fs.ErrPermission) {
				log.Panicf("Permission was denied. %s was not made.\n", outName)
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
description = A Minecraft texture pack converted to Mineclonia
author = Unknown
release = 01`, inName, outName)

	fmt.Printf("Pack info:\n%s\n", packConfigFile)
	err := os.WriteFile(outName+"/texture_pack.conf", []byte(packConfigFile), 0644)
	if err != nil {
		log.Panic(err)
	}

	// pack icon
	//copyTexture(inName+"/pack.png", outName+"/screenshot.png")
	if src, err := imaging.Open(inName + "/pack.png"); err != nil {
		fmt.Println("Pack icon error~")
	} else {
		background := imaging.Fill(src, 350, 233, imaging.Center, imaging.Lanczos)
		background = imaging.Blur(background, 10)
		foreground := imaging.Resize(src, 233, 0, imaging.Lanczos)

		dst := imaging.New(350, 233, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, background, image.Pt(0, 0))
		dst = imaging.OverlayCenter(dst, foreground, 1.0)
		err = imaging.Save(dst, outName+"/screenshot.png")
	}

	g := map[string]string{
		"block":    "/assets/minecraft/textures/block/",
		"item":     "/assets/minecraft/textures/item/",
		"particle": "/assets/minecraft/textures/particle/",
	}

	f := map[string]string{
		"amethyst":       "/ITEMS/mcl_amethyst/textures/",
		"anvils":         "/ITEMS/mcl_anvils/textures/",
		"blast_furnace":  "/ITEMS/mcl_blast_furnace/textures/",
		"copper":         "/ITEMS/mcl_copper/textures/",
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

	//srcLocation := inName + "/assets/minecraft/textures/"
	blocksAndItems := [...][2]string{
		// mcl_amethyst
		{g["block"] + "amethyst_block.png", f["amethyst"] + "mcl_amethyst_amethyst_block.png"},
		{g["block"] + "large_amethyst_bud.png", f["amethyst"] + "mcl_amethyst_amethyst_bud_large.png"},
		{g["block"] + "medium_amethyst_bud.png", f["amethyst"] + "mcl_amethyst_amethyst_bud_medium.png"},
		{g["block"] + "small_amethyst_bud.png", f["amethyst"] + "mcl_amethyst_amethyst_bud_small.png"},
		{g["block"] + "amethyst_cluster.png", f["amethyst"] + "mcl_amethyst_amethyst_cluster.png"},
		{g["item"] + "amethyst_shard.png", f["amethyst"] + "mcl_amethyst_amethyst_shard.png"},
		{g["block"] + "budding_amethyst.png", f["amethyst"] + "mcl_amethyst_budding_amethyst.png"},
		{g["block"] + "calcite.png", f["amethyst"] + "mcl_amethyst_calcite_block.png"},
		{g["block"] + "tinted_glass.png", f["amethyst"] + "mcl_amethyst_tinted_glass.png"},
		// mcl_anvils
		// FIX: anvil damaged texture is not wide enough for all the model.
		{g["block"] + "anvil.png", f["anvils"] + "mcl_anvils_anvil_base.png"},
		{g["block"] + "anvil.png", f["anvils"] + "mcl_anvils_anvil_side.png"},
		//{g["block"] + "anvil_top.png", f["anvils"] + "mcl_anvils_anvil_top_damaged_0.png"},
		//{g["block"] + "chipped_anvil_top.png", f["anvils"] + "mcl_anvils_anvil_top_damaged_1.png"},
		//{g["block"] + "damaged_anvil_top.png", f["anvils"] + "mcl_anvils_anvil_top_damaged_2.png"},
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
		{g["block"] + "blast_furnace_front.png", f["blast_furnace"] + "blast_furnace_front.png"},
		{g["block"] + "blast_furnace_front_on.png", f["blast_furnace"] + "blast_furnace_front_on.png"},
		{g["block"] + "blast_furnace_side.png", f["blast_furnace"] + "blast_furnace_side.png"},
		{g["block"] + "blast_furnace_top.png", f["blast_furnace"] + "blast_furnace_top.png"},
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
		//{g["block"] + "", f["copper"] + "mcl_copper_anti_oxidation_particle.png"}, //no match?
		{g["block"] + "copper_block.png", f["copper"] + "mcl_copper_block.png"},
		{g["block"] + "copper_bulb.png", f["copper"] + "mcl_copper_block_bulb_off.png"},
		{g["block"] + "copper_bulb_lit.png", f["copper"] + "mcl_copper_block_bulb_on.png"},
		{g["block"] + "copper_bulb_powered.png", f["copper"] + "mcl_copper_block_bulb_powered_off.png"},
		{g["block"] + "copper_bulb_lit_powered.png", f["copper"] + "mcl_copper_block_bulb_powered_on.png"},
		{g["block"] + "chiseled_copper.png", f["copper"] + "mcl_copper_block_chiseled.png"},
		{g["block"] + "cut_copper.png", f["copper"] + "mcl_copper_block_cut.png"},
		{g["block"] + "copper_grate.png", f["copper"] + "mcl_copper_block_grate.png"},
		{g["block"] + "raw_copper_block.png", f["copper"] + "mcl_copper_block_raw.png"},
		{g["item"] + "copper_door.png", f["copper"] + "mcl_copper_door.png"},
		{g["block"] + "copper_door_bottom.png", f["copper"] + "mcl_copper_door_bottom.png"},
		{g["item"] + "exposed_copper_door.png", f["copper"] + "mcl_copper_door_exposed.png"},
		{g["block"] + "exposed_copper_door_bottom.png", f["copper"] + "mcl_copper_door_exposed_bottom.png"},
		//{g["block"] + "", f["copper"] + ""},
		//{g["item"] + "", f["copper"] + ""},
		// mcl_core
		{g["block"] + "acacia_leaves.png", f["core"] + "default_acacia_leaves.png"},
		{g["block"] + "acacia_sapling.png", f["core"] + "default_acacia_sapling.png"},
		{g["block"] + "acacia_log.png", f["core"] + "default_acacia_tree.png"},
		{g["block"] + "acacia_log_top.png", f["core"] + "default_acacia_tree_top.png"},
		{g["block"] + "acacia_planks.png", f["core"] + "default_acacia_wood.png"},
		{g["item"] + "apple.png", f["core"] + "default_apple.png"},
		{g["block"] + "bricks.png", f["core"] + "default_brick.png"},
		{g["block"] + "clay.png", f["core"] + "default_clay.png"},
		{g["item"] + "brick.png", f["core"] + "default_clay_brick.png"},
		{g["item"] + "clay_ball.png", f["core"] + "default_clay_lump.png"},
		{g["block"] + "coal_block.png", f["core"] + "default_coal_block.png"},
		{g["item"] + "coal.png", f["core"] + "default_coal_lump.png"},
		{g["block"] + "cobblestone.png", f["core"] + "default_cobble.png"},
		{g["item"] + "diamond.png", f["core"] + "default_diamond.png"},
		{g["block"] + "diamond_block.png", f["core"] + "default_diamond_block.png"},
		{g["block"] + "dirt.png", f["core"] + "default_dirt.png"},
		{g["block"] + "dead_bush.png", f["core"] + "default_dry_shrub.png"},
		{g["item"] + "flint.png", f["core"] + "default_flint.png"},
		{g["block"] + "glass.png", f["core"] + "default_glass.png"},
		//{g["block"] + "", f["core"] + "default_glass_detail.png"},		//no match?
		{g["block"] + "gold_block.png", f["core"] + "default_gold_block.png"},
		{g["item"] + "gold_ingot.png", f["core"] + "default_gold_ingot.png"},
		{g["block"] + "gravel.png", f["core"] + "default_gravel.png"},
		{g["block"] + "ice.png", f["core"] + "default_ice.png"},
		{g["block"] + "jungle_leaves.png", f["core"] + "default_jungleleaves.png"},
		{g["block"] + "jungle_sapling.png", f["core"] + "default_junglesapling.png"},
		{g["block"] + "jungle_log.png", f["core"] + "default_jungletree.png"},
		{g["block"] + "jungle_log_top.png", f["core"] + "default_jungletree_top.png"},
		{g["block"] + "jungle_planks.png", f["core"] + "default_junglewood.png"},
		{g["block"] + "ladder.png", f["core"] + "default_ladder.png"},
		{g["block"] + "lava_flow.png", f["core"] + "default_lava_flowing_animated.png"}, //special attention
		{g["block"] + "lava_still.png", f["core"] + "default_lava_source_animated.png"}, //special attention
		{g["block"] + "oak_leaves.png", f["core"] + "default_leaves.png"},
		{g["block"] + "mossy_cobblestone.png", f["core"] + "default_mossycobble.png"},
		{g["block"] + "obsidian.png", f["core"] + "default_obsidian.png"},
		{g["item"] + "paper.png", f["core"] + "default_paper.png"},
		{g["block"] + "sand.png", f["core"] + "default_sand.png"},
		{g["block"] + "oak_sapling.png", f["core"] + "default_sapling.png"},
		{g["block"] + "snow.png", f["core"] + "default_snow.png"},
		{g["block"] + "iron_block.png", f["core"] + "default_steel_block.png"},
		{g["item"] + "iron_ingot.png", f["core"] + "default_steel_ingot.png"},
		{g["item"] + "stick.png", f["core"] + "default_stick.png"},
		{g["block"] + "stone_bricks.png", f["core"] + "default_stone_brick.png"},
		{g["block"] + "oak_log.png", f["core"] + "default_tree.png"},
		{g["block"] + "oak_log_top.png", f["core"] + "default_tree_top.png"},
		{g["block"] + "water_flow.png", f["core"] + "default_water_flowing_animated.png"}, //special attention
		{g["block"] + "water_still.png", f["core"] + "default_water_source_animated.png"}, //special attention
		{g["block"] + "oak_planks.png", f["core"] + "default_wood.png"},
		{g["block"] + "andesite.png", f["core"] + "mcl_core_andesite.png"},
		{g["block"] + "polished_andesite.png", f["core"] + "mcl_core_andesite_smooth.png"},
		{g["item"] + "golden_apple.png", f["core"] + "mcl_core_apple_golden.png"},
		{g["item"] + "barrier.png", f["core"] + "mcl_core_barrier.png"},
		{g["block"] + "bedrock.png", f["core"] + "mcl_core_bedrock.png"},
		{g["block"] + "bone_block_side.png", f["core"] + "mcl_core_bone_block_side.png"},
		{g["block"] + "bone_block_top.png", f["core"] + "mcl_core_bone_block_top.png"},
		{g["item"] + "bowl.png", f["core"] + "mcl_core_bowl.png"},
		{g["block"] + "cactus_bottom.png", f["core"] + "mcl_core_cactus_bottom.png"},
		{g["block"] + "cactus_side.png", f["core"] + "mcl_core_cactus_side.png"},
		{g["block"] + "cactus_top.png", f["core"] + "mcl_core_cactus_top.png"},
		{g["item"] + "charcoal.png", f["core"] + "mcl_core_charcoal.png"},
		{g["block"] + "coal_ore.png", f["core"] + "mcl_core_coal_ore.png"},
		{g["block"] + "coarse_dirt.png", f["core"] + "mcl_core_coarse_dirt.png"},
		{g["block"] + "crying_obsidian.png", f["core"] + "mcl_core_crying_obsidian.png"}, //special attention? might be ok
		//{g["block"] + "", f["core"] + "mcl_core_crying_obsidian_tear.png"},		//no match?
		//{g["block"] + "", f["core"] + "mcl_core_crying_obsidian_tear2.png"},		//no match?
		//{g["block"] + "", f["core"] + "mcl_core_crying_obsidian_tear3.png"},		//no match?
		{g["block"] + "diamond_ore.png", f["core"] + "mcl_core_diamond_ore.png"},
		{g["block"] + "diorite.png", f["core"] + "mcl_core_diorite.png"},
		{g["block"] + "polished_diorite.png", f["core"] + "mcl_core_diorite_smooth.png"},
		{g["block"] + "podzol_side.png", f["core"] + "mcl_core_dirt_podzol_side.png"},
		{g["block"] + "podzol_top.png", f["core"] + "mcl_core_dirt_podzol_top.png"},
		{g["item"] + "emerald.png", f["core"] + "mcl_core_emerald.png"},
		{g["block"] + "emerald_block.png", f["core"] + "mcl_core_emerald_block.png"},
		{g["block"] + "emerald_ore.png", f["core"] + "mcl_core_emerald_ore.png"},
		{g["block"] + "frosted_ice_0.png", f["core"] + "mcl_core_frosted_ice_0.png"},
		{g["block"] + "frosted_ice_1.png", f["core"] + "mcl_core_frosted_ice_1.png"},
		{g["block"] + "frosted_ice_2.png", f["core"] + "mcl_core_frosted_ice_2.png"},
		{g["block"] + "frosted_ice_3.png", f["core"] + "mcl_core_frosted_ice_3.png"},
		//Glass TODO: All glass has normal and "detail" textures. Must create "detail".
		{g["block"] + "black_stained_glass.png", f["core"] + "mcl_core_glass_black.png"},
		//{g["block"] + "black_stained_glass.png", f["core"] + "mcl_core_glass_black_detail.png"},
		{g["block"] + "blue_stained_glass.png", f["core"] + "mcl_core_glass_blue.png"},
		{g["block"] + "blue_stained_glass.png", f["core"] + "mcl_core_glass_blue_detail.png"},
		{g["block"] + "brown_stained_glass.png", f["core"] + "mcl_core_glass_brown.png"},
		{g["block"] + "brown_stained_glass.png", f["core"] + "mcl_core_glass_brown_detail.png"},
		{g["block"] + "cyan_stained_glass.png", f["core"] + "mcl_core_glass_cyan.png"},
		{g["block"] + "cyan_stained_glass.png", f["core"] + "mcl_core_glass_cyan_detail.png"},
		{g["block"] + "gray_stained_glass.png", f["core"] + "mcl_core_glass_gray.png"},
		{g["block"] + "gray_stained_glass.png", f["core"] + "mcl_core_glass_gray_detail.png"},
		{g["block"] + "green_stained_glass.png", f["core"] + "mcl_core_glass_green.png"},
		{g["block"] + "green_stained_glass.png", f["core"] + "mcl_core_glass_green_detail.png"},
		{g["block"] + "light_blue_stained_glass.png", f["core"] + "mcl_core_glass_light_blue.png"},
		{g["block"] + "light_blue_stained_glass.png", f["core"] + "mcl_core_glass_light_blue_detail.png"},
		{g["block"] + "lime_stained_glass.png", f["core"] + "mcl_core_glass_lime.png"},
		{g["block"] + "lime_stained_glass.png", f["core"] + "mcl_core_glass_lime_detail.png"},
		{g["block"] + "magenta_stained_glass.png", f["core"] + "mcl_core_glass_magenta.png"},
		{g["block"] + "magenta_stained_glass.png", f["core"] + "mcl_core_glass_magenta_detail.png"},
		{g["block"] + "orange_stained_glass.png", f["core"] + "mcl_core_glass_orange.png"},
		{g["block"] + "orange_stained_glass.png", f["core"] + "mcl_core_glass_orange_detail.png"},
		{g["block"] + "pink_stained_glass.png", f["core"] + "mcl_core_glass_pink.png"},
		{g["block"] + "pink_stained_glass.png", f["core"] + "mcl_core_glass_pink_detail.png"},
		{g["block"] + "purple_stained_glass.png", f["core"] + "mcl_core_glass_purple.png"},
		{g["block"] + "purple_stained_glass.png", f["core"] + "mcl_core_glass_purple_detail.png"},
		{g["block"] + "red_stained_glass.png", f["core"] + "mcl_core_glass_red.png"},
		{g["block"] + "red_stained_glass.png", f["core"] + "mcl_core_glass_red_detail.png"},
		{g["block"] + "light_gray_stained_glass.png", f["core"] + "mcl_core_glass_silver.png"},
		{g["block"] + "light_gray_stained_glass.png", f["core"] + "mcl_core_glass_silver_detail.png"},
		{g["block"] + "white_stained_glass.png", f["core"] + "mcl_core_glass_white.png"},
		{g["block"] + "white_stained_glass.png", f["core"] + "mcl_core_glass_white_detail.png"},
		{g["block"] + "yellow_stained_glass.png", f["core"] + "mcl_core_glass_yellow.png"},
		{g["block"] + "yellow_stained_glass.png", f["core"] + "mcl_core_glass_yellow_detail.png"},
		//Glass TODO
		{g["item"] + "gold_nugget.png", f["core"] + "mcl_core_gold_nugget.png"},
		{g["block"] + "gold_ore.png", f["core"] + "mcl_core_gold_ore.png"},
		{g["block"] + "granite.png", f["core"] + "mcl_core_granite.png"},
		{g["block"] + "polished_granite.png", f["core"] + "mcl_core_granite_smooth.png"},
		{g["block"] + "grass_block_side_overlay.png", f["core"] + "mcl_core_grass_block_side_overlay.png"},
		{g["block"] + "grass_block_top.png", f["core"] + "mcl_core_grass_block_top.png"},
		{g["block"] + "dirt_path_side.png", f["core"] + "mcl_core_grass_path_side.png"},
		{g["block"] + "dirt_path_top.png", f["core"] + "mcl_core_grass_path_top.png"},
		{g["block"] + "grass_block_snow.png", f["core"] + "mcl_core_grass_side_snowed.png"},
		{g["block"] + "packed_ice.png", f["core"] + "mcl_core_ice_packed.png"},
		{g["item"] + "iron_nugget.png", f["core"] + "mcl_core_iron_nugget.png"},
		{g["block"] + "iron_ore.png", f["core"] + "mcl_core_iron_ore.png"},
		{g["item"] + "lapis_lazuli.png", f["core"] + "mcl_core_lapis.png"},
		{g["block"] + "lapis_block.png", f["core"] + "mcl_core_lapis_block.png"},
		{g["block"] + "lapis_ore.png", f["core"] + "mcl_core_lapis_ore.png"},
		{g["block"] + "dark_oak_leaves.png", f["core"] + "mcl_core_leaves_big_oak.png"},
		{g["block"] + "birch_leaves.png", f["core"] + "mcl_core_leaves_birch.png"},
		{g["block"] + "spruce_leaves.png", f["core"] + "mcl_core_leaves_spruce.png"},
		{g["item"] + "light_00.png", f["core"] + "mcl_core_light_0.png"},
		{g["item"] + "light_01.png", f["core"] + "mcl_core_light_1.png"},
		{g["item"] + "light_02.png", f["core"] + "mcl_core_light_2.png"},
		{g["item"] + "light_03.png", f["core"] + "mcl_core_light_3.png"},
		{g["item"] + "light_04.png", f["core"] + "mcl_core_light_4.png"},
		{g["item"] + "light_05.png", f["core"] + "mcl_core_light_5.png"},
		{g["item"] + "light_06.png", f["core"] + "mcl_core_light_6.png"},
		{g["item"] + "light_07.png", f["core"] + "mcl_core_light_7.png"},
		{g["item"] + "light_08.png", f["core"] + "mcl_core_light_8.png"},
		{g["item"] + "light_09.png", f["core"] + "mcl_core_light_9.png"},
		{g["item"] + "light_10.png", f["core"] + "mcl_core_light_10.png"},
		{g["item"] + "light_11.png", f["core"] + "mcl_core_light_11.png"},
		{g["item"] + "light_12.png", f["core"] + "mcl_core_light_12.png"},
		{g["item"] + "light_13.png", f["core"] + "mcl_core_light_13.png"},
		{g["item"] + "light_14.png", f["core"] + "mcl_core_light_14.png"}, //no light_15 in Mineclonia
		{g["block"] + "dark_oak_log.png", f["core"] + "mcl_core_log_big_oak.png"},
		{g["block"] + "dark_oak_log_top.png", f["core"] + "mcl_core_log_big_oak_top.png"},
		{g["block"] + "birch_log.png", f["core"] + "mcl_core_log_birch.png"},
		{g["block"] + "birch_log_top.png", f["core"] + "mcl_core_log_birch_top.png"},
		{g["block"] + "spruce_log.png", f["core"] + "mcl_core_log_spruce.png"},
		{g["block"] + "spruce_log_top.png", f["core"] + "mcl_core_log_spruce_top.png"},
		//{g["block"] + "", f["core"] + "mcl_core_mycelium_particle.png"}, 	//special attention
		{g["block"] + "mycelium_side.png", f["core"] + "mcl_core_mycelium_side.png"},
		{g["block"] + "mycelium_top.png", f["core"] + "mcl_core_mycelium_top.png"},
		//{g["block"] + "", f["core"] + "mcl_core_palette_grass.png"}, 	//special attention
		//{g["block"] + "", f["core"] + "mcl_core_palette_leaves.png"}, 		//special attention
		{g["block"] + "sugar_cane.png", f["core"] + "mcl_core_papyrus.png"},
		{g["block"] + "dark_oak_planks.png", f["core"] + "mcl_core_planks_big_oak.png"},
		{g["block"] + "birch_planks.png", f["core"] + "mcl_core_planks_birch.png"},
		{g["block"] + "spruce_planks.png", f["core"] + "mcl_core_planks_spruce.png"},
		{g["block"] + "red_sand.png", f["core"] + "mcl_core_red_sand.png"},
		{g["block"] + "red_sandstone_bottom.png", f["core"] + "mcl_core_red_sandstone_bottom.png"},
		{g["block"] + "chiseled_red_sandstone.png", f["core"] + "mcl_core_red_sandstone_carved.png"},
		{g["block"] + "red_sandstone.png", f["core"] + "mcl_core_red_sandstone_normal.png"},
		{g["block"] + "cut_red_sandstone.png", f["core"] + "mcl_core_red_sandstone_smooth.png"},
		{g["block"] + "red_sandstone_top.png", f["core"] + "mcl_core_red_sandstone_top.png"},
		{g["block"] + "redstone_ore.png", f["core"] + "mcl_core_redstone_ore.png"},
		{g["item"] + "sugar_cane.png", f["core"] + "mcl_core_reeds.png"},
		{g["block"] + "sandstone_bottom.png", f["core"] + "mcl_core_sandstone_bottom.png"},
		{g["block"] + "chiseled_sandstone.png", f["core"] + "mcl_core_sandstone_carved.png"},
		{g["block"] + "sandstone.png", f["core"] + "mcl_core_sandstone_normal.png"},
		{g["block"] + "cut_sandstone.png", f["core"] + "mcl_core_sandstone_smooth.png"},
		{g["block"] + "sandstone_top.png", f["core"] + "mcl_core_sandstone_top.png"},
		{g["block"] + "dark_oak_sapling.png", f["core"] + "mcl_core_sapling_big_oak.png"},
		{g["block"] + "birch_sapling.png", f["core"] + "mcl_core_sapling_birch.png"},
		{g["block"] + "spruce_sapling.png", f["core"] + "mcl_core_sapling_spruce.png"},
		{g["block"] + "slime_block.png", f["core"] + "mcl_core_slime.png"},
		{g["block"] + "chiseled_stone_bricks.png", f["core"] + "mcl_core_stonebrick_carved.png"},
		{g["block"] + "cracked_stone_bricks.png", f["core"] + "mcl_core_stonebrick_cracked.png"},
		{g["block"] + "mossy_stone_bricks.png", f["core"] + "mcl_core_stonebrick_mossy.png"},
		{g["block"] + "stripped_acacia_log.png", f["core"] + "mcl_core_stripped_acacia_side.png"},
		{g["block"] + "stripped_acacia_log_top.png", f["core"] + "mcl_core_stripped_acacia_top.png"},
		{g["block"] + "stripped_birch_log.png", f["core"] + "mcl_core_stripped_birch_side.png"},
		{g["block"] + "stripped_birch_log_top.png", f["core"] + "mcl_core_stripped_birch_top.png"},
		{g["block"] + "stripped_dark_oak_log.png", f["core"] + "mcl_core_stripped_dark_oak_side.png"},
		{g["block"] + "stripped_dark_oak_log_top.png", f["core"] + "mcl_core_stripped_dark_oak_top.png"},
		{g["block"] + "stripped_jungle_log.png", f["core"] + "mcl_core_stripped_jungle_side.png"},
		{g["block"] + "stripped_jungle_log_top.png", f["core"] + "mcl_core_stripped_jungle_top.png"},
		{g["block"] + "stripped_oak_log.png", f["core"] + "mcl_core_stripped_oak_side.png"},
		{g["block"] + "stripped_oak_log_top.png", f["core"] + "mcl_core_stripped_oak_top.png"},
		{g["block"] + "stripped_spruce_log.png", f["core"] + "mcl_core_stripped_spruce_side.png"},
		{g["block"] + "stripped_spruce_log_top.png", f["core"] + "mcl_core_stripped_spruce_top.png"},
		{g["item"] + "sugar.png", f["core"] + "mcl_core_sugar.png"},
		{g["block"] + "vine.png", f["core"] + "mcl_core_vine.png"}, //special attention
		//{g["block"] + "", f["core"] + "mcl_core_void.png"},         //no match
		{g["block"] + "cobweb.png", f["core"] + "mcl_core_web.png"},
		{g["block"] + "grass_block_side_overlay.png", f["core"] + "mcl_dirt_grass_shadow.png"}, //special attention?
		{g["particle"] + "lava.png", f["core"] + "mcl_particles_lava.png"},
		//{g["block"] + "", f["core"] + "mcl_stairs_andesite_smooth_slab.png"}, //special attention
		//{g["block"] + "", f["core"] + "mcl_stairs_diorite_smooth_slab.png"},  //special attention
		//{g["block"] + "", f["core"] + "mcl_stairs_granite_smooth_slab.png"},  //special attention

		//{g["block"] + "", f["core"] + ""},
		//{g["item"] + "", f["core"] + ""},

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
		{g["block"] + "furnace_top.png", f["furnaces"] + "default_furnace_bottom.png"}, //Minecraft does not have this texture.
		{g["block"] + "furnace_front.png", f["furnaces"] + "default_furnace_front.png"},
		{g["block"] + "furnace_front_on.png", f["furnaces"] + "default_furnace_front_active.png"},
		{g["block"] + "furnace_side.png", f["furnaces"] + "default_furnace_side.png"},
		{g["block"] + "furnace_top.png", f["furnaces"] + "default_furnace_top.png"},
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
		{g["block"] + "mud.png", f["mud"] + "mcl_mud.png"},
		{g["block"] + "mud_bricks.png", f["mud"] + "mcl_mud_bricks.png"},
		{g["block"] + "packed_mud.png", f["mud"] + "mcl_mud_packed_mud.png"},
		// mcl_mushrooms
		{g["block"] + "brown_mushroom.png", f["mushrooms"] + "farming_mushroom_brown.png"},
		{g["block"] + "red_mushroom.png", f["mushrooms"] + "farming_mushroom_red.png"},
		{g["item"] + "mushroom_stew.png", f["mushrooms"] + "farming_mushroom_stew.png"},
		{g["block"] + "mushroom_block_inside.png", f["mushrooms"] + "mcl_mushrooms_mushroom_block_inside.png"},
		{g["block"] + "brown_mushroom_block.png", f["mushrooms"] + "mcl_mushrooms_mushroom_block_skin_brown.png"},
		{g["block"] + "red_mushroom_block.png", f["mushrooms"] + "mcl_mushrooms_mushroom_block_skin_red.png"},
		{g["block"] + "mushroom_stem.png", f["mushrooms"] + "mcl_mushrooms_mushroom_block_skin_stem.png"},
		// mcl_nether
		// mcl_ocean
		// mcl_panes
		// mcl_portals
		// mcl_potions
		// mcl_pottery_sherds
		// mcl_raw_ores
		{g["item"] + "raw_gold.png", f["raw_ores"] + "mcl_raw_ores_raw_gold.png"},
		{g["block"] + "raw_gold_block.png", f["raw_ores"] + "mcl_raw_ores_raw_gold_block.png"},
		{g["item"] + "raw_iron.png", f["raw_ores"] + "mcl_raw_ores_raw_iron.png"},
		{g["block"] + "raw_iron_block.png", f["raw_ores"] + "mcl_raw_ores_raw_iron_block.png"},
		// mcl_sculk
		// mcl_shields
		// mcl_signs
		// mcl_smithing_table
		{g["block"] + "smithing_table_bottom.png", f["smithing_table"] + "mcl_smithing_table_bottom.png"},
		{g["block"] + "smithing_table_front.png", f["smithing_table"] + "mcl_smithing_table_front.png"},
		{g["item"] + "empty_slot_smithing_template_armor_trim.png", f["smithing_table"] + "mcl_smithing_table_inventory_trim_bg.png"},
		{g["block"] + "smithing_table_side.png", f["smithing_table"] + "mcl_smithing_table_side.png"},
		{g["block"] + "smithing_table_top.png", f["smithing_table"] + "mcl_smithing_table_top.png"},
		// mcl_smoker
		{g["block"] + "smoker_bottom.png", f["smoker"] + "smoker_bottom.png"},
		{g["block"] + "smoker_front.png", f["smoker"] + "smoker_front.png"},
		{g["block"] + "smoker_front_on.png", f["smoker"] + "smoker_front_on.png"},
		{g["block"] + "smoker_side.png", f["smoker"] + "smoker_side.png"},
		{g["block"] + "smoker_top.png", f["smoker"] + "smoker_top.png"},
		// mcl_sponges
		// mcl_spyglass
		// mcl_stairs
		// mcl_stonecutter
		// mcl_sus_nodes
		// mcl_sus_stew
		{g["item"] + "suspicious_stew.png", f["sus_stew"] + "sus_stew.png"},
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
		copyTexture(inName+e[0], outName+e[1])
	}

	//special casses
	////anvil
	anvil := func() {
		abase, err := imaging.Open(inName + g["block"] + "anvil.png")
		if err != nil {
			fmt.Println("AnvilBase error~", g["block"]+"anvil.png")
			return
		}
		a0, err := imaging.Open(inName + g["block"] + "anvil_top.png")
		if err != nil {
			fmt.Println("Anvil0 error~")
			return
		}
		a1, err := imaging.Open(inName + g["block"] + "chipped_anvil_top.png")
		if err != nil {
			fmt.Println("Anvil1 error~")
			return
		}
		a2, err := imaging.Open(inName + g["block"] + "damaged_anvil_top.png")
		if err != nil {
			fmt.Println("Anvil2 error~")
			return
		}
		anvilX := abase.Bounds().Dx()
		anvilY := abase.Bounds().Dy()

		dst := imaging.New(anvilX, anvilY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, abase, image.Pt(0, 0))
		dst = imaging.OverlayCenter(dst, a0, 1.0)

		if err = imaging.Save(dst, outName+f["anvils"]+"mcl_anvils_anvil_top_damaged_0.png"); err != nil {
			fmt.Println("Anvil undamaged failed!")
		}
		dst = imaging.OverlayCenter(dst, a1, 1.0)
		if err = imaging.Save(dst, outName+f["anvils"]+"mcl_anvils_anvil_top_damaged_1.png"); err != nil {
			fmt.Println("Anvil damaged1 failed!")
		}
		dst = imaging.OverlayCenter(dst, a2, 1.0)
		if err = imaging.Save(dst, outName+f["anvils"]+"mcl_anvils_anvil_top_damaged_2.png"); err != nil {
			fmt.Println("Anvil damaged2 failed!")
		}
	}
	anvil()

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
