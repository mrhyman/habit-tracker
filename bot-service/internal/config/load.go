package config

import (
	"flag"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

const (
	configPathCLIArg  = "config-path"
	configPathENV     = "CONFIG_PATH"
	configPathDefault = "./config/local.yml"
)

var (
	log *slog.Logger
)

func init() {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	log = slog.New(slog.NewJSONHandler(os.Stdout, opts))
}

func MustLoad() *Config {

	viper.SetConfigFile(getConfigPath())

	if err := viper.ReadInConfig(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	return config
}

func getConfigPath() string {

	var res string

	flag.StringVar(&res, configPathCLIArg, "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv(configPathENV)
		if res != "" {
			log.Info("Config file not set, using ENV variable")
			return res
		}
	}

	if res == "" {
		log.Info("Config file not set, using default")
		return configPathDefault
	}

	return res
}
