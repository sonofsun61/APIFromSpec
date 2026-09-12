package main

import (
	"github.com/sonofsun61/APIFromSpec.git/internal/app"
	"github.com/sonofsun61/APIFromSpec.git/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()
	application := app.NewApp(cfg)
	application.Run()
}