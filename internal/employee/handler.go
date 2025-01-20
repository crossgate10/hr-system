package employee

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, service Service) {
	h := &Handler{service: service}
	router.POST("/employees", h.AddEmployee)
	router.GET("/employees/:id", h.GetEmployeeDetails)
	router.PUT("/employees/:id", h.UpdateEmployee)
	router.DELETE("/employees/:id", h.DeleteEmployee)
	router.GET("/employees", h.ListEmployees)
}

type Handler struct {
	service Service
}

type AddEmployeeRequest struct {
	Name        string  `json:"name" binding:"required" example:"John Doe"`
	Position    string  `json:"position" binding:"required" example:"Software Engineer"`
	ContactInfo string  `json:"contact_info" binding:"required" example:"john.doe@example.com"`
	Salary      float64 `json:"salary" binding:"required" example:"50000"`
}

type AddEmployeeResponse struct {
	ID          int     `json:"id" example:"1"`
	Name        string  `json:"name" example:"John Doe"`
	Position    string  `json:"position" example:"Software Engineer"`
	ContactInfo string  `json:"contact_info" example:"john.doe@example.com"`
	Salary      float64 `json:"salary" example:"50000"`
}

func (h *Handler) AddEmployee(c *gin.Context) {
	var req AddEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEmp, err := h.service.AddEmployee(c.Request.Context(), req.Name, req.Position, req.ContactInfo, req.Salary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := AddEmployeeResponse{
		ID:          createdEmp.ID,
		Name:        createdEmp.Name,
		Position:    createdEmp.Position,
		ContactInfo: createdEmp.ContactInfo,
		Salary:      createdEmp.Salary,
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetEmployeeDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}

	emp, err := h.service.GetEmployeeDetails(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emp)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}

	var emp Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	emp.ID = id

	if err := h.service.UpdateEmployeeDetails(c.Request.Context(), &emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emp)
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}

	if err := h.service.RemoveEmployee(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) ListEmployees(c *gin.Context) {
	filters := ListEmployeesFilter{
		Name:     c.Query("name"),
		Position: c.Query("position"),
	}

	minSalary := c.Query("minSalary")
	if minSalary != "" {
		if minSalary, err := strconv.ParseFloat(minSalary, 64); err == nil {
			filters.MinSalary = minSalary
		}
	}

	maxSalary := c.Query("maxSalary")
	if maxSalary != "" {
		if maxSalary, err := strconv.ParseFloat(maxSalary, 64); err == nil {
			filters.MaxSalary = maxSalary
		}
	}

	employees, err := h.service.ListEmployees(c.Request.Context(), &filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}
