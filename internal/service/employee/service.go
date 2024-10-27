package employee

import (
	"web-employee/internal/repository"
	"web-employee/internal/service"
)

type serv struct {
	EmployeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) service.EmployeeService {
	return &serv{EmployeeRepository: employeeRepository}
}
