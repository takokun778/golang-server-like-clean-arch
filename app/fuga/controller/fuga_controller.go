package controller

import (
	common "clean/app/domain/common"
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"
	"context"
)

type fugaController struct {
	fugaUsecase fuga.Usecase
	pbFuga.UnimplementedFugaServiceServer
}

func NewFugaController(fugaUsecase fuga.Usecase) *fugaController {
	res := new(fugaController)
	res.fugaUsecase = fugaUsecase
	return res
}

func (c *fugaController) Create(ctx context.Context, req *pbFuga.CreateRequest) (*pbFuga.CreateResponse, error) {
	input := fuga.CreateUsecaseInput{
		Name:   fuga.NewName(req.Name),
		Number: fuga.NewNumber(req.Number),
	}

	output, appErr := c.fugaUsecase.Create(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbFuga.CreateResponse)
	res.Fuga = FugaTranslator{}.ToProto(output.Fuga)
	return res, nil
}

func (c *fugaController) Fetch(ctx context.Context, req *pbFuga.FetchRequest) (*pbFuga.FetchResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := fuga.FetchUsecaseInput{
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

func (c *fugaController) FetchAll(ctx context.Context, req *pbFuga.FetchAllRequest) (*pbFuga.FetchAllResponse, error) {
	input := fuga.FetchAllUsecaseInput{}

	output, appErr := c.fugaUsecase.FetchAll(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbFuga.FetchAllResponse)
	res.Fugas = FugaTranslator{}.ToProtoList(output.FugaList)
	return res, nil
}

func (c *fugaController) Update(ctx context.Context, req *pbFuga.UpdateRequest) (*pbFuga.UpdateResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := fuga.UpdateUsecaseInput{
		Id:     id,
		Name:   fuga.NewName(req.Name),
		Number: fuga.NewNumber(req.Number),
	}

	output, appErr := c.fugaUsecase.Update(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbFuga.UpdateResponse)
	res.Fuga = FugaTranslator{}.ToProto(output.Fuga)
	return res, nil
}

func (c *fugaController) Delete(ctx context.Context, req *pbFuga.DeleteRequest) (*pbFuga.DeleteResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := fuga.DeleteUsecaseInput{
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
