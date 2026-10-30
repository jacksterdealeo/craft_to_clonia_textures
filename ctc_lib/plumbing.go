package ctc_lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type readWriteError struct {
	files   []string
	message string
}

func (e *readWriteError) Error() string {
	return fmt.Sprintf("%s has %d fails:\n\t%v", e.message, len(e.files), strings.Join(e.files[:], "\n\t"))
}

// Checks if packs in input directory are just asset folders.
func AssetsOnlyChecks(packsPath string) (err error) {
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
