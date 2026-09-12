package config

import "os"

type Config struct {
	Port string
}

func MustLoadConfig() *Config {
	return &Config{
		Port: os.Getenv("PORT"),
	}
}