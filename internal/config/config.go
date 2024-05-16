package config

import (
	"flag"
	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

func MustLoad() *Config {
	path := fetchPath()

	_, err := os.Stat(path)
	if err != nil {
		panic("config file not found: " + err.Error())
	}
	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read file: " + err.Error())
	}

	var cfg Config

	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to unmarshal cfg file: " + err.Error())
	}
	return &cfg
}

func fetchPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path to cfg")
	flag.Parse()

	if path == "" {
		path = ifNotExists()
	}

	return path
}

func ifNotExists() string {
	if err := godotenv.Load(); err != nil {
		panic("failed to load env")
	}
	path := os.Getenv("CONFIG_PATH")
	return path
}
