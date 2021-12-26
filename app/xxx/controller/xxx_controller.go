package controller

import (
	"xxx/app/domain/xxx"
)

type XxxController struct {
	xxxUsecase xxx.Usecase
}

func NewXxxController(xxxUsecase xxx.Usecase) *XxxController {
	res := new(XxxController)
	res.xxxUsecase = xxxUsecase
	return res
}
