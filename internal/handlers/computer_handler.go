package handlers

import (
	"greenbone-computer-inventory/internal/models"
	"greenbone-computer-inventory/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ComputerHandler struct {
	repo *repository.ComputerRepository
}

func NewComputerHandler(repo *repository.ComputerRepository) *ComputerHandler {
	return &ComputerHandler{repo: repo}
}

func (h *ComputerHandler) CreateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.ShouldBindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&computer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if computer.EmployeeAbbreviation != nil {
		count, err := h.repo.CountByEmployee(*computer.EmployeeAbbreviation)
		if err == nil && count >= 3 {
			go h.sendNotification(*computer.EmployeeAbbreviation, count)
		}
	}

	c.JSON(http.StatusCreated, computer)
}

func (h *ComputerHandler) GetComputer(c *gin.Context) {
	id := c.Param("id")
	computer, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Computer not found"})
		return
	}
	c.JSON(http.StatusOK, computer)
}

func (h *ComputerHandler) GetAllComputers(c *gin.Context) {
	computers, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, computers)
}

func (h *ComputerHandler) GetComputersByEmployee(c *gin.Context) {
	employee := c.Param("employee")
	computers, err := h.repo.GetByEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, computers)
}

func (h *ComputerHandler) UpdateComputer(c *gin.Context) {
	id := c.Param("id")
	var computer models.Computer
	if err := c.ShouldBindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	computer.ID = id
	if err := h.repo.Update(&computer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if computer.EmployeeAbbreviation != nil {
		count, err := h.repo.CountByEmployee(*computer.EmployeeAbbreviation)
		if err == nil && count >= 3 {
			go h.sendNotification(*computer.EmployeeAbbreviation, count)
		}
	}

	c.JSON(http.StatusOK, computer)
}

func (h *ComputerHandler) DeleteComputer(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *ComputerHandler) sendNotification(employeeAbbr string, count int) {
}
