package envhelper

import (
	"fmt"
	"os"
	"strings"
)

func SetGodoEnv(key string, value *string) error {

	if value == nil {
		return fmt.Errorf("in SetGodoEnv, a value was passed a nil string pointer")
	}

	key = strings.ToUpper(key)
	envVariable := fmt.Sprintf("GODO_%s", key)

	err := os.Setenv(envVariable, *value)
	if err != nil {
		return err
	}

	return nil
}
