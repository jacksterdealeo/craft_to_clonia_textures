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

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
)

var (
	nowShort       = time.Now().Format("2Jan06")
	ConfigLocation = "config.json"
)

type readWriteError struct {
	files   []string
	message string
}

func (e *readWriteError) Error() string {
	return fmt.Sprintf("%s has %d fails:\n\t%v", e.message, len(e.files), strings.Join(e.files[:], "\n\t"))
}

func main() {
	var (
		config *configure.Config
		err    error
	)

	if version == "scriptless" {
		version += " " + nowShort
	}

	fmt.Printf("\x1b[0m"+
		`⠀⠀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⡀⠀⣤⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣄⠀⠀⠀⠀⠀⠀
⠀⣾⠋⠙⠇⢠⣤⣤⣤⠀⣤⣤⡄⢀⣀⣿⣉⠁⣤⣿⣤⡄⠀⠀⠀⠀⢠⣼⣧⣤⠀⣠⣤⣄⠀⠀⠀⠀⠀⣸⠏⠉⠷⠀⢸⡇⠀⠀⣠⣦⣤⠀⢠⣠⣤⡄⢀⣈⣏⠀⢀⣤⣤⣄⠀
⠸⣯⠀⠀⠀⠀⣿⠁⠁⢀⣵⠾⣿⠈⠉⣿⠉⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⠠⣟⡀⢹⡇⠀⠀⠀⠀⣿⠀⠀⠀⠀⢸⡇⠀⠀⣿⠀⢹⡇⢸⠉⢸⡇⠘⠉⣿⠀⠈⣥⠶⣿⠀
⠀⢿⣦⣤⡦⣤⣿⣤⡄⠸⣧⣴⣿⠀⠀⣿⡀⠀⠀⢿⣤⡄⠀⠀⠀⠀⠀⠸⣧⣤⠀⢿⣥⣾⠃⠀⠀⠀⠀⠸⣧⣤⣶⠀⠸⣧⣤⠀⢿⣤⣼⠃⢸⠀⢸⡇⢠⣴⣿⣤⠘⣧⣤⣿⡄
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣶⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀
Minecraft to Mineclonia Texture Pack Converter
	release: %s
	Using Minecraft Pack Version: %s`, version, mc_version)
	fmt.Print("\n\n")

	config, err = configure.ReadConfigFile(ConfigLocation)
	if errors.Is(err, os.ErrNotExist) {
		config = configure.NewConfig()
		fmt.Println("Making the config.json file. Directories are unlikely to match your own.")
		if saveErr := config.SaveConfig(ConfigLocation); saveErr != nil {
			log.Fatal(saveErr)
		}
	} else if err != nil {
		log.Fatal("coudn't read config.", err)
	}

	fmt.Println(config)
	fmt.Print("\n\n")

	if !config.DefinedInput {
		config.InputDir = "./input/"
		if fs.ValidPath("input") {
			err := os.Mkdir("input", 0755)
			if err != nil {
				if errors.Is(err, fs.ErrPermission) {
					log.Fatal("Permission was denied. \"input\" was not made.")
				} else if errors.Is(err, fs.ErrExist) {
					fmt.Println("Folder \"input\" already exists.")
				} else {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Made the input folder!")
			}
		}
	}

	if !config.DefinedOutput {
		config.OutputDir = "./output/"
		if fs.ValidPath("output") {
			if err := os.Mkdir("output", 0755); err != nil {
				if errors.Is(err, fs.ErrPermission) {
					log.Fatal("Permission was denied. \"output\" was not made.")
				} else if errors.Is(err, fs.ErrExist) {
					fmt.Println("Folder \"output\" already exists.")
				} else {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Made the output folder!")
			}
		}
	}

	inputFiles, err := os.ReadDir(config.InputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, inputFile := range inputFiles {
		fileExt := filepath.Ext(inputFile.Name())
		if !inputFile.IsDir() && (fileExt == ".zip" || fileExt == ".jar") {
			if _, err := os.Stat(config.InputDir + "/" + FileNameWithoutExt(inputFile.Name())); errors.Is(err, os.ErrNotExist) {
				fmt.Println("Unzipping:", inputFile.Name())
				if err2 := unzipSource(config.InputDir+"/"+inputFile.Name(), config.InputDir+"/"+FileNameWithoutExt(inputFile.Name())); err2 != nil {
					fmt.Println("Extraction Error:", err)
				}
			} else {
				fmt.Println(inputFile.Name(), "was already decompressed! :D")
			}
		}
	}

	if err := assetsOnlyChecks(config.InputDir); err != nil {
		fmt.Println(err)
	}

	var outputDir *os.File
	outputDir, err = os.Open(config.OutputDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			log.Println(config.InputDir + "Output folder doesn't exist.")
		}
		log.Fatal(err)
	}
	defer outputDir.Close()

	inputFiles, err = os.ReadDir(config.InputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, inputFile := range inputFiles {
		if inputFile.IsDir() {
			fmt.Println(inputFile.Name())
			if config.ExportMineclonia {
				o := fmt.Sprintf("%s_mc_to_clonia", strings.ReplaceAll(strings.ToLower(inputFile.Name()), " ", "_"))

				if config.ForceClearOldConvertPacks {
					if err := os.RemoveAll(filepath.Join(config.OutputDir, o)); err != nil {
						fmt.Println(err)
					} else {
					}
				}

				convertPackClonia(inputFile.Name(), o, config)
			}
			if config.ExportMinetestGame {
				o := fmt.Sprintf("%s_mc_to_mtg", strings.ReplaceAll(strings.ToLower(inputFile.Name()), " ", "_"))

				if config.ForceClearOldConvertPacks {
					if err := os.RemoveAll(filepath.Join(config.OutputDir, o)); err != nil {
						fmt.Println(err)
					}
				}

				convertPackMTG(inputFile.Name(), o, config)
			}

			/*
				// TODO: ADD VOXELIBRE AS A SEPERATE OPTION
				if config.ExportVoxeLibre {
					o := fmt.Sprintf("%s_mc_to_vl", strings.ReplaceAll(strings.ToLower(inputFile.Name()), " ", "_"))

					if config.ForceClearOldConvertPacks {
						if err := os.RemoveAll(filepath.Join(config.OutputDir, o)); err != nil {
							fmt.Println(err)
						}
					}

					// convertPackVL(inputFile.Name(), o, config)
				}
			*/

			fmt.Print("Done!\n\n")
		}
	}
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
