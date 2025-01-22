package main

import ()

var (
	mtgPaths = map[string]string{
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

	minetestGameItems = [...]simpleConversion{
		//{"bed", "", "beds", "beds_bed.png", 1},  // no match
		//{"bed", "", "beds", "beds_bed_fancy.png", 1},  // no match
		{"bed", "", "beds", "beds_bed_foot.png", 1},
		{"bed", "", "beds", "beds_bed_head.png", 1},
		{"bed", "", "beds", "beds_bed_side_bottom.png", 1},
		{"bed", "", "beds", "beds_bed_side_bottom_r.png", 1},
		{"bed", "", "beds", "beds_bed_side_top.png", 1},
		{"bed", "", "beds", "beds_bed_side_top_r.png", 1},
		{"bed", "", "beds", "beds_bed_side1.png", 1},
		{"bed", "", "beds", "beds_bed_side2.png", 1},
		{"bed", "", "beds", "beds_bed_top_bottom.png", 1},
		{"bed", "", "beds", "beds_bed_top_top.png", 1},
		{"bed", "", "beds", "beds_bed_top1.png", 1},
		{"bed", "", "beds", "beds_bed_top2.png", 1},
		{"bed", "", "beds", "beds_bed_under.png", 1},
	}
)
