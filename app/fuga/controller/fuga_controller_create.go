package controller

import (
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"
	"context"
)

func (c *fugaController) Create(ctx context.Context, req *pbFuga.CreateRequest) (*pbFuga.CreateResponse, error) {
	input := fuga.UsecaseCreateInput{
		Name:   fuga.Name(req.Name),
		Number: fuga.Number(req.Number),
	}

	output, err := c.fugaUsecase.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	res := new(pbFuga.CreateResponse)
	res.Fuga = FugaTranslator{}.ToProto(output.Fuga)
	return res, nil
}
