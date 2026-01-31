package data

var SimpleMobs = [...]StaticTexture{
	// {"entity", "allay/allay.png", "mobs_mc", ".png"}, //no match
	{"entity", "axolotl/axolotl_blue.png", "mobs_mc", "mobs_mc_axolotl_purple.png"},
	{"entity", "axolotl/axolotl_cyan.png", "mobs_mc", "mobs_mc_axolotl_white.png"},
	{"entity", "axolotl/axolotl_gold.png", "mobs_mc", "mobs_mc_axolotl_yellow.png"},
	{"entity", "axolotl/axolotl_lucy.png", "mobs_mc", "mobs_mc_axolotl_pink.png"},
	{"entity", "axolotl/axolotl_wild.png", "mobs_mc", "mobs_mc_axolotl_brown.png"},
	// {"entity", "axolotl/.png", "mobs_mc", "mobs_mc_axolotl_black.png"},  // no match
	// {"entity", "axolotl/.png", "mobs_mc", "mobs_mc_axolotl_green.png"},  // no match
	{"entity", "bear/polarbear.png", "mobs_mc", "mobs_mc_polarbear.png"},

	{"entity", "cat/all_black.png", "mobs_mc", "mobs_mc_cat_all_black.png"},
	{"entity", "cat/black.png", "mobs_mc", "mobs_mc_cat_black.png"},
	{"entity", "cat/british_shorthair.png", "mobs_mc", "mobs_mc_cat_british_shorthair.png"},
	{"entity", "cat/calico.png", "mobs_mc", "mobs_mc_cat_calico.png"},
	{"entity", "cat/cat_collar.png", "mobs_mc", "mobs_mc_cat_collar.png"},
	{"entity", "cat/jellie.png", "mobs_mc", "mobs_mc_cat_jellie.png"},
	{"entity", "cat/ocelot.png", "mobs_mc", "mobs_mc_cat_ocelot.png"},
	{"entity", "cat/persian.png", "mobs_mc", "mobs_mc_cat_persian.png"},
	{"entity", "cat/ragdoll.png", "mobs_mc", "mobs_mc_cat_ragdoll.png"},
	{"entity", "cat/red.png", "mobs_mc", "mobs_mc_cat_red.png"},
	{"entity", "cat/siamese.png", "mobs_mc", "mobs_mc_cat_siamese.png"},
	{"entity", "cat/tabby.png", "mobs_mc", "mobs_mc_cat_tabby.png"},
	{"entity", "cat/white.png", "mobs_mc", "mobs_mc_cat_white.png"},

	// {"entity", "cow/brown_mooshroom.png", "mobs_mc", ".png"},  // no match
	// {"entity", "cow/cold_cow.png", "mobs_mc", ".png"},  // no match
	// {"entity", "cow/red_mooshroom.png", "mobs_mc", ".png"},  // no match
	// {"entity", "cow/temperate_cow.png", "mobs_mc", "mobs_mc_cow.png"},  // special attention
	// {"entity", "cow/warm_cow.png", "mobs_mc", ".png"},  // no match
	// {"entity", "cow/.png", "mobs_mc", ".png"},  // no match

	{"entity", "creeper/creeper.png", "mobs_mc", "mobs_mc_creeper.png"},
	{"entity", "creeper/creeper_armor.png", "mobs_mc", "mobs_mc_creeper_charge.png"},

	// {"entity", "chicken/cold_chicken.png", "mobs_mc", ".png"},  // no match
	{"entity", "chicken/temperate_chicken.png", "mobs_mc", "mobs_mc_chicken.png"},
	// {"entity", "chicken/warm_chicken.png", "mobs_mc", ".png"},  // no match

	// {"entity", "enderdragon/dragon.png", "mobs_mc", "mobs_mc_dragon.png"},  // special attention
	{"entity", "enderdragon/dragon_fireball.png", "mobs_mc", "mobs_mc_dragon_fireball.png"},
	{"entity", "enderman/enderman.png", "mobs_mc", "mobs_mc_enderman.png"},
	// {"entity", ".png", "mobs_mc", "mobs_mc_enderman_eyes.png"},  // no match

	{"entity", "fish/cod.png", "mobs_mc", "extra_mobs_cod.png"},
	{"entity", "fish/pufferfish.png", "mobs_mc", ".png"},
	{"entity", "fish/salmon.png", "mobs_mc", "extra_mobs_salmon.png"},
	{"entity", "fish/tropical_a.png", "mobs_mc", "extra_mobs_tropical_fish_a.png"},
	{"entity", "fish/tropical_a_pattern_1.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_a_1.png"},
	{"entity", "fish/tropical_a_pattern_2.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_a_2.png"},
	{"entity", "fish/tropical_a_pattern_3.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_a_3.png"},
	{"entity", "fish/tropical_a_pattern_4.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_a_4.png"},
	{"entity", "fish/tropical_a_pattern_5.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_a_5.png"},
	{"entity", "fish/tropical_a_pattern_6.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_a_6.png"},
	{"entity", "fish/tropical_b.png", "mobs_mc", "extra_mobs_tropical_fish_b.png"},
	{"entity", "fish/tropical_b_pattern_1.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_b_1.png"},
	{"entity", "fish/tropical_b_pattern_2.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_b_2.png"},
	{"entity", "fish/tropical_b_pattern_3.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_b_3.png"},
	{"entity", "fish/tropical_b_pattern_4.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_b_4.png"},
	{"entity", "fish/tropical_b_pattern_5.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_b_5.png"},
	{"entity", "fish/tropical_b_pattern_6.png", "mobs_mc", "extra_mobs_tropical_fish_pattern_b_6.png"},

	{"entity", "ghast/ghast.png", "mobs_mc", "mobs_mc_ghast.png"},
	{"entity", "ghast/ghast_shooting.png", "mobs_mc", "mobs_mc_ghast_firing.png"},

	{"entity", "hoglin/hoglin.png", "mobs_mc", "extra_mobs_hoglin.png"},
	{"entity", "hoglin/zoglin.png", "mobs_mc", "extra_mobs_zoglin.png"},

	/* special attention The Clonia versions have saddles on all of them. They may not support higher-rez textures.
	{"entity", "horse/donkey.png", "mobs_mc", "mobs_mc_donkey.png"},
	{"entity", "horse/horse_black.png", "mobs_mc", "mobs_mc_horse_black.png"},
	{"entity", "horse/horse_brown.png", "mobs_mc", "mobs_mc_horse_brown.png"},
	{"entity", "horse/horse_chestnut.png", "mobs_mc", "mobs_mc_horse_chestnut.png"},
	{"entity", "horse/horse_creamy.png", "mobs_mc", "mobs_mc_horse_creamy.png"},
	{"entity", "horse/horse_darkbrown.png", "mobs_mc", "mobs_mc_horse_darkbrown.png"},
	{"entity", "horse/horse_gray.png", "mobs_mc", "mobs_mc_horse_gray.png"},
	{"entity", "horse/horse_markings_blackdots.png", "mobs_mc", "mobs_mc_horse_markings_blackdots.png"},
	{"entity", "horse/horse_markings_white.png", "mobs_mc", "mobs_mc_horse_markings_white.png"},
	{"entity", "horse/horse_markings_whitedots.png", "mobs_mc", "mobs_mc_horse_markings_whitedots.png"},
	{"entity", "horse/horse_markings_whitefield.png", "mobs_mc", "mobs_mc_horse_markings_whitefield.png"},
	{"entity", "horse/horse_skeleton.png", "mobs_mc", "mobs_mc_horse_skeleton.png"},
	{"entity", "horse/horse_white.png", "mobs_mc", "mobs_mc_horse_white.png"},
	{"entity", "horse/horse_zombie.png", "mobs_mc", "mobs_mc_horse_zombie.png"},
	{"entity", "horse/mule.png", "mobs_mc", "mobs_mc_mule.png"},
	*/

	{"entity", "illager/evoker.png", "mobs_mc", "mobs_mc_evoker.png"},
	{"entity", "illager/evoker_fangs.png", "mobs_mc", "mobs_mc_evoker_fangs.png"},
	{"entity", "illager/illusioner.png", "mobs_mc", "mobs_mc_illusionist.png"},
	{"entity", "illager/pillager.png", "mobs_mc", "mobs_mc_pillager.png"},
	// {"entity", "illager/ravager.png", "mobs_mc", ".png"},  // no match
	// {"entity", "illager/vex.png", "mobs_mc", "mobs_mc_vex.png"},  // special attention
	// {"entity", "illager/vex_charging.png", "mobs_mc", "mobs_mc_vex_charging.png"},  // special attention
	{"entity", "illager/vindicator.png", "mobs_mc", "mobs_mc_vindicator.png"},

	{"entity", "iron_golem/iron_golem.png", "mobs_mc", "mobs_mc_iron_golem.png"},
	{"entity", "iron_golem/iron_golem_crackiness_high.png", "mobs_mc", "mobs_mc_iron_golem_crack_high.png"},
	{"entity", "iron_golem/iron_golem_crackiness_low.png", "mobs_mc", "mobs_mc_iron_golem_crack_low.png"},
	{"entity", "iron_golem/iron_golem_crackiness_medium.png", "mobs_mc", "mobs_mc_iron_golem_crack_medium.png"},

	{"entity", "llama/brown.png", "mobs_mc", "mobs_mc_llama_brown.png"},
	{"entity", "llama/creamy.png", "mobs_mc", "mobs_mc_llama_creamy.png"},
	{"entity", "llama/creamy.png", "mobs_mc", "mobs_mc_llama.png"},
	{"entity", "llama/gray.png", "mobs_mc", "mobs_mc_llama_gray.png"},
	{"entity", "llama/white.png", "mobs_mc", "mobs_mc_llama_white.png"},
	{"entity", "llama/white.png", "mobs_mc", "mobs_mc_llama_white.png"},
	{"llama_body", "black.png", "mobs_mc", "mobs_mc_llama_decor_black.png"},
	{"llama_body", "blue.png", "mobs_mc", "mobs_mc_llama_decor_blue.png"},
	{"llama_body", "brown.png", "mobs_mc", "mobs_mc_llama_decor_brown.png"},
	{"llama_body", "cyan.png", "mobs_mc", "mobs_mc_llama_decor_cyan.png"},
	{"llama_body", "gray.png", "mobs_mc", "mobs_mc_llama_decor_gray.png"},
	{"llama_body", "green.png", "mobs_mc", "mobs_mc_llama_decor_green.png"},
	{"llama_body", "light_blue.png", "mobs_mc", "mobs_mc_llama_decor_light_blue.png"},
	{"llama_body", "light_gray.png", "mobs_mc", "mobs_mc_llama_decor_light_gray.png"},
	{"llama_body", "lime.png", "mobs_mc", "mobs_mc_llama_decor_lime.png"},
	{"llama_body", "magenta.png", "mobs_mc", "mobs_mc_llama_decor_magenta.png"},
	{"llama_body", "orange.png", "mobs_mc", "mobs_mc_llama_decor_orange.png"},
	{"llama_body", "pink.png", "mobs_mc", "mobs_mc_llama_decor_pink.png"},
	{"llama_body", "purple.png", "mobs_mc", "mobs_mc_llama_decor_purple.png"},
	{"llama_body", "red.png", "mobs_mc", "mobs_mc_llama_decor_red.png"},
	{"llama_body", "trader_llama.png", "mobs_mc", "mobs_mc_llama_decor_wandering_trader.png"},
	{"llama_body", "white.png", "mobs_mc", "mobs_mc_llama_decor_white.png"},
	{"llama_body", "yellow.png", "mobs_mc", "mobs_mc_llama_decor_yellow.png"},

	{"entity", "parrot/parrot_blue.png", "mobs_mc", "mobs_mc_parrot_blue.png"},
	{"entity", "parrot/parrot_green.png", "mobs_mc", "mobs_mc_parrot_green.png"},
	{"entity", "parrot/parrot_grey.png", "mobs_mc", "mobs_mc_parrot_grey.png"},
	{"entity", "parrot/parrot_red_blue.png", "mobs_mc", "mobs_mc_parrot_red_blue.png"},
	{"entity", "parrot/parrot_yellow_blue.png", "mobs_mc", "mobs_mc_parrot_yellow_blue.png"},

	// {"entity", "pig/cold_pig.png", "mobs_mc", ".png"}, no match
	{"entity", "pig/pig_saddle.png", "mobs_mc", "mobs_mc_pig_saddle.png"},
	// {"entity", "pig/temperate_pig.png", "mobs_mc", "mobs_mc_pig.png"},  // special attention
	// {"entity", "pig/warm_pig.png", "mobs_mc", ".png"},  // no match

	{"entity", "piglin/piglin.png", "mobs_mc", "extra_mobs_piglin.png"},
	{"entity", "piglin/piglin_brute.png", "mobs_mc", "extra_mobs_piglin_brute.png"},
	{"entity", "piglin/zombified_piglin.png", "mobs_mc", "extra_mobs_zombified_piglin.png"},

	// {"entity", "player/slim/alex.png", "other", "mcl_skins_character_1.png"},  // special attention
	// {"entity", "player/wide/steve.png", "other", "character.png"},  // special attention

	{"entity", "rabbit/black.png", "mobs_mc", "mobs_mc_rabbit_black.png"},
	{"entity", "rabbit/brown.png", "mobs_mc", "mobs_mc_rabbit_brown.png"},
	{"entity", "rabbit/caerbannog.png", "mobs_mc", "mobs_mc_rabbit_caerbannog.png"},
	{"entity", "rabbit/gold.png", "mobs_mc", "mobs_mc_rabbit_gold.png"},
	{"entity", "rabbit/salt.png", "mobs_mc", "mobs_mc_rabbit_salt.png"},
	{"entity", "rabbit/toast.png", "mobs_mc", "mobs_mc_rabbit_toast.png"},
	{"entity", "rabbit/white.png", "mobs_mc", "mobs_mc_rabbit_white.png"},
	{"entity", "rabbit/white_splotched.png", "mobs_mc", "mobs_mc_rabbit_white_splotched.png"},

	{"entity", "sheep/sheep.png", "mobs_mc", "mobs_mc_sheep.png"},
	{"entity", "sheep/sheep_wool.png", "mobs_mc", "mobs_mc_sheep_fur.png"},
	{"entity", "sheep/sheep_wool_undercoat.png", "mobs_mc", "mobs_mc_sheep_sheared.png"},

	// {"entity", "skeleton/bogged.png", "mobs_mc", ".png"},  // no match
	// {"entity", "skeleton/bogged_overlay.png", "mobs_mc", ".png"},  // no match
	{"entity", "skeleton/skeleton.png", "mobs_mc", "mobs_mc_skeleton.png"},
	{"entity", "skeleton/stray.png", "mobs_mc", "mobs_mc_stray.png"},
	{"entity", "skeleton/stray_overlay.png", "mobs_mc", "mobs_mc_stray_overlay.png"},
	{"entity", "skeleton/wither_skeleton.png", "mobs_mc", "mobs_mc_wither_skeleton.png"},

	// {"entity", "slime/magmacube.png", "mobs_mc", "mobs_mc_magmacube.png"}, // special attention
	// {"entity", "slime/slime.png", "mobs_mc", "mobs_mc_slime.png"}, // special attention

	{"entity", "spider/cave_spider.png", "mobs_mc", "mobs_mc_cave_spider.png"},
	{"entity", "spider/spider.png", "mobs_mc", "mobs_mc_spider.png"},

	{"entity", "squid/glow_squid.png", "mobs_mc", "extra_mobs_glow_squid.png"},
	{"entity", "squid/squid.png", "mobs_mc", "mobs_mc_squid.png"},

	{"entity", "strider/strider.png", "mobs_mc", "extra_mobs_strider.png"},
	{"entity", "strider/strider_cold.png", "mobs_mc", "extra_mobs_strider_cold.png"},

	// Villagers are weird.
	// TODO: Look into their mcmeta file.
	{"villager", "villager.png", "mobs_mc", "mobs_mc_villager.png"},
	{"villager", "villager.png", "mobs_mc", "mobs_mc_villager_base.png"},
	{"villager", "profession/armorer.png", "mobs_mc", "mobs_mc_villager_profession_armorer.png"},
	{"villager", "profession/butcher.png", "mobs_mc", "mobs_mc_villager_profession_butcher.png"},
	{"villager", "profession/cartographer.png", "mobs_mc", "mobs_mc_villager_profession_cartographer.png"},
	{"villager", "profession/cleric.png", "mobs_mc", "mobs_mc_villager_profession_cleric.png"},
	{"villager", "profession/farmer.png", "mobs_mc", "mobs_mc_villager_profession_farmer.png"},
	{"villager", "profession/fisherman.png", "mobs_mc", "mobs_mc_villager_profession_fisherman.png"},
	{"villager", "profession/fletcher.png", "mobs_mc", "mobs_mc_villager_profession_fletcher.png"},
	{"villager", "profession/leatherworker.png", "mobs_mc", "mobs_mc_villager_profession_leatherworker.png"},
	{"villager", "profession/librarian.png", "mobs_mc", "mobs_mc_villager_profession_librarian.png"},
	{"villager", "profession/mason.png", "mobs_mc", "mobs_mc_villager_profession_mason.png"},
	{"villager", "profession/nitwit.png", "mobs_mc", "mobs_mc_villager_profession_nitwit.png"},
	{"villager", "profession/shepherd.png", "mobs_mc", "mobs_mc_villager_profession_shepherd.png"},
	{"villager", "profession/toolsmith.png", "mobs_mc", "mobs_mc_villager_profession_toolsmith.png"},
	{"villager", "profession/weaponsmith.png", "mobs_mc", "mobs_mc_villager_profession_weaponsmith.png"},
	{"villager", "type/desert.png", "mobs_mc", "mobs_mc_villager_desert.png"},
	{"villager", "type/jungle.png", "mobs_mc", "mobs_mc_villager_jungle.png"},
	{"villager", "type/plains.png", "mobs_mc", "mobs_mc_villager_plains.png"},
	{"villager", "type/savanna.png", "mobs_mc", "mobs_mc_villager_savanna.png"},
	{"villager", "type/snow.png", "mobs_mc", "mobs_mc_villager_snow.png"},
	{"villager", "type/swamp.png", "mobs_mc", "mobs_mc_villager_swamp.png"},
	{"villager", "type/taiga.png", "mobs_mc", "mobs_mc_villager_taiga.png"},

	{"entity", "wither/wither.png", "mobs_mc", "mobs_mc_wither.png"},
	{"entity", "wither/wither_armor.png", "mobs_mc", "mobs_mc_wither_armor.png"},
	{"entity", "wither/wither_invulnerable.png", "mobs_mc", "mobs_mc_wither_invulnerable.png"},

	{"entity", "wolf/wolf.png", "mobs_mc", "mobs_mc_wolf.png"},
	{"entity", "wolf/wolf_angry.png", "mobs_mc", "mobs_mc_wolf_angry.png"},
	{"entity", "wolf/wolf_angry.png", "mobs_mc", "mobs_mc_wolf_angry_eyes.png"},
	// {"entity", "wolf/wolf_armor_crackiness_high.png", "mobs_mc", ".png"},  // no match
	// {"entity", "wolf/wolf_armor_crackiness_low.png", "mobs_mc", ".png"},  // no match
	// {"entity", "wolf/wolf_armor_crackiness_medium.png", "mobs_mc", ".png"},  // no match
	{"entity", "wolf/wolf_ashen.png", "mobs_mc", "mobs_mc_wolf_ashen.png"},
	{"entity", "wolf/wolf_ashen_angry.png", "mobs_mc", "mobs_mc_wolf_ashen_angry.png"},
	{"entity", "wolf/wolf_ashen_tame.png", "mobs_mc", "mobs_mc_wolf_ashen_tame.png"},
	{"entity", "wolf/wolf_black.png", "mobs_mc", "mobs_mc_wolf_black.png"},
	{"entity", "wolf/wolf_black_angry.png", "mobs_mc", "mobs_mc_wolf_black_angry.png"},
	{"entity", "wolf/wolf_black_tame.png", "mobs_mc", "mobs_mc_wolf_black_tame.png"},
	{"entity", "wolf/wolf_chestnut.png", "mobs_mc", "mobs_mc_wolf_chestnut.png"},
	{"entity", "wolf/wolf_chestnut_angry.png", "mobs_mc", "mobs_mc_wolf_chestnut_angry.png"},
	{"entity", "wolf/wolf_chestnut_tame.png", "mobs_mc", "mobs_mc_wolf_chestnut_tame.png"},
	{"entity", "wolf/wolf_collar.png", "mobs_mc", "mobs_mc_wolf_collar.png"},
	{"entity", "wolf/wolf_rusty.png", "mobs_mc", "mobs_mc_wolf_rusty.png"},
	{"entity", "wolf/wolf_rusty_angry.png", "mobs_mc", "mobs_mc_wolf_rusty_angry.png"},
	{"entity", "wolf/wolf_rusty_tame.png", "mobs_mc", "mobs_mc_wolf_rusty_tame.png"},
	{"entity", "wolf/wolf_snowy.png", "mobs_mc", "mobs_mc_wolf_snowy.png"},
	{"entity", "wolf/wolf_snowy_angry.png", "mobs_mc", "mobs_mc_wolf_snowy_angry.png"},
	{"entity", "wolf/wolf_snowy_tame.png", "mobs_mc", "mobs_mc_wolf_snowy_tame.png"},
	{"entity", "wolf/wolf_spotted.png", "mobs_mc", "mobs_mc_wolf_spotted.png"},
	{"entity", "wolf/wolf_spotted_angry.png", "mobs_mc", "mobs_mc_wolf_spotted_angry.png"},
	{"entity", "wolf/wolf_spotted_tame.png", "mobs_mc", "mobs_mc_wolf_spotted_tame.png"},
	{"entity", "wolf/wolf_striped.png", "mobs_mc", "mobs_mc_wolf_striped.png"},
	{"entity", "wolf/wolf_striped_angry.png", "mobs_mc", "mobs_mc_wolf_striped_angry.png"},
	{"entity", "wolf/wolf_striped_tame.png", "mobs_mc", "mobs_mc_wolf_striped_tame.png"},
	{"entity", "wolf/wolf_tame.png", "mobs_mc", "mobs_mc_wolf_striped_tame.png"},
	{"entity", "wolf/wolf_woods.png", "mobs_mc", "mobs_mc_wolf_woods.png"},
	{"entity", "wolf/wolf_woods_angry.png", "mobs_mc", "mobs_mc_wolf_woods_angry.png"},
	{"entity", "wolf/wolf_woods_tame.png", "mobs_mc", "mobs_mc_wolf_woods_tame.png"},

	// {"entity", "zombie/drowned.png", "mobs_mc", ""}, // no match
	// {"entity", "zombie/drowned_outer_layer.png", "mobs_mc", ""}, // no match
	{"entity", "zombie/husk.png", "mobs_mc", "mobs_mc_husk.png"},
	{"entity", "zombie/zombie.png", "mobs_mc", "mobs_mc_zombie.png"},

	{"zombie_villager", "zombie_villager.png", "mobs_mc", "mobs_mc_zombie_villager.png"},
	{"zombie_villager", "zombie_villager.png", "mobs_mc", "mobs_mc_zombie_villager_base.png"},
	{"zombie_villager", "profession/armorer.png", "mobs_mc", "mobs_mc_zombie_villager_profession_armorer.png"},
	{"zombie_villager", "profession/butcher.png", "mobs_mc", "mobs_mc_zombie_villager_profession_butcher.png"},
	{"zombie_villager", "profession/cartographer.png", "mobs_mc", "mobs_mc_zombie_villager_profession_cartographer.png"},
	{"zombie_villager", "profession/cleric.png", "mobs_mc", "mobs_mc_zombie_villager_profession_cleric.png"},
	{"zombie_villager", "profession/farmer.png", "mobs_mc", "mobs_mc_zombie_villager_profession_farmer.png"},
	{"zombie_villager", "profession/fisherman.png", "mobs_mc", "mobs_mc_zombie_villager_profession_fisherman.png"},
	{"zombie_villager", "profession/fletcher.png", "mobs_mc", "mobs_mc_zombie_villager_profession_fletcher.png"},
	{"zombie_villager", "profession/leatherworker.png", "mobs_mc", "mobs_mc_zombie_villager_profession_leatherworker.png"},
	{"zombie_villager", "profession/librarian.png", "mobs_mc", "mobs_mc_zombie_villager_profession_librarian.png"},
	{"zombie_villager", "profession/mason.png", "mobs_mc", "mobs_mc_zombie_villager_profession_mason.png"},
	{"zombie_villager", "profession/nitwit.png", "mobs_mc", "mobs_mc_zombie_villager_profession_nitwit.png"},
	{"zombie_villager", "profession/shepherd.png", "mobs_mc", "mobs_mc_zombie_villager_profession_shepherd.png"},
	{"zombie_villager", "profession/toolsmith.png", "mobs_mc", "mobs_mc_zombie_villager_profession_toolsmith.png"},
	{"zombie_villager", "profession/weaponsmith.png", "mobs_mc", "mobs_mc_zombie_villager_profession_weaponsmith.png"},

	{"zombie_villager", "type/desert.png", "mobs_mc", "mobs_mc_zombie_villager_desert.png"},
	{"zombie_villager", "type/jungle.png", "mobs_mc", "mobs_mc_zombie_villager_jungle.png"},
	{"zombie_villager", "type/plains.png", "mobs_mc", "mobs_mc_zombie_villager_plains.png"},
	{"zombie_villager", "type/savanna.png", "mobs_mc", "mobs_mc_zombie_villager_savanna.png"},
	{"zombie_villager", "type/snow.png", "mobs_mc", "mobs_mc_zombie_villager_snow.png"},
	{"zombie_villager", "type/swamp.png", "mobs_mc", "mobs_mc_zombie_villager_swamp.png"},
	{"zombie_villager", "type/taiga.png", "mobs_mc", "mobs_mc_zombie_villager_taiga.png"},

	// {"entity", "bat.png", "mobs_mc", "mobs_mc_bat.png"},  // special attention
	{"entity", "blaze.png", "mobs_mc", "mobs_mc_blaze.png"},
	{"entity", "dolphin.png", "mobs_mc", "extra_mobs_dolphin.png"},
	{"entity", "endermite.png", "mobs_mc", "mobs_mc_endermite.png"},
	{"entity", "guardian.png", "mobs_mc", "mobs_mc_guardian.png"},
	{"entity", "guardian_elder.png", "mobs_mc", "mobs_mc_guardian_elder.png"},
	{"entity", "silverfish.png", "mobs_mc", "mobs_mc_silverfish.png"},
	{"entity", "snow_golem.png", "mobs_mc", "mobs_mc_snowman.png"},
	{"entity", "spider_eyes.png", "mobs_mc", "mobs_mc_spider_eyes.png"},
	{"entity", "wandering_trader.png", "mobs_mc", "mobs_mc_villager_wandering_trader.png"},
	{"entity", "witch.png", "mobs_mc", "mobs_mc_witch.png"},

	// {"entity", ".png", "mobs_mc", ".png"},
}
