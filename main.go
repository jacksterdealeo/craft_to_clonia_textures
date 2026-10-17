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

func expect(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
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
	fmt.Println()

	var config *configure.Config
	var err error
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

	fmt.Print("Config: ", config, "\n\n")
	convertPacks(config)
}

func convertPacks(c *configure.Config) {
	inputFiles, err := os.ReadDir(c.InputDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Yellow + "Extracting Zips" + Reset)

	for _, inputFile := range inputFiles {
		fileExt := filepath.Ext(inputFile.Name())
		if !inputFile.IsDir() && (fileExt == ".zip" || fileExt == ".jar") {
			if _, statErr := os.Stat(filepath.Join(c.InputDir,
				FileNameWithoutExt(inputFile.Name()))); errors.Is(statErr, os.ErrNotExist) {
				fmt.Println("📁 Unzipping: " + inputFile.Name())
				if unzipErr := unzipSource(
					filepath.Join(
						c.InputDir, inputFile.Name()),
					filepath.Join(
						c.InputDir, FileNameWithoutExt(inputFile.Name()))); unzipErr != nil {
					fmt.Println(Red+"Extraction Error: "+Reset, unzipErr)
				}
			} else {
				fmt.Println("📂 " + inputFile.Name() + " was already decompressed! Skipping this archive.")
			}
		}
	}

	fmt.Println("\n" + Yellow + "Converting Packs" + Reset)

	if err := assetsOnlyChecks(c.InputDir); err != nil {
		fmt.Println(err)
	}

	var outputDir *os.File
	outputDir, err = os.Open(c.OutputDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			log.Println(c.InputDir, "Output folder doesn't exist.")
		}
		log.Fatal(err)
	}
	defer outputDir.Close()

	inputFiles, err = os.ReadDir(c.InputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, inputFile := range inputFiles {
		if inputFile.IsDir() {
			fmt.Println(Yellow + "Converting: " + inputFile.Name() + Reset)
			if c.ExportMineclonia {
				o := fmt.Sprintf("%s_mc_to_clonia", strings.ReplaceAll(strings.ToLower(inputFile.Name()), " ", "_"))

				if c.ForceClearOldConvertPacks {
					if err := os.RemoveAll(filepath.Join(c.OutputDir, o)); err != nil {
						fmt.Println(err)
					} else {
					}
				}

				convertPackClonia(inputFile.Name(), o, c)
			}

			if c.ExportMinetestGame {
				o := fmt.Sprintf("%s_mc_to_mtg", strings.ReplaceAll(strings.ToLower(inputFile.Name()), " ", "_"))

				if c.ForceClearOldConvertPacks {
					if err := os.RemoveAll(filepath.Join(c.OutputDir, o)); err != nil {
						fmt.Println(err)
					}
				}

				convertPackMTG(inputFile.Name(), o, c)
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

			fmt.Print(Green + "Finished Pack" + Reset + "\n\n")
		}
	}
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
