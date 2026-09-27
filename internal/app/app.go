package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	productRepo := postgres.NewPostgresProductRepository(pool)
	imageRepo := postgres.NewPostgresImageRepository(pool)

	clientService := service.NewClientService(clientRepo)
	supplierService := service.NewSupplierService(supplierRepo)
	productService := service.NewProductService(productRepo)
	imageService := service.NewImageService(imageRepo)

	clientHandler := handler.NewClientHandler(clientService, validate)
	supplierHandler := handler.NewSupplierHandler(supplierService, validate)
	productHandler := handler.NewProductHandler(productService, validate)
	imageHandler := handler.NewImageHandler(imageService, validate)

	router := handler.SetUpRouter(clientHandler, supplierHandler, productHandler, imageHandler)
	return &App{
		config: cfg,
		pool:   pool,
		router: router,
	}
}

func (a *App) Run() error {
	server := &http.Server{
		Addr: ":8080",
		Handler: a.router,
	}
	go func() {
		log.Println("server started")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Server is shutting down...")

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}
	a.pool.Close()
	log.Println("Server has been stopped")
	return nil
}
