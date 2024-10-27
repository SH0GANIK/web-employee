package employee

import (
	"context"
	"web-employee/internal/model"
)

func (r *repo) CreateEmployee(ctx context.Context, info *model.Employee) (string, error) {
	query :=
		`INSERT INTO employees
		(name, surname, phone, companyid, passport_type, passport_number, department_name, department_phone)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	var id string
	err := r.db.QueryRow(ctx, query, info.Name, info.Surname, info.Phone, info.CompanyId, info.Passport.Type, info.Passport.Number, info.Department.Name, info.Department.Phone).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
