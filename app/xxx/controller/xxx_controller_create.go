package controller

import (
	"context"

	"xxx/app/domain/xxx"
)

type CreatePort struct {
	Name   string
	Number int
}

func (c *XxxController) Create(ctx context.Context, port *CreatePort) {
	name, err := xxx.NewName(port.Name)

	if err != nil {
		// error
		return
	}

	number, err := xxx.NewNumber(port.Number)

	if err != nil {
		// error
		return
	}

	dto := &xxx.UsecaseCreateDto{
		Name:   name,
		Number: number,
	}

	c.xxxUsecase.Create(ctx, dto)
}
