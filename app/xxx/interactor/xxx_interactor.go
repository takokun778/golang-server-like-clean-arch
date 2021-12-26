package interactor

import (
	"xxx/app/domain/xxx"
	"xxx/app/xxx/presenter"
)

type xxxInteractor struct {
	xxxRepository xxx.Repository
	xxxPresenter  presenter.XxxPresenter
}

func NewXxxInteractor(xxxRepository xxx.Repository, xxxPresentor presenter.XxxPresenter) xxx.Usecase {
	interactor := new(xxxInteractor)
	interactor.xxxRepository = xxxRepository
	interactor.xxxPresenter = xxxPresentor
	return interactor
}
