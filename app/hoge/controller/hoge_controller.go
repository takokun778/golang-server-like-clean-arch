package controller

import (
	common "clean/app/domain/common"
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"
	"context"
)

type hogeController struct {
	hogeUsecase hoge.Usecase
	pbHoge.UnimplementedHogeServiceServer
}

func NewHogeController(hogeUsecase hoge.Usecase) *hogeController {
	res := new(hogeController)
	res.hogeUsecase = hogeUsecase
	return res
}

func (c *hogeController) Create(ctx context.Context, req *pbHoge.CreateRequest) (*pbHoge.CreateResponse, error) {
	input := hoge.CreateUsecaseInput{
		Name:   hoge.NewName(req.Name),
		Number: hoge.NewNumber(req.Number),
	}

	output, appErr := c.hogeUsecase.Create(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbHoge.CreateResponse)
	res.Hoge = HogeTranslator{}.ToProto(output.Hoge)
	return res, nil
}

func (c *hogeController) Fetch(ctx context.Context, req *pbHoge.FetchRequest) (*pbHoge.FetchResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := hoge.FetchUsecaseInput{
		Id: id,
	}

	output, appErr := c.hogeUsecase.Fetch(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbHoge.FetchResponse)
	res.Hoge = HogeTranslator{}.ToProto(output.Hoge)
	return res, nil
}

func (c *hogeController) FetchAll(ctx context.Context, req *pbHoge.FetchAllRequest) (*pbHoge.FetchAllResponse, error) {
	input := hoge.FetchAllUsecaseInput{}

	output, appErr := c.hogeUsecase.FetchAll(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbHoge.FetchAllResponse)
	res.Hoges = HogeTranslator{}.ToProtoList(output.HogeList)
	return res, nil
}

func (c *hogeController) Update(ctx context.Context, req *pbHoge.UpdateRequest) (*pbHoge.UpdateResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := hoge.UpdateUsecaseInput{
		Id:     id,
		Name:   hoge.NewName(req.Name),
		Number: hoge.NewNumber(req.Number),
	}

	output, appErr := c.hogeUsecase.Update(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbHoge.UpdateResponse)
	res.Hoge = HogeTranslator{}.ToProto(output.Hoge)
	return res, nil
}

func (c *hogeController) Delete(ctx context.Context, req *pbHoge.DeleteRequest) (*pbHoge.DeleteResponse, error) {
	id, err := common.ParseId(req.Id)

	if err != nil {
		return nil, err
	}

	input := hoge.DeleteUsecaseInput{
		Id: id,
	}

	output, appErr := c.hogeUsecase.Delete(ctx, input)
	if appErr != nil {
		return nil, appErr
	}

	res := new(pbHoge.DeleteResponse)
	res.Hoge = HogeTranslator{}.ToProto(output.Hoge)
	return res, nil
}
