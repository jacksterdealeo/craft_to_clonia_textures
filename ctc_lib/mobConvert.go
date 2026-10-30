package ctc_lib

import "codeberg.org/ostech/craft_to_clonia_textures/ctc_lib/data"

func mobConvert(inputPackLocation string, outputPackLocation string) *readWriteError {
	copyTextureFails := []string{}
	successes := 0

	for _, texture := range data.SimpleMobs {
		err := texture.Convert(inputPackLocation, outputPackLocation)
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
