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

	config.CurrentGroup = newGroup
	jsonPath := filepath.Join(dirPath, ".godoCfg.json") // Open to JSON file

	file, err := os.Open(jsonPath)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := json.Marshal(config) // Encode to JSON
	if err != nil {
		return err
	}

	err = os.WriteFile(file.Name(), data, 0777) // Write to JSON data into file
	if err != nil {
		return err
	}

	return nil // Sucessfully changed currentGroup in config
}
