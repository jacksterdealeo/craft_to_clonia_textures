package stitches

import (
	"log"
	"os"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
)

func RWEmbedded(input_pack_path, output_pack_path string) error {
	stitch := "RWEmbedded"
	// craft_path := "whoKnows"
	clonia_path := "vl"

	out_path := output_pack_path + data.CloniaPaths[clonia_path]

	readFails := make([][2]string, 0)
	saveFails := make([][2]string, 0)

	x16Dir := "textures/16"
	entries, err := data.TexturesFS.ReadDir(x16Dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			log.Println(entry.Name(), "is a directory! This does not happen during normal usage!")
			continue
		}
		contents, err := data.TexturesFS.ReadFile(x16Dir +"/"+ entry.Name())
		if err != nil {
			readFails = append(readFails, [2]string{x16Dir, entry.Name()})
			continue
		}
		if err := os.WriteFile(out_path+entry.Name(), contents, 0644); err != nil {
    	saveFails = append(saveFails, [2]string{out_path, entry.Name()})
    }

	}

	return multiErrMsg(stitch, readFails, saveFails)
}
