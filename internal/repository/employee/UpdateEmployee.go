package employee

import (
	"context"
	"fmt"
	"web-employee/internal/model"
)

func (r *repo) UpdateEmployee(ctx context.Context, id string, employee *model.Employee) error {
	query := `UPDATE employees 
		SET name = COALESCE(NULLIF($1, ''), name), surname = COALESCE(NULLIF($2, ''), surname), phone = COALESCE(NULLIF($3, ''), phone), passport_type = COALESCE(NULLIF($4, ''), passport_type), passport_number = COALESCE(NULLIF($5, ''), passport_number), department_name = COALESCE(NULLIF($6, ''), department_name), department_phone = COALESCE(NULLIF($7, ''), department_phone) WHERE id = $8`
	res, err := r.db.Exec(ctx, query, employee.Name, employee.Surname, employee.Phone, employee.Passport.Type, employee.Passport.Number, employee.Department.Name, employee.Department.Phone, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("employee with id %s not found", id)
	}
	return nil
}
