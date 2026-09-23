package main

import (
	"github.com/sonofsun61/APIFromSpec/internal/app"
	"github.com/sonofsun61/APIFromSpec/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()
	application := app.NewApp(cfg)
	if err := application.Run(); err != nil {
		panic(err)
	}
}
