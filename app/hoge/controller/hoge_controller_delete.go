package controller

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"
	"context"
)

func (c *hogeController) Delete(ctx context.Context, req *pbHoge.DeleteRequest) (*pbHoge.DeleteResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := hoge.UsecaseDeleteInput{
		Id: id,
	}

	output, err := c.hogeUsecase.Delete(ctx, input)
	if err != nil {
		return nil, err
	}

	res := new(pbHoge.DeleteResponse)
	res.Hoge = HogeTranslator{}.ToProto(output.Hoge)
	return res, nil
}
