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

	craftPaths := map[string]string{
		"block":    "/assets/minecraft/textures/block/",
		"item":     "/assets/minecraft/textures/item/",
		"particle": "/assets/minecraft/textures/particle/",
	}

	cloniaPaths := map[string]string{
		"amethyst":       "/ITEMS/mcl_amethyst/textures/",
		"anvils":         "/ITEMS/mcl_anvils/textures/",
		"blast_furnace":  "/ITEMS/mcl_blast_furnace/textures/",
		"copper":         "/ITEMS/mcl_copper/textures/",
		"core":           "/ITEMS/mcl_core/textures/",
		"deepslate":      "/ITEMS/mcl_deepslate/textures/",
		"furnaces":       "/ITEMS/mcl_furnaces/textures/",
		"mud":            "/ITEMS/mcl_mud/textures/",
		"mushrooms":      "/ITEMS/mcl_mushrooms/textures/",
		"raw_ores":       "/ITEMS/mcl_raw_ores/textures/",
		"smithing_table": "/ITEMS/mcl_smithing_table/textures/",
		"smoker":         "/ITEMS/mcl_smoker/textures/",
		"sus_stew":       "/ITEMS/mcl_sus_stew/textures/",
	}
	for _, v := range cloniaPaths {
		if err := os.MkdirAll(outName+v, 0755); err != nil {
			log.Panic(err)
		}
	}

	//srcLocation := inName + "/assets/minecraft/textures/"
	blocksAndItems := [...][4]string{
		// mcl_amethyst
		{"block", "amethyst_block.png", "amethyst", "mcl_amethyst_amethyst_block.png"},
		{"block", "large_amethyst_bud.png", "amethyst", "mcl_amethyst_amethyst_bud_large.png"},
		{"block", "medium_amethyst_bud.png", "amethyst", "mcl_amethyst_amethyst_bud_medium.png"},
		{"block", "small_amethyst_bud.png", "amethyst", "mcl_amethyst_amethyst_bud_small.png"},
		{"block", "amethyst_cluster.png", "amethyst", "mcl_amethyst_amethyst_cluster.png"},
		{"item", "amethyst_shard.png", "amethyst", "mcl_amethyst_amethyst_shard.png"},
		{"block", "budding_amethyst.png", "amethyst", "mcl_amethyst_budding_amethyst.png"},
		{"block", "calcite.png", "amethyst", "mcl_amethyst_calcite_block.png"},
		{"block", "tinted_glass.png", "amethyst", "mcl_amethyst_tinted_glass.png"},
		// mcl_anvils
		// FIX: anvil damaged texture is not wide enough for all the model.
		{"block", "anvil.png", "anvils", "mcl_anvils_anvil_base.png"},
		{"block", "anvil.png", "anvils", "mcl_anvils_anvil_side.png"},
		//{"block", "anvil_top.png", "anvils", "mcl_anvils_anvil_top_damaged_0.png"},
		//{"block", "chipped_anvil_top.png", "anvils", "mcl_anvils_anvil_top_damaged_1.png"},
		//{"block", "damaged_anvil_top.png", "anvils", "mcl_anvils_anvil_top_damaged_2.png"},
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
		{"block", "blast_furnace_front.png", "blast_furnace", "blast_furnace_front.png"},
		{"block", "blast_furnace_front_on.png", "blast_furnace", "blast_furnace_front_on.png"},
		{"block", "blast_furnace_side.png", "blast_furnace", "blast_furnace_side.png"},
		{"block", "blast_furnace_top.png", "blast_furnace", "blast_furnace_top.png"},
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
		//{"block", "", "copper", "mcl_copper_anti_oxidation_particle.png"}, //no match?
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
		// mcl_core
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
		//{"block", "", "core", "default_glass_detail.png"},		//no match?
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
		{"block", "lava_flow.png", "core", "default_lava_flowing_animated.png"}, //special attention
		{"block", "lava_still.png", "core", "default_lava_source_animated.png"}, //special attention
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
		//{"block", "", "core", "mcl_core_crying_obsidian_tear.png"},		//no match?
		//{"block", "", "core", "mcl_core_crying_obsidian_tear2.png"},		//no match?
		//{"block", "", "core", "mcl_core_crying_obsidian_tear3.png"},		//no match?
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
		//{"block", "black_stained_glass.png", "core", "mcl_core_glass_black_detail.png"},
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
		{"block", "vine.png", "core", "mcl_core_vine.png"}, //special attention
		//{"block", "", "core", "mcl_core_void.png"},         //no match
		{"block", "cobweb.png", "core", "mcl_core_web.png"},
		{"block", "grass_block_side_overlay.png", "core", "mcl_dirt_grass_shadow.png"}, //special attention?
		{"particle", "lava.png", "core", "mcl_particles_lava.png"},
		//{"block", "", "core", "mcl_stairs_andesite_smooth_slab.png"}, //special attention
		//{"block", "", "core", "mcl_stairs_diorite_smooth_slab.png"},  //special attention
		//{"block", "", "core", "mcl_stairs_granite_smooth_slab.png"},  //special attention

		//{"block", "", "core", ""},
		//{"item", "", "core", ""},

		// mcl_crafting_table
		// mcl_crimson
		// mcl_deepslate
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
		{"block", "furnace_top.png", "furnaces", "default_furnace_bottom.png"}, //Minecraft does not have this texture.
		{"block", "furnace_front.png", "furnaces", "default_furnace_front.png"},
		{"block", "furnace_front_on.png", "furnaces", "default_furnace_front_active.png"},
		{"block", "furnace_side.png", "furnaces", "default_furnace_side.png"},
		{"block", "furnace_top.png", "furnaces", "default_furnace_top.png"},
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
		{"block", "mud.png", "mud", "mcl_mud.png"},
		{"block", "mud_bricks.png", "mud", "mcl_mud_bricks.png"},
		{"block", "packed_mud.png", "mud", "mcl_mud_packed_mud.png"},
		// mcl_mushrooms
		{"block", "brown_mushroom.png", "mushrooms", "farming_mushroom_brown.png"},
		{"block", "red_mushroom.png", "mushrooms", "farming_mushroom_red.png"},
		{"item", "mushroom_stew.png", "mushrooms", "farming_mushroom_stew.png"},
		{"block", "mushroom_block_inside.png", "mushrooms", "mcl_mushrooms_mushroom_block_inside.png"},
		{"block", "brown_mushroom_block.png", "mushrooms", "mcl_mushrooms_mushroom_block_skin_brown.png"},
		{"block", "red_mushroom_block.png", "mushrooms", "mcl_mushrooms_mushroom_block_skin_red.png"},
		{"block", "mushroom_stem.png", "mushrooms", "mcl_mushrooms_mushroom_block_skin_stem.png"},
		// mcl_nether
		// mcl_ocean
		// mcl_panes
		// mcl_portals
		// mcl_potions
		// mcl_pottery_sherds
		// mcl_raw_ores
		{"item", "raw_gold.png", "raw_ores", "mcl_raw_ores_raw_gold.png"},
		{"block", "raw_gold_block.png", "raw_ores", "mcl_raw_ores_raw_gold_block.png"},
		{"item", "raw_iron.png", "raw_ores", "mcl_raw_ores_raw_iron.png"},
		{"block", "raw_iron_block.png", "raw_ores", "mcl_raw_ores_raw_iron_block.png"},
		// mcl_sculk
		// mcl_shields
		// mcl_signs
		// mcl_smithing_table
		{"block", "smithing_table_bottom.png", "smithing_table", "mcl_smithing_table_bottom.png"},
		{"block", "smithing_table_front.png", "smithing_table", "mcl_smithing_table_front.png"},
		{"item", "empty_slot_smithing_template_armor_trim.png", "smithing_table", "mcl_smithing_table_inventory_trim_bg.png"},
		{"block", "smithing_table_side.png", "smithing_table", "mcl_smithing_table_side.png"},
		{"block", "smithing_table_top.png", "smithing_table", "mcl_smithing_table_top.png"},
		// mcl_smoker
		{"block", "smoker_bottom.png", "smoker", "smoker_bottom.png"},
		{"block", "smoker_front.png", "smoker", "smoker_front.png"},
		{"block", "smoker_front_on.png", "smoker", "smoker_front_on.png"},
		{"block", "smoker_side.png", "smoker", "smoker_side.png"},
		{"block", "smoker_top.png", "smoker", "smoker_top.png"},
		// mcl_sponges
		// mcl_spyglass
		// mcl_stairs
		// mcl_stonecutter
		// mcl_sus_nodes
		// mcl_sus_stew
		{"item", "suspicious_stew.png", "sus_stew", "sus_stew.png"},
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
		copyTexture(inName+craftPaths[e[0]]+e[1], outName+cloniaPaths[e[2]]+e[3])
	}

	//special casses
	anvil := func() {
		abase, err := imaging.Open(inName + craftPaths["block"] + "anvil.png")
		if err != nil {
			fmt.Println("AnvilBase error~", "block, anvil.png")
			return
		}
		a0, err := imaging.Open(inName + craftPaths["block"] + "anvil_top.png")
		if err != nil {
			fmt.Println("Anvil0 error~")
			return
		}
		a1, err := imaging.Open(inName + craftPaths["block"] + "chipped_anvil_top.png")
		if err != nil {
			fmt.Println("Anvil1 error~")
			return
		}
		a2, err := imaging.Open(inName + craftPaths["block"] + "damaged_anvil_top.png")
		if err != nil {
			fmt.Println("Anvil2 error~")
			return
		}
		anvilX := abase.Bounds().Dx()
		anvilY := abase.Bounds().Dy()

		dst := imaging.New(anvilX, anvilY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, abase, image.Pt(0, 0))
		dst = imaging.OverlayCenter(dst, a0, 1.0)

		if err = imaging.Save(dst, outName+cloniaPaths["anvils"]+"mcl_anvils_anvil_top_damaged_0.png"); err != nil {
			fmt.Println("Anvil undamaged failed!")
		}
		dst = imaging.OverlayCenter(dst, a1, 1.0)
		if err = imaging.Save(dst, outName+cloniaPaths["anvils"]+"mcl_anvils_anvil_top_damaged_1.png"); err != nil {
			fmt.Println("Anvil damaged1 failed!")
		}
		dst = imaging.OverlayCenter(dst, a2, 1.0)
		if err = imaging.Save(dst, outName+cloniaPaths["anvils"]+"mcl_anvils_anvil_top_damaged_2.png"); err != nil {
			fmt.Println("Anvil damaged2 failed!")
		}
	}
	anvil()

	water := func() {
		/*
			craft water
			  still   :  16 x 512
			  flowing :  32 x 1024
			clonia water (and river water)
			  still   :  16 x 256
			  flowing :  16 x 1024
		*/
		wStill, err := imaging.Open(inName + craftPaths["block"] + "water_still.png")
		if err != nil {
			fmt.Println("water_still.png error~", craftPaths["block"]+"water.png")
		} else {
			wStillX := wStill.Bounds().Dx()
			wStillY := wStill.Bounds().Dy()
			dst := imaging.New(wStillX, wStillY/2, color.NRGBA{0, 0, 0, 0})
			//For still water, I use every other frame.
			for i := 0; i < wStillY/wStillX; i++ {
				a := imaging.Crop(wStill, image.Rect(0, (i*2)*wStillX, wStillX, ((i*2)+1)*wStillX))
				dst = imaging.Overlay(dst, a, image.Point{0, i * wStillX}, 1.0)
			}
			plainWater := imaging.AdjustFunc(dst,
				func(c color.NRGBA) color.NRGBA {
					r := int(c.R) - 105
					g := int(c.G) - 40
					b := int(c.B) + 20
					if r < 0 {
						r = 0
					}
					if g < 0 {
						g = 0
					}
					if b > 255 {
						b = 255
					}
					return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
				})
			if err = imaging.Save(plainWater, outName+cloniaPaths["core"]+"default_water_source_animated.png"); err != nil {
				fmt.Println("default_water_source_animated.png save failed!")
			}

			riverWater := imaging.AdjustFunc(dst,
				func(c color.NRGBA) color.NRGBA {
					r := int(c.R) - 105
					g := int(c.G) - 0
					b := int(c.B) + 45
					if r < 0 {
						r = 0
					}
					if g < 0 {
						g = 0
					}
					if b > 255 {
						b = 255
					}
					return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
				})
			if err = imaging.Save(riverWater, outName+cloniaPaths["core"]+"default_river_water_source_animated.png"); err != nil {
				fmt.Println("default_river_water_source_animated.png save failed!")
			}

		}

		wFlowing, err := imaging.Open(inName + craftPaths["block"] + "water_still.png")
		_ = wFlowing
		if err != nil {
			fmt.Println("FlowingWater error~", craftPaths["block"]+"water_still.png")
			return
		}
	}
	water()
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
