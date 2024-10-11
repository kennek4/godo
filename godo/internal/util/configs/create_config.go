package configs

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func CreateConfigFile(dirPath string, defaultTable string) error {

	defaultConfig := ConfigFile{
		CurrentGroup: defaultTable,
	}

	path := filepath.Join(dirPath, ".godoCfg.json")
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := json.Marshal(defaultConfig) // Convert struct to JSON
	if err != nil {
		return err
	}

	err = os.WriteFile(file.Name(), data, 0777) // Write JSON to file
	if err != nil {
		return err
	}

	return nil
}
