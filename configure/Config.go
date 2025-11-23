package configure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DefinedInput       bool
	DefinedOutput      bool
	ExportMinetestGame bool
	ExportMineclonia   bool

	InputDir  string
	OutputDir string

	HUDOnFireAnimationFrames int
	SpearVersion             string
}

func NewConfig() *Config {
	var config = Config{
		DefinedInput:       false,
		DefinedOutput:      false,
		ExportMinetestGame: false,
		ExportMineclonia:   true,

		HUDOnFireAnimationFrames: 8,
		SpearVersion:             "short",
	}

	if userHomeDir, err := os.UserHomeDir(); err != nil {
		fmt.Println("You have no home directory? %w", err)
		config.InputDir = ("./input/")
		config.OutputDir = ("./output/")
	} else {
		config.InputDir =
			filepath.Join(userHomeDir, ".minecraft", "resourcepacks")
		config.OutputDir =
			filepath.Join(userHomeDir, ".var", "app", "org.luanti.luanti", ".minetest", "textures")
	}
	return &config
}

// Reads the config file.
// If feilds are missing from the file, it will add them.
func ReadConfigFile(configLocation string) (*Config, error) {
	var config *Config
	var err error

	if _, statErr := os.Stat(configLocation); errors.Is(statErr, os.ErrNotExist) {
		return nil, statErr
	}

	configFile, err := os.ReadFile(configLocation)
	if err != nil {
		fmt.Println("couldn't read config file:", err)
		return nil, err
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("couldn't marshal config json:", err)
		return nil, err
	}

	needsUpdate, updatedConfigFile, err := LegacyJsonConfigFileUpdater(configFile, config)
	if err != nil {
		fmt.Println("config was not updated:", err)
	} else if needsUpdate {
		configFile = updatedConfigFile
		if err := os.WriteFile(configLocation, []byte(configFile), 0644); err != nil {
			fmt.Println("couldn't update config file:", err)
			return nil, err
		}
	}

	config.InputDir = CutHomePath(config.InputDir) + "/"
	config.OutputDir = CutHomePath(config.OutputDir) + "/"

	if !config.DefinedInput {
		config.InputDir = ("./input/")
	}
	if !config.DefinedOutput {
		config.InputDir = ("./output/")
	}

	return config, nil
}

func (c *Config) SaveConfig(saveLocation string) error {
	configData, err := json.MarshalIndent(*c, "", "")
	if err != nil {
		return err
	}
	if err := os.WriteFile(saveLocation, []byte(configData), 0644); err != nil {
		return err
	}
	return nil
}

func (c *Config) String() string {
	if configData, err := json.MarshalIndent(*c, "", ""); err != nil {
		return err.Error()
	} else {
		return (string(configData))
	}
}

// Only checks lines that end in commas. Sorry if you don't like that kind of formatting.
func LegacyJsonConfigFileUpdater(file []byte, c *Config) (needsUpdate bool, updatedFile []byte, err error) {
	needsUpdate = false

	referenceBytes, err := json.MarshalIndent(*NewConfig(), "", "")
	if err != nil {
		log.Panic(err)
	}

	referenceLineCount := bytes.Count(referenceBytes, []byte(",\n"))
	fileLineCount := bytes.Count(file, []byte(",\n"))
	if referenceLineCount != fileLineCount {
		needsUpdate = true
	}

	if bytes.Contains(file, []byte("ExportMinetest_Game")) {
		needsUpdate = true
	}

	// Start of field checks
	if c.HUDOnFireAnimationFrames == 0 {
		c.HUDOnFireAnimationFrames = 8
	}
	if c.SpearVersion == "" {
		c.SpearVersion = "short"
	}
	// End of field checks

	if !needsUpdate {
		return
	}

	updatedFile, err = json.MarshalIndent(*c, "", "")
	if err != nil {
		return
	}
	return true, updatedFile, nil
}
