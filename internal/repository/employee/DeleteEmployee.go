package employee

import (
	"context"
	"fmt"
)

func (r *repo) DeleteEmployee(ctx context.Context, id string) error {
	query := `DELETE FROM employees WHERE id = $1`
	res, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("employee with id %s not found", id)
	}
	return nil
}
