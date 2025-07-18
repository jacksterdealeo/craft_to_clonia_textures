package data

var CloniaPaths = map[string]string{
	// -- ENTITIES
	"boats":     "/ENTITIES/mcl_boats/",
	"charges":   "/ENTITIES/mcl_charges/",
	"minecarts": "/ENTITIES/mcl_minecarts/",
	"mobs_mc":   "/ENTITIES/mobs_mc/",

	// -- HUD
	"achievements": "/HUD/mcl_achievements/",
	"craftguide":   "/HELP/mcl_craftguide/",
	"criticals":    "/PLAYER/mcl_criticals/",
	"doc":          "/HELP/doc/doc/",
	"experience":   "/HUD/mcl_experience/",
	"hbarmor":      "/HUD/mcl_hbarmor/",

	"hudbars":           "/HUD/hudbars/",
	"hud_base_textures": "/HUD/mcl_base_textures/",
	"hunger":            "/PLAYER/mcl_hunger/",
	"inventory":         "/HUD/mcl_inventory/",
	"offhand":           "/HUD/mcl_offhand/",

	// -- IDK
	"other": "/other_textures/", // use when the original directory is unknown or inconvenient
	"skins": "/PLAYER/mcl_skins/",

	// -- ITEMS
	"amethyst":    "/ITEMS/mcl_amethyst/",
	"anvils":      "/ITEMS/mcl_anvils/",
	"armor":       "/ITEMS/mcl_armor/",
	"armor_stand": "/ITEMS/mcl_armor_stand/",
	"bamboo":      "/ITEMS/mcl_bamboo/",

	"banners":  "/ITEMS/mcl_banners/",
	"barrels":  "/ITEMS/mcl_barrels/",
	"beds":     "/ITEMS/mcl_beds/",
	"beehives": "/ITEMS/mcl_beehives/",
	"bells":    "/ITEMS/mcl_bells/",

	"blackstone":    "/ITEMS/mcl_blackstone/",
	"blast_furnace": "/ITEMS/mcl_blast_furnace/",
	"bone_meal":     "/ITEMS/mcl_bone_meal/",
	"books":         "/ITEMS/mcl_books/",
	"bows":          "/ITEMS/mcl_bows/",

	"brewing":           "/ITEMS/mcl_brewing/",
	"buckets":           "/ITEMS/mcl_buckets/",
	"cake":              "/ITEMS/mcl_cake/",
	"campfires":         "/ITEMS/mcl_campfires/",
	"candles":           "/ITEMS/mcl_candles/",
	"cartography_table": "/ITEMS/mcl_cartography_table/",

	"cauldrons":      "/ITEMS/mcl_cauldrons/",
	"cherry_blossom": "/ITEMS/mcl_cherry_blossom/",
	"chests":         "/ITEMS/mcl_chests/",
	"clock":          "/ITEMS/mcl_clock/",
	"cocoas":         "/ITEMS/mcl_cocoas/",

	"colorblocks": "/ITEMS/mcl_colorblocks/",
	"compass":     "/ITEMS/mcl_compass/",
	"composters":  "/ITEMS/mcl_composters/",
	"copper":      "/ITEMS/mcl_copper/",
	"core":        "/ITEMS/mcl_core/",

	"crafting_table": "/ITEMS/mcl_crafting_table/",
	"crimson":        "/ITEMS/mcl_crimson/",
	"deepslate":      "/ITEMS/mcl_deepslate/",
	"dripstone":      "/ITEMS/mcl_dripstone/",
	"doors":          "/ITEMS/mcl_doors/",

	"dyes":       "/ITEMS/mcl_dyes/",
	"enchanting": "/ITEMS/mcl_enchanting/",
	"end":        "/ITEMS/mcl_end/",
	"farming":    "/ITEMS/mcl_farming/",
	"fences":     "/ITEMS/mcl_fences/",

	"fire":            "/ITEMS/mcl_fire/",
	"fireworks":       "/ITEMS/mcl_fireworks/",
	"fishing":         "/ITEMS/mcl_fishing/",
	"fletching_table": "/ITEMS/mcl_fletching_table/",
	"flowerpots":      "/ITEMS/mcl_flowerpots/",

	"flowers":    "/ITEMS/mcl_flowers/",
	"furnaces":   "/ITEMS/mcl_furnaces/",
	"honey":      "/ITEMS/mcl_honey/",
	"hopper":     "/ITEMS/mcl_hoppers/",
	"itemframes": "/ITEMS/mcl_itemframes/",

	"jukebox":        "/ITEMS/mcl_jukebox/",
	"lanterns":       "/ITEMS/mcl_lanterns/",
	"lectern":        "/ITEMS/mcl_lectern/",
	"lightning_rods": "/ITEMS/mcl_lightning_rods/",
	"loom":           "/ITEMS/mcl_loom/",

	"lush_caves":  "/ITEMS/mcl_lush_caves/",
	"mangrove":    "/ITEMS/mcl_mangrove/",
	"maps":        "/ITEMS/mcl_maps/",
	"mobs":        "/ENTITIES/mcl_mobs/textures/",
	"mobitems":    "/ITEMS/mcl_mobitems/",
	"mobspawners": "/ITEMS/mcl_mobspawners/",

	"mud":       "/ITEMS/mcl_mud/",
	"mushrooms": "/ITEMS/mcl_mushrooms/",
	"nether":    "/ITEMS/mcl_nether/",
	"ocean":     "/ITEMS/mcl_ocean/",
	"panes":     "/ITEMS/mcl_panes/",

	"portals":        "/ITEMS/mcl_portals/",
	"potions":        "/ITEMS/mcl_potions/",
	"pottery_sherds": "/ITEMS/mcl_pottery_sherds/",
	"powder_snow":    "/ITEMS/mcl_powder_snow/",
	"raw_ores":       "/ITEMS/mcl_raw_ores/",

	"sculk":          "/ITEMS/mcl_sculk/",
	"shields":        "/ITEMS/mcl_shields/",
	"signs":          "/ITEMS/mcl_signs/",
	"sponges":        "/ITEMS/mcl_sponges/",
	"spyglass":       "/ITEMS/mcl_spyglass/",
	"smithing_table": "/ITEMS/mcl_smithing_table/",

	"smoker":      "/ITEMS/mcl_smoker/",
	"stairs":      "/ITEMS/mcl_stairs/",
	"stonecutter": "/ITEMS/mcl_stonecutter/",
	"sus_nodes":   "/ITEMS/mcl_sus_nodes/",
	"sus_stew":    "/ITEMS/mcl_sus_stew/",

	"throwing": "/ITEMS/mcl_throwing/",
	"tnt":      "/ITEMS/mcl_tnt/",
	"tools":    "/ITEMS/mcl_tools/",
	"torches":  "/ITEMS/mcl_torches/",
	"totems":   "/ITEMS/mcl_torches/",

	"vaults":  "/ITEMS/mcl_vaults/",
	"walls":   "/ITEMS/mcl_walls/",
	"wool":    "/ITEMS/mcl_wool/",
	"xfences": "/ITEMS/mclx_fences/",
	"xstairs": "/ITEMS/mclx_stairs/",

	//"": "/ITEMS//",
	"REDSTONE": "/ITEMS/REDSTONE/",

	// -- ENVIRONMENT
	"lightning":     "/ENVIRONMENT/mcl_lightning/textures/",
	"moon":          "/ENVIRONMENT/mcl_moon/textures/",
	"raids":         "/ENVIRONMENT/mcl_raids/textures/",
	"void_damage":   "/ENVIRONMENT/mcl_void_damage/textures/",
	"weather":       "/ENVIRONMENT/mcl_weather/textures/",
	"zombie_sieges": "/ENVIRONMENT/mcl_zombie_sieges/textures/",

	// mods
	"copper_stuff":    "/modded/mcl_copper_stuff/",
	"emerald_stuff":   "/modded/mcl_emerald_stuff/",
	"rose_gold_stuff": "/modded/mcl_rose_gold/",
	"travelnet":       "/modded/travelnet/",
}
