package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB_SOURCE      string `mapstructure:"DB_SOURCE"`
	DB_DRIVER      string `mapstructure:"DB_DRIVER"`
	SERVER_ADDRESS string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}
