package main

import "codeberg.org/ostech/craft_to_clonia_textures/data"

func MobConvert(inPath string, outPath string) *readWriteError {
	copyTextureFails := []string{}
	successes := 0

	for _, texture := range data.SimpleMobs {
		err := copyTexture(
			inPath+data.CraftPaths[texture.InPath]+texture.InTexture,
			outPath+data.CloniaPaths[texture.OutPath]+texture.OutTexture,
		)
		if err != nil {
			copyTextureFails = append(copyTextureFails, err.Error()+" ~ "+texture.InPath+"::"+texture.InTexture)
		} else {
			successes += 1
		}
	}

	if len(copyTextureFails) > 0 {
		return &readWriteError{copyTextureFails, "mob textures"}
	}
	return nil
}
