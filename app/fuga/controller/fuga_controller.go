package controller

import (
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"
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
