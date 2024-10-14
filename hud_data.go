package main

import ()

func BasicHud() [][4]string {
	equivalents := [][4]string{
		// -- mcl_base_textures
		{"", "/assets/minecraft/textures/gui/sprites/hud/air.png",
			"", "/HUD/mcl_base_textures/textures/bubble.png"},
		//{"block", "destroy_stage_0.png", "", "crack_anylength.png"},
		//{"block", "destroy_stage_9.png", "", "crack_anylength.png"},
		{"", "assets/minecraft/textures/gui/sprites/hud/crosshair.png",
			"", "/HUD/mcl_base_textures/textures/crosshair.png"},
		{"", "assets/minecraft/textures/gui/sprites/hud/heart/full.png",
			"", "/HUD/mcl_base_textures/textures/heart.png"},
		//{"", ".png", "", "mcl_base_textures_background.png"},
		//{"", ".png", "", "mcl_base_textures_background9.png"},
		//{"", ".png", "", "mcl_base_textures_button9.png"},
		//{"", ".png", "", "mcl_base_textures_button9_pressed.png"},
		{"", "assets/minecraft/textures/gui/sprites/hud/crosshair.png",
			"", "/HUD/mcl_base_textures/textures/object_crosshair.png"},
		//{"", ".png", "", "smoke_puff.png"},

		//{"", ".png", "", ".png"},
	}
	return equivalents
}
