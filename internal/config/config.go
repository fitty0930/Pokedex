package config

import "io/ioutil"

// DbConfig ...
type DbConfig struct {
	Driver string `yaml:"driver"`
}

// Config ...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml: "version"`
}

// LoadConfig ...
func LoadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
}
