package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Nginx Nginx `yaml:"nginx"`
}

type Nginx struct {
	Location string `yaml:"location"`
	Template string `yaml:"template"`
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		slog.Info("Error loading .env file", err)
	}

	configPath := "config.yml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		slog.Error("Error reading config file", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		slog.Error("Error unmarshalling config file", err)
	}

	return config
}
