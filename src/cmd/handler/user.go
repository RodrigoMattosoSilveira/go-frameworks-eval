package handler

import (
	"errors"
	"strconv"

	"gofr.dev/pkg/gofr"

	"madronetek.com/go-frameworks-eval/cmd/model"
	"madronetek.com/go-frameworks-eval/cmd/service"
)

type handler struct {
	service service.ServiceUserInt
}

// New - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func New(s  service.ServiceUserInt) handler {
	return handler{service: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var user model.User

	err := ctx.Bind(&user)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	resp, err := h.service.Create(ctx, &user)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	_id := ctx.PathParam("id")
	if _id == "" {
		return nil, errors.New("missing param: ID")
	}

	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	var user model.User

	_id := ctx.PathParam("id")
	if _id == "" {
		return nil, errors.New("missing param: ID")
	}

	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		return nil, err
	}

	user.Id = id

	err = ctx.Bind(&user)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	resp, err := h.service.Update(ctx, &user)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	_id := ctx.PathParam("id")
	if _id == "" {
		return nil, errors.New("missing param: ID")
	}

	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		return nil, err
	}
	
	err = h.service.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return id, nil
}