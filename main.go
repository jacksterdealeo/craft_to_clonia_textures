package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	now      = time.Now().Format("01-02-2006 15:04:05")
	nowShort = time.Now().Format("2Jan06")
)

type readWriteError struct {
	files   []string
	message string
}

func (e *readWriteError) Error() string {
	return fmt.Sprintf("%s has %d fails:\n\t%v", e.message, len(e.files), strings.Join(e.files[:], "\n\t"))
}

func main() {
	fmt.Printf("\x1b[0m"+
		`⠀⠀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⡀⠀⣤⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣄⠀⠀⠀⠀⠀⠀
⠀⣾⠋⠙⠇⢠⣤⣤⣤⠀⣤⣤⡄⢀⣀⣿⣉⠁⣤⣿⣤⡄⠀⠀⠀⠀⢠⣼⣧⣤⠀⣠⣤⣄⠀⠀⠀⠀⠀⣸⠏⠉⠷⠀⢸⡇⠀⠀⣠⣦⣤⠀⢠⣠⣤⡄⢀⣈⣏⠀⢀⣤⣤⣄⠀
⠸⣯⠀⠀⠀⠀⣿⠁⠁⢀⣵⠾⣿⠈⠉⣿⠉⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⠠⣟⡀⢹⡇⠀⠀⠀⠀⣿⠀⠀⠀⠀⢸⡇⠀⠀⣿⠀⢹⡇⢸⠉⢸⡇⠘⠉⣿⠀⠈⣥⠶⣿⠀
⠀⢿⣦⣤⡦⣤⣿⣤⡄⠸⣧⣴⣿⠀⠀⣿⡀⠀⠀⢿⣤⡄⠀⠀⠀⠀⠀⠸⣧⣤⠀⢿⣥⣾⠃⠀⠀⠀⠀⠸⣧⣤⣶⠀⠸⣧⣤⠀⢿⣤⣼⠃⢸⠀⢸⡇⢠⣴⣿⣤⠘⣧⣤⣿⡄
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣶⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀
Minecraft to Mineclonia Texture Pack Converter
	Version: %s
	Using Minecraft Pack Version: %s`, version, mc_version)
	fmt.Print("\n\n")

	if config, err := loadJsonConfig(); err != nil {
		fmt.Println(err)
	} else {
		Config = config
	}

	if !Config.DefinedInput {
		Config.InputDir = "./input/"
		if fs.ValidPath("input") {
			err := os.Mkdir("input", 0755)
			if err != nil {
				if errors.Is(err, fs.ErrPermission) {
					log.Panicf("Permission was denied. %s was not made.\n", "input")
				} else if errors.Is(err, fs.ErrExist) {
					fmt.Printf("Folder %s already exists.\n", "input")
				} else {
					fmt.Printf("How.\n")
					log.Panic(err)
				}
			} else {
				fmt.Println("Made the input folder!")
			}
		}
	}

	if !Config.DefinedOutput {
		Config.OutputDir = "./output/"
		if fs.ValidPath("output") {
			if err := os.Mkdir("output", 0755); err != nil {
				if errors.Is(err, fs.ErrPermission) {
					log.Panicf("Permission was denied. %s was not made.\n", "output")
				} else if errors.Is(err, fs.ErrExist) {
					fmt.Printf("Folder %s already exists.\n", "output")
				} else {
					fmt.Printf("How.\n")
					log.Panic(err)
				}
			} else {
				fmt.Println("Made the output folder!")
			}
		}
	}

	var dir *os.File
	var err error
	if !Config.DefinedInput {
		dir, err = os.Open("./input/")
		if err != nil {
			log.Panic("Error:", err)
			return
		}
	} else {
		dir, err = os.Open(Config.InputDir)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				log.Println(Config.InputDir + "\n\nInput folder from config doesn't exist.")
			} else {
				log.Panic("Error:", err)
				return
			}
		}
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() &&
			filepath.Ext(file.Name()) == ".zip" {
			if _, err := os.Stat(Config.InputDir + "/" + FileNameWithoutExt(file.Name())); errors.Is(err, os.ErrNotExist) {
				fmt.Println("Unzipping:", file.Name())
				if err2 := unzipSource(Config.InputDir+"/"+file.Name(), Config.InputDir+"/"+FileNameWithoutExt(file.Name())); err2 != nil {
					fmt.Println("Extraction Error:", err)
				}
			} else {
				fmt.Println(file.Name(), "was already decompressed! :D")
			}
		}
	}

	dir, err = os.Open(Config.InputDir)
	if err != nil {
		log.Panic("Error:", err)
		return
	}
	defer dir.Close()
	files, err = dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
			if Config.ExportMineclonia {
				o := fmt.Sprintf("%s_mc_to_clonia", strings.ReplaceAll(strings.ToLower(file.Name()), " ", "_"))
				convertPackClonia(file.Name(), o)
			}
			if Config.ExportMinetest_Game {
				o := fmt.Sprintf("%s_mc_to_mtg", strings.ReplaceAll(strings.ToLower(file.Name()), " ", "_"))
				convertPackMTG(file.Name(), o)
			}

			fmt.Print("Done!\n\n")
		}
	}
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
