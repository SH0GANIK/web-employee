package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-employee/internal/model"
)

func (h *Handler) CreateEmployee(c *gin.Context) {
	ctx := c.Request.Context()
	var employee model.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.CreateEmployee(ctx, &employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) GetEmployeesByCompany(c *gin.Context) {
	ctx := c.Request.Context()
	employees, err := h.service.GetEmployeesByCompany(ctx, c.Param("company_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) GetEmployeesByDepartment(c *gin.Context) {
	ctx := c.Request.Context()
	employees, err := h.service.GetEmployeesByDepartment(ctx, c.Param("company_id"), c.Param("department_name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	ctx := c.Request.Context()
	err := h.service.DeleteEmployee(ctx, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	ctx := c.Request.Context()
	var employee model.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateEmployee(ctx, c.Param("id"), &employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
