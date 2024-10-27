package employee

import (
	"context"
	"web-employee/internal/model"
)

func (r *serv) GetEmployeesByDepartment(ctx context.Context, companyId string, departmentName string) ([]*model.Employee, error) {
	employees, err := r.EmployeeRepository.GetEmployeesByDepartment(ctx, companyId, departmentName)
	if err != nil {
		return nil, err
	}
	return employees, nil
}
