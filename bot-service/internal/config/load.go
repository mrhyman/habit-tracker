package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func MustLoad() *Config {
	viper.SetConfigFile("config/config_local.yml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("Fatal error decoding envs: %s \n", err))
	}

	return config
}
