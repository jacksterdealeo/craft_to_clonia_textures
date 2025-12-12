package data

var SimpleHUD = [...]SimpleConversion{
	// -- craftguide
	{"item", "knowledge_book.png", "craftguide", "craftguide_book.png", 1},
	// -- criticals
	{"particle", "critical_hit.png", "criticals", "mcl_particles_crit.png", 1},
	// -- hbarmor
	//{"", ".png", "hbarmor", "hbarmor_bar.png", 1},
	{"hud", "armor_empty.png", "hbarmor", "hbarmor_bgicon.png", 1},
	{"hud", "armor_full.png", "hbarmor", "hbarmor_icon.png", 1},
	// -- hudbars -- incomplete
	{"hud", "air.png", "hudbars", "hudbars_icon_breath.png", 1},
	{"hud", "air_empty.png", "hudbars", "hudbars_bgicon_breath.png", 1},
	//{"hud", "heart/full.png", "hudbars", "hudbars_icon_health.png", 1},
	{"hud", "heart/container.png", "hudbars", "hudbars_bgicon_health.png", 1},
	// -- mcl_base_textures
	{"hud", "air.png", "hud_base_textures", "bubble.png", 1},
	//{"block", "destroy_stage_0.png", "", "crack_anylength.png", 1},
	//{"block", "destroy_stage_9.png", "", "crack_anylength.png", 1},
	// {"hud", "crosshair.png", "hud_base_textures", "crosshair.png", 1}, // requires sizing
	//{"hud", "heart/full.png", "hud_base_textures", "heart.png", 1},
	//{"", ".png", "", "mcl_base_textures_background.png", 1},
	//{"", ".png", "", "mcl_base_textures_background9.png", 1},
	//{"", ".png", "", "mcl_base_textures_button9.png", 1},
	//{"", ".png", "", "mcl_base_textures_button9_pressed.png", 1},
	{"hud", "crosshair.png", "hud_base_textures", "object_crosshair.png", 1},
	//{"", ".png", "", "smoke_puff.png", 1},
	//{"", ".png", "", ".png", 1},
	// -- hunger
	{"hud", "food_empty.png", "hunger", "hbhunger_bgicon.png", 1},
	// {"hud", "food_full.png", "hunger", "hbhunger_icon.png", 1},
	// {"hud", "food_full_hunger.png", "hunger", "mcl_hunger_icon_foodpoison.png", 1},
	// -- inventory -- incomplete
	{"item", "empty_armor_slot_boots.png", "inventory", "mcl_inventory_empty_armor_slot_boots.png", 1},
	{"item", "empty_armor_slot_chestplate.png", "inventory", "mcl_inventory_empty_armor_slot_chestplate.png", 1},
	{"item", "empty_armor_slot_helmet.png", "inventory", "mcl_inventory_empty_armor_slot_helmet.png", 1},
	{"item", "empty_armor_slot_leggings.png", "inventory", "mcl_inventory_empty_armor_slot_leggings.png", 1},
	{"item", "empty_armor_slot_shield.png", "inventory", "mcl_inventory_empty_armor_slot_shield.png", 1},
	// {"hud", "hotbar.png", "inventory", "mcl_inventory_hotbar.png", -1},
	{"hud", "hotbar_selection.png", "inventory", "mcl_inventory_hotbar_selected.png", 1},
	// -- weather
	{"particle", "generic_3.png", "weather", "weather_pack_snow_snowflake1.png", 1},
	{"particle", "generic_5.png", "weather", "weather_pack_snow_snowflake2.png", 1},
	// Turns out, you have to scale the moon.
	// {"environment", "moon_phases.png", "moon", "mcl_moon_moon_phases.png", 1},
	// Rain is a bit harder to do.
	// -- mobs
	{"particle", "generic_7.png", "mobs", "mcl_particles_mob_death.png", 1},
	{"particle", "damage.png", "mobs", "mobs_blood.png", 1},

	// Requires recolor.
	// {"particle", "glint.png", "mobs_mc", "extra_mobs_glow_squid_glint1.png", 1},
	// {"particle", "glint.png", "mobs_mc", "extra_mobs_glow_squid_glint2.png", 1},
	// {"particle", "glint.png", "mobs_mc", "extra_mobs_glow_squid_glint3.png", 1},
	// {"particle", "glint.png", "mobs_mc", "extra_mobs_glow_squid_glint4.png", 1},

	// -- misc
	{"particle", "critical_hit.png", "potions", "mcl_particles_instant_effect.png", 1},
	{"particle", "effect_3.png", "potions", "mcl_particles_effect.png", 1},
}
