package controller

import (
	"xxx/app/common/controller"
	"xxx/app/domain/xxx"
	pbXxxx "xxx/proto/xxx"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type xxxController struct {
	pbXxxx.UnimplementedXxxServiceServer
	*controller.Translator
	xxxUsecaseCreate  xxx.UsecaseCreate
	xxxUsecaseRead    xxx.UsecaseRead
	xxxUsecaseReadAll xxx.UsecaseReadAll
	xxxUsecaseUpdate  xxx.UsecaseUpdate
	xxxUsecaseDelete  xxx.UsecaseDelete
}

func NewXxxController(
	xxxUsecaseCreate xxx.UsecaseCreate,
	xxxUsecaseRead xxx.UsecaseRead,
	xxxUsecaseReadAll xxx.UsecaseReadAll,
	xxxUsecaseUpdate xxx.UsecaseUpdate,
	xxxUsecaseDelete xxx.UsecaseDelete,
) *xxxController {
	res := new(xxxController)
	res.Translator = controller.SetTranslator()
	res.xxxUsecaseCreate = xxxUsecaseCreate
	res.xxxUsecaseRead = xxxUsecaseRead
	res.xxxUsecaseReadAll = xxxUsecaseReadAll
	res.xxxUsecaseUpdate = xxxUsecaseUpdate
	res.xxxUsecaseDelete = xxxUsecaseDelete
	return res
}

func (*xxxController) translateToProto(xxx xxx.Props) *pbXxxx.Xxx {
	proto := new(pbXxxx.Xxx)
	proto.Id = xxx.Id().Value().String()
	proto.Name = xxx.Name().Value()
	proto.Number = int32(xxx.Number().Value())
	proto.CreatedAt = timestamppb.New(xxx.CreatedAt().Value())
	proto.UpdatedAt = timestamppb.New(xxx.UpdatedAt().Value())
	return proto
}

func (c *xxxController) translateToProtos(xxxs []xxx.Props) []*pbXxxx.Xxx {
	proto := make([]*pbXxxx.Xxx, 0)

	for _, xxx := range xxxs {
		proto = append(proto, c.translateToProto(xxx))
	}

	return proto
}
