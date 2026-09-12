package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sonofsun61/APIFromSpec.git/internal/config"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg config.Config) *App {
	return &App{
		cfg: &cfg,
	}
}

func (app *App) Run() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := chi.NewRouter()

	server := http.Server {
		Addr: "",
		Handler: router,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go func() {
		logger.Info("Server started on", "address", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Server startup failed")
		}
	}()
	<- ctx.Done()
	logger.Info("Recieved a stop signal. Gracefil Shutdown started...")
	shudownCtx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if err := server.Shutdown(shudownCtx); err != nil {
		logger.Warn("Graceful shutdown failed")
	}
	logger.Info("Shutdown completed")
}