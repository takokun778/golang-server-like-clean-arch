package controller

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"
	"context"
)

func (c *hogeController) Fetch(ctx context.Context, req *pbHoge.FetchRequest) (*pbHoge.FetchResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := hoge.UsecaseFetchInput{
		Id: id,
	}

	output, err := c.hogeUsecase.Fetch(ctx, input)
	if err != nil {
		return nil, err
	}

	res := new(pbHoge.FetchResponse)
	res.Hoge = HogeTranslator{}.ToProto(output.Hoge)
	return res, nil
}
