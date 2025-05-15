package main

import (
	"strings"
	"sync"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	"codeberg.org/ostech/craft_to_clonia_textures/stitches"
	imaging "github.com/disintegration/imaging"
)

// anvil_fix(inputPackLocation+craftPaths["block"], outputPackLocation+cloniaPaths["anvils"]),
func mcStitches(input_pack_path, output_pack_path string, log *strings.Builder) {
	var wg sync.WaitGroup

	errors_c := make(chan *readWriteError, 32)
	wg.Add(1)
	go func() {
		defer wg.Done()
		stitch := "anvil textures"
		in_path := input_pack_path + data.CraftPaths["block"]
		out_path := output_pack_path + data.CloniaPaths["anvils"]
		abase, err := imaging.Open(in_path + "anvil.png")
		if err != nil {
			errors_c <- (&readWriteError{[]string{"block::anvil.png failed to open! Skipping the rest!"}, stitch})
			return
		}
		a0, err := imaging.Open(in_path + "anvil_top.png")
		if err != nil {
			errors_c <- (&readWriteError{[]string{"block::anvil_top.png failed to open! Skipping the rest!"}, stitch})
			return
		}
		a1, err := imaging.Open(in_path + "chipped_anvil_top.png")
		if err != nil {
			errors_c <- (&readWriteError{[]string{"block::chipped_anvil_top.png failed to open! Skipping the rest!"}, stitch})
			return
		}
		a2, err := imaging.Open(in_path + "damaged_anvil_top.png")
		if err != nil {
			errors_c <- (&readWriteError{[]string{"block::damaged_anvil_top.png failed to open!"}, stitch})
			return
		}

		new_a0, new_a1, new_a2 := stitches.Anvil(abase, a0, a1, a2)

		if err = imaging.Save(new_a0, out_path+"mcl_anvils_anvil_top_damaged_0.png"); err != nil {
			errors_c <- (&readWriteError{[]string{"mcl_anvils_anvil_top_damaged_0.png failed to save! Skipping the rest!"}, stitch})
			return
		}
		if err = imaging.Save(new_a1, out_path+"mcl_anvils_anvil_top_damaged_1.png"); err != nil {
			errors_c <- (&readWriteError{[]string{"mcl_anvils_anvil_top_damaged_1.png failed to save! Skipping the rest!"}, stitch})
			return
		}
		if err = imaging.Save(new_a2, out_path+"mcl_anvils_anvil_top_damaged_2.png"); err != nil {
			errors_c <- (&readWriteError{[]string{"mcl_anvils_anvil_top_damaged_2.png failed to save!"}, stitch})
			return
		}
	}()

	// errors_c <- (&readWriteError{[]string{"test.png failed to save!"}, "discovery"})

	wg.Wait()
	close(errors_c)
	for e := range errors_c {
		log.WriteString(e.Error())
	}
}
