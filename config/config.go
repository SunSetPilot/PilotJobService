package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppName   string
	DebugMode bool
	MySQL     string
	LogPath   string
}

func MustLoad(configFile string, config *Config) {
	var err error
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	if err = viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	viper.WatchConfig()
	return
}
