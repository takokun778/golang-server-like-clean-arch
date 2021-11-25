package interactor

import (
	"context"
	"time"

	dError "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

type xxxInteractorUpdate struct {
	xxxRepository xxx.Repository
}

func NewXxxInteractorUpdate(xxxRepository xxx.Repository) xxx.UsecaseUpdate {
	interactor := new(xxxInteractorUpdate)
	interactor.xxxRepository = xxxRepository
	return interactor
}

func (u *xxxInteractorUpdate) Handle(ctx context.Context, input *xxx.UsecaseUpdateInput) (*xxx.UsecaseUpdateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := u.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: input.Id})

	if err != nil {
		return nil, dError.NewInternalServerError(err, "")
	}

	dst := xxx.Reconstruct(src).Update(input.Name, input.Number)

	_, err = u.xxxRepository.Update(timeOutCtx, &xxx.RepositoryUpdateItem{Xxx: dst.Props()})

	if err != nil {
		return nil, dError.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseUpdateOutput{
		Xxx: dst.Props(),
	}, nil
}
