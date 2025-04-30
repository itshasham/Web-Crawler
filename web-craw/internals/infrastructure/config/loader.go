package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName string `yaml:"app_name"`
	Port    string `yaml:"port"`
	DBURI   string `yaml:"db_uri"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
