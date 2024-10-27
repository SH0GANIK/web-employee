package employee

import (
	"context"
	"web-employee/internal/model"
)

func (r *serv) GetEmployeesByCompany(ctx context.Context, companyId string) ([]*model.Employee, error) {
	employee, err := r.EmployeeRepository.GetEmployeesByCompany(ctx, companyId)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
