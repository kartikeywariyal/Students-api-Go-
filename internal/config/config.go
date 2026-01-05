package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type HttpServer struct {
	Address string
}
type Config struct {
	Env         string
	StoragePath string
	Port        string
	HttpServer  HttpServer
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/local.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic("cannot read config: " + err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic("cannot parse config: " + err.Error())
	}

	return &cfg
}
