package service

import (
	"gofr.dev/pkg/gofr"

	"madronetek.com/go-frameworks-eval/cmd/model"
)

type ServiceUserInt interface {
	Create(ctx *gofr.Context, user *model.User) (*model.User, error)
	GetAll(ctx *gofr.Context) ([]model.User, error)
	GetByID(ctx *gofr.Context, id int64) (*model.User, error)
	Update(ctx *gofr.Context, user *model.User) (*model.User, error)
	Delete(ctx *gofr.Context, id int64) error
}
