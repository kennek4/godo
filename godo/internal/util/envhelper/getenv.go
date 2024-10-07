package envhelper

import (
	"github.com/spf13/viper"
)

func GetGodoEnv(variable string) any {
	result := viper.Get(variable)
	return result
}
