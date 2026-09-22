package main

import (
	"context"
	"fmt"

	"github.com/sonofsun61/APIFromSpec/internal/config"
	"github.com/sonofsun61/APIFromSpec/internal/database"
)

func main() {
	cfg := config.MustLoadConfig()
	ctx := context.Background()
	pool := database.MustConnectToDatabase(ctx, cfg.ConnString)
	defer pool.Close()
	fmt.Println("connected to database successfully")
}