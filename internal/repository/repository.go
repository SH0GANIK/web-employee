package repository

import (
	"context"
	"web-employee/internal/model"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, info *model.Employee) (string, error)
	DeleteEmployee(ctx context.Context, id string) error
	GetEmployeesByCompany(ctx context.Context, companyId string) ([]*model.Employee, error)
	GetEmployeesByDepartment(ctx context.Context, companyId string, departmentName string) ([]*model.Employee, error)
	UpdateEmployee(ctx context.Context, id string, employee *model.Employee) error
}
