package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
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
