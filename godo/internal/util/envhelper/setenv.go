package envhelper

import (
	"fmt"
	"os"
)

func SetEnvString(key string, value string) error {
	envVariable := fmt.Sprintf("GODO_%s", key)
	err := os.Setenv(envVariable, value)
	if err != nil {
		return err
	}

	return nil
}
