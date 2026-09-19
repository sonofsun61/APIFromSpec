package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
}

func MustLoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:123@localhost:8080/shopapi?sslomode=disabled"
	}
	return &Config{
		Port: port,
		DatabaseURL: dbURL,
	}
}
