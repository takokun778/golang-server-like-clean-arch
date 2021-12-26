package xxx

import (
	"context"

	"xxx/app/domain/xxx"
	"xxx/app/infra/proto"
	"xxx/app/xxx/presenter"
	pbXxx "xxx/proto/xxx"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type xxxAdapter struct{}

func NewXxxAdapter() *xxxAdapter {
	adapter := new(xxxAdapter)
	return adapter
}

func (s *xxxAdapter) Create(ctx context.Context, dto *presenter.PresenterCreateDto, err error) {
	if err != nil {
		proto.SetError(ctx, err)
		return
	}

	res := &pbXxx.CreateResponse{
		Xxx: translate(dto.Xxx),
	}

	setCreateResponse(ctx, res)
}

func (s *xxxAdapter) Update(ctx context.Context, dto *presenter.PresenterUpdateDto, err error) {
	if err != nil {
		proto.SetError(ctx, err)
		return
	}

	res := &pbXxx.UpdateResponse{
		Xxx: translate(dto.Xxx),
	}

	setUpdateResponse(ctx, res)
}

func (s *xxxAdapter) Read(ctx context.Context, dto *presenter.PresenterReadDto, err error) {
	if err != nil {
		proto.SetError(ctx, err)
		return
	}

	res := &pbXxx.ReadResponse{
		Xxx: translate(dto.Xxx),
	}

	setReadResponse(ctx, res)
}

func (s *xxxAdapter) ReadAll(ctx context.Context, dto *presenter.PresenterReadAllDto, err error) {
	if err != nil {
		proto.SetError(ctx, err)
		return
	}

	res := &pbXxx.ReadAllResponse{
		Xxxs: translateArray(dto.Xxxs),
	}

	setReadAllResponse(ctx, res)
}

func (s *xxxAdapter) Delete(ctx context.Context, dto *presenter.PresenterDeleteDto, err error) {
	if err != nil {
		proto.SetError(ctx, err)
		return
	}

	res := &pbXxx.DeleteResponse{
		Xxx: translate(dto.Xxx),
	}

	setDeleteResponse(ctx, res)
}

func translate(xxx xxx.Props) *pbXxx.Xxx {
	proto := new(pbXxx.Xxx)
	proto.Id = xxx.Id().Value().String()
	proto.Name = xxx.Name().Value()
	proto.Number = int32(xxx.Number().Value())
	proto.CreatedAt = timestamppb.New(xxx.CreatedAt().Value())
	proto.UpdatedAt = timestamppb.New(xxx.UpdatedAt().Value())
	return proto
}

func translateArray(xxxs []xxx.Props) []*pbXxx.Xxx {
	proto := make([]*pbXxx.Xxx, 0)

	for _, xxx := range xxxs {
		proto = append(proto, translate(xxx))
	}

	return proto
}
