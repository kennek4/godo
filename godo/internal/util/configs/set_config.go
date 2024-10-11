package configs

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SetCurrentGroup(newGroup string, dirPath string) error {
	config, err := getConfig(dirPath)
	if err != nil {
		return err
	}

	config.CurrentGroup = newGroup // Change values in JSON struct

	jsonPath := filepath.Join(dirPath, ".godoCfg.json")

	file, err := os.Open(jsonPath) // Open the JSON file
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := json.Marshal(config) // Encode to JSON
	if err != nil {
		return err
	}

	err = os.WriteFile(file.Name(), data, 0777) // Write JSON data into file
	if err != nil {
		return err
	}

	return nil // Sucessfully changed currentGroup in config
}
