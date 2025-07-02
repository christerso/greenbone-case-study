package main

import (
	"greenbone-computer-inventory/internal/config"
	"greenbone-computer-inventory/internal/database"
	"greenbone-computer-inventory/pkg/logger"
	"log/slog"
	"os"
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
	defer db.Close()

	log.Info("Server starting", slog.String("port", cfg.Port))
}
