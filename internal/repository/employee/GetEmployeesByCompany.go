package employee

import (
	"context"
	"web-employee/internal/model"
)

func (r *repo) GetEmployeesByCompany(ctx context.Context, companyId string) ([]*model.Employee, error) {
	query := `SELECT id, name, surname, phone, companyid, passport_type, passport_number, department_name, department_phone FROM employees WHERE companyid = $1`
	rows, err := r.db.Query(ctx, query, companyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employees := make([]*model.Employee, 0)
	for rows.Next() {
		employee := &model.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Surname, &employee.Phone, &employee.CompanyId, &employee.Passport.Type, &employee.Passport.Number, &employee.Department.Name, &employee.Department.Phone)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
