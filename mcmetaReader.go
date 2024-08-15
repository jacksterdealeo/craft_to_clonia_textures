package main

import (
	"encoding/json"
	"os"
)

type Animation struct {
	Frames []int `json:"frames"`
}

type MCMETABlockData struct {
	Animation Animation `json:"animation"`
}

func McmetaReader(filePath string) (Frames []int, Error error) {
	file, err := os.ReadFile(filePath + ".mcmeta")
	if err != nil {
		return []int{}, err
	}

	jsonData := MCMETABlockData{}

	json.Unmarshal([]byte(file), &jsonData)
	return (jsonData.Animation.Frames), nil
}
