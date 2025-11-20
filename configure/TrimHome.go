package configure

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CutHomePath(path string) string {
	if cut, found := strings.CutPrefix(path, "~/"); found {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Panicf("You have no home directory? Maybe don't use \"~/\" in %v.\n %v\n", path, err)
		}
		return filepath.Join(home, cut) + "/"
	}
	return path
}
