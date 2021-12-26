package xxx

import (
	"context"

	"xxx/app/infra/proto"
	xc "xxx/app/xxx/controller"
	pbXxx "xxx/proto/xxx"
)

type xxxService struct {
	pbXxx.UnimplementedXxxServiceServer
	xxxController *xc.XxxController
}

func NewXxxService(xxxController *xc.XxxController) *xxxService {
	service := new(xxxService)
	service.xxxController = xxxController
	return service
}

func (s *xxxService) Create(ctx context.Context, req *pbXxx.CreateRequest) (*pbXxx.CreateResponse, error) {
	port := &xc.CreatePort{
		Name:   req.Name,
		Number: int(req.Number),
	}

	s.xxxController.Create(ctx, port)

	errRes, err := proto.GetError(ctx)

	if errRes != nil {
		return nil, errRes
	}

	if err != nil {
		return nil, err
	}

	res, err := getCreateResponse(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *xxxService) Update(ctx context.Context, req *pbXxx.UpdateRequest) (*pbXxx.UpdateResponse, error) {
	port := &xc.UpdatePort{
		Id:     req.Id,
		Name:   req.Name,
		Number: int(req.Number),
	}

	s.xxxController.Update(ctx, port)

	errRes, err := proto.GetError(ctx)

	if errRes != nil {
		return nil, errRes
	}

	if err != nil {
		return nil, err
	}

	res, err := getUpdateResponse(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *xxxService) Read(ctx context.Context, req *pbXxx.ReadRequest) (*pbXxx.ReadResponse, error) {
	port := &xc.ReadPort{
		Id: req.Id,
	}

	s.xxxController.Read(ctx, port)

	errRes, err := proto.GetError(ctx)

	if errRes != nil {
		return nil, errRes
	}

	if err != nil {
		return nil, err
	}

	res, err := getReadResponse(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *xxxService) ReadAll(ctx context.Context, req *pbXxx.ReadAllRequest) (*pbXxx.ReadAllResponse, error) {
	port := &xc.ReadAllPort{}

	s.xxxController.ReadAll(ctx, port)

	errRes, err := proto.GetError(ctx)

	if errRes != nil {
		return nil, errRes
	}

	if err != nil {
		return nil, err
	}

	res, err := getReadAllResponse(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *xxxService) Delete(ctx context.Context, req *pbXxx.DeleteRequest) (*pbXxx.DeleteResponse, error) {
	port := &xc.DeletePort{
		Id: req.Id,
	}

	s.xxxController.Delete(ctx, port)

	errRes, err := proto.GetError(ctx)

	if errRes != nil {
		return nil, errRes
	}

	if err != nil {
		return nil, err
	}

	res, err := getDeleteResponse(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}
