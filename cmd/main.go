package main

import (
	"context"
	"errors"
	"greenbone-computer-inventory/internal/api"
	"greenbone-computer-inventory/internal/config"
	"greenbone-computer-inventory/internal/database"
	"greenbone-computer-inventory/internal/handlers"
	"greenbone-computer-inventory/internal/repository"
	"greenbone-computer-inventory/pkg/logger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.New()

	cfg := config.Load()
	log.Info("Configuration loaded", slog.String("port", cfg.Port))

	db, err := database.Connect(cfg.DatabaseURL, log)
	if err != nil {
		log.Error("Failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer func() {
		log.Info("Closing database connection")
		err := db.Close()
		if err != nil {
			log.Error("Failed to close database connection", slog.String("error", err.Error()))
			return
		}
	}()

	repo := repository.NewComputerRepository(db)
	handler := handlers.NewComputerHandler(repo)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	api.SetupRoutes(router, handler, db)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Info("Server starting", slog.String("port", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("Server failed to start", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", slog.String("error", err.Error()))
	}

	log.Info("Server exited")
}
