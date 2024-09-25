package main

import (
	"fmt"
	imaging "github.com/disintegration/imaging"
	//_ "image/png"
)

func animated_texture_fix(inName string, outName string) error {

	craftPaths := CraftPaths()
	cloniaPaths := CloniaPaths()

	animated := [][4]string{
		{"block", "respawn_anchor_top.png", "beds", "respawn_anchor_top_on.png"},
		{"block", "soul_fire_0.png", "blackstone", "soul_fire_basic_flame_animated.png"},
		{"block", "fire_0.png", "fire", "fire_basic_flame_animated.png"},
		{"block", "fire_0.png", "fire", "mcl_burning_entity_flame_animated.png"},
		{"block", "fire_0.png", "fire", "mcl_burning_hud_flame_animated.png"},
		{"block", "magma.png", "nether", "mcl_nether_magma.png"},
	}

	for _, e := range animated {
		img, err := imaging.Open(inName + craftPaths[e[0]] + e[1])
		if err != nil {
			fmt.Println(e[1], "couldn't open!")
		} else {
			if err = imaging.Save(img, outName+cloniaPaths[e[2]]+e[3]); err != nil {
				fmt.Println(e[3], "failed to save!")
			}
		}
	}
	return nil
}
