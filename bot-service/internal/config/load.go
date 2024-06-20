package config

import (
	"context"
	"flag"
	"fmt"
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
	ctx := context.Background()
	configFilePath := getConfigPath(ctx)
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("Can't read config file %s", configFilePath),
			slog.String("err", err.Error()),
		)
		os.Exit(1)
	}

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("Can't parse config file %s", configFilePath),
			slog.String("err", err.Error()),
		)
		os.Exit(1)
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
			log.InfoContext(ctx, "Config file not set, using ENV variable")
			return res
		}
	}

	if res == "" {
		log.InfoContext(ctx, "Config file not set, using default")
		return configPathDefault
	}

	return res
}
