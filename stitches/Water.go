package stitches

import (
	"image"
	"image/color"
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

// TODO: Make this read animation frames.
func RWWater(input_pack_path, output_pack_path string, _ *configure.Config) error {
	inPath := filepath.Join(input_pack_path, data.CraftPaths["block"])
	outPath := filepath.Join(output_pack_path, data.CloniaPaths["core"])
	voxPath := filepath.Join(output_pack_path, data.CloniaPaths["vl"])
	/*
		craft water
		  still   :  16 x 512
		  flowing :  32 x 1024
		clonia water (and river water)
		  still   :  16 x 256
		  flowing :  16 x 1024
	*/
	wStill, err := imaging.Open(filepath.Join(inPath, "water_still.png"))
	if err != nil {
		return openErrMsg("Water", "block", "water_still.png")
	} else {
		wStillX := wStill.Bounds().Dx()
		wStillY := wStill.Bounds().Dy()
		dst := imaging.New(wStillX, wStillY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Overlay(dst, wStill, image.Point{0, 0}, 1.0)
		if err = imaging.Save(dst, filepath.Join(voxPath, "mcl_core_water_source_animation.png")); err != nil {
			return saveErrMsg("Water", "vl", "mcl_core_water_source_animation.png")
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
		if err = imaging.Save(plainWater, filepath.Join(outPath, "default_water_source_animated.png")); err != nil {
			return saveErrMsg("Water", "core", "default_water_source_animated.png")
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
		if err = imaging.Save(riverWater, filepath.Join(outPath, "default_river_water_source_animated.png")); err != nil {
			return saveErrMsg("Water", "core", "default_river_water_source_animated.png")
		}
	}

	wFlowing, err := imaging.Open(filepath.Join(inPath, "water_flow.png"))
	if err != nil {
		return openErrMsg("Water", "block", "water_flow.png")
	} else {
		wFlowingX := wFlowing.Bounds().Dx()
		wFlowingY := wFlowing.Bounds().Dy()
		dst := imaging.New(wFlowingX/2, wFlowingY, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Overlay(dst, wFlowing, image.Point{0, 0}, 1.0)
		if err = imaging.Save(dst, filepath.Join(voxPath, "mcl_core_water_flow_animation.png")); err != nil {
			return saveErrMsg("Water", "vl", "mcl_core_water_flow_animation.png")
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
		if err = imaging.Save(plainWater, filepath.Join(outPath, "default_water_flowing_animated.png")); err != nil {
			return saveErrMsg("Water", "core", "default_water_flowing_animated.png")
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
		if err = imaging.Save(riverWater, filepath.Join(outPath, "default_river_water_flowing_animated.png")); err != nil {
			return saveErrMsg("Water", "core", "default_river_water_flowing_animated.png")
		}
	}

	waterBucket, err := imaging.Open(filepath.Join(input_pack_path, data.CraftPaths["item"], "water_bucket.png"))
	if err != nil {
		return openErrMsg("Water", "item", "water_bucket.png")
	}
	riverWaterBucket := imaging.AdjustFunc(waterBucket,
		func(c color.NRGBA) color.NRGBA {
			if !(c.B/10 > c.R/8 && c.B/10 > c.G/8) {
				return c
			}

			r := int(c.R) + 40
			g := int(c.G) + 40
			b := int(c.B) + 10
			if r > 255 {
				r = 255
			}
			if g > 255 {
				g = 255
			}
			if b > 255 {
				b = 255
			}
			return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
		})
	if err = imaging.Save(riverWaterBucket, filepath.Join(output_pack_path, data.CloniaPaths["buckets"], "bucket_river_water.png")); err != nil {
		return saveErrMsg("Water", "buckets", "bucket_river_water.png")
	}
	if err = imaging.Save(riverWaterBucket, filepath.Join(output_pack_path, data.CloniaPaths["vl"], "mcl_buckets_river_water_bucket.png")); err != nil {
		return saveErrMsg("Water", "vl", "mcl_buckets_river_water_bucket.png")
	}

	return nil
}
