package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

type xxxInteractorDelete struct {
	xxxRepository xxx.Repository
}

func NewXxxInteractorDelete(xxxRepository xxx.Repository) xxx.UsecaseDelete {
	interactor := new(xxxInteractorDelete)
	interactor.xxxRepository = xxxRepository
	return interactor
}

func (u *xxxInteractorDelete) Handle(ctx context.Context, input *xxx.UsecaseDeleteInput) (*xxx.UsecaseDeleteOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: input.Id})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	err = u.xxxRepository.Delete(timeOutCtx, &xxx.RepositoryDeleteItem{Id: input.Id})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseDeleteOutput{
		Xxx: result,
	}, nil
}
