package usecase

import (
	"xxx/app/domain/xxx"
)

type xxxUsecase struct {
	xxxRepository xxx.Repository
}

func NewXxxUsecase(xxxRepository xxx.Repository) xxx.Usecase {
	usecase := new(xxxUsecase)
	usecase.xxxRepository = xxxRepository
	return usecase
}
