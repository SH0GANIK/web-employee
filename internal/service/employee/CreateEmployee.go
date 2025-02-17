package employee

import (
	"context"
	"web-employee/internal/model"
)

func (r *serv) CreateEmployee(ctx context.Context, info *model.Employee) (string, error) {
	id, err := r.EmployeeRepository.CreateEmployee(ctx, info)
	if err != nil {
		return "", err
	}
	return id, nil
}
