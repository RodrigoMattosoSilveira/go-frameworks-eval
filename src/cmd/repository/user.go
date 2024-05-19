package repository

import (
	"database/sql"
	"fmt"

	"gofr.dev/pkg/gofr"

	"madronetek.com/go-frameworks-eval/cmd/model"
)

type repository struct{}

// New is a factory function for store layer that returns a interface type, UserInt
func New() RepositoryUserInt {
	return repository{}
}

// A RepositoryInt interface method
//
// Inserts a record in the user table
func (repo repository) Create(ctx *gofr.Context, user *model.User) (*model.User, error) {
	result, err := ctx.SQL.ExecContext(ctx, "INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	row := ctx.SQL.QueryRow("SELECT * FROM user WHERE Id = ?", id)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("id %d: no such user", id)
		}
		return user, fmt.Errorf("id %d: %v", id, err)
	}

	return user, nil
}

// A RepositoryInt interface method
func (s repository) GetByID(ctx *gofr.Context, id int64) (*model.User, error) {
	panic("unimplemented")
}

// A RepositoryInt interface method
func (s repository) GetAll(ctx *gofr.Context) ([]model.User, error) {
	panic("unimplemented")
}

// A RepositoryInt interface method
func (s repository) Update(ctx *gofr.Context, order *model.User) (*model.User, error) {
	panic("unimplemented")
}

// A RepositoryInt interface method
func (s repository) Delete(ctx *gofr.Context, id int64) error {
	panic("unimplemented")
}

// func (s repository) GetAll(ctx *gofr.Context) ([]model.User, error) {
// 	rows, err := ctx.SQL.QueryContext(ctx, "SELECT id, cust_id, products, status FROM orders WHERE deleted_at IS NULL")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var orders []model.User

// 	for rows.Next() {
// 		var o model.User

// 		err = rows.Scan(&o.ID, &o.CustomerID, pq.Array(&o.Products), &o.Status)
// 		if err != nil {
// 			return nil, err
// 		}

// 		orders = append(orders, o)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return orders, nil
// }

// func (s repository) GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Order, error) {
// 	var order model.OrUserder

// 	err := ctx.SQL.QueryRowContext(ctx, "SELECT id, cust_id, products, status FROM orders WHERE id=$1 and deleted_at IS NULL", id).
// 		Scan(&order.ID, &order.CustomerID, pq.Array(&order.Products), &order.Status)

// 	switch {
// 	case err == sql.ErrNoRows:
// 		return nil, err
// 	case err != nil:
// 		return nil, err
// 	}

// 	return &order, nil
// }

// func (s repository) Update(ctx *gofr.Context, order *model.User) (*model.User, error) {
// 	updatedAt := time.Now().UTC().Format(time.RFC3339)

// 	res, err := ctx.SQL.ExecContext(ctx, "UPDATE orders SET cust_id=$1, products=$2, status=$3, updated_at=$4 WHERE id=$5 and deleted_at IS NULL",
// 		order.CustomerID, pq.Array(order.Products), order.Status, updatedAt, order.ID)
// 	if err != nil {
// 		return nil, errors.New("DB error")
// 	}

// 	rowsAffected, _ := res.RowsAffected()

// 	if rowsAffected == 0 {
// 		return nil, errors.New("entity not found")
// 	}

// 	return order, nil
// }

// func (s repository) Delete(ctx *gofr.Context, id uuid.UUID) error {
// 	deletedAt := time.Now().UTC().Format(time.RFC3339)
// 	updatedAt := deletedAt

// 	res, err := ctx.SQL.ExecContext(ctx, "UPDATE orders SET deleted_at=$1, updated_at=$2 WHERE id=$3 and deleted_at IS NULL", deletedAt, updatedAt, id)
// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected, _ := res.RowsAffected()

// 	if rowsAffected == 0 {
// 		return err
// 	}

// 	return nil
// }
