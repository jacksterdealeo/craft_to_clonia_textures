package data

var (
	// Copper Stuff
	CopperStuffWithFallback = [...]ConversionWithFallback{
		// copper_horse_armor.png
		{[2]SimpleTexture{{"item", "copper_boots.png"}, {"item", "iron_boots.png"}}, "copper_stuff", "mcl_copper_stuff_inv_boots_copper.png", 0},
		{[2]SimpleTexture{{"item", "copper_chestplate.png"}, {"item", "iron_chestplate.png"}}, "copper_stuff", "mcl_copper_stuff_inv_chestplate_copper.png", 0},
		{[2]SimpleTexture{{"item", "copper_helmet.png"}, {"item", "iron_helmet.png"}}, "copper_stuff", "mcl_copper_stuff_inv_helmet_copper.png", 0},
		{[2]SimpleTexture{{"item", "copper_leggings.png"}, {"item", "iron_leggings.png"}}, "copper_stuff", "mcl_copper_stuff_inv_leggings_copper.png", 0},

		{[2]SimpleTexture{{"item", "copper_hoe.png"}, {"item", "iron_hoe.png"}}, "copper_stuff", "mcl_copper_stuff_copper_hoe.png", 0},
		{[2]SimpleTexture{{"item", "copper_axe.png"}, {"item", "iron_axe.png"}}, "copper_stuff", "mcl_copper_stuff_copper_axe.png", 0},
		{[2]SimpleTexture{{"item", "copper_pickaxe.png"}, {"item", "iron_pickaxe.png"}}, "copper_stuff", "mcl_copper_stuff_copper_pickaxe.png", 0},
		{[2]SimpleTexture{{"item", "copper_shovel.png"}, {"item", "iron_shovel.png"}}, "copper_stuff", "mcl_copper_stuff_copper_shovel.png", 0},
		{[2]SimpleTexture{{"item", "copper_sword.png"}, {"item", "iron_sword.png"}}, "copper_stuff", "mcl_copper_stuff_copper_sword.png", 0},

		{[2]SimpleTexture{{"item", "copper_shears_not_real_object_to_please_the_function_params.png"}, {"item", "shears.png"}}, "copper_stuff", "mcl_copper_stuff_copper_shears.png", 0},
		{[2]SimpleTexture{{"item", "copper_nugget.png"}, {"item", "iron_nugget.png"}}, "copper_stuff", "mcl_copper_stuff_copper_nugget.png", 0},
		{[2]SimpleTexture{{"block", "copper_bars.png"}, {"block", "iron_bars.png"}}, "copper_stuff", "xpanes_pane_copper.png", 0},
	}

	// Emerald Stuff
	EmeraldStuffMod = [...]SimpleConversion{
		{"item", "diamond_boots.png", "emerald_stuff", "mcl_emerald_stuff_inv_boots_emerald.png", 1},
		{"item", "diamond_chestplate.png", "emerald_stuff", "mcl_emerald_stuff_inv_chestplate_emerald.png", 1},
		{"item", "diamond_helmet.png", "emerald_stuff", "mcl_emerald_stuff_inv_helmet_emerald.png", 1},
		{"item", "diamond_leggings.png", "emerald_stuff", "mcl_emerald_stuff_inv_leggings_emerald.png", 1},

		{"item", "diamond_axe.png", "emerald_stuff", "mcl_emerald_stuff_axe.png", 1},
		{"item", "diamond_pickaxe.png", "emerald_stuff", "mcl_emerald_stuff_pick.png", 1},
		{"item", "diamond_shovel.png", "emerald_stuff", "mcl_emerald_stuff_shovel.png", 1},
		{"item", "diamond_sword.png", "emerald_stuff", "mcl_emerald_stuff_sword.png", 1},
		{"item", "diamond_hoe.png", "emerald_stuff", "mcl_emerald_stuff_hoe.png", 1},
		{"item", "diamond_horse_armor.png", "emerald_stuff", "mcl_emerald_stuff_emerald_horse_armor.png", 1},
	}
	EmeraldStuffModShortSpear = SimpleConversion{"item", "diamond_spear.png", "vl", "mcl_emerald_stuff_spear.png", 1}
	EmeraldStuffModLongSpear  = SimpleConversion{"item", "diamond_spear_in_hand.png", "vl", "mcl_emerald_stuff_spear.png", 1}

	// Rose Gold Stuff
	RoseGoldStuffMod_NetheriteToRoseGold = [...]SimpleConversion{
		{"item", "netherite_boots.png", "rose_gold_stuff", "mcl_rose_gold_inv_boots_rose_gold.png", 1},
		{"item", "netherite_chestplate.png", "rose_gold_stuff", "mcl_rose_gold_inv_chestplate_rose_gold.png", 1},
		{"item", "netherite_helmet.png", "rose_gold_stuff", "mcl_rose_gold_inv_helmet_rose_gold.png", 1},
		{"item", "netherite_leggings.png", "rose_gold_stuff", "mcl_rose_gold_inv_leggings_rose_gold.png", 1},
	}
	RoseGoldStuffMod_CopperToRoseGoldExposed = [...]SimpleConversion{
		{"block", "raw_copper_block.png", "rose_gold_stuff", "mcl_rose_gold_raw_rose_gold_ore_block_exposed.png", 1},
		{"block", "oxidized_copper.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_block_exposed.png", 1},
	}
	RoseGoldStuffMod_CopperToRoseGold = [...]SimpleConversion{
		{"block", "raw_copper_block.png", "rose_gold_stuff", "mcl_rose_gold_raw_rose_gold_ore_block.png", 1},
		{"block", "copper_block.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_block.png", 1},
		{"item", "raw_copper.png", "rose_gold_stuff", "mcl_rose_gold_raw_rose_gold_ore.png", 1},
		{"item", "copper_ingot.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_ingot.png", 1},
	}
	RoseGoldStuffMod_IronToRoseGold = [...]SimpleConversion{
		{"item", "shears.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_shears.png", 1},
		/* ARMOR
		{"item", "iron_boots.png", "rose_gold_stuff", "mcl_rose_gold_inv_boots_rose_gold.png", 1},
		{"item", "iron_chestplate.png", "rose_gold_stuff", "mcl_rose_gold_inv_chestplate_rose_gold.png", 1},
		{"item", "iron_helmet.png", "rose_gold_stuff", "mcl_rose_gold_inv_helmet_rose_gold.png", 1},
		{"item", "iron_leggings.png", "rose_gold_stuff", "mcl_rose_gold_inv_leggings_rose_gold.png", 1},
		*/
		{"item", "iron_hoe.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_hoe.png", 1},
		{"item", "iron_axe.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_axe.png", 1},
		{"item", "iron_pickaxe.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_pick.png", 1},
		{"item", "iron_shovel.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_shovel.png", 1},
		{"item", "iron_sword.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_sword.png", 1},
		{"block", "lantern.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_lantern.png", 1},
		{"item", "lantern.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_lantern_inv.png", 1},
	}
	RoseGoldStuffMod_IronToRoseGoldNoFilter = [...]SimpleConversion{
		{"item", "iron_nugget.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_nugget.png", 1},
		{"block", "iron_chain.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_chain.png", 1},
		{"item", "iron_chain.png", "rose_gold_stuff", "mcl_rose_gold_rose_gold_chain_inv.png", 1},
	}
)

