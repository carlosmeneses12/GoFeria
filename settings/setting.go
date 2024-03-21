package settings

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed setting.yaml
var settingsFile []byte

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Setting struct {
	Port string         `yaml:"port"`
	DB   DatabaseConfig `yaml:"database"`
}

func New() (*Setting, error) {
	var s Setting

	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil

}

