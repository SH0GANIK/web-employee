package employee

import (
	"github.com/jackc/pgx/v5"
	"web-employee/internal/repository"
)

type repo struct {
	db *pgx.Conn
}

func NewEmployeeRepository(db *pgx.Conn) repository.EmployeeRepository {
	return &repo{db: db}
}
