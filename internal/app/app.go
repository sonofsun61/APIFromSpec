package app

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonofsun61/APIFromSpec/internal/config"
	"github.com/sonofsun61/APIFromSpec/internal/database"
	"github.com/sonofsun61/APIFromSpec/internal/handler"
	"github.com/sonofsun61/APIFromSpec/internal/repository/postgres"
	"github.com/sonofsun61/APIFromSpec/internal/service"
)

type App struct {
	config *config.Config
	pool   *pgxpool.Pool
	router http.Handler
}

func NewApp(cfg *config.Config) *App {
	validate := validator.New()
	pool := database.MustConnectToDatabase(context.Background(), cfg.ConnString)
	clientRepo := postgres.NewPostgresClientRepository(pool)
	supplierRepo := postgres.NewPostgresSupplierRepository(pool)
	clientService := service.NewClientService(clientRepo)
	supplierService := service.NewSupplierService(supplierRepo)
	clientHandler := handler.NewClientHandler(clientService, validate)
	supplierHandler := handler.NewSupplierHandler(supplierService, validate)
	router := handler.SetUpRouter(clientHandler, supplierHandler)
	return &App{
		config: cfg,
		pool:   pool,
		router: router,
	}
}

func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.router)
}
