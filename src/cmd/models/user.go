package models

import (
	"errors"
	"fmt"
	"log/slog"
	"net/mail"
	"time"

	"gofr.dev/pkg/gofr"
) 

type User struct {
  Id        int64     `json:"id"`
  Name      string    `json:"name"`
  Email     string    `json:"email"`
  Password  string    `json:"password"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) Create(ctx *gofr.Context) (interface{}, error) {
	err := ctx.Bind(&user)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	if !validEmail(user.Email) {
		errorString := fmt.Sprintf("Invalid email: %s", user.Email);
		return nil, errors.New(errorString)
	}
	name := ctx.Param("name")
	slog.Info(fmt.Sprintf("Create, name = %s ", name))


	return "user Create called", nil

}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
    return err == nil
}
