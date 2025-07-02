package main

import (
	"greenbone-computer-inventory/internal/api"
	"greenbone-computer-inventory/internal/config"
	"greenbone-computer-inventory/internal/database"
	"greenbone-computer-inventory/internal/handlers"
	"greenbone-computer-inventory/internal/repository"
	"greenbone-computer-inventory/pkg/logger"
	"log/slog"
	"os"

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

	log.Info("Server starting", slog.String("port", cfg.Port))
	err = router.Run(":" + cfg.Port)
	if err != nil {
		return
	}
}
