package main

import (
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
		db.Close()
	}()

	repo := repository.NewComputerRepository(db)
	handler := handlers.NewComputerHandler(repo)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := router.Group("/api")
	{
		api.POST("/computers", handler.CreateComputer)
		api.GET("/computers", handler.GetAllComputers)
		api.GET("/computers/:id", handler.GetComputer)
		api.PUT("/computers/:id", handler.UpdateComputer)
		api.DELETE("/computers/:id", handler.DeleteComputer)
		api.GET("/employees/:employee/computers", handler.GetComputersByEmployee)
	}

	log.Info("Server starting", slog.String("port", cfg.Port))
	err = router.Run(":" + cfg.Port)
	if err != nil {
		return
	}
}
