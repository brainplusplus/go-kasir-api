package config

import (
	"fmt"
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
		fmt.Println("Local .env file found, loading config from file...")
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	} else {
		fmt.Println("No .env file found, using OS Environment Variables")
	}

	var config Config
	// Manual mapping to ensure keys are read from Env even if .env is missing
	config.Port = viper.GetString("PORT")
	config.Storage = viper.GetString("STORAGE")
	config.DBHost = viper.GetString("DB_HOST")
	config.DBUser = viper.GetString("DB_USER")
	config.DBPassword = viper.GetString("DB_PASSWORD")
	config.DBName = viper.GetString("DB_NAME")
	config.DBPort = viper.GetString("DB_PORT")
	config.DBSSLMode = viper.GetString("DB_SSLMODE")

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

	fmt.Printf("Config loaded from: %s\n", config.Storage)
	fmt.Printf("Running on Port: %s\n", config.Port)

	return &config, nil
}
