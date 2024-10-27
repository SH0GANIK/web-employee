package employee

import (
	"context"
	"web-employee/internal/model"
)

func (r *serv) UpdateEmployee(ctx context.Context, id string, employee *model.Employee) error {
	err := r.EmployeeRepository.UpdateEmployee(ctx, id, employee)
	if err != nil {
		return err
	}
	return nil
}
