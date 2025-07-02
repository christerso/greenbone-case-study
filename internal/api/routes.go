package api

import (
	"database/sql"
	"greenbone-computer-inventory/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *handlers.ComputerHandler, db *sql.DB) {
	router.GET("/health", healthHandler(db))

	api := router.Group("/api")
	{
		api.POST("/computers", handler.CreateComputer)
		api.GET("/computers", handler.GetAllComputers)
		api.GET("/computers/:id", handler.GetComputer)
		api.PUT("/computers/:id", handler.UpdateComputer)
		api.DELETE("/computers/:id", handler.DeleteComputer)
		api.GET("/employees/:employee/computers", handler.GetComputersByEmployee)
	}
}

func healthHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := gin.H{"status": "ok"}

		if err := db.Ping(); err != nil {
			status["status"] = "error"
			status["database"] = "disconnected"
			c.JSON(503, status)
			return
		}

		status["database"] = "connected"
		c.JSON(200, status)
	}
}
