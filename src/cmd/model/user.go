package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
) 

type User struct {
  Id        int64     `json:"id"`
  Name      string    `json:"name"`
  Email     string    `json:"email"`
  Password  string    `json:"password"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt time.Time `json:"deleted_at"`
}

func (u User) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *User) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &u)
}