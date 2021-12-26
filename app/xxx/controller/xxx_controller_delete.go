package controller

import (
	"context"

	"xxx/app/domain/xxx"
)

type DeletePort struct {
	Id string
}

func (c *XxxController) Delete(ctx context.Context, port *DeletePort) {
	id, err := xxx.ParseId(port.Id)

	if err != nil {
		// error
		return
	}

	dto := &xxx.UsecaseDeleteDto{
		Id: id,
	}

	c.xxxUsecase.Delete(ctx, dto)
}
