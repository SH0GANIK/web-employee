package employee

import "context"

func (r *serv) DeleteEmployee(ctx context.Context, id string) error {
	err := r.EmployeeRepository.DeleteEmployee(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
