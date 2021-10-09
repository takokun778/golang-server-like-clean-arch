package controller

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"
	"context"
)

func (c *hogeController) Update(ctx context.Context, req *pbHoge.UpdateRequest) (*pbHoge.UpdateResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := hoge.UsecaseUpdateInput{
		Id:     id,
		Name:   hoge.Name(req.Name),
		Number: hoge.Number(req.Number),
	}

	output, appErr := c.hogeUsecase.Update(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbHoge.UpdateResponse)
	res.Hoge = HogeTranslator{}.ToProto(output.Hoge)
	return res, nil
}
