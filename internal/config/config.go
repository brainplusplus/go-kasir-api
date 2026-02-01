package config

import (
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
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		// Just explicitly return default values or proceed if env vars are set
		// But usually we prefer to error if .env is missing AND env vars aren't set?
		// For now we can ignore error if file not found, but good to know.
		// However, returning error and handling in main is better.
		// If .env doesn't exist, we might be relying purely on Env vars.
		// Let's be lenient if file is missing but strict on unmarshaling.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
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
