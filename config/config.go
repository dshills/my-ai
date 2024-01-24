package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	APIKey       string `json:"apikey"`
	Model        string `json:"model"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}

func Load(path string) (*Configuration, error) {
	str, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config.LoadWithPath: %w", err)
	}
	conf := Configuration{}
	if err = json.Unmarshal(str, &conf); err != nil {
		return nil, fmt.Errorf("config.LoadWithPath: %w", err)
	}
	return &conf, nil
}
