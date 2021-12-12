package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxx "xxx/proto/xxx"
)

func (c *xxxController) Read(ctx context.Context, req *pbXxx.ReadRequest) (*pbXxx.ReadResponse, error) {
	id, err := xxx.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := &xxx.UsecaseReadInput{
		Id: id,
	}

	output, err := c.xxxUsecase.Read(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxx.ReadResponse{
		Xxx: c.Proto.Translate(output.Xxx),
	}, nil
}
