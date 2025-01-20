package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Redis struct {
		Addr string `mapstructure:"addr"`
	} `mapstructure:"redis"`
	Database struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"database"`
}

func LoadConfig() error {
	absPath, err := filepath.Abs("./configs")
	if err != nil {
		return err
	}

	viper.AddConfigPath(absPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}

func Get() Config {
	return config
}
