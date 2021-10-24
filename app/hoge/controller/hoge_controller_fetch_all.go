package controller

import (
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"
	"context"
)

func (c *hogeController) FetchAll(ctx context.Context, req *pbHoge.FetchAllRequest) (*pbHoge.FetchAllResponse, error) {
	input := hoge.UsecaseFetchAllInput{}

	output, err := c.hogeUsecase.FetchAll(ctx, input)
	if err != nil {
		return nil, err
	}

	res := new(pbHoge.FetchAllResponse)
	res.Hoges = HogeTranslator{}.ToProtoList(output.HogeList)
	return res, nil
}
