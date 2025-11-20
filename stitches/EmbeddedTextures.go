package stitches

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

// TODO: ACTUALLY FIX ERROR HANDLING! ADD AN ISSUE ON THE CODEBERG PAGE!

func RWEmbedded(input_pack_path, output_pack_path string, config *configure.Config) error {
	stitch := "RWEmbedded"
	// craft_path := "whoKnows"
	clonia_path := "embedded"

	out_path := output_pack_path + data.CloniaPaths[clonia_path]

	fails := make([]error, 0)

	x16Dir := "textures/16"
	entries, err := data.TexturesFS.ReadDir(x16Dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		contents, err := data.TexturesFS.ReadFile(x16Dir + "/" + entry.Name())
		if err != nil {
			fails = append(fails, fmt.Errorf("%v, %v, %w", x16Dir, entry.Name(), err))
			continue
		}
		if err := os.WriteFile(out_path+entry.Name(), contents, 0644); err != nil {
			fails = append(fails, fmt.Errorf("%v, %v, %w", out_path, entry.Name(), err))
		}
	}

	switch config.SpearVersion {
	case "old", "embedded", "deprecated":
		x16SpearDir := "textures/16/deprecated_spears"
		entries, err := data.TexturesFS.ReadDir(x16SpearDir)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			contents, err := data.TexturesFS.ReadFile(x16SpearDir + "/" + entry.Name())
			if err != nil {
				fails = append(fails, fmt.Errorf("%v, %v, %w", x16Dir, entry.Name(), err))
				continue
			}
			if err := os.WriteFile(out_path+entry.Name(), contents, 0644); err != nil {
				fails = append(fails, fmt.Errorf("%v, %v, %w", out_path, entry.Name(), err))
			}
		}

	case "long":
		// I am going to ignore animations for this.
		for _, e := range data.LongSpears {

			img, err := imaging.Open(filepath.Join(input_pack_path, e.ReadPath()))
			if err != nil {
				fails = append(fails, fmt.Errorf("%v, %v, %w", "shortSpears", e.OutTexture, err))
			}
			imgX := img.Bounds().Dx()
			// imgY := img.Bounds().Dy()

			outImg := imaging.New(imgX, imgX, color.NRGBA{0, 0, 0, 0})
			outImg = imaging.Rotate270(imaging.Overlay(outImg, img, image.Point{0, 0}, 1.0))

			if err = imaging.Save(outImg, filepath.Join(output_pack_path, e.SavePath())); err != nil {
				fails = append(fails, fmt.Errorf("%v, %v, %w", "shortSpears", e.OutTexture, err))
			}
		}

	default:
		for _, e := range data.ShortSpears {
			err := CopyTextureAnimated(filepath.Join(input_pack_path, e.ReadPath()),
				filepath.Join(output_pack_path, e.SavePath()), e.FramesAllowed)
			if err != nil {
				fails = append(fails, fmt.Errorf("%v, %v, %w", "shortSpears", e.OutTexture, err))
			}
		}

	}

	return fmt.Errorf("%v %v", stitch, fails)
}
