package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	Storage    string `mapstructure:"STORAGE"` // "memory" or "postgres"
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Set Defaults
	if config.Port == "" {
		config.Port = "8080"
	}
	if config.Storage == "postgres" && config.DBSSLMode == "" {
		config.DBSSLMode = "disable"
	}
	if config.Storage == "" {
		config.Storage = "memory"
	}

	return &config, nil
}
