package controller

import (
	"xxx/app/domain/xxx"
	"xxx/app/infra/proto"
	pbXxx "xxx/proto/xxx"
)

type xxxController struct {
	*proto.Proto
	xxxUsecase xxx.Usecase
	pbXxx.UnimplementedXxxServiceServer
}

func NewXxxController(xxxUsecase xxx.Usecase) *xxxController {
	res := new(xxxController)
	res.Proto = proto.NewProto()
	res.xxxUsecase = xxxUsecase
	return res
}
