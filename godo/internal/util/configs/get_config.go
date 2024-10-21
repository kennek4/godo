package configs

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

func getConfig(dirPath string) (cfg ConfigFile, err error) {

	var configData ConfigFile
	configPath := filepath.Join(dirPath, ".godoCfg.json")

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Something went wrong with reading a file, %s\n", err)
	}

	defer file.Close()

	jsonData, err := io.ReadAll(file) // Convert JSON data to []byte
	if err != nil {
		log.Fatal("Something went wrong with converting file to []byte")
	}

	err = json.Unmarshal(jsonData, &configData) // Add []byte to ConfigFile struct
	if err != nil {
		log.Fatalf("Something went wrong with unmarshaling a file, %s\n", err)
	}

	return configData, nil
}

func GetCurrentGroup(dirPath string) string {
	config, err := getConfig(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	return config.CurrentGroup
}
