package config

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"main/pkg"
	"os"

	"github.com/spf13/viper"
)

const (
	configPathCLIArg  = "config-path"
	configPathENV     = "CONFIG_PATH"
	configPathDefault = "./config/local.yml"
)

func MustLoad() *Config {
	ctx := context.Background()
	configFilePath := getConfigPath(ctx)
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		pkg.LogFatal(
			ctx,
			fmt.Sprintf("Can't read config file %s", configFilePath),
			err,
		)
	}

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		pkg.LogFatal(
			ctx,
			fmt.Sprintf("Can't parse config file %s", configFilePath),
			err,
		)
	}

	return config
}

func getConfigPath(ctx context.Context) string {
	var res string

	flag.StringVar(&res, configPathCLIArg, "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv(configPathENV)
		if res != "" {
			slog.InfoContext(ctx, "Config file not set, using ENV variable")
			return res
		}
	}

	if res == "" {
		slog.InfoContext(ctx, "Config file not set, using default")
		return configPathDefault
	}

	return res
}
