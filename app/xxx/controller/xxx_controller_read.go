package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxxx "xxx/proto/xxx"
)

func (c *xxxController) Read(ctx context.Context, req *pbXxxx.ReadRequest) (*pbXxxx.ReadResponse, error) {
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

	return &pbXxxx.ReadResponse{
		Xxx: c.translateToProto(output.Xxx),
	}, nil
}
