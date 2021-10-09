package controller

import (
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"
	"context"
)

func (c *hogeController) FetchAll(ctx context.Context, req *pbHoge.FetchAllRequest) (*pbHoge.FetchAllResponse, error) {
	input := hoge.UsecaseFetchAllInput{}

	output, appErr := c.hogeUsecase.FetchAll(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbHoge.FetchAllResponse)
	res.Hoges = HogeTranslator{}.ToProtoList(output.HogeList)
	return res, nil
}
