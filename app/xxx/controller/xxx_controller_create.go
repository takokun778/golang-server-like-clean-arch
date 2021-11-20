package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxxx "xxx/proto/xxx"
)

func (c *xxxController) Create(ctx context.Context, req *pbXxxx.CreateRequest) (*pbXxxx.CreateResponse, error) {
	input := &xxx.UsecaseCreateInput{
		Name:   xxx.Name(req.Name),
		Number: xxx.Number(req.Number),
	}

	output, err := c.xxxUsecase.Create(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxxx.CreateResponse{
		Xxx: c.translateToProto(output.Xxx),
	}, nil
}
