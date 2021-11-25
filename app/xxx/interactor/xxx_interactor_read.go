package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

type xxxInteractorRead struct {
	xxxRepository xxx.Repository
}

func NewXxxInteractorRead(xxxRepository xxx.Repository) xxx.UsecaseRead {
	interactor := new(xxxInteractorRead)
	interactor.xxxRepository = xxxRepository
	return interactor
}

func (u *xxxInteractorRead) Handle(ctx context.Context, input *xxx.UsecaseReadInput) (*xxx.UsecaseReadOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: input.Id})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseReadOutput{
		Xxx: result,
	}, nil
}
