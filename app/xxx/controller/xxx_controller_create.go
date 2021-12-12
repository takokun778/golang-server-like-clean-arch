package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxx "xxx/proto/xxx"
)

func (c *xxxController) Create(ctx context.Context, req *pbXxx.CreateRequest) (*pbXxx.CreateResponse, error) {
	name, err := xxx.NewName(req.Name)

	if err != nil {
		return nil, err
	}

	number, err := xxx.NewNumber(int(req.Number))

	if err != nil {
		return nil, err
	}

	input := &xxx.UsecaseCreateInput{
		Name:   name,
		Number: number,
	}

	output, err := c.xxxUsecase.Create(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxx.CreateResponse{
		Xxx: c.Proto.Translate(output.Xxx),
	}, nil
}
