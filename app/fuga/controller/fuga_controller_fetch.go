package controller

import (
	common "clean/app/domain/common"
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"
	"context"
)

func (c *fugaController) Fetch(ctx context.Context, req *pbFuga.FetchRequest) (*pbFuga.FetchResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := fuga.UsecaseFetchInput{
		Id: id,
	}

	output, appErr := c.fugaUsecase.Fetch(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbFuga.FetchResponse)
	res.Fuga = FugaTranslator{}.ToProto(output.Fuga)
	return res, nil
}
