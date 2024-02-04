package config

import (
	"flag"
	"github.com/MaksKazantsev/go-crud/internal/server"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Env         string             `yaml:"env"`
	StoragePath string             `yaml:"storage_path"`
	Port        string             `yaml:"port"`
	Address     string             `yaml:"address"`
	HTTPServer  *server.HTTPServer `yaml:"HTTPServer"`
}

func MustLoad() *Config {
	path := fetchPath()
	if path == "" {
		panic("no config file provided!")
	}
	_, err := os.Stat(path)
	if err != nil {
		panic("Cant find the config file")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		panic("Cant read the config file")
	}
	var cfg Config
	if err = yaml.Unmarshal(content, &cfg); err != nil {
		panic("failed to unmarshal config path")
	}
	return &cfg
}

func fetchPath() string {
	var path string
	flag.StringVar(&path, "c", "", "config path")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	return path
}
