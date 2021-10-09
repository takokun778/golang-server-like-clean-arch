package controller

import (
	common "clean/app/domain/common"
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"
	"context"
)

func (c *fugaController) Update(ctx context.Context, req *pbFuga.UpdateRequest) (*pbFuga.UpdateResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := fuga.UsecaseUpdateInput{
		Id:     id,
		Name:   fuga.Name(req.Name),
		Number: fuga.Number(req.Number),
	}

	output, appErr := c.fugaUsecase.Update(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbFuga.UpdateResponse)
	res.Fuga = FugaTranslator{}.ToProto(output.Fuga)
	return res, nil
}
