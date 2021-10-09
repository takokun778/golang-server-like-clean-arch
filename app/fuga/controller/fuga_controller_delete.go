package controller

import (
	common "clean/app/domain/common"
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"
	"context"
)

func (c *fugaController) Delete(ctx context.Context, req *pbFuga.DeleteRequest) (*pbFuga.DeleteResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := fuga.UsecaseDeleteInput{
		Id: id,
	}

	output, appErr := c.fugaUsecase.Delete(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbFuga.DeleteResponse)
	res.Fuga = FugaTranslator{}.ToProto(output.Fuga)
	return res, nil
}
