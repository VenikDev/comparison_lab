package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	CONFIG_FOR_SEARCH = "\\config_for_search"
	JSON_KEY_WORD     = "\\key_word.json"
	JSON_LABORATORY   = "\\laboratory.json"
)

func getWorkDir() string {
	wordDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wordDir
}

func pathToConfig() string {
	wordDir := getWorkDir()
	return wordDir + CONFIG_FOR_SEARCH
}

func ParseKeyValues() []KeyValue {
	pathToConfig := pathToConfig()

	jsonFile, err := os.Open(pathToConfig + JSON_KEY_WORD)
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our KeyValue array
	arrKeyValues := make([]KeyValue, 2)

	err = json.Unmarshal(byteValue, &arrKeyValues)
	if err != nil {
		panic(err)
	}

	return arrKeyValues
}

func ParseLabs() []Laboratory {
	pathToConfig := pathToConfig()

	jsonFile, err := os.Open(pathToConfig + JSON_LABORATORY)
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our KeyValue array
	arrLabs := make([]Laboratory, 2)

	err = json.Unmarshal(byteValue, &arrLabs)
	if err != nil {
		panic(err)
	}

	return arrLabs
}