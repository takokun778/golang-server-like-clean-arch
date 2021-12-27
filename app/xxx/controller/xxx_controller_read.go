package controller

import (
	"context"

	"xxx/app/domain/xxx"
)

type ReadPort struct {
	Id string
}

func (c *XxxController) Read(ctx context.Context, port *ReadPort) {
	id, err := xxx.ParseId(port.Id)

	if err != nil {
		c.xxxUsecase.Read(ctx, nil, err)
		return
	}

	dto := &xxx.UsecaseReadDto{
		Id: id,
	}

	c.xxxUsecase.Read(ctx, dto, nil)
}
