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
	textures := [...][2]string{
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

		//{"", f["core"] + "default_clay.png"},
		//{"", f["core"] + "default_clay_brick.png"},
		//{"", f["core"] + "default_clay_lump.png"},
		//{"", f["core"] + "default_coal_block.png"},
		//{"", f["core"] + "default_coal_lump.png"},
		//{"", f["core"] + "default_cobble.png"},

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
	_ = textures

	for _, e := range textures {
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
