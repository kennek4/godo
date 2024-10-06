package envhelper

import (
	"os"

	"github.com/spf13/viper"
)

func GetEnvString(variable string) (result string, err error) {
	err = viper.Get()
	return result, nil
}