var (
	Travelnet = [...]SimpleConversion{
		{"block", ".png", "travelnet", "travelnet_bottom.png", 1},
		{"block", ".png", "travelnet", "travelnet_elevator_door_glass.png", 1},
		{"block", ".png", "travelnet", "travelnet_elevator_front.png", 1},
		{"block", ".png", "travelnet", "travelnet_elevator_inside_ceiling.png", 1},
		{"block", ".png", "travelnet", "travelnet_elevator_inside_controls.png", 1},
		{"block", ".png", "travelnet", "travelnet_elevator_inside_floor.png", 1},
		{"block", ".png", "travelnet", "travelnet_elevator_inv.png", 1},
		{"block", ".png", "travelnet", "travelnet_elevator_sides_outside.png", 1},
		{"block", ".png", "travelnet", "travelnet_flash.png", 1},
		{"block", ".png", "travelnet", "travelnet_inv_base.png", 1},
		{"block", ".png", "travelnet", "travelnet_inv_colorable.png", 1},
		{"block", ".png", "travelnet", "travelnet_top.png", 1},
		{"block", ".png", "travelnet", "travelnet_travelnet_back.png", 1},
		{"block", ".png", "travelnet", "travelnet_travelnet_back_color.png", 1},
		{"block", ".png", "travelnet", "travelnet_travelnet_front.png", 1},
		{"block", ".png", "travelnet", "travelnet_travelnet_front_color.png", 1},
		{"block", ".png", "travelnet", "travelnet_travelnet_side.png", 1},
		{"block", ".png", "travelnet", "travelnet_travelnet_side_color.png", 1},
	}
)
