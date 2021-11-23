package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxx "xxx/proto/xxx"
)

func (c *xxxController) Update(ctx context.Context, req *pbXxx.UpdateRequest) (*pbXxx.UpdateResponse, error) {
	id, err := xxx.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	name, err := xxx.NewName(req.Name)

	if err != nil {
		return nil, err
	}

	number, err := xxx.NewNumber(int(req.Number))

	if err != nil {
		return nil, err
	}

	input := &xxx.UsecaseUpdateInput{
		Id:     id,
		Name:   name,
		Number: number,
	}

	output, err := c.xxxUsecase.Update(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxx.UpdateResponse{
		Xxx: c.translateToProto(output.Xxx),
	}, nil
}
