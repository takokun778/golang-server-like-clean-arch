package usecase

import (
	"clean/app/domain/fuga"
)

type fugaUsecase struct {
	fugaRepository fuga.Repository
}

func NewFugaUsecase(fugaRepository fuga.Repository) fuga.Usecase {
	usecase := new(fugaUsecase)
	usecase.fugaRepository = fugaRepository
	return usecase
}
