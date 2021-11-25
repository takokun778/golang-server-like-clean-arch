package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

type xxxInteractorReadAll struct {
	xxxRepository xxx.Repository
}

func NewXxxInteractorReadAll(xxxRepository xxx.Repository) xxx.UsecaseReadAll {
	interactor := new(xxxInteractorReadAll)
	interactor.xxxRepository = xxxRepository
	return interactor
}

func (u *xxxInteractorReadAll) Handle(ctx context.Context, input *xxx.UsecaseReadAllInput) (*xxx.UsecaseReadAllOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.FindAll(timeOutCtx, &xxx.RepositoryFindAllItem{})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseReadAllOutput{
		Xxxs: result,
	}, nil
}
