package env

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
