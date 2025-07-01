package main

import (
	"strings"
	"sync"

	"codeberg.org/ostech/craft_to_clonia_textures/stitches"
)

// Runs all the functions in the "stitches.EveryStitch" array.
func mcStitches(input_pack_path, output_pack_path string, err_log *strings.Builder) {
	number_of_stitches := len(stitches.EveryStitch)
	errors_chan := make(chan error, number_of_stitches)

	var wg sync.WaitGroup
	wg.Add(number_of_stitches)
	for _, e := range stitches.EveryStitch {
		go func() {
			if err := e(input_pack_path, output_pack_path); err != nil {
				errors_chan <- err
			}
			wg.Done()
		}()
	}
	wg.Wait()

	close(errors_chan)
	for e := range errors_chan {
		err_log.WriteString(e.Error())
		err_log.WriteString("\n")
	}
}
