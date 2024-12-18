package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type config struct {
	DefinedInput  bool
	DefinedOutput bool
	InputDir      string
	OutputDir     string
}

var Config *config = &config{
	DefinedInput:  false,
	DefinedOutput: false,
	InputDir:      "input",
	OutputDir:     "output",
}

var (
	ConfigLocation = "config.json"
)

func loadJsonConfig() (*config, error) {
	if _, statErr := os.Stat(ConfigLocation); errors.Is(statErr, os.ErrNotExist) {
		fmt.Println("Making the config.json file.")
		if userHomeDir, err := os.UserHomeDir(); err != nil {
			fmt.Println("You have no home directory?", err)
			return Config, err
		} else {
			Config.InputDir =
				userHomeDir + "/.minecraft/resourcepacks"
			Config.OutputDir =
				userHomeDir + "/.var/app/net.minetest.Minetest/.minetest/textures"
		}

		configData, err := json.Marshal(Config)

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
	fmt.Println("FILE DATA: ", string(configData))
	err = json.Unmarshal([]byte(configData), &Config)
	if err != nil {
		fmt.Println("Couldn't Marshal config json :", err)
		return nil, err
	}
	if err := os.WriteFile(ConfigLocation, []byte(configData), 0644); err != nil {
		return nil, err
	}

	fmt.Println("REAL DATA: ", *Config)
	if !Config.DefinedInput {
		Config.InputDir = "input"
	}
	if !Config.DefinedOutput {
		Config.OutputDir = "output"
	}
	return Config, nil
}
