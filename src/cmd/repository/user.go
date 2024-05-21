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
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("id %d: no such user", id)
		}
		return user, fmt.Errorf("id %d: %v", id, err)
	}

	return user, nil
}

// A RepositoryInt interface method
func (repo repository) GetByID(ctx *gofr.Context, id int64) (*model.User, error) {

	var user model.User;
	
	row := ctx.SQL.QueryRow("SELECT * FROM user WHERE Id = ?", id)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active); err != nil {
		if err == sql.ErrNoRows {
			return &user, fmt.Errorf("id %d: no such user", id)
		}
		return &user, fmt.Errorf("id %d: %v", id, err)
	}

	return &user, nil

	// panic("unimplemented")
}

// A RepositoryInt interface method
func (s repository) GetAll(ctx *gofr.Context) ([]model.User, error) {
	rows, err := ctx.SQL.QueryContext(ctx, "SELECT * FROM user WHERE active LIKE 'yes'")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User


		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
	// panic("unimplemented")
}

// A RepositoryInt interface method
// 
//  Since this is a framework evaluation I'll not optimize to update only the attributes that changed
// TODO figure out a way to update only the attributes that changed
// 
func (s repository) Update(ctx *gofr.Context, newUser *model.User) (*model.User, error) {
	var user model.User

	_, err := ctx.SQL.ExecContext(ctx,
		"UPDATE user SET name = ?, email = ?, password = ? WHERE id = ? and active LIKE ?", 
		newUser.Name, newUser.Email, newUser.Password, newUser.Id, "yes")
	if err != nil {
		// return nil, fmt.Errorf(`Update: unable to update user id: %d`, newUser.Id)
		return nil, fmt.Errorf(err.Error())
	}

	row := ctx.SQL.QueryRow("SELECT * FROM user WHERE Id = ? and active LIKE ?", newUser.Id, "yes")
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Update; unable to find updated user id %d", newUser.Id)
		}
		return nil, fmt.Errorf("id %d: %v", newUser.Id, err)
	}

	return &user, nil
}

// A RepositoryInt interface method
func (s repository) Delete(ctx *gofr.Context, id int64) (*model.User, error) {

	var user model.User;
	
	row := ctx.SQL.QueryRow("SELECT * FROM user WHERE Id = ? and active LIKE ?", id, "yes")
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Delete: user %d not found", id)
		}
		return nil, err
	}

	_, err := ctx.SQL.Exec("UPDATE user SET active = ? WHERE id = ?", "no", id)
	if err != nil {
		return nil, err
	}

	row = ctx.SQL.QueryRow("SELECT * FROM user WHERE Id = ? and active LIKE ?", id, "no")
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Active); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("DELETE; unable to find deleted user id %d", id)
		}
		return nil, fmt.Errorf("id %d: %v", id, err)
	}

	return &user, nil

}