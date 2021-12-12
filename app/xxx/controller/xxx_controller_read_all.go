package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxx "xxx/proto/xxx"
)

func (c *xxxController) ReadAll(ctx context.Context, req *pbXxx.ReadAllRequest) (*pbXxx.ReadAllResponse, error) {
	input := &xxx.UsecaseReadAllInput{}

	output, err := c.xxxUsecase.ReadAll(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxx.ReadAllResponse{
		Xxxs: c.Proto.TranslateArray(output.Xxxs),
	}, nil
}
