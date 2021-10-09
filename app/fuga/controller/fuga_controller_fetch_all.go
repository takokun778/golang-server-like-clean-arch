package controller

import (
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"
	"context"
)

func (c *fugaController) FetchAll(ctx context.Context, req *pbFuga.FetchAllRequest) (*pbFuga.FetchAllResponse, error) {
	input := fuga.UsecaseFetchAllInput{}

	output, appErr := c.fugaUsecase.FetchAll(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbFuga.FetchAllResponse)
	res.Fugas = FugaTranslator{}.ToProtoList(output.FugaList)
	return res, nil
}
