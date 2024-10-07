package configs

import "github.com/spf13/viper"

func SetCurrentGroup(newGroup string) {
	viper.GetViper().Set("group", newGroup)
}
