package service

import (
	"gofr.dev/pkg/gofr"

	"madronetek.com/go-frameworks-eval/cmd/model"
)

type service struct {
	store ServiceUserInt
}

// New - is a factory function to inject store in service.
func New(s ServiceUserInt)  ServiceUserInt {
	return service{store: s}
}

func (s service) Create(ctx *gofr.Context, order *model.User) (*model.User, error) {
	return s.store.Create(ctx, order)
}

func (s service) GetAll(ctx *gofr.Context) ([]model.User, error) {
	return s.store.GetAll(ctx)
}

func (s service) GetByID(ctx *gofr.Context, id int64) (*model.User, error) {
	return s.store.GetByID(ctx, id)
}

func (s service) Update(ctx *gofr.Context, order *model.User) (*model.User, error) {
	return s.store.Update(ctx, order)
}

func (s service) Delete(ctx *gofr.Context, id int64) error {
	return s.store.Delete(ctx, id)
}