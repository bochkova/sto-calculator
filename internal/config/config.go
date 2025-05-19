package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type LoggerConfig struct {
	Level string `yaml:"level"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	Logger   LoggerConfig   `yaml:"logging"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func GetConfig(filePath string) (*Config, error) {
	config := &Config{}

	err := cleanenv.ReadConfig(filePath, config)
	if err != nil {
		return nil, fmt.Errorf("read config error: %v", err)
	}

	return config, nil
}
