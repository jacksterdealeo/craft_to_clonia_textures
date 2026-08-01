package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

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

	// GUI CODE HERE

	a := app.NewWithID("Craft to Clonia Textures")
	w := a.NewWindow("Craft to Clonia Textures")

	progress := widget.NewProgressBarInfinite()

	exportMinetestGameBool := binding.NewBool()
	exportMinecloniaBool := binding.NewBool()
	inputDirString := binding.NewString()
	outputDirString := binding.NewString()
	forceClearOldConvertPacksBool := binding.NewBool()
	hudOnFireAnimationFramesString := binding.NewString()
	spearVersionString := binding.NewString()
	enlargeObjectCrosshair := binding.NewBool()

	expect(exportMinetestGameBool.Set(config.ExportMinetestGame))
	expect(exportMinecloniaBool.Set(config.ExportMineclonia))
	expect(inputDirString.Set(config.InputDir))
	expect(outputDirString.Set(config.OutputDir))
	expect(forceClearOldConvertPacksBool.Set(config.ForceClearOldConvertPacks))
	expect(hudOnFireAnimationFramesString.Set(strconv.Itoa(config.HUDOnFireAnimationFrames)))
	expect(spearVersionString.Set(config.SpearVersion))
	expect(enlargeObjectCrosshair.Set(config.EnlargeObjectCrosshair))

	saveToConfig := func() {
		go func() {
			freshConfig := configure.NewConfig()
			// Save window state into config here:
			config.ExportMinetestGame, err = exportMinetestGameBool.Get()
			expect(err)

			config.ExportMineclonia, err = exportMinecloniaBool.Get()
			expect(err)
			config.InputDir, err = inputDirString.Get()
			expect(err)
			config.OutputDir, err = outputDirString.Get()
			expect(err)
			config.ForceClearOldConvertPacks, err = forceClearOldConvertPacksBool.Get()
			expect(err)

			a, err := hudOnFireAnimationFramesString.Get()
			expect(err)
			config.HUDOnFireAnimationFrames, err = strconv.Atoi(a)
			if err != nil {
				expect(hudOnFireAnimationFramesString.Set(strconv.Itoa(freshConfig.HUDOnFireAnimationFrames)))
			}

			config.SpearVersion, err = spearVersionString.Get()
			expect(err)
			config.EnlargeObjectCrosshair, err = enlargeObjectCrosshair.Get()
			expect(err)

			expect(config.SaveConfig(ConfigLocation))
			fyne.Do(func() {
				progress.Start()
			})
			convertPacks(config)
			fyne.Do(func() {
				progress.Stop()
			})
		}()
	}

	w.SetContent(container.NewVBox(
		widget.NewLabel("Craft to Clonia Textures"),

		widget.NewCheckWithData("Export for Minetest Game?", exportMinetestGameBool),
		widget.NewCheckWithData("Export for Mineclonia?", exportMinecloniaBool),

		widget.NewEntryWithData(inputDirString),
		widget.NewButton("Open Input Directory", func() {
			dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				if list == nil {
					log.Println("Cancelled")
					return
				}

				children, err := list.List()
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				out := fmt.Sprintf("Folder %s (%d children):\n%s", list.Name(), len(children), list.String())
				dialog.ShowInformation("Folder Open", out, w)
			}, w)
		}),

		widget.NewEntryWithData(outputDirString),
		widget.NewButton("Open Output Directory", func() {
			dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				if list == nil {
					log.Println("Cancelled")
					return
				}

				children, err := list.List()
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				out := fmt.Sprintf("Folder %s (%d children):\n%s", list.Name(), len(children), list.String())
				dialog.ShowInformation("Folder Open", out, w)
			}, w)
		}),
		widget.NewCheckWithData("Delete all old converted packs from output directory?", forceClearOldConvertPacksBool),
		widget.NewEntryWithData(hudOnFireAnimationFramesString),
		widget.NewEntryWithData(spearVersionString),
		widget.NewCheckWithData("Enlarge crosshair when looking at an object?", enlargeObjectCrosshair),

		widget.NewButton("Save Config and Convert", saveToConfig),
		progress,
	))
	progress.Stop()

	w.Resize(fyne.NewSize(800, 400))
	w.ShowAndRun()

	//

}

func convertPacks(c *configure.Config) {
	fmt.Println("Converting the Packs")
	inputFiles, err := os.ReadDir(c.InputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, inputFile := range inputFiles {
		fileExt := filepath.Ext(inputFile.Name())
		if !inputFile.IsDir() && (fileExt == ".zip" || fileExt == ".jar") {
			if _, statErr := os.Stat(filepath.Join(c.InputDir,
				FileNameWithoutExt(inputFile.Name()))); errors.Is(statErr, os.ErrNotExist) {
				fmt.Println("Unzipping:", inputFile.Name())
				if unzipErr := unzipSource(
					filepath.Join(
						c.InputDir, inputFile.Name()),
					filepath.Join(
						c.InputDir, FileNameWithoutExt(inputFile.Name()))); unzipErr != nil {
					fmt.Println("Extraction Error:", unzipErr)
				}
			} else {
				fmt.Println(inputFile.Name(), "was already decompressed! 🐭 Skipping this archive.")
			}
		}
	}

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
			fmt.Println(inputFile.Name())
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

			fmt.Print("Done!\n\n")
		}
	}
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
