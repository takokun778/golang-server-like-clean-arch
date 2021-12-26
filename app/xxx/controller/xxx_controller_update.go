package controller

import (
	"context"

	"xxx/app/domain/xxx"
)

type UpdatePort struct {
	Id     string
	Name   string
	Number int
}

func (c *XxxController) Update(ctx context.Context, port *UpdatePort) {
	id, err := xxx.ParseId(port.Id)

	if err != nil {
		// error
		return
	}

	name, err := xxx.NewName(port.Name)

	if err != nil {
		// error
		return
	}

	number, err := xxx.NewNumber(int(port.Number))

	if err != nil {
		// error
		return
	}

	dto := &xxx.UsecaseUpdateDto{
		Id:     id,
		Name:   name,
		Number: number,
	}

	c.xxxUsecase.Update(ctx, dto)
}
