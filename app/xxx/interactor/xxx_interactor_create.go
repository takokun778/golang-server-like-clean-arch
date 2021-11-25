package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

type xxxInteractorCreate struct {
	xxxRepository xxx.Repository
}

func NewXxxInteractorCreate(xxxRepository xxx.Repository) xxx.UsecaseCreate {
	interactor := new(xxxInteractorCreate)
	interactor.xxxRepository = xxxRepository
	return interactor
}

func (u *xxxInteractorCreate) Handle(ctx context.Context, input *xxx.UsecaseCreateInput) (*xxx.UsecaseCreateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := xxx.Create(input.Name, input.Number)

	repoInput := &xxx.RepositorySaveItem{
		Xxx: result.Props(),
	}

	_, err := u.xxxRepository.Save(timeOutCtx, repoInput)

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseCreateOutput{
		Xxx: result.Props(),
	}, nil
}
