package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Checks if packs in input directory are just asset folders.
func assetsOnlyChecks(packsPath string) (err error) {
	inputDir, err := os.ReadDir(packsPath)
	if err != nil {
		return err
	}
	for _, texturePack := range inputDir {
		texturePackDir, err := os.ReadDir(
			filepath.Join(packsPath, texturePack.Name()))
		for _, inputFile := range texturePackDir {
			if !(inputFile.Name() == "minecraft" && inputFile.IsDir()) {
				continue
			}

			fmt.Println("You appear to have an assets-only pack in your input. Fixing that now.")
			// move the minecraftPath into the assetsPath
			minecraftPath := filepath.Join(packsPath, texturePack.Name(), "minecraft")
			assetsPath := filepath.Join(packsPath, texturePack.Name(), "assets")
			// if _, err := os.Stat(assetsPath); !errors.Is(err, os.ErrNotExist) {
			// 	continue
			// }
			if err := os.MkdirAll(assetsPath, 0750); err != nil {
				return err
			}
			if err = os.Rename(minecraftPath, filepath.Join(assetsPath, inputFile.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}
