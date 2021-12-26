package controller

import (
	"context"

	"xxx/app/domain/xxx"
)

type ReadAllPort struct {
}

func (c *XxxController) ReadAll(ctx context.Context, port *ReadAllPort) {
	dto := &xxx.UsecaseReadAllDto{}

	c.xxxUsecase.ReadAll(ctx, dto)
}
