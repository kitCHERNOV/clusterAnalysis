package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	ClusterCount int `yaml:"n_clusters"`
}

func LoadConfig() *Config {
	cfg := &Config{
		ClusterCount: 3, // as default
	}

	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		log.Fatal(err)
		return nil
	}

	return cfg
}
