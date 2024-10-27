package handler

import (
	"github.com/gin-gonic/gin"
	"web-employee/internal/service"
)

type Handler struct {
	service service.EmployeeService
}

func NewHandler(service service.EmployeeService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	employees := router.Group("/employees")
	{
		employees.GET("/company/:company_id", h.GetEmployeesByCompany)
		employees.GET("/company/:company_id/:department_name", h.GetEmployeesByDepartment)
		employees.POST("", h.CreateEmployee)
		employees.DELETE("/:id", h.DeleteEmployee)
		employees.PUT("/:id", h.UpdateEmployee)
	}
	return router
}
