package configs

import "github.com/spf13/viper"

func GetValueFromKey(key string) (value string) {
	return viper.GetViper().GetString(key)
}
