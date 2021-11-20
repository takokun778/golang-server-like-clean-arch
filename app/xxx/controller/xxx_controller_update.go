package controller

import (
	"context"

	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
	pbXxx "xxx/proto/xxx"
)

func (c *xxxController) Update(ctx context.Context, req *pbXxx.UpdateRequest) (*pbXxx.UpdateResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := &xxx.UsecaseUpdateInput{
		Id:     id,
		Name:   xxx.Name(req.Name),
		Number: xxx.Number(req.Number),
	}

	output, err := c.xxxUsecase.Update(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxx.UpdateResponse{
		Xxx: c.translateToProto(output.Xxx),
	}, nil
}
