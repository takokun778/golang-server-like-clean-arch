package controller

import (
	"context"

	"xxx/app/domain/xxx"
	pbXxxx "xxx/proto/xxx"
)

func (c *xxxController) Delete(ctx context.Context, req *pbXxxx.DeleteRequest) (*pbXxxx.DeleteResponse, error) {
	id, err := xxx.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := &xxx.UsecaseDeleteInput{
		Id: id,
	}

	output, err := c.xxxUsecaseDelete.Handle(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pbXxxx.DeleteResponse{
		Xxx: c.translateToProto(output.Xxx),
	}, nil
}
