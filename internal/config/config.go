package config

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type HttpServer struct {
	Address string
}
type Config struct {
	Env         string     `yaml:"env"`
	StoragePath string     `yaml:"storagepath"`
	Port        string     `yaml:"port"`
	HttpServer  HttpServer `yaml:"httpserver"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		root, err := rootDir()
		if err != nil {
			panic("cannot resolve project root: " + err.Error())
		}
		configPath = filepath.Join(root, "config", "local.yaml")
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

func rootDir() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", os.ErrInvalid
	}

	// internal/config -> project root
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
	return root, nil
}
