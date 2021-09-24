package usecase

import (
	"clean/app/domain/hoge"
)

type hogeUsecase struct {
	hogeRepository hoge.Repository
}

func NewHogeUsecase(hogeRepository hoge.Repository) hoge.Usecase {
	usecase := new(hogeUsecase)
	usecase.hogeRepository = hogeRepository
	return usecase
}
