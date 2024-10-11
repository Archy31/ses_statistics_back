package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"ses_back/internal/models"
)

var Config models.AppConfig

func GetConfigPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	configPath := filepath.Join(wd, "config.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Panic(err)
	}

	return configPath, nil
}

func ReadConfig() {
	YamlPath, err := GetConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(YamlPath)

	if err != nil {
		log.Panic(err)
	}

	err = yaml.Unmarshal(file, &Config)

	if err != nil {
		log.Panic(err)
	}
}
