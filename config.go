package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type config struct {
	DefinedInput        bool
	DefinedOutput       bool
	ExportMinetest_Game bool
	ExportMineclonia    bool
	InputDir            string
	OutputDir           string
}

var Config *config = &config{
	DefinedInput:        false,
	DefinedOutput:       false,
	ExportMinetest_Game: false,
	ExportMineclonia:    true,
	InputDir:            "./input/",
	OutputDir:           "./output/",
}

var (
	ConfigLocation = "config.json"
)

func loadJsonConfig() (*config, error) {
	if _, statErr := os.Stat(ConfigLocation); errors.Is(statErr, os.ErrNotExist) {
		fmt.Println("Making the config.json file. Directories are unlikely to match your own.")
		if userHomeDir, err := os.UserHomeDir(); err != nil {
			fmt.Println("You have no home directory?", err)
			return Config, err
		} else {
			Config.InputDir =
				userHomeDir + "/.minecraft/resourcepacks/"
			Config.OutputDir =
				userHomeDir + "/.var/app/org.luanti.luanti/.minetest/textures/"
		}

		configData, err := json.MarshalIndent(Config, "", "")

		if err != nil {
			fmt.Println("Couldn't Marshal config json :", err)
			return nil, err
		}
		if err := os.WriteFile(ConfigLocation, []byte(configData), 0644); err != nil {
			return nil, err
		}
		return Config, nil
	}

	configData, err := os.ReadFile(ConfigLocation)
	if err != nil {
		fmt.Println(err)
		return Config, err
	}
	fmt.Printf("CONFIG FILE: %v\n", string(configData))
	err = json.Unmarshal([]byte(configData), &Config)
	if err != nil {
		fmt.Println("Couldn't Marshal config json :", err)
		return nil, err
	}
	if err := os.WriteFile(ConfigLocation, []byte(configData), 0644); err != nil {
		return nil, err
	}

	// Sometimes people forget to add a slash to the end of their directories.
	// If they didn't forget, the extra slash is ignored anyway.
	Config.InputDir += "/"
	Config.OutputDir += "/"

	if !Config.DefinedInput {
		Config.InputDir = "input"
	}
	if !Config.DefinedOutput {
		Config.OutputDir = "output"
	}
	return Config, nil
}
