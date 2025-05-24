package stitches

import "fmt"

func openErrMsg(stitch, dir, file string) error {
	return fmt.Errorf("%v Stitch > \tCould not open! ~ %v::%v", stitch, dir, file)
}

func saveErrMsg(stitch, dir, file string) error {
	return fmt.Errorf("%v Stitch > \tCould not save! ~ %v::%v", stitch, dir, file)
}
