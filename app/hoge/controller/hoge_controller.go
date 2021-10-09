package controller

import (
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"
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
