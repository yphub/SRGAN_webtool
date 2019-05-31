package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config ...
type Config struct {
	HTTPAddress string
	TFServing   ConfigTFServing
	StaticFile  string
}

// ConfigTFServing ...
type ConfigTFServing struct {
	Address   string
	ModelName string
	SigName   string
}

// LoadConfig ...
func LoadConfig(file string) (*Config, error) {
	config := &Config{}

	configBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
