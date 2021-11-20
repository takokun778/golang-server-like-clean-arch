package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxxx "xxx/proto/xxx"
)

func (c *xxxController) ReadAll(ctx context.Context, req *pbXxxx.ReadAllRequest) (*pbXxxx.ReadAllResponse, error) {
	input := &xxx.UsecaseReadAllInput{}

	output, err := c.xxxUsecase.ReadAll(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxxx.ReadAllResponse{
		Xxxs: c.translateToProtos(output.Xxxs),
	}, nil
}
