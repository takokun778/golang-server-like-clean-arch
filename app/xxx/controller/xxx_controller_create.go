package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxxx "xxx/proto/xxx"
)

func (c *xxxController) Create(ctx context.Context, req *pbXxxx.CreateRequest) (*pbXxxx.CreateResponse, error) {
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

	output, err := c.xxxUsecaseCreate.Handle(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxxx.CreateResponse{
		Xxx: c.translateToProto(output.Xxx),
	}, nil
}
