package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Fields struct {
	Server struct {
		Port string `yaml:"port" envconfig:"PORT"`
	} `yaml:"server"`
	Database struct {
		Primary struct {
			Username       string `yaml:"user" envconfig:"DB_USER"`
			Password       string `yaml:"pass" envconfig:"DB_PASS"`
			ConnectionPath string `yaml:"connectionPath" envconfig:"DB_CONN"`
		} `yaml:"primary"`
	} `yaml:"database"`
}

func populateFields(config *Config) {
	fields := &Fields{}

	loadFieldsFromYaml(fields)
	loadFieldsFromEnv(fields)
	// TODO: loadFromSecretsManager

	config.Fields = fields
}

func loadFieldsFromYaml(config *Fields) {
	filePath := fmt.Sprintf("config.%s.yml", os.Getenv("ENV"))
	file, err := os.Open(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println("Config Yaml not found:", filePath)
			return
		} else {
			panic(err)
		}
	}
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		panic(err)
	}
	file.Close()
}

func loadFieldsFromEnv(config *Fields) {
	err := envconfig.Process("", config)
	if err != nil {
		panic(err)
	}
}
