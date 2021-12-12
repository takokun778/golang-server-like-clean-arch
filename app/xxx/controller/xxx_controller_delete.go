package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxx "xxx/proto/xxx"
)

func (c *xxxController) Delete(ctx context.Context, req *pbXxx.DeleteRequest) (*pbXxx.DeleteResponse, error) {
	id, err := xxx.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := &xxx.UsecaseDeleteInput{
		Id: id,
	}

	output, err := c.xxxUsecase.Delete(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxx.DeleteResponse{
		Xxx: c.Proto.Translate(output.Xxx),
	}, nil
}
