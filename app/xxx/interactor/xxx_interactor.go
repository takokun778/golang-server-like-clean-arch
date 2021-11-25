package interactor

import (
	"xxx/app/domain/xxx"
)

type xxxInteractor struct {
	xxxRepository xxx.Repository
}

func NewXxxInteractor(xxxRepository xxx.Repository) xxx.Usecase {
	interactor := new(xxxInteractor)
	interactor.xxxRepository = xxxRepository
	return interactor
}
