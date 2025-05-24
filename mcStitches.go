package main

import (
	"strings"
	"sync"

	"codeberg.org/ostech/craft_to_clonia_textures/stitches"
)

func mcStitches(input_pack_path, output_pack_path string, err_log *strings.Builder) {
	StitchFuncsToExec := stitches.EveryStitch

	errors_chan := make(chan error, len(StitchFuncsToExec))
	var wg sync.WaitGroup
	wg.Add(len(StitchFuncsToExec))
	for _, e := range StitchFuncsToExec {
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
