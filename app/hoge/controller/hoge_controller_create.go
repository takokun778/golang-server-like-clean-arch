package controller

import (
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"
	"context"
)

func (c *hogeController) Create(ctx context.Context, req *pbHoge.CreateRequest) (*pbHoge.CreateResponse, error) {
	input := hoge.UsecaseCreateInput{
		Name:   hoge.Name(req.Name),
		Number: hoge.Number(req.Number),
	}

	output, err := c.hogeUsecase.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	res := new(pbHoge.CreateResponse)
	res.Hoge = HogeTranslator{}.ToProto(output.Hoge)
	return res, nil
}
