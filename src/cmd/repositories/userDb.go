package repositories

import (
	"database/sql"
	"fmt"

	"madronetek.com/go-frameworks-eval/cmd/models"
	"madronetek.com/go-frameworks-eval/cmd/storage"
)

func CreateUserDb(user models.User) (models.User, error) {
	db := storage.GetDB()

	// MySQL
	// 
	result, err := db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	row := db.QueryRow("SELECT * FROM user WHERE Id = ?", id)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("id %d: no such user", id)
		}
		return user, fmt.Errorf("id %d: %v", id, err)
	}

	return user, nil
}

func UpdateUserDb(user models.User, id int64) (models.User, error) {
	db := storage.GetDB()

	// MySQL
	sqlStatement := `
	UPDATE user 
	SET 
		name = ?, 
		email = ?, 
		password = ?
	WHERE id = ?`
	_, err := db.Exec(sqlStatement, user.Name, user.Email, user.Password, id)
	if err != nil {
		return user, err
	}


	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("id %d: no such user", user.Id)
		}
	
		return user, fmt.Errorf("id %d: %v", user.Id, err)
	}

	return user, nil
}
