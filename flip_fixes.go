package main

import (
	"fmt"
	imaging "github.com/disintegration/imaging"
	//_ "image/png"
)

func flip_fixes(inName string, outName string) {

	craftPaths := CraftPaths()
	cloniaPaths := CloniaPaths()

	flips := [][4]string{
		////mcl_bamboo
		{"block", "bamboo_door_bottom.png", "bamboo", "mcl_bamboo_door_bottom.png"},
		{"block", "bamboo_door_top.png", "bamboo", "mcl_bamboo_door_top.png"},
		////mcl_cherry_blossom
		{"block", "cherry_door_bottom.png", "cherry_blossom", "mcl_cherry_blossom_door_bottom.png"},
		{"block", "cherry_door_top.png", "cherry_blossom", "mcl_cherry_blossom_door_top.png"},
		////mcl_crimson
		{"block", "crimson_door_bottom.png", "crimson", "mcl_crimson_crimson_door_bottom.png"},
		{"block", "crimson_door_top.png", "crimson", "mcl_crimson_crimson_door_top.png"},
		{"block", "warped_door_bottom.png", "crimson", "mcl_crimson_warped_door_bottom.png"},
		{"block", "warped_door_top.png", "crimson", "mcl_crimson_warped_door_top.png"},
		////mcl_doors
		{"block", "acacia_door_bottom.png", "doors", "mcl_doors_door_acacia_lower.png"},
		{"block", "acacia_door_top.png", "doors", "mcl_doors_door_acacia_upper.png"},
		{"block", "birch_door_bottom.png", "doors", "mcl_doors_door_birch_lower.png"},
		{"block", "birch_door_top.png", "doors", "mcl_doors_door_birch_upper.png"},
		{"block", "dark_oak_door_bottom.png", "doors", "mcl_doors_door_dark_oak_lower.png"},
		{"block", "dark_oak_door_top.png", "doors", "mcl_doors_door_dark_oak_upper.png"},
		{"block", "jungle_door_bottom.png", "doors", "mcl_doors_door_jungle_lower.png"},
		{"block", "jungle_door_top.png", "doors", "mcl_doors_door_jungle_upper.png"},
		{"block", "spruce_door_bottom.png", "doors", "mcl_doors_door_spruce_lower.png"},
		{"block", "spruce_door_top.png", "doors", "mcl_doors_door_spruce_upper.png"},
		{"block", "oak_door_bottom.png", "doors", "mcl_doors_door_wood_lower.png"},
		{"block", "oak_door_top.png", "doors", "mcl_doors_door_wood_upper.png"},
		////mcl_mangrove
		{"block", "mangrove_door_bottom.png", "mangrove", "mcl_mangrove_door_bottom.png"},
		{"block", "mangrove_door_top.png", "mangrove", "mcl_mangrove_door_top.png"},
	}

	for _, e := range flips {
		img, err := imaging.Open(inName + craftPaths[e[0]] + e[1])
		if err != nil {
			fmt.Println(e[1], "couldn't open!")
			return
		}
		img = imaging.FlipH(img)
		//anvilX := abase.Bounds().Dx()
		//anvilY := abase.Bounds().Dy()

		if err = imaging.Save(img, outName+cloniaPaths[e[2]]+e[3]); err != nil {
			fmt.Println(e[3], "failed to save!")
		}
	}
}
