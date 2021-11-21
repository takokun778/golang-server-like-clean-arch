package controller

import (
	"xxx/app/common/controller"
	"xxx/app/domain/xxx"
	pbXxxx "xxx/proto/xxx"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type xxxController struct {
	*controller.Translator
	xxxUsecase xxx.Usecase
	pbXxxx.UnimplementedXxxServiceServer
}

func NewXxxController(xxxUsecase xxx.Usecase) *xxxController {
	res := new(xxxController)
	res.Translator = controller.SetTranslator()
	res.xxxUsecase = xxxUsecase
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
