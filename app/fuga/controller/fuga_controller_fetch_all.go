package controller

import (
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"
	"context"
)

func (c *fugaController) FetchAll(ctx context.Context, req *pbFuga.FetchAllRequest) (*pbFuga.FetchAllResponse, error) {
	input := fuga.UsecaseFetchAllInput{}

	output, err := c.fugaUsecase.FetchAll(ctx, input)
	if err != nil {
		return nil, err
	}

	res := new(pbFuga.FetchAllResponse)
	res.Fugas = FugaTranslator{}.ToProtoList(output.FugaList)
	return res, nil
}
