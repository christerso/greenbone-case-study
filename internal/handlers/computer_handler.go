package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"greenbone-computer-inventory/internal/models"
	"greenbone-computer-inventory/internal/repository"
	"net"
	"net/http"
	"regexp"
	"strings"

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

	if err := h.ValidateComputer(&computer); err != nil {
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

	if err := h.ValidateComputer(&computer); err != nil {
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

type NotificationPayload struct {
	Level                string `json:"level"`
	EmployeeAbbreviation string `json:"employeeAbbreviation"`
	Message              string `json:"message"`
}

func (h *ComputerHandler) sendNotification(employeeAbbr string, count int) {
	payload := NotificationPayload{
		Level:                "warning",
		EmployeeAbbreviation: employeeAbbr,
		Message:              fmt.Sprintf("Employee %s has %d computers assigned", employeeAbbr, count),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	resp, err := http.Post("http://localhost:8080/api/notify", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func (h *ComputerHandler) ValidateComputer(computer *models.Computer) error {
	if strings.TrimSpace(computer.ComputerName) == "" {
		return fmt.Errorf("computer_name is required")
	}

	if strings.TrimSpace(computer.IPAddress) == "" {
		return fmt.Errorf("ip_address is required")
	}

	if net.ParseIP(computer.IPAddress) == nil {
		return fmt.Errorf("invalid ip_address format")
	}

	if strings.TrimSpace(computer.MACAddress) == "" {
		return fmt.Errorf("mac_address is required")
	}

	macRegex := regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	if !macRegex.MatchString(computer.MACAddress) {
		return fmt.Errorf("invalid mac_address format")
	}

	if computer.EmployeeAbbreviation != nil {
		abbr := strings.TrimSpace(*computer.EmployeeAbbreviation)
		if len(abbr) != 3 {
			return fmt.Errorf("employee_abbreviation must be exactly 3 characters")
		}
		*computer.EmployeeAbbreviation = abbr
	}

	return nil
}
