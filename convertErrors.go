package main

import "fmt"

type ErrConvertFail int

const (
	ErrMissingRead ErrConvertFail = iota
	ErrInvalidRead
	ErrFailedSave
	ErrInvalidSaveDir
	ErrIDK
)

var (
	convertFailName = map[ErrConvertFail]string{
		ErrMissingRead:    "File was missing!",
		ErrInvalidRead:    "File couldn't be read!",
		ErrFailedSave:     "File couldn't be saved!",
		ErrInvalidSaveDir: "Directory was invalid!",
		ErrIDK:            "This shouldn't be possible.",
	}
)

func (e ErrConvertFail) Error() string {
	return fmt.Sprintln(convertFailName[e])
}
