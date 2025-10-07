package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type ConfigProvider interface {
	Load() (*Config, error)
}

type FileConfigProvider struct {
	FilePath string
}

func (f FileConfigProvider) Load() (*Config, error) {
	data, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	return &config, err
}

type EnvConfigProvider struct{}

func (e EnvConfigProvider) Load() (*Config, error) {
	return &Config{
		Port: os.Getenv("APP_PORT"),
		Host: os.Getenv("APP_HOST"),
	}, nil
}

func main() {
	var provider ConfigProvider = EnvConfigProvider{}
	config, _ := provider.Load()
	fmt.Printf("Config: %+v\n", config)
}
