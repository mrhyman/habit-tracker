package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	configPathCLIArg  = "config-path"
	configPathENV     = "CONFIG_PATH"
	configPathDefault = "./config/local.yml"
)

func MustLoad() *Config {
	viper.SetConfigFile(getConfigPath())

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Fatal error decoding envs: %s \n", err)
	}

	return config
}

func getConfigPath() string {

	var res string

	flag.StringVar(&res, configPathCLIArg, "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv(configPathENV)
	}

	if res == "" {
		log.Println("Config file not set, using default")
		res = configPathDefault
	}

	return res
}
